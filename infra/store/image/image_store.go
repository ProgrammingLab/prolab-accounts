package imagestore

import (
	"bytes"
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"strings"
	"time"

	"github.com/minio/minio-go"
	"github.com/pkg/errors"
	"golang.org/x/image/draw"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc/grpclog"

	"github.com/ProgrammingLab/prolab-accounts/infra/store"
)

type imageStoreImpl struct {
	ctx        context.Context
	cli        *minio.Client
	bucketName string
}

// NewImageStore returns new image store
func NewImageStore(ctx context.Context, cli *minio.Client, bucket string) store.ImageStore {
	return &imageStoreImpl{
		ctx:        ctx,
		cli:        cli,
		bucketName: bucket,
	}
}

func (s *imageStoreImpl) CreateImage(img []byte) (filename string, err error) {
	name, err := generateFilename()
	if err != nil {
		return "", errors.WithStack(err)
	}
	r := bytes.NewReader(img)
	return s.createImage(name, r)
}

func (s *imageStoreImpl) DeleteImage(filename string) error {
	err := s.cli.RemoveObject(s.bucketName, filename)
	return errors.WithStack(err)
}

var (
	imageSizes = []int{
		64,
		128,
		256,
		512,
	}
)

func (s *imageStoreImpl) MigrateImages() error {
	grpclog.Infoln("image migration started")

	doneCh := make(chan struct{})
	defer close(doneCh)
	n := 0
	keys := make(map[string]struct{})
	for info := range s.cli.ListObjectsV2(s.bucketName, "", true, doneCh) {
		key := info.Key
		keys[key] = struct{}{}
	}

	grpclog.Infof("found %v images", len(keys))

	for key := range keys {
		if strings.HasSuffix(key, "px") {
			continue
		}

		migrated := true
		for _, px := range imageSizes {
			_, ok := keys[filenameWithPx(key, px)]
			migrated = migrated && ok
		}
		if migrated {
			continue
		}

		grpclog.Infof("migrating %v", key)
		err := s.migrateImage(key)
		if err != nil {
			return err
		}
		grpclog.Infof("migrated %v", key)
		n++
	}

	grpclog.Infof("migrated %v images!", n)

	return nil
}

func (s *imageStoreImpl) migrateImage(key string) error {
	i := strings.LastIndex(key, ".")
	name := key[:i]
	obj, err := s.cli.GetObjectWithContext(s.ctx, s.bucketName, key, minio.GetObjectOptions{})
	if err != nil {
		return errors.WithStack(err)
	}
	defer obj.Close()

	_, err = s.createImage(name, obj)
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (s *imageStoreImpl) createImage(name string, img io.Reader) (filename string, err error) {
	var buf bytes.Buffer
	tee := io.TeeReader(img, &buf)
	src, ext, err := image.Decode(tee)
	if err != nil {
		return "", errors.WithStack(err)
	}

	eg := errgroup.Group{}

	orgName := name + "." + ext
	eg.Go(func() error {
		opt := minio.PutObjectOptions{
			ContentType: "image/" + ext,
		}
		_, err := s.cli.PutObjectWithContext(s.ctx, s.bucketName, orgName, &buf, -1, opt)
		return errors.WithStack(err)
	})

	for _, size := range imageSizes {
		px := size
		eg.Go(func() error {
			start := time.Now()
			img := s.Resize(src, px)
			dur := time.Since(start)
			grpclog.Infof("image resize %v: %v", px, dur)

			start = time.Now()
			err := s.putImage(img, filenameWithPx(name+"."+ext, px), ext)
			if err != nil {
				return errors.WithStack(err)
			}
			dur = time.Since(start)
			grpclog.Infof("image put %v: %v", px, dur)

			return nil
		})
	}

	err = eg.Wait()
	if err != nil {
		return "", err
	}
	return orgName, nil
}

func (s *imageStoreImpl) Resize(src image.Image, size int) image.Image {
	srcW, srcH := src.Bounds().Dx(), src.Bounds().Dy()
	if srcW <= size && srcH <= size {
		return src
	}

	var (
		w int
		h int
	)
	// 長辺がsizeになるように比を変えずに縮小する
	if srcW < srcH {
		h = size
		w = srcW * size / srcH
	} else {
		w = size
		h = srcH * size / srcW
	}

	dst := image.NewRGBA(image.Rect(0, 0, w, h))
	draw.ApproxBiLinear.Scale(dst, dst.Bounds(), src, src.Bounds(), draw.Over, nil)
	return dst
}

func (s *imageStoreImpl) putImage(img image.Image, filename, ext string) error {
	var buf bytes.Buffer

	var err error
	switch ext {
	case "gif":
		err = errors.WithStack(gif.Encode(&buf, img, nil))
	case "jpeg":
		err = errors.WithStack(jpeg.Encode(&buf, img, nil))
	case "png":
		err = errors.WithStack(png.Encode(&buf, img))
	default:
		err = errors.WithStack(image.ErrFormat)
	}
	if err != nil {
		return errors.WithStack(err)
	}

	opt := minio.PutObjectOptions{
		ContentType: "image/" + ext,
	}
	_, err = s.cli.PutObjectWithContext(s.ctx, s.bucketName, filename, &buf, int64(buf.Len()), opt)
	return errors.WithStack(err)
}

func generateFilename() (string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", errors.WithStack(err)
	}

	res := base64.RawURLEncoding.EncodeToString(b)

	return string(res), nil
}

func filenameWithPx(filename string, px int) string {
	return fmt.Sprintf("%v_%vpx", filename, px)
}

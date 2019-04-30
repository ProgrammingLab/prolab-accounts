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
	r := bytes.NewReader(img)
	return s.createImage(r)
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

func (s *imageStoreImpl) createImage(img io.Reader) (filename string, err error) {
	var buf bytes.Buffer
	tee := io.TeeReader(img, &buf)
	src, ext, err := image.Decode(tee)
	if err != nil {
		return "", errors.WithStack(err)
	}
	name, err := generateFilename()
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
			img := s.resize(src, px)
			err := s.putImage(img, fmt.Sprintf("%v.%v_%vpx", name, ext, px), ext)
			return errors.WithStack(err)
		})
	}

	err = eg.Wait()
	if err != nil {
		return "", err
	}
	return orgName, nil
}

func (s *imageStoreImpl) resize(src image.Image, size int) image.Image {
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
	draw.CatmullRom.Scale(dst, dst.Bounds(), src, src.Bounds(), draw.Over, nil)
	return dst
}

func (s *imageStoreImpl) putImage(img image.Image, filename, ext string) error {
	r, w := io.Pipe()
	go func() {
		var err error
		defer func() {
			e := w.CloseWithError(err)
			grpclog.Error(e)
		}()
		switch ext {
		case "gif":
			err = errors.WithStack(gif.Encode(w, img, nil))
		case "jpeg":
			err = errors.WithStack(jpeg.Encode(w, img, nil))
		case "png":
			err = errors.WithStack(png.Encode(w, img))
		default:
			err = errors.WithStack(image.ErrFormat)
		}
	}()

	opt := minio.PutObjectOptions{
		ContentType: "image/" + ext,
	}
	_, err := s.cli.PutObjectWithContext(s.ctx, s.bucketName, filename, r, -1, opt)
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

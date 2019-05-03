package imagestore

import (
	"bufio"
	"context"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"os"
	"testing"

	"github.com/pkg/errors"
	"golang.org/x/image/draw"
	"google.golang.org/grpc/grpclog"
)

var fileNames = []string{
	"./cases/ramen.jpg",
	"./cases/ramen.png",
	"./cases/hanabi.jpg",
}

type Case struct {
	Name  string
	Ext   string
	Image image.Image
}

func BenchmarkImageStoreImpl_Resize(b *testing.B) {
	jpg, err := os.Open("./cases/ramen.jpg")
	if err != nil {
		b.Fatal(err)
	}
	defer jpg.Close()

	jpgImg, _, err := image.Decode(jpg)
	if err != nil {
		b.Fatal(err)
	}

	png, err := os.Open("./cases/ramen.png")
	if err != nil {
		b.Fatal(err)
	}
	defer png.Close()

	pngImg, _, err := image.Decode(png)
	if err != nil {
		b.Fatal(err)
	}

	is := &imageStoreImpl{
		ctx: context.TODO(),
	}

	for _, size := range imageSizes {
		b.Run(fmt.Sprintf("resize jpeg to %vpx", size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				is.Resize(jpgImg, size)
			}
		})
	}

	for _, size := range imageSizes {
		b.Run(fmt.Sprintf("resize png to %vpx", size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				is.Resize(pngImg, size)
			}
		})
	}
}

func TestResize(t *testing.T) {
	cases := []struct {
		Name string
		draw.Interpolator
	}{
		{
			Name:         "NearestNeighbor",
			Interpolator: draw.NearestNeighbor,
		},
		{
			Name:         "ApproxBiLinear",
			Interpolator: draw.ApproxBiLinear,
		},
		{
			Name:         "BiLinear",
			Interpolator: draw.BiLinear,
		},
		{
			Name:         "CatmullRom",
			Interpolator: draw.CatmullRom,
		},
	}

	src, err := os.Open("./cases/ramen.jpg")
	if err != nil {
		t.Fatal(err)
	}
	defer src.Close()

	srcImg, _, err := image.Decode(src)
	if err != nil {
		t.Fatal(err)
	}

	_ = os.Mkdir("./out", 0700)

	size := 512
	for _, c := range cases {
		srcW, srcH := srcImg.Bounds().Dx(), srcImg.Bounds().Dy()

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
		c.Scale(dst, dst.Bounds(), srcImg, srcImg.Bounds(), draw.Over, nil)

		f, err := os.Create(fmt.Sprintf("./out/ramen_%v.jpg", c.Name))
		if err != nil {
			t.Fatal(err)
		}

		err = jpeg.Encode(f, dst, nil)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func BenchmarkImageStoreImpl_DecodeImage(b *testing.B) {
	jpg, err := os.Open("./cases/ramen.jpg")
	if err != nil {
		b.Fatal(err)
	}
	defer jpg.Close()

	b.Run("decode jpeg", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, err := jpg.Seek(0, 0)
			if err != nil {
				b.Fatal(err)
			}

			_, _, err = image.Decode(jpg)
			if err != nil {
				b.Fatal(err)
			}
		}
	})

	png, err := os.Open("./cases/ramen.png")
	if err != nil {
		b.Fatal(err)
	}
	defer png.Close()

	b.Run("decode png", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, err := png.Seek(0, 0)
			if err != nil {
				b.Fatal(err)
			}

			_, _, err = image.Decode(png)
			if err != nil {
				b.Fatal(err)
			}
		}
	})
}

type MockWriter struct{}

func (w MockWriter) Write(p []byte) (n int, err error) {
	return len(p), nil
}

func BenchmarkImageStoreImpl_EncodeImage(b *testing.B) {
	jpgFile, err := os.Open("./cases/ramen.jpg")
	if err != nil {
		b.Fatal(err)
	}
	defer jpgFile.Close()

	jpgImg, _, err := image.Decode(jpgFile)
	if err != nil {
		b.Fatal(err)
	}

	b.Run("encode jpeg", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			err := jpeg.Encode(MockWriter{}, jpgImg, nil)
			if err != nil {
				b.Fatal(err)
			}
		}
	})

	pngFile, err := os.Open("./cases/ramen.png")
	if err != nil {
		b.Fatal(err)
	}
	defer pngFile.Close()

	pngImg, _, err := image.Decode(pngFile)
	if err != nil {
		b.Fatal(err)
	}

	b.Run("encode png", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			err := png.Encode(MockWriter{}, pngImg)
			if err != nil {
				b.Fatal(err)
			}
		}
	})
}

func BenchmarkImageStoreImpl_putImage(b *testing.B) {
	cases := []Case{}
	for _, n := range fileNames {
		func() {
			f, err := os.Open(n)
			if err != nil {
				b.Fatal(err)
			}
			defer func() {
				_ = f.Close()
			}()
			img, ext, err := image.Decode(f)
			if err != nil {
				b.Fatal(err)
			}

			cases = append(cases, Case{
				Name:  f.Name(),
				Ext:   ext,
				Image: img,
			})
		}()
	}

	for _, c := range cases {
		cur := c
		b.Run(fmt.Sprintf("put image %v", c.Name), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				r, w := io.Pipe()
				defer r.Close()
				go func() {
					var err error
					defer func() {
						e := w.CloseWithError(err)
						if e != nil {
							grpclog.Error(e)
						}
					}()
					bw := bufio.NewWriter(w)
					switch cur.Ext {
					case "gif":
						err = errors.WithStack(gif.Encode(bw, cur.Image, nil))
					case "jpeg":
						err = errors.WithStack(jpeg.Encode(bw, cur.Image, nil))
					case "png":
						err = errors.WithStack(png.Encode(bw, cur.Image))
					default:
						err = errors.WithStack(image.ErrFormat)
					}
				}()

				buf := make([]byte, 4096)
				var err error
				for ; err == nil; _, err = r.Read(buf) {
				}
				if err != io.EOF {
					b.Fatal(err)
				}
			}
		})
	}
}

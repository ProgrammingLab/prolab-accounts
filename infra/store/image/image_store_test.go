package imagestore

import (
	"context"
	"fmt"
	"image"
	"image/jpeg"
	"os"
	"testing"

	"golang.org/x/image/draw"
)

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
			jpg.Seek(0, 0)
			_, _, err := image.Decode(jpg)
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
			png.Seek(0, 0)
			_, _, err := image.Decode(png)
			if err != nil {
				b.Fatal(err)
			}
		}
	})
}

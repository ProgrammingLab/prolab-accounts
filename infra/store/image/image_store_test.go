package imagestore

import (
	"context"
	"fmt"
	"image"
	_ "image/jpeg"
	"os"
	"testing"
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

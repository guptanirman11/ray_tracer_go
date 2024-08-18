package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
)

func main() {
	// TODO
	imgWidth := 256
	imgHeight := 256

	img := image.NewRGBA(image.Rect(0, 0, imgWidth, imgHeight))

	for y := 0; y < imgHeight; y++ {
		for x := 0; x < imgWidth; x++ {

			r := float64(x) / float64(imgWidth-1)
			g := float64(y) / float64(imgHeight-1)
			b := 0.2

			img.Set(x, y, color.RGBA{
				uint8(r * 255.999),
				uint8(g * 255.999),
				uint8(b * 255.999),
				255,
			})
		}

	}
	f, err := os.Create("output.png")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	err = png.Encode(f, img)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Done")
}

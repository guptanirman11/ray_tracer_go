package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
)

func main() {
	imgWidth := 400
	imgHeight := 400

	img := image.NewRGBA(image.Rect(0, 0, imgWidth, imgHeight))

	for j := 0; j < imgHeight; j++ {
		for i := 0; i < imgWidth; i++ {

			v := Vector{x: float64(i) / float64(imgWidth-1), y: float64(j) / float64(imgHeight-1), z: 0.2}

			img.Set(i, j, color.RGBA{
				uint8(v.x * 255.999),
				uint8(v.y * 255.999),
				uint8(v.z * 255.999),
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

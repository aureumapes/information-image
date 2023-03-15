package internal

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

func ToImage(writeTo, file string) {
	fileBytes, _ := os.ReadFile(file)
	currColor := color.RGBA{
		R: 0,
		G: 0,
		B: 0,
		A: 255,
	}
	currX := 0
	img := image.NewRGBA(image.Rectangle{image.Point{0, 0}, image.Point{len(fileBytes), 1}})
	for _, textByte := range fileBytes {
		currColor.R, currColor.G, currColor.B = textByte, textByte, textByte
		img.Set(currX, 0, currColor)
		currColor = color.RGBA{
			R: 0,
			G: 0,
			B: 0,
			A: 255,
		}
		currX++
	}
	f, _ := os.Create(writeTo)
	png.Encode(f, img)

}

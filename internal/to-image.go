package internal

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

func ToImage(writeTo, file string) {
	fileBytes, _ := os.ReadFile(file)

	img := image.NewRGBA(image.Rectangle{image.Point{0, 0}, image.Point{1000, 1000}})
	currX := img.Bounds().Min.X
	currY := img.Bounds().Min.Y

	for _, textByte := range fileBytes {
		img.Set(currX, currY, color.RGBA{textByte, textByte, textByte, 0xff})
		currX++
		if currX >= img.Bounds().Max.X {
			currX = 0
			currY++
		}
	}
	f, _ := os.Create(writeTo)
	png.Encode(f, img)

}

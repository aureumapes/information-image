package internal

import (
	"image"
	"image/color"
	"image/png"
	"math"
	"os"
)

func ToImage(writeTo, file string) {
	fileBytes, _ := os.ReadFile(file)

	width := int(math.Sqrt(float64(len(fileBytes))))
	height := width + (len(fileBytes) - width*width)

	img := image.NewRGBA(image.Rectangle{image.Point{0, 0}, image.Point{width, height}})
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

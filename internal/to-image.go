package internal

import (
	"image"
	"image/color"
	"image/png"
	"os"
	"math"
)

func ToImage(writeTo, file string) {
	fileBytes, _ := os.ReadFile(file)
	width_height := int(math.Sqrt(float64(len(fileBytes))))+2

	img := image.NewRGBA(image.Rectangle{image.Point{0, 0}, image.Point{width_height, width_height}})
	currX := img.Bounds().Min.X
	currY := img.Bounds().Min.Y

	for _, textByte := range fileBytes {
		img.Set(currX, currY, color.RGBA{textByte, textByte, textByte, 255})
		currX++
		if currX >= img.Bounds().Max.X {
			currX = 0
			currY++
		}
	}
	img.Set(width_height-1, width_height-1, color.RGBA{0xff, 0x00, 0x00, 0xaa})
	f, err := os.Create(writeTo)
	if err != nil {
		panic(err)	
	}
	png.Encode(f, img)
}

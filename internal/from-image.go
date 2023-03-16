package internal

import (
	"bytes"
	"image"
	"image/png"
	"os"
)

func FromImage(writeFrom, file string) {
	imageBytes, _ := os.ReadFile(writeFrom)
	img, _ := png.Decode(bytes.NewReader(imageBytes))
	//fileBytes := []byte{}

	bytes := getBytes(img)
	returned := string(bytes)

	print(returned)
}

func getBytes(img image.Image) []byte {
	var bytes []byte
	height, width := img.Bounds().Max.X, img.Bounds().Max.Y
	for y := 0; y <= height; y++ {
		for x := 0; x <= width; x++ {
			r, _, _, _ := img.At(x, y).RGBA()
			bytes = append(bytes, byte(r))
		}
	}
	return bytes
}

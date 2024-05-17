package internal

import (
	"bytes"
	"image"
	"image/png"
	"os"
)

func FromImage(writeFrom string) {
	imageBytes, _ := os.ReadFile(writeFrom)
	img, _ := png.Decode(bytes.NewReader(imageBytes))

	bytes := getBytes(img)

	println(string(bytes))
}

func getBytes(img image.Image) []byte {
	var bytes []byte
	height, width := img.Bounds().Max.X, img.Bounds().Max.Y
	for y := 0; y <= height; y++ {
		for x := 0; x <= width; x++ {
			r, _, _, a := img.At(x, y).RGBA()
			if a == 65535 {
				bytes = append(bytes, byte(r))
			}
		}
	}
	return bytes
}

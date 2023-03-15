package internal

import (
	"bytes"
	"image/png"
	"os"
)

func FromImage(writeFrom, file string) {
	imageBytes, _ := os.ReadFile(writeFrom)
	img, _ := png.Decode(bytes.NewReader(imageBytes))
	fileBytes := []byte{}

	currX := img.Bounds().Min.X
	currY := img.Bounds().Min.Y

	for currY <= img.Bounds().Max.Y && currX <= img.Bounds().Max.X {
		r, _, _, _ := img.At(currX, currY).RGBA()
		fileBytes = append(fileBytes, byte(r))
		currX++
		if currX < img.Bounds().Min.X {
			currY++
			if currY < img.Bounds().Min.Y {
				break
			}
			currX = img.Bounds().Max.X
		}

	}
	print(string(fileBytes))
}

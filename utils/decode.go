package utils

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"strings"
)

// DecodeImage takes a file and decodes it
func DecodeImage(filename string) image.Image {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		panic(fmt.Sprintf("File [%s] does not exist", filename))
	}
	ext := getExtension(filename)
	reader, _ := os.Open(filename)
	defer reader.Close()
	if ext == "png" {
		img, err := png.Decode(reader)
		if err != nil {
			panic(err)
		}
		return img
	}
	img, err := jpeg.Decode(reader)
	if err != nil {
		panic(err)
	}
	return img
}

func getExtension(filename string) string {
	splitString := strings.Split(filename, ".")
	ext := strings.ToLower(splitString[len(splitString)-1])
	if ext == "png" || ext == "jpg" || ext == "jpeg" {
		return ext
	}
	panic(fmt.Sprintf("Unsupported file type: [%s]", ext))
}

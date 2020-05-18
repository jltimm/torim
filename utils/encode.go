package utils

import (
	"image"
	"image/color"
	"image/jpeg"
	"os"
)

//EncodeDiff creates an image from the 2D diff array
func EncodeDiff(filename string, diffs [][2]int, img1Bounds, img2Bounds image.Rectangle) {
	imgBoundsX := img1Bounds.Max.X
	if img2Bounds.Max.X > imgBoundsX {
		imgBoundsX = img2Bounds.Max.X
	}
	imgBoundsY := img1Bounds.Max.Y
	if img2Bounds.Max.Y > imgBoundsY {
		imgBoundsY = img2Bounds.Max.Y
	}
	overlayImage := image.NewNRGBA(image.Rect(0, 0, imgBoundsX, imgBoundsY))
	black := color.NRGBA{0, 0, 0, 0}
	for x := 0; x < imgBoundsX; x++ {
		for y := 0; y < imgBoundsY; y++ {
			overlayImage.Set(x, y, black)
		}
	}
	white := color.NRGBA{255, 255, 255, 255}
	for i := 0; i < len(diffs); i++ {
		overlayImage.Set(diffs[i][0], diffs[i][1], white)
	}
	outputFile, _ := os.Create(filename)
	defer outputFile.Close()
	jpeg.Encode(outputFile, overlayImage, nil)
}

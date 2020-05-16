package compare

import (
	"image"
)

// FindDifferences compares two images, finding the differences
func FindDifferences(img1 image.Image, img2 image.Image) [][2]int {
	img1Bounds := img1.Bounds()
	diffs := [][2]int{}
	for x := 0; x < img1Bounds.Max.X; x++ {
		for y := 0; y < img1Bounds.Max.Y; y++ {
			r1, g1, b1, a1 := img1.At(x, y).RGBA()
			r2, g2, b2, a2 := img2.At(x, y).RGBA()
			if r1 != r2 || g1 != g2 || b1 != b2 || a1 != a2 {
				diff := [2]int{x, y}
				diffs = append(diffs, diff)
			}
		}
	}
	return diffs
}

package main

import (
	"fmt"
	"image"
	"os"

	"./compare"
	"./utils"
)

func main() {
	args := os.Args[1:]
	if len(args) == 2 {
		fmt.Println("Comparing the files")
		img1chan := make(chan image.Image)
		img2chan := make(chan image.Image)
		go utils.DecodeImage(args[0], img1chan)
		go utils.DecodeImage(args[1], img2chan)
		img1, img2 := <-img1chan, <-img2chan
		diffs := compare.FindDifferences(img1, img2)
		utils.EncodeDiff("test.jpg", diffs, img1.Bounds(), img2.Bounds())
	}
}

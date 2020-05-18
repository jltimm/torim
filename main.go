package main

import (
	"fmt"
	"os"

	"./compare"
	"./utils"
)

func main() {
	args := os.Args[1:]
	if len(args) == 2 {
		fmt.Println("Comparing the files")
		img1 := utils.DecodeImage(args[0])
		img2 := utils.DecodeImage(args[1])
		diffs := compare.FindDifferences(img1, img2)
		utils.EncodeDiff("test.jpg", diffs, img1.Bounds(), img2.Bounds())
	}
}

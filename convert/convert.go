package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"os"
)

func init() {
	// damn important or else At(), Bounds() functions will
	// caused memory pointer error!!
	image.RegisterFormat("jpeg", "jpeg", jpeg.Decode, jpeg.DecodeConfig)
}

func main() {
	imgfile, err := os.Open("./images/5.jpg")

	if err != nil {
		fmt.Println("img.jpg file not found!")
		os.Exit(1)
	}

	defer imgfile.Close()

	img, _, err := image.Decode(imgfile)

	fmt.Println(img.At(10, 10))

	bounds := img.Bounds()

	fmt.Println(bounds)

	canvas := image.NewAlpha(bounds)

	// is this image opaque
	op := canvas.Opaque()

	fmt.Println(op)
}

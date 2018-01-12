package main

import (
	"fmt"
	//"github.com/nfnt/resize"
	"image"
	"image/gif"
	"image/jpeg"
	"log"
	"os"
)

func main() {
	myImage := "sanju.jpg"
	fmt.Println("The image is:", myImage)
	imageToBeRead := readTheImage(myImage)
	fmt.Println("Reading image data ...")
	fmt.Println("Converting  image to thumbnail  ...")
	convertToGIF(imageToBeRead)
}

func readTheImage(ImageFile string) (image image.Image) {
	// open "test.jpg"
	imgFile, err := os.Open(ImageFile)
	if err != nil {
		log.Fatal(err)
	}

	// decode jpeg into image.Image
	img, err := jpeg.Decode(imgFile)
	if err != nil {
		log.Fatal(err)
	}
	imgFile.Close()

	return img
}

func convertToGIF(img image.Image) {
	out, err := os.Create("sanju.gif")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	var opt gif.Options
	opt.NumColors = 256
	err = gif.Encode(out, img, &opt)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(img, "\n success... \n ")

}

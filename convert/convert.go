package main

import (
	"fmt"
	//"github.com/nfnt/resize"
	"image"
	"image/jpeg"
	"image/png"
	"log"
	"os"
)

func main() {
	myImage := "./images/5.jpg"
	fmt.Println("The image is:", myImage)
	imageToBeRead := readTheImage(myImage)
	fmt.Println("Reading image data ...")
	fmt.Println("Converting  image to png  ...")
	convertToPNG(imageToBeRead)
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

func convertToPNG(img image.Image) {
	out, err := os.Create("./images/5.png")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	err = png.Encode(out, img)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(img, "\n success... \n ")

}

package main

import (
	"fmt"
	//"github.com/nfnt/resize"
	"image"
	"image/jpeg"
	"log"
	"os"

	"github.com/nfnt/resize"
)

func main() {
	myImage := "sanju.jpg"
	fmt.Println("The image is:", myImage)
	imageToBeRead := readTheImage(myImage)
	fmt.Println("Reading image data ...")
	fmt.Println("Resizing  ...")
	resizeImage(imageToBeRead)
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

func resizeImage(img image.Image) {

	m := resize.Resize(100, 100, img, resize.NearestNeighbor)

	out, err := os.Create("sanju_resized.jpg")
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	// write new image to file
	anyError := jpeg.Encode(out, m, nil)
	if anyError != nil {
		fmt.Println("Resizing did not work ")
		os.Exit(0)
	}
	fmt.Println("success...")

}

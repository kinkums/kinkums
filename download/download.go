package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	url := "https://github.com/kinkums/kinkums/tree/master/webupload/imageOutput/sanju.jpg"
	response, er := http.Get(url)
	if er != nil {
		log.Fatal(er)
	}
	defer response.Body.Close()
	file, err := os.Create("sanju.jpg")
	if err != nil {
		log.Fatal(err)
	}

	_, err1 := io.Copy(file, response.Body)
	if err1 != nil {
		log.Fatal(err1)
	}
	file.Close()
	fmt.Println("Image downloaded")

	myImage := "sanju.jpg"
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
	out, err := os.Create("Picture 400.png")
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

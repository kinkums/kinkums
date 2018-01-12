package main

import (
	"fmt"
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

}

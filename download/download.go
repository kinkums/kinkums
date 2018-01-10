package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	url := "https://github.com/kinkums/kinkums/tree/master/webupload/test/Picture002.jpg"
	response, er := http.Get(url)
	if er != nil {
		log.Fatal(er)
	}
	defer response.Body.Close()
	file, err := os.Create("./images/pic2.jpg")
	if err != nil {
		log.Fatal(err)
	}

	_, err1 := io.Copy(file, response.Body)
	if err1 != nil {
		log.Fatal(err1)
	}
	file.Close()
	fmt.Println("Image downloaded")

	/*f, err := os.Create("./convert/pic2.png")
		if err != nil {
			log.Fatal(err)
		}
	  img
		if err := png.Encode(f, "./images/pic2.jpg"); err != nil {
			f.Close()
			log.Fatal(err)
		}

		if err := f.Close(); err != nil {
			log.Fatal(err)
		}*/
}

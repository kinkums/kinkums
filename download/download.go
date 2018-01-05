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
		log.Fatal("%s", er)
	}
	defer response.Body.Close()
	file, err := os.Create("/images/pic2.jpg")
	if err != nil {
		log.Fatal("%s", err)
	}

	_, err1 := io.Copy(file, response.Body)
	if err1 != nil {
		log.Fatal("%s", err1)
	}
	file.Close()
	fmt.Println("Image downloaded")
}

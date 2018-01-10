package main

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
)

func upload(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, _ := template.ParseFiles("upload.gtpl")
		t.Execute(w, nil)
		//fmt.Println(" Method is Get ")
	} else if r.Method == "POST" {
		file, handler, err := r.FormFile("uploadfile")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		//fmt.Println(" Method is Post ")
		fmt.Fprintf(w, "%v", handler.Header)
		f, err := os.OpenFile("./imageOutput/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		io.Copy(f, file)
	} else {
		fmt.Println("Unknown HTTP " + r.Method + " Method")
	}
}

func main() {
	http.HandleFunc("/mediaservice", upload)
	http.ListenAndServe(":9090", nil)
}

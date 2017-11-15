package main

import (
	. "fmt"
	"io/ioutil"
)

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	filename := "c:\\gocode\\src\\github.com\\kinkums\\kinkums\\goweb\\" + p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	fileName := title + ",txt"
	body, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func main() {
	p1 := &Page{Title: "TestPage", Body: []byte("This is a sample web page.")}
	p1.save()
	p2, _ := loadPage("TestPage")
	Println(string(p2.Body))
}
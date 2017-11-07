package main

import "fmt"

const (
	pi       = 3.14
	language = "Go"
)

const (
	_  = iota
	KB = 1 << (iota * 10)
	MB = 1 << (iota * 10)
	GB = 1 << (iota * 10)
)

const p string = "death and taxes"

func main() {
	const q = 42
	fmt.Println(" p - ", p)
	fmt.Println(" q - ", q)
	fmt.Println(" language - ", language)
	fmt.Println(" pi - ", pi)
	fmt.Println(" KB- ", KB)
	fmt.Println(" MB - ", MB)
	fmt.Println(" GB - ", GB)
}

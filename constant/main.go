package main

import "fmt"

const (
	pi       = 3.14
	language = "Go"
)

const p string = "death and taxes"

func main() {
	const q = 42
	fmt.Println(" p - ", p)
	fmt.Println(" q - ", q)
	fmt.Println(" language - ", language)
	fmt.Println(" pi - ", pi)
}

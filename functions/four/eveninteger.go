package main

import "fmt"

func half(n int) (int, bool) {
	return n / 2, n%2 == 0
}

func main() {
	h, even := half(2)
	fmt.Println(h, even)
	i, odd := half(5)
	fmt.Println(i, odd)
}

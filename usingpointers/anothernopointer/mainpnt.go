package main

import "fmt"

func zero1(x int) {
	fmt.Println("%p\n", &x)
	fmt.Println(&x)
	x = 0
}

func main() {
	x := 5
	fmt.Println("%p\n", &x)
	fmt.Println(&x)
	zero1(x)
	fmt.Println(x)
}

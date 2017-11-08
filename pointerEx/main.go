package main

import "fmt"

func main() {
	a := 43
	fmt.Println("Value of a is ", a)
	fmt.Println("Address of a is ", &a)
	var b *int = &a
	fmt.Println("Value of b is ", b)
}

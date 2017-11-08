package main

import "fmt"

func main() {
	a := 43
	fmt.Println("Value of a is ", a)
	fmt.Println("Address of a is ", &a)
	var b *int = &a
	fmt.Println("Value of b is ", b)
	fmt.Println("Value of *b is", *b)
	*b = 42
	fmt.Println("New Value of a after reassigning b is ", a)
}

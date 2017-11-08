package main

import "fmt"

func zero(x *int) {
	*x = 0
}

func main() {
	x := 5
	zero(&x)
	fmt.Println("Value of x is ", x)
}

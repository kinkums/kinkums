package main

import "fmt"

func max(numbers ...int) int {
	var largest int
	for _, v := range numbers {
		if v > largest {
			largest = v
		}

	}
	return largest
}

func main() {
	greatest := max(15, 25, 2, 8, 13, 95, 7, 111, 3)
	fmt.Println("Largest Number is: ", greatest)
}

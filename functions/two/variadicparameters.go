package main

import "fmt"

func average(sf ...float64) float64 {
	total := 0.0
	for _, v := range sf {
		total += v
	}
	return total / float64(len(sf))
}

func main() {
	n := average(15, 25, 32, 41, 18, 19, 66)
	fmt.Println("Average: ", n)
}

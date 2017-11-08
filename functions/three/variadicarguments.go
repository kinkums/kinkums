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
	data := []float64{30, 40, 50, 60, 70, 80, 90}
	n := average(data...)
	fmt.Println("Average is: ", n)
}

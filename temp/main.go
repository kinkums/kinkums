package main

import "fmt"

func main() {
	var str1 string
	fmt.Println("Enter your Name ")
	fmt.Scanf("%s", &str1)
	if str1 == "Rad" {
		fmt.Printf("My Name is %s", str1)
	} else {
		fmt.Printf("My Name is not Rad")
	}

}

package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Print("Enter your choice (C/E) ")
	a := checkResponse()
	fmt.Println(" Value of A is ", a)
}

func checkResponse() bool {
	var userInput string
	_, err := fmt.Scan(&userInput)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Type of UI is %T", userInput)
	fmt.Println("User Input values are ", userInput)
	//const cont string = "C"
	//const exits string = "E"
	if userInput == "C" {
		fmt.Println("Came into if")
		return true
	} else if userInput == "E" {
		fmt.Println("Came into else")
		return false
	} else {
		fmt.Println("Valid user values are C or E")
		return checkResponse()
	}
}

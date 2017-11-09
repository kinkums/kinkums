package main

import "fmt"

import "log"
import "time"

func main() {

	c1 := make(chan string)
	c2 := make(chan string)
	go func() {
		for {
			time.Sleep(time.Second * 10)
			c1 <- "eligible for a 5 minute break"

		}
	}()
	go func() {
		for {
			time.Sleep(time.Second * 25)
			c2 <- "eligible for a 30 minute break"

		}
	}()
	go func() {
		for {
			select {

			case msg2 := <-c2:

				myOption := checkResponse()
				fmt.Println("Option selected is ", myOption)
				if myOption == true {
					fmt.Println(msg2)
				} else {
					break
				}
			case msg1 := <-c1:
				fmt.Println(msg1)
			}
		}
	}()
	var input string
	fmt.Scan(&input)
}

func checkResponse() bool {
	var userInput string
	fmt.Print("Do you want to continue / exit (C/E)")

	_, err := fmt.Scan(&userInput)
	fmt.Println("2")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("User Input value ", userInput)

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

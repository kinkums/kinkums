package main

import (
	"fmt"
	"log"
	"time"
)

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
				fmt.Print("Do you want to continue / exit (Press C for continue, any character for exit)")
				var userInput string
				_, err := fmt.Scan(&userInput)
				fmt.Println("User Input value ", userInput)
				if err != nil {
					log.Fatal(err)
				}

				//myOption := checkResponse()
				//fmt.Println("Option selected is ", myOption)
				//fmt.Println("Input Value ", myOption)
				if userInput == "C" {
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
	fmt.Scanln(&input)
}

/*func checkResponse() string {
	var userInput string
	_, err := fmt.Scan(&userInput)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("User Input value ", userInput)

	if userInput == "C" {
		fmt.Println("Came into if")
		return "C"
	} else {
		fmt.Println("Came into else")
		return "X"
	}
}*/

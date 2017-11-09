package main

import "fmt"

import "log"
import "time"

func main() {

	c1 := make(chan string)
	c2 := make(chan string)
	go func() {
		for {
			time.Sleep(time.Minute * 2)
			c1 <- "eligible for a 5 minute break"

		}
	}()
	go func() {
		for {
			time.Sleep(time.Minute * 5)
			c2 <- "eligible for a 30 minute break"

		}
	}()
	go func() {
		for {
			select {

			case msg2 := <-c2:
				fmt.Print("Do you want to continue / exit (C/E)")
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
	fmt.Scanln(&input)
}

func checkResponse() bool {
	var userInput string
	_, errorOccured := fmt.Scanln(&userInput)
	if errorOccured != nil {
		log.Fatal(errorOccured)
	}
	const cont string = "C"
	const exits string = "E"
	if userInput == cont {
		return true
	} else if userInput == exits {
		return false
	} else {
		fmt.Println("Valid user values are C or E")
		return checkResponse()
	}
}

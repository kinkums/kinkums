package main

import (
	"fmt"
	"os"
	"time"
)

func main() {

	c1 := make(chan string)
	c2 := make(chan string)
	done := make(chan struct{})
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
			time.Sleep(time.Second * 25)
			fmt.Println("Press any key to exit or do nothing for further")
			os.Stdin.Read(make([]byte, 1))
			close(done)

		}
	}()
	go func() {
		for {
			select {
			case msg3 := <-done:
				fmt.Println(msg3)
			case msg2 := <-c2:
				/*fmt.Print("Enter your choice (C/E) ")
				var userInput string
				fmt.Scan("%s", &userInput)
				//a := checkResponse()
				fmt.Println(" Value of UserInput is %s", userInput)*/
				fmt.Println(msg2)
				/*if a {
					fmt.Println(msg2)
				} else {
					break
				}*/
			case msg1 := <-c1:
				fmt.Println(msg1)
			}

		}

	}()
	var input string
	fmt.Scanln(&input)
}

/*func checkResponse() bool {
	var userInput string
	_, err := fmt.Scan(&userInput)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println("Type of UI is %T", userInput)
	fmt.Println("User Input values are ", userInput)
	//const cont string = "C"
	//const exits string = "E"
	if userInput == "C" {
		fmt.Println("Came into if")
		return true
	} else {
		fmt.Println("Came into else")
		return false
	} /*else {
		fmt.Println("Valid user values are C or E")
		return checkResponse()
	}
}*/

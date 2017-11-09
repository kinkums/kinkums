package main

import (
	"fmt"
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
			fmt.Print("Enter 1 to Continue")
			var i string
			_, err := fmt.Scanln(&i)
			fmt.Println("Value of dump is ", i)
			fmt.Println("Value of err is ", err)

		}
	}()

	go func() {
		for {
			select {

			case msg2 := <-c2:

				fmt.Println(msg2)
				//var dump int

			case msg1 := <-c1:
				fmt.Println(msg1)
			}
		}
	}()
	/*fmt.Print("Enter your choice (C/E) ")
	a := checkResponse()
	fmt.Println(" Value of A is ", a)*/

	var input string
	fmt.Scanln(&input)
}

/*func checkResponse() bool {
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
}*/

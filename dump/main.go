package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)
	//c3 := make(chan string)
	var i string

	go func() {
		for {
			time.Sleep(time.Second * 11)
			c1 <- "eligible for a 5 minute break"

		}
	}()

	go func() {
		for {
			time.Sleep(time.Second * 17)
			//c2 <- "eligible for a 30 minute break"
			c2 <- i
		}
	}()

	/*go func() {
		for {

			c3 <- i
		}
	}()*/

	go func() {
		for {
			select {

			//case msg3 := <-c3:
			//fmt.Println("Returned Value ", msg3)
			case msg2 := <-c2:
				fmt.Println(msg2)
			//i := "Yes"
			case msg1 := <-c1:
				fmt.Println(msg1)
			}
		}
	}()
	/*fmt.Print("Enter your choice (C/E) ")
	a := checkResponse()
	fmt.Println(" Value of A is ", a)*/

	time.Sleep(time.Second * 17)
	fmt.Println("Enter a value ")

	fmt.Scan(&i)
	//fmt.Println(i)
	var input string
	fmt.Scanln(&input)
}

/*func value() string {
	fmt.Println(" Enter a String Value")
	var data string
	fmt.Scanf("%s", data)
	return data
}*/

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

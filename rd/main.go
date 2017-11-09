package main

import (
	"fmt"
	"time"
)

func main() {
	var myOption bool
	c1 := make(chan string)
	c2 := make(chan string)
	go func() {
		for {
			c1 <- "eligible for a 5 minute break"
			time.Sleep(time.Minute * 2)
		}
	}()
	go func() {
		for {
			c2 <- "eligible for a 30 minute break"
			time.Sleep(time.Minute * 5)
		}
	}()
	go func() {
		for {
			select {

			case msg2 := <-c2:
				fmt.Print("Do you want to continue (Y/N)")
				fmt.Scan(&myOption)
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

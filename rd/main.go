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
			case msg1 := <-c1:
				fmt.Println(msg1)
			case msg2 := <-c2:
				fmt.Println(msg2)
			}
		}
	}()
	var input string
	fmt.Scanln(&input)
}

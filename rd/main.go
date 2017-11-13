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
			time.Sleep(time.Hour * 1)
			c1 <- "eligible for a 5 minute break"
			time.Sleep(time.Minute * 5)
			c1 <- "Get back to work"
		}
	}()

	go func() {
		for {
			time.Sleep(time.Hour * 3)
			c2 <- "you are eligible for a 30 minute break\n Press any key if you want to exit or do nothing to continue\n"
			time.Sleep(time.Minute * 30)
			c2 <- "Get back to work"
		}
	}()

	go func() {
		for {
			time.Sleep(time.Hour * 3)
			os.Stdin.Read(make([]byte, 1))
		}
	}()

	go func() {
		for {
			select {
			case <-done:
				close(done)
			case msg2 := <-c2:
				fmt.Println(msg2)
			case msg1 := <-c1:
				fmt.Println(msg1)
			}
		}
	}()

	var input string
	fmt.Scanln(&input)
}

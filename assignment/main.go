package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	clock()
}

var breakTime int
var timeCheck int

func clock() {
	st := time.Now()
	chan1 := make(chan int)
	for {
		time.Sleep(10 * time.Second)         //For Testing: change the sleep time to (1 * time.Seconds)
		end := int(time.Since(st).Seconds()) //For Testing: change this value to time.Since(st).Seconds()

		go addBreak(chan1)
		if (end % 15) == 0 { //For testing: changing this value to a smaller number
			chan1 <- 5
			if (end % 25) == 0 {
				chan1 <- 30
				go removeBreak(chan1)
			}

		}
	}
}
func addBreak(chan1 chan int) {
	br := <-chan1
	breakTime += br //this variable will be increased by 5 mins every hour and after 3 hours, it will be increased by 30 mins
	fmt.Println("break time ", breakTime)
}

func removeBreak(chan1 chan int) {
	brTime := <-chan1
	var option string
	if brTime < timeCheck {
		fmt.Println("Do you want to quit (y/n)?")
		fmt.Scanf("%s", option)
		if option == "Y" {
			os.Exit(1)
		}

	}
}

/* You can write some methods for :
1.  If user takes break then total time taken for break should be substracted from breakTime
2.  if breakTime < totalTimeTakenForBreak then ask user whether he wants to quit application.
os.Exit can be used here. */

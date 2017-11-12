package main

import (
	"fmt"
	"os"
	"time"
)



func main() {
	c1 := make(chan string)
	c2 := make(chan string)
	//c3 := make(chan string)
	done := make(chan struct{})

	go func() {
		for {
			time.Sleep(time.Second * 11)
			c1 <- "eligible for a 5 minute break"

		}
	}()

	go func() {
		for {
			time.Sleep(time.Second * 17)
			c2 <- "eligible for a 30 minute break"
			//c2 <- i

			//fmt.Println("Press Y to close, any other key to continue ")

			//os.Stdin.Read(make([]byte, 1))
			//fmt.Println("A", xx)
			//c2 <- "Application closing"

			/*} else {
				c2 <- "eligible for a 30 minute break"
			}*/

		}

	}()

	go func() {
		for {
			os.Stdin.Read(make([]byte, 1))
			close(done)
		}
	}()

	func cancelled() bool {
		select {
		case <-done:
			return true
		default:
			return false
		}
	}

	go func() {
		for {
			select {
			//case msg3 := <-done:
			//fmt.Println("App closed", msg3)
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

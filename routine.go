//go:build routine
// +build routine

package main

import (
	"fmt"
	"time"
)

func main() {
	
	// go routine. Sometimes in Go, the main function finish and exit even before the go
	// routine func is run.
	go func(str1 string) {
		fmt.Println(str1)
	}("Hello everybody")

	time.Sleep(time.Second)
	fmt.Println("Go routine and channel")
	fmt.Println("done")

	// go channel
	msgChan := make(chan string) // this is an unbuffered channel, which mean it does not have a fixed size
	// msgChan <- "Ran" // send data to the channel
	// name := <-msgChan
	// msgChan <- "Takahashi"
	// lastName := <-msgChan

		// the above will cause a deadlock because line 25 will block everything
		// we need to either use go routine or buffered channel to concurrently send and receive data
		// channel works followed first-in-first-out principle

	go func() {
			msgChan <- "Ran"
			msgChan <- "Takahashi"
	}()

	name := <-msgChan
	lastName := <-msgChan

	fmt.Println("Player", name, lastName)

	// buffered channel means it has a fixed size. If defined 2, only 2 can be send to the channel
	intBuffered := make(chan int, 2)
	intBuffered <- 2
	intBuffered <- 4
	fmt.Println("first element", <-intBuffered) // this means that an element in the channel is out
	// so we can send 9 to the channel
	// send to channel chan<-; receive from channel <-chan
	intBuffered <- 9
	// intBuffered <- 1010 // this will cause deadlock because now the channel already full with 2 values.
	fmt.Println(<-intBuffered)
	fmt.Println(<-intBuffered)
}

//go:build rountine
// +build rountine

package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 2; i++ {
		fmt.Println(from, ":", i)
	}
}
func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true // Block until we receive a notification from the worker on the channel.
}
func main() {

	//	f("direct")

	// a go rountine is a lightweight thread of execution
	// to invoke a go routine, use keyword "go sth"

	// go f("g")

	// go func(msg string) {
	// 	fmt.Println(msg)
	// }("going")

	// time.Sleep(time.Second)
	// fmt.Println("done")

	/*
	   Channels behave as a first-in-first-out queue.
	   So, when one goroutine writes data to the channel,
	    the other goroutine reads the data from the
	    channel in the same order it was written.
	*/
	msgChan := make(chan string)

	go func() {
		// time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		fmt.Println("Tell me why")
		msgChan <- "hello" // Write data to the channel
		fmt.Println("Tell me why")
		msgChan <- "world" // Write data to the channel
		msgChan <- "Finland"
	}()

	// msg1 := <-msgChan
	// msg2 := <-msgChan

	// fmt.Println(msg1)

	// Buffered channels
	// messages := make(chan string, 2)
	// messages <- "buffered"
	// messages <- "channel"
	// messages <- "test" // this will cause error

	// fmt.Println(<-messages)
	// fmt.Println(<-messages)
	// fmt.Println(<-messages)

	// channel synchronization

	done := make(chan bool, 1)
	go worker(done)

	<-done

}

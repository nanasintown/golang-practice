//go:build select
// +build select

package main

import (
	"fmt"
	"time"
)

// go select allows wait on multiple channel
// similar to switch, but this is for channel

func main() {
	c1 := make(chan string)
	c2 := make(chan string)
	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received c1", msg1)
		case msg2 := <-c2:
			fmt.Println("received c2", msg2)
		}
	}
}

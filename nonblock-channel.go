//go:build ignore
// +build ignore

package main

import "fmt"

/* main idea of non-blocking select is to simultaneously check multiple
channels for readiness

*/

func main() {
	messages := make(chan string)
	signals := make(chan bool)

	select {
	case msg := <-messages: // if messages actually has a value
		fmt.Println("Received message", msg)
	default:
		fmt.Println("no message received")
	}
	// the case above will print out the default case because messages has no value

	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("sent", msg)
	default:
		fmt.Println("No message sent")
	}

	//the case above will print out the default case
	// because <messages> has no buffer, and there is no receiver here. So <msg> cant be sent to the channel

	select {
	case msg := <-messages:
		fmt.Println("Received", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}

}

//go:build enums
// +build enums

package main

import "fmt"

// enum has a fixed number of possible values, with unique names
type ServerState int // this is an enum

// those are possible values, as constant type
const (
	StateIdle = iota // this is a keyword.  iota is a special constant that is automatically incremented for each constant in a block.
	StateConnected
	StateError
	StateRetrying
)

var stateName = map[ServerState]string{
	StateIdle:      "idle",
	StateConnected: "connected",
	StateError:     "error",
	StateRetrying:  "retrying",
}

func (ss ServerState) String() string {
	return stateName[ss]
} // this satisfies the fmt.Stringer interface, so we can print the ServerState directly
// it's similar to overide toString()

func transition(s ServerState) ServerState {
	switch s {
	case StateIdle:
		return StateConnected
	case StateConnected, StateRetrying:
		return StateIdle
	case StateError:
		return StateError
	default:
		panic(fmt.Errorf("unknown state: %s", s))
	}
}

func main() {
	fmt.Println("Interfaces")
	ns := transition(StateIdle)
	fmt.Println(ns)

	ns2 := transition(ns)
	fmt.Println(ns2)
}

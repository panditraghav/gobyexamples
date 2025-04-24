package main

import "fmt"

/*
Enumerated types (enums) are a special case of sum types.
An enum is a type that has a fixed number of possible values,
each with a distinct name. Go doesn’t have an enum type as a
distinct language feature, but enums are simple to implement
using existing language idioms.
*/

type ServerState int

/*
const (
	StateIdle      ServerState = 0
	StateConnected ServerState = 1
	StateError     ServerState = 2
	StateRetrying  ServerState = 3
)
*/
// This is same as above, iota generates successive constant values
const (
	StateIdle ServerState = iota
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
}

func main() {
	ns := nextState(StateIdle)
	fmt.Println(ns)

	ns2 := nextState(ns)
	fmt.Println(ns2)
}

func nextState(ss ServerState) ServerState {
	switch ss {
	case StateIdle:
		return StateConnected
	case StateConnected, StateRetrying:
		return StateIdle
	case StateError:
		return StateError
	default:
		panic(fmt.Errorf("Unknown state: %s", ss))
	}
}

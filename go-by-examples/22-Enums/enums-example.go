package main

import (
	"fmt"
	enums_test "myapp/go-by-examples/22-Enums/gpt-examples"
	my_enum_example "myapp/go-by-examples/22-Enums/my_enum_example"
)

// An enum is a type that has a fixed set of values, each with a distinct name.
// Go doesn't have enums, but you can use constants to achieve the same effect.
type ServiceState int

const (
	StateIdle = iota
	StateConnected
	StateError
	StateRetrying
)

var stateName = map[ServiceState]string{
	StateIdle: "Idle",
	StateConnected: "Connected",
	StateError: "Error",
	StateRetrying: "Retrying",
}

// The function name is String, and it returns a string 
func (ss ServiceState) String() string {
	return stateName[ss]
}

func main() {
	ns := transition(StateIdle)
	fmt.Println(ns)

	ns2 := transition(ns)
	fmt.Println(ns2)

	enums_test.EnumExample()

	http_status := my_enum_example.OK
	fmt.Println("httpStatusCode 200: ", http_status.String())
}

func transition(s ServiceState) ServiceState{
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
	return StateConnected
}
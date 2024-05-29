package main

import (
	"fmt"
	"time"
)

// A goroutine is a lightweight thread of execution

func f(from string) {
	for i := 0; i < 3; i++ {
		time.Sleep(time.Second)
		fmt.Println(from, ":", i)
	}
}

func main() {
	go f("goroutine2")
	// Call it in a usual way, running it synchronously.
	f("direct")

	// To invoke this function in a goroutine, use go f(s), this runs concurrently.
	go f("goroutine")

	//You can also start a goroutine for an anonymous function call.
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	time.Sleep(time.Second * 4)
	fmt.Println("done")
	
}
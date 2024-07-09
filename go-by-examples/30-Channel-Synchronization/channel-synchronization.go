package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Println("working...")
	time.Sleep(time.Second * 2)
	fmt.Println("done")

	done <- true
}

func main() {
	done := make(chan bool, 1)
	fmt.Println("Starting worker")
	go worker(done)
	fmt.Println("triggering worker in goroutine")
	/*
	You're mostly correct, but let me clarify a few points:

    The line `<-done` in the `main` function is where the program receives from the `done` channel. 
	This is a blocking operation, meaning that the `main` function will wait at this line until something is sent into the `done` channel. 
	This is a common pattern in Go to wait for a goroutine to finish execution. 
    
	Once the `worker` goroutine sends `true` into the `done` channel, the `main` function can continue execution, printing "Worker finished" and then exiting.

	So, you're on the right track! The key point to understand is that `<-done` is used to wait for the `worker` goroutine to finish.
	*/
	<- done
	fmt.Println("Worker finished")
}
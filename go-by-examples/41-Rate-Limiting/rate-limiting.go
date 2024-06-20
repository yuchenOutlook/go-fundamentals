package main

import (
	"fmt"
	"time"
)

func main() {
	requests := make(chan int, 5)
	// First we’ll look at basic rate limiting. Suppose we want to limit our handling of 
	// incoming requests. We’ll serve these requests off a channel of the same name.
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	// close the channel to indicate that's all the values we'll send.
	close(requests)

	// here is the function signiture of time.Tick() : func time.Tick(d time.Duration) <-chan time.Time
	// limiter is a channel that will receive a value every 200 milliseconds. This is the regulator in our rate limiting scheme.
	limiter := time.Tick(time.Millisecond * 1000)

	for req := range requests {
		// By blocking on a receive from the limiter channel 
		// before serving each request, we limit ourselves to 1 request every 200 milliseconds.
		<-limiter
		fmt.Println("request", req, time.Now())
	}

	// we may want to allow short bursts of requests in our rate limiting scheme while preserving the overall rate limit.

	// While Go does not set a hard limit on the size of a buffered channel, 
	// the practical limits are defined by your system’s memory and the maximum integer size. 
	// The buffer sizes you choose should be based on your application's requirements and the 
	// available system resources.
	burstyLimiter := make(chan time.Time, 10000)

	for i := 0; i < 10000; i++ {
		burstyLimiter <- time.Now()
	}

	go func() {
		// Every 1000 milliseconds we’ll try to add a new value to burstyLimiter, 
		// up to its limit of 3.
		for t := range time.Tick(time.Millisecond * 100) {
			burstyLimiter <- t
		}
	}()

	burstyRequests := make(chan int, 20000)
	for i := 1; i <= 20000; i++ {
		burstyRequests <- i
	}

	close(burstyRequests)
	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}
}
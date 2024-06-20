package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	// This is a atomic counter that will be accessed by multiple goroutines.
	// Here we use an unsigned integer to represent the counter.
	var ops atomic.Uint64

	// use the waitGroup to wait for all goroutines to finish.
	var wg sync.WaitGroup

	// here we create 50 go routines that each increment the counter exactly 1000 times.
	for i := 0; i < 50; i++ {
		// Increment the waitGroup counter for each goroutine.
		wg.Add(1)

		// fire off a goroutine.
		go func(i int) {

			// for each iteration, increment the counter exactly 1000 times.
			for c :=0; c < 1000; c++ {

				// To atomically increment the counter.
				ops.Add(1)
			}
			// Decrement the waitGroup counter when the goroutine is done.
			wg.Done()

			// You will find the ops.load is changing as the goroutines are running concurrently.
			// and the ith goroutine is not in sequence as they are running concurrently.
			fmt.Printf("this is %dth goroutine done. ops has total of %d count \n", i, ops.Load())
		}(i)
	}

	// wait for all goroutines to finish.
	wg.Wait()

	// It’s safe to access ops now because we know no other goroutine is writing to it.
	fmt.Println("ops: ", ops.Load())
}
package main

/*
To wait for multiple goroutines to finish, we can use a wait group.
worker function is the function we will run in every go routine.

*/
import (
	"fmt"
	"sync"
	"time"
)

func worker(id int) {
	fmt.Printf("Worker %d starting\n", id)

	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func main() {
	// The waitGroup is to wait for all goroutines to finish.
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		// Increment the waitGroup counter.
		wg.Add(1)

		// fire off a goroutine.
		go func() {
			// wg.Done decrements the counter.
			// defer wg.Done() statement is executed immediately.
			// However, because it is deferred, the actual call to wg.Done() will be delayed 
			// until the surrounding function (the goroutine) returns.
			defer wg.Done()

			// Call the worker function.
			worker(i)
		}()
	}

	wg.Wait()
}
package main

// https://gobyexample.com/closing-channels
// Closing a channel indicates that no more values will be sent on it.
import (
	"fmt"
	"time"
)

func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		// This runs inside a goroutine and waits for all the jobs to be completed.
		// it will keep running until it receives all the jobs. because it is a blocking operation.
		// the for loop will keep running until it receives all the jobs.
		for {
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	for j := 1; j <= 3; j++ {
		time.Sleep(1 * time.Second)
		jobs <- j
		fmt.Println("sent job", j)
	}
	close(jobs)
	fmt.Println("sent all jobs")

	<-done

	_, ok := <-jobs
	fmt.Println("Received more jobs:", ok)
}
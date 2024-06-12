package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

func main() {

	const numJobs = 15 
	const numWorkers = 3
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	for w := 1; w <= numWorkers; w++ {
		go worker(w, jobs, results)
	}
	
	for j := 1; j <= numJobs; j++ {
		time.Sleep(time.Millisecond * 100)
		jobs <- j
		fmt.Printf("%d job sent to the jobs channel \n", j)
	}
	close(jobs)

	for a := 1; a <= numJobs; a++ {
		<-results
	}
}
package main

import (
	"fmt"
	"sync"
	"time"
)

type Container struct {
	mu 			sync.Mutex
	counters 	map[string]int
}

func (c *Container) inc(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counters[key]++
}

func main() {
	c := Container{
		counters: map[string]int{
			"a": 0,
			"b": 0,
		},
	}

	var wg sync.WaitGroup

	doIncrement := func(name string, n int) {
		for i:=0; i<n; i++ {
			c.inc(name)
		}
		time.Sleep(1 * time.Second)
		fmt.Printf("done incrementing %v %d times\n", name, n)
		fmt.Println("counter info: ", c.counters)
		wg.Done()
	}

	wg.Add(3)
	/*
	In the given example, you have three goroutines concurrently trying to increment the values 
	in the counters map of the Container struct. The Container struct has a mutex (mu) to 
	synchronize access to the counters map, ensuring that only one goroutine can access or 
	modify the map at a time. This prevents race conditions and ensures the 
	integrity of the data.
	*/
	go doIncrement("a", 1000)
	go doIncrement("a", 1000)
	go doIncrement("b", 1000)

	wg.Wait()
	fmt.Println(c.counters)
}
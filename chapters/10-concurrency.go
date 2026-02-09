package main

import (
	"fmt"
	"sync"
	"time"
)

// worker simulates a task that takes time and reports its result via a channel
func worker(id int, wg *sync.WaitGroup, results chan<- string) {
	// Signal that this goroutine is done when the function exits
	defer wg.Done()

	fmt.Printf("Worker %d: starting task...\n", id)

	// Simulate work
	time.Sleep(time.Second * 2)

	// Send result to the channel
	results <- fmt.Sprintf("Worker %d: task complete", id)
}

func main() {
	// sync.WaitGroup is used to wait for all goroutines to finish
	var wg sync.WaitGroup

	// Buffered channel to collect results from workers
	// capacity of 3 matches our loop count
	results := make(chan string, 3)

	fmt.Println("Main: starting workers")

	for i := 1; i <= 3; i++ {
		// Increment the WaitGroup counter
		wg.Add(1)

		// Start a worker in a new goroutine
		go worker(i, &wg, results)
	}

	// Start a separate goroutine to wait for all workers
	// This prevents blocking the main thread while waiting
	go func() {
		wg.Wait()
		// Close the channel once all workers are done
		// This signals the range loop below to terminate
		close(results)
	}()

	// Receive results from the channel as they arrive
	// This loop continues until the channel is closed
	for res := range results {
		fmt.Println(res)
	}

	fmt.Println("Main: all workers finished")
}

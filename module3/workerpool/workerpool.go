package main

import (
	"fmt"
	"sync"
)

func worker(id int, jobs <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	for message := range jobs {
		fmt.Printf("Worker %d: %s\n", id, message)
	}

	fmt.Printf("Worker %d stopped\n", id)
}

func main() {

	n := 3 // number of workers

	jobs := make(chan string)
	var wg sync.WaitGroup

	// Create n workers
	wg.Add(n)

	for i := 1; i <= n; i++ {
		go worker(i, jobs, &wg)
	}

	// Send messages
	jobs <- "Hello"
	jobs <- "Go"
	jobs <- "Worker Pool"
	jobs <- "Jai Shri Ram"

	// Main closes the channel
	close(jobs)

	// Wait for all workers to stop
	wg.Wait()

	fmt.Println("Main stopped")
}

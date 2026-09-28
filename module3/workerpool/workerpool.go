package main

import (
	"fmt"
)

func worker(jobs <-chan int) {
	for {
		job := <-jobs
		fmt.Println("Worker received : ", job)
	}
}

func main() {
	jobs := make(chan int)

	go worker(jobs)

	jobs <- 100
	jobs <- 200
	jobs <- 300
	jobs <- 400
}

// func work(ind int, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	fmt.Println("worker ", ind)
// }

// func main() {
// 	n := 5
// 	var wg sync.WaitGroup
// 	wg.Add(n)
// 	for o := 1; o <= n; o++ {
// 		go work(o, &wg)
// 	}
// 	wg.Wait()
// }

// func worker(id int, jobs <-chan int, wg *sync.WaitGroup) {
// 	defer wg.Done()

// 	fmt.Println("Worker", id, "started")

// 	for job := range jobs {
// 		fmt.Println("Worker", id, "processed", job)
// 	}
// }

// func main() {
// 	n := 5

// 	jobs := make(chan int)

// 	var wg sync.WaitGroup

// 	// Create 5 workers
// 	wg.Add(n)

// 	for i := 1; i <= n; i++ {
// 		go worker(i, jobs, &wg)
// 	}

// 	// Send 10 jobs
// 	for i := 1; i <= 10; i++ {
// 		jobs <- i
// 	}

// 	// Tell workers there are no more jobs
// 	close(jobs)

// 	// Wait for all workers to finish
// 	wg.Wait()
// }

// func worker(id int, jobs <-chan string, wg *sync.WaitGroup) {
// 	defer wg.Done()

// 	for message := range jobs {
// 		fmt.Printf("Worker %d: %s\n", id, message)
// 	}

// 	fmt.Printf("Worker %d stopped\n", id)
// }

// func main() {

// 	n := 3 // number of workers

// 	jobs := make(chan string)
// 	var wg sync.WaitGroup

// 	// Create n workers
// 	wg.Add(n)

// 	for i := 1; i <= n; i++ {
// 		go worker(i, jobs, &wg)
// 	}

// 	// Send messages
// 	jobs <- "Hello"
// 	jobs <- "Go"
// 	jobs <- "Worker Pool"
// 	jobs <- "Jai Shri Ram"

// 	// Main closes the channel
// 	close(jobs)

// 	// Wait for all workers to stop
// 	wg.Wait()

// 	fmt.Println("Main stopped")
// }

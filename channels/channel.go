package main

import "sync"

func main() {
	count := 0
	var wg sync.WaitGroup
	var mu sync.Mutex
	wg.Add(1)

	go func() {
		defer wg.Done()
		for i := 0; i < 1000; i++ {
			mu.Lock()
			count++
			mu.Unlock()
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 1000; i++ {
			mu.Lock()
			count++
			mu.Unlock()
		}
	}()

	wg.Wait()
}

//1.
//sends genuinely blocks with no receiver .
// ch := make(chan string)
// fmt.Println("before sending msg to channel")
// ch <- "Hello, Channel!" // Send a message to the channel
// fmt.Println("never reached")
// 	By default (make(chan int)), a channel is unbuffered. This means:

// A send blocks until some other goroutine is ready to receive at that exact moment.
// A receive blocks until some other goroutine is ready to send.

//2.
// package main
// import "fmt"

// func main() {
// 	ch := make(chan int)
// 	go func() {
// 		ch <- 5 // now a receiver exists, so this can proceed
// 	}()
// 	value := <-ch
// 	fmt.Println("got:", value)
// }

//3. buffered channel , no deadlock msg
// ch := make(chan int, 3)
// ch <- 1
// ch <- 2
// ch <- 3
// fmt.Println("Channel length:", len(ch))

//4.
// close(ch) marks a channel as done. You can still receive remaining buffered values after closing, but you can never send again (sending on a closed channel panics).
// Receivers can detect closure with the two-value form:
// go
// v, ok := <-ch
// ok is false once the channel is closed and drained — meaning you've received every value that was sent before the close.
// Rule of thumb: only the sender should close a channel, never the receiver — and only close a channel once. Closing an already-closed channel panics.
// ch := make(chan int, 3)
// ch <- 1
// ch <- 2
// ch <- 3
// close(ch)
// ch <- 4
// v1, ok1 := <-ch
// fmt.Println(v1, ok1)
// v2, ok2 := <-ch
// fmt.Println(v2, ok2)
// v3, status := <-ch
// fmt.Println(v3, status)
// v4, status := <-ch
// fmt.Println(v4, status)

//5. range over channel
// ch := make(chan string, 3)
// 	ch <- "Hello, Channel!" // Send a message to the channel
// 	ch <- "Bolo ! Jai Shri Ram"
// 	ch <- "Jai Shree Ram"
// 	close(ch)

// 	for val := range ch {
// 		println(val)
// 	}

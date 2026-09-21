package main

import (
	"fmt"
	"sync"
)

func main() {
	count := 0
	var wg sync.WaitGroup
	var mu sync.Mutex
	wg.Add(3)

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
	fmt.Println("count", count)
}

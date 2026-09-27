package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(4)
	go func() {
		fmt.Println(2)
		wg.Done()
		wg.Done()
	}()

	go func() {
		fmt.Println(4)
		wg.Done()
	}()
	wg.Wait()
	fmt.Println(6)
}

//1. Example of race condition.
// count := 0
// go func() {
// 	for i := 0; i < 1000; i++ {
// 		count++
// 	}
// }()

// go func() {
// 	for i := 0; i < 1000; i++ {
// 		count++
// 	}
// }()
// fmt.Println("count", count)

// 2.example of waitgroup(mechansim to wait for multiple goroutines to finish executing)
// 	count := 0
// 	var wg sync.WaitGroup
// 	wg.Add(2)

// 	go func() {
// 		for i := 0; i < 1000; i++ {
// 			count++
// 		}
// 		defer wg.Done()
// 	}()

// 	go func() {
// 		for i := 0; i < 1000; i++ {
// 			count++
// 		}
// 		defer wg.Done()
// 	}()
// 	wg.Wait()
// 	fmt.Println("count", count)

//3. Example of mutex lock to avoid race condition
// count := 0
// var wg sync.WaitGroup
// var mu sync.Mutex
// wg.Add(2)

// go func() {
// 	for i := 0; i < 1000; i++ {
// 		mu.Lock()
// 		count++
// 		mu.Unlock()
// 	}
// 	defer wg.Done()
// }()

// go func() {
// 	for i := 0; i < 1000; i++ {
// 		mu.Lock()
// 		count++
// 		mu.Unlock()
// 	}
// 	defer wg.Done()
// }()

// wg.Wait()
// fmt.Println("count", count)

//4.sync.cond
// var mu sync.Mutex
// cond := sync.NewCond(&mu)
// dataready := false
// data := ""

// go func() {
// 	mu.Lock()
// 	for !dataready {
// 		cond.Wait()
// 	}
// 	mu.Unlock()
// 	fmt.Println("data consumed by consumer : ", data)
// }()

// go func() {
// 	time.Sleep(2 * time.Second)
// 	mu.Lock()
// 	data = "Jai Shree Ram"
// 	fmt.Println("data produced by producer : ", data)
// 	dataready = true
// 	cond.Signal()
// 	mu.Unlock()

// }()
// fmt.Println("first statement to be printed")
// time.Sleep(4 * time.Second)
// fmt.Println("last statement to be printed")

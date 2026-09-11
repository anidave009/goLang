package main

import (
	"fmt"
	"sync"
	"time"
)

func task_goroutineWaitGroup(name string, wg *sync.WaitGroup) {
	defer wg.Done() // Decrement the counter when the goroutine completes
	for i := 0; i <= 5; i++ {
		fmt.Println(name, ":", i)
		time.Sleep(2 * time.Second)
	}
}

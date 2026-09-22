package main

import (
	"fmt"
	"sync"
	"time"
)

type MyWaitGroup struct {
	mu    sync.Mutex
	cond  *sync.Cond
	count int
}

func NewMyWaitGroup() *MyWaitGroup {
	wg := &MyWaitGroup{}
	wg.cond = sync.NewCond(&wg.mu)
	return wg
}

func (wg *MyWaitGroup) Add(delta int) {
	wg.mu.Lock()
	defer wg.mu.Unlock()
	wg.count += delta
	if wg.count < 0 {
		panic("MyWaitGroup: negative counter")
	}
	if wg.count == 0 {
		wg.cond.Broadcast()
	}
}

func (wg *MyWaitGroup) Done() {
	wg.Add(-1)
}

func (wg *MyWaitGroup) Wait() {
	wg.mu.Lock()
	defer wg.mu.Unlock()
	for wg.count > 0 {
		wg.cond.Wait()
	}
}

func main() {
	wg := NewMyWaitGroup()

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			time.Sleep(time.Duration(id) * 2 * time.Second)
			fmt.Printf("worker %d done\n", id)
		}(i)
	}

	fmt.Println("waiting for all workers...")
	wg.Wait()
	fmt.Println("all workers finished")
}

//Step 1 . Understand what waitgroup does .
// count := 0
// var wg sync.WaitGroup
// var mu sync.Mutex
// wg.Add(3)

// go func() {
// 	defer wg.Done()
// 	for i := 0; i < 1000; i++ {
// 		mu.Lock()
// 		count++
// 		mu.Unlock()
// 	}
// }()

// go func() {
// 	defer wg.Done()
// 	for i := 0; i < 1000; i++ {
// 		mu.Lock()
// 		count++
// 		mu.Unlock()
// 	}
// }()
// wg.Wait()
// fmt.Println("count", count)

//Step 2 : Waitgroup handled with help of mutex lock ,
//so it results in busy waiting & unnexessary polling & cpu usage .
// package main

// import (
// 	"fmt"
// 	"sync"
// 	"time"
// )

// type MyWaitGroup struct {
// 	mu    sync.Mutex
// 	count int
// }

// // constructor like function in go
// func NewMyWaitGroup() *MyWaitGroup {
// 	return &MyWaitGroup{}
// }

// func (wg *MyWaitGroup) Add(delta int) {
// 	wg.mu.Lock()
// 	defer wg.mu.Unlock()
// 	wg.count += delta
// 	if wg.count < 0 {
// 		panic("MyWaitGrp : negative counter")
// 	}
// }

// func (wg *MyWaitGroup) Done() {
// 	wg.Add(-1)
// }

// func (wg *MyWaitGroup) Wait() {
// 	for {
// 		wg.mu.Lock()
// 		if wg.count == 0 {
// 			wg.mu.Unlock()
// 			return
// 		}
// 		wg.mu.Unlock()
// 		// briefly yield so this loop doesn't hog the CPU core completely;
// 		// still "busy waiting" in spirit, just slightly less wasteful
// 		time.Sleep(time.Millisecond)
// 	}
// }

// func main() {
// 	wg := NewMyWaitGroup()

// 	for i := 1; i <= 5; i++ {
// 		wg.Add(1)
// 		go func(id int) {
// 			defer wg.Done()
// 			time.Sleep(time.Duration(id) * 10 * time.Second)
// 			fmt.Printf("worker %d done\n", id)
// 		}(i)
// 	}

// 	fmt.Println("waiting for all workers...")
// 	wg.Wait()
// 	fmt.Println("all workers finished")
// }

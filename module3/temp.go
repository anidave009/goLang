package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(ind int, job <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for val := range job {
		fmt.Println("val : ", val)
		fmt.Println("worker id : ", ind)
		time.Sleep(1000 * time.Millisecond)
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(4)
	job := make(chan int)
	for i := 1; i <= 4; i++ {
		go worker(i, job, &wg)
	}
	job <- 111
	job <- 222
	job <- 333
	job <- 444
	job <- 555
	job <- 666
	job <- 777
	job <- 888

	close(job)
	wg.Wait()
	fmt.Println("main ends")
}

// 1.
// 	ch := make(chan int)
// 	ch <- 23
// 	val := <-ch
// 	fmt.Println(val)

//2.
// ch := make(chan int)
// val := <-ch
// ch <- 133
// fmt.Println(val)

//3.
// ch := make(chan int)
// ch <- 24
// go func() {
// 	val := <-ch
// 	fmt.Println(val)
// }()

//4 .
// ch := make(chan int)

// go func() {
// 	val := <-ch
// 	fmt.Println(val)
// }()

// ch <- 24

//5.
// ch := make(chan int)
// val := <-ch

// go func() {
// 	ch <- 211
// }()

// fmt.Println(val)

package main

import (
	"fmt"
	"time"
)

func task_sleep(name string) {
	for i := 0; i < 5; i++ {
		fmt.Println(name, ":", i)
		time.Sleep(1 * time.Second)
	}
}

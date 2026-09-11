package main

import (
	"fmt"
	"strings"
)

func main() {

	var slice_A = []int{}
	fmt.Println(len(slice_A))
	fmt.Println(cap(slice_A))
	fmt.Println(slice_A)

	var slice_1 = []int{1, 2, 3, 4, 5}
	fmt.Println(len(slice_1))
	fmt.Println(cap(slice_1))
	fmt.Println(slice_1)

	var arr = [5]int{100, 101, 102, 103, 104}
	var slice_2 = arr[1:4]
	fmt.Println(len(slice_2))
	fmt.Println(cap(slice_2))
	fmt.Println(slice_2)

	fmt.Println(" Hello")

	s := "hello"
	fmt.Println(s[0])

	words := strings.Fields("Go makes programming fun")

	names := make([]string, 0)

	names = append(names, "Alice")
	names = append(names, "Bob")

	fmt.Println(names) // [Alice Bob]

	for _, word := range words {
		fmt.Println(word)
	}
}

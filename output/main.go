package main

import (
	"fmt"
)

func main() {
	var i, j string = "Hello", "World!"

	fmt.Print(i)
	fmt.Print(j)

	fmt.Print(i, "\n")
	fmt.Print(j, "\n")

	fmt.Print(i, "\n", j)

	fmt.Print(i, " ", j)

	var a, b int = 1, 2
	fmt.Print(a, b, "\n")

	fmt.Printf("i has value: %v and type: %T and value : %#v\n", i, i, i)
	fmt.Printf("a has value: %v and type: %T\n", a)
}

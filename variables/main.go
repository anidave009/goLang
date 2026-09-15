package main

import (
	"fmt"
	"gopractice/variables/maths"
)

func main() {

	// How to declare variables in Go:
	//var variable_name type = value

	var age int = 23
	var name string = "Aniket"

	name1 := "Dave"
	fmt.Println(age)
	fmt.Println(name)
	fmt.Println(name1)

	var age1 int = 22
	age2 := 19 //inferred type

	fmt.Println(age1)
	fmt.Println(age2)

	fmt.Println(maths.Add(2, 3))

	var (
		a int
		b string = "Shri ram"
		c float32
		d bool
		e int    = 23
		f string = "I"
	)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)

	var E, F = 23, "I"
	fmt.Println(E)
	fmt.Println(F)

	T, Z := 45, "Hello"
	fmt.Println(T)
	fmt.Println(Z)

}

package main

import "fmt"

func main() {
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
}

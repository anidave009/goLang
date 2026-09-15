package main

import "fmt"

func main() {

	var arr [3]int
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3

	fmt.Println(arr)

	var arr1 = [3]int{1, 2, 3}
	fmt.Println(arr1)

	var arr2 = [...]int{1, 2, 3, 4, 5}
	fmt.Println(arr2)

	arr3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(arr3)

	var arr4 = [5]int{}
	for i := 0; i < len(arr3); i++ {
		arr4[i] = arr3[i] * 2
	}
	fmt.Println(arr4)
}

// var arr = [4]int{1, 2, 3, 4}          //specifying the size of the array
// var arr1 = [...]int{1, 2, 3, 4, 5, 6} //not specifying the size of the array, compiler will determine the size of the array based on the number of elements in the array

// var tmp [3]int //declaring an array of size 3 and initializing it with default values of 0
// tmp[0] = 10

// arr2 := [5]int{1, 2, 3, 4, 5}      //without using the var keyword, we can declare and initialize an array in a single line using the := operator
// arr3 := [...]int{1, 2, 3, 4, 5, 6} //without using the var keyword, we can declare and initialize an array in a single line using the := operator and not specifying the size of the array, compiler will determine the size of the array based on the number of elements in the array
// fmt.Println(arr)
// fmt.Println(arr1)
// fmt.Println(arr2)
// fmt.Println(arr3)

// temp := [3]int{}
// fmt.Println(temp)

// temp1 := [...]int{0: 20, 2: 10, 6: 30} //particular index of the array can be initialized with a value, if the index is not specified, the value will be assigned to the next available index
// fmt.Println(temp1)

// for i := 0; i < len(temp1); i++ { //looping through the array using the length of the array
// 	fmt.Print(temp1[i], " ")
// }
// fmt.Println("/////")
// fmt.Println(tmp)

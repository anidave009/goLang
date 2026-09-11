package main

func main() {
	//Task 1 with goroutines using sleep function
	// go task_sleep("Task A")
	// go task_sleep("Task B")
	// fmt.Println("Main function is running concurrently with tasks")
	// time.Sleep(3 * time.Second) // Wait for goroutines to finish //so we let the main function sleep for 3 seconds to allow the goroutines to complete their execution before the program exits
	// fmt.Println("Main function completed")

	// Task 2 with goroutines using WaitGroup
	// var wg sync.WaitGroup
	// wg.Add(1) // Set the number of goroutines to wait for

	// go task_goroutineWaitGroup("Task A", &wg)
	// go task_goroutineWaitGroup("Task B", &wg)

	// fmt.Println("Main function is running concurrently with tasks")

	// wg.Wait() // Wait for all goroutines to finish
	// fmt.Println("Main function completed")

	//Task 3 with goroutines & channels
	// result := make(chan string) // Create a channel to receive results from goroutines
	// go gochanneltask(result)
	// message := <-result  // Receive the message from the channel
	// fmt.Println(message) // Print the received message

	// message2 := <-result  // Receive the second message from the channel
	// fmt.Println(message2) // Print the second received message

	//Task 4 : Send Numbers (1-5) from a channel.
	// numbers := make(chan int)
	// go sendNumbers(numbers)
	// for i := 0; i < 5; i++ {
	// 	num := <-numbers
	// 	fmt.Println(num)
	// }

	// //Task 5 : Calculate Product of two numbers using channels
	// productChannel := make(chan int)
	// go calcProduct(productChannel, 5, 10)
	// product := <-productChannel
	// fmt.Println(product)

	// //Task 6 : Create 2 goroutines & let them coordinate via one channels
	// messageChannel := make(chan string)
	// go sendMessage(messageChannel, "Hello from goroutine 1")
	// go sendMessage(messageChannel, "Hello from goroutine 2")
	// message1 := <-messageChannel
	// message2 := <-messageChannel
	// fmt.Println(message1)
	// fmt.Println(message2)

	//Unbuffered channel
	// ch := make(chan int)
	// ch <- 10
	// num := <-ch
	// fmt.Println(num)

	//ubfchnl with goroutine
	// ubffCh := make(chan int)
	// go func() {
	// 	ubffCh <- 20
	// }()
	// num1 := <-ubffCh
	// fmt.Println(num1)

	//Waiting to receive but no one has sent anything so deadlock
	// ch := make(chan int)
	// fmt.Println("Waiting to receive")
	// num2 := <-ch
	// fmt.Println(num2)

	//with help of another goroutine one can receive without deadlock
	// ch2 := make(chan string)
	// go func() {
	// 	ch2 <- "Jai Shri Ram"
	// }()

	// stg := <-ch2
	// fmt.Println(stg)
}

//Output without goroutines:
// Task A : 0
// Task A : 1
// Task A : 2
// Task A : 3
// Task A : 4
// Task B : 0
// Task B : 1
// Task B : 2
// Task B : 3
// Task B : 4

// Output with goroutines:
// Task B : 0
// Task A : 0
// Task A : 1
// Task B : 1
// Task B : 2
// Task A : 2
// Task A : 3
// Task B : 3
// Task A : 4
// Task B : 4
// Main function completed

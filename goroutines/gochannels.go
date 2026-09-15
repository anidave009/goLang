package main

func gochanneltask(result chan string) {
	result <- "Task completed" // Send a message to the channel
	result <- "Bolo ! Jai Shri Ram"
}

func sendNumbers(numbers chan int) {
	for i := 11; i <= 15; i++ {
		numbers <- i // Send numbers 1 to 5 to the channel
	}
}

func calcProduct(productChannel chan int, x int, y int) {
	product := x * y
	productChannel <- product // Send the product to the channel
}

func sendMessage(msgChannel chan string, message string) {
	msgChannel <- message // Send a message to the channel
}

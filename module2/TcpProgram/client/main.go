package main

import (
	"fmt"
	"net"
)

func main() {
	// Connect to the TCP server
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}

	defer conn.Close()

	// Message to send
	message := "Jai Shri Ram"

	// Send message to server
	_, err = conn.Write([]byte(message))
	if err != nil {
		fmt.Println("Error sending message:", err)
		return
	}

	fmt.Println("Message sent:", message)

	// Buffer to store server response
	buffer := make([]byte, 1024)

	// Read server response
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	// Print the response
	response := string(buffer[:n])

	fmt.Println("Server response:")
	fmt.Println(response)
}

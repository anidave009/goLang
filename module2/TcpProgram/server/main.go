package main

import (
	"fmt"
	"net"
	"strings"
	"time"
)

func main() {
	// Create TCP server and listen on port 8080
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}

	defer listener.Close()

	fmt.Println("Server is listening on port 8080...")

	// Keep accepting new clients
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		// Handle each client concurrently
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	fmt.Println("Client connected:", conn.RemoteAddr())
	time.Sleep(10 * time.Second)

	// Buffer to store incoming data
	buffer := make([]byte, 1024)

	// Read data sent by the client
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error reading from client:", err)
		return
	}

	// Convert received bytes into a string
	message := string(buffer[:n])

	fmt.Println("Received:", message)

	// Create the response
	response := strings.ToUpper(message) + "\n" +
		message + "\n" +
		strings.ToLower(message)

	// Send response back to client
	_, err = conn.Write([]byte(response))
	if err != nil {
		fmt.Println("Error writing to client:", err)
		return
	}

	fmt.Println("Response sent")
}

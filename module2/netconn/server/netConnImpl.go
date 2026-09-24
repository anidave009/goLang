package main

import (
	"fmt"
	"net"
	"time"
)

// Our custom connection
type MyConn struct {
	conn net.Conn
}

// Implement net.Conn methods

func (c MyConn) Read(b []byte) (int, error) {

	// Read the client's request first
	n, err := c.conn.Read(b)
	if err != nil {
		return n, err
	}

	// Get the second when request was received
	second := time.Now().Second()

	var result string

	if second%2 == 0 {
		result = "EVEN"
	} else {
		result = "ODD"
	}

	// Replace the received data with our result
	copy(b, result)

	return len(result), nil
}

func (c MyConn) Write(b []byte) (int, error) {
	return c.conn.Write(b)
}

func (c MyConn) Close() error {
	return c.conn.Close()
}

func (c MyConn) LocalAddr() net.Addr {
	return c.conn.LocalAddr()
}

func (c MyConn) RemoteAddr() net.Addr {
	return c.conn.RemoteAddr()
}

func (c MyConn) SetDeadline(t time.Time) error {
	return c.conn.SetDeadline(t)
}

func (c MyConn) SetReadDeadline(t time.Time) error {
	return c.conn.SetReadDeadline(t)
}

func (c MyConn) SetWriteDeadline(t time.Time) error {
	return c.conn.SetWriteDeadline(t)
}

func main() {

	// Create TCP listener
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer listener.Close()

	fmt.Println("Server listening on port 8080...")

	for {
		// Get the actual TCP connection
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}

		// Wrap it with our custom connection
		myConn := MyConn{
			conn: conn,
		}

		go handleConnection(myConn)
	}
}

func handleConnection(conn net.Conn) {

	defer conn.Close()

	buffer := make([]byte, 1024)

	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Result:", string(buffer[:n]))

	// Send result back to client
	conn.Write(buffer[:n])
}

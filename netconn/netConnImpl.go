package main

import (
	"net"
	"time"
)

type MyConn struct{}

func (c MyConn) Read(b []byte) (int, error) {
	return 0, nil
}

func (c MyConn) Write(b []byte) (int, error) {
	return 0, nil
}

func (c MyConn) Close() error {
	return nil
}

func (c MyConn) LocalAddr() net.Addr {
	return nil
}

func (c MyConn) RemoteAddr() net.Addr {
	return nil
}

func (c MyConn) SetDeadline(t time.Time) error {
	return nil
}

func (c MyConn) SetReadDeadline(t time.Time) error {
	return nil
}

func (c MyConn) SetWriteDeadline(t time.Time) error {
	return nil
}

func main() {
	conn := MyConn{}

	// Prove that MyConn implements net.Conn
	var networkConn net.Conn = conn

	_ = networkConn
}

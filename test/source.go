package main

import (
	"fmt"
	"net"
	"time"
)

func handleError(err error, message string) {
	if err != nil {
		fmt.Println(message)
	}
}

func handleClient(c net.Conn) {
	// Wait 1 second and send reply "X"
	time.Sleep(time.Second * 1)
	c.Write([]byte("X"))
}

func main() {
	// Open TCP port 1234
	listener, err := net.Listen("tcp", ":1234")
	handleError(err, "ERROR")

	for {
		// Wait for connection
		conn, err := listener.Accept()
		handleError(err, "ERROR")

		// Fork handler
		go handleClient(conn)
	}
}

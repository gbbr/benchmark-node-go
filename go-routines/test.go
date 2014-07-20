package main

import (
	"fmt"
	"net"
)

const CONNECTIONS = 200

func handleError(err error, message string) {
	if err != nil {
		fmt.Println(message)
	}
}

func readResponse(conn net.Conn, solved chan byte) {
	buf := make([]byte, 1024)
	conn.Read(buf)
	fmt.Printf("%s", buf)
	solved <- 1
}

func main() {
	var count byte = 0
	solved := make(chan byte)

	for i := 0; i < CONNECTIONS; i++ {
		conn, err := net.Dial("tcp", "127.0.0.1:1234")
		handleError(err, "Unable to connect")
		go readResponse(conn, solved)
	}

	for {
		count += <-solved
		if count == CONNECTIONS {
			break
		}
	}
}

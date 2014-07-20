package main

import (
	"fmt"
	"net"
	"time"
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
	var count uint32 = 0
	solved := make(chan byte)
	startTime := time.Now()

	for i := 0; i < CONNECTIONS; i++ {
		conn, err := net.Dial("tcp", "127.0.0.1:1234")
		handleError(err, "Unable to connect")
		go readResponse(conn, solved)
	}

	for {
		count += uint32(<-solved)
		if count == CONNECTIONS {
			break
		}
	}

	fmt.Printf("Time: %s\n", time.Now().Sub(startTime).String())
}

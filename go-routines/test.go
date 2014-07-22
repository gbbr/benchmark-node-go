package main

import (
	"fmt"
	"net"
	"sync"
)

const CONNECTIONS = 200

var w sync.WaitGroup

func handleError(err error, message string) {
	if err != nil {
		fmt.Println(message)
	}
}

func readResponse(conn net.Conn, solved chan byte) {
	buf := make([]byte, 1024)
	conn.Read(buf)
	fmt.Printf("%s", buf)
	w.Done()
}

func main() {
	w.Add(CONNECTIONS)

	solved := make(chan byte)

	for i := 0; i < CONNECTIONS; i++ {
		conn, err := net.Dial("tcp", "127.0.0.1:1234")
		handleError(err, "Unable to connect")
		go readResponse(conn, solved)
	}

	w.Wait()
}

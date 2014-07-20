package main

import (
	"net"
	"time"
)

func handleClient(c net.Conn) {
	time.Sleep(time.Second * 1)
	c.Write([]byte("X"))
	c.Close()
}

func main() {
	listener, _ := net.Listen("tcp", ":1234")
	for {
		conn, _ := listener.Accept()
		go handleClient(conn)
	}
}

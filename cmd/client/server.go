package main

import (
	"fmt"
	"net"
)

const (
	addr = "127.0.0.1:7000"
)

func main() {
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			panic(err)
		}
		go handler(conn)
	}
}

func handler(conn net.Conn)  {
	conn.Write([]byte("hello world"))
	buffer := make([]byte, 1024)
	c, err := conn.Read(buffer)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(buffer[0:c]), conn.RemoteAddr())
	conn.Close()
}

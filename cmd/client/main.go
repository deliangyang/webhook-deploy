package main

import (
	"fmt"
	"net"
)

const (
	address = "127.0.0.1:7000"
)

func main() {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		panic(err)
	}

	defer conn.Close()
	buffer := make([]byte, 1024)
	c, err := conn.Read(buffer)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(buffer[0:c]))
	conn.Write([]byte("hello world, from client"))

}
package main

import (
	"fmt"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":4040")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		conn.Write([]byte("Nice to meet you!"))
		conn.Close()
	}

}

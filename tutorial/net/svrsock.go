package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	listener, err := net.Listen("tcp", ":4040")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer listener.Close()
	fmt.Println("echo server running on 4040...")

	// connection loop
	for {
		// wait and accept new client connection
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}

		// handle connection
		// read text and echo it
		go func(c net.Conn) {
			defer conn.Close()
			reader := bufio.NewReader(c)
			for {
				line, err := reader.ReadString('\n')
				if err != nil {
					fmt.Println(err)
					return
				}
				if _, err := fmt.Fprint(conn, line); err != nil {
					fmt.Println(err)
					return
				}
			}
		}(conn)
	}
}

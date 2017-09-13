package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

var host, port = "127.0.0.1", "4040"
var addr = net.JoinHostPort(host, port)

func main() {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer conn.Close()
	fmt.Println("Connected to echo server at:", addr)
	fmt.Println("Press ctrl-c to exit")

	for {
		cmdReader := bufio.NewReader(os.Stdin)
		conReader := bufio.NewReader(conn)

		// read command-line input
		data, err := cmdReader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			continue
		}

		// send data to server
		if _, err := fmt.Fprint(conn, data); err != nil {
			fmt.Println(err)
			continue
		}

		// read server response
		resp, err := conReader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Print(resp)
	}
}

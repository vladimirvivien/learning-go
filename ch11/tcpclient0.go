package main

import (
	"fmt"
	"net"
	"time"
)

var host, port = "127.0.0.1", "4040"
var addr = net.JoinHostPort(host, port)
var deadline = time.Now().Add(time.Millisecond * 700)

const prompt = "curr"
const buffLen = 1024

func main() {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()
	fmt.Println("Connected to Global Currency Service")
	var cmd, param string

	// repl
	for {
		fmt.Print(prompt, "> ")
		_, err = fmt.Scanf("%s %s", &cmd, &param)
		if err != nil {
			fmt.Println("Usage: GET <search string or *>")
			continue
		}
		// send command line
		cmdLine := fmt.Sprintf("%s %s", cmd, param)
		if n, err := conn.Write([]byte(cmdLine)); n == 0 || err != nil {
			fmt.Println(err)
			return
		}

		// stream and display response
		//conbuf := bufio.NewReader(conn)
		conn.SetReadDeadline(time.Now().Add(time.Millisecond * 5000))
		for {
			buff := make([]byte, buffLen)
			n, err := conn.Read(buff)
			if err != nil {
				break
			}
			fmt.Print(string(buff[0:n]))
			conn.SetReadDeadline(time.Now().Add(time.Millisecond * 700))
		}
	}

}

package main

import (
	"fmt"
	"net"
	"os"
	"strings"
	"time"

	curr "github.com/vladimirvivien/learning-go/ch11/curr0"
)

var currencies = curr.Load("./data.csv")

func main() {
	ln, err := net.Listen("tcp", ":4040")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer ln.Close()
	fmt.Println("Global Currency Service... Listening on port 4040")

	// connection-loop - handle incoming requests
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			conn.Close()
			continue
		}
		fmt.Println("Connected to ", conn.RemoteAddr())
		// delegate connection to a goroutine
		go handleConnection(conn)
	}
}

// handle client connection
func handleConnection(conn net.Conn) {
	defer conn.Close()

	// loop to stay connected with client
	for {
		cmdLine := make([]byte, (1024 * 4))
		n, err := conn.Read(cmdLine)
		if n == 0 || err != nil {
			return
		}
		cmd, param := parseCommand(string(cmdLine[0:n]))
		if cmd == "" {
			continue
		}

		// execute command
		switch strings.ToUpper(cmd) {
		case "GET":
			result := curr.Find(currencies, param)
			if len(result) == 0 {
				conn.Write([]byte("Nothing found\n"))
				continue
			}
			// stream result to client
			for _, cur := range result {
				_, err := fmt.Fprintf(
					conn,
					"%s %s %s %s\n",
					cur.Name, cur.Code, cur.Number, cur.Country,
				)
				if err != nil {
					return
				}

				// reset deadline while writing,
				// causes server to close conn if client is gone
				conn.SetWriteDeadline(time.Now().Add(time.Second * 5))
			}
			// reset read deadline for next read
			conn.SetWriteDeadline(time.Time{})

		default:
			conn.Write([]byte("Invalid command\n"))
		}
	}
}

func parseCommand(cmdLine string) (cmd, param string) {
	parts := strings.Split(cmdLine, " ")
	if len(parts) != 2 {
		return "", ""
	}
	cmd = strings.TrimSpace(parts[0])
	param = strings.TrimSpace(parts[1])
	return
}

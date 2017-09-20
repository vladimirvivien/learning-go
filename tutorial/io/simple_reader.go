package main

import (
	"fmt"
	"io"
	"strings"
	"os"
)

func main() {
	reader := strings.NewReader("Clear is better than clever")
	p := make([]byte, 4)
	
	for {
		n, err := reader.Read(p)
		if err != nil{
		    if err == io.EOF {
			fmt.Println(string(p[:n]))
			break
		    }
		    fmt.Pritnln(err)
		    os.Exit(1)
		}
		fmt.Println(string(p[:n]))
	}
}

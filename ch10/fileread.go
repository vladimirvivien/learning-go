package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fin, err := os.Open("../ch05/dict.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer fin.Close()
	p := make([]byte, 1024)
	for {
		n, err := fin.Read(p)
		if err == io.EOF {
			break
		}
		fmt.Print(string(p[:n]))
	}
}

package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	f1, err := os.Open("./file0.go")
	if err != nil {
		fmt.Println("Unable to open file:", err)
		os.Exit(1)
	}
	defer f1.Close()

	n, err := io.Copy(os.Stdout, f1)
	if err != nil {
		fmt.Println("Failed to copy:", err)
		os.Exit(1)
	}

	fmt.Printf("Copied %d bytes from %s \n", n, f1.Name())
}

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

	f2, err := os.Create("./file0.bkp")
	if err != nil {
		fmt.Println("Unable to create file:", err)
		os.Exit(1)
	}
	defer f2.Close()

	n, err := io.Copy(f2, f1)
	if err != nil {
		fmt.Println("Failed to copy:", err)
		os.Exit(1)
	}

	fmt.Printf("Copied %d bytes from %s to %s\n", n, f1.Name(), f2.Name())
}

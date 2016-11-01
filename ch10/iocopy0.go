package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	data := strings.NewReader("Write me down.")
	file, err := os.Create("./iocopy.data")
	if err != nil {
		fmt.Println("Unable to create file:", err)
		os.Exit(1)
	}
	io.Copy(file, data)
}

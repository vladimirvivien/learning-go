package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Create("./iopipe.data")
	if err != nil {
		fmt.Println("File creation failed:", err)
		os.Exit(1)
	}

	pr, pw := io.Pipe() // pipe reader/writer
	go func() {
		fmt.Fprint(pw, "Pipe streaming")
		pw.Close()
	}()

	wait := make(chan struct{})
	go func() {
		io.Copy(file, pr)
		pr.Close()
		close(wait)
	}()
	<-wait
}

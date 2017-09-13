package main

import (
	"bytes"
	"io"
	"os"
)

func main() {
	proverbs := new(bytes.Buffer)
	proverbs.WriteString("Channels orchestrate mutexes serialize\n")
	proverbs.WriteString("Cgo is not Go\n")
	proverbs.WriteString("Errors are values\n")
	proverbs.WriteString("Don't panic\n")

	piper, pipew := io.Pipe()

	// write in writer end of pipe
	go func() {
		defer pipew.Close()
		io.Copy(pipew, proverbs)
	}()

	// read from reader end of pipe.
	io.Copy(os.Stdout, piper)
	piper.Close()
}

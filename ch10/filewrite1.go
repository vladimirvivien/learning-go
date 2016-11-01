package main

import (
	"fmt"
	"os"
)

func main() {
	data := [][]byte{
		[]byte("The quick brown fox\n"),
		[]byte("jumps over the lazy dog\n"),
	}
	fout, err := os.Create("./filewrite.data")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer fout.Close()
	for _, out := range data {
		fout.Write(out)
	}
}

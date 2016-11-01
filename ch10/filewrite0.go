package main

import (
	"fmt"
	"os"
)

func main() {
	rows := []string{
		"The quick brown fox",
		"jumps over the lazy dog",
	}

	fout, err := os.Create("./filewrite.data")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer fout.Close()

	for _, row := range rows {
		fout.WriteString(row)
	}
}

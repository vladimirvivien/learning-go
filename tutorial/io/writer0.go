package main

import (
	"bytes"
	"fmt"
)

func main() {
	proverbs := []string{
		"Channels orchestrate mutexes serialize",
		"Cgo is not Go",
		"Errors are values",
		"Don't panic",
	}

	// bytes.Buffer an in-mem io.Writer
	buf := new(bytes.Buffer)

	for i, p := range proverbs {
		fmt.Fprintf(buf, "Proverb %d = %s\n", i, p)
	}
	fmt.Println(buf.String())
}

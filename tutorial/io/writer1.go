package main

import (
	"fmt"
	"os"
)

func main() {
	proverbs := []string{
		"Channels orchestrate mutexes serialize",
		"Cgo is not Go",
		"Errors are values",
		"Don't panic",
	}

	for i, p := range proverbs {
		fmt.Fprintf(os.Stdout, "Proverb %d = %s\n", i, p)
	}
}

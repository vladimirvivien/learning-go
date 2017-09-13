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

	file, err := os.Create("proverbs.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	for i, p := range proverbs {
		fmt.Fprintf(file, "Proverb %d = %s\n", i, p)
	}
	fmt.Println("proverbs.txt created")
}

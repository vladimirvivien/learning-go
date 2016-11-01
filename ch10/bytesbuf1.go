package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	var books bytes.Buffer
	books.WriteString("The Great Gatsby\n")
	books.WriteString("1984\n")
	books.WriteString("A Take of Two Cities\n")
	books.WriteString("Les Miserables\n")
	books.WriteString("The Call of the Wild\n")

	file, err := os.Create("./books.txt")
	if err != nil {
		fmt.Println("Unable to create file:", err)
		return
	}
	defer file.Close()
	books.WriteTo(file)
}

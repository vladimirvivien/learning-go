package main

import (
	"bytes"
	"os"
)

func main() {
	var books bytes.Buffer
	books.WriteString("The Great Gatsby")
	books.WriteString("1984")
	books.WriteString("A Tale of Two Cities")
	books.WriteString("Les Miserables")
	books.WriteString("The Call of the Wild")

	books.WriteTo(os.Stdout)
}

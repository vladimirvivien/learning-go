package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"
)

type Name struct {
	First, Last string
}

func (n *Name) UnmarshalJSON(data []byte) error {
	var name string
	err := json.Unmarshal(data, &name)
	if err != nil {
		fmt.Println(err)
		return err
	}
	parts := strings.Split(name, ", ")
	n.Last, n.First = parts[0], parts[1]
	return nil
}

type Book struct {
	Title       string
	PageCount   int
	ISBN        string
	Authors     []Name
	Publisher   string
	PublishDate time.Time
}

func main() {
	file, err := os.Open("book.dat")
	if err != nil {
		fmt.Println(err)
		return
	}

	var books []Book
	dec := json.NewDecoder(file)
	if err := dec.Decode(&books); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(books)

}

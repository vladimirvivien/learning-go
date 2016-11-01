package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type Name struct {
	First, Last string
}

func (n *Name) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s, %s\"", n.Last, n.First)), nil
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
	books := []Book{
		Book{
			Title:       "Leaning Go",
			PageCount:   375,
			ISBN:        "9781784395438",
			Authors:     []Name{{"Vladimir", "Vivien"}},
			Publisher:   "Packt",
			PublishDate: time.Date(2016, time.July, 0, 0, 0, 0, 0, time.UTC),
		},
		Book{
			Title:       "The Go Programming Language",
			PageCount:   380,
			ISBN:        "9780134190440",
			Authors:     []Name{{"Alan", "Donavan"}, {"Brian", "Kernighan"}},
			Publisher:   "Addison-Wesley",
			PublishDate: time.Date(2015, time.October, 26, 0, 0, 0, 0, time.UTC),
		},
		Book{
			Title:       "Introducing Go",
			PageCount:   124,
			ISBN:        "978-1491941959",
			Authors:     []Name{{"Caleb", "Doxsey"}},
			Publisher:   "O'Reilly",
			PublishDate: time.Date(2016, time.January, 0, 0, 0, 0, 0, time.UTC),
		},
	}

	file, err := os.Create("book.dat")
	if err != nil {
		fmt.Println(err)
		return
	}
	enc := json.NewEncoder(file)
	if err := enc.Encode(books); err != nil {
		fmt.Println(err)
	}
}

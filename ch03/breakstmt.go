package main

import (
	"fmt"
)

var words = [][]string{
	{"break", "lake", "go", "right", "strong", "kite", "hello"},
	{"fix", "river", "stop", "left", "weak", "flight", "bye"},
	{"fix", "lake", "slow", "middle", "sturdy", "high", "hello"},
}

func search(w string) {
DoSearch:
	for i := 0; i < len(words); i++ {
		for k := 0; k < len(words[i]); k++ {
			if words[i][k] == w {
				fmt.Println("Found", w)
				break DoSearch
			}
		}
	}
}

func main() {
	search("lake")
}

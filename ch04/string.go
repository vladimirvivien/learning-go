package main

import "fmt"

var (
	txt  = "水 and 火"
	txt2 = "\u6C34\x20brings\x20\x6c\x69\x66\x65."
	txt3 = `
	\u6C34\x20
	brings\x20
	\x6c\x69\x66\x65.
	`
)

func main() {
	fmt.Printf("%s (%d)\n", txt, len(txt))
	for i := 0; i < len(txt); i++ {
		fmt.Printf("%U ", txt[i])
	}
	fmt.Println()
	fmt.Println(txt2)
	fmt.Println(txt3)
}

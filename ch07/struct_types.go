package main

import "fmt"

var (
	empty    struct{}
	car      struct{ make, model string }
	currency struct {
		name, country string
		code          int
	}
	node struct {
		edges  []string
		weight int
	}
	person struct {
		name    string
		address struct {
			street      string
			city, state string
			postal      string
		}
	}
)

func main() {
	fmt.Printf("emtpy %T\n", empty)
	fmt.Printf("car %T\n", car)
	fmt.Printf("currency %T\n", currency)
	fmt.Printf("node %T\n", node)
	fmt.Printf("person %T\n", person)
}

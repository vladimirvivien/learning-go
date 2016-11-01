package main

import "fmt"

var (
	empty = struct{}{}

	car = struct{ make, model string }{make: "Ford", model: "F150"}

	currency = struct {
		name, country string
		code          int
	}{
		"USD", "United States",
		840,
	}

	node = struct {
		edges  []string
		weight int
	}{
		edges: []string{"north", "south", "west"},
	}

	person = struct {
		name    string
		address struct {
			street      string
			city, state string
			postal      string
		}
	}{
		name: "Tim the Robot",
		address: struct {
			street      string
			city, state string
			postal      string
		}{street: "Main Street"},
	}
)

func main() {
	fmt.Printf("car %T\n", car)
	fmt.Printf("currency %T\n", currency)
	fmt.Printf("node %T\n", node)
	fmt.Printf("person %T\n", person)
}

package main

import "fmt"

type person struct {
	name    string
	address address
}

type address struct {
	street      string
	city, state string
	postal      string
}

func makePerson() person {
	addr := address{
		city:   "Goville",
		state:  "Go",
		postal: "12345",
	}
	return person{
		name:    "vladimir vivien",
		address: addr,
	}
}

func main() {
	p := makePerson()
	fmt.Println(p)
}

package main

import "fmt"

type person struct {
	name  string
	title string
	address
}

type address struct {
	street      string
	city, state string
	postal      string
}

func updateName(p *person, name string) {
	p.name = name
}

func main() {
	p := new(person)
	p.name = "uknown"
	p.title = "author"
	p.street = "12345 Main street"
	p.city, p.state = "Goville", "Go"
	p.postal = "12345"
	fmt.Println(p)
	updateName(p, "Vladimir Vivien")
	fmt.Println(p)
}

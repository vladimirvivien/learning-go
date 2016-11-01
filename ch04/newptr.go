package main

import "fmt"

func main() {
	intptr := new(int)
	*intptr = 44

	p := new(struct{ first, last string })
	p.first = "Samuel"
	p.last = "Pierre"

	fmt.Printf("Value %d, type %T\n", *intptr, intptr)
	fmt.Printf("Person %+v\n", p)
}

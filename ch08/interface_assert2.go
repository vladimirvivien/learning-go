package main

import (
	"fmt"
)

type food interface {
	eat()
}

type veggie string

func (v veggie) eat() {
	fmt.Println("Eating", v)
}

type meat string

func (m meat) eat() {
	fmt.Println("Eating tasty", m)
}

func eat(f food) {
	switch morsel := f.(type) {
	case veggie:
		if morsel == "okra" {
			fmt.Println("Yuk! not eating ", morsel)
		} else {
			morsel.eat()
		}
	case meat:
		if morsel == "beef" {
			fmt.Println("Yuk! not eating ", morsel)
		} else {
			morsel.eat()
		}
	default:
		fmt.Println("Not eating whatever that is: ", f)
	}
}

func main() {
	eat(veggie("carrot"))
	eat(meat("lamb"))
	eat(veggie("okra"))
	eat(meat("beef"))
}

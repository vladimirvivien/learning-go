package main

import (
	"fmt"
	"strings"
)

func main() {
	a := 3
	double(&a)
	fmt.Println(a)
	p := &struct{ first, last string }{"Max", "Planck"}
	cap(p)
	fmt.Println(p)
}

func double(x *int) {
	*x = *x * 2
}

func cap(p *struct{ first, last string }) {
	p.first = strings.ToUpper(p.first)
	p.last = strings.ToUpper(p.last)
}

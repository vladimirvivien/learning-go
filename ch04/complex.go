package main

import "fmt"

func main() {
	a := -3.5 + 2i
	fmt.Printf("%v\n", a)
	fmt.Printf("%+g, %+g\n", real(a), imag(a))
}

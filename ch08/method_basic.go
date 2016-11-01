package main

import "fmt"

type gallon float64

func (g gallon) quart() float64 {
	return float64(g * 4)
}

func main() {
	gal := gallon(5)
	fmt.Println(gal.quart())
}

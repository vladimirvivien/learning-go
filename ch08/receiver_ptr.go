package main

import "fmt"

type gallon float64

func (g gallon) quart() float64 {
	return float64(g * 4)
}

func (g gallon) half() {
	g = gallon(g * 0.5)
}

func (g *gallon) double() {
	*g = gallon(*g * 2)
}

func main() {
	var gal gallon = 5
	gal.half()
	fmt.Println(gal)
	gal.double()
	fmt.Println(gal)
}

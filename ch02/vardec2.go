package main

import "fmt"

var name string = "Earth"
var desc string = "Planet"
var radius int32 = 6378
var mass float64 = 5.972E+24
var active bool = true
var satellites = []string{
	"Moon",
}

func main() {
	fmt.Println(name)
	fmt.Println(desc)
	fmt.Println("Radius (km)", radius)
	fmt.Println("Mass (kg)", mass)
	fmt.Println("Satellites", satellites)
}

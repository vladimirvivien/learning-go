package main

import "fmt"

var name = "Mars"
var desc = "Planet"
var radius = 3396.2
var mass = 6.4185e23
var active = true
var satellites = []string{
	"Phobos",
	"Deimos",
}

func main() {
	fmt.Println(name)
	fmt.Println(desc)
	fmt.Println("Radius (km)", radius)
	fmt.Println("Mass (kg)", mass)
	fmt.Println("Satellites", satellites)
}

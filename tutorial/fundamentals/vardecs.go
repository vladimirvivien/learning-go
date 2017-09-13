package main

import "fmt"

var name, desc string   // declares two variables of type string
var radius int32        // variable of type int32
var mass float64        // variable of type float64
var active bool         // variable of type bool
var satellites []string // variable of type []string

func main() {
	name = "Sun"
	desc = "Star"
	radius = 685800
	mass = 1.989E+30
	active = true
	satellites = []string{
		"Mercury",
		"Venus",
		"Earth",
		"Mars",
		"Jupiter",
		"Saturn",
		"Uranus",
		"Neptune",
	}
	fmt.Println(name)
	fmt.Println(desc)
	fmt.Println("Radius (km)", radius)
	fmt.Println("Mass (kg)", mass)
	fmt.Println("Satellites", satellites)
}

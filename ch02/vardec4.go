package main

import "fmt"

func main() {
	name := "Neptune"
	desc := "Planet"
	radius := 24764
	mass := 1.024e26
	active := true
	satellites := []string{
		"Naiad", "Thalassa", "Despina", "Galatea", "Larissa", "S/2004 N 1", "Proteus",
		"Triton", "Nereid", "Halimede", "Sao", "Laomedeia", "Neso", "Psamathe",
	}

	fmt.Println(name)
	fmt.Println(desc)
	fmt.Println("Radius (km)", radius)
	fmt.Println("Mass (kg)", mass)
	fmt.Println("Satellites", satellites)
	fmt.Println("Active", active)
}

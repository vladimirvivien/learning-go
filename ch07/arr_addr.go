package main

import "fmt"

type galaxies [14]string

func printGalaxies(names *galaxies) {
	for _, g := range names {
		fmt.Println(g)
	}
}

func main() {
	namedGalaxies := &galaxies{
		"Andromeda",
		"Black Eye",
		"Bode's",
		"Cartwheel",
		"Cigar",
		"Comet",
		"Hoag's",
		"Magellanic",
		"Mayall's",
		"Pinwheel",
		"Sombrero",
		"Sunflower",
		"Tadpole",
		"Whirpool",
	}
	fmt.Println("Named Galaxies")
	fmt.Println("--------------")
	printGalaxies(namedGalaxies)
}

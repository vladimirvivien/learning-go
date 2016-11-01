package main

import (
	"fmt"
)

type diameter int

type name struct {
	long   string
	short  string
	symbol rune
}

type planet struct {
	diameter
	name
	desc string
}

func nameStr(n name) string {
	return fmt.Sprintf(
		"long name: %s short name: %s symbol: %v",
		n.long, n.short, string(n.symbol),
	)
}

func main() {

	earth := planet{
		diameter: 7926,
		name: name{
			long:   "Earth",
			short:  "E",
			symbol: '\u2641',
		},
		desc: "Third rock from the Sun",
	}

	jupiter := planet{}
	jupiter.diameter = 88846
	jupiter.name.long = "Jupiter"
	jupiter.name.short = "J"
	jupiter.name.symbol = '\u2643'
	jupiter.desc = "A ball of gas"

	saturn := planet{}
	saturn.diameter = 120536
	saturn.long = "Saturn"
	saturn.short = "S"
	saturn.symbol = '\u2644'
	saturn.desc = "Slow mover"

	fmt.Printf("Planet %v, diam %d, desc: %s\n", nameStr(earth.name), earth.diameter, earth.desc)
	fmt.Printf("Planet %v, diam %d, desc: %s\n", nameStr(jupiter.name), jupiter.diameter, jupiter.desc)
	fmt.Printf("Planet %v, diam %d, desc: %s\n", nameStr(saturn.name), saturn.diameter, saturn.desc)

}

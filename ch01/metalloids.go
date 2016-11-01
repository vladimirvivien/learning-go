// This program prints molecular information for known metalloids
// including atomic number, mass, and atom count found
// in 100 grams of each element using the mole unit.
// See http://en.wikipedia.org/wiki/Mole_(unit)
package main

import "fmt"

const avogadro float64 = 6.0221413e+23
const grams = 100.0

type amu float64

func (mass amu) float() float64 {
	return float64(mass)
}

type metalloid struct {
	name   string
	number int32
	weight amu
}

var metalloids = []metalloid{
	metalloid{"Boron", 5, 10.81},
	metalloid{"Silicon", 14, 28.085},
	metalloid{"Germanium", 32, 74.63},
	metalloid{"Arsenic", 33, 74.921},
	metalloid{"Antimony", 51, 121.760},
	metalloid{"Tellerium", 52, 127.60},
	metalloid{"Polonium", 84, 209.0},
}

// finds # of moles
func moles(mass amu) float64 {
	return float64(mass) / grams
}

// returns # of atoms moles
func atoms(moles float64) float64 {
	return moles * avogadro
}

// return column headers
func headers() string {
	return fmt.Sprintf(
		"%-10s %-10s %-10s Atoms in %.2f Grams\n",
		"Element", "Number", "AMU", grams,
	)
}

func main() {
	fmt.Print(headers())

	for _, m := range metalloids {
		fmt.Printf(
			"%-10s %-10d %-10.3f %e\n",
			m.name, m.number, m.weight.float(), atoms(moles(m.weight)),
		)
	}
}

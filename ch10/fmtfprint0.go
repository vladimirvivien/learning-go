package main

import (
	"fmt"
	"os"
)

type metalloid struct {
	name   string
	number int32
	weight float64
}

func main() {
	var metalloids = []metalloid{
		{"Boron", 5, 10.81},
		{"Silicon", 14, 28.085},
		{"Germanium", 32, 74.63},
		{"Arsenic", 33, 74.921},
		{"Antimony", 51, 121.760},
		{"Tellerium", 52, 127.60},
		{"Polonium", 84, 209.0},
	}
	file, err := os.Create("./metalloids.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		os.Exit(1)
	}
	defer file.Close()

	for _, m := range metalloids {
		fmt.Fprintf(
			file,
			"%-10s %-10d %-10.3f\n",
			m.name, m.number, m.weight,
		)
	}
}

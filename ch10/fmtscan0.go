package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	var name, hasRing string
	var diam, moons int

	// read data
	data, err := os.Open("./planets.txt")
	if err != nil {
		fmt.Println("Unable to open planet data:", err)
		return
	}
	defer data.Close()

	// print data
	fmt.Printf(
		"%-10s %-10s %-6s %-6s\n",
		"Planet", "Diameter", "Moons", "Ring?",
	)
	for {
		_, err := fmt.Fscanf(
			data,
			"%s %d %d %s\n",
			&name, &diam, &moons, &hasRing,
		)
		if err != nil {
			if err == io.EOF {
				break
			} else {
				fmt.Println("Scan error:", err)
				return
			}
		}
		fmt.Printf(
			"%-10s %-10d %-6d %-6s\n",
			name, diam, moons, hasRing,
		)
	}

	// ask question
	fmt.Println("Which planet has the most moons? ")
	_, err = fmt.Scanf("%s", &name)
	if err != nil {
		fmt.Println("Invalid input provided, try again.")
		return
	}
	if "jupiter" == strings.ToLower(name) {
		fmt.Println("Corret, it's Jupiter!")
	} else {
		fmt.Println("No, look at the chart.")
	}
}

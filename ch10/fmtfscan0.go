package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	var name string
	var diam int
	var moons int
	var hasRing string

	data, err := os.Open("./planets.txt")
	if err != nil {
		fmt.Println("Unable to open planet data:", err)
		return
	}
	defer data.Close()

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

}

package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	var name, hasRing string
	var diam, moons int

	file, err := os.Open("./planets.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	fmt.Printf("%-10s %-10s %-6s %-6s\n", "Planet", "Diameter", "Moons", "Ring?")
	for {
		_, err := fmt.Fscanf(file, "%s %d %d %s\n", &name, &diam, &moons, &hasRing)
		if err != nil {
			if err == io.EOF {
				break
			} else {
				fmt.Println(err)
				os.Exit(1)
			}
		}
		fmt.Printf("%-10s %-10d %-6d %-6s\n", name, diam, moons, hasRing)
	}

}

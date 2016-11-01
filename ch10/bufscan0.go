package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("./planets.txt")
	if err != nil {
		fmt.Println("Unable to open file:", err)
		return
	}
	defer file.Close()

	fmt.Printf(
		"%-10s %-10s %-6s %-6s\n",
		"Planet", "Diameter", "Moons", "Ring?",
	)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		fields := strings.Split(scanner.Text(), " ")
		fmt.Printf(
			"%-10s %-10s %-6s %-6s\n",
			fields[0], fields[1], fields[2], fields[3],
		)
	}
}

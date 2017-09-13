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
	reader := bufio.NewScanner(file)

	for {
		line, err := reader.ReadString("\n")
		if err != nil {
			if err == io.EOF {
				break
			} else {
				fmt.Println(err)
				os.Exit(1)
			}
		}
		fmt.Println(line)
	}

}

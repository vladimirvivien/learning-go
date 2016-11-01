package main

import "fmt"

func do(steps ...string) {
	defer fmt.Println("All done!")
	for _, s := range steps {
		defer fmt.Println(s)
	}

	fmt.Println("Starting")
}

func main() {
	do(
		"Find key",
		"Aplly break",
		"Put key in ignition",
		"Start car",
	)
}

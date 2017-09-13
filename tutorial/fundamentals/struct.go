package main

import "fmt"

func main() {
	truck := struct {
		year        int
		make, model string
	}{
		make:  "Ford",
		model: "F150",
		year:  2017,
	}

	fmt.Printf("Truck: %d %s %s\n", truck.year, truck.make, truck.model)
}

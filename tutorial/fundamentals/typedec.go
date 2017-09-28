package main

import "fmt"

type car struct {
	year        int
	make, model string
}

func main() {
	ford := car{
		year:  2001,
		make:  "Ford",
		model: "F150",
	}

	fmt.Println(ford.make, ford.model)

	chevy := car{
		year:  2012,
		make:  "Chevrolet",
		model: "Silverado",
	}

	fmt.Println(chevy)
}

package main

import (
	"fmt"
	"math"
)

func dbl(val float64) {
	val = 2 * val // update param
	fmt.Printf("dbl()=%.5f\n", val)
}

func main() {
	p := math.Pi
	fmt.Printf("before dbl() p = %.5f\n", p)
	dbl(p)
	fmt.Printf("after dbl() p = %.5f\n", p)
}

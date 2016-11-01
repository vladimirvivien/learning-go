package main

import (
	"fmt"
	"math"
)

func main() {
	for i := 0.0; i < 360.0; i += 45.0 {
		rad := func() float64 {
			return i * math.Pi / 180
		}()
		fmt.Printf("%.2f Deg = %.2f Rad\n", i, rad)
	}
}

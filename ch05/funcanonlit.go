package main

import "fmt"

func main() {
	fmt.Printf(
		"94 (°F) = %.2f (°C)\n",
		func(f float64) float64 {
			return (f - 32.0) * (5.0 / 9.0)
		}(94),
	)
}

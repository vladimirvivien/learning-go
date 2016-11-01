package main

import (
	"fmt"
	"math"
)

func isPrime(n int) bool {
	lim := int(math.Sqrt(float64(n)))
	for p := 2; p <= lim; p++ {
		if (n % p) == 0 {
			return false
		}

	}
	return true
}

func main() {
	prime := 37
	fmt.Printf("isPrime(%d) = %v\n", prime, isPrime(prime))
}

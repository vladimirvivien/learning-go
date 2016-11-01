// Calculate sum of all multiple of 3 and 5 less MAX value.
// See https://projecteuler.net/problem=1
package main

import (
	"fmt"
)

const MAX = 1000

func main() {
	work := make(chan int, MAX)
	result := make(chan int)

	// Create channel of multiples of 3 and 5
	// concurrently using goroutine
	go func() {
		for i := 1; i < MAX; i++ {
			if (i%3) == 0 || (i%5) == 0 {
				work <- i // push for work
			}
		}
		close(work)
	}()

	// Concurrently sum up work and put result
	// in channel result
	go func() {
		r := 0
		for i := range work {
			r = r + i
		}
		result <- r
	}()

	// Wait for result, then print
	fmt.Println("Total:", <-result)
}

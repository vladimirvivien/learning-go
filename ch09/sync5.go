// Calculate sum of all multiple of 3 and 5 less than MAX value.
// See https://projecteuler.net/problem=1
package main

import (
	"fmt"
	"sync"
)

const MAX = 1000

func main() {
	values := make(chan int, MAX)
	result := make(chan int, 2)
	var wg sync.WaitGroup
	wg.Add(2)

	go func() { // gen multiple of 3 & 5 values
		for i := 1; i < MAX; i++ {
			if (i%3) == 0 || (i%5) == 0 {
				values <- i // push downstream
			}
		}
		close(values)
	}()

	work := func() { // work unit, calc partial result
		defer wg.Done()
		r := 0
		for i := range values {
			r += i
		}
		result <- r
	}

	// distribute work to two goroutines
	go work()
	go work()

	wg.Wait()                    // wait for both groutines
	total := <-result + <-result // gather partial results
	fmt.Println("Total:", total)
}

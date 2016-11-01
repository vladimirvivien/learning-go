// Calculate sum of all multiple of 3 and 5 less than MAX value.
// See https://projecteuler.net/problem=1
package main

import (
	"fmt"
	"sync"
)

const MAX = 1000
const workers = 2

func main() {
	values := make(chan int)
	result := make(chan int, workers)
	var wg sync.WaitGroup

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

	wg.Add(workers)
	for i := 0; i < workers; i++ {
		go work() // execute on its own thread
	}

	wg.Wait() // wait for both groutines
	close(result)
	total := 0
	for pr := range result {
		total += pr
	}
	fmt.Println("Total:", total)
}

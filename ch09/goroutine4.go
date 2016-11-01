package main

import (
	"fmt"
)

func main() {

	starts := []int{10,40,70,100}

	for _, j := range starts{
		go func() {
			count(j, j+20, 10)
		}()
	}

	fmt.Scanln() // wait for enter
}

func count(start, stop, delta int) {
	for i := start; i <= stop; i += delta {
		fmt.Println(i)
	}
}

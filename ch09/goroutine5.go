package main

import (
	"fmt"
)

func main() {
	starts := []int{10,40,70,100}

	for _, j := range starts{
		go func(s int) {
			count(s, s+20, 10)
		}(j)
	}

	fmt.Scanln() // wait for enter
}

func count(start, stop, delta int) {
	for i := start; i <= stop; i += delta {
		fmt.Println(i)
	}
}

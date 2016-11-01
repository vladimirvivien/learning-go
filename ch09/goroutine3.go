package main

import (
	"fmt"
)

func main() {

	start := 0
	stop := 50
	step := 5

	go func() {
		count(start, stop, step)
	}()

	fmt.Scanln() // wait for enter
}

func count(start, stop, delta int) {
	for i := start; i <= stop; i += delta {
		fmt.Println(i)
	}
}

package main

import (
	"fmt"
)

func main() {
	go count(10, 30, 10)

	go func() {
		count(40, 60, 10)
	}()

	go func() {
		count(70, 120, 20)
	}()

	fmt.Scanln() // wait for enter
}

func count(start, stop, delta int) {
	for i := start; i <= stop; i += delta {
		fmt.Println(i)
	}
}

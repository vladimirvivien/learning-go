package main

import (
	"fmt"
)

func main() {
	go count(10, 30, 10)
	go count(40, 60, 10)
	go count(70, 120, 20)
	fmt.Scanln() // blocks for kb input
}

func count(start, stop, delta int) {
	for i := start; i <= stop; i += delta {
		fmt.Println(i)
	}
}

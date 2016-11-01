package main

import (
	"fmt"
)

func main() {

	go func() {
		count(60, 100, 10)
	}()


	fmt.Scanln() // wait for enter
}

func count(start, stop, delta int) {
	for i := start; i <= stop; i += delta {
		fmt.Println(i)
	}
}

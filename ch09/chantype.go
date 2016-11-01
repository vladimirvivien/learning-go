package main

import "fmt"

var ints chan int

func main() {
	ints = make(chan int, 10)

	// produce
	for i := 0; i < 10; i++ {
		ints <- i
	}
	close(ints)

	// consume
	for i := range ints {
		fmt.Println(i)
	}
}

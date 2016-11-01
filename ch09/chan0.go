package main

import (
	"fmt"
)

func main() {
	ch := make(chan int, 4)
	ch <- 2
	ch <- 4
	ch <- 6
	ch <- 8

	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)

}

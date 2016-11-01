package main

import (
	"fmt"
)

func main() {
	ch := make(chan int, 4)
	ch <- 2
	ch <- 4
	close(ch)

	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch) // channel closed, returns zero value for element

}

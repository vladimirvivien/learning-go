package main

import (
	"fmt"
)

func main() {
	// execute within same routine (1)
	// each calls execute serially
	count(10, 50, 10)
	count(60, 100, 10)
	count(110, 200, 20)

	// end of main()
	// routine 1 exits
}

func count(start, stop, delta int) {
	for i := start; i <= stop; i += delta {
		fmt.Println(i)
	}
}

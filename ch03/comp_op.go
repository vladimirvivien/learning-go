package main

import (
	"fmt"
)

func main() {
	//var c1 = 12.0
	//var c2 = 44
	//var c3 = c1*2 + 3
	//fmt.Printf("c3 = %v (%T)\n", c3, c3)
	//e := 4.09
	var c0 uint = 32
	var c1 float32 = 1 << c0
	fmt.Printf("c1=%v,%T", c1, c1)
}

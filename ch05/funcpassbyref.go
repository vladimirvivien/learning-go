package main

import "fmt"

// modify the value referenced by pointer
func half(val *float64) {
	fmt.Printf("call half(%f)\n", *val)
	*val = *val / 2
}

func main() {
	num := 2.807770
	fmt.Printf("num=%f\n", num)
	half(&num)
	fmt.Printf("half(num)=%f\n", num)
}

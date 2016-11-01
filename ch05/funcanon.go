package main

import "fmt"

var (
	mul = func(op0, op1 int) int {
		return op0 * op1
	}

	sqr = func(val int) int {
		return mul(val, val)
	}
)

func main() {
	fmt.Printf("mul(25,7) = %d\n", mul(25, 7))
	fmt.Printf("sqr(13) = %d\n", sqr(13))
}

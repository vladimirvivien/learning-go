package main

import "fmt"

func add(op0 int, op1 int) int {
	return op0 + op1
}

func sub(op0, op1 int) int {
	return op0 - op1
}

func main() {
	var opAdd func(int, int) int = add
	opSub := sub
	fmt.Printf("op0(12,44)=%d\n", opAdd(12, 44))
	fmt.Printf("sub(99,13)=%d\n", opSub(99, 13))
}

package main

import "fmt"

func main() {
	structPtr := &struct{ x, y int }{44, 55}
	pairPtr := &[2]string{"A", "B"}
	fmt.Printf("struct=%#v, type=%T\n", structPtr, structPtr)
	fmt.Printf("pairPtr=%#v, type=%T\n", pairPtr, pairPtr)
}

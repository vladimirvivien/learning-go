package main

import "fmt"

type signal int

func main() {
	var count int32
	var actual int
	var test int32 = actual + count

	var sig signal
	var event int = sig

	fmt.Println(test)
	fmt.Println(event)
}

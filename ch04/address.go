package main

import (
	"fmt"
)

func main() {
	var a int = 1024
	var aptr *int = &a

	fmt.Printf("a=%v\n", a)
	fmt.Printf("aptr=%v\n", aptr)
}

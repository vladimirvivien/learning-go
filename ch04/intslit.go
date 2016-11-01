package main

import "fmt"

func main() {
	vals := []int{
		1024,
		0x0FF1CE,
		0x8BADF00D,
		0xBEEF,
		0777,
	}
	for _, i := range vals {
		if i == 0xBEEF {
			fmt.Printf("Got %d\n", i)
			break
		}
	}
}

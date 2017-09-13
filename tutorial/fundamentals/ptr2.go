package main

import "fmt"

func main() {
	score := 32
	adjust(score)
	fmt.Println(score) // score not updated
}

func adjust(val int) {
	val = val * 4
}

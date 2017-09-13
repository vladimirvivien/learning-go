package main

import "fmt"

func main() {
	score := 32
	scorePtr := &score
	adjust(scorePtr)
	fmt.Println(score) // score is updated
}

func adjust(val *int) {
	*val = *val * 4
}

package main

import "fmt"

func main() {
	score := 32
	scorePtr := &score    // pointer assigned address
	fmt.Println(scorePtr) // prints address
	*scorePtr = 44        // pointer dereferenced with value
	fmt.Println(score)    // prints updated score
}

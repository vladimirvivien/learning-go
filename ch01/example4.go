package main

import "fmt"

var steps int = 21

type Message string

func (msg Message) String() string { return string("<<" + msg + ">>") }
func main() {
	msg := Message("Mastering Go in")
	fmt.Println(msg, steps, "steps.")
}

package main

import "fmt"

var steps int = 21

type Message string

func (msg Message) Say() { fmt.Print(msg) }
func main() {
	msg := Message("Mastering Go in")
	msg.Say()
	fmt.Println(steps, "steps.")
}

package main

import "fmt"

func main() {
	steps := [3]string{"SEND", "RCVD", "WAIT"} // size 3 array, initialized
	fmt.Println(steps[1])                      // prints "RCVD"
	steps[3] = "PAUSE"                         // index out of range error
}

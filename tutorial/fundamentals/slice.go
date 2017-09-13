package main

import "fmt"

func main() {
	steps := []string{"SEND", "RCVD", "WAIT"} // slice initialized with 3 elements
	fmt.Println(steps[1])                     // prints "RCVD"
	steps = append(steps, "PAUSE")            // slice expanded to size 4
	steps[3] = "RESUME"                       // updates value at index 4

	actions := make([]string, 2) // initializes a slice of size 2
	actions[0] = "PRINT"
	actions[1] = "LOG"
	actions = append(actions, "ADD") // expands slice, size now 3
}

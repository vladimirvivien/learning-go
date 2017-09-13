package main

import "fmt"

func main() {
	steps := [3]string{"SEND", "RCVD", "WAIT"}
	for index, value := range steps {
		fmt.Printf("step %d = %s\n", index, value)
	}
}

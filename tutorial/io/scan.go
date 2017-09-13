package main

import (
	"fmt"
	"os"
)

func main() {
	var choice int
	fmt.Println("A square is what?")
	fmt.Print("Enter 1=quadrilateral 2=rectagonal: ")

	n, err := fmt.Scanf("%d", &choice)
	if n != 1 || err != nil {
		fmt.Println("invalid choice")
		os.Exit(1)
	}
	if choice == 1 {
		fmt.Println("You are correct")
	} else {
		fmt.Println("sorry")
	}
}

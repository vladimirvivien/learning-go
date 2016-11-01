package main

import "fmt"

func do(steps ...string) {
	defer fmt.Println("Go! Go! Go!")
	defer func() {
		fmt.Println("Hurry!")
		fmt.Println("We are late.")
	}()

	for _, s := range steps {
		fmt.Println(s)
	}
}

func main() {

	do(
		"Find pants",
		"Find shirt",
		"Get dressed",
		"Eat breakfast",
		"Find keys",
		"Get in car",
		"Start car",
		"Leave for work",
	)

}

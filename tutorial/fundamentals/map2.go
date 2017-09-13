package main

import "fmt"

func main() {
	hist := make(map[string]int)
	hist["Jan"] = 100
	hist["Feb"] = 445
	hist["Mar"] = 514

	for key, value := range hist {
		fmt.Printf("histogram[%s] = %d\n", key, value)
	}
}

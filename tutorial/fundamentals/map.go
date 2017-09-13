package main

import "fmt"

func main() {
	ratings := map[string][]int{
		"men":   {32, 55, 12, 55, 42, 53},
		"women": {44, 42, 23, 41, 65, 44},
	}

	ratings["children"] = []int{2, 34, 5, 43, 64, 22}

	fmt.Println(ratings)
}

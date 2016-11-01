package main

import "fmt"

var (
	histogram map[string]int = map[string]int{
		"Jan": 100, "Feb": 445, "Mar": 514, "Apr": 233,
		"May": 321, "Jun": 644, "Jul": 113, "Aug": 734,
		"Sep": 553, "Oct": 344, "Nov": 831, "Dec": 312,
	}

	table = map[string][]int{
		"Men":   []int{32, 55, 12, 55, 42, 53},
		"Women": []int{44, 42, 23, 41, 65, 44},
	}
)

func main() {
	fmt.Println(histogram)
	fmt.Println(table)
}

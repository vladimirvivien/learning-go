package main

import "fmt"

var val [100]int
var days [7]string
var weekdays [5]string
var truth [256]bool
var histogram [5]map[string]int
var board [4][2]int
var matrix [2][2][2][2]byte

func main() {
	fmt.Printf("%T: %v\n", val, val)
	fmt.Printf("%T: %v\n", days, days)
	fmt.Printf("%T: %v\n", weekdays, weekdays)
	fmt.Printf("%T: %v\n", truth, truth)
	fmt.Printf("%T: %v\n", histogram, histogram)

	printDays(days)
}

func printDays(days [7]string) {
	for i, val := range days {
		fmt.Printf("Day %d = %s\n", i, val)
	}
}

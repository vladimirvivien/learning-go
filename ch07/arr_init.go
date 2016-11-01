package main

import "fmt"

var val [100]int = [100]int{44, 72, 12, 55, 64, 1, 4, 90, 13, 54}
var msg = [12]rune{0: 'H', 2: 'E', 4: 'L', 6: 'O', 8: '!'}
var days [7]string = [7]string{
	"Monday",
	"Tuesday",
	"Wednesday",
	"Thursday",
	"Friday",
	"Saturday",
	"Sunday",
}
var weekdays = [...]string{
	"Monday",
	"Tuesday",
	"Wednesday",
	"Thursday",
	"Friday",
}

var truth = [256]bool{true}

var histogram = [5]map[string]int{
	map[string]int{"A": 12, "B": 1, "D": 15},
	map[string]int{"man": 1344, "women": 844, "children": 577, "pets": 150},
}

var board = [4][2]int{
	{33, 23},
	{62, 2},
	{23, 4},
	{51, 88},
}
var matrix = [2][2][2][2]byte{
	{{{4, 4}, {3, 5}}, {{55, 12}, {22, 4}}},
	{{{2, 2}, {7, 9}}, {{43, 0}, {88, 7}}},
}

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

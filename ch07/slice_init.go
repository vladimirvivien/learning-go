package main

import "fmt"

var (
	ids    []string = []string{"fe225", "ac144", "3b12c"}
	vector          = []float64{12.4, 44, 126, 2, 11.5}
	months          = []string{
		"Jan", "Feb", "Mar", "Apr",
		"May", "Jun", "Jul", "Aug",
		"Sep", "Oct", "Nov", "Dec",
	}
	// slice of map type
	tables = []map[string][]int{
		{
			"age": {53, 13, 5, 55, 45, 62, 34, 7},
			"pay": {124, 66, 777, 531, 933, 231},
		},
		{
			"points":   {3, 62, 88, 90, 32, 62, 8},
			"timenano": {12332, 43433, 524, 4443},
		},
	}
	board = [][]int{
		{712, 5, 344},
		{},
	}
	graph = [][][][]int{
		{{{44}, {3, 5}}, {{55, 12, 3}, {22, 4}}},
		{{{22, 12, 9, 19}, {7, 9}}, {{43, 0, 44, 12}, {7}}},
	}
)

func print(strs []string) {
	for _, str := range strs {
		fmt.Println(str)
	}
}

func main() {
	// all values are unitialized in this example
	fmt.Printf("Ids %T : %v\n", ids, ids)
	fmt.Printf("Vector %T : %v\n", vector, vector)
	fmt.Printf("Months %T : %v\n", months, months)
	fmt.Printf("tables %T : %v\n", tables, tables)
	fmt.Printf("Board %T : %v\n", board, board)
	fmt.Printf("Graph %T : %v\n", graph, graph)

	print(months)
}

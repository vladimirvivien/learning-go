package main

import "fmt"

var (
	image  []byte
	ids    []string
	vector []float64
	months []string
	q1     []string
	tables []map[string][]int
	board  [][]int
	graph  [][][][]int
)

func print(strs []string) {
	for _, str := range strs {
		fmt.Println(str)
	}
}

func main() {
	// all values are unitialized in this example
	fmt.Printf("Image %T : %v\n", image, image)
	fmt.Printf("Ids %T : %v\n", ids, ids)
	fmt.Printf("Vector %T : %v\n", vector, vector)
	fmt.Printf("Months %T : %v\n", months, months)
	fmt.Printf("Q1 %T : %v\n", q1, q1)
	fmt.Printf("Tables %T : %v\n", tables, tables)
	fmt.Printf("Board %T : %v\n", board, board)
	fmt.Printf("Graph %T : %v\n", graph, graph)

	print(months)
	print(q1)
}

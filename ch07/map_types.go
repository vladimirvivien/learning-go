package main

import "fmt"

var (
	legends     map[int]string
	histogram   map[string]int
	calibration map[float64]bool
	matrix      map[[2][2]int]bool                          // map with array key type
	table       map[string][]string                         // map of string slices
	log         map[struct{ name string }]map[string]string // map (with struct key) of map of string
)

func main() {
	fmt.Printf("%T\n", legends)
	fmt.Printf("%T\n", histogram)
	fmt.Printf("%T\n", calibration)
	fmt.Printf("%T\n", matrix)
	fmt.Printf("%T\n", table)
	fmt.Printf("%T\n", log)
}

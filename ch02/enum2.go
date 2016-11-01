package main

import (
	"fmt"
)

const (
	StarHyperGiant = 2.0 * iota
	StarSuperGiant
	StarBrightGiant
	StarGiant
	StarSubGiant
)
const (
	StarDwarf = iota
	StarSubDwarf
	StarWhiteDwarf
	StarRedDwarf
	StarBrownDwarf
)

func main() {
	fmt.Printf("StarHyperGiant = %v [%T]\n", StarHyperGiant, StarHyperGiant)
	fmt.Printf("StarSuperGiant = %v [%T]\n", StarSuperGiant, StarSuperGiant)
	fmt.Printf("StarBrightGiant = %v [%T]\n", StarBrightGiant, StarBrightGiant)
	fmt.Printf("StarGiant = %v [%T]\n", StarGiant, StarGiant)
	fmt.Printf("StarSubGiant = %v [%T]\n", StarSubGiant, StarSubGiant)
	fmt.Printf("StarDwarf = %v [%T]\n", StarDwarf, StarDwarf)
	fmt.Printf("StarSubDwarf = %v [%T]\n", StarSubDwarf, StarSubDwarf)
	fmt.Printf("StarWhiteDwarf = %v [%T]\n", StarWhiteDwarf, StarWhiteDwarf)
	fmt.Printf("StarRedDwarf = %v [%T]\n", StarRedDwarf, StarRedDwarf)
	fmt.Printf("StarBrownDwarf = %v [%T]\n", StarBrownDwarf, StarBrownDwarf)
}

package main

import (
	"fmt"
)

const (
	StarHyperGiant = iota
	StarSuperGiant
	StarBrightGiant
	StarGiant
	StarSubGiant
	StarDwarf
	StarSubDwarf
	StarWhiteDwarf
	StarRedDwarf
	StarBrownDwarf
)

func main() {
	fmt.Println(StarHyperGiant)
	fmt.Println(StarSuperGiant)
	fmt.Println(StarBrightGiant)
	fmt.Println(StarGiant)
	fmt.Println(StarSubGiant)
	fmt.Println(StarDwarf)
	fmt.Println(StarSubDwarf)
	fmt.Println(StarWhiteDwarf)
	fmt.Println(StarRedDwarf)
	fmt.Println(StarBrownDwarf)
}

package main

import "fmt"

type matrix [2][2][2][2]byte

func main() {
	var mat1 matrix
	mat1 = initMat()
	fmt.Println(mat1)
}

func initMat() matrix {
	return matrix{
		{{{4, 4}, {3, 5}}, {{55, 12}, {22, 4}}},
		{{{2, 2}, {7, 9}}, {{43, 0}, {88, 7}}},
	}
}

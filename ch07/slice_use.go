package main

import (
	"fmt"
)

func scale(factor float64, vector []float64) []float64 {
	for i, _ := range vector {
		vector[i] *= factor
	}
	return vector
}

func scale2(factor float64, vector []float64) {
	for i, _ := range vector {
		vector[i] *= factor
	}
}

func join(v1, v2 []float64) []float64 {
	return append(v1, v2...)
}

func clone(v []float64) (result []float64) {
	result = make([]float64, len(v), cap(v))
	copy(result, v)
	return
}

func add(v1, v2 []float64) []float64 {
	if len(v1) != len(v2) {
		panic("Size mismatch")
	}
	result := make([]float64, len(v1))
	for i := range v1 {
		result[i] = v1[i] + v2[i]
	}
	return result
}

func main() {
	h := []float64{12.5, 18.4, 7.0}
	h[0] = 15
	fmt.Println(h[0])

	h10 := scale(2.0, h)
	fmt.Println(h10)

	h2 := clone(h)
	fmt.Println(add(h, h2))
	fmt.Println("h2", h2)
	scale2(0.5, h2)
	fmt.Println("h2", h2)

}

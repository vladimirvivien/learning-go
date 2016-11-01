package main

import "fmt"

func avg(nums ...float64) float64 {
	n := len(nums)
	t := 0.0
	for _, v := range nums {
		t += v
	}
	return t / float64(n)
}

func sum(nums ...float64) float64 {
	var sum float64
	for _, v := range nums {
		sum += v
	}
	return sum
}

func main() {
	fmt.Printf("avg([1, 2.5, 3.75]) =%.2f\n", avg(1, 2.5, 3.75))
	points := []float64{9, 4, 3.7, 7.1, 7.9, 9.2, 10}
	fmt.Printf("sum(%v) = %.2f\n", points, sum(points...))
}

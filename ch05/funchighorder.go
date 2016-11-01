package main

import "fmt"

func apply(nums []int, f func(int) int) func() {
	for i, v := range nums {
		nums[i] = f(v)
	}
	return func() {
		fmt.Println(nums)
	}
}

func main() {
	nums := []int{4, 32, 11, 77, 556, 3, 19, 88, 422}
	result := apply(nums, func(i int) int {
		return i / 2
	})
	result()
}

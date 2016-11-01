package main

import "fmt"
import "math/rand"
import "time"

const (
	size     = 1000
	variance = 100000
)

var (
	nums [size]int
)

func init() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		nums[i] = rand.Intn(variance)
	}
}

func max(nums [size]int) int {
	temp := nums[0]
	for _, val := range nums {
		if val > temp {
			temp = val
		}
	}
	return temp
}

func main() {
	fmt.Println(nums)
	fmt.Println("Max num is: ", max(nums))
}

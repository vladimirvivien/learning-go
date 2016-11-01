package main

import (
	"fmt"
	"math"
)

func NewBigNumber() *float64 {
	big := math.MaxFloat64
	return &big
}

func half(val *float64) {
	*val = *val / 2
}

func double(val float64) {
	val = val * 2
}

func main() {
	bigNumber := NewBigNumber()
	fmt.Println("BigNumber =", *bigNumber)
	half(bigNumber)
	fmt.Println("Half BigNumber =", *bigNumber)
	double(*bigNumber)
	fmt.Println("BigNumber =", *bigNumber)
}

package main

import "math/big"

// recursive factorial with big package
func f(n int) *big.Int {
	if n == 0 || n == 1 {
		return big.NewInt(1)
	}
	next := big.NewInt(int64(n))
	var result big.Int
	return result.Mul(f(n-1), next)
}

func main() {}

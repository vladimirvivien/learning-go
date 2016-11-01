package main

import "fmt"

var (
	op0 = 3
	op1 = 10
)

func main() {
	q, r := divmod2(op0, op1)
	fmt.Printf("%d/%d quotient=%d remainder=%d", op0, op1, q, r)
}

/*
A Euclidian division adapted from the Wikipedia entry
http://en.wikipedia.org/wiki/Division_algorithm.
You also can find other implementations in the Go libraries
https://golang.org/src/pkg/math/big/int.go

// This version only works if dividend and divisor > 0.
// other conditions may cause panic or other unwanted behavior

*/
func divmod(dvdn, dvsr int) (q, r int) {
	r = dvdn
	for r >= dvsr {
		q++
		r = r - dvsr
	}
	return
}

func divmod2(dvdn, dvsr int) (int, int) {
	q := 0
	r := 0
	r = dvdn
	for r >= dvsr {
		q++
		r = r - dvsr
	}
	return q, r
}

package testexample

// A Euclidean division adapted from the Wikipedia entry
// http://en.wikipedia.org/wiki/Division_algorithm.
// You also can find other implementations in the Go libraries
// https://golang.org/src/pkg/math/big/int.go

// DivMod performs a Eucledan division producing a quotient and remainder.
// This version only works if dividend and divisor > 0.
func DivMod(dvdn, dvsr int) (q, r int) {
	r = dvdn
	for r >= dvsr {
		q += 1
		r = r - dvsr
	}
	return
}

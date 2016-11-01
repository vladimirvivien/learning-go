package main

import "fmt"

func div(dvdn, dvsr int) (q, r int) {
	r = dvdn
	for r >= dvsr {
		q++
		r = r - dvsr
	}
	return
}

func main() {
	q, r := div(71, 5)
	fmt.Printf("div(71,5) -> q = %d, r = %d\n", q, r)
}

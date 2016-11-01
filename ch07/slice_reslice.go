package main

import "fmt"

// reslice from existing slice
var (
	halfyr = []string{
		"Jan", "Feb", "Mar",
		"Apr", "May", "Jun",
	}

	q1 = halfyr[:3]
	q2 = halfyr[3:]
)

func main() {
	fmt.Printf("halfyr [%d,%d]--> %v\n", len(halfyr), cap(halfyr), halfyr)
	fmt.Printf("q1 [%d,%d]--> %v\n", len(q1), cap(q1), q1)
	fmt.Printf("q2 [%d,%d]--> %v\n", len(q2), cap(q2), q2)
}

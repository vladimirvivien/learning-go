package main

import "fmt"

func main() {
	months := make([]string, 6, 12)
	months[0] = "Jan"
	months[1] = "Feb"
	months[2] = "Mar"
	months[3] = "Apr"
	months[4] = "May"
	months[5] = "Jun"

	q1 := months[:3]
	q2 := months[3:6]

	fmt.Printf("months [%d,%d]--> %v\n", len(months), cap(months), months)
	fmt.Printf("q1 [%d,%d]--> %v\n", len(q1), cap(q1), q1)
	fmt.Printf("q2 [%d,%d]--> %v\n", len(q2), cap(q2), q2)
}

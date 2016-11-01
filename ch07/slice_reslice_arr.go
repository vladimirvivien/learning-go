package main

import "fmt"

// init slice from array
var (
	months [12]string = [12]string{
		"Jan", "Feb", "Mar",
		"Apr", "May", "Jun",
		"Jul", "Aug", "Sep",
		"Oct", "Nov", "Dec",
	}

	halfyr = months[:6]
	q1     = halfyr[:3]
	q2     = halfyr[3:6]
	q3     = months[6:9]
	q4     = months[9:]

	summer1 = months[5:8:8]
)

func main() {
	fmt.Printf("months [%d,%d]--> %v\n", len(months), cap(months), months)
	fmt.Printf("halfyr [%d,%d]--> %v\n", len(halfyr), cap(halfyr), halfyr)
	fmt.Printf("q1 [%d,%d]--> %v\n", len(q1), cap(q1), q1)
	fmt.Printf("q2 [%d,%d]--> %v\n", len(q2), cap(q2), q2)
	fmt.Printf("q3 [%d,%d]--> %v\n", len(q3), cap(q3), q3)
	fmt.Printf("q4 [%d,%d]--> %v\n", len(q4), cap(q4), q4)
	fmt.Printf("summer1[%d,%d]--> %v\n", len(summer1), cap(summer1), summer1)
}

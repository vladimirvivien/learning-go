package main

import "fmt"

func main() {
	months := [12]string{
		"Jan", "Feb", "Mar", "Apr", "May", "Jun",
		"Jul", "Aug", "Sep", "Oct", "Nov", "Dec",
	}
	summer := months[5:8]

	fmt.Println("--Summer Months--")

	for _, month := range summer {
		fmt.Println(month)
	}
}

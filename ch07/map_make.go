package main

import "fmt"

func main() {
	hist := make(map[string]int, 6)
	hist["Jan"] = 100
	hist["Feb"] = 445
	hist["Mar"] = 514
	hist["Apr"] = 233
	hist["May"] = 321
	hist["Jun"] = 644
	hist["Jul"] = 113
	hist["Aug"] = 734
	hist["Sep"] = 553
	hist["Oct"] = 344
	hist["Nov"] = 831
	hist["Dec"] = 312

	fmt.Println(hist)
	fmt.Println(len(hist))
}

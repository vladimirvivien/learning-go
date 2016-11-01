package main

import "fmt"

func avogadro() float64 {
	return 6.02214129e23
}

func main() {
	fmt.Printf("avogadro() = %e 1/mol\n", avogadro())
}

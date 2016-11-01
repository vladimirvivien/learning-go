package main

import (
	"fmt"
)

var (
	bksp = '\b'
	tab  = '\t'
	nwln = '\n'

	char1 = 'ɸ'
	char2 = 'আ'
	char3 = '語'

	char4 = '\u0369'
	char5 = '\xFA'
	char6 = '\045'
)

func main() {
	fmt.Println(bksp)
	fmt.Println(tab)
	fmt.Println(nwln)

	fmt.Println(char1)
	fmt.Println(char2)
	fmt.Println(char3)

	fmt.Println(char4)
	fmt.Println(char5)
	fmt.Println(char6)
}

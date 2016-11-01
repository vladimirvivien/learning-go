package main

import "fmt"

var valPtr *float32
var countPtr *int
var person *struct {
	name string
	age  int
}
var matrix *[1024]int
var row []*int64

func main() {
	fmt.Println(valPtr, countPtr, person, matrix, row)
}

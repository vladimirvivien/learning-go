package main

import (
	"fmt"
	"math"
	"unsafe"
)

var _ int8 = 12
var _ int16 = -400
var _ int32 = 12022
var _ int64 = 1 << 33
var _ int = 3 + 1415

var _ uint8 = 18
var _ uint16 = 44
var _ uint32 = 133121
var i uint64 = 23113233
var _ uint = 7542
var _ byte = 255
var _ uintptr = unsafe.Sizeof(i)

var _ float32 = 0.5772156649
var _ float64 = math.Pi

var _ complex64 = 3.5 + 2i
var _ complex128 = -5.0i

func main() {
	fmt.Println("all types declared!")
}

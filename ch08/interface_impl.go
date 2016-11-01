package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
	perim() float64
}

type rect struct {
	name           string
	length, height float64
}

func (r *rect) area() float64 {
	return r.length * r.height
}

func (r *rect) perim() float64 {
	return 2*r.length + 2*r.height
}

func (r *rect) String() string {
	return fmt.Sprintf(
		"%s[lenght:%.2f height:%.2f]",
		r.name, r.length, r.height,
	)
}

type triangle struct {
	name    string
	a, b, c float64
}

func (t *triangle) area() float64 {
	return 0.5 * (t.a * t.b)
}

func (t *triangle) perim() float64 {
	return t.a + t.b + math.Sqrt((t.a*t.a)+(t.b*t.b))
}

func (t *triangle) String() string {
	return fmt.Sprintf(
		"%s[sides: a=%.2f b=%.2f c=%.2f]",
		t.name, t.a, t.b, t.c,
	)
}

func shapeInfo(s shape) string {
	return fmt.Sprintf(
		"Area = %.2f, Perim = %.2f",
		s.area(), s.perim(),
	)
}

func main() {
	r := &rect{"Square", 4.0, 4.0}
	fmt.Println(r, "=>", shapeInfo(r))

	t := &triangle{"Right Triangle", 1, 2, 3}
	fmt.Println(t, "=>", shapeInfo(t))
}

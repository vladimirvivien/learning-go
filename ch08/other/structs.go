package main

import "fmt"

const (
	galunit = "GAL(US)"
	qrtunit = "QRT(US)"
	cupunit = "CUP(US)"
	ozunit  = "OZ(US)"
)

type volume struct {
	unit string
	qty  float64
}

func (v volume) String() string {
	return fmt.Sprintf("%.2f %s", v.qty, v.unit)
}

// gallon type
type gallon volume

func newGal(qty float64) gallon {
	return gallon{qty: qty, unit: galunit}
}
func (g gallon) String() string {
	return volume(g).String()
}

func (v gallon) toQuart() quart {
	return newQrt(v.qty * 4.0)
}

// quart type
type quart volume

func newQrt(qty float64) quart {
	return quart{qty: qty, unit: qrtunit}
}
func (q quart) String() string {
	return volume(q).String()
}
func (q quart) toCup() cup {
	return newCup(q.qty * 4.0)
}

// cup type
type cup volume

func newCup(qty float64) cup {
	return cup{qty: qty, unit: cupunit}
}
func (c cup) String() string {
	return volume(c).String()
}
func (c cup) toOunce() ounce {
	return newOz(c.qty * 8.0)
}

//ounce type
type ounce volume

func newOz(qty float64) ounce {
	return ounce{qty: qty, unit: ozunit}
}
func (o ounce) String() string {
	return volume(o).String()
}
func (o ounce) toCup() cup {
	return newCup(o.qty * 0.125)
}

func main() {
	gals := newGal(5)
	fmt.Printf("%v = %v\n", gals, gals.toQuart())
	ozs := gals.toQuart().toCup().toOunce()
	fmt.Printf("%v = %v\n", gals, ozs)
}

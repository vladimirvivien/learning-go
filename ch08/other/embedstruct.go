package main

import "fmt"

const (
	galunit = "GAL(US)"
	qrtunit = "QRT(US)"
	cupunit = "CUP(US)"
	ozunit  = "OZ(US)"
)

// base type volume
type volume struct {
	unit string
	qty  float64
}

func (v volume) String() string {
	return fmt.Sprintf("%.2f %s", v.qty, v.unit)
}

// gallon type
type gallon struct {
	volume    // embed volume type
	qrtFactor float64
}

func newGal(qty float64) gallon {
	var gal gallon
	gal.qty = qty
	gal.unit = galunit
	gal.qrtFactor = 4.0
	return gal
}
func (v gallon) toQuart() quart {
	return newQrt(v.qty * v.qrtFactor)
}

// quart type
type quart struct {
	volume
	galFactor float64
	cupFactor float64
}

func newQrt(qty float64) quart {
	return quart{
		volume:    volume{qty: qty, unit: qrtunit},
		galFactor: 0.25,
		cupFactor: 4.00,
	}
}
func (q quart) toCup() cup {
	return newCup(q.qty * q.cupFactor)
}

// cup type
type cup struct {
	volume
	qrtFactor float64
	ozFactor  float64
}

func newCup(qty float64) cup {
	return cup{
		volume:    volume{qty: qty, unit: cupunit},
		qrtFactor: 0.25,
		ozFactor:  8.00,
	}
}
func (c cup) toOunce() ounce {
	return newOz(c.qty * c.ozFactor)
}

//ounce type
type ounce struct {
	volume
	cupFactor float64
}

func newOz(qty float64) ounce {
	return ounce{
		volume:    volume{qty: qty, unit: ozunit},
		cupFactor: 0.125,
	}
}

func (o ounce) toCup() cup {
	return newCup(o.qty * o.cupFactor)
}

func main() {
	gals := newGal(5)
	fmt.Printf("%v = %v\n", gals, gals.toQuart())
	ozs := gals.toQuart().toCup().toOunce()
	fmt.Printf("%v = %v\n", gals, ozs)
}

package main

import "fmt"

// Using US volume units

type ounce float64
type cup float64
type quart float64
type gallon float64

func galsToQrts(gal gallon) quart {
	return quart(gal * 4.0)
}

func qrtsToCups(qrts quart) cup {
	return cup(qrts * 4.0)
}

func cupsToOz(c cup) ounce {
	return ounce(c * 8.0)
}
func main() {
	gal := gallon(5)
	qrt := galsToQrts(gal)
	fmt.Printf("%.2f gallons = %.2f quarts\n", gal, qrt)
	ozs := cupsToOz(qrtsToCups(galsToQrts(gal)))
	fmt.Printf("%.2f gallons = %.2f ounces\n", gal, ozs)
}

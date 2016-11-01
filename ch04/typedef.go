package main

import "fmt"

type fahrenheit float64
type celsius float64
type kelvin float64

func fharToCel(f fahrenheit) celsius {
	return celsius((f - 32) * 5 / 9)
}

func fharToKel(f fahrenheit) celsius {
	return celsius((f-32)*5/9 + 273.15)
}

func celToFahr(c celsius) fahrenheit {
	return fahrenheit(c*5/9 + 32)
}

func celToKel(c celsius) kelvin {
	return kelvin(c + 273.15)
}

func main() {
	var c celsius = 32.0
	f := fahrenheit(122)
	fmt.Printf("%.2f \u00b0C = %.2f \u00b0K\n", c, celToKel(c))
	fmt.Printf("%.2f \u00b0F = %.2f \u00b0C\n", f, fharToCel(f))
}

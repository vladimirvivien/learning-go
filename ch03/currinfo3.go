package main

import "fmt"

/*
PrintCurrencyInfo outputs the currency information
for a given country.  This version of the function
accepts three string parameters (curr, ctry, code), where
curr=currency, ctry=country, and code=currency code.
*/
func PrintCurrencyInfo(curr, ctry, code string) {
	fmt.Println("The", ctry, "currency is the", curr, "(", code, ")")
}

// To run this program, use
// go run currinfo3.go
func main() {
	currency := "US Dollar"    // country currency
	country := "United States" // country name
	code := "USD"              // currency code
	PrintCurrencyInfo(currency, country, code)
}

package main

//import "fmt"
//import "github.com/vladimirvivien/mastering-go/ch02lib"

import "github.com/vladimirvivien/currency/lib"

func main() {
	currency := "US Dollar"
	country := "United States"
	code := "USD"
	//fmt.Println("The currency of", country, "is the", currency, "(", code, ")")
	lib.PrintCurrencyInfo(country, currency, code)

}

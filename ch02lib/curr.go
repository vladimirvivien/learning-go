package lib

import "fmt"

func PrintCurrencyInfo(country, currency, code string) {
	fmt.Println("The currency of", country, "is the", currency, "(", code, ")")
}

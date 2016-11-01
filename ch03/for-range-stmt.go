package main

import (
	"fmt"
	"strings"
)

type Curr struct {
	Currency string
	Name     string
	Country  string
	Number   int
}

var currencies = []Curr{
	Curr{"NOK", "Norwegian Krone", "Norwary", 578},
	Curr{"KES", "Kenyan Shilling", "Kenya", 404},
	Curr{"DZD", "Algerian Dinar", "Algeria", 12},
	Curr{"AUD", "Australian Dollar", "Australia", 36},
	Curr{"MXN", "Mexican Peso", "Mexico", 484},
	Curr{"GRD", "Drachma", "Greece", 978},
	Curr{"KHR", "Riel", "Cambodia", 116},
	Curr{"SZL", "Lilangeni", "Swaziland", 748},
	Curr{"GBP", "Pound Sterling", "Isle of Man", 826},
	Curr{"HTG", "Gourde", "Haiti", 332},
	Curr{"BWP", "Pula", "Botswana", 72},
	Curr{"CLP", "Chilean Peso", "Chile", 152},
	Curr{"HKD", "Hong Kong Dollar", "Hong Koong", 344},
	Curr{"HTG", "Gourde", "Haiti", 332},
	Curr{"TRY", "Turkish Lira", "Turkey", 949},
	Curr{"EUR", "Euro", "Belgium", 978},
	Curr{"JMD", "Jamaican Dollar", "Jamaica", 388},
	Curr{"ALL", "Lek", "Albania", 8},
	Curr{"GEL", "Lari", "Georgia", 981},
	Curr{"KFM", "Coromo Franc", "Comoros", 174},
	Curr{"NZD", "New Zeland Dollar", "Tokelau", 554},
}

var sortedCurrs []Curr

func main() {
	updateCurrency(978, Curr{"EUR", "Euro", "Greece", 978})
	printCurrencies()
}

func updateCurrency(number int, newCurr Curr) {
	for i, v := range currencies {
		if number == v.Number {
			currencies[i] = newCurr
			fmt.Printf("%v -> %v\n", v, newCurr)
			break
		}
	}
}

func printCurrencies() {
	for i := range currencies {
		fmt.Printf("%d: %v\n", i, currencies[i])
	}
}

func find(filter string) {
	for _, c := range currencies {
		switch {
		case strings.Contains(c.Currency, filter),
			strings.Contains(c.Name, filter),
			strings.Contains(c.Country, filter):
			fmt.Println("Found", c)
		}
	}
}

func doEmptyRange() {
	for range currencies {
		fmt.Println("A currency found.")
	}
}

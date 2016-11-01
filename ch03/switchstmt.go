package main

import "fmt"

type Curr struct {
	Currency string
	Name     string
	Country  string
	Number   int
}

var currencies = []Curr{
	Curr{"DZD", "Algerian Dinar", "Algeria", 12},
	Curr{"AUD", "Australian Dollar", "Australia", 36},
	Curr{"EUR", "Euro", "Belgium", 978},
	Curr{"CLP", "Chilean Peso", "Chile", 152},
	Curr{"EUR", "Euro", "Greece", 978},
	Curr{"HTG", "Gourde", "Haiti", 332},
	Curr{"HKD", "Hong Kong Dollar", "Hong Koong", 344},
	Curr{"KES", "Kenyan Shilling", "Kenya", 404},
	Curr{"MXN", "Mexican Peso", "Mexico", 484},
	Curr{"USD", "US Dollar", "United States", 840},
	Curr{"EUR", "Euro", "Italy", 978},
}

func isDollar(curr Curr) bool {
	var result bool
	switch curr {
	default:
		result = false
	case Curr{"AUD", "Australian Dollar", "Australia", 36}:
		result = true
	case Curr{"HKD", "Hong Kong Dollar", "Hong Koong", 344}:
		result = true
	case Curr{"USD", "US Dollar", "United States", 840}:
		result = true
	}
	return result
}

func isDollar2(curr Curr) bool {
	switch curr {
	case Curr{"AUD", "Australian Dollar", "Australia", 36}:
		fallthrough
	case Curr{"HKD", "Hong Kong Dollar", "Hong Koong", 344}:
		fallthrough
	case Curr{"USD", "US Dollar", "United States", 840}:
		return true
	default:
		return false
	}
}

func isEuro(curr Curr) bool {
	switch curr {
	case currencies[2], currencies[4], currencies[10]:
		return true
	default:
		return false
	}
}

func main() {
	curr := Curr{"EUR", "Euro", "Italy", 978}
	if isDollar(curr) {
		fmt.Printf("%+v is Dollar currency\n", curr)
	} else if isEuro(curr) {
		fmt.Printf("%+v is Euro currency\n", curr)
	} else {
		fmt.Println("Currency is not Dollar or Euro")
	}

	dol := Curr{"AUD", "Australian Dollar", "Australia", 36}
	if isDollar2(dol) {
		fmt.Println("Dollar currency found:", dol)
	}
}

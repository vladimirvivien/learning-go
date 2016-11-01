package main

import (
	"errors"
	"fmt"
)

type Currency struct {
	Name    string
	Country string
	Number  int
}

var CAD = Currency{Name: "Canadian Dollar", Country: "Canada", Number: 124}
var FJD = Currency{Name: "Fiji Dollar", Country: "Fiji", Number: 242}
var JMD = Currency{Name: "Jamaican Dollar", Country: "Jamaica", Number: 388}
var USD = Currency{Name: "US Dollar", Country: "USA", Number: 840}

func main() {
	num := 840
	if num < 100 || num > 900 {
		fmt.Println("Invalid currency number")
	} else {
		if curr, err := findCurr(num); err == nil {
			fmt.Println("Currency found:", curr)
		} else {
			fmt.Println("Currency code", num, " not found")
		}
	}
}

func findCurr(number int) (Currency, error) {
	if CAD.Number == number {
		return CAD, nil
	} else if FJD.Number == number {
		return FJD, nil
	} else if JMD.Number == number {
		return JMD, nil
	} else if USD.Number == number {
		return USD, nil
	} else {
		return Currency{}, errors.New("Currency not found")
	}
}

package main

import (
	"fmt"
	"time"
)

type coord struct {
	long, lat float64
}

type loc struct {
	coord coord
	desc  string
	date  time.Time
}

func main() {
	loc1 := loc{
		coord: coord{
			35.027356,
			-85.157655,
		},
		desc: "Gas station",
		date: time.Now(),
	}

	loc1.desc = "Local Gas station"

	fmt.Printf("Loc coords long = %f, lat = %f\n", loc1.coord.long, loc1.coord.lat)
	fmt.Println("Location description = ", loc1.desc)
	fmt.Println("Visitation date/time = ", loc1.date.String())
}

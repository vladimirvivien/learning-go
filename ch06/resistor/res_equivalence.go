package resistor

// Rser returns equivalent resistance for resistors
// arrange in series in a circuit
func Rser(resists ...float64) (Rtotal float64) {
	for _, r := range resists {
		Rtotal = Rtotal + r
	}
	return
}

// Rpara return equivalent resistance for resistors
// arranged in parallel in a circuit
func Rpara(resists ...float64) (Rtotal float64) {
	for _, r := range resists {
		Rtotal = Rtotal + recip(r)
	}
	return
}

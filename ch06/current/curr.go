package current

// I returns current using Ohm's law
// when volt and impendence are known
func I(v, r float64) float64 {
	return v / r
}

// Ipv returns current when volt and power are known.
func Ivp(v, p float64) float64 {
	return p / v
}

// Ipara returns total current in parallel circuit
func Ipara(currents ...float64) (Itotal float64) {
	for _, i := range currents {
		Itotal = Itotal + i
	}
	return
}

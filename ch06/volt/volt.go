package volt

// V returns the votage of circuit
// using Ohm's law
func V(i, r float64) float64 {
	return i * r
}

// Vser returns the total voltage across a
// a series circuit
func Vser(volts ...float64) (Vtotal float64) {
	for _, v := range volts {
		Vtotal = Vtotal + v
	}
	return
}

// Vip returns the voltage when power p and
// current i are known
func Vpi(p, i float64) float64 {
	return p / i
}

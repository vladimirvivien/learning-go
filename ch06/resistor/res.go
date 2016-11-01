package resistor

// R returns resistance using Ohm's law
// when volt v and current i are known.
func R(v, i float64) float64 {
	return v / i
}

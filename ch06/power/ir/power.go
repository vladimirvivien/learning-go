package ir

// P calculates power when current i
// and resistance r are know for a circuit
func P(i, r float64) float64 {
	return i * i * r
}

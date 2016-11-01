package vr

// P returns power consumed in a circuit
// when voltage r and resistance r are known.
func P(v, r float64) float64 {
	return (v * v) / r
}

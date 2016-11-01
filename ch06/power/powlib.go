package power

// P returns power consumed in a circuit when
// voltage v and current i are known
func P(v, i float64) float64 {
	return v * i
}

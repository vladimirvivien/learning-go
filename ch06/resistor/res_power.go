package resistor

var Rpi func(float64, float64) float64

func init() {
	Rpi = func(p, i float64) float64 {
		return p / (i * i)
	}
}

// Rvp calculates resistance in a circuit
// when volt v and power p are known
func Rvp(v, p float64) float64 {
	return (v * v) / p
}

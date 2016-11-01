package testexample

import "testing"

func TestDivide(t *testing.T) {
	dvnd := 40
	for dvsor := 1; dvsor < dvnd; dvsor++ {
		q, r := DivMod(dvnd, dvsor)
		if (dvnd % dvsor) != r {
			t.Fatalf("%d/%d q=%d, r=%d, bad remainder.", dvnd, dvsor, q, r)
		}
	}
}

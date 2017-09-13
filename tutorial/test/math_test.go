package main

import "testing"

func TestSum(t *testing.T) {
	expected := 16
	actual := sum(5, 5, 3, 3)
	if actual != float64(expected) {
		t.Errorf("expecting 16, go %f", actual)
	}
}

func TestMax(t *testing.T) {
	expected := float64(1.2)
	actual := max(1.2, 0.3, 1.02, 0.20, 0.175)
	if actual != expected {
		t.Errorf("expecting 1.2, got %f", actual)
	}
}

func TestAvg(t *testing.T) {
	cases := []struct {
		nums []float64
		avg  float64
	}{
		{[]float64{5, 5, 3, 3}, 4},
		{[]float64{3, 9, 3}, 5},
		{[]float64{3.5, 1.5, 3.2, 1.8}, 2.5},
	}

	for _, c := range cases {
		actual := avg(c.nums...)
		if c.avg != actual {
			t.Errorf("unexpected result: need %f got %f", c.avg, actual)
		}
	}
}

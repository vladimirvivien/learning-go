package vector

import (
	"math"
	"testing"
	"os"
)

func TestNewVector(t *testing.T) {
	v := New(1, 2, 4)
	if len(v) != 3 {
		t.Errorf("Expecting vector size %d, but got %d", 3, len(v))
	}
}

func TestVectorString(t *testing.T) {
	v := New(1, 2, 3, 4)
	if v.String() != "[1,2,3,4]" {
		t.Logf("Expecting [1,2,3,4], but got %s", v.String())
		t.Fail()
	}
}

func TestVectorEqual(t *testing.T) {
	v1 := New(45, 44, 90)
	var v2 SimpleVector = []float64{45, 44, 90}
	t.Log(v1.Angle(v2), zero, v1.Angle(v2) <= zero)
	if !v1.Eq(v2) {
		t.Logf("Vectors are expected to be eqal")
		t.Fail()
	}
}

func TestVectorEqual2(t *testing.T) {
	v1 := New(45, 44, 90)
	var v2 SimpleVector = []float64{45, 44, 90}
	if !v1.Eq2(v2) {
		t.Logf("Vectors are expected to be eqal")
		t.Fail()
	}
}

func TestVectorAdd(t *testing.T) {
	v1 := New(8.218, -9.341)
	v2 := New(-1.129, 2.111)
	v3 := v1.Add(v2)
	expect := New(
		v1[0]+v2[0],
		v1[1]+v2[1],
	)

	if !v3.Eq(expect) {
		t.Logf("Addition failed, expecting %s, got %s", expect, v3)
		t.Fail()
	}
	t.Log(v1, "+", v2, v3)
}

func TestVectorSub(t *testing.T) {
	v1 := New(7.119, 8.215)
	v2 := New(-8.223, 0.878)
	v3 := v1.Sub(v2)
	expect := New(
		v1[0]-v2[0],
		v1[1]-v2[1],
	)
	if !v3.Eq(expect) {
		t.Log("Subtraction failed, expecting %s, got %s", expect, v3)
		t.Fail()
	}
	t.Log(v1, "-", v2, "=", v3)
}

func TestVectorScale(t *testing.T) {
	v := New(1.671, -1.012, -0.318)
	v.Scale(7.41)
	expect := New(
		7.41*1.671,
		7.41*-1.012,
		7.41*-0.318,
	)
	if !v.Eq(expect) {
		t.Logf("Scalar mul failed, expecting %s, got %s", expect, v)
		t.Fail()
	}
	t.Log("1.671,-1.012, -0.318 Scale", 7.41, "=", v)
}

func TestVectorMag(t *testing.T) {
	cases := []struct{
		vec SimpleVector
		expected float64

	}{
    	{New(1.2, 3.4), math.Sqrt(1.2*1.2 + 3.4*3.4)},
		{New(-0.21, 7.47), math.Sqrt(-0.21*-0.21 + 7.47*7.47)},
		{New(1.43, -5.40), math.Sqrt(1.43*1.43 + -5.40*-5.40)},
		{New(-2.07, -9.0), math.Sqrt(-2.07*-2.07 + -9.0*-9.0)},
	}
	for _, c := range cases {
		mag := c.vec.Mag()
		if mag != c.expected {
			t.Errorf("Magnitude failed, execpted %d, got %d", c.expected, mag)
		}
	}
}

func TestVectorUnit(t *testing.T) {
	v := New(5.581, -2.136)
	mag := v.Mag()
	expect := New((1/mag)*v[0], (1/mag)*v[1])
	if !v.Unit().Eq(expect) {
		t.Logf("Vector Unit failed, expecting %s, got %s", expect, v.Unit())
		t.Fail()
	}
	t.Logf("Vector = %v; Unit vector = %v\n", v, expect)
}

func TestVectorDotProd(t *testing.T) {
	v1 := New(7.887, 4.138)
	v2 := New(-8.802, 6.776)
	actual := v1.DotProd(v2)
	expect := v1[0]*v2[0] + v1[1]*v2[1]
	if actual != expect {
		t.Logf("DotPoduct2 failed, expecting %v, got %v", expect, actual)
		t.Fail()
	}
	t.Log(v1, "DotProd", v2, "=", actual)
}

func TestVectorAngle(t *testing.T) {
	if os.Getenv("RUN_ANGLE") == "" {
		t.Skipf("Env variable RUN_ANGLE not set, skipping:")
	}
	v1 := New(3.183, -7.627)
	v2 := New(-2.668, 5.319)
	actual := v1.Angle(v2)
	expect := math.Acos(v1.DotProd(v2) / (v1.Mag() * v2.Mag()))
	if actual != expect {
		t.Logf("Vector angle failed, expecting %d, got %d", expect, actual)
		t.Fail()
	}
	t.Log("Angle between", v1, "and", v2, "=", actual)
}

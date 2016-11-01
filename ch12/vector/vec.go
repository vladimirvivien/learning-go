package vector

import (
	"bytes"
	"math"
	"strconv"
)

const (
	zero = 1.0e-7 // zero tolerance
)

type Vector interface {
	String() string
	Eq(other Vector) bool
	Add(other Vector) Vector
	Sub(other Vector) Vector
	Scale(factor float64)
	DotProd(other Vector) float64
	Angle(other Vector) float64
	Mag() float64
	Unit() Vector
}

// the Vector type is a slice of float64
type SimpleVector []float64

func New(elems ...float64) SimpleVector {
	return SimpleVector(elems)
}

func (v SimpleVector) assertLenMatch(other Vector) {
	if len(v) != len(other.(SimpleVector)) {
		panic("Vector length mismatch")
	}
}

// String returns a string representation of the Vector
func (v SimpleVector) String() string {
	buff := bytes.NewBufferString("[")
	for i, val := range v {
		buff.WriteString(strconv.FormatFloat(val, 'g', -1, 64))
		if i < len(v)-1 {
			buff.WriteRune(',')
		}
	}
	buff.WriteRune(']')
	return buff.String()
}

// Eq compares vector magnitude and directions
func (v SimpleVector) Eq(other Vector) bool {
	ang := v.Angle(other)
	if math.IsNaN(ang) {
		return v.Mag() == other.Mag()
	}
	return v.Mag() == other.Mag() && ang <= zero
}

// Eq compares each vector components for equality
func (v SimpleVector) Eq2(other Vector) bool {
	v.assertLenMatch(other)
	otherVec := other.(SimpleVector)
	for i, val := range v {
		if val != otherVec[i] {
			return false
		}
	}
	return true
}

// Test for the zero vector
func (v SimpleVector) IsZero() bool {
	return v.Mag() <= zero
}

// Add returns the sum of two vectors
func (v SimpleVector) Add(other Vector) Vector {
	v.assertLenMatch(other)
	otherVec := other.(SimpleVector)
	result := make([]float64, len(v))
	for i, val := range v {
		result[i] = val + otherVec[i]
	}
	return SimpleVector(result)
}

// Sub returns the subtraction of a vector from another
func (v SimpleVector) Sub(other Vector) Vector {
	v.assertLenMatch(other)
	otherVec := other.(SimpleVector)
	result := make([]float64, len(v))
	for i, val := range v {
		result[i] = val - otherVec[i]
	}
	return SimpleVector(result)
}

// Scale scales the vector
func (v SimpleVector) Scale(scale float64) {
	for i := range v {
		v[i] = v[i] * scale
	}
}

// Mag computes the magnitude of the vector
func (v SimpleVector) Mag() (result float64) {
	for _, v := range v {
		result += (v * v)
	}
	return math.Sqrt(result)
}

// Unit returns the normalization of the vector
func (v SimpleVector) Unit() Vector {
	var result SimpleVector = make([]float64, len(v))
	copy(result, v)
	mag := result.Mag()
	result.Scale(1 / mag)
	return result
}

// DotProd calculates the dot product  using sum of prudcts
func (v SimpleVector) DotProd(other Vector) (result float64) {
	v.assertLenMatch(other)
	otherVec := other.(SimpleVector)
	for i, val := range v {
		result += val * otherVec[i]
	}
	return
}

// Angle calculates the angle between two vectors (Rad)
func (v SimpleVector) Angle(other Vector) float64 {
	return math.Acos(v.DotProd(other) / (v.Mag() * other.Mag()))
}

func (v SimpleVector) Proj(base Vector) Vector {
	baseUnit := base.Unit()
	baseUnit.Scale(v.DotProd(baseUnit))
	return baseUnit
}

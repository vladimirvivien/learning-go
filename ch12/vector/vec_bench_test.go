package vector

import (
	"math/rand"
	"testing"
	"time"
)

func BenchmarkVectorEqual1(b *testing.B) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < b.N; i++ {
		v1 := New(r.Float64(), r.Float64())
		v1.Eq(v1)
	}
}

func BenchmarkVectorEqual2(b *testing.B) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < b.N; i++ {
		v1 := New(r.Float64(), r.Float64())
		v1.Eq2(v1)
	}
}

func BenchmarkVectorAdd(b *testing.B) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < b.N; i++ {
		v1 := New(r.Float64(), r.Float64())
		v2 := New(r.Float64(), r.Float64())
		v1.Add(v2)
	}
}

func BenchmarkVectorSub(b *testing.B) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < b.N; i++ {
		v1 := New(r.Float64(), r.Float64())
		v2 := New(r.Float64(), r.Float64())
		v1.Sub(v2)
	}
}

func BenchmarkVectorScale(b *testing.B) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < b.N; i++ {
		v1 := New(r.Float64(), r.Float64())
		v1.Scale(r.Float64())
	}
}

func BenchmarkVectorMag(b *testing.B) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < b.N; i++ {
		v1 := New(r.Float64(), r.Float64())
		v1.Mag()
	}
}

func BenchmarkVectorUnit(b *testing.B) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < b.N; i++ {
		v1 := New(r.Float64(), r.Float64())
		v1.Unit()
	}
}

func BenchmarkVectorDotProd(b *testing.B) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < b.N; i++ {
		v1 := New(r.Float64(), r.Float64())
		v2 := New(r.Float64(), r.Float64())
		v1.DotProd(v2)
	}
}

func BenchmarkVectorAngle(b *testing.B) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < b.N; i++ {
		v1 := New(r.Float64(), r.Float64())
		v2 := New(r.Float64(), r.Float64())
		v1.Angle(v2)
	}
}

package main

import (
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/vladimirvivien/learning-go/ch12/vector"
)

func TestVectorAdd(t *testing.T) {
	reqBody := "[[1,2],[3,4]]"
	req, err := http.NewRequest("POST", "http://0.0.0.0/", strings.NewReader(reqBody))
	if err != nil {
		t.Fatal(err)
	}
	actual := vector.New(1, 2).Add(vector.New(3, 4))
	w := httptest.NewRecorder()
	add(w, req)
	if actual.String() != strings.TrimSpace(w.Body.String()) {
		t.Fatalf("Expecting actual %s, got %s", actual.String(), w.Body.String())
	}
}

func TestVectorSub(t *testing.T) {
	reqBody := "[[99,2.4],[3.22,4]]"
	req, err := http.NewRequest("POST", "http://0.0.0.0/", strings.NewReader(reqBody))
	if err != nil {
		t.Fatal(err)
	}
	actual := vector.New(99, 2.4).Sub(vector.New(3.22, 4))
	w := httptest.NewRecorder()
	sub(w, req)
	if actual.String() != strings.TrimSpace(w.Body.String()) {
		t.Fatalf("Expecting actual %s, got %s", actual.String(), w.Body.String())
	}
}

func TestVectorDotProd(t *testing.T) {
	reqBody := "[[-9,12.33],[3.22,6]]"
	req, err := http.NewRequest("POST", "http://0.0.0.0/", strings.NewReader(reqBody))
	if err != nil {
		t.Fatal(err)
	}
	actual := vector.New(-9, 12.33).DotProd(vector.New(3.22, 6))
	w := httptest.NewRecorder()
	dotProd(w, req)
	result, err := strconv.ParseFloat(strings.TrimSpace(w.Body.String()), 64)
	if err != nil {
		t.Fatal(err)
	}
	if actual != result {
		t.Fatalf("Expecting actual %d, got %d", actual, result)
	}
}

func TestVectorMag(t *testing.T) {
	reqBody := "[9,12.33,7]"
	req, err := http.NewRequest("POST", "http://0.0.0.0/", strings.NewReader(reqBody))
	if err != nil {
		t.Fatal(err)
	}
	actual := vector.New(9, 12.33, 7).Mag()
	w := httptest.NewRecorder()
	mag(w, req)
	result, err := strconv.ParseFloat(strings.TrimSpace(w.Body.String()), 64)
	if err != nil {
		t.Fatal(err)
	}
	if actual != result {
		t.Fatalf("Expecting actual %d, got %d", actual, result)
	}
}

func TestVectorUnit(t *testing.T) {
	reqBody := "[122, 99, -3, 77]"
	req, err := http.NewRequest("POST", "http://0.0.0.0/", strings.NewReader(reqBody))
	if err != nil {
		t.Fatal(err)
	}
	actual := vector.New(122, 99, -3, 77).Unit()
	w := httptest.NewRecorder()
	unit(w, req)
	if actual.String() != strings.TrimSpace(w.Body.String()) {
		t.Fatalf("Expecting actual %s, got %s", actual.String(), w.Body.String())
	}
}

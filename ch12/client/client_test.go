package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/vladimirvivien/learning-go/ch12/vector"
)

func TestClientAdd(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(
		func(resp http.ResponseWriter, req *http.Request) {
			// test incoming request path
			if req.URL.Path != "/vec/add" {
				t.Errorf("unexpected request path %s", req.URL.Path)
				return
			}
			// test incoming params
			body, _ := ioutil.ReadAll(req.Body)
			params := strings.TrimSpace(string(body))
			if params != "[[1,2],[3,4]]" {
				t.Errorf("unexpected params '%v'", params)
				return
			}
			// send result
			result := vector.New(1, 2).Add(vector.New(3, 4))
			err := json.NewEncoder(resp).Encode(&result)
			if err != nil {
				t.Fatal(err)
				return
			}
		},
	))
	defer server.Close()
	client := newVecClient(server.URL)
	expected := vector.New(1, 2).Add(vector.New(3, 4))
	result, err := client.add(vector.New(1, 2), vector.New(3, 4))
	if err != nil {
		t.Fatal(err)
	}
	if !result.Eq(expected) {
		t.Errorf("Expecting %s, got %s", expected, result)
	}
}

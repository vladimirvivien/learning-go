package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/vladimirvivien/learning-go/ch12/vector"
)

func add(resp http.ResponseWriter, req *http.Request) {
	var params []vector.SimpleVector
	if err := json.NewDecoder(req.Body).Decode(&params); err != nil {
		resp.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(resp, "Unable to parse request: %s\n", err)
		return
	}
	if len(params) != 2 {
		resp.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(resp, "Expected 2 or more vectors")
		return
	}
	result := params[0].Add(params[1])
	if err := json.NewEncoder(resp).Encode(&result); err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(resp, err.Error())
		return
	}
}

func sub(resp http.ResponseWriter, req *http.Request) {
	var params []vector.SimpleVector
	if err := json.NewDecoder(req.Body).Decode(&params); err != nil {
		resp.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(resp, "Unable to parse request: %s\n", err)
		return
	}
	if len(params) != 2 {
		resp.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(resp, "Expected 2 or more vectors")
		return
	}
	result := params[0].Sub(params[1])
	if err := json.NewEncoder(resp).Encode(&result); err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(resp, err.Error())
		return
	}
}

func dotProd(resp http.ResponseWriter, req *http.Request) {
	var params []vector.SimpleVector
	if err := json.NewDecoder(req.Body).Decode(&params); err != nil {
		resp.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(resp, "Unable to parse request: %s\n", err)
		return
	}
	if len(params) != 2 {
		resp.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(resp, "Expected 2 or more vectors")
		return
	}
	result := params[0].DotProd(params[1])
	if err := json.NewEncoder(resp).Encode(&result); err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(resp, err.Error())
		return
	}
}

func mag(resp http.ResponseWriter, req *http.Request) {
	var param vector.SimpleVector
	if err := json.NewDecoder(req.Body).Decode(&param); err != nil {
		resp.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(resp, "Unable to parse request: %s\n", err)
		return
	}
	result := param.Mag()
	if err := json.NewEncoder(resp).Encode(&result); err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(resp, err.Error())
		return
	}
}

func unit(resp http.ResponseWriter, req *http.Request) {
	var param vector.SimpleVector
	if err := json.NewDecoder(req.Body).Decode(&param); err != nil {
		resp.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(resp, "Unable to parse request: %s\n", err)
		return
	}
	result := param.Unit()
	if err := json.NewEncoder(resp).Encode(&result); err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(resp, err.Error())
		return
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/vec/add", add)
	mux.HandleFunc("/vec/sub", sub)
	mux.HandleFunc("/vec/dotprod", dotProd)
	mux.HandleFunc("/vec/mag", mag)
	mux.HandleFunc("/vec/unit", unit)

	if err := http.ListenAndServe(":4040", mux); err != nil {
		fmt.Println(err)
	}
}

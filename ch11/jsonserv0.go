package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/vladimirvivien/learning-go/ch11/curr1"
)

// Simple json server for currency service
// As in tcpserver example, returns []curr1.Currency
// This time, however, data is encoded as JSON
// Test with:
// curl-X POST -d '{"get":"Euro"}' http://localhost:4040/currency

var currencies = curr1.Load("./data.csv")

func currs(resp http.ResponseWriter, req *http.Request) {
	var currRequest curr1.CurrencyRequest
	dec := json.NewDecoder(req.Body)
	if err := dec.Decode(&currRequest); err != nil {
		resp.WriteHeader(http.StatusBadRequest)
		fmt.Println(err)
		return
	}

	result := curr1.Find(currencies, currRequest.Get)
	enc := json.NewEncoder(resp)
	if err := enc.Encode(&result); err != nil {
		fmt.Println(err)
		resp.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/currency", currs)

	if err := http.ListenAndServe(":4040", mux); err != nil {
		fmt.Println(err)
	}
}

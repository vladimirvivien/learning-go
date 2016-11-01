package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/vladimirvivien/learning-go/ch11/curr1"
)

// Simple json server for currency service
// As in tcpserver example, returns []curr1.Currency
// This time, however, data is encoded as JSON
// Test with:
// curl-X POST -d '{"get":"Euro"}' http://localhost:4040/currency

var currencies = curr1.Load("./data.csv")

// api endpoint for service
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

// serves HTML gui
func gui(resp http.ResponseWriter, req *http.Request) {
	file, err := os.Open("./currency.html")
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err)
		return
	}
	io.Copy(resp, file)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", gui)
	mux.HandleFunc("/currency", currs)

	if err := http.ListenAndServe(":4040", mux); err != nil {
		fmt.Println(err)
	}
}

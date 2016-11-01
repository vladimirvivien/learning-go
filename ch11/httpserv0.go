package main

import "fmt"
import "net/http"

// msg implements handler
type msg string

func (m msg) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Add("Content-Type", "text/html")
	resp.WriteHeader(http.StatusOK)
	fmt.Fprint(resp, m)
}

// creates server and starts
func main() {
	msgHandler := msg("Hello from high above!")
	server := http.Server{Addr: ":4040", Handler: msgHandler}
	server.ListenAndServe()
}

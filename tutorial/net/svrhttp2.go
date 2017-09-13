package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(writer http.ResponseWriter, r *http.Request) {
		fmt.Fprint(writer, "/now\n/date\n/time")
	})

	mux.HandleFunc("/now", func(writer http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(writer, "%s", time.Now())
	})

	mux.HandleFunc("/date", func(writer http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(writer, "%s", time.Now().Format("2006-01-02"))
	})

	mux.HandleFunc("/time", func(writer http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(writer, "%s", time.Now().Format("03:04:05 MST"))
	})

	server := http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  time.Second * 5,
		WriteTimeout: time.Second * 3,
	}

	fmt.Println("time api server running on :8080")
	if err := server.ListenAndServe(); err != nil {
		fmt.Println(err)
	}
}

package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, r *http.Request) {
		fmt.Fprint(writer, "/now\n/date\n/time")
	})

	http.HandleFunc("/now", func(writer http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(writer, "%s", time.Now())
	})

	http.HandleFunc("/date", func(writer http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(writer, "%s", time.Now().Format("2006-01-02"))
	})

	http.HandleFunc("/time", func(writer http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(writer, "%s", time.Now().Format("03:04:05 MST"))
	})

	fmt.Println("time api server running on :8080")
	if err := (http.ListenAndServe(":8080", nil)); err != nil {
		fmt.Println(err)
	}
}

package main

import (
	"fmt"
	"net/http"
	"regexp"
	"os"
	"log"
)

var (
	pages = "https://vladimirvivien.github.io/learning-go/"
	repo = "https://github.com/vladimirvivien/learning-go"
	exRegex = regexp.MustCompile(`^\/ch\d\d\/.+\.go$`) // regex for examples
)

func main() {
        http.HandleFunc("/", mux)

        port := os.Getenv("PORT")
        if port == "" {
                port = "80"
                log.Printf("Defaulting to port %s", port)
        }

        log.Printf("Listening on port %s", port)
        if err := http.ListenAndServe(":"+port, nil); err != nil {
                log.Fatal(err)
        }
}

func mux(w http.ResponseWriter, r *http.Request) {
	if exRegex.MatchString(r.URL.Path) {
		examples(w, r)
	}else{
	    frontpage(w, r)
	}
}

func examples(resp http.ResponseWriter, req *http.Request) {
	path := req.URL.Path
	location := fmt.Sprintf("%s/blob/master%s", repo, path)
	resp.Header().Add("Location", location)
	resp.WriteHeader(http.StatusTemporaryRedirect)
}

func frontpage(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Add("Location", pages)
	resp.WriteHeader(http.StatusTemporaryRedirect)
}

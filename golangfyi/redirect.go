package hello

import (
	"fmt"
	"net/http"
	"regexp"
)

var (
	pages = "https://vladimirvivien.github.io/learning-go/"
	repo = "https://github.com/vladimirvivien/learning-go"
	exRegex = regexp.MustCompile(`^\/ch\d\d\/.+\.go$`) // regex for examples
)

func init() {
	http.HandleFunc("/", mux)
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

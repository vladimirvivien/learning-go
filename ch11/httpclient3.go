package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	client := &http.Client{}

	req, err := http.NewRequest("GET", "http://tools.ietf.org/rfc/rfc7540.txt", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Accept", "text/plain")
	req.Header.Add("User-Agent", "SampleClient/1.0")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	io.Copy(os.Stdout, resp.Body)
}

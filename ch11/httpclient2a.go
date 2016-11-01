package main

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"time"
	"os"
)

func main() {
	client := &http.Client{
		Transport: &http.Transport{
			DisableKeepAlives: true,
			Dial: (&net.Dialer{
			   Timeout:   30 * time.Second,
	   		}).Dial,
		},
	}
	resp, err := client.Get("http://tools.ietf.org/rfc/rfc7540.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	io.Copy(os.Stdout, resp.Body)
}

package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	client := http.Client{}
	resp, err := client.Get("http://gutenberg.org/cache/epub/16328/pg16328.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	print(resp.Body)
}
func print(body io.ReadCloser) {
	defer body.Close()
	io.Copy(os.Stdout, body)
}

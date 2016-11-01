package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"os"
)

func main() {
	in := bytes.NewBufferString("Hello, zip me!")
	var out bytes.Buffer

	zip := gzip.NewWriter(&out)
	defer zip.Close()

	if count, err := io.Copy(zip, in); err == nil {
		fmt.Println("Compressed", count, "bytes.")
	} else {
		fmt.Println("Gzip failed:", err)
		os.Exit(1)
	}
}

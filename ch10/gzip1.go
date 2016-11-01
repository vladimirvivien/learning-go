package main

import (
	"compress/gzip"
	"fmt"
	"io"
	"os"
)

func main() {
	filein, err := os.Open("./gzip1.go")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer filein.Close()

	// zip content to output file
	fileout, err := os.Create("./out.gz")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer fileout.Close()

	zip := gzip.NewWriter(fileout)
	zip.Name = fileout.Name()
	defer zip.Close()

	if count, err := io.Copy(zip, filein); err == nil {
		fmt.Printf("Gzipd file %s with %d bytes", fileout.Name(), count)
	} else {
		fmt.Println("Gzip failed:", err)
		os.Exit(1)
	}
}

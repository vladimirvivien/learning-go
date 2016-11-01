package main

import (
	"compress/gzip"
	"crypto/md5"
	"crypto/sha1"
	"fmt"
	"io"
	"os"
)

func main() {
	fin, err := os.Open("./ioteerdr.go")
	if err != nil {
		fmt.Println("Unable to open file:", err)
		os.Exit(1)
	}
	defer fin.Close()

	fout, err := os.Create("./teereader.gz")
	if err != nil {
		fmt.Println("Unable to create file:", err)
		os.Exit(1)
	}
	defer fout.Close()

	zip := gzip.NewWriter(fout)
	defer zip.Close()

	sha := sha1.New()
	md := md5.New()
	data := io.TeeReader(io.TeeReader(fin, md), sha) // compose stream
	io.Copy(zip, data)

	fmt.Printf("Calculated SHA1 %x\n", sha.Sum(nil))
	fmt.Printf("Calculated MD5 %x\n", md.Sum(nil))
}

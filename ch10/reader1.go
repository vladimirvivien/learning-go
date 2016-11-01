package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type alphaReader struct {
	src io.Reader
}

func NewAlphaReader(source io.Reader) *alphaReader {
	return &alphaReader{source}
}

func (a *alphaReader) Read(p []byte) (int, error) {
	if len(p) == 0 {
		return 0, nil
	}
	count, err := a.src.Read(p) // p is populated with source data
	if err != nil {
		return count, err
	}
	for i := 0; i < len(p); i++ {
		if (p[i] >= 'A' && p[i] <= 'Z') ||
			(p[i] >= 'a' && p[i] <= 'z') {
			continue
		} else {
			p[i] = 0
		}
	}
	return count, io.EOF
}

func main() {
	str := strings.NewReader("Hello! Where is the sun?")
	alpha := NewAlphaReader(str)
	io.Copy(os.Stdout, alpha)
	fmt.Println()
}

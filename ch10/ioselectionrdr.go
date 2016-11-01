package main

import (
    "io"
    "os"
    "strings"
)

func main() {
    str := strings.NewReader("The quick brown"+
        "fox jumps over the lazy dog")
    section := io.NewSectionReader(str, 19, 23)
    io.Copy(os.Stdout, section)
}

package main

import (
    "io"
    "os"
    "strings"
)
func main() {
    str := strings.NewReader("The quick brown " +
        "fox jumps over the lazy dog")
    limited := &io.LimitedReader{R: str, N: 19}
    io.Copy(os.Stdout, limited)
}

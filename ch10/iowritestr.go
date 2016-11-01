package main
import (
    "fmt"
    "io"
    "os"
)
func main() {
    fout, err := os.Create("./iowritestr.data")
    if err != nil {
    	fmt.Println(err)
    	os.Exit(1)
    }
    defer fout.Close()
    io.WriteString(fout, "Hello there!\n")
}

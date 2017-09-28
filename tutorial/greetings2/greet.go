package main

import (
	"fmt"
	"os"

	"github.com/vladimirvivien/learning-go/tutorial/greetlib"
)

func main() {
	lang := "English"
	if len(os.Args) >= 2 {
		lang = os.Args[1]
	}
	fmt.Println(greetlib.GreetIn(lang))
}

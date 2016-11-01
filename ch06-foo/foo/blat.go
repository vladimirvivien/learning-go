package foo

import (
	"fmt"
	"foo/bar/bazz"
)

func fooIt() {
	fmt.Println("Foo!")
	bazz.Qux()
}

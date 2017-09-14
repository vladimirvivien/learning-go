# Your first Go program
To get started with Go, let us write the obligatory *Hello World* program.  For now, create directory  `$HOME/go/src/hello` .  Later we will have a full discussion on code organization and Go packages.
```
mkdir -p $HOME/go/src/hello
```
Next create and save the following code in a file called  `hello_world.go`.
```go
package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
}
```
After the source code is saved, we can compile and run it with the `go run` command as shown in the following example:
```sh
$> go run hello_world.go
Hello, World!
```
The `go run` command, shown above, is a convenience tool that compiles of the source code and immediately runs the resulting program.   This command can be handy when prototyping ideas or running a simple Go program.  As we will see later though, there are other and more suitable ways to compile your Go code and packages for distribution.

### Next
Continue to the next section to read about the [Go source file](./source_file.md).
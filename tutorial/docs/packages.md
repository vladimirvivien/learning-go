# Packages
In Go, a package is the unit of code that can be compiled either as an executable `program` or a reusable code `library`.  Packages are directories in the workspace under `$HOME/go/src`.  A package can consist of one or more source files which are compiled into one logical unit.  The membership of a source to a package is declared in the with the `package` directive (discussed later).  All files in a package must declare the same package name or the compiler will not be happy.  

## Package import path and the default name
The `import path` of a package is the unique directory path of the package in the workspace, relative to path `$HOME/go/src` as shown in the following table:

| Workspace Path | Import Path | Default Name
|---|---|---|
|`$HOME/go/src/foo`|`foo`|`foo`|
|`$HOME/go/src/foo/bar`|`foo/bar`|`bar`|
|`$HOME/go/src/foo/bar/bazz`|`foo/bar/bazz`|`bazz`|

The previous table shows examples of workspace paths, their resolved import paths, and their default package name.  These concepts are crucial to Go and its tools.  The import path is used in Go source files when specifying packages to `import` for instance.  The default package name is resolved as the last element of the import path and is used in source files as an identifier for the package.  For instance, the following example uses import path  `foo/bar/bazz` which is assigned package identifier `bazz` to access package elements using  a dot-notation:
```go
package main
import "foo/bar/bazz"
func main() {
    bazz.Blat()
}
...
```
It should be noted that the default package name in a source file can be explicitly specified by preceding the import path with an identifier as shown below:
```go
package main
import ba "foo/bar/bazz"
func main() {
    ba.Blat()
}
...
```
## Naming your packages
An accepted practice in Go is to give the package path a unique name to avoid name collisions.  This is specially important if you plan to distribute your code for others to consume.  The most common approach is to include a unique identifier such as a source code repository and username as part of the path.  Others also use a company name or a project name, when naming the package directory.  

|Import path|Qualifier|
|---|---|
|github.com/vladimirvivien/go-tutorial|github.com/vladimirvivien|
|github.com/stretchr/testify|github.com/stretchr|
|k8s.io/client-go/pkg/api/v1|k8s.io/client-go|
|gopkg.in/yaml.v1/|gopkg.in|

## A Go program
As mentioned earlier,  a program is a package that can be compiled into executable code.  All source files of a program must declare `package main` and at most one source file in the directory must include the special function `main()`.  

For instance, the following shows the layout for a program in directory  `greetings/` :  
```sh 
$HOME/go/src/
 +-github.com/vladimirvivien/go-tutorial/
   +-greetings/
     +-greet.go
``` 

The program package directory is `/greetings` with one source file,  `greet.go`, shown below:  
```go
package main

import (
	"fmt"
	"os"
)

var greetings = map[string]string{
	"English": "Hello, World!",
	"French":  "Salut Monde",
	"Chinese": "世界您好",
	"Klingon": "qo' vIvan",
	"Hindi":   "हैलो वर्ल्ड",
	"Korean":  "안녕하세요",
	"Russian": "привет мир",
	"Swahili": "Wapendwa Dunia",
	"Spanish": "Hola Mundo",
	"Turkish": "Merhaba Dünya",
}

func main() {
	lang := "English"
	if len(os.Args) >= 2 {
		lang = os.Args[1]
	}
	if greeting, ok := greetings[lang]; ok {
		fmt.Println(greeting)
	} else {
		fmt.Println(greetings["English"])
	}
}
```
### Compiling and running the program
We can compile the program package, along with its dependencies, using the `go build` command-line tool by specifying the relative path of the package or its *import path*.  For instance, the following will compile the program in package `greetings`:
```sh
$ cd $HOME/go/src/github.com/vladimirvivien/go-tutorial/greetings
$ go build .
```
Or,
```sh
$ cd $HOME/go/src/github.com/vladimirvivien/go-tutorial
$ go build ./greetings
```
In the previous commands are equivalent. They `.` specify the relative directory path of package `greetings` to compile.  By default, the `go build` command creates a binary with the same name as the package.

```sh
> ls -l
total 1520
-rw-rw-r-- 1 vvivien vvivien     582 Jul  8 08:30 greet.go
-rwxrwxr-x 1 vvivien vvivien 1551885 Jul  8 09:05 greetings
```
We can run the greetings program as follows:

```sh
> ./greetings Korean
안녕하세요
```
>Note that we can use the `-o` flag to specify the name of the binary compiled. 

We can also compile the program by specifying its full import path:

```sh
$ go build github.com/vladimirvivien/go-tutorial/greetings
```
We can also use `go install` which is a tool that compiles the package and its dependencies and installs the resulting binary in `$HOME/go/bin`.  This command also caches any dependent packages in the workspace to avoid future unnecessary compilation.

> See `go help build` and `go help install` for detail.
 
## A Go library
Libraries are packages that use the same `go` command tools and are compiled into archive files (instead of executable code) that can be reused by other packages.  To demonstrate a library, we will rewrite the previous greeting program.  In this version, we will extract the greeting functionality and place it into library `greetlib` so that it can be imported by other packages:  
```sh 
$HOME/go/src/
 +-github.com/vladimirvivien/go-tutorial
   +-greetlib
     +-lib.go
   +-greetings2/
     +-greet.go
```
The source code for the library package is stored in source file `greetlib/lib.go` as shown below:
```go
package greetlib

var greetings = map[string]string{
	"English": "Hello, World!",
	"French":  "Salut Monde",
	"Chinese": "世界您好",
	"Klingon": "qo' vIvan",
	"Hindi":   "हैलो वर्ल्ड",
	"Korean":  "안녕하세요",
	"Russian": "привет мир",
	"Swahili": "Wapendwa Dunia",
	"Spanish": "Hola Mundo",
	"Turkish": "Merhaba Dünya",
}

// GreetIn returns a greeting in specified lang
func GreetIn(lang string) string {
	if greeting, ok := greetings[lang]; ok {
		return greeting
	}
	return greetings["English"]
}
```
As a convention, and to make things easy, the source files of a package declare a package name that matches the directory where they are located.  The previous source snippet, for instance, declares `package greetlib` since the directory where the file is located is called `greetlib`.

The program which uses the library is in package `greetings2`.  The source code in that package imports the library by specifying its import path as `github.com/vladimirvivien/go-tutorial/greetlib` to access its exported code elements.

```go
package main

import (
	"fmt"
	"os"
	
	"github.com/vladimirvivien/go-tutorial/greetlib"
)

func main() {
	lang := "English"
	if len(os.Args) >= 2 {
		lang = os.Args[1]
	}
	fmt.Println(greetlib.GreetIn(lang))
}
```
### Compiling a library
Compiling a library is done using the `go build` tool by specifying the package's import path or its relative directory path as was done before.  For instance, the following would compile the `greetlib` package:

```sh
$ cd $HOME/go/src/github.com/vladimirvivien/go-tutorial/greetlib
$ go build . 
```
Or we can ensure that the compiled artifacts is cached in the workspace and is available for other packages by using the `go install` command:
```sh
$ cd $HOME/go/src/github.com/vladimirvivien/go-tutorial/
$ go install ./greetlib
```
It should be noted that when we compile the program, the `go` tool will properly resolve the dependency to the `greetlib` library and compile that as well.  So, the following will compile and install both the program and its dependent library together:
```
$ cd HOME/go/src/github.com/vladimirvivien/go-tutorial/
$ go install ./greetings2
```

> See `go help build` and `go help install` for detail.

## Package element visibility
Go has a simple rule for element visibility when a package is imported from another:

>*Capitalized identifiers are visible to other packages*

For instance, we can see this in action from the `greetlib` package example above. Variable `greetings` is an identifier with lowercase, therefore it cannot be accessed from other packages.

```go
var greetings = map[string]string{
	"English": "Hello, World!",
	"French":  "Salut Monde",
	"Chinese": "世界您好",
	...
}
```
On the other hand, function `GreetIn()` , shown below is capitalized which means it can be accessed by other packages.

```go
func GreetIn(lang string) string {
	if greeting, ok := greetings[lang]; ok {
		return greeting
	}
	return greetings["English"]
}
```
## Remote packages
Go includes the `go get` tool which can retrieve and install packages stored on a remote source control repository server such as Git or Mercurial.  The tool uses the import path to figure out where the file is located on the server.  For instance, the following command will pull and install package `greetings2` from this repository:
```
$ go get github.com/vladimirvivien/go-tutorial/greetings2
```
Since package `greetings2` imports package `github.com/vladimirvivien/go-tutorial/greetlib`,  `go get` will transitively attempt to resolve, download, and install package `greetlib` if it is not found in the workspace.

> see `go help get` for more detail.

### Next
Continue to the next section to learn about [Go language fundamentals](./fundamentals.md).

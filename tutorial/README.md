# Getting Started With Go 
This document is a collection of topics designed to help newcomers to get started with the Go programming language.  All code samples, discussed in the sections below, can be found in this repository.

## About Go
Go was created as system language at Google in 2007 primarily by Robert Griesemer, Rob Pike, Ken Thomson.  The language was an answer to handle the needs of application development at Google-scale. The designers of Go wanted to mitigate issues such as slow build cycle, language complexity, magical syntax while creating a new language that is simple, safe, consistent, and predictable. 

> “Go is an attempt to combine the safety and performance of a statically-typed language with the expressiveness and convenience of a dynamically-typed interpreted language.” -- Rob Pike

- BSD-style licensed open source language
- It borrows ideas from many languages before it
- Go has a simple and concise syntax (that slightly resemble C)
- Short mortal-decipherable spec document
- Strongly and statically typed for safety
- Type declaration with dynamic language feel
- Support for objects (not class-based)
- Near-zero compilation time with script language-like performance
- Compiled to native binaries for fast runtime execution
- Easy cross compilation for popular OS and architectures
- Uses channels with send-receive idiom
- Simple concurrency idioms

## Where to start
There are several ways to start with the Go programming language

- Download Go tools from [golang.org/dl](https://golang.org/dl/) [download and install the Go toolchain.  Configure your local environment and use your own editor to craft sweet Go code]
- Try it out at the Playground ([play.golang.org](https://play.golang.org)) [ for instant gratification, you can try this browser-based IDE that let you run your own code to test and share ideas]
- Take the Go Tour at [tour.golang.org](https://tour.golang.org) [ a step-by-step exploration of the features of the language with browser-based editor]

### Supported OSs / Architectures
The Go compiler can target 8 architectures: 

- arm
- arm64
- mips
- mips64
- ppc64
- s390
- x86
- x86_64

The Go compiler can target 9 operating systems:

- DragonFly BSD
- FreeBSD
- Linux
- NetBSD
- OpenBSD
- OS X
- Plan 9
- Solaris
- Windows

This means you can create your program from the comfort your Mac running OS X on a x86_64 and compile that code to run on Linux running on an ARM6.

### Installing Go
Go can be installed on your local machine in either of the following ways:
- From archive files - for Linux, Mac OS, FreeBSD, and Windows
- Installer - available for Mac OS and Windows
- Build from source 
Use URL [https://golang.org/doc/install](https://golang.org/doc/install) to follow instruction on your preferred method of installation for your target environment.

### Setup your workspace
Before you get introduced to your first Go program, let us take a moment to discuss your environment for local development.  After you install your Go toolchain, it is crucial that you properly setup your workspace.  It is an arbitrary directory which stores your Go source files, organized as packages, and built artifacts (such as object files and executable binaries) as shown in the following sample workspace:
```sh
$HOME/go
 +- src/
 |  +- github.com/vladimirvivien/go-tutorial
 |  |  +- hello/hello_world.go 
 |  |  +- greetings2/greet.go
 |  +- github.com/vladimirvivien/automi
 |  |  +- operators/batch/funcs.go
 |  |  +- stream/stream.go
 |  + ... (may contain many more source packages) ...
 |
 +- pkg/
 |  +- darwin_amd64
 |     +- github.com/vladimirvivien/go-tutorial/greetings2.a
 |     +- github.com/vladimirvivien/automi/operators/batch.a
 |     +- github.com/vladimirvivien/automi/stream.a
 |     +- ... (each package compiled as an object file) ...
 |
 +- bin/
    +- hello
    +- ... (compiled programs stored here) ...
```
### Setting GOPATH
By convention (and for the remainder of this discussion), the Go toolchain assumes the workspace is located at `$HOME/go`. If different, you will need to set env variable `$GOPATH` pointing to the workspace location.  You can use the `go env` command to verify the location of your `GOPATH` as shown below.
```sh
go env GOPATH
/Users/vvivien/DEV/go
```

### Editors and IDE
Because of its minimal syntax, working with Go can be done using your favorite text editor.  Some editors include features such as syntax higlights while others have deep integrations with the Go toolchain to provide an IDE-like experience.  Some of the more well-known editors with Go support include:
- Atom
- Emacs
- Sublime Text
- TextWrangler
- Vim/Neovim
- ... and more

If you feel more comfortable with working with a full IDE, you will be well-served there as well.  There are several IDE's providing full support for Go including the followings:
- IntelliJ Gogland - Go IDE
- Visual Studio Code - Support with vscode-go plugin
- GoClipse - Eclipse based Go IDE

## Your first Go program
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

## The Go source file
It is important to understand the make up of a Go source file.  The following figure highlights the major attributes of a Go source file.

![Go Source](go-source.png)

#### `package`
All Go source files must start with the he `package` directive.  It declares the name of a package to which the source in the file belongs.  In the source snippet above, the source belongs to package `main` which is a reserved name for packages that get compiled to runnable binaries.  

#### `import`
The `import` statement specifies an  `import path` for packages that are used in the source code.  Your code can import from the standard library or other user-provided packages that are in the workspace.  The code above imports package `fmt` so that it can invoke function `fmt.Println`.

#### `func main()`
Go functions are declared using the `func` keyword.  Function `main` is a special function that is used as the entry point of an executable program.  Function `main` must be defined in the main package.

#### File `hello_world.go`
Go source files can have arbitrary names followed by the `.go` extension.  There is no relations between the code elements in the source file and its name.  However, by convention, the file is usually named something meaningful that is usually kept short and use the underscore (`_`) as word separator.

#### `// Comment`
Go uses C-style comments which are used by Go tools to generate documentation automatically.

#### Optional semi-colon
One more thing that is notable in Go sources is the lack of semi-colon.  In idiomatic Go, semi-colons are optional and are always omitted.  However, the Go compiler inserts them during compilation as they are required by Go's formal grammar.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 

## Packages
In Go, a package is the unit of code that can be compiled either as an executable `program` or a reusable code `library`.  Packages are directories in the workspace under `$HOME/go/src`.  A package can consist of one or more source files which are compiled into one logical unit.  The membership of a source to a package is declared in the with the `package` directive (discussed later).  All files in a package must declare the same package name or the compiler will not be happy.  

### Package import path and the default name
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
### Naming your packages
An accepted practice in Go is to give the package path a unique name to avoid name collisions.  This is specially important if you plan to distribute your code for others to consume.  The most common approach is to include a unique identifier such as a source code repository and username as part of the path.  Others also use a company name or a project name, when naming the package directory.  

|Import path|Qualifier|
|---|---|
|github.com/vladimirvivien/go-tutorial|github.com/vladimirvivien|
|github.com/stretchr/testify|github.com/stretchr|
|k8s.io/client-go/pkg/api/v1|k8s.io/client-go|
|gopkg.in/yaml.v1/|gopkg.in|

### A Go program
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
#### Compiling and running the program
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
 
### A Go library
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
```
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
#### Compiling a library
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

### Package element visibility
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
### Remote packages
Go includes the `go get` tool which can retrieve and install packages stored on a remote source control repository server such as Git or Mercurial.  The tool uses the import path to figure out where the file is located on the server.  For instance, the following command will pull and install package `greetings2` from this repository:
```
$ go get github.com/vladimirvivien/go-tutorial/greetings2
```
Since package `greetings2` imports package `github.com/vladimirvivien/go-tutorial/greetlib`,  `go get` will transitively attempt to resolve, download, and install package `greetlib` if it is not found in the workspace.

> see `go help get` for more detail.

## Language Fundamentals
This section covers the fundamentals of the Go language including data types, variable declaration, and other language construct that are crucial in understanding the language.

### Variables
Go is a *strongly typed* language where all variables must have a value and a type.  When a variable is declared, it must receive a type and a value.  The following shows a long form of the declaration where the type is explicitly provided and the value is subsequently given:
```go
package main

import "fmt"

var name, desc string		// declares two variables of type string
var radius int32			// variable of type int32
var mass float64			// variable of type float64
var active bool				// variable of type bool
var satellites []string		// variable of type []string

func main() {
	name = "Sun"
	desc = "Star"
	radius = 685800
	mass = 1.989E+30
	active = true
	satellites = []string{
		"Mercury",
		"Venus",
		"Earth",
		"Mars",
		"Jupiter",
		"Saturn",
		"Uranus",
		"Neptune",
	}
	fmt.Println(name)
	fmt.Println(desc)
	fmt.Println("Radius (km)", radius)
	fmt.Println("Mass (kg)", mass)
	fmt.Println("Satellites", satellites)
}
```
The previous program shows the *long way*	of declaring variables without explicit initial values.  Each type however, has a default value, known as the *zero-value*, that is assigned to the variable when no explicit initialized values are provided.  For numeric values it is `0`, for string values it's the empty string `""`, boolean value is `false`.

The language also offers an expressive syntax for variable declaration, that can feel like dynamic language,  where the type can be inferred and the value can be assigned in one statement as shown below.

```go
var name = "Mars"		// inferred as type string
var desc = "Planet"		// inferred as type string
var radius = 3396.2		// inferred as type float64
var mass = 6.4185e23	// inferred as type float64
var active = true		// inferred as type bool
```
When variables are declared inside a function, the declaration can get even shorter by dropping the `var` keyword as shown in the following example.  This forms combines the declaration and initialization of variable in one step.  Even without the type information, the compiler uses the literal text to infer a type for each variable.  This is one of the most idiomatic version of type declaration that you will encounter.

```go
package main

import "fmt"

func main() {
	name := "Neptune"    // type string
	desc := "Planet"     // type string
	radius := 24764      // type int
	mass := 1.024e26     // type float64
	active := true       // type bool
	
	satellites := []string{
		"Naiad", "Thalassa", "Despina", "Galatea"
		"Triton", "Nereid", "Halimede", "Sao",
	}
...
}
```
> Note that operator `:=` only initializes the variable.  Further update of the variable must be done using the `=` operator.

#### The blank identifier `_`	
A declared variable must be subsequently used in an expression or a statement or failure to do so will result in a compilation error.  For instance, the following snippet will not compile because variable `b` is declared and not used.
```go
func main() {
    a := 12
    b := a
    fmt.Println (a)
}
```
In some situations however, you may need to use a placeholder instead of an actual variable. Go supports the use of a blank identifier `_` which can be bound to a value and does not require subsequent usage.  The blank identifier provides an idiom where values are discarded as shown in the following snippet:
```go
package main

import (
	"fmt"
	"path"
)

func main() {
	_, f := path.Split("a/b/c/file.ext")
	fmt.Println(f)
}
```
In the previous example, function `path.Split(path)(dirName,fileName)` returns two values `dirName` and `fileName`.  In the code, however, we discard of the value for `dirName` by binding it to the blank identifier.
 
### Primitive data types
Go support several *numeric types*:
- *Signed* integers: `int8`, `int16`, `int32`, `int64` , and `int`
- *Unsigned* integers: `uint8`, `uint16`, `uint32`, `uint64` , and `uint`
- *Character representation*: type `rune` an `int32` alias
- *Byte values*: type `byte` an alias for `uint8`
- *Floating point* types `float32` and `float64`
- *Complex numbers* types `complex64` and `complex128` 

*Boolean type*:
- Go has a boolean type `bool` representing values `true` or `false`.

*String type*:
To represent text Go uses type `string` which stores a sequence of `rune` capable of UTF-8 encoded string values.

Primitive type examples: 
```go
var color uint32 = 0xFEFEFE  	// hex, uint32
var mod = 0466               	// octal, int
count := 1245                	// decimal, int
avogadro := 6.0221409e+1	 	// float64
value := "automobile"			// string
```
### Composite types
Composite types are used to store sequences of values of primitive types.  Composite literals values are contained within curly-braces preceded by the type as shown below.

#### *Array*
Type *array* represents a fixed-size sequenced values numerically indexed.  
```
func main(){
    steps := [3]string{"SEND", "RCVD", "WAIT"} 	// size 3 array, initialized
    fmt.Println(steps[1]) 			// prints "RCVD" 
    steps[3] = "PAUSE"				// index out of range error
}
```
The size of an array remains static throughout its lifetime.  Attempt to access data beyond the declared bound of the array will result in a runtime error as shown above.  

Go provides `for-range` construct to walk the items of an array.  With each iteration, the for-range expression emits the current index and the value of the item in the array as shown below:
```go
func main(){
    steps := [3]string{"SEND", "RCVD", "WAIT"}
    for index, value := range steps {
        fmt.Printf("step %d = %s\n", index, value)
    }   
}
```

#### *Slice*
A slice is a dynamically-sized array.  The slice omits its size as part of its type declaration as shown in the snippet below.   Slices can be initialized with a composite literal or with the `make()` built-in function:
```
steps := []string{"SEND", "RCVD", "WAIT"} 	// slice initialized with 3 elements
fmt.Println(steps[1]) 						// prints "RCVD" 
steps = append(steps, "PAUSE")				// slice expanded to size 4
steps[3] = "RESUME"							// updates value at index 4

actions := make([]int, 2)					// initializes a slice of size 2
actions[0] = "PRINT"
actions[1] = "LOG"
actions = append(actions, "ADD")			// expands slice, size now 3
```
Go also supports slice expressions which can be used to create new slices from arrays or other slices.  For instance, slice `summer` is created by slicing existing array `months`
```
months := [12]string{
    	"Jan", "Feb", "Mar", "Apr", "May", "Jun",
    	"Jul", "Aug", "Sep", "Oct", "Nov", "Dec",
}
summer := months[5:8]
```
Similar to arrays, Go can use the `for-range` construct to iterate over slice items as shown in the following example:
```go
func main() {
	months := [12]string{
		"Jan", "Feb", "Mar", "Apr", "May", "Jun",
		"Jul", "Aug", "Sep", "Oct", "Nov", "Dec",
	}
	summer := months[5:8]
	
	fmt.Println("--Summer Months--")
	
	for _, month := range summer {
	    fmt.Println(month)
	}
}
```
Note that in this example, we assign the index value of the item to the blank identifier, `_`, so that we can ignore it.

#### *Map*
A map is a dynamically-sized composite type that stores elements of arbitrary types that are indexes using a values of  type.  A map can be initialized using a composite literal:
```
ratings := map[string][]int{
	"men":   {32, 55, 12, 55, 42, 53},
	"women": {44, 42, 23, 41, 65, 44},
}

ratings["children"] = []int{2,34,5,43,64,22}
```
A map can also be initialized using the built-in function `make()` as shown below:
```
hist := make(map[string]int)
hist["Jan"] = 100
hist["Feb"] = 445
hist["Mar"] = 514
```
The `for-range` statement can also be used with map values.  In this context, the range expression emits the key and the value of each element in the map as shown in the following example:
```go
func main() {
	hist := make(map[string]int)
	hist["Jan"] = 100
	hist["Feb"] = 445
	hist["Mar"] = 514
	
	for key, value := range hist {
	    fmt.Printf("histogram[%s] = %d\n", key, value)
	}
}
```
#### *Struct*
The *struct* type is a composite that stores named elements of diverse types known as fields.  The following example creates variable `truck` as type `struct{year int; make, model string}` and initializes it with a composite literal.
```
func main() {
	truck := struct {
		year        int
		make, model string
	}{
		make:  "Ford",
		model: "F150",
		year:  2017,
	}

	fmt.Printf("Truck: %d %s %s\n", truck.year, truck.make, truck.model)
}
```
The struct uses the dot notation to access field members of the struct.

### The pointer type
Go supports a type pointer which is a value that may be used to reference the memory address where the data is located. Go uses the `*` operator to designate a type as a pointer of that type.  The followings are examples of declaration of pointer type where `scorePtr` is a pointer to type `float32`:
```
var scorePtr *float32
```
Pointer variables can only be assigned address values of its declared type.  One way to do so in Go is to use the address operator `&` (ampersand) to obtain the address of a variable as shown in the following example:
```
score := 32
scorePtr = &score		// pointer assigned address
*scorePtr = 44			// pointer dereferenced with value
```
While Go only support pass-by-value when calling a function/method, pointers can be used to create a pass-by-reference idiom.  For instance, in the following, variable `score` will not be updated after a call to function `adjust()` because the function receives a copy of the value via parameter `val`:
```go
func main() {
	score := 32
	adjust(score)
	fmt.Println(score)    // score not updated
}

func adjust(val int) {
	val = val * 4
}
```
However, we can modify function `adjust()` to receive a value of type `*int` instead which causes a copy of the address of the original value to be passed to the function thus mimicking pass-by-reference as shown below.
```go
func main() {
	score := 32
	scorePtr := &score
	adjust(scorePtr)
	fmt.Println(score)    // score is updated
}

func adjust(val *int) {
	*val = *val * 4
}
``` 


### The function type
In Go, a function is also a type that can be assigned to a variable or stored for later use.  A function can be *named* or be assigned to a identifier as shown in the following example:
```go
func main() {
	printLn := func(val string) {
		fmt.Println(len(val))
	}
	run(printLn)
}

func run(f func(string)) {
	if f != nil {
		f("Hello")
	}
}
```
In Go, a function can *return* a list of value of different types.  This idiom is often used as a way of handling errors from a function (or method) call.  For instance, the following function returns two values, one is the expected `int` value, and the other one is an `error` type used to signal any exceptional faults caused by the function call.
```go
package main
import (
	"fmt"
	"os"
	"errors"
)

func main() {
	val, err := div(4, 0)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(val)
}

func div(op0, op1 int) (int, error) {
	if op1 == 0 {
		return 0, errors.New("div by zero")
	}
	return op0 / op1, nil
}
```
### Methods 
Methods are functions that are attached to a type.  Most Go types can receive a function via a special parameter, called a receiver parameter, that associate the function to the type.  The following example shows that type `*car` can receive method `drive()`:
```go
type car struct {
	make,
	model string
}

func (c *car) drive() {
    fmt.Println("driving a", c.make, c.model)
}

func main() {
	ford := &car{
		make:  "Ford",
		model: "F150",
	}

	ford.drive()
}
```
So, variable `ford` of type `*car` can invoke function `ford.drive()`.

### Deferring functions and method calls
Function (or method) calls can be deferred using the `defer` statement which is called right before returning from the surrounding function.  Deferring function calls is an idiom common in Go to implement lifecycle logic such as clean up logic, closing a network connections, closing channel, deleting unwanted test files, or any other tasks that should always happen when the called function is returning.  

The following uses a defer statement to print a separator right before function `print()` exits.
```go
func main() {
    print([]string{"Hello", "Goodbye!"})
    print([]string{"North", "South", "East", "West"})
    print([]string{"Sad", "Happy", "Cry"})
}

func print(data []string) {
    defer fmt.Println("----")
    for _, d := range data {
        fmt.Println(d)
    }
}
```
A more realistic example shows a simple and HTTP client that retrieves the full text of literary classic Beowulf from the Project Gutenberg's website.  Notice how `defer` is used, in function `print()` to close the resource before exiting the function. 
```go
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	client := http.Client{}
	resp, err := client.Get("http://gutenberg.org/cache/epub/16328/pg16328.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	print(resp.Body)
}
func print(body io.ReadCloser) {
	defer body.Close()
	io.Copy(os.Stdout, body)
}
```
### Type declaration
Go allows a type declaration to receive an identifier so that the type may be reused by referring to its name.  For instance, type `struct{year int; make, model string}` can be assigned a name `car` so that subsequent variable declarations only needs to use the type name as shown below.

```go
package main
import "fmt"

type car struct {
    year int,
	make,
	model string
}

func main() {
	ford := car{
	    year: 2001,
		make:  "Ford",
		model: "F150",
	}

	fmt.Println(ford.make, ford.model)

	chevy := car{
	    year: 2012,
		make:  "Chevrolet",
		model: "Silverado",
	}

	fmt.Println(chevy)
}
```
### Interfaces 
Another feature that gets celebrated in Go are interfaces.  An interface in Go is a type that represents a set of zero or more methods.  Any other type that implements that set of methods automatically implements the interface.  For instance, built-in package `io` defines interfaces `io.Reader` and `io.Writer` for IO input and output respectively. 
```go
// io.Reader
type Reader interface {
    Read(p []byte) (n int, err error)
}

// io.Writer
type Writer interface {
    Write(p []byte) (n int, err error)
}
```
Therefore, any type that implements method `Read(p []byte) (n int, err error)` automatically implements `io.Reader` and any type that implements method `Write(p []byte) (n int, err error)` is considered an `io.Writer` by the type system.  

For instance, built-in function `io.WriterString(w io.Writer, s string)`, from the `os` package, uses any value `w` that implements `io.Writer` to output a string.  We use that function in wrapper function `WriteData(io.Writer,string[])` to output the string slice to any value that implements the io.Writer interface:
```go
package main

import (
	"fmt"
	"io"
	"os"
	"bytes"
)

func main() {
    proverbs := []string{
        "Channels orchestrate mutexes serialize", 
        "Cgo is not Go",
        "Errors are values",
        "Don't panic",
    }
    
    // write to in-mem buffer
    buf := new(bytes.Buffer)
    WriteData(buf, proverbs)

    // write to stdout
    WriteData(os.Stdout, proverbs)

    // open file for writing
    file, err := os.Create("./textfile.txt")
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    defer file.Close()

    // write to file
    WriteData(file, proverbs)
}

func WriteData(writer io.Writer, data []string) {
    for _, d := range data {
        io.WriteString(writer, d)
    }
}
```
The previous code uses function `WriteData(writer io.Writer, data []string)` to output string slice `data` to different target including an in-memory buffer `buf` of type `*bytes.Buffer`, standard output `os.Stdout`, and file object `file` of type `*os.File` which are all values that implement `io.Writer`.

### Flow control
Go supports the expected flow control from a modern language for branching and looping as outlined in the his section.
#### if statement
As you would expect, Go has the normal *if* statement as shown below:
```go
if len(os.Args) >= 2 {
	lang = os.Args[1]
}
```
Another idiomatic version of the `if` statement uses an initializer expression as shown below:
```go
if result, err := div(4, 0); err != nil {
	fmt.Println(err)
}
```
While this version of the if statement it compact, it captures the variables which go out of scope at the end of the if statement.

#### switch statement
Go supports multi-way branching using a `switch` statement as shown in the following example.  Notice the case clause can evaluate any comparable types (not just numeric) and multiple cases can be evaluated separated by a comma:  
```go
func main() {
	states := []string{"S", "P", "P", "E", "", "X"}
	for _, s := range states {
		fmt.Println(next(s))
	}
}

func next(state string) string {
	switch state {
	case "S":
		return "START"
	case "P":
		return "PROC"
	case "E":
		return "PAUSE"
	case "X", "H":
		return "STOP"
	default:
		return "CONTINUE"
	}
}

```
Go also supports an expression-less switch statement that can be used as a replacement for if-else chains:
```go
func main() {
	a := 12
	b := 0
	c := b
	switch {
	case a == b:
		fmt.Println("same")
	case c == b, b == 0:
		fmt.Println("starting")
	default:
		fmt.Println("nothing to do")
	}
}
```
#### for statement
Go offers the traditional *for* statement that loops sequentially after testing a given condition. 
The following shows several forms of the for statement:
```go
// semantically equivalent to while loop
func main() {
	a := 12
	for a > 6 {
		fmt.Println(a)
		a--
	}
}
```
The *for* statement can also include an initializer and update condition as shown below:
```go
func main(){
    nums := []int{2, 34, 5, 43, 64, 22}
    for i := 0; i < 6; i = i + 2 {
	    fmt.Println(nums[i])
    }
}
```
## Concurrency with goroutines and channels
Concurrency is considered to be the one of the most attractive features of the Go programming language.  Its adopters revel in the simplicity of its primitives to express correct concurrency.  Go has its own concurrency primitive called the *goroutine* that allows a program to launch a function/method (or routine) to execute independently from its calling function. 

For instance, the following program uses the *go* statement to launch three *goroutines* that get executed concurrently, along with the main function.  The `fmt.Scanln()` function call is used for its side-effect of blocking the `main` function which causes it to wait for the goroutines to have time to complete.
```go
package main

import (
	"fmt"
)

func main() {
	go count(10, 30, 10)
	go count(40, 60, 10)
	go count(70, 120, 20)
	fmt.Scanln() // blocks for kb input, used only for side effect
}

func count(start, stop, delta int) {
	for i := start; i <= stop; i += delta {
		fmt.Println(i)
	}
}
```
### Channels
One area where Go diverges from other modern programming languages is the way it handles synchronization and data sharing between running goroutines.  Go uses a primitive known as a *channel* as a conduit that can send and receive data to communicate between goroutines.  This notion is captured in the popular slogan:

> Do not communicate by sharing memory; instead, share memory by communicating

Channels are strongly typed entities that only allow data of the declared type to pass through.  Go uses the `<-` operator to send or receive data from a channel depending on where the channel operand is placed as illustrated below:
```go
ch <- 5 // sends 5 to channel variable ch
value := <- ch // receives 5 from channel
```

Channels can be declared to be either *buffered* or *unbuffered*, each with its own characteristics. 
For instance, the following declares and initializes variable `intChan` as an unbuffered channel of type integer while `bufChan` is a buffered channel of size 5.
```go
intChan := make(chan int)		// unbuffered
bufChan := make(chan bool, 5)	// buffered
```
A buffered channel can receive up to `N` items after which subsequent send operations will block until the channel is drained by at least `N-1` item.  For instance, in the following snippet channel `intsCh` can receive up to 3 items before it blocks.
```
func main() {
    intsCh := make(chan int, 3)
    intsCh <- 12
    intsCh <- 5
    intsCh <- 55
    
    // intsCh <- 90       // this would block
    
    fmt.Println(<- ch)    // drain N-1
    intsCh <- 44          // works
}
```

An unbuffered channel blocks immediately after a send operation until the item is received.  For instance, the following would cause a deadlock immediately:
```go
func main() {
    intsCh := make(chan int)
    intsCh <- 5			// blocks main() forever
    fmt.Println(<- intsCh) 
}
```
A simple strategy to avoid deadlock when working with channels (unbuffered or buffered) is to place send operations in their own goroutine.  For instance, the previous is re-written where it does not block function `main()`:

```go
func main() {
    c := make(chan int)
    go func(){
        c <- 5  // send 5
    }()
    
    fmt.Println(<- c) // receives 5
}
```

One common usage of unbuffered channels is the synchronization between  goroutines.  In the following example, `doneCh` is declared as a channel that can receive and send elements of type `bool`.  The channel is used as a synchronization mechanism between functions `main()` and goroutine `go display()`.
```go
package main

import (
	"fmt"
)

func main() {
    doneCh := make(chan bool)
    go display(doneCh) // launches new goroutine
    <-doneCh //waits until a value is received
}

func display(done chan bool) {
   for _, char := range "Channel Synchronization" {
       fmt.Println(string(char))
   }
   done <- true  
}
```
In the previous example, function `main()` will block and wait at receive operation `<-doneCh` until the goroutine, launched by `go display()`, executes send operation `done <- true`.  

Another important channel characteristic we can use, when doing synchronization, is the fact that a closed channel does not block a receive operation.  So, we can rewrite the previous example by closing the channel instead of sending a dummy value as shown below:
```go
func main() {
    doneCh := make(chan bool)
    go display(doneCh) // launches new goroutine
    <-doneCh //waits until a value is received
}

func display(done chan bool) {
   for _, char := range "Channel Synchronization" {
       fmt.Println(string(char))
   }
   close(done)  // channel is closed
}
```
## Data IO and communication
As mentioned earlier in the *Interfaces* section, Go uses interfaces  `io.Reader` and `io.Writer` to achieve streaming data input and output.  This section provides more examples on how to the built-in packages for data input and communication.

### Reading formatted data Standard In
The `fmt` package can be used to read (or *scan*) formatted text input from standard input (provided as `os.Stdin`).  The following code shows how to use `fmt.Scanf()` to input a single integer value from standard in:
```go
func main() {
	var choice int
	fmt.Println("A square is what?")
	fmt.Print("Enter 1=quadrilateral 2=rectagonal: ")

	n, err := fmt.Scanf("%d", &choice) // uses global reader os.Stdin
	if n != 1 || err != nil {
		fmt.Println("invalid choice")
		os.Exit(1)
	}
	if choice == 1 {
		fmt.Println("You are correct")
	} else {
		fmt.Println("sorry")
	}
}
```
### Reading formatted data from any io.Reader
We can `fmt.Fscanf()` to read formatted text from an arbitrary reader.  Luckily, Go provides `*os.File` type, which implements `io.Reader`, represent a file from the filesystem.  Once we have a handle to a file, it's easy to use the functions in `fmt` for formatted reading.  For instance we have a space-delimited file like the following:

We can use the `fmt.Fscanf()` to read into memory as shown in the following example:
```go
func main() {
	var name, hasRing string
	var diam, moons int

	// open data file
	file, err := os.Open("./planets.txt") // creates *os.File
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	// print formatted header
	fmt.Printf("%-10s %-10s %-6s %-6s\n", 
	    "Planet", "Diameter", "Moons", "Ring?",
	)
	for {
		// scan value into variables
		_, err := fmt.Fscanf(file, "%s %d %d %s\n", 
		    &name, &diam, &moons, &hasRing,
		)
		
		// check for io.EOF (end of file)
		// which is returned as an error value
		if err != nil {
			if err == io.EOF {
				break
			} else {
				fmt.Println(err)
				os.Exit(1)
			}
		}
		fmt.Printf("%-10s %-10d %-6d %-6s\n", name, diam, moons, hasRing)
	}
}
```
### Using buffered IO for reading
Go also supports buffered IO via the  `bytes` package. The package includes functions and types that make it easy to work with text content.  For instance, we can redo the previous program to read the content of `planets.txt` a line at a time delimited by byte value  `'\n'` :
```go
func main() {
	file, err := os.Open("./planets.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()
	reader := bufio.NewReader(file) 

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			} else {
				fmt.Println(err)
				os.Exit(1)
			}
		}
		fmt.Print(line)
	}
}
```

### Writing data
We have been using package `fmt`  throughout this document to print textual data to standard output using functions `fmt.Printf()` and `fmt.Println()`.  We can also print textual output to any arbitrary `io.Writer` using functions `fmt.Fprint` or `fmt.Fprintf()`.  For instance, the following snippet uses `fmt.Fprintf()` to write data to an in-memory buffer of type bytes.Buffer which implements `io.Writer`:
```go
package main

import (
	"fmt"
	"bytes"
)

func main() {
	proverbs := []string{
		"Channels orchestrate mutexes serialize",
		"Cgo is not Go",
		"Errors are values",
		"Don't panic",
	}
	
    // bytes.Buffer an in-mem io.Writer
    buf := new(bytes.Buffer)
    
    for i, p := range proverbs {
        fmt.Fprintf(buf, "Proverb %d = %s\n", i, p)
    }
    fmt.Println(buf.String())
}
```
We can modify the previous program to output the formatted text to the standard output instead of an in-memory structure using `os.Stdout`, an `io.Writer`:
```go
package main

import (
	"fmt"
	"os"
)

func main() {
	proverbs := []string{
		"Channels orchestrate mutexes serialize",
		"Cgo is not Go",
		"Errors are values",
		"Don't panic",
	}
	
    for i, p := range proverbs {
        fmt.Fprintf(os.Stdout, "Proverb %d = %s\n", i, p)
    }
}
```
The previous program can be updated further to write to a `*os.File` object which also implements `io.Writer`:

```go
func main() {
	proverbs := []string{
		"Channels orchestrate mutexes serialize",
		"Cgo is not Go",
		"Errors are values",
		"Don't panic",
	}

	file, err := os.Create("proverbs.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	for i, p := range proverbs {
		fmt.Fprintf(file, "Proverb %d = %s\n", i, p)
	}
	fmt.Println("proverbs.txt created")
}
```
## Data communications
Go comes with a formidable set of functionalities for network communication using the `net` package.  We will only scratch the surface with a couple of examples to give you an idea what is possible with a few lines of code.

### A simple socket server 
To get started, let us create a simple echo server and client.  The server is implemented with a few lines of code, but is capable of scaling to handle large number of connections using a goroutine as shown below:
```go
package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	listener, err := net.Listen("tcp", ":4040")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer listener.Close()
	fmt.Println("echo server running on 4040...")

	// connection loop
	for {
		// wait and accept new client connection
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}

		// handle connection
		// read text and echo it
		go func(c net.Conn) {
			defer conn.Close()
			reader := bufio.NewReader(c)
			for {
				line, err := reader.ReadString('\n')
				if err != nil {
					fmt.Println(err)
					return
				}
				if _, err := fmt.Fprint(conn, line); err != nil {
					fmt.Println(err)
					return
				}
			}
		}(conn)
	}
}
```
When the echo server program is started, a telnet session can be used to test it as shown below:
```sh
> telnet localhost 4040
Trying ::1...
Connected to localhost.
Escape character is '^]'.
hello
hello
I like your hat
I like your hat
```
###A simple socket client
Writing a socket client in Go is just as easy.  The following is a simple client to the echo server above.  It captures the keyboard input using buffered IO reader `cmdReader`, then sends the data to the server, and lastly, prints the response from the server captured with `conReader`.
```go
package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

var host, port = "127.0.0.1", "4040"
var addr = net.JoinHostPort(host, port)

func main() {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer conn.Close()
	fmt.Println("Connected to echo server at:", addr)

	for {
		cmdReader := bufio.NewReader(os.Stdin)
		conReader := bufio.NewReader(conn)

		// read command-line input
		data, err := cmdReader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			continue
		}

		// send data to server
		if _, err := fmt.Fprint(conn, data); err != nil {
			fmt.Println(err)
			continue
		}

		// read server response
		resp, err := conReader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Print(resp)
	}
}
```
### A simple web server
Another well-appointed capability of Go is its support for creating HTTP-based servers and clients using the `http` package.  The following shows a program that implements a time API server which returns the current date/time of the server.  
```go
package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, r *http.Request) {
		fmt.Fprint(writer, "/now\n/date\n/time")
	})

	http.HandleFunc("/now", func(writer http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(writer, "%s", time.Now())
	})

	http.HandleFunc("/date", func(writer http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(writer, "%s", time.Now().Format("2006-01-02"))
	})

	http.HandleFunc("/time", func(writer http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(writer, "%s", time.Now().Format("03:04:05 MST"))
	})

	fmt.Println("time api server running on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println(err)
	}
}
```
When we run the program, it immediately blocks to listen and serve incoming HTTP requests on port `8080`.  Using `cURL` we can test the program as shown below:
```sh
$ curl localhost:8080
/now
/date
/time

$ curl localhost:8080/now
2017-08-02 08:32:14.854542532 -0400 EDT

$ curl localhost:8080/time
08:33:34 ED

$ curl localhost:8080/date
2017-08-02
```
The previous program uses a default package-provided server object with default configurations.  The URL routes are handled by a default packaged-provider handler using function `http.HandleFunc` which takes a path and the function used to handle HTTP request.  

However, the program can be re-written using an explicit declarations of an `http.Server` value along with an `http.Handler` object used for route multiplexing called `http.ServeMux`.  This explicit declaration of these objects provides deeper control and configuration of the server as shown below:
```go
func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(writer http.ResponseWriter, r *http.Request) {
		fmt.Fprint(writer, "/now\n/date\n/time")
	})

	mux.HandleFunc("/now", func(writer http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(writer, "%s", time.Now())
	})

	mux.HandleFunc("/date", func(writer http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(writer, "%s", time.Now().Format("2006-01-02"))
	})

	mux.HandleFunc("/time", func(writer http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(writer, "%s", time.Now().Format("03:04:05 MST"))
	})

	server := http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  time.Second * 5,
		WriteTimeout: time.Second * 3,
	}

	fmt.Println("time api server running on :8080")
	if err := server.ListenAndServe(); err != nil {
		fmt.Println(err)
	}
}
```
### Writing HTTP clients
The Go `http` package comes with everything you need to write robust HTTP client programs.  The following shows a simple client, using default configurations, to the time api server above. 
```go
func main() {
	resp, err := http.Get("http://127.0.0.1:8080/now")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	if _, err := io.Copy(os.Stdout, resp.Body); err != nil {
		fmt.Println(err)
	}
}
```
The previous code uses a default client object provided by the `http` package.  If you want want more control over client configuration, you can explicitly declare your own client object as is shown below:
```go
func main() {
	client := &http.Client{
		Timeout: 21 * time.Second,
	}
	resp, err := client.Get("http://127.0.0.1:8080/date")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer resp.Body.Close()
	if _, err := io.Copy(os.Stdout, resp.Body); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
```
## Testing
Code testing is a crucial feature of the Go ecosystem. Support for code testing is baked in the tooling and the standard library.  Test functions are regular functions that have a specific signature so that the `go test` tool can compile them.  Test functions are saved in files with names ending in `*_test.go` so that test sources are not included in the compiled binary for distribution.

Let us consider the following source code for a library which contains simple functions for sum, max, and average, saved in a file called `math.go`:
```go
package main

func sum(nums ...float64) float64 {
	var sum float64
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	return sum
}

func max(nums ...float64) float64 {
	var max float64
	for i := 0; i < len(nums); i++ {
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}

func avg(nums ...float64) float64 {
	return sum(nums...) / float64(len(nums))
}
```
To test these functions, we simply need to do the followings:

- create a Go source file with a name that ends in `*_test.go`
- Next, in that test source file, create functions with name prefixes `TestXXX(t *testing.T)`

The following source code, saved in file `math_test.go` shows the test function used to exercise the logic in the previous code:
```go
package main

import "testing"

func TestSum(t *testing.T) {
	expected := 16
	actual := sum(5, 5, 3, 3)
	if actual != float64(expected) {
		t.Errorf("expecting 16, go %f", actual)
	}
}

func TestMax(t *testing.T) {
	expected := float64(1.2)
	actual := max(1.2, 0.3, 1.02, 0.20, 0.175)
	if actual != expected {
		t.Errorf("expecting 1.2, got %f", actual)
	}
}

func TestAvg(t *testing.T) {
	cases := []struct {
		nums []float64
		avg  float64
	}{
		{[]float64{5, 5, 3, 3}, 4},
		{[]float64{3, 9, 3}, 5},
		{[]float64{3.5, 1.5, 3.2, 1.8}, 2.5},
	}

	for _, c := range cases {
		actual := avg(c.nums...)
		if c.avg != actual {
			t.Errorf("unexpected result: need %f got %f", c.avg, actual)
		}
	}
}
```
The test functions are setup to compare results from the actual functions with known expected values.  Notice that test function `TestAvg()` uses an approach called `table-driven` test where multiple possibilities are tested in one function.

The test framework builtin Go supports many more features including: 
- HTTP testing
- Code coverage
- Test benchmark
- Debugging / Profiling

## Standard library
We only scratched the surface with this introduction to Go and its built-in packages.  The following lists additional packages that are part of the standard library:

- Regular expressions with search and replace
- Full support for secure socket programming with TCP/UDP, IPv4, and IPv6
- APIs for writing production-ready HTTP/HTTPS services and clients
- Synchronization primitives (mutex, atomic, etc.)
- General-purpose template framework with HTML support
- Data serialization (JSON, XML, binary, etc)
- Support for multiple RPC wire formats
- Compression algorithms support: tar, zip/gzip, zlib, etc.
- Cryptography support for most major algorithms and hash functions
- Access to OS-level processes, environment info, and signaling
- And much more

## Conclusion
This document is an attempt to cover helpful topics that a Go newcomer would need to know when getting started.  After going through the material here, the reader should be able to understand the fundamentals of the Go programming language and other critical idioms necessary to start creating Go programs.  However, this is by any stretch of the imagination complete.  After going through the material here, it is highly recommend that you grab a good book ([hopefully mine](https://www.packtpub.com/application-development/learning-go-programming)) to continue your journey into the Go programming language.

Hope you enjoyed it!

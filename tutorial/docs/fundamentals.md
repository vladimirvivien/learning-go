# Language Fundamentals
This section covers the fundamentals of the Go language including data types, variable declaration, and other language construct that are crucial in understanding the language.

## Variables
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
var name = "Mars"	// inferred as type string
var desc = "Planet"	// inferred as type string
var radius = 3396.2	// inferred as type float64
var mass = 6.4185e23	// inferred as type float64
var active = true	// inferred as type bool
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

### The blank identifier `_`	
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
 
## Primitive data types
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
var color uint32 = 0xFEFEFE	// hex, uint32
var mod = 0466			// octal, int
count := 1245			// decimal, int
avogadro := 6.0221409e+1	// float64
value := "automobile"		// string
```
## Composite types
Composite types are used to store sequences of values of primitive types.  Composite literals values are contained within curly-braces preceded by the type as shown below.

### *Array*
Type *array* represents a fixed-size sequenced values numerically indexed.  
```go
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

### *Slice*
A slice is a dynamically-sized array.  The slice omits its size as part of its type declaration as shown in the snippet below.   Slices can be initialized with a composite literal or with the `make()` built-in function:
```go
steps := []string{"SEND", "RCVD", "WAIT"}	// slice initialized with 3 elements
fmt.Println(steps[1]) 				// prints "RCVD" 
steps = append(steps, "PAUSE")			// slice expanded to size 4
steps[3] = "RESUME"				// updates value at index 4

actions := make([]int, 2)			// initializes a slice of size 2
actions[0] = "PRINT"
actions[1] = "LOG"
actions = append(actions, "ADD")		// expands slice, size now 3
```
Go also supports slice expressions which can be used to create new slices from arrays or other slices.  For instance, slice `summer` is created by slicing existing array `months`
```go
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

### *Map*
A map is a dynamically-sized composite type that stores elements of arbitrary types that are indexes using a values of  type.  A map can be initialized using a composite literal:
```go
ratings := map[string][]int{
	"men":   {32, 55, 12, 55, 42, 53},
	"women": {44, 42, 23, 41, 65, 44},
}

ratings["children"] = []int{2,34,5,43,64,22}
```
A map can also be initialized using the built-in function `make()` as shown below:
```go
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
### *Struct*
The *struct* type is a composite that stores named elements of diverse types known as fields.  The following example creates variable `truck` as type `struct{year int; make, model string}` and initializes it with a composite literal.
```go
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

## The pointer type
Go supports a type pointer which is a value that may be used to reference the memory address where the data is located. Go uses the `*` operator to designate a type as a pointer of that type.  The followings are examples of declaration of pointer type where `scorePtr` is a pointer to type `float32`:
```go
var scorePtr *float32
```
Pointer variables can only be assigned address values of its declared type.  One way to do so in Go is to use the address operator `&` (ampersand) to obtain the address of a variable as shown in the following example:
```go
func main() {
	score := 32
	scorePtr := &score    // pointer assigned address
	fmt.Println(scorePtr) // prints address
	*scorePtr = 44        // pointer dereferenced with value
	fmt.Println(score)    // prints updated score
}
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

## The function type
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
## Methods 
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

## Deferring functions and method calls
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
## Type declaration
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
## Interfaces 
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

### Next
The next section continues with Go [flow control](./flow_control.md) statements

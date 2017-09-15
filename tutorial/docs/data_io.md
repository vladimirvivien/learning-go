## Data IO
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
### Next
Continue to the next section to learn about Go [network programming](./networking.md).
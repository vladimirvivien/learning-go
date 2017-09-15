## Concurrency
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

### Next
The next section shows how to use the [Go streaming IO API](./data_io.md).
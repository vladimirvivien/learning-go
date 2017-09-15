# Flow control
Go supports the expected flow control from a modern language for branching and looping as outlined in the his section.
## if statement
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

## switch statement
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
## for statement
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

### Next
Next, let us go beyond fundamentals and learn about [Go concurrency](./concurrency.md).
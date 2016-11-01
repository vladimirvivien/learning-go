package main

import "fmt"
import "math"
import "errors"

// The followings are a collection of functions
// that implement simple operations for illustration purposes.
// This contrive set of examples is designed to illustrate the
// different form and usage of Go functions.

var (
	mul = func(op0, op1 int) int {
		return op0 * op1
	}

	sqr = func(val int) int { return mul(val, val) }
)

// printPi prints Go's internal representation of Pi.
func printPi() {
	fmt.Printf("printPi() %v\n", math.Pi)
}

func avogadro() float64 {
	return 6.02214129e23
}

// fib is a simple interative implementation of a function
// that lists all fibonacci numbers up to n position in the sequence
func fib(n int) {
	fmt.Printf("fib(%d): [", n)

	var p0, p1 uint64 = 0, 1
	fmt.Printf("%d %d ", p0, p1)
	for i := 2; i <= n; i++ {
		p0, p1 = p1, p0+p1
		fmt.Printf("%d ", p1)
	}

	fmt.Println("]")
}

// isPrime tests number n for primality using the "trial division" method
// to test all factors up to Sqrt(n)
func isPrime(n int) bool {
	lim := int(math.Sqrt(float64(n)))
	for p := 2; p <= lim; p++ {
		if (n % p) == 0 {
			return false
		}

	}
	return true
}

// add implements the addition operation returning the value
// op0 + op1
func add(op0 int, op1 int) int {
	return op0 + op1
}

// sub returns the value of op0 - op1
func sub(op0, op1 int) int {
	return op0 - op1
}

// div is a simplistic implementtion of a Eucledian
// division algorithm that returns both quotient and remainer.
// See example divide.go
func div(op0, op1 int) (int, int) {
	r := op0
	q := 0
	for r >= op1 {
		q++
		r = r - op1
	}
	return q, r
}

// div2 is the same as above.  However, the returned values
// are setup with named identifiers.
func div2(dvdn, dvsr int) (q, r int) {
	r = dvdn
	for r >= dvsr {
		q++
		r = r - dvsr
	}
	return
}

// avg demonstrates variadic parameters.
// It returns the average of all values passed to function.
func avg(nums ...float64) float64 {
	n := len(nums)
	t := 0.0
	for _, v := range nums {
		t += v
	}
	return t / float64(n)
}

// sum demonstrates variadic parameters
// it returns the sum of all numers passed to function
func sum(nums ...float64) float64 {
	var sum float64
	for _, v := range nums {
		sum += v
	}
	return sum
}

// doUnary uses first-order functionality of Go
// to implement a function that can apply unary
// operation to any given integer (i.e. sqr(3), log(1000), etc)
func doUnary(f func(int) int, val int) int {
	return f(val)
}

// doBinary illustrates first-order function capability in Go.
// It invokes invokes any function that takes two ints and return
// an integer as a result (i.e. 2 + 3, or 9 * 5, etc)
func doBinary(f func(int, int) int, op1, op2 int) int {
	return f(op1, op2)
}

// apply uses a functional idiom to apply the
// provided function on each numer in nums
func apply(nums []int, f func(int) int) {
	for i, v := range nums {
		nums[i] = f(v)
	}
}

// filter uses Go's first-order function support to
// build a functional idiom where the provided number list
// by applying f()
func filter(nums []int, f func(int) bool) []int {
	var result []int
	for _, v := range nums {
		if f(v) {
			result = append(result, v)
		}
	}
	return result
}

func main() {
	// simple func calls
	printPi()
	fmt.Printf("avogadro() = %e 1/mol\n", avogadro())
	fib(41)
	prime := 37
	fmt.Printf("isPrime(%d) returns %v\n", prime, isPrime(prime))

	// func type, assignment
	var addition func(int, int) int // a function type
	addition = add
	fmt.Println("func add(op0,op1 int) int {return op0+op1}")
	fmt.Println("addition := add")
	fmt.Printf("addition(3,7) = %d\n", addition(3, 7))
	fmt.Printf("doBinary(f func(int,int)int, int,int) -> doBinary(addition,12,4) = %d\n", doBinary(addition, 12, 4))
	fmt.Println("mul := func(i,j int) int { return i * j}")
	fmt.Printf("mul(25,7) = %d\n", mul(25, 7))
	fmt.Println("sqr := func (i int) int {return mul(i,i)")
	fmt.Printf("sql(13) = %d\n", sqr(13))
	fmt.Printf("doUnary(f func(int)int, int) -> doUnary(sqr, 12) = %d\n", doUnary(sqr, 12))

	// multi-result functions
	q, r := div(71, 5)
	fmt.Printf("div(71,5) -> q = %d, r = %d\n", q, r)

	// anonymous function def/call
	fmt.Printf(
		"94 (°F) = %.2f (°C)\n",
		func(f float64) float64 {
			return (f - 32.0) * (5.0 / 9.0)
		}(94),
	)

	// higher-order idioms using
	nums := []int{4, 32, 11, 77, 556, 3, 19, 88, 422, 21, 52, 97, 123}
	fmt.Println("nums := ", nums)

	fmt.Println("Squqre nums: apply(nums,func(int i)int{return i * i})")
	apply(nums, func(i int) int {
		return i * i
	})
	fmt.Println(nums)

	fmt.Println("Half nums: apply(nums,func(int i)int{return i/2})")
	apply(nums, func(i int) int {
		return i / 2
	})
	fmt.Println(nums)

	fmt.Println("filter(nums, func(i int)bool{return (i %% 2) == 0})")
	fmt.Printf("%v\n", filter(nums, func(i int) bool { return (i % 2) == 0 }))

	// passing in variadic parameter values directly
	fmt.Println("avg([1, 2.5, 3.75, 4.1, 5.9, 9.2, 8.7]) =",
		avg(1, 2.5, 3.75, 4.1, 5.9, 9.2, 8.7))

	// passing variadic param values as a variable
	points := []float64{9, 4, 3.7, 7.1, 7.9, 9.2, 10}
	fmt.Printf("sum(%v) = %f\n", points, sum(points...))

	var deg float64 = 290.2
	if rad := func() float64 { return deg * math.Pi / 180 }(); rad > 3 {
		fmt.Println("Calculated rad", rad, "too large.")
	} else {
		fmt.Println("Calculated radian:", rad)
	}
}

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

## Othere features
The test framework that comes with Go supports many more features including: 
- HTTP testing
- Code coverage
- Test benchmark
- Debugging / Profiling

### Next
See what other APIs are available in Go's [standard library](./stdlib.md).
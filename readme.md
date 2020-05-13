This package is distributed in the hope that it will be useful. For that it has a simple goal: to fast boilerplate the local development of Go local solutions for Hackerrank tests.

First, create a new object of type Tests with a directory that contains inputs and outputs of Hackerrank tests.
```
var tests Tests = hackerrank.NewTests("inputs_outputs_files_path")
```
This Tests object is created traversing the directory looking for tests, based on the input and output files.
It considers one test each filename on the format `"input" + "##" + ".txt"` and look for output equivalent file.
This means "input00.txt" is considered, but "Xput00.txt" is not considered.

For each test, it will have properties of a name (made of the test file number) and two file readers. These file readers have methods that are actually used.

```
Tests.Test.Name
Tests.Test.In.NextLine()
Tests.Test.Out.NextLine()
```

With these tests, call Run with a function with signature func f(hackerrank.Test).
With each test, read input lines, treat accordingly, and call a function with signature equal to the hackerrank function with the final solver.

### Example using this package
```go
package main

import (
    "fmt"
    "log"
    "strconv"
    "github.com/sztelzer/hackerrank"
)

func main() {
	tests := hackerrank.NewTests("../tests")
	tests.Run(Solution)
}

func Solution(t hackerrank.Test) {
    s1 := t.In.NextLine()
    s2 := t.In.NextLine()
    // treat these strings as needed by the solver

    n1, _ := strconv.atoi(s1)

    result := Solver(n1, s2)
    should := t.Out.NextLine()

    fmt.Println(should, result, should==result)
}

// build this function as exactly as hackerrank expects
func Solver(n int, s string) string {
	// return your solution
}
```
Now you can copy solver body to hackerrank.

# go-assert

`go-assert` is a simple Go package that provides assertion functions to enforce conditions in your code. If a condition is not met, the package will trigger a panic, making it clear that a critical assumption has been violated. This can be particularly useful for defensive programming, and ensuring code correctness.

## Why Use Assertions?

Assertions provide several benefits when incorporated into your code:
- Early Bug Detection: Assertions help catch logical errors and unexpected conditions early in the development process. By asserting the expected state of your program, you can detect issues immediately during testing or runtime rather than allowing them to propagate and potentially cause more serious issues later.
- Documentation of Assumptions: Assert statements serve as documentation of your code's assumptions and expectations. They make it clear to other developers (including your future self) what conditions must hold true for the code to function correctly.
- Debugging Aid: When an assertion fails, it provides valuable information about what went wrong, including the expected versus actual values and a descriptive message. This helps pinpoint the root cause of the issue quickly.
- Code Maintainability: Assertions can improve code maintainability by enforcing invariants and reducing the likelihood of subtle bugs creeping into the codebase over time. They serve as guardrails that ensure the integrity of your program's internal state.

## Features

- Assert that a condition is true
- Assert that two values are equal or not equal
- Assert that a value is nil or not nil
- Assert that a boolean value is true or false

## Installation

To install the package, use `go get`:

```sh
go get github.com/MikolajKorbanek/go-assert
```

Then, import the package in your Go code:

```go
import "github.com/MikolajKorbanek/go-assert"
```

## Usage

Here are some examples of how to use the `go-assert` package in your code:

### Assert

```go
package main

import "github.com/MikolajKorbanek/go-assert"

func main() {
    x := 5
    assert.Assert(x > 0, "x should be greater than 0")
}
```

### AssertEqual

```go
package main

import "github.com/MikolajKorbanek/go-assert"

func main() {
    y := "hello"
    assert.AssertEqual("hello", y, "y should be 'hello'")
}
```

### AssertNotEqual

```go
package main

import "github.com/MikolajKorbanek/go-assert"

func main() {
    z := 10
    assert.AssertNotEqual(5, z, "z should not be 5")
}
```

### AssertNil

```go
package main

import "github.com/MikolajKorbanek/go-assert"

func main() {
    var ptr *int
    assert.AssertNil(ptr, "ptr should be nil")
}
```

### AssertNotNil

```go
package main

import "github.com/MikolajKorbanek/go-assert"

func main() {
    ptr := new(int)
    assert.AssertNotNil(ptr, "ptr should not be nil")
}
```

### AssertFalse

```go
package main

import "github.com/MikolajKorbanek/go-assert"

func main() {
    flag := false
    assert.AssertFalse(flag, "flag should be false")
}
```

## Contributing

Contributions are welcome! Please feel free to submit a pull request or open an issue on GitHub.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file

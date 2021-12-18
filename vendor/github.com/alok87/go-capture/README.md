# go-capturer

Capture `os.Stdout` and/or `os.Stderr` in Go.
This package is useful for writing tests which print some outputs using `fmt` package.

## Usage

This package provides `Stdout()`, `Stderr()` and `Output()` functions to capture outputs.

```go
package main

import (
	"fmt"
	"os"

	"github.com/alok87/go-capture"
)

func ExampleCaptureStdout() {
	out := capture.Stdout(func() {
		fmt.Fprint(os.Stdout, "foo")
	})

	fmt.Println(out)
	// Output: foo
}

func ExampleCaptureStderr() {
	out := capture.Stderr(func() {
		fmt.Fprint(os.Stderr, "bar")
	})

	fmt.Println(out)
	// Output: bar
}

func ExampleCaptureOutput() {
	out := capture.Output(func() {
		fmt.Fprint(os.Stdout, "foo")
		fmt.Fprint(os.Stderr, "bar")
	})

	fmt.Println(out)
	// Output: foobar
}
```

## Installation

```
$ go get github.com/alok87/go-capture
```

## Contributing

Bug reports and pull requests are welcome on GitHub at https://github.com/alok87/go-capture.

```
go test ./...
```


## License

The package is available as open source under the terms of the [MIT License](http://opensource.org/licenses/MIT).

# walk

[![CI](https://github.com/MJKWoolnough/python/actions/workflows/go-checks.yml/badge.svg)](https://github.com/MJKWoolnough/python/actions)
[![Go Reference](https://pkg.go.dev/badge/vimagination.zapto.org/python.svg)](https://pkg.go.dev/vimagination.zapto.org/python/walk)

--
    import "vimagination.zapto.org/python/walk"

Package walk provides a python type walker.

## Highlights

 - Simple interface to allow control over walking through parsed python.
 - Allows modification to the tree as it's being walked.

## Usage

```go
package main

import (
	"fmt"

	"vimagination.zapto.org/parser"
	"vimagination.zapto.org/python"
	"vimagination.zapto.org/python/walk"
)

func main() {
	src := `a = 'b' - "c"`
	tk := parser.NewStringTokeniser(src)

	p, _ := python.Parse(&tk)

	var walkFn walk.Handler

	walkFn = walk.HandlerFunc(func(t python.Type) error {
		switch t := t.(type) {
		case *python.AddExpression:
			if t.Add != nil {
				t.Add.Data = "+"
			}
		case *python.Atom:
			if t.Literal != nil {
				switch t.Literal.Data {
				case "'b'":
					t.Literal.Data = "'Hello'"
				case `"c"`:
					t.Literal.Data = `", world"`
				}
			}
		}

		return walk.Walk(t, walkFn)
	})

	walk.Walk(p, walkFn)

	fmt.Printf("%+s", p)

	// Output:
	// a = 'Hello' + ", world"
}
```

## Documentation

Full API docs can be found at:

https://pkg.go.dev/vimagination.zapto.org/python/walk

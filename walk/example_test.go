package walk_test

import (
	"fmt"

	"vimagination.zapto.org/parser"
	"vimagination.zapto.org/python"
	"vimagination.zapto.org/python/walk"
)

func Example() {
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

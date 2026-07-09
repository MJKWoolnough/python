# python

[![CI](https://github.com/MJKWoolnough/python/actions/workflows/go-checks.yml/badge.svg)](https://github.com/MJKWoolnough/python/actions)
[![Go Reference](https://pkg.go.dev/badge/vimagination.zapto.org/python.svg)](https://pkg.go.dev/vimagination.zapto.org/python)

--
    import "vimagination.zapto.org/python"

Package python implements a python tokeniser and parser.

## Highlights

 - Parse python code into AST.
 - Modify parsed code.
 - Consistent python formatting.

## Usage

```go
package main

import (
	"fmt"

	"vimagination.zapto.org/parser"
	"vimagination.zapto.org/python"
)

func main() {
	src := `for name in ["Alice", "Bob", "Charlie"]: print("Hello,", name)`

	tk := parser.NewStringTokeniser(src)

	ast, err := python.Parse(&tk)
	if err != nil {
		fmt.Println(err)

		return
	}

	python.UnwrapConditional(python.UnwrapConditional(ast.Statements[0].CompoundStatement.For.Suite.StatementList.Statements[0].AssignmentStatement.StarredExpression.Expression.ConditionalExpression).(*python.PrimaryExpression).Call.ArgumentList.PositionalArguments[0].AssignmentExpression.Expression.ConditionalExpression).(*python.Atom).Literal.Data = `"Hi,"`

	fmt.Printf("%+s", ast)

	// Output:
	// for name in ["Alice", "Bob", "Charlie"]: print("Hi,", name)
}
```

## Documentation

Full API docs can be found at:

https://pkg.go.dev/vimagination.zapto.org/python

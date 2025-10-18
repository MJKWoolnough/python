package python_test

import (
	"fmt"

	"vimagination.zapto.org/parser"
	"vimagination.zapto.org/python"
)

func Example() {
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

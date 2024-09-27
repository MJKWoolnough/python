package python

import (
	"errors"
	"reflect"
	"testing"

	"vimagination.zapto.org/parser"
)

type sourceFn struct {
	Source string
	Fn     func(*test, Tokens)
}

type test struct {
	Tokens               *pyParser
	Output               Type
	AssignmentExpression *AssignmentExpression
	TokenSkip            int
	Err                  error
}

func doTests(t *testing.T, tests []sourceFn, fn func(*test) (Type, error)) {
	t.Helper()

	var err error

	for n, tt := range tests {
		var ts test

		tk := parser.NewStringTokeniser(tt.Source)

		ts.Tokens, err = newPyParser(&tk)
		if err != nil {
			t.Errorf("test %d: unexpected error: %s", n+1, err)

			continue
		}

		tt.Fn(&ts, ts.Tokens.Tokens[:cap(ts.Tokens.Tokens)])

		if output, err := fn(&ts); !errors.Is(err, ts.Err) {
			t.Errorf("test %d: expecting error: %v, got %v", n+1, ts.Err, err)
		} else if ts.Output != nil && !reflect.DeepEqual(output, ts.Output) {
			t.Errorf("test %d: expecting \n%+v\n...got...\n%+v", n+1, ts.Output, output)
		}
	}
}

func wrapConditionalExpressionError(err Error) error {
	switch err.Parsing {
	case "Atom":
		err = Error{
			Err:     err,
			Parsing: "PrimaryExpression",
			Token:   err.Token,
		}

		fallthrough
	case "PrimaryExpression":
		err = Error{
			Err:     err,
			Parsing: "PowerExpression",
			Token:   err.Token,
		}

		fallthrough
	case "PowerExpression":
		err = Error{
			Err:     err,
			Parsing: "UnaryExpression",
			Token:   err.Token,
		}

		fallthrough
	case "UnaryExpression":
		err = Error{
			Err:     err,
			Parsing: "MultiplyExpression",
			Token:   err.Token,
		}

		fallthrough
	case "MultiplyExpression":
		err = Error{
			Err:     err,
			Parsing: "AddExpression",
			Token:   err.Token,
		}

		fallthrough
	case "AddExpression":
		err = Error{
			Err:     err,
			Parsing: "ShiftExpression",
			Token:   err.Token,
		}

		fallthrough
	case "ShiftExpression":
		err = Error{
			Err:     err,
			Parsing: "AndExpression",
			Token:   err.Token,
		}

		fallthrough
	case "AndExpression":
		err = Error{
			Err:     err,
			Parsing: "XorExpression",
			Token:   err.Token,
		}

		fallthrough
	case "XorExpression":
		err = Error{
			Err:     err,
			Parsing: "OrExpression",
			Token:   err.Token,
		}

		fallthrough
	case "OrExpression":
		err = Error{
			Err:     err,
			Parsing: "Comparison",
			Token:   err.Token,
		}

		fallthrough
	case "Comparison":
		err = Error{
			Err:     err,
			Parsing: "NotTest",
			Token:   err.Token,
		}

		fallthrough
	case "NotTest":
		err = Error{
			Err:     err,
			Parsing: "AndTest",
			Token:   err.Token,
		}

		fallthrough
	case "AndTest":
		err = Error{
			Err:     err,
			Parsing: "OrTest",
			Token:   err.Token,
		}

		fallthrough
	case "OrTest":
		err = Error{
			Err:     err,
			Parsing: "ConditionalExpression",
			Token:   err.Token,
		}

	}

	return err
}

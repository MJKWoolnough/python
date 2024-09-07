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
	Tokens *pyParser
	Output Type
	Err    error
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

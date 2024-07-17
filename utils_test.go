package python

import (
	"errors"
	"strconv"
	"testing"
)

func TestUnquote(t *testing.T) {
	for n, test := range [...]struct {
		Input, Output string
		Err           error
	}{
		{
			Input:  "\"abc\"",
			Output: "abc",
		},
		{
			Input:  "\"ab\\\"c\"",
			Output: "ab\"c",
		},
		{
			Input:  "'ab\\\"c'",
			Output: "ab\"c",
		},
		{
			Input:  "'ab\\'c'",
			Output: "ab'c",
		},
		{
			Input: "\"ab\nc\"",
			Err:   strconv.ErrSyntax,
		},
		{
			Input:  "\"\"\"ab\nc\"\"\"",
			Output: "ab\nc",
		},
		{
			Input: "'ab\nc'",
			Err:   strconv.ErrSyntax,
		},
		{
			Input:  "'''ab\nc'''",
			Output: "ab\nc",
		},
		{
			Input: "\"abc\\\"",
			Err:   strconv.ErrSyntax,
		},
		{
			Input:  "r\"abc\\\"",
			Output: "abc\\",
		},
		{
			Input:  "R'abc\\'",
			Output: "abc\\",
		},
	} {
		output, err := Unquote(test.Input)

		if !errors.Is(test.Err, err) {
			t.Errorf("test %d: expecting error %q, got %q", n+1, test.Err, err)
		} else if test.Output != output {
			t.Errorf("test %d: expecting output %q, got %q", n+1, test.Output, output)
		}
	}
}

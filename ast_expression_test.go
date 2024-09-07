package python

import (
	"testing"
)

func TestAtom(t *testing.T) {
	doTests(t, []sourceFn{
		{`a`, func(t *test, tk Tokens) { // 1
			t.Output = Atom{
				Identifier: &tk[0],
				Tokens:     tk[:1],
			}
		}},
		{`1`, func(t *test, tk Tokens) { // 2
			t.Output = Atom{
				Literal: &tk[0],
				Tokens:  tk[:1],
			}
		}},
		{`"abc"`, func(t *test, tk Tokens) { // 3
			t.Output = Atom{
				Literal: &tk[0],
				Tokens:  tk[:1],
			}
		}},
	}, func(t *test) (Type, error) {
		var a Atom

		err := a.parse(t.Tokens)

		return a, err
	})
}

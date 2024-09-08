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
		{`{}`, func(t *test, tk Tokens) { // 4
			t.Output = Atom{
				Enclosure: &Enclosure{
					DictDisplay: &DictDisplay{},
					Tokens:      tk[:2],
				},
				Tokens: tk[:2],
			}
		}},
	}, func(t *test) (Type, error) {
		var a Atom

		err := a.parse(t.Tokens)

		return a, err
	})
}

func TestPrimaryExpression(t *testing.T) {
	doTests(t, []sourceFn{
		{`a`, func(t *test, tk Tokens) { // 1
			t.Output = PrimaryExpression{
				Atom: &Atom{
					Identifier: &tk[0],
					Tokens:     tk[:1],
				},
				Tokens: tk[:1],
			}
		}},
		{`a.b`, func(t *test, tk Tokens) { // 2
			t.Output = PrimaryExpression{
				PrimaryExpression: &PrimaryExpression{
					Atom: &Atom{
						Identifier: &tk[0],
						Tokens:     tk[:1],
					},
					Tokens: tk[:1],
				},
				AttributeRef: &tk[2],
				Tokens:       tk[:3],
			}
		}},
		{`a[b]`, func(t *test, tk Tokens) { // 3
			t.Output = PrimaryExpression{
				PrimaryExpression: &PrimaryExpression{
					Atom: &Atom{
						Identifier: &tk[0],
						Tokens:     tk[:1],
					},
					Tokens: tk[:1],
				},
				Slicing: &SliceList{
					SliceItems: []SliceItem{
						{
							Expression: &Expression{
								ConditionalExpression: WrapConditional(&PrimaryExpression{
									Atom: &Atom{
										Identifier: &tk[2],
										Tokens:     tk[2:3],
									},
									Tokens: tk[2:3],
								}),
								Tokens: tk[2:3],
							},
							Tokens: tk[2:3],
						},
					},
					Tokens: tk[2:3],
				},
				Tokens: tk[:4],
			}
		}},
	}, func(t *test) (Type, error) {
		var pe PrimaryExpression

		err := pe.parse(t.Tokens)

		return pe, err
	})
}

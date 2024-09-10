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
		{`True`, func(t *test, tk Tokens) { // 5
			t.Output = Atom{
				Literal: &tk[0],
				Tokens:  tk[:1],
			}
		}},
		{`nonlocal`, func(t *test, tk Tokens) { // 6
			t.Err = Error{
				Err: Error{
					Err:     ErrInvalidEnclosure,
					Parsing: "Enclosure",
					Token:   tk[0],
				},
				Parsing: "Atom",
				Token:   tk[0],
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
		{`a(b)`, func(t *test, tk Tokens) { // 4
			t.Output = PrimaryExpression{
				PrimaryExpression: &PrimaryExpression{
					Atom: &Atom{
						Identifier: &tk[0],
						Tokens:     tk[:1],
					},
					Tokens: tk[:1],
				},
				Call: &ArgumentListOrComprehension{
					ArgumentList: &ArgumentList{
						PositionalArguments: []PositionalArgument{
							{
								AssignmentExpression: &AssignmentExpression{
									Expression: Expression{
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
								Tokens: tk[2:3],
							},
						},
						Tokens: tk[2:3],
					},
					Tokens: tk[2:3],
				},
				Tokens: tk[:4],
			}
		}},
		{`nonlocal`, func(t *test, tk Tokens) { // 5
			t.Err = Error{
				Err: Error{
					Err: Error{
						Err:     ErrInvalidEnclosure,
						Parsing: "Enclosure",
						Token:   tk[0],
					},
					Parsing: "Atom",
					Token:   tk[0],
				},
				Parsing: "PrimaryExpression",
				Token:   tk[0],
			}
		}},
		{`a.nonlocal`, func(t *test, tk Tokens) { // 6
			t.Err = Error{
				Err:     ErrMissingIdentifier,
				Parsing: "PrimaryExpression",
				Token:   tk[2],
			}
		}},
		{`a[nonlocal]`, func(t *test, tk Tokens) { // 7
			t.Err = Error{
				Err: Error{
					Err: Error{
						Err: Error{
							Err: wrapConditionalExpressionError(Error{
								Err: Error{
									Err:     ErrInvalidEnclosure,
									Parsing: "Enclosure",
									Token:   tk[2],
								},
								Parsing: "Atom",
								Token:   tk[2],
							}),
							Parsing: "Expression",
							Token:   tk[2],
						},
						Parsing: "SliceItem",
						Token:   tk[2],
					},
					Parsing: "ExpressionList",
					Token:   tk[2],
				},
				Parsing: "PrimaryExpression",
				Token:   tk[2],
			}
		}},
	}, func(t *test) (Type, error) {
		var pe PrimaryExpression

		err := pe.parse(t.Tokens)

		return pe, err
	})
}

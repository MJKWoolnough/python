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
	Expression           *Expression
	OrigTokens           *pyParser
	Decorators           *Decorators
	TokenSkip            int
	AllowTypeAnnotations bool
	Async                bool
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

		ts.Tokens = ts.Tokens.NewGoal()

		for range ts.TokenSkip {
			ts.Tokens.Next()
		}

		if output, err := fn(&ts); !reflect.DeepEqual(err, ts.Err) {
			t.Errorf("test %d: expecting error: %v, got %v", n+1, ts.Err, err)
		} else if ts.Output != nil && !reflect.DeepEqual(output, ts.Output) {
			t.Errorf("test %d: expecting \n%+v\n...got...\n%+v", n+1, ts.Output, output)
		}
	}
}

func wrapConditionalExpressionError(err Error) Error {
	switch err.Parsing {
	case "Enclosure":
		err = Error{
			Err:     err,
			Parsing: "Atom",
			Token:   err.Token,
		}

		fallthrough
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

func TestParse(t *testing.T) {
	const (
		errorA = "Tokens: error at position 1 (1:1):\ninvalid character"
		errorB = "File: error at position 1 (1:1):\nStatement: error at position 1 (1:1):\nStatementList: error at position 1 (1:1):\nSimpleStatement: error at position 1 (1:1):\nNonLocalStatement: error at position 9 (1:9):\nmissing identifier"
	)

	var e Error

	tk := parser.NewStringTokeniser("!")

	if _, err := Parse(&tk); err == nil {
		t.Error("1: expecting error, got nil")
	} else if !errors.As(err, &e) {
		t.Errorf("1: expecting Error type, got %T", err)
	} else if errStr := e.Error(); errStr != errorA {
		t.Errorf("1: expecting error %q, got %q", errorA, errStr)
	}

	tk = parser.NewStringTokeniser("nonlocal")

	if _, err := Parse(&tk); err == nil {
		t.Error("2: expecting error, got nil")
	} else if !errors.As(err, &e) {
		t.Errorf("2: expecting Error type, got %T", err)
	} else if errStr := e.Error(); errStr != errorB {
		t.Errorf("2: expecting error %q, got %q", errorB, errStr)
	}
}

func TestFile(t *testing.T) {
	doTests(t, []sourceFn{
		{`a`, func(t *test, tk Tokens) { // 1
			t.Output = File{
				Statements: []Statement{
					{
						StatementList: &StatementList{
							Statements: []SimpleStatement{
								{
									Type: StatementAssignment,
									AssignmentStatement: &AssignmentStatement{
										StarredExpression: &StarredExpression{
											Expression: &Expression{
												ConditionalExpression: WrapConditional(&Atom{
													Identifier: &tk[0],
													Tokens:     tk[:1],
												}),
												Tokens: tk[:1],
											},
											Tokens: tk[:1],
										},
										Tokens: tk[:1],
									},
									Tokens: tk[:1],
								},
							},
							Tokens: tk[:1],
						},
						Tokens: tk[:1],
					},
				},
				Tokens: tk[:1],
			}
		}},
		{"a\nb", func(t *test, tk Tokens) { // 2
			t.Output = File{
				Statements: []Statement{
					{
						StatementList: &StatementList{
							Statements: []SimpleStatement{
								{
									Type: StatementAssignment,
									AssignmentStatement: &AssignmentStatement{
										StarredExpression: &StarredExpression{
											Expression: &Expression{
												ConditionalExpression: WrapConditional(&Atom{
													Identifier: &tk[0],
													Tokens:     tk[:1],
												}),
												Tokens: tk[:1],
											},
											Tokens: tk[:1],
										},
										Tokens: tk[:1],
									},
									Tokens: tk[:1],
								},
							},
							Tokens: tk[:1],
						},
						Tokens: tk[:1],
					},
					{
						StatementList: &StatementList{
							Statements: []SimpleStatement{
								{
									Type: StatementAssignment,
									AssignmentStatement: &AssignmentStatement{
										StarredExpression: &StarredExpression{
											Expression: &Expression{
												ConditionalExpression: WrapConditional(&Atom{
													Identifier: &tk[2],
													Tokens:     tk[2:3],
												}),
												Tokens: tk[2:3],
											},
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
				},
				Tokens: tk[:3],
			}
		}},
		{"a\n\nb", func(t *test, tk Tokens) { // 3
			t.Output = File{
				Statements: []Statement{
					{
						StatementList: &StatementList{
							Statements: []SimpleStatement{
								{
									Type: StatementAssignment,
									AssignmentStatement: &AssignmentStatement{
										StarredExpression: &StarredExpression{
											Expression: &Expression{
												ConditionalExpression: WrapConditional(&Atom{
													Identifier: &tk[0],
													Tokens:     tk[:1],
												}),
												Tokens: tk[:1],
											},
											Tokens: tk[:1],
										},
										Tokens: tk[:1],
									},
									Tokens: tk[:1],
								},
							},
							Tokens: tk[:1],
						},
						Tokens: tk[:1],
					},
					{
						StatementList: &StatementList{
							Statements: []SimpleStatement{
								{
									Type: StatementAssignment,
									AssignmentStatement: &AssignmentStatement{
										StarredExpression: &StarredExpression{
											Expression: &Expression{
												ConditionalExpression: WrapConditional(&Atom{
													Identifier: &tk[3],
													Tokens:     tk[3:4],
												}),
												Tokens: tk[3:4],
											},
											Tokens: tk[3:4],
										},
										Tokens: tk[3:4],
									},
									Tokens: tk[3:4],
								},
							},
							Tokens: tk[3:4],
						},
						Tokens: tk[3:4],
					},
				},
				Tokens: tk[:4],
			}
		}},
		{"a \nb", func(t *test, tk Tokens) { // 4
			t.Output = File{
				Statements: []Statement{
					{
						StatementList: &StatementList{
							Statements: []SimpleStatement{
								{
									Type: StatementAssignment,
									AssignmentStatement: &AssignmentStatement{
										StarredExpression: &StarredExpression{
											Expression: &Expression{
												ConditionalExpression: WrapConditional(&Atom{
													Identifier: &tk[0],
													Tokens:     tk[:1],
												}),
												Tokens: tk[:1],
											},
											Tokens: tk[:1],
										},
										Tokens: tk[:1],
									},
									Tokens: tk[:1],
								},
							},
							Tokens: tk[:1],
						},
						Tokens: tk[:1],
					},
					{
						StatementList: &StatementList{
							Statements: []SimpleStatement{
								{
									Type: StatementAssignment,
									AssignmentStatement: &AssignmentStatement{
										StarredExpression: &StarredExpression{
											Expression: &Expression{
												ConditionalExpression: WrapConditional(&Atom{
													Identifier: &tk[3],
													Tokens:     tk[3:4],
												}),
												Tokens: tk[3:4],
											},
											Tokens: tk[3:4],
										},
										Tokens: tk[3:4],
									},
									Tokens: tk[3:4],
								},
							},
							Tokens: tk[3:4],
						},
						Tokens: tk[3:4],
					},
				},
				Tokens: tk[:4],
			}
		}},
		{"a\n#A comment\nb", func(t *test, tk Tokens) { // 5
			t.Output = File{
				Statements: []Statement{
					{
						StatementList: &StatementList{
							Statements: []SimpleStatement{
								{
									Type: StatementAssignment,
									AssignmentStatement: &AssignmentStatement{
										StarredExpression: &StarredExpression{
											Expression: &Expression{
												ConditionalExpression: WrapConditional(&Atom{
													Identifier: &tk[0],
													Tokens:     tk[:1],
												}),
												Tokens: tk[:1],
											},
											Tokens: tk[:1],
										},
										Tokens: tk[:1],
									},
									Tokens: tk[:1],
								},
							},
							Tokens: tk[:1],
						},
						Tokens: tk[:1],
					},
					{
						StatementList: &StatementList{
							Statements: []SimpleStatement{
								{
									Type: StatementAssignment,
									AssignmentStatement: &AssignmentStatement{
										StarredExpression: &StarredExpression{
											Expression: &Expression{
												ConditionalExpression: WrapConditional(&Atom{
													Identifier: &tk[4],
													Tokens:     tk[4:5],
												}),
												Tokens: tk[4:5],
											},
											Tokens: tk[4:5],
										},
										Tokens: tk[4:5],
									},
									Tokens: tk[4:5],
								},
							},
							Tokens: tk[4:5],
						},
						Comments: Comments{tk[2]},
						Tokens:   tk[2:5],
					},
				},
				Tokens: tk[:5],
			}
		}},
		{"a #A comment\nb", func(t *test, tk Tokens) { // 6
			t.Output = File{
				Statements: []Statement{
					{
						StatementList: &StatementList{
							Statements: []SimpleStatement{
								{
									Type: StatementAssignment,
									AssignmentStatement: &AssignmentStatement{
										StarredExpression: &StarredExpression{
											Expression: &Expression{
												ConditionalExpression: WrapConditional(&Atom{
													Identifier: &tk[0],
													Tokens:     tk[:1],
												}),
												Tokens: tk[:1],
											},
											Tokens: tk[:1],
										},
										Tokens: tk[:1],
									},
									Comments: Comments{tk[2]},
									Tokens:   tk[:3],
								},
							},
							Tokens: tk[:3],
						},
						Tokens: tk[:3],
					},
					{
						StatementList: &StatementList{
							Statements: []SimpleStatement{
								{
									Type: StatementAssignment,
									AssignmentStatement: &AssignmentStatement{
										StarredExpression: &StarredExpression{
											Expression: &Expression{
												ConditionalExpression: WrapConditional(&Atom{
													Identifier: &tk[4],
													Tokens:     tk[4:5],
												}),
												Tokens: tk[4:5],
											},
											Tokens: tk[4:5],
										},
										Tokens: tk[4:5],
									},
									Tokens: tk[4:5],
								},
							},
							Tokens: tk[4:5],
						},
						Tokens: tk[4:5],
					},
				},
				Tokens: tk[:5],
			}
		}},
		{`nonlocal`, func(t *test, tk Tokens) { // 7
			t.Err = Error{
				Err: Error{
					Err: Error{
						Err: Error{
							Err: Error{
								Err:     ErrMissingIdentifier,
								Parsing: "NonLocalStatement",
								Token:   tk[1],
							},
							Parsing: "SimpleStatement",
							Token:   tk[0],
						},
						Parsing: "StatementList",
						Token:   tk[0],
					},
					Parsing: "Statement",
					Token:   tk[0],
				},
				Parsing: "File",
				Token:   tk[0],
			}
		}},
		{"a\n# A Comment", func(t *test, tk Tokens) { // 8
			t.Output = File{
				Statements: []Statement{
					{
						StatementList: &StatementList{
							Statements: []SimpleStatement{
								{
									Type: StatementAssignment,
									AssignmentStatement: &AssignmentStatement{
										StarredExpression: &StarredExpression{
											Expression: &Expression{
												ConditionalExpression: WrapConditional(&Atom{
													Identifier: &tk[0],
													Tokens:     tk[:1],
												}),
												Tokens: tk[:1],
											},
											Tokens: tk[:1],
										},
										Tokens: tk[:1],
									},
									Tokens: tk[:1],
								},
							},
							Tokens: tk[:1],
						},
						Tokens: tk[:1],
					},
				},
				Comments: Comments{tk[2]},
				Tokens:   tk[:3],
			}
		}},
		{"a # A Comment\n# B Comment\n\n# EOF Comment", func(t *test, tk Tokens) { // 9
			t.Output = File{
				Statements: []Statement{
					{
						StatementList: &StatementList{
							Statements: []SimpleStatement{
								{
									Type: StatementAssignment,
									AssignmentStatement: &AssignmentStatement{
										StarredExpression: &StarredExpression{
											Expression: &Expression{
												ConditionalExpression: WrapConditional(&Atom{
													Identifier: &tk[0],
													Tokens:     tk[:1],
												}),
												Tokens: tk[:1],
											},
											Tokens: tk[:1],
										},
										Tokens: tk[:1],
									},
									Comments: Comments{tk[2], tk[4]},
									Tokens:   tk[:5],
								},
							},
							Tokens: tk[:5],
						},
						Tokens: tk[:5],
					},
				},
				Comments: Comments{tk[7]},
				Tokens:   tk[:8],
			}
		}},
	}, func(t *test) (Type, error) {
		var f File

		err := f.parse(t.Tokens)

		return f, err
	})
}

func TestErrUnwrap(t *testing.T) {
	err := Error{
		Err: Error{
			Err: Error{
				Err: Error{
					Err: Error{
						Err:     ErrMissingIdentifier,
						Parsing: "NonLocalStatement",
					},
					Parsing: "SimpleStatement",
				},
				Parsing: "StatementList",
			},
			Parsing: "Statement",
		},
		Parsing: "File",
	}

	if !errors.Is(err, ErrMissingIdentifier) {
		t.Errorf("error could not be correctly unwrapped")
	}
}

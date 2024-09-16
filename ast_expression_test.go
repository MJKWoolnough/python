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
		{`a(nonlocal)`, func(t *test, tk Tokens) { // 8
			t.Err = Error{
				Err: Error{
					Err: Error{
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
								Parsing: "AssignmentExpression",
								Token:   tk[2],
							},
							Parsing: "PositionalArgument",
							Token:   tk[2],
						},
						Parsing: "ArgumentList",
						Token:   tk[2],
					},
					Parsing: "ArgumentListOrComprehension",
					Token:   tk[2],
				},
				Parsing: "PrimaryExpression",
				Token:   tk[2],
			}
		}},
		{`a()`, func(t *test, tk Tokens) { // 9
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
						Tokens: tk[2:2],
					},
					Tokens: tk[2:2],
				},
				Tokens: tk[:3],
			}
		}},
		{`a(a for i in x)`, func(t *test, tk Tokens) { // 10
			t.Output = PrimaryExpression{
				PrimaryExpression: &PrimaryExpression{
					Atom: &Atom{
						Identifier: &tk[0],
						Tokens:     tk[:1],
					},
					Tokens: tk[:1],
				},
				Call: &ArgumentListOrComprehension{
					Comprehension: &Comprehension{
						AssignmentExpression: AssignmentExpression{
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
						ComprehensionFor: ComprehensionFor{
							TargetList: TargetList{
								Targets: []Target{
									{
										PrimaryExpression: &PrimaryExpression{
											Atom: &Atom{
												Identifier: &tk[6],
												Tokens:     tk[6:7],
											},
											Tokens: tk[6:7],
										},
										Tokens: tk[6:7],
									},
								},
								Tokens: tk[6:7],
							},
							OrTest: WrapConditional(&PrimaryExpression{
								Atom: &Atom{
									Identifier: &tk[10],
									Tokens:     tk[10:11],
								},
								Tokens: tk[10:11],
							}).OrTest,
							Tokens: tk[4:11],
						},
						Tokens: tk[2:11],
					},
					Tokens: tk[2:11],
				},
				Tokens: tk[:12],
			}
		}},
		{`a(a for 1 in x)`, func(t *test, tk Tokens) { // 11
			t.Err = Error{
				Err: Error{
					Err: Error{
						Err: Error{
							Err: Error{
								Err: Error{
									Err:     ErrMissingIdentifier,
									Parsing: "Target",
									Token:   tk[6],
								},
								Parsing: "TargetList",
								Token:   tk[6],
							},
							Parsing: "ComprehensionFor",
							Token:   tk[6],
						},
						Parsing: "Comprehension",
						Token:   tk[4],
					},
					Parsing: "ArgumentListOrComprehension",
					Token:   tk[2],
				},
				Parsing: "PrimaryExpression",
				Token:   tk[2],
			}
		}},
		{`a(a for i() in x)`, func(t *test, tk Tokens) { // 12
			t.Err = Error{
				Err: Error{
					Err: Error{
						Err: Error{
							Err: Error{
								Err: Error{
									Err:     ErrMissingIdentifier,
									Parsing: "Target",
									Token:   tk[6],
								},
								Parsing: "TargetList",
								Token:   tk[6],
							},
							Parsing: "ComprehensionFor",
							Token:   tk[6],
						},
						Parsing: "Comprehension",
						Token:   tk[4],
					},
					Parsing: "ArgumentListOrComprehension",
					Token:   tk[2],
				},
				Parsing: "PrimaryExpression",
				Token:   tk[2],
			}
		}},
		{`a.b[c](d).e`, func(t *test, tk Tokens) { // 13
			t.Output = PrimaryExpression{
				PrimaryExpression: &PrimaryExpression{
					PrimaryExpression: &PrimaryExpression{
						PrimaryExpression: &PrimaryExpression{
							PrimaryExpression: &PrimaryExpression{
								Atom: &Atom{
									Identifier: &tk[0],
									Tokens:     tk[:1],
								},
								Tokens: tk[:1],
							},
							AttributeRef: &tk[2],
							Tokens:       tk[:3],
						},
						Slicing: &SliceList{
							SliceItems: []SliceItem{
								{
									Expression: &Expression{
										ConditionalExpression: WrapConditional(&PrimaryExpression{
											Atom: &Atom{
												Identifier: &tk[4],
												Tokens:     tk[4:5],
											},
											Tokens: tk[4:5],
										}),
										Tokens: tk[4:5],
									},
									Tokens: tk[4:5],
								},
							},
							Tokens: tk[4:5],
						},
						Tokens: tk[:6],
					},
					Call: &ArgumentListOrComprehension{
						ArgumentList: &ArgumentList{
							PositionalArguments: []PositionalArgument{
								{
									AssignmentExpression: &AssignmentExpression{
										Expression: Expression{
											ConditionalExpression: WrapConditional(&PrimaryExpression{
												Atom: &Atom{
													Identifier: &tk[7],
													Tokens:     tk[7:8],
												},
												Tokens: tk[7:8],
											}),
											Tokens: tk[7:8],
										},
										Tokens: tk[7:8],
									},
									Tokens: tk[7:8],
								},
							},
							Tokens: tk[7:8],
						},
						Tokens: tk[7:8],
					},
					Tokens: tk[:9],
				},
				AttributeRef: &tk[10],
				Tokens:       tk[:11],
			}
		}},
	}, func(t *test) (Type, error) {
		var pe PrimaryExpression

		err := pe.parse(t.Tokens)

		return pe, err
	})
}

func TestPowerExpression(t *testing.T) {
	doTests(t, []sourceFn{
		{`a`, func(t *test, tk Tokens) { // 1
			t.Output = PowerExpression{
				PrimaryExpression: PrimaryExpression{
					Atom: &Atom{
						Identifier: &tk[0],
						Tokens:     tk[:1],
					},
					Tokens: tk[:1],
				},
				Tokens: tk[:1],
			}
		}},
		{`await a`, func(t *test, tk Tokens) { // 2
			t.Output = PowerExpression{
				AwaitExpression: true,
				PrimaryExpression: PrimaryExpression{
					Atom: &Atom{
						Identifier: &tk[2],
						Tokens:     tk[2:3],
					},
					Tokens: tk[2:3],
				},
				Tokens: tk[:3],
			}
		}},
		{`a ** b`, func(t *test, tk Tokens) { // 3
			t.Output = PowerExpression{
				PrimaryExpression: PrimaryExpression{
					Atom: &Atom{
						Identifier: &tk[0],
						Tokens:     tk[:1],
					},
					Tokens: tk[:1],
				},
				UnaryExpression: &UnaryExpression{
					PowerExpression: &PowerExpression{
						PrimaryExpression: PrimaryExpression{
							Atom: &Atom{
								Identifier: &tk[4],
								Tokens:     tk[4:5],
							},
							Tokens: tk[4:5],
						},
						Tokens: tk[4:5],
					},
					Tokens: tk[4:5],
				},
				Tokens: tk[:5],
			}
		}},
		{`await a ** b`, func(t *test, tk Tokens) { // 4
			t.Output = PowerExpression{
				AwaitExpression: true,
				PrimaryExpression: PrimaryExpression{
					Atom: &Atom{
						Identifier: &tk[2],
						Tokens:     tk[2:3],
					},
					Tokens: tk[2:3],
				},
				UnaryExpression: &UnaryExpression{
					PowerExpression: &PowerExpression{
						PrimaryExpression: PrimaryExpression{
							Atom: &Atom{
								Identifier: &tk[6],
								Tokens:     tk[6:7],
							},
							Tokens: tk[6:7],
						},
						Tokens: tk[6:7],
					},
					Tokens: tk[6:7],
				},
				Tokens: tk[:7],
			}
		}},
		{`a ** b ** c`, func(t *test, tk Tokens) { // 5
			t.Output = PowerExpression{
				PrimaryExpression: PrimaryExpression{
					Atom: &Atom{
						Identifier: &tk[0],
						Tokens:     tk[:1],
					},
					Tokens: tk[:1],
				},
				UnaryExpression: &UnaryExpression{
					PowerExpression: &PowerExpression{
						PrimaryExpression: PrimaryExpression{
							Atom: &Atom{
								Identifier: &tk[4],
								Tokens:     tk[4:5],
							},
							Tokens: tk[4:5],
						},
						UnaryExpression: &UnaryExpression{
							PowerExpression: &PowerExpression{
								PrimaryExpression: PrimaryExpression{
									Atom: &Atom{
										Identifier: &tk[8],
										Tokens:     tk[8:9],
									},
									Tokens: tk[8:9],
								},
								Tokens: tk[8:9],
							},
							Tokens: tk[8:9],
						},
						Tokens: tk[4:9],
					},
					Tokens: tk[4:9],
				},
				Tokens: tk[:9],
			}
		}},
		{`nonlocal`, func(t *test, tk Tokens) { // 6
			t.Err = Error{
				Err: Error{
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
				},
				Parsing: "PowerExpression",
				Token:   tk[0],
			}
		}},
		{`1 ** nonlocal`, func(t *test, tk Tokens) { // 7
			t.Err = Error{
				Err: Error{
					Err: Error{
						Err: Error{
							Err: Error{
								Err: Error{
									Err:     ErrInvalidEnclosure,
									Parsing: "Enclosure",
									Token:   tk[4],
								},
								Parsing: "Atom",
								Token:   tk[4],
							},
							Parsing: "PrimaryExpression",
							Token:   tk[4],
						},
						Parsing: "PowerExpression",
						Token:   tk[4],
					},
					Parsing: "UnaryExpression",
					Token:   tk[4],
				},
				Parsing: "PowerExpression",
				Token:   tk[4],
			}
		}},
	}, func(t *test) (Type, error) {
		var pe PowerExpression

		err := pe.parse(t.Tokens)

		return pe, err
	})
}

func TestUnaryExpression(t *testing.T) {
	doTests(t, []sourceFn{
		{`a`, func(t *test, tk Tokens) { // 1
			t.Output = UnaryExpression{
				PowerExpression: &PowerExpression{
					PrimaryExpression: PrimaryExpression{
						Atom: &Atom{
							Identifier: &tk[0],
							Tokens:     tk[:1],
						},
						Tokens: tk[:1],
					},
					Tokens: tk[:1],
				},
				Tokens: tk[:1],
			}
		}},
		{`~True`, func(t *test, tk Tokens) { // 2
			t.Output = UnaryExpression{
				Unary: &tk[0],
				UnaryExpression: &UnaryExpression{
					PowerExpression: &PowerExpression{
						PrimaryExpression: PrimaryExpression{
							Atom: &Atom{
								Literal: &tk[1],
								Tokens:  tk[1:2],
							},
							Tokens: tk[1:2],
						},
						Tokens: tk[1:2],
					},
					Tokens: tk[1:2],
				},
				Tokens: tk[:2],
			}
		}},
		{`-1`, func(t *test, tk Tokens) { // 3
			t.Output = UnaryExpression{
				Unary: &tk[0],
				UnaryExpression: &UnaryExpression{
					PowerExpression: &PowerExpression{
						PrimaryExpression: PrimaryExpression{
							Atom: &Atom{
								Literal: &tk[1],
								Tokens:  tk[1:2],
							},
							Tokens: tk[1:2],
						},
						Tokens: tk[1:2],
					},
					Tokens: tk[1:2],
				},
				Tokens: tk[:2],
			}
		}},
		{`+ a`, func(t *test, tk Tokens) { // 4
			t.Output = UnaryExpression{
				Unary: &tk[0],
				UnaryExpression: &UnaryExpression{
					PowerExpression: &PowerExpression{
						PrimaryExpression: PrimaryExpression{
							Atom: &Atom{
								Identifier: &tk[2],
								Tokens:     tk[2:3],
							},
							Tokens: tk[2:3],
						},
						Tokens: tk[2:3],
					},
					Tokens: tk[2:3],
				},
				Tokens: tk[:3],
			}
		}},
		{`~-a`, func(t *test, tk Tokens) { // 4
			t.Output = UnaryExpression{
				Unary: &tk[0],
				UnaryExpression: &UnaryExpression{
					Unary: &tk[1],
					UnaryExpression: &UnaryExpression{
						PowerExpression: &PowerExpression{
							PrimaryExpression: PrimaryExpression{
								Atom: &Atom{
									Identifier: &tk[2],
									Tokens:     tk[2:3],
								},
								Tokens: tk[2:3],
							},
							Tokens: tk[2:3],
						},
						Tokens: tk[2:3],
					},
					Tokens: tk[1:3],
				},
				Tokens: tk[:3],
			}
		}},
		{`nonlocal`, func(t *test, tk Tokens) { // 5
			t.Err = Error{
				Err: Error{
					Err: Error{
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
					},
					Parsing: "PowerExpression",
					Token:   tk[0],
				},
				Parsing: "UnaryExpression",
				Token:   tk[0],
			}
		}},
		{`+nonlocal`, func(t *test, tk Tokens) { // 6
			t.Err = Error{
				Err: Error{
					Err: Error{
						Err: Error{
							Err: Error{
								Err: Error{
									Err:     ErrInvalidEnclosure,
									Parsing: "Enclosure",
									Token:   tk[1],
								},
								Parsing: "Atom",
								Token:   tk[1],
							},
							Parsing: "PrimaryExpression",
							Token:   tk[1],
						},
						Parsing: "PowerExpression",
						Token:   tk[1],
					},
					Parsing: "UnaryExpression",
					Token:   tk[1],
				},
				Parsing: "UnaryExpression",
				Token:   tk[1],
			}
		}},
	}, func(t *test) (Type, error) {
		var ue UnaryExpression

		err := ue.parse(t.Tokens)

		return ue, err
	})
}

func TestMultiplyExpression(t *testing.T) {
	doTests(t, []sourceFn{
		{`a`, func(t *test, tk Tokens) { // 1
			t.Output = MultiplyExpression{
				UnaryExpression: UnaryExpression{
					PowerExpression: &PowerExpression{
						PrimaryExpression: PrimaryExpression{
							Atom: &Atom{
								Identifier: &tk[0],
								Tokens:     tk[:1],
							},
							Tokens: tk[:1],
						},
						Tokens: tk[:1],
					},
					Tokens: tk[:1],
				},
				Tokens: tk[:1],
			}
		}},
		{`a * b`, func(t *test, tk Tokens) { // 2
			t.Output = MultiplyExpression{
				UnaryExpression: UnaryExpression{
					PowerExpression: &PowerExpression{
						PrimaryExpression: PrimaryExpression{
							Atom: &Atom{
								Identifier: &tk[0],
								Tokens:     tk[:1],
							},
							Tokens: tk[:1],
						},
						Tokens: tk[:1],
					},
					Tokens: tk[:1],
				},
				Multiply: &tk[2],
				MultiplyExpression: &MultiplyExpression{
					UnaryExpression: UnaryExpression{
						PowerExpression: &PowerExpression{
							PrimaryExpression: PrimaryExpression{
								Atom: &Atom{
									Identifier: &tk[4],
									Tokens:     tk[4:5],
								},
								Tokens: tk[4:5],
							},
							Tokens: tk[4:5],
						},
						Tokens: tk[4:5],
					},
					Tokens: tk[4:5],
				},
				Tokens: tk[:5],
			}
		}},
		{`a @ b`, func(t *test, tk Tokens) { // 4
			t.Output = MultiplyExpression{
				UnaryExpression: UnaryExpression{
					PowerExpression: &PowerExpression{
						PrimaryExpression: PrimaryExpression{
							Atom: &Atom{
								Identifier: &tk[0],
								Tokens:     tk[:1],
							},
							Tokens: tk[:1],
						},
						Tokens: tk[:1],
					},
					Tokens: tk[:1],
				},
				Multiply: &tk[2],
				MultiplyExpression: &MultiplyExpression{
					UnaryExpression: UnaryExpression{
						PowerExpression: &PowerExpression{
							PrimaryExpression: PrimaryExpression{
								Atom: &Atom{
									Identifier: &tk[4],
									Tokens:     tk[4:5],
								},
								Tokens: tk[4:5],
							},
							Tokens: tk[4:5],
						},
						Tokens: tk[4:5],
					},
					Tokens: tk[4:5],
				},
				Tokens: tk[:5],
			}
		}},
		{`a//b`, func(t *test, tk Tokens) { // 5
			t.Output = MultiplyExpression{
				UnaryExpression: UnaryExpression{
					PowerExpression: &PowerExpression{
						PrimaryExpression: PrimaryExpression{
							Atom: &Atom{
								Identifier: &tk[0],
								Tokens:     tk[:1],
							},
							Tokens: tk[:1],
						},
						Tokens: tk[:1],
					},
					Tokens: tk[:1],
				},
				Multiply: &tk[1],
				MultiplyExpression: &MultiplyExpression{
					UnaryExpression: UnaryExpression{
						PowerExpression: &PowerExpression{
							PrimaryExpression: PrimaryExpression{
								Atom: &Atom{
									Identifier: &tk[2],
									Tokens:     tk[2:3],
								},
								Tokens: tk[2:3],
							},
							Tokens: tk[2:3],
						},
						Tokens: tk[2:3],
					},
					Tokens: tk[2:3],
				},
				Tokens: tk[:3],
			}
		}},
		{`a / b`, func(t *test, tk Tokens) { // 6
			t.Output = MultiplyExpression{
				UnaryExpression: UnaryExpression{
					PowerExpression: &PowerExpression{
						PrimaryExpression: PrimaryExpression{
							Atom: &Atom{
								Identifier: &tk[0],
								Tokens:     tk[:1],
							},
							Tokens: tk[:1],
						},
						Tokens: tk[:1],
					},
					Tokens: tk[:1],
				},
				Multiply: &tk[2],
				MultiplyExpression: &MultiplyExpression{
					UnaryExpression: UnaryExpression{
						PowerExpression: &PowerExpression{
							PrimaryExpression: PrimaryExpression{
								Atom: &Atom{
									Identifier: &tk[4],
									Tokens:     tk[4:5],
								},
								Tokens: tk[4:5],
							},
							Tokens: tk[4:5],
						},
						Tokens: tk[4:5],
					},
					Tokens: tk[4:5],
				},
				Tokens: tk[:5],
			}
		}},
		{`a % b`, func(t *test, tk Tokens) { // 7
			t.Output = MultiplyExpression{
				UnaryExpression: UnaryExpression{
					PowerExpression: &PowerExpression{
						PrimaryExpression: PrimaryExpression{
							Atom: &Atom{
								Identifier: &tk[0],
								Tokens:     tk[:1],
							},
							Tokens: tk[:1],
						},
						Tokens: tk[:1],
					},
					Tokens: tk[:1],
				},
				Multiply: &tk[2],
				MultiplyExpression: &MultiplyExpression{
					UnaryExpression: UnaryExpression{
						PowerExpression: &PowerExpression{
							PrimaryExpression: PrimaryExpression{
								Atom: &Atom{
									Identifier: &tk[4],
									Tokens:     tk[4:5],
								},
								Tokens: tk[4:5],
							},
							Tokens: tk[4:5],
						},
						Tokens: tk[4:5],
					},
					Tokens: tk[4:5],
				},
				Tokens: tk[:5],
			}
		}},
		{`a / b @ c * d`, func(t *test, tk Tokens) { // 8
			t.Output = MultiplyExpression{
				UnaryExpression: UnaryExpression{
					PowerExpression: &PowerExpression{
						PrimaryExpression: PrimaryExpression{
							Atom: &Atom{
								Identifier: &tk[0],
								Tokens:     tk[:1],
							},
							Tokens: tk[:1],
						},
						Tokens: tk[:1],
					},
					Tokens: tk[:1],
				},
				Multiply: &tk[2],
				MultiplyExpression: &MultiplyExpression{
					UnaryExpression: UnaryExpression{
						PowerExpression: &PowerExpression{
							PrimaryExpression: PrimaryExpression{
								Atom: &Atom{
									Identifier: &tk[4],
									Tokens:     tk[4:5],
								},
								Tokens: tk[4:5],
							},
							Tokens: tk[4:5],
						},
						Tokens: tk[4:5],
					},
					Multiply: &tk[6],
					MultiplyExpression: &MultiplyExpression{
						UnaryExpression: UnaryExpression{
							PowerExpression: &PowerExpression{
								PrimaryExpression: PrimaryExpression{
									Atom: &Atom{
										Identifier: &tk[8],
										Tokens:     tk[8:9],
									},
									Tokens: tk[8:9],
								},
								Tokens: tk[8:9],
							},
							Tokens: tk[8:9],
						},
						Multiply: &tk[10],
						MultiplyExpression: &MultiplyExpression{
							UnaryExpression: UnaryExpression{
								PowerExpression: &PowerExpression{
									PrimaryExpression: PrimaryExpression{
										Atom: &Atom{
											Identifier: &tk[12],
											Tokens:     tk[12:13],
										},
										Tokens: tk[12:13],
									},
									Tokens: tk[12:13],
								},
								Tokens: tk[12:13],
							},
							Tokens: tk[12:13],
						},
						Tokens: tk[8:13],
					},
					Tokens: tk[4:13],
				},
				Tokens: tk[:13],
			}
		}},
		{`nonlocal`, func(t *test, tk Tokens) { // 9
			t.Err = Error{
				Err: Error{
					Err: Error{
						Err: Error{
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
						},
						Parsing: "PowerExpression",
						Token:   tk[0],
					},
					Parsing: "UnaryExpression",
					Token:   tk[0],
				},
				Parsing: "MultiplyExpression",
				Token:   tk[0],
			}
		}},
		{`1 * nonlocal`, func(t *test, tk Tokens) { // 10
			t.Err = Error{
				Err: Error{
					Err: Error{
						Err: Error{
							Err: Error{
								Err: Error{
									Err: Error{
										Err:     ErrInvalidEnclosure,
										Parsing: "Enclosure",
										Token:   tk[4],
									},
									Parsing: "Atom",
									Token:   tk[4],
								},
								Parsing: "PrimaryExpression",
								Token:   tk[4],
							},
							Parsing: "PowerExpression",
							Token:   tk[4],
						},
						Parsing: "UnaryExpression",
						Token:   tk[4],
					},
					Parsing: "MultiplyExpression",
					Token:   tk[4],
				},
				Parsing: "MultiplyExpression",
				Token:   tk[4],
			}
		}},
	}, func(t *test) (Type, error) {
		var me MultiplyExpression

		err := me.parse(t.Tokens)

		return me, err
	})
}

func TestAddExpression(t *testing.T) {
	doTests(t, []sourceFn{
		{`a`, func(t *test, tk Tokens) { // 1
			t.Output = AddExpression{
				MultiplyExpression: MultiplyExpression{
					UnaryExpression: UnaryExpression{
						PowerExpression: &PowerExpression{
							PrimaryExpression: PrimaryExpression{
								Atom: &Atom{
									Identifier: &tk[0],
									Tokens:     tk[:1],
								},
								Tokens: tk[:1],
							},
							Tokens: tk[:1],
						},
						Tokens: tk[:1],
					},
					Tokens: tk[:1],
				},
				Tokens: tk[:1],
			}
		}},
		{`a+b`, func(t *test, tk Tokens) { // 2
			t.Output = AddExpression{
				MultiplyExpression: MultiplyExpression{
					UnaryExpression: UnaryExpression{
						PowerExpression: &PowerExpression{
							PrimaryExpression: PrimaryExpression{
								Atom: &Atom{
									Identifier: &tk[0],
									Tokens:     tk[:1],
								},
								Tokens: tk[:1],
							},
							Tokens: tk[:1],
						},
						Tokens: tk[:1],
					},
					Tokens: tk[:1],
				},
				Add: &tk[1],
				AddExpression: &AddExpression{
					MultiplyExpression: MultiplyExpression{
						UnaryExpression: UnaryExpression{
							PowerExpression: &PowerExpression{
								PrimaryExpression: PrimaryExpression{
									Atom: &Atom{
										Identifier: &tk[2],
										Tokens:     tk[2:3],
									},
									Tokens: tk[2:3],
								},
								Tokens: tk[2:3],
							},
							Tokens: tk[2:3],
						},
						Tokens: tk[2:3],
					},
					Tokens: tk[2:3],
				},
				Tokens: tk[:3],
			}
		}},
		{`a - b`, func(t *test, tk Tokens) { // 3
			t.Output = AddExpression{
				MultiplyExpression: MultiplyExpression{
					UnaryExpression: UnaryExpression{
						PowerExpression: &PowerExpression{
							PrimaryExpression: PrimaryExpression{
								Atom: &Atom{
									Identifier: &tk[0],
									Tokens:     tk[:1],
								},
								Tokens: tk[:1],
							},
							Tokens: tk[:1],
						},
						Tokens: tk[:1],
					},
					Tokens: tk[:1],
				},
				Add: &tk[2],
				AddExpression: &AddExpression{
					MultiplyExpression: MultiplyExpression{
						UnaryExpression: UnaryExpression{
							PowerExpression: &PowerExpression{
								PrimaryExpression: PrimaryExpression{
									Atom: &Atom{
										Identifier: &tk[4],
										Tokens:     tk[4:5],
									},
									Tokens: tk[4:5],
								},
								Tokens: tk[4:5],
							},
							Tokens: tk[4:5],
						},
						Tokens: tk[4:5],
					},
					Tokens: tk[4:5],
				},
				Tokens: tk[:5],
			}
		}},
		{`a + b-c`, func(t *test, tk Tokens) { // 4
			t.Output = AddExpression{
				MultiplyExpression: MultiplyExpression{
					UnaryExpression: UnaryExpression{
						PowerExpression: &PowerExpression{
							PrimaryExpression: PrimaryExpression{
								Atom: &Atom{
									Identifier: &tk[0],
									Tokens:     tk[:1],
								},
								Tokens: tk[:1],
							},
							Tokens: tk[:1],
						},
						Tokens: tk[:1],
					},
					Tokens: tk[:1],
				},
				Add: &tk[2],
				AddExpression: &AddExpression{
					MultiplyExpression: MultiplyExpression{
						UnaryExpression: UnaryExpression{
							PowerExpression: &PowerExpression{
								PrimaryExpression: PrimaryExpression{
									Atom: &Atom{
										Identifier: &tk[4],
										Tokens:     tk[4:5],
									},
									Tokens: tk[4:5],
								},
								Tokens: tk[4:5],
							},
							Tokens: tk[4:5],
						},
						Tokens: tk[4:5],
					},
					Add: &tk[5],
					AddExpression: &AddExpression{
						MultiplyExpression: MultiplyExpression{
							UnaryExpression: UnaryExpression{
								PowerExpression: &PowerExpression{
									PrimaryExpression: PrimaryExpression{
										Atom: &Atom{
											Identifier: &tk[6],
											Tokens:     tk[6:7],
										},
										Tokens: tk[6:7],
									},
									Tokens: tk[6:7],
								},
								Tokens: tk[6:7],
							},
							Tokens: tk[6:7],
						},
						Tokens: tk[6:7],
					},
					Tokens: tk[4:7],
				},
				Tokens: tk[:7],
			}
		}},
		{`nonlocal`, func(t *test, tk Tokens) { // 5
			t.Err = Error{
				Err: Error{
					Err: Error{
						Err: Error{
							Err: Error{
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
							},
							Parsing: "PowerExpression",
							Token:   tk[0],
						},
						Parsing: "UnaryExpression",
						Token:   tk[0],
					},
					Parsing: "MultiplyExpression",
					Token:   tk[0],
				},
				Parsing: "AddExpression",
				Token:   tk[0],
			}
		}},
		{`1 + nonlocal`, func(t *test, tk Tokens) { // 6
			t.Err = Error{
				Err: Error{
					Err: Error{
						Err: Error{
							Err: Error{
								Err: Error{
									Err: Error{
										Err: Error{
											Err:     ErrInvalidEnclosure,
											Parsing: "Enclosure",
											Token:   tk[4],
										},
										Parsing: "Atom",
										Token:   tk[4],
									},
									Parsing: "PrimaryExpression",
									Token:   tk[4],
								},
								Parsing: "PowerExpression",
								Token:   tk[4],
							},
							Parsing: "UnaryExpression",
							Token:   tk[4],
						},
						Parsing: "MultiplyExpression",
						Token:   tk[4],
					},
					Parsing: "AddExpression",
					Token:   tk[4],
				},
				Parsing: "AddExpression",
				Token:   tk[4],
			}
		}},
	}, func(t *test) (Type, error) {
		var ae AddExpression

		err := ae.parse(t.Tokens)

		return ae, err
	})
}

func TestShiftExpression(t *testing.T) {
	doTests(t, []sourceFn{
		{`a`, func(t *test, tk Tokens) { // 1
			t.Output = ShiftExpression{
				AddExpression: AddExpression{
					MultiplyExpression: MultiplyExpression{
						UnaryExpression: UnaryExpression{
							PowerExpression: &PowerExpression{
								PrimaryExpression: PrimaryExpression{
									Atom: &Atom{
										Identifier: &tk[0],
										Tokens:     tk[:1],
									},
									Tokens: tk[:1],
								},
								Tokens: tk[:1],
							},
							Tokens: tk[:1],
						},
						Tokens: tk[:1],
					},
					Tokens: tk[:1],
				},
				Tokens: tk[:1],
			}
		}},
		{`a<<b`, func(t *test, tk Tokens) { // 2
			t.Output = ShiftExpression{
				AddExpression: AddExpression{
					MultiplyExpression: MultiplyExpression{
						UnaryExpression: UnaryExpression{
							PowerExpression: &PowerExpression{
								PrimaryExpression: PrimaryExpression{
									Atom: &Atom{
										Identifier: &tk[0],
										Tokens:     tk[:1],
									},
									Tokens: tk[:1],
								},
								Tokens: tk[:1],
							},
							Tokens: tk[:1],
						},
						Tokens: tk[:1],
					},
					Tokens: tk[:1],
				},
				Shift: &tk[1],
				ShiftExpression: &ShiftExpression{
					AddExpression: AddExpression{
						MultiplyExpression: MultiplyExpression{
							UnaryExpression: UnaryExpression{
								PowerExpression: &PowerExpression{
									PrimaryExpression: PrimaryExpression{
										Atom: &Atom{
											Identifier: &tk[2],
											Tokens:     tk[2:3],
										},
										Tokens: tk[2:3],
									},
									Tokens: tk[2:3],
								},
								Tokens: tk[2:3],
							},
							Tokens: tk[2:3],
						},
						Tokens: tk[2:3],
					},
					Tokens: tk[2:3],
				},
				Tokens: tk[:3],
			}
		}},
		{`a >> b`, func(t *test, tk Tokens) { // 3
			t.Output = ShiftExpression{
				AddExpression: AddExpression{
					MultiplyExpression: MultiplyExpression{
						UnaryExpression: UnaryExpression{
							PowerExpression: &PowerExpression{
								PrimaryExpression: PrimaryExpression{
									Atom: &Atom{
										Identifier: &tk[0],
										Tokens:     tk[:1],
									},
									Tokens: tk[:1],
								},
								Tokens: tk[:1],
							},
							Tokens: tk[:1],
						},
						Tokens: tk[:1],
					},
					Tokens: tk[:1],
				},
				Shift: &tk[2],
				ShiftExpression: &ShiftExpression{
					AddExpression: AddExpression{
						MultiplyExpression: MultiplyExpression{
							UnaryExpression: UnaryExpression{
								PowerExpression: &PowerExpression{
									PrimaryExpression: PrimaryExpression{
										Atom: &Atom{
											Identifier: &tk[4],
											Tokens:     tk[4:5],
										},
										Tokens: tk[4:5],
									},
									Tokens: tk[4:5],
								},
								Tokens: tk[4:5],
							},
							Tokens: tk[4:5],
						},
						Tokens: tk[4:5],
					},
					Tokens: tk[4:5],
				},
				Tokens: tk[:5],
			}
		}},
		{`a << b>>c`, func(t *test, tk Tokens) { // 4
			t.Output = ShiftExpression{
				AddExpression: AddExpression{
					MultiplyExpression: MultiplyExpression{
						UnaryExpression: UnaryExpression{
							PowerExpression: &PowerExpression{
								PrimaryExpression: PrimaryExpression{
									Atom: &Atom{
										Identifier: &tk[0],
										Tokens:     tk[:1],
									},
									Tokens: tk[:1],
								},
								Tokens: tk[:1],
							},
							Tokens: tk[:1],
						},
						Tokens: tk[:1],
					},
					Tokens: tk[:1],
				},
				Shift: &tk[2],
				ShiftExpression: &ShiftExpression{
					AddExpression: AddExpression{
						MultiplyExpression: MultiplyExpression{
							UnaryExpression: UnaryExpression{
								PowerExpression: &PowerExpression{
									PrimaryExpression: PrimaryExpression{
										Atom: &Atom{
											Identifier: &tk[4],
											Tokens:     tk[4:5],
										},
										Tokens: tk[4:5],
									},
									Tokens: tk[4:5],
								},
								Tokens: tk[4:5],
							},
							Tokens: tk[4:5],
						},
						Tokens: tk[4:5],
					},
					Shift: &tk[5],
					ShiftExpression: &ShiftExpression{
						AddExpression: AddExpression{
							MultiplyExpression: MultiplyExpression{
								UnaryExpression: UnaryExpression{
									PowerExpression: &PowerExpression{
										PrimaryExpression: PrimaryExpression{
											Atom: &Atom{
												Identifier: &tk[6],
												Tokens:     tk[6:7],
											},
											Tokens: tk[6:7],
										},
										Tokens: tk[6:7],
									},
									Tokens: tk[6:7],
								},
								Tokens: tk[6:7],
							},
							Tokens: tk[6:7],
						},
						Tokens: tk[6:7],
					},
					Tokens: tk[4:7],
				},
				Tokens: tk[:7],
			}
		}},
	}, func(t *test) (Type, error) {
		var se ShiftExpression

		err := se.parse(t.Tokens)

		return se, err
	})
}

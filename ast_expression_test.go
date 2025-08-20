package python

import "testing"

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
		{`a . b`, func(t *test, tk Tokens) { // 3
			t.Output = PrimaryExpression{
				PrimaryExpression: &PrimaryExpression{
					Atom: &Atom{
						Identifier: &tk[0],
						Tokens:     tk[:1],
					},
					Tokens: tk[:1],
				},
				AttributeRef: &tk[4],
				Tokens:       tk[:5],
			}
		}},
		{`a[b]`, func(t *test, tk Tokens) { // 4
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
								ConditionalExpression: WrapConditional(&Atom{
									Identifier: &tk[2],
									Tokens:     tk[2:3],
								}),
								Tokens: tk[2:3],
							},
							Tokens: tk[2:3],
						},
					},
					Tokens: tk[1:4],
				},
				Tokens: tk[:4],
			}
		}},
		{`a(b)`, func(t *test, tk Tokens) { // 5
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
						},
						Tokens: tk[2:3],
					},
					Tokens: tk[1:4],
				},
				Tokens: tk[:4],
			}
		}},
		{"(a # A\n. # B\nb)", func(t *test, tk Tokens) { // 6
			t.Output = PrimaryExpression{
				PrimaryExpression: &PrimaryExpression{
					Atom: &Atom{
						Identifier: &tk[1],
						Tokens:     tk[1:2],
					},
					Tokens: tk[1:2],
				},
				AttributeRef: &tk[9],
				Comments:     [2]Comments{{tk[3]}, {tk[7]}},
				Tokens:       tk[1:10],
			}
		}},
		{"(a # A\n[b])", func(t *test, tk Tokens) { // 7
			t.Output = PrimaryExpression{
				PrimaryExpression: &PrimaryExpression{
					Atom: &Atom{
						Identifier: &tk[1],
						Tokens:     tk[1:2],
					},
					Tokens: tk[1:2],
				},
				Slicing: &SliceList{
					SliceItems: []SliceItem{
						{
							Expression: &Expression{
								ConditionalExpression: WrapConditional(&Atom{
									Identifier: &tk[6],
									Tokens:     tk[6:7],
								}),
								Tokens: tk[6:7],
							},
							Tokens: tk[6:7],
						},
					},
					Tokens: tk[5:8],
				},
				Comments: [2]Comments{{tk[3]}},
				Tokens:   tk[1:8],
			}
		}},
		{"(a # A\n(b))", func(t *test, tk Tokens) { // 8
			t.Output = PrimaryExpression{
				PrimaryExpression: &PrimaryExpression{
					Atom: &Atom{
						Identifier: &tk[1],
						Tokens:     tk[1:2],
					},
					Tokens: tk[1:2],
				},
				Call: &ArgumentListOrComprehension{
					ArgumentList: &ArgumentList{
						PositionalArguments: []PositionalArgument{
							{
								AssignmentExpression: &AssignmentExpression{
									Expression: Expression{
										ConditionalExpression: WrapConditional(&Atom{
											Identifier: &tk[6],
											Tokens:     tk[6:7],
										}),
										Tokens: tk[6:7],
									},
									Tokens: tk[6:7],
								},
								Tokens: tk[6:7],
							},
						},
						Tokens: tk[6:7],
					},
					Tokens: tk[5:8],
				},
				Comments: [2]Comments{{tk[3]}},
				Tokens:   tk[1:8],
			}
		}},
		{`nonlocal`, func(t *test, tk Tokens) { // 9
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
		{`a.nonlocal`, func(t *test, tk Tokens) { // 10
			t.Err = Error{
				Err:     ErrMissingIdentifier,
				Parsing: "PrimaryExpression",
				Token:   tk[2],
			}
		}},
		{`a[nonlocal]`, func(t *test, tk Tokens) { // 11
			t.Err = Error{
				Err: Error{
					Err: Error{
						Err: Error{
							Err: wrapConditionalExpressionError(Error{
								Err:     ErrInvalidEnclosure,
								Parsing: "Enclosure",
								Token:   tk[2],
							}),
							Parsing: "Expression",
							Token:   tk[2],
						},
						Parsing: "SliceItem",
						Token:   tk[2],
					},
					Parsing: "SliceList",
					Token:   tk[2],
				},
				Parsing: "PrimaryExpression",
				Token:   tk[1],
			}
		}},
		{`a(nonlocal)`, func(t *test, tk Tokens) { // 12
			t.Err = Error{
				Err: Error{
					Err: Error{
						Err: Error{
							Err: Error{
								Err: Error{
									Err: wrapConditionalExpressionError(Error{
										Err:     ErrInvalidEnclosure,
										Parsing: "Enclosure",
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
				Token:   tk[1],
			}
		}},
		{`a()`, func(t *test, tk Tokens) { // 13
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
					Tokens: tk[1:3],
				},
				Tokens: tk[:3],
			}
		}},
		{`a(a for i in x)`, func(t *test, tk Tokens) { // 14
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
								ConditionalExpression: WrapConditional(&Atom{
									Identifier: &tk[2],
									Tokens:     tk[2:3],
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
							OrTest: WrapConditional(&Atom{
								Identifier: &tk[10],
								Tokens:     tk[10:11],
							}).OrTest,
							Tokens: tk[4:11],
						},
						Tokens: tk[2:11],
					},
					Tokens: tk[1:12],
				},
				Tokens: tk[:12],
			}
		}},
		{`a(a for 1 in x)`, func(t *test, tk Tokens) { // 15
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
				Token:   tk[1],
			}
		}},
		{`a(a for i() in x)`, func(t *test, tk Tokens) { // 16
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
				Token:   tk[1],
			}
		}},
		{`a.b[c](d).e`, func(t *test, tk Tokens) { // 17
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
										ConditionalExpression: WrapConditional(&Atom{
											Identifier: &tk[4],
											Tokens:     tk[4:5],
										}),
										Tokens: tk[4:5],
									},
									Tokens: tk[4:5],
								},
							},
							Tokens: tk[3:6],
						},
						Tokens: tk[:6],
					},
					Call: &ArgumentListOrComprehension{
						ArgumentList: &ArgumentList{
							PositionalArguments: []PositionalArgument{
								{
									AssignmentExpression: &AssignmentExpression{
										Expression: Expression{
											ConditionalExpression: WrapConditional(&Atom{
												Identifier: &tk[7],
												Tokens:     tk[7:8],
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
						Tokens: tk[6:9],
					},
					Tokens: tk[:9],
				},
				AttributeRef: &tk[10],
				Tokens:       tk[:11],
			}
		}},
	}, func(t *test) (Type, error) {
		var pe PrimaryExpression

		if t.Tokens.Peek().Data == "(" {
			t.Tokens.Tokens = t.Tokens.Tokens[1:1]
		}

		err := pe.parse(t.Tokens)

		return pe, err
	})
}

func TestPrimaryExpressionIsIdentifier(t *testing.T) {
	for n, test := range [...]struct {
		Input       PrimaryExpression
		IsIdentifer bool
	}{
		{
			Input: PrimaryExpression{
				Atom: &Atom{
					Identifier: new(Token),
				},
			},
			IsIdentifer: true,
		},
		{
			Input: PrimaryExpression{
				PrimaryExpression: &PrimaryExpression{
					Atom: &Atom{
						Identifier: new(Token),
					},
				},
			},
			IsIdentifer: true,
		},
		{
			Input: PrimaryExpression{
				Atom: &Atom{
					Literal: new(Token),
				},
			},
			IsIdentifer: false,
		},
		{
			Input: PrimaryExpression{
				PrimaryExpression: &PrimaryExpression{
					Atom: &Atom{
						Literal: new(Token),
					},
				},
			},
			IsIdentifer: false,
		},
		{},
	} {
		if test.Input.IsIdentifier() != test.IsIdentifer {
			t.Errorf("test %d: expecting IsIdentifier() to return %v, got %v", n+1, test.IsIdentifer, !test.IsIdentifer)
		}
	}
}

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
					DictDisplay: &DictDisplay{
						Tokens: tk[1:1],
					},
					Tokens: tk[:2],
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

func TestEnclosure(t *testing.T) {
	doTests(t, []sourceFn{
		{`()`, func(t *test, tk Tokens) { // 1
			t.Output = Enclosure{
				ParenthForm: &StarredExpression{
					Tokens: tk[1:1],
				},
				Tokens: tk[:2],
			}
		}},
		{`( )`, func(t *test, tk Tokens) { // 2
			t.Output = Enclosure{
				ParenthForm: &StarredExpression{
					Tokens: tk[2:2],
				},
				Tokens: tk[:3],
			}
		}},
		{`(yield a)`, func(t *test, tk Tokens) { // 3
			t.Output = Enclosure{
				YieldAtom: &YieldExpression{
					ExpressionList: &ExpressionList{
						Expressions: []Expression{
							{
								ConditionalExpression: WrapConditional(&Atom{
									Identifier: &tk[3],
									Tokens:     tk[3:4],
								}),
								Tokens: tk[3:4],
							},
						},
						Tokens: tk[3:4],
					},
					Tokens: tk[1:4],
				},
				Tokens: tk[:5],
			}
		}},
		{`( yield a )`, func(t *test, tk Tokens) { // 4
			t.Output = Enclosure{
				YieldAtom: &YieldExpression{
					ExpressionList: &ExpressionList{
						Expressions: []Expression{
							{
								ConditionalExpression: WrapConditional(&Atom{
									Identifier: &tk[4],
									Tokens:     tk[4:5],
								}),
								Tokens: tk[4:5],
							},
						},
						Tokens: tk[4:5],
					},
					Tokens: tk[2:5],
				},
				Tokens: tk[:7],
			}
		}},
		{`(a for b in c)`, func(t *test, tk Tokens) { // 5
			t.Output = Enclosure{
				GeneratorExpression: &GeneratorExpression{
					Expression: Expression{
						ConditionalExpression: WrapConditional(&Atom{
							Identifier: &tk[1],
							Tokens:     tk[1:2],
						}),
						Tokens: tk[1:2],
					},
					ComprehensionFor: ComprehensionFor{
						TargetList: TargetList{
							Targets: []Target{
								{
									PrimaryExpression: &PrimaryExpression{
										Atom: &Atom{
											Identifier: &tk[5],
											Tokens:     tk[5:6],
										},
										Tokens: tk[5:6],
									},
									Tokens: tk[5:6],
								},
							},
							Tokens: tk[5:6],
						},
						OrTest: WrapConditional(&Atom{
							Identifier: &tk[9],
							Tokens:     tk[9:10],
						}).OrTest,
						Tokens: tk[3:10],
					},
					Tokens: tk[1:10],
				},
				Tokens: tk[:11],
			}
		}},
		{`( a for b in c )`, func(t *test, tk Tokens) { // 6
			t.Output = Enclosure{
				GeneratorExpression: &GeneratorExpression{
					Expression: Expression{
						ConditionalExpression: WrapConditional(&Atom{
							Identifier: &tk[2],
							Tokens:     tk[2:3],
						}),
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
						OrTest: WrapConditional(&Atom{
							Identifier: &tk[10],
							Tokens:     tk[10:11],
						}).OrTest,
						Tokens: tk[4:11],
					},
					Tokens: tk[2:11],
				},
				Tokens: tk[:13],
			}
		}},
		{`(a)`, func(t *test, tk Tokens) { // 7
			t.Output = Enclosure{
				ParenthForm: &StarredExpression{
					Expression: &Expression{
						ConditionalExpression: WrapConditional(&Atom{
							Identifier: &tk[1],
							Tokens:     tk[1:2],
						}),
						Tokens: tk[1:2],
					},
					Tokens: tk[1:2],
				},
				Tokens: tk[:3],
			}
		}},
		{`( a )`, func(t *test, tk Tokens) { // 8
			t.Output = Enclosure{
				ParenthForm: &StarredExpression{
					Expression: &Expression{
						ConditionalExpression: WrapConditional(&Atom{
							Identifier: &tk[2],
							Tokens:     tk[2:3],
						}),
						Tokens: tk[2:3],
					},
					Tokens: tk[2:3],
				},
				Tokens: tk[:5],
			}
		}},
		{`[]`, func(t *test, tk Tokens) { // 9
			t.Output = Enclosure{
				ListDisplay: &FlexibleExpressionListOrComprehension{
					Tokens: tk[1:1],
				},
				Tokens: tk[:2],
			}
		}},
		{`[ ]`, func(t *test, tk Tokens) { // 10
			t.Output = Enclosure{
				ListDisplay: &FlexibleExpressionListOrComprehension{
					Tokens: tk[2:2],
				},
				Tokens: tk[:3],
			}
		}},
		{`[a]`, func(t *test, tk Tokens) { // 11
			t.Output = Enclosure{
				ListDisplay: &FlexibleExpressionListOrComprehension{
					FlexibleExpressionList: &FlexibleExpressionList{
						FlexibleExpressions: []FlexibleExpression{
							{
								AssignmentExpression: &AssignmentExpression{
									Expression: Expression{
										ConditionalExpression: WrapConditional(&Atom{
											Identifier: &tk[1],
											Tokens:     tk[1:2],
										}),
										Tokens: tk[1:2],
									},
									Tokens: tk[1:2],
								},
								Tokens: tk[1:2],
							},
						},
						Tokens: tk[1:2],
					},
					Tokens: tk[1:2],
				},
				Tokens: tk[:3],
			}
		}},
		{`[ a ]`, func(t *test, tk Tokens) { // 12
			t.Output = Enclosure{
				ListDisplay: &FlexibleExpressionListOrComprehension{
					FlexibleExpressionList: &FlexibleExpressionList{
						FlexibleExpressions: []FlexibleExpression{
							{
								AssignmentExpression: &AssignmentExpression{
									Expression: Expression{
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
						},
						Tokens: tk[2:3],
					},
					Tokens: tk[2:3],
				},
				Tokens: tk[:5],
			}
		}},
		{`{}`, func(t *test, tk Tokens) { // 13
			t.Output = Enclosure{
				DictDisplay: &DictDisplay{
					Tokens: tk[1:1],
				},
				Tokens: tk[:2],
			}
		}},
		{`{ }`, func(t *test, tk Tokens) { // 14
			t.Output = Enclosure{
				DictDisplay: &DictDisplay{
					Tokens: tk[2:2],
				},
				Tokens: tk[:3],
			}
		}},
		{`{**a}`, func(t *test, tk Tokens) { // 15
			t.Output = Enclosure{
				DictDisplay: &DictDisplay{
					DictItems: []DictItem{
						{
							OrExpression: &WrapConditional(&Atom{
								Identifier: &tk[2],
								Tokens:     tk[2:3],
							}).OrTest.AndTest.NotTest.Comparison.OrExpression,
							Tokens: tk[1:3],
						},
					},
					Tokens: tk[1:3],
				},
				Tokens: tk[:4],
			}
		}},
		{`{ **a }`, func(t *test, tk Tokens) { // 16
			t.Output = Enclosure{
				DictDisplay: &DictDisplay{
					DictItems: []DictItem{
						{
							OrExpression: &WrapConditional(&Atom{
								Identifier: &tk[3],
								Tokens:     tk[3:4],
							}).OrTest.AndTest.NotTest.Comparison.OrExpression,
							Tokens: tk[2:4],
						},
					},
					Tokens: tk[2:4],
				},
				Tokens: tk[:6],
			}
		}},
		{`{*a}`, func(t *test, tk Tokens) { // 17
			t.Output = Enclosure{
				SetDisplay: &FlexibleExpressionListOrComprehension{
					FlexibleExpressionList: &FlexibleExpressionList{
						FlexibleExpressions: []FlexibleExpression{
							{
								StarredExpression: &StarredExpression{
									StarredList: &StarredList{
										StarredItems: []StarredItem{
											{
												OrExpr: &WrapConditional(&Atom{
													Identifier: &tk[2],
													Tokens:     tk[2:3],
												}).OrTest.AndTest.NotTest.Comparison.OrExpression,
												Tokens: tk[1:3],
											},
										},
										Tokens: tk[1:3],
									},
									Tokens: tk[1:3],
								},
								Tokens: tk[1:3],
							},
						},
						Tokens: tk[1:3],
					},
					Tokens: tk[1:3],
				},
				Tokens: tk[:4],
			}
		}},
		{`{ *a }`, func(t *test, tk Tokens) { // 18
			t.Output = Enclosure{
				SetDisplay: &FlexibleExpressionListOrComprehension{
					FlexibleExpressionList: &FlexibleExpressionList{
						FlexibleExpressions: []FlexibleExpression{
							{
								StarredExpression: &StarredExpression{
									StarredList: &StarredList{
										StarredItems: []StarredItem{
											{
												OrExpr: &WrapConditional(&Atom{
													Identifier: &tk[3],
													Tokens:     tk[3:4],
												}).OrTest.AndTest.NotTest.Comparison.OrExpression,
												Tokens: tk[2:4],
											},
										},
										Tokens: tk[2:4],
									},
									Tokens: tk[2:4],
								},
								Tokens: tk[2:4],
							},
						},
						Tokens: tk[2:4],
					},
					Tokens: tk[2:4],
				},
				Tokens: tk[:6],
			}
		}},
		{`{a:=b}`, func(t *test, tk Tokens) { // 19
			t.Output = Enclosure{
				SetDisplay: &FlexibleExpressionListOrComprehension{
					FlexibleExpressionList: &FlexibleExpressionList{
						FlexibleExpressions: []FlexibleExpression{
							{
								AssignmentExpression: &AssignmentExpression{
									Identifier: &tk[1],
									Expression: Expression{
										ConditionalExpression: WrapConditional(&Atom{
											Identifier: &tk[3],
											Tokens:     tk[3:4],
										}),
										Tokens: tk[3:4],
									},
									Tokens: tk[1:4],
								},
								Tokens: tk[1:4],
							},
						},
						Tokens: tk[1:4],
					},
					Tokens: tk[1:4],
				},
				Tokens: tk[:5],
			}
		}},
		{`{ a:=b }`, func(t *test, tk Tokens) { // 20
			t.Output = Enclosure{
				SetDisplay: &FlexibleExpressionListOrComprehension{
					FlexibleExpressionList: &FlexibleExpressionList{
						FlexibleExpressions: []FlexibleExpression{
							{
								AssignmentExpression: &AssignmentExpression{
									Identifier: &tk[2],
									Expression: Expression{
										ConditionalExpression: WrapConditional(&Atom{
											Identifier: &tk[4],
											Tokens:     tk[4:5],
										}),
										Tokens: tk[4:5],
									},
									Tokens: tk[2:5],
								},
								Tokens: tk[2:5],
							},
						},
						Tokens: tk[2:5],
					},
					Tokens: tk[2:5],
				},
				Tokens: tk[:7],
			}
		}},
		{`{a:b}`, func(t *test, tk Tokens) { // 21
			t.Output = Enclosure{
				DictDisplay: &DictDisplay{
					DictItems: []DictItem{
						{
							Key: &Expression{
								ConditionalExpression: WrapConditional(&Atom{
									Identifier: &tk[1],
									Tokens:     tk[1:2],
								}),
								Tokens: tk[1:2],
							},
							Value: &Expression{
								ConditionalExpression: WrapConditional(&Atom{
									Identifier: &tk[3],
									Tokens:     tk[3:4],
								}),
								Tokens: tk[3:4],
							},
							Tokens: tk[1:4],
						},
					},
					Tokens: tk[1:4],
				},
				Tokens: tk[:5],
			}
		}},
		{`{ a:b }`, func(t *test, tk Tokens) { // 22
			t.Output = Enclosure{
				DictDisplay: &DictDisplay{
					DictItems: []DictItem{
						{
							Key: &Expression{
								ConditionalExpression: WrapConditional(&Atom{
									Identifier: &tk[2],
									Tokens:     tk[2:3],
								}),
								Tokens: tk[2:3],
							},
							Value: &Expression{
								ConditionalExpression: WrapConditional(&Atom{
									Identifier: &tk[4],
									Tokens:     tk[4:5],
								}),
								Tokens: tk[4:5],
							},
							Tokens: tk[2:5],
						},
					},
					Tokens: tk[2:5],
				},
				Tokens: tk[:7],
			}
		}},
		{`(yield nonlocal)`, func(t *test, tk Tokens) { // 23
			t.Err = Error{
				Err: Error{
					Err: Error{
						Err: Error{
							Err: wrapConditionalExpressionError(Error{
								Err:     ErrInvalidEnclosure,
								Parsing: "Enclosure",
								Token:   tk[3],
							}),
							Parsing: "Expression",
							Token:   tk[3],
						},
						Parsing: "ExpressionList",
						Token:   tk[3],
					},
					Parsing: "YieldExpression",
					Token:   tk[3],
				},
				Parsing: "Enclosure",
				Token:   tk[1],
			}
		}},
		{`(nonlocal for)`, func(t *test, tk Tokens) { // 24
			t.Err = Error{
				Err: Error{
					Err: Error{
						Err: wrapConditionalExpressionError(Error{
							Err:     ErrInvalidEnclosure,
							Parsing: "Enclosure",
							Token:   tk[1],
						}),
						Parsing: "Expression",
						Token:   tk[1],
					},
					Parsing: "GeneratorExpression",
					Token:   tk[1],
				},
				Parsing: "Enclosure",
				Token:   tk[1],
			}
		}},
		{`(nonlocal)`, func(t *test, tk Tokens) { // 25
			t.Err = Error{
				Err: Error{
					Err: Error{
						Err: wrapConditionalExpressionError(Error{
							Err:     ErrInvalidEnclosure,
							Parsing: "Enclosure",
							Token:   tk[1],
						}),
						Parsing: "Expression",
						Token:   tk[1],
					},
					Parsing: "StarredExpression",
					Token:   tk[1],
				},
				Parsing: "Enclosure",
				Token:   tk[1],
			}
		}},
		{`(a b)`, func(t *test, tk Tokens) { // 26
			t.Err = Error{
				Err:     ErrMissingClosingParen,
				Parsing: "Enclosure",
				Token:   tk[3],
			}
		}},
		{`[nonlocal]`, func(t *test, tk Tokens) { // 27
			t.Err = Error{
				Err: Error{
					Err: Error{
						Err: Error{
							Err: Error{
								Err: Error{
									Err: wrapConditionalExpressionError(Error{
										Err:     ErrInvalidEnclosure,
										Parsing: "Enclosure",
										Token:   tk[1],
									}),
									Parsing: "Expression",
									Token:   tk[1],
								},
								Parsing: "AssignmentExpression",
								Token:   tk[1],
							},
							Parsing: "FlexibleExpression",
							Token:   tk[1],
						},
						Parsing: "FlexibleExpressionList",
						Token:   tk[1],
					},
					Parsing: "FlexibleExpressionListOrComprehension",
					Token:   tk[1],
				},
				Parsing: "Enclosure",
				Token:   tk[1],
			}
		}},
		{`[a for b in c d]`, func(t *test, tk Tokens) { // 28
			t.Err = Error{
				Err:     ErrMissingClosingBracket,
				Parsing: "Enclosure",
				Token:   tk[11],
			}
		}},
		{`{nonlocal}`, func(t *test, tk Tokens) { // 29
			t.Err = Error{
				Err: Error{
					Err: Error{
						Err: wrapConditionalExpressionError(Error{
							Err:     ErrInvalidEnclosure,
							Parsing: "Enclosure",
							Token:   tk[1],
						}),
						Parsing: "Expression",
						Token:   tk[1],
					},
					Parsing: "AssignmentExpression",
					Token:   tk[1],
				},
				Parsing: "Enclosure",
				Token:   tk[1],
			}
		}},
		{`{a:nonlocal}`, func(t *test, tk Tokens) { // 30
			t.Err = Error{
				Err: Error{
					Err: Error{
						Err: Error{
							Err: wrapConditionalExpressionError(Error{
								Err:     ErrInvalidEnclosure,
								Parsing: "Enclosure",
								Token:   tk[3],
							}),
							Parsing: "Expression",
							Token:   tk[3],
						},
						Parsing: "DictItem",
						Token:   tk[3],
					},
					Parsing: "DictDisplay",
					Token:   tk[1],
				},
				Parsing: "Enclosure",
				Token:   tk[1],
			}
		}},
		{`{*nonlocal}`, func(t *test, tk Tokens) { // 31
			t.Err = Error{
				Err: Error{
					Err: Error{
						Err: Error{
							Err: Error{
								Err: Error{
									Err: Error{
										Err: Error{
											Err: Error{
												Err: Error{
													Err: Error{
														Err: Error{
															Err: Error{
																Err: Error{
																	Err: Error{
																		Err: Error{
																			Err: Error{
																				Err: Error{
																					Err:     ErrInvalidEnclosure,
																					Parsing: "Enclosure",
																					Token:   tk[2],
																				},
																				Parsing: "Atom",
																				Token:   tk[2],
																			},
																			Parsing: "PrimaryExpression",
																			Token:   tk[2],
																		},
																		Parsing: "PowerExpression",
																		Token:   tk[2],
																	},
																	Parsing: "UnaryExpression",
																	Token:   tk[2],
																},
																Parsing: "MultiplyExpression",
																Token:   tk[2],
															},
															Parsing: "AddExpression",
															Token:   tk[2],
														},
														Parsing: "ShiftExpression",
														Token:   tk[2],
													},
													Parsing: "AndExpression",
													Token:   tk[2],
												},
												Parsing: "XorExpression",
												Token:   tk[2],
											},
											Parsing: "OrExpression",
											Token:   tk[2],
										},
										Parsing: "StarredItem",
										Token:   tk[2],
									},
									Parsing: "StarredList",
									Token:   tk[1],
								},
								Parsing: "StarredExpression",
								Token:   tk[1],
							},
							Parsing: "FlexibleExpression",
							Token:   tk[1],
						},
						Parsing: "FlexibleExpressionList",
						Token:   tk[1],
					},
					Parsing: "FlexibleExpressionListOrComprehension",
					Token:   tk[1],
				},
				Parsing: "Enclosure",
				Token:   tk[1],
			}
		}},
		{`{a for b in c d}`, func(t *test, tk Tokens) { // 32
			t.Err = Error{
				Err:     ErrMissingClosingBrace,
				Parsing: "Enclosure",
				Token:   tk[11],
			}
		}},
		{"(#abc\n)", func(t *test, tk Tokens) { // 33
			t.Output = Enclosure{
				ParenthForm: &StarredExpression{
					Tokens: tk[1:1],
				},
				Comments: [2]Comments{{tk[1]}},
				Tokens:   tk[:4],
			}
		}},
		{"[#abc\n]", func(t *test, tk Tokens) { // 34
			t.Output = Enclosure{
				ListDisplay: &FlexibleExpressionListOrComprehension{
					Tokens: tk[1:1],
				},
				Comments: [2]Comments{{tk[1]}},
				Tokens:   tk[:4],
			}
		}},
		{"{#abc\n}", func(t *test, tk Tokens) { // 35
			t.Output = Enclosure{
				DictDisplay: &DictDisplay{
					Tokens: tk[1:1],
				},
				Comments: [2]Comments{{tk[1]}},
				Tokens:   tk[:4],
			}
		}},
		{"(#abc\na\n#def\n)", func(t *test, tk Tokens) { // 36
			t.Output = Enclosure{
				ParenthForm: &StarredExpression{
					Expression: &Expression{
						ConditionalExpression: WrapConditional(&Atom{
							Identifier: &tk[3],
							Tokens:     tk[3:4],
						}),
						Tokens: tk[3:4],
					},
					Tokens: tk[3:4],
				},
				Comments: [2]Comments{{tk[1]}, {tk[5]}},
				Tokens:   tk[:8],
			}
		}},
		{"[#abc\na\n#def\n]", func(t *test, tk Tokens) { // 37
			t.Output = Enclosure{
				ListDisplay: &FlexibleExpressionListOrComprehension{
					FlexibleExpressionList: &FlexibleExpressionList{
						FlexibleExpressions: []FlexibleExpression{
							{
								AssignmentExpression: &AssignmentExpression{
									Expression: Expression{
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
						},
						Tokens: tk[3:4],
					},
					Tokens: tk[3:4],
				},
				Comments: [2]Comments{{tk[1]}, {tk[5]}},
				Tokens:   tk[:8],
			}
		}},
		{"{#abc\n**a\n#def\n}", func(t *test, tk Tokens) { // 38
			t.Output = Enclosure{
				DictDisplay: &DictDisplay{
					DictItems: []DictItem{
						{
							OrExpression: &WrapConditional(&Atom{
								Identifier: &tk[4],
								Tokens:     tk[4:5],
							}).OrTest.AndTest.NotTest.Comparison.OrExpression,
							Tokens: tk[3:5],
						},
					},
					Tokens: tk[3:5],
				},
				Comments: [2]Comments{{tk[1]}, {tk[6]}},
				Tokens:   tk[:9],
			}
		}},
		{"{#abc\n*a\n#def\n}", func(t *test, tk Tokens) { // 39
			t.Output = Enclosure{
				SetDisplay: &FlexibleExpressionListOrComprehension{
					FlexibleExpressionList: &FlexibleExpressionList{
						FlexibleExpressions: []FlexibleExpression{
							{
								StarredExpression: &StarredExpression{
									StarredList: &StarredList{
										StarredItems: []StarredItem{
											{
												OrExpr: &WrapConditional(&Atom{
													Identifier: &tk[4],
													Tokens:     tk[4:5],
												}).OrTest.AndTest.NotTest.Comparison.OrExpression,
												Tokens: tk[3:5],
											},
										},
										Tokens: tk[3:5],
									},
									Tokens: tk[3:5],
								},
								Tokens: tk[3:5],
							},
						},
						Tokens: tk[3:5],
					},
					Tokens: tk[3:5],
				},
				Comments: [2]Comments{{tk[1]}, {tk[6]}},
				Tokens:   tk[:9],
			}
		}},
		{"{#abc\na:b\n#def\n}", func(t *test, tk Tokens) { // 40
			t.Output = Enclosure{
				DictDisplay: &DictDisplay{
					DictItems: []DictItem{
						{
							Key: &Expression{
								ConditionalExpression: WrapConditional(&Atom{
									Identifier: &tk[3],
									Tokens:     tk[3:4],
								}),
								Tokens: tk[3:4],
							},
							Value: &Expression{
								ConditionalExpression: WrapConditional(&Atom{
									Identifier: &tk[5],
									Tokens:     tk[5:6],
								}),
								Tokens: tk[5:6],
							},
							Tokens: tk[3:6],
						},
					},
					Tokens: tk[3:6],
				},
				Comments: [2]Comments{{tk[1]}, {tk[7]}},
				Tokens:   tk[:10],
			}
		}},
	}, func(t *test) (Type, error) {
		var e Enclosure

		err := e.parse(t.Tokens)

		return e, err
	})
}

func TestFlexibleExpressionListOrComprehension(t *testing.T) {
	doTests(t, []sourceFn{
		{`a`, func(t *test, tk Tokens) { // 1
			t.Output = FlexibleExpressionListOrComprehension{
				FlexibleExpressionList: &FlexibleExpressionList{
					FlexibleExpressions: []FlexibleExpression{
						{
							AssignmentExpression: &AssignmentExpression{
								Expression: Expression{
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
					},
					Tokens: tk[:1],
				},
				Tokens: tk[:1],
			}
		}},
		{`a for b in c`, func(t *test, tk Tokens) { // 2
			t.Output = FlexibleExpressionListOrComprehension{
				Comprehension: &Comprehension{
					AssignmentExpression: AssignmentExpression{
						Expression: Expression{
							ConditionalExpression: WrapConditional(&Atom{
								Identifier: &tk[0],
								Tokens:     tk[:1],
							}),
							Tokens: tk[:1],
						},
						Tokens: tk[:1],
					},
					ComprehensionFor: ComprehensionFor{
						TargetList: TargetList{
							Targets: []Target{
								{
									PrimaryExpression: &PrimaryExpression{
										Atom: &Atom{
											Identifier: &tk[4],
											Tokens:     tk[4:5],
										},
										Tokens: tk[4:5],
									},
									Tokens: tk[4:5],
								},
							},
							Tokens: tk[4:5],
						},
						OrTest: WrapConditional(&Atom{
							Identifier: &tk[8],
							Tokens:     tk[8:9],
						}).OrTest,
						Tokens: tk[2:9],
					},
					Tokens: tk[:9],
				},
				Tokens: tk[:9],
			}
		}},
		{`a async for b in c`, func(t *test, tk Tokens) { // 3
			t.Output = FlexibleExpressionListOrComprehension{
				Comprehension: &Comprehension{
					AssignmentExpression: AssignmentExpression{
						Expression: Expression{
							ConditionalExpression: WrapConditional(&Atom{
								Identifier: &tk[0],
								Tokens:     tk[:1],
							}),
							Tokens: tk[:1],
						},
						Tokens: tk[:1],
					},
					ComprehensionFor: ComprehensionFor{
						Async: true,
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
						OrTest: WrapConditional(&Atom{
							Identifier: &tk[10],
							Tokens:     tk[10:11],
						}).OrTest,
						Tokens: tk[2:11],
					},
					Tokens: tk[:11],
				},
				Tokens: tk[:11],
			}
		}},
		{`a`, func(t *test, tk Tokens) { // 4
			t.AssignmentExpression = &AssignmentExpression{
				Expression: Expression{
					ConditionalExpression: WrapConditional(&Atom{
						Identifier: &tk[0],
						Tokens:     tk[:1],
					}),
					Tokens: tk[:1],
				},
				Tokens: tk[:1],
			}
			t.TokenSkip = 1
			t.Output = FlexibleExpressionListOrComprehension{
				FlexibleExpressionList: &FlexibleExpressionList{
					FlexibleExpressions: []FlexibleExpression{
						{
							AssignmentExpression: &AssignmentExpression{
								Expression: Expression{
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
					},
					Tokens: tk[:1],
				},
				Tokens: tk[:1],
			}
		}},
		{`a for b in c`, func(t *test, tk Tokens) { // 5
			t.AssignmentExpression = &AssignmentExpression{
				Expression: Expression{
					ConditionalExpression: WrapConditional(&Atom{
						Identifier: &tk[0],
						Tokens:     tk[:1],
					}),
					Tokens: tk[:1],
				},
				Tokens: tk[:1],
			}
			t.TokenSkip = 1
			t.Output = FlexibleExpressionListOrComprehension{
				Comprehension: &Comprehension{
					AssignmentExpression: AssignmentExpression{
						Expression: Expression{
							ConditionalExpression: WrapConditional(&Atom{
								Identifier: &tk[0],
								Tokens:     tk[:1],
							}),
							Tokens: tk[:1],
						},
						Tokens: tk[:1],
					},
					ComprehensionFor: ComprehensionFor{
						TargetList: TargetList{
							Targets: []Target{
								{
									PrimaryExpression: &PrimaryExpression{
										Atom: &Atom{
											Identifier: &tk[4],
											Tokens:     tk[4:5],
										},
										Tokens: tk[4:5],
									},
									Tokens: tk[4:5],
								},
							},
							Tokens: tk[4:5],
						},
						OrTest: WrapConditional(&Atom{
							Identifier: &tk[8],
							Tokens:     tk[8:9],
						}).OrTest,
						Tokens: tk[2:9],
					},
					Tokens: tk[:9],
				},
				Tokens: tk[:9],
			}
		}},
		{`a async for b in c`, func(t *test, tk Tokens) { // 6
			t.AssignmentExpression = &AssignmentExpression{
				Expression: Expression{
					ConditionalExpression: WrapConditional(&Atom{
						Identifier: &tk[0],
						Tokens:     tk[:1],
					}),
					Tokens: tk[:1],
				},
				Tokens: tk[:1],
			}
			t.TokenSkip = 1
			t.Output = FlexibleExpressionListOrComprehension{
				Comprehension: &Comprehension{
					AssignmentExpression: AssignmentExpression{
						Expression: Expression{
							ConditionalExpression: WrapConditional(&Atom{
								Identifier: &tk[0],
								Tokens:     tk[:1],
							}),
							Tokens: tk[:1],
						},
						Tokens: tk[:1],
					},
					ComprehensionFor: ComprehensionFor{
						Async: true,
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
						OrTest: WrapConditional(&Atom{
							Identifier: &tk[10],
							Tokens:     tk[10:11],
						}).OrTest,
						Tokens: tk[2:11],
					},
					Tokens: tk[:11],
				},
				Tokens: tk[:11],
			}
		}},
		{`nonlocal`, func(t *test, tk Tokens) { // 7
			t.Err = Error{
				Err: Error{
					Err: Error{
						Err: Error{
							Err: Error{
								Err: wrapConditionalExpressionError(Error{
									Err:     ErrInvalidEnclosure,
									Parsing: "Enclosure",
									Token:   tk[0],
								}),
								Parsing: "Expression",
								Token:   tk[0],
							},
							Parsing: "AssignmentExpression",
							Token:   tk[0],
						},
						Parsing: "FlexibleExpression",
						Token:   tk[0],
					},
					Parsing: "FlexibleExpressionList",
					Token:   tk[0],
				},
				Parsing: "FlexibleExpressionListOrComprehension",
				Token:   tk[0],
			}
		}},
		{`a for nonlocal in b`, func(t *test, tk Tokens) { // 8
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
								Parsing: "Target",
								Token:   tk[4],
							},
							Parsing: "TargetList",
							Token:   tk[4],
						},
						Parsing: "ComprehensionFor",
						Token:   tk[4],
					},
					Parsing: "Comprehension",
					Token:   tk[2],
				},
				Parsing: "FlexibleExpressionListOrComprehension",
				Token:   tk[0],
			}
		}},
	}, func(t *test) (Type, error) {
		var f FlexibleExpressionListOrComprehension

		err := f.parse(t.Tokens, t.AssignmentExpression)

		return f, err
	})
}

func TestFlexibleExpressionList(t *testing.T) {
	doTests(t, []sourceFn{
		{`a`, func(t *test, tk Tokens) { // 1
			t.Output = FlexibleExpressionList{
				FlexibleExpressions: []FlexibleExpression{
					{
						AssignmentExpression: &AssignmentExpression{
							Expression: Expression{
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
				},
				Tokens: tk[:1],
			}
		}},
		{`a,b`, func(t *test, tk Tokens) { // 2
			t.Output = FlexibleExpressionList{
				FlexibleExpressions: []FlexibleExpression{
					{
						AssignmentExpression: &AssignmentExpression{
							Expression: Expression{
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
					{
						AssignmentExpression: &AssignmentExpression{
							Expression: Expression{
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
				},
				Tokens: tk[:3],
			}
		}},
		{`a, b, c`, func(t *test, tk Tokens) { // 3
			t.Output = FlexibleExpressionList{
				FlexibleExpressions: []FlexibleExpression{
					{
						AssignmentExpression: &AssignmentExpression{
							Expression: Expression{
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
					{
						AssignmentExpression: &AssignmentExpression{
							Expression: Expression{
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
					{
						AssignmentExpression: &AssignmentExpression{
							Expression: Expression{
								ConditionalExpression: WrapConditional(&Atom{
									Identifier: &tk[6],
									Tokens:     tk[6:7],
								}),
								Tokens: tk[6:7],
							},
							Tokens: tk[6:7],
						},
						Tokens: tk[6:7],
					},
				},
				Tokens: tk[:7],
			}
		}},
		{`a,`, func(t *test, tk Tokens) { // 4
			t.Output = FlexibleExpressionList{
				FlexibleExpressions: []FlexibleExpression{
					{
						AssignmentExpression: &AssignmentExpression{
							Expression: Expression{
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
				},
				Tokens: tk[:2],
			}
		}},
		{`nonlocal`, func(t *test, tk Tokens) { // 5
			t.Err = Error{
				Err: Error{
					Err: Error{
						Err: Error{
							Err: wrapConditionalExpressionError(Error{
								Err:     ErrInvalidEnclosure,
								Parsing: "Enclosure",
								Token:   tk[0],
							}),
							Parsing: "Expression",
							Token:   tk[0],
						},
						Parsing: "AssignmentExpression",
						Token:   tk[0],
					},
					Parsing: "FlexibleExpression",
					Token:   tk[0],
				},
				Parsing: "FlexibleExpressionList",
				Token:   tk[0],
			}
		}},
		{`a,nonlocal`, func(t *test, tk Tokens) { // 6
			t.Err = Error{
				Err: Error{
					Err: Error{
						Err: Error{
							Err: wrapConditionalExpressionError(Error{
								Err:     ErrInvalidEnclosure,
								Parsing: "Enclosure",
								Token:   tk[2],
							}),
							Parsing: "Expression",
							Token:   tk[2],
						},
						Parsing: "AssignmentExpression",
						Token:   tk[2],
					},
					Parsing: "FlexibleExpression",
					Token:   tk[2],
				},
				Parsing: "FlexibleExpressionList",
				Token:   tk[2],
			}
		}},
	}, func(t *test) (Type, error) {
		var f FlexibleExpressionList

		err := f.parse(t.Tokens)

		return f, err
	})
}

func TestFlexibleExpression(t *testing.T) {
	doTests(t, []sourceFn{
		{`a`, func(t *test, tk Tokens) { // 1
			t.Output = FlexibleExpression{
				AssignmentExpression: &AssignmentExpression{
					Expression: Expression{
						ConditionalExpression: WrapConditional(&Atom{
							Identifier: &tk[0],
							Tokens:     tk[:1],
						}),
						Tokens: tk[:1],
					},
					Tokens: tk[:1],
				},
				Tokens: tk[:1],
			}
		}},
		{`*a`, func(t *test, tk Tokens) { // 2
			t.Output = FlexibleExpression{
				StarredExpression: &StarredExpression{
					StarredList: &StarredList{
						StarredItems: []StarredItem{
							{
								OrExpr: &WrapConditional(&Atom{
									Identifier: &tk[1],
									Tokens:     tk[1:2],
								}).OrTest.AndTest.NotTest.Comparison.OrExpression,
								Tokens: tk[:2],
							},
						},
						Tokens: tk[:2],
					},
					Tokens: tk[:2],
				},
				Tokens: tk[:2],
			}
		}},
		{`nonlocal`, func(t *test, tk Tokens) { // 3
			t.Err = Error{
				Err: Error{
					Err: Error{
						Err: wrapConditionalExpressionError(Error{
							Err:     ErrInvalidEnclosure,
							Parsing: "Enclosure",
							Token:   tk[0],
						}),
						Parsing: "Expression",
						Token:   tk[0],
					},
					Parsing: "AssignmentExpression",
					Token:   tk[0],
				},
				Parsing: "FlexibleExpression",
				Token:   tk[0],
			}
		}},
		{`*nonlocal`, func(t *test, tk Tokens) { // 4
			t.Err = Error{
				Err: Error{
					Err: Error{
						Err: Error{
							Err: Error{
								Err: Error{
									Err: Error{
										Err: Error{
											Err: Error{
												Err: Error{
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
													Parsing: "MultiplyExpression",
													Token:   tk[1],
												},
												Parsing: "AddExpression",
												Token:   tk[1],
											},
											Parsing: "ShiftExpression",
											Token:   tk[1],
										},
										Parsing: "AndExpression",
										Token:   tk[1],
									},
									Parsing: "XorExpression",
									Token:   tk[1],
								},
								Parsing: "OrExpression",
								Token:   tk[1],
							},
							Parsing: "StarredItem",
							Token:   tk[1],
						},
						Parsing: "StarredList",
						Token:   tk[0],
					},
					Parsing: "StarredExpression",
					Token:   tk[0],
				},
				Parsing: "FlexibleExpression",
				Token:   tk[0],
			}
		}},
	}, func(t *test) (Type, error) {
		var f FlexibleExpression

		err := f.parse(t.Tokens)

		return f, err
	})
}

func TestComprehension(t *testing.T) {
	doTests(t, []sourceFn{
		{`a for b in c`, func(t *test, tk Tokens) { // 1
			t.Output = Comprehension{
				AssignmentExpression: AssignmentExpression{
					Expression: Expression{
						ConditionalExpression: WrapConditional(&Atom{
							Identifier: &tk[0],
							Tokens:     tk[:1],
						}),
						Tokens: tk[:1],
					},
					Tokens: tk[:1],
				},
				ComprehensionFor: ComprehensionFor{
					TargetList: TargetList{
						Targets: []Target{
							{
								PrimaryExpression: &PrimaryExpression{
									Atom: &Atom{
										Identifier: &tk[4],
										Tokens:     tk[4:5],
									},
									Tokens: tk[4:5],
								},
								Tokens: tk[4:5],
							},
						},
						Tokens: tk[4:5],
					},
					OrTest: WrapConditional(&Atom{
						Identifier: &tk[8],
						Tokens:     tk[8:9],
					}).OrTest,
					Tokens: tk[2:9],
				},
				Tokens: tk[:9],
			}
		}},
		{`a for b in c`, func(t *test, tk Tokens) { // 2
			t.AssignmentExpression = &AssignmentExpression{
				Expression: Expression{
					ConditionalExpression: WrapConditional(&Atom{
						Identifier: &tk[0],
						Tokens:     tk[:1],
					}),
					Tokens: tk[:1],
				},
				Tokens: tk[:1],
			}
			t.TokenSkip = 1
			t.Output = Comprehension{
				AssignmentExpression: AssignmentExpression{
					Expression: Expression{
						ConditionalExpression: WrapConditional(&Atom{
							Identifier: &tk[0],
							Tokens:     tk[:1],
						}),
						Tokens: tk[:1],
					},
					Tokens: tk[:1],
				},
				ComprehensionFor: ComprehensionFor{
					TargetList: TargetList{
						Targets: []Target{
							{
								PrimaryExpression: &PrimaryExpression{
									Atom: &Atom{
										Identifier: &tk[4],
										Tokens:     tk[4:5],
									},
									Tokens: tk[4:5],
								},
								Tokens: tk[4:5],
							},
						},
						Tokens: tk[4:5],
					},
					OrTest: WrapConditional(&Atom{
						Identifier: &tk[8],
						Tokens:     tk[8:9],
					}).OrTest,
					Tokens: tk[2:9],
				},
				Tokens: tk[:9],
			}
		}},
		{`nonlocal for b in c`, func(t *test, tk Tokens) { // 3
			t.Err = Error{
				Err: Error{
					Err: Error{
						Err: wrapConditionalExpressionError(Error{
							Err:     ErrInvalidEnclosure,
							Parsing: "Enclosure",
							Token:   tk[0],
						}),
						Parsing: "Expression",
						Token:   tk[0],
					},
					Parsing: "AssignmentExpression",
					Token:   tk[0],
				},
				Parsing: "Comprehension",
				Token:   tk[0],
			}
		}},
		{`a for nonlocal in c`, func(t *test, tk Tokens) { // 4
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
							Parsing: "Target",
							Token:   tk[4],
						},
						Parsing: "TargetList",
						Token:   tk[4],
					},
					Parsing: "ComprehensionFor",
					Token:   tk[4],
				},
				Parsing: "Comprehension",
				Token:   tk[2],
			}
		}},
	}, func(t *test) (Type, error) {
		var c Comprehension

		err := c.parse(t.Tokens, t.AssignmentExpression)

		return c, err
	})
}

func TestComprehensionFor(t *testing.T) {
	doTests(t, []sourceFn{
		{`for a in b`, func(t *test, tk Tokens) { // 1
			t.Output = ComprehensionFor{
				TargetList: TargetList{
					Targets: []Target{
						{
							PrimaryExpression: &PrimaryExpression{
								Atom: &Atom{
									Identifier: &tk[2],
									Tokens:     tk[2:3],
								},
								Tokens: tk[2:3],
							},
							Tokens: tk[2:3],
						},
					},
					Tokens: tk[2:3],
				},
				OrTest: WrapConditional(&Atom{
					Identifier: &tk[6],
					Tokens:     tk[6:7],
				}).OrTest,
				Tokens: tk[:7],
			}
		}},
		{`async for a in b`, func(t *test, tk Tokens) { // 2
			t.Output = ComprehensionFor{
				Async: true,
				TargetList: TargetList{
					Targets: []Target{
						{
							PrimaryExpression: &PrimaryExpression{
								Atom: &Atom{
									Identifier: &tk[4],
									Tokens:     tk[4:5],
								},
								Tokens: tk[4:5],
							},
							Tokens: tk[4:5],
						},
					},
					Tokens: tk[4:5],
				},
				OrTest: WrapConditional(&Atom{
					Identifier: &tk[8],
					Tokens:     tk[8:9],
				}).OrTest,
				Tokens: tk[:9],
			}
		}},
		{`for a in b if c`, func(t *test, tk Tokens) { // 3
			t.Output = ComprehensionFor{
				TargetList: TargetList{
					Targets: []Target{
						{
							PrimaryExpression: &PrimaryExpression{
								Atom: &Atom{
									Identifier: &tk[2],
									Tokens:     tk[2:3],
								},
								Tokens: tk[2:3],
							},
							Tokens: tk[2:3],
						},
					},
					Tokens: tk[2:3],
				},
				OrTest: WrapConditional(&Atom{
					Identifier: &tk[6],
					Tokens:     tk[6:7],
				}).OrTest,
				ComprehensionIterator: &ComprehensionIterator{
					ComprehensionIf: &ComprehensionIf{
						OrTest: WrapConditional(&Atom{
							Identifier: &tk[10],
							Tokens:     tk[10:11],
						}).OrTest,
						Tokens: tk[8:11],
					},
					Tokens: tk[8:11],
				},
				Tokens: tk[:11],
			}
		}},
		{`async for a in b for c in d`, func(t *test, tk Tokens) { // 4
			t.Output = ComprehensionFor{
				Async: true,
				TargetList: TargetList{
					Targets: []Target{
						{
							PrimaryExpression: &PrimaryExpression{
								Atom: &Atom{
									Identifier: &tk[4],
									Tokens:     tk[4:5],
								},
								Tokens: tk[4:5],
							},
							Tokens: tk[4:5],
						},
					},
					Tokens: tk[4:5],
				},
				OrTest: WrapConditional(&Atom{
					Identifier: &tk[8],
					Tokens:     tk[8:9],
				}).OrTest,
				ComprehensionIterator: &ComprehensionIterator{
					ComprehensionFor: &ComprehensionFor{
						TargetList: TargetList{
							Targets: []Target{
								{
									PrimaryExpression: &PrimaryExpression{
										Atom: &Atom{
											Identifier: &tk[12],
											Tokens:     tk[12:13],
										},
										Tokens: tk[12:13],
									},
									Tokens: tk[12:13],
								},
							},
							Tokens: tk[12:13],
						},
						OrTest: WrapConditional(&Atom{
							Identifier: &tk[16],
							Tokens:     tk[16:17],
						}).OrTest,
						Tokens: tk[10:17],
					},
					Tokens: tk[10:17],
				},
				Tokens: tk[:17],
			}
		}},
		{`for a in b async for c in d`, func(t *test, tk Tokens) { // 5
			t.Output = ComprehensionFor{
				TargetList: TargetList{
					Targets: []Target{
						{
							PrimaryExpression: &PrimaryExpression{
								Atom: &Atom{
									Identifier: &tk[2],
									Tokens:     tk[2:3],
								},
								Tokens: tk[2:3],
							},
							Tokens: tk[2:3],
						},
					},
					Tokens: tk[2:3],
				},
				OrTest: WrapConditional(&Atom{
					Identifier: &tk[6],
					Tokens:     tk[6:7],
				}).OrTest,
				ComprehensionIterator: &ComprehensionIterator{
					ComprehensionFor: &ComprehensionFor{
						Async: true,
						TargetList: TargetList{
							Targets: []Target{
								{
									PrimaryExpression: &PrimaryExpression{
										Atom: &Atom{
											Identifier: &tk[12],
											Tokens:     tk[12:13],
										},
										Tokens: tk[12:13],
									},
									Tokens: tk[12:13],
								},
							},
							Tokens: tk[12:13],
						},
						OrTest: WrapConditional(&Atom{
							Identifier: &tk[16],
							Tokens:     tk[16:17],
						}).OrTest,
						Tokens: tk[8:17],
					},
					Tokens: tk[8:17],
				},
				Tokens: tk[:17],
			}
		}},
		{"(# A\nfor # B\na #C\nin # D\nb # E\n)", func(t *test, tk Tokens) { // 6
			t.Output = ComprehensionFor{
				TargetList: TargetList{
					Targets: []Target{
						{
							PrimaryExpression: &PrimaryExpression{
								Atom: &Atom{
									Identifier: &tk[7],
									Tokens:     tk[7:8],
								},
								Tokens: tk[7:8],
							},
							Tokens: tk[7:8],
						},
					},
					Comments: [2]Comments{nil, {tk[9]}},
					Tokens:   tk[7:10],
				},
				OrTest: WrapConditional(&Atom{
					Identifier: &tk[15],
					Tokens:     tk[15:16],
				}).OrTest,
				Comments: [5]Comments{{tk[1]}, nil, {tk[5]}, {tk[13]}, {tk[17]}},
				Tokens:   tk[1:18],
			}
		}},
		{"(# A\nasync # B\nfor # C\na # D\nin # E\nb # F\n)", func(t *test, tk Tokens) { // 7
			t.Output = ComprehensionFor{
				Async: true,
				TargetList: TargetList{
					Targets: []Target{
						{
							PrimaryExpression: &PrimaryExpression{
								Atom: &Atom{
									Identifier: &tk[11],
									Tokens:     tk[11:12],
								},
								Tokens: tk[11:12],
							},
							Tokens: tk[11:12],
						},
					},
					Comments: [2]Comments{nil, {tk[13]}},
					Tokens:   tk[11:14],
				},
				OrTest: WrapConditional(&Atom{
					Identifier: &tk[19],
					Tokens:     tk[19:20],
				}).OrTest,
				Comments: [5]Comments{{tk[1]}, {tk[5]}, {tk[9]}, {tk[17]}, {tk[21]}},
				Tokens:   tk[1:22],
			}
		}},
		{`async a in b if c`, func(t *test, tk Tokens) { // 8
			t.Err = Error{
				Err:     ErrMissingFor,
				Parsing: "ComprehensionFor",
				Token:   tk[2],
			}
		}},
		{`for nonlocal in a if b`, func(t *test, tk Tokens) { // 9
			t.Err = Error{
				Err: Error{
					Err: Error{
						Err: Error{
							Err: Error{
								Err: Error{
									Err:     ErrInvalidEnclosure,
									Parsing: "Enclosure",
									Token:   tk[2],
								},
								Parsing: "Atom",
								Token:   tk[2],
							},
							Parsing: "PrimaryExpression",
							Token:   tk[2],
						},
						Parsing: "Target",
						Token:   tk[2],
					},
					Parsing: "TargetList",
					Token:   tk[2],
				},
				Parsing: "ComprehensionFor",
				Token:   tk[2],
			}
		}},
		{`for a in nonlocal if b`, func(t *test, tk Tokens) { // 10
			t.Err = Error{
				Err: wrapConditionalExpressionError(Error{
					Err:     ErrInvalidEnclosure,
					Parsing: "Enclosure",
					Token:   tk[6],
				}).Err,
				Parsing: "ComprehensionFor",
				Token:   tk[6],
			}
		}},
		{`for a in b if nonlocal`, func(t *test, tk Tokens) { // 11
			t.Err = Error{
				Err: Error{
					Err: Error{
						Err: wrapConditionalExpressionError(Error{
							Err:     ErrInvalidEnclosure,
							Parsing: "Enclosure",
							Token:   tk[10],
						}).Err,
						Parsing: "ComprehensionIf",
						Token:   tk[10],
					},
					Parsing: "ComprehensionIterator",
					Token:   tk[8],
				},
				Parsing: "ComprehensionFor",
				Token:   tk[8],
			}
		}},
		{`for a`, func(t *test, tk Tokens) { // 12
			t.Err = Error{
				Err:     ErrMissingIn,
				Parsing: "ComprehensionFor",
				Token:   tk[3],
			}
		}},
	}, func(t *test) (Type, error) {
		var c ComprehensionFor

		if t.Tokens.Peek().Data == "(" {
			t.Tokens.Tokens = t.Tokens.Tokens[1:1]
		}

		err := c.parse(t.Tokens)

		return c, err
	})
}

func TestComprehensionIterator(t *testing.T) {
	doTests(t, []sourceFn{
		{`if a`, func(t *test, tk Tokens) { // 1
			t.Output = ComprehensionIterator{
				ComprehensionIf: &ComprehensionIf{
					OrTest: WrapConditional(&Atom{
						Identifier: &tk[2],
						Tokens:     tk[2:3],
					}).OrTest,
					Tokens: tk[:3],
				},
				Tokens: tk[:3],
			}
		}},
		{`for a in b`, func(t *test, tk Tokens) { // 2
			t.Output = ComprehensionIterator{
				ComprehensionFor: &ComprehensionFor{
					TargetList: TargetList{
						Targets: []Target{
							{
								PrimaryExpression: &PrimaryExpression{
									Atom: &Atom{
										Identifier: &tk[2],
										Tokens:     tk[2:3],
									},
									Tokens: tk[2:3],
								},
								Tokens: tk[2:3],
							},
						},
						Tokens: tk[2:3],
					},
					OrTest: WrapConditional(&Atom{
						Identifier: &tk[6],
						Tokens:     tk[6:7],
					}).OrTest,
					Tokens: tk[:7],
				},
				Tokens: tk[:7],
			}
		}},
		{`async for a in b`, func(t *test, tk Tokens) { // 3
			t.Output = ComprehensionIterator{
				ComprehensionFor: &ComprehensionFor{
					Async: true,
					TargetList: TargetList{
						Targets: []Target{
							{
								PrimaryExpression: &PrimaryExpression{
									Atom: &Atom{
										Identifier: &tk[4],
										Tokens:     tk[4:5],
									},
									Tokens: tk[4:5],
								},
								Tokens: tk[4:5],
							},
						},
						Tokens: tk[4:5],
					},
					OrTest: WrapConditional(&Atom{
						Identifier: &tk[8],
						Tokens:     tk[8:9],
					}).OrTest,
					Tokens: tk[:9],
				},
				Tokens: tk[:9],
			}
		}},
		{`if nonlocal`, func(t *test, tk Tokens) { // 4
			t.Err = Error{
				Err: Error{
					Err: wrapConditionalExpressionError(Error{
						Err:     ErrInvalidEnclosure,
						Parsing: "Enclosure",
						Token:   tk[2],
					}).Err,
					Parsing: "ComprehensionIf",
					Token:   tk[2],
				},
				Parsing: "ComprehensionIterator",
				Token:   tk[0],
			}
		}},
		{`for nonlocal in a`, func(t *test, tk Tokens) { // 5
			t.Err = Error{
				Err: Error{
					Err: Error{
						Err: Error{
							Err: Error{
								Err: Error{
									Err: Error{
										Err:     ErrInvalidEnclosure,
										Parsing: "Enclosure",
										Token:   tk[2],
									},
									Parsing: "Atom",
									Token:   tk[2],
								},
								Parsing: "PrimaryExpression",
								Token:   tk[2],
							},
							Parsing: "Target",
							Token:   tk[2],
						},
						Parsing: "TargetList",
						Token:   tk[2],
					},
					Parsing: "ComprehensionFor",
					Token:   tk[2],
				},
				Parsing: "ComprehensionIterator",
				Token:   tk[0],
			}
		}},
	}, func(t *test) (Type, error) {
		var c ComprehensionIterator

		err := c.parse(t.Tokens)

		return c, err
	})
}

func TestComprehensionIf(t *testing.T) {
	doTests(t, []sourceFn{
		{`if a`, func(t *test, tk Tokens) { // 1
			t.Output = ComprehensionIf{
				OrTest: WrapConditional(&Atom{
					Identifier: &tk[2],
					Tokens:     tk[2:3],
				}).OrTest,
				Tokens: tk[:3],
			}
		}},
		{`if a if b`, func(t *test, tk Tokens) { // 2
			t.Output = ComprehensionIf{
				OrTest: WrapConditional(&Atom{
					Identifier: &tk[2],
					Tokens:     tk[2:3],
				}).OrTest,

				ComprehensionIterator: &ComprehensionIterator{
					ComprehensionIf: &ComprehensionIf{
						OrTest: WrapConditional(&Atom{
							Identifier: &tk[6],
							Tokens:     tk[6:7],
						}).OrTest,
						Tokens: tk[4:7],
					},
					Tokens: tk[4:7],
				},
				Tokens: tk[:7],
			}
		}},
		{`if a for b in c`, func(t *test, tk Tokens) { // 3
			t.Output = ComprehensionIf{
				OrTest: WrapConditional(&Atom{
					Identifier: &tk[2],
					Tokens:     tk[2:3],
				}).OrTest,
				ComprehensionIterator: &ComprehensionIterator{
					ComprehensionFor: &ComprehensionFor{
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
						OrTest: WrapConditional(&Atom{
							Identifier: &tk[10],
							Tokens:     tk[10:11],
						}).OrTest,
						Tokens: tk[4:11],
					},
					Tokens: tk[4:11],
				},
				Tokens: tk[:11],
			}
		}},
		{`if a async for b in c`, func(t *test, tk Tokens) { // 4
			t.Output = ComprehensionIf{
				OrTest: WrapConditional(&Atom{
					Identifier: &tk[2],
					Tokens:     tk[2:3],
				}).OrTest,
				ComprehensionIterator: &ComprehensionIterator{
					ComprehensionFor: &ComprehensionFor{
						Async: true,
						TargetList: TargetList{
							Targets: []Target{
								{
									PrimaryExpression: &PrimaryExpression{
										Atom: &Atom{
											Identifier: &tk[8],
											Tokens:     tk[8:9],
										},
										Tokens: tk[8:9],
									},
									Tokens: tk[8:9],
								},
							},
							Tokens: tk[8:9],
						},
						OrTest: WrapConditional(&Atom{
							Identifier: &tk[12],
							Tokens:     tk[12:13],
						}).OrTest,
						Tokens: tk[4:13],
					},
					Tokens: tk[4:13],
				},
				Tokens: tk[:13],
			}
		}},
		{"(# A\nif # B\na # C\n)", func(t *test, tk Tokens) { // 5
			t.Output = ComprehensionIf{
				OrTest: WrapConditional(&Atom{
					Identifier: &tk[7],
					Tokens:     tk[7:8],
				}).OrTest,
				Comments: [3]Comments{{tk[1]}, {tk[5]}, {tk[9]}},
				Tokens:   tk[1:10],
			}
		}},
		{"(# A\nif # B\na # C\nfor b in c # D\n\n# E\n)", func(t *test, tk Tokens) { // 6
			t.Output = ComprehensionIf{
				OrTest: WrapConditional(&Atom{
					Identifier: &tk[7],
					Tokens:     tk[7:8],
				}).OrTest,
				ComprehensionIterator: &ComprehensionIterator{
					ComprehensionFor: &ComprehensionFor{
						TargetList: TargetList{
							Targets: []Target{
								{
									PrimaryExpression: &PrimaryExpression{
										Atom: &Atom{
											Identifier: &tk[13],
											Tokens:     tk[13:14],
										},
										Tokens: tk[13:14],
									},
									Tokens: tk[13:14],
								},
							},
							Tokens: tk[13:14],
						},
						OrTest: WrapConditional(&Atom{
							Identifier: &tk[17],
							Tokens:     tk[17:18],
						}).OrTest,
						Comments: [5]Comments{{tk[9]}, nil, nil, nil, {tk[19]}},
						Tokens:   tk[9:20],
					},
					Tokens: tk[9:20],
				},
				Comments: [3]Comments{{tk[1]}, {tk[5]}, {tk[21]}},
				Tokens:   tk[1:22],
			}
		}},
		{`a`, func(t *test, tk Tokens) { // 7
			t.Err = Error{
				Err:     ErrMissingIf,
				Parsing: "ComprehensionIf",
				Token:   tk[0],
			}
		}},
		{`if nonlocal`, func(t *test, tk Tokens) { // 8
			t.Err = Error{
				Err: wrapConditionalExpressionError(Error{
					Err:     ErrInvalidEnclosure,
					Parsing: "Enclosure",
					Token:   tk[2],
				}).Err,
				Parsing: "ComprehensionIf",
				Token:   tk[2],
			}
		}},
		{`if a if nonlocal`, func(t *test, tk Tokens) { // 9
			t.Err = Error{
				Err: Error{
					Err: Error{
						Err: wrapConditionalExpressionError(Error{
							Err:     ErrInvalidEnclosure,
							Parsing: "Enclosure",
							Token:   tk[6],
						}).Err,
						Parsing: "ComprehensionIf",
						Token:   tk[6],
					},
					Parsing: "ComprehensionIterator",
					Token:   tk[4],
				},
				Parsing: "ComprehensionIf",
				Token:   tk[4],
			}
		}},
	}, func(t *test) (Type, error) {
		var c ComprehensionIf

		if t.Tokens.Peek().Data == "(" {
			t.Tokens.Tokens = t.Tokens.Tokens[1:1]
		}

		err := c.parse(t.Tokens)

		return c, err
	})
}

func TestDictDisplay(t *testing.T) {
	doTests(t, []sourceFn{
		{`a: b`, func(t *test, tk Tokens) { // 1
			t.Output = DictDisplay{
				DictItems: []DictItem{
					{
						Key: &Expression{
							ConditionalExpression: WrapConditional(&Atom{
								Identifier: &tk[0],
								Tokens:     tk[:1],
							}),
							Tokens: tk[:1],
						},
						Value: &Expression{
							ConditionalExpression: WrapConditional(&Atom{
								Identifier: &tk[3],
								Tokens:     tk[3:4],
							}),
							Tokens: tk[3:4],
						},
						Tokens: tk[:4],
					},
				},
				Tokens: tk[:4],
			}
		}},
		{`a :b, c:d`, func(t *test, tk Tokens) { // 2
			t.Output = DictDisplay{
				DictItems: []DictItem{
					{
						Key: &Expression{
							ConditionalExpression: WrapConditional(&Atom{
								Identifier: &tk[0],
								Tokens:     tk[:1],
							}),
							Tokens: tk[:1],
						},
						Value: &Expression{
							ConditionalExpression: WrapConditional(&Atom{
								Identifier: &tk[3],
								Tokens:     tk[3:4],
							}),
							Tokens: tk[3:4],
						},
						Tokens: tk[:4],
					},
					{
						Key: &Expression{
							ConditionalExpression: WrapConditional(&Atom{
								Identifier: &tk[6],
								Tokens:     tk[6:7],
							}),
							Tokens: tk[6:7],
						},
						Value: &Expression{
							ConditionalExpression: WrapConditional(&Atom{
								Identifier: &tk[8],
								Tokens:     tk[8:9],
							}),
							Tokens: tk[8:9],
						},
						Tokens: tk[6:9],
					},
				},
				Tokens: tk[:9],
			}
		}},
		{`a: b for c in d`, func(t *test, tk Tokens) { // 3
			t.Output = DictDisplay{
				DictItems: []DictItem{
					{
						Key: &Expression{
							ConditionalExpression: WrapConditional(&Atom{
								Identifier: &tk[0],
								Tokens:     tk[:1],
							}),
							Tokens: tk[:1],
						},
						Value: &Expression{
							ConditionalExpression: WrapConditional(&Atom{
								Identifier: &tk[3],
								Tokens:     tk[3:4],
							}),
							Tokens: tk[3:4],
						},
						Tokens: tk[:4],
					},
				},
				DictComprehension: &ComprehensionFor{
					TargetList: TargetList{
						Targets: []Target{
							{
								PrimaryExpression: &PrimaryExpression{
									Atom: &Atom{
										Identifier: &tk[7],
										Tokens:     tk[7:8],
									},
									Tokens: tk[7:8],
								},
								Tokens: tk[7:8],
							},
						},
						Tokens: tk[7:8],
					},
					OrTest: WrapConditional(&Atom{
						Identifier: &tk[11],
						Tokens:     tk[11:12],
					}).OrTest,
					Tokens: tk[5:12],
				},
				Tokens: tk[:12],
			}
		}},
		{`a: b async for c in d`, func(t *test, tk Tokens) { // 4
			t.Output = DictDisplay{
				DictItems: []DictItem{
					{
						Key: &Expression{
							ConditionalExpression: WrapConditional(&Atom{
								Identifier: &tk[0],
								Tokens:     tk[:1],
							}),
							Tokens: tk[:1],
						},
						Value: &Expression{
							ConditionalExpression: WrapConditional(&Atom{
								Identifier: &tk[3],
								Tokens:     tk[3:4],
							}),
							Tokens: tk[3:4],
						},
						Tokens: tk[:4],
					},
				},
				DictComprehension: &ComprehensionFor{
					Async: true,
					TargetList: TargetList{
						Targets: []Target{
							{
								PrimaryExpression: &PrimaryExpression{
									Atom: &Atom{
										Identifier: &tk[9],
										Tokens:     tk[9:10],
									},
									Tokens: tk[9:10],
								},
								Tokens: tk[9:10],
							},
						},
						Tokens: tk[9:10],
					},
					OrTest: WrapConditional(&Atom{
						Identifier: &tk[13],
						Tokens:     tk[13:14],
					}).OrTest,
					Tokens: tk[5:14],
				},
				Tokens: tk[:14],
			}
		}},
		{`**a`, func(t *test, tk Tokens) { // 5
			t.Output = DictDisplay{
				DictItems: []DictItem{
					{
						OrExpression: &WrapConditional(&Atom{
							Identifier: &tk[1],
							Tokens:     tk[1:2],
						}).OrTest.AndTest.NotTest.Comparison.OrExpression,
						Tokens: tk[:2],
					},
				},
				Tokens: tk[:2],
			}
		}},
		{`a: b, ** c`, func(t *test, tk Tokens) { // 6
			t.Output = DictDisplay{
				DictItems: []DictItem{
					{
						Key: &Expression{
							ConditionalExpression: WrapConditional(&Atom{
								Identifier: &tk[0],
								Tokens:     tk[:1],
							}),
							Tokens: tk[:1],
						},
						Value: &Expression{
							ConditionalExpression: WrapConditional(&Atom{
								Identifier: &tk[3],
								Tokens:     tk[3:4],
							}),
							Tokens: tk[3:4],
						},
						Tokens: tk[:4],
					},
					{
						OrExpression: &WrapConditional(&Atom{
							Identifier: &tk[8],
							Tokens:     tk[8:9],
						}).OrTest.AndTest.NotTest.Comparison.OrExpression,
						Tokens: tk[6:9],
					},
				},
				Tokens: tk[:9],
			}
		}},
		{`**a, b:c`, func(t *test, tk Tokens) { // 7
			t.Output = DictDisplay{
				DictItems: []DictItem{
					{
						OrExpression: &WrapConditional(&Atom{
							Identifier: &tk[1],
							Tokens:     tk[1:2],
						}).OrTest.AndTest.NotTest.Comparison.OrExpression,
						Tokens: tk[:2],
					},
					{
						Key: &Expression{
							ConditionalExpression: WrapConditional(&Atom{
								Identifier: &tk[4],
								Tokens:     tk[4:5],
							}),
							Tokens: tk[4:5],
						},
						Value: &Expression{
							ConditionalExpression: WrapConditional(&Atom{
								Identifier: &tk[6],
								Tokens:     tk[6:7],
							}),
							Tokens: tk[6:7],
						},
						Tokens: tk[4:7],
					},
				},
				Tokens: tk[:7],
			}
		}},
		{`a: b`, func(t *test, tk Tokens) { // 8
			t.Expression = &Expression{
				ConditionalExpression: WrapConditional(&Atom{
					Identifier: &tk[0],
					Tokens:     tk[:1],
				}),
				Tokens: tk[:1],
			}
			t.TokenSkip = 1
			t.Output = DictDisplay{
				DictItems: []DictItem{
					{
						Key: &Expression{
							ConditionalExpression: WrapConditional(&Atom{
								Identifier: &tk[0],
								Tokens:     tk[:1],
							}),
							Tokens: tk[:1],
						},
						Value: &Expression{
							ConditionalExpression: WrapConditional(&Atom{
								Identifier: &tk[3],
								Tokens:     tk[3:4],
							}),
							Tokens: tk[3:4],
						},
						Tokens: tk[:4],
					},
				},
				Tokens: tk[:4],
			}
		}},
		{`a :b, c:d`, func(t *test, tk Tokens) { // 9
			t.Expression = &Expression{
				ConditionalExpression: WrapConditional(&Atom{
					Identifier: &tk[0],
					Tokens:     tk[:1],
				}),
				Tokens: tk[:1],
			}
			t.TokenSkip = 1
			t.Output = DictDisplay{
				DictItems: []DictItem{
					{
						Key: &Expression{
							ConditionalExpression: WrapConditional(&Atom{
								Identifier: &tk[0],
								Tokens:     tk[:1],
							}),
							Tokens: tk[:1],
						},
						Value: &Expression{
							ConditionalExpression: WrapConditional(&Atom{
								Identifier: &tk[3],
								Tokens:     tk[3:4],
							}),
							Tokens: tk[3:4],
						},
						Tokens: tk[:4],
					},
					{
						Key: &Expression{
							ConditionalExpression: WrapConditional(&Atom{
								Identifier: &tk[6],
								Tokens:     tk[6:7],
							}),
							Tokens: tk[6:7],
						},
						Value: &Expression{
							ConditionalExpression: WrapConditional(&Atom{
								Identifier: &tk[8],
								Tokens:     tk[8:9],
							}),
							Tokens: tk[8:9],
						},
						Tokens: tk[6:9],
					},
				},
				Tokens: tk[:9],
			}
		}},
		{`a: b,`, func(t *test, tk Tokens) { // 10
			t.Expression = &Expression{
				ConditionalExpression: WrapConditional(&Atom{
					Identifier: &tk[0],
					Tokens:     tk[:1],
				}),
				Tokens: tk[:1],
			}
			t.TokenSkip = 1
			t.Output = DictDisplay{
				DictItems: []DictItem{
					{
						Key: &Expression{
							ConditionalExpression: WrapConditional(&Atom{
								Identifier: &tk[0],
								Tokens:     tk[:1],
							}),
							Tokens: tk[:1],
						},
						Value: &Expression{
							ConditionalExpression: WrapConditional(&Atom{
								Identifier: &tk[3],
								Tokens:     tk[3:4],
							}),
							Tokens: tk[3:4],
						},
						Tokens: tk[:4],
					},
				},
				Tokens: tk[:4],
			}
		}},
		{`a: nonlocal`, func(t *test, tk Tokens) { // 11
			t.Err = Error{
				Err: Error{
					Err: Error{
						Err: wrapConditionalExpressionError(Error{
							Err:     ErrInvalidEnclosure,
							Parsing: "Enclosure",
							Token:   tk[3],
						}),
						Parsing: "Expression",
						Token:   tk[3],
					},
					Parsing: "DictItem",
					Token:   tk[3],
				},
				Parsing: "DictDisplay",
				Token:   tk[0],
			}
		}},
		{`a: b, c: d for e in f`, func(t *test, tk Tokens) { // 12
			t.Err = Error{
				Err:     ErrInvalidKeyword,
				Parsing: "DictDisplay",
				Token:   tk[11],
			}
		}},
		{`**a for e in f`, func(t *test, tk Tokens) { // 13
			t.Err = Error{
				Err:     ErrInvalidKeyword,
				Parsing: "DictDisplay",
				Token:   tk[3],
			}
		}},
		{`a: b for nonlocal in f`, func(t *test, tk Tokens) { // 14
			t.Err = Error{
				Err: Error{
					Err: Error{
						Err: Error{
							Err: Error{
								Err: Error{
									Err: Error{
										Err:     ErrInvalidEnclosure,
										Parsing: "Enclosure",
										Token:   tk[7],
									},
									Parsing: "Atom",
									Token:   tk[7],
								},
								Parsing: "PrimaryExpression",
								Token:   tk[7],
							},
							Parsing: "Target",
							Token:   tk[7],
						},
						Parsing: "TargetList",
						Token:   tk[7],
					},
					Parsing: "ComprehensionFor",
					Token:   tk[7],
				},
				Parsing: "DictDisplay",
				Token:   tk[5],
			}
		}},
	}, func(t *test) (Type, error) {
		var d DictDisplay

		err := d.parse(t.Tokens, t.OrigTokens, t.Expression)

		return d, err
	})
}

func TestDictItem(t *testing.T) {
	doTests(t, []sourceFn{
		{`a: b`, func(t *test, tk Tokens) { // 1
			t.Output = DictItem{
				Key: &Expression{
					ConditionalExpression: WrapConditional(&Atom{
						Identifier: &tk[0],
						Tokens:     tk[:1],
					}),
					Tokens: tk[:1],
				},
				Value: &Expression{
					ConditionalExpression: WrapConditional(&Atom{
						Identifier: &tk[3],
						Tokens:     tk[3:4],
					}),
					Tokens: tk[3:4],
				},
				Tokens: tk[:4],
			}
		}},
		{`a: b`, func(t *test, tk Tokens) { // 2
			t.Expression = &Expression{
				ConditionalExpression: WrapConditional(&Atom{
					Identifier: &tk[0],
					Tokens:     tk[:1],
				}),
				Tokens: tk[:1],
			}
			t.TokenSkip = 1
			t.Output = DictItem{
				Key: &Expression{
					ConditionalExpression: WrapConditional(&Atom{
						Identifier: &tk[0],
						Tokens:     tk[:1],
					}),
					Tokens: tk[:1],
				},
				Value: &Expression{
					ConditionalExpression: WrapConditional(&Atom{
						Identifier: &tk[3],
						Tokens:     tk[3:4],
					}),
					Tokens: tk[3:4],
				},
				Tokens: tk[:4],
			}
		}},
		{`**a`, func(t *test, tk Tokens) { // 3
			t.Output = DictItem{
				OrExpression: &WrapConditional(&Atom{
					Identifier: &tk[1],
					Tokens:     tk[1:2],
				}).OrTest.AndTest.NotTest.Comparison.OrExpression,
				Tokens: tk[:2],
			}
		}},
		{`nonlocal: b`, func(t *test, tk Tokens) { // 4
			t.Err = Error{
				Err: Error{
					Err: wrapConditionalExpressionError(Error{
						Err:     ErrInvalidEnclosure,
						Parsing: "Enclosure",
						Token:   tk[0],
					}),
					Parsing: "Expression",
					Token:   tk[0],
				},
				Parsing: "DictItem",
				Token:   tk[0],
			}
		}},
		{`a: nonlocal`, func(t *test, tk Tokens) { // 5
			t.Err = Error{
				Err: Error{
					Err: wrapConditionalExpressionError(Error{
						Err:     ErrInvalidEnclosure,
						Parsing: "Enclosure",
						Token:   tk[3],
					}),
					Parsing: "Expression",
					Token:   tk[3],
				},
				Parsing: "DictItem",
				Token:   tk[3],
			}
		}},
		{`**nonlocal`, func(t *test, tk Tokens) { // 6
			t.Err = Error{
				Err: Error{
					Err: Error{
						Err: Error{
							Err: Error{
								Err: Error{
									Err: Error{
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
										Parsing: "MultiplyExpression",
										Token:   tk[1],
									},
									Parsing: "AddExpression",
									Token:   tk[1],
								},
								Parsing: "ShiftExpression",
								Token:   tk[1],
							},
							Parsing: "AndExpression",
							Token:   tk[1],
						},
						Parsing: "XorExpression",
						Token:   tk[1],
					},
					Parsing: "OrExpression",
					Token:   tk[1],
				},
				Parsing: "DictItem",
				Token:   tk[1],
			}
		}},
		{`a`, func(t *test, tk Tokens) { // 7
			t.Err = Error{
				Err:     ErrMissingColon,
				Parsing: "DictItem",
				Token:   tk[1],
			}
		}},
	}, func(t *test) (Type, error) {
		var d DictItem

		err := d.parse(t.Tokens, t.Expression)

		return d, err
	})
}

func TestGeneratorExpression(t *testing.T) {
	doTests(t, []sourceFn{
		{`a for b in c`, func(t *test, tk Tokens) { // 1
			t.Output = GeneratorExpression{
				Expression: Expression{
					ConditionalExpression: WrapConditional(&Atom{
						Identifier: &tk[0],
						Tokens:     tk[:1],
					}),
					Tokens: tk[:1],
				},
				ComprehensionFor: ComprehensionFor{
					TargetList: TargetList{
						Targets: []Target{
							{
								PrimaryExpression: &PrimaryExpression{
									Atom: &Atom{
										Identifier: &tk[4],
										Tokens:     tk[4:5],
									},
									Tokens: tk[4:5],
								},
								Tokens: tk[4:5],
							},
						},
						Tokens: tk[4:5],
					},
					OrTest: WrapConditional(&Atom{
						Identifier: &tk[8],
						Tokens:     tk[8:9],
					}).OrTest,
					Tokens: tk[2:9],
				},
				Tokens: tk[0:9],
			}
		}},
		{"(# A\na # B\nfor b in c\n# C\n)", func(t *test, tk Tokens) { // 2
			t.Output = GeneratorExpression{
				Expression: Expression{
					ConditionalExpression: WrapConditional(&Atom{
						Identifier: &tk[3],
						Tokens:     tk[3:4],
					}),
					Tokens: tk[3:4],
				},
				ComprehensionFor: ComprehensionFor{
					TargetList: TargetList{
						Targets: []Target{
							{
								PrimaryExpression: &PrimaryExpression{
									Atom: &Atom{
										Identifier: &tk[9],
										Tokens:     tk[9:10],
									},
									Tokens: tk[9:10],
								},
								Tokens: tk[9:10],
							},
						},
						Tokens: tk[9:10],
					},
					OrTest: WrapConditional(&Atom{
						Identifier: &tk[13],
						Tokens:     tk[13:14],
					}).OrTest,
					Tokens: tk[7:14],
				},
				Comments: [3]Comments{{tk[1]}, {tk[5]}, {tk[15]}},
				Tokens:   tk[1:16],
			}
		}},
		{`nonlocal for b in c`, func(t *test, tk Tokens) { // 3
			t.Err = Error{
				Err: Error{
					Err: wrapConditionalExpressionError(Error{
						Err:     ErrInvalidEnclosure,
						Parsing: "Enclosure",
						Token:   tk[0],
					}),
					Parsing: "Expression",
					Token:   tk[0],
				},
				Parsing: "GeneratorExpression",
				Token:   tk[0],
			}
		}},
		{`a for b in nonlocal`, func(t *test, tk Tokens) { // 4
			t.Err = Error{
				Err: Error{
					Err: wrapConditionalExpressionError(Error{
						Err:     ErrInvalidEnclosure,
						Parsing: "Enclosure",
						Token:   tk[8],
					}).Err,
					Parsing: "ComprehensionFor",
					Token:   tk[8],
				},
				Parsing: "GeneratorExpression",
				Token:   tk[2],
			}
		}},
	}, func(t *test) (Type, error) {
		var g GeneratorExpression

		if t.Tokens.Peek().Data == "(" {
			t.Tokens.Tokens = t.Tokens.Tokens[1:1]
		}

		err := g.parse(t.Tokens)

		return g, err
	})
}

func TestArgumentListOrComprehension(t *testing.T) {
	doTests(t, []sourceFn{
		{`()`, func(t *test, tk Tokens) { // 1
			t.Output = ArgumentListOrComprehension{
				ArgumentList: &ArgumentList{
					Tokens: tk[1:1],
				},
				Tokens: tk[:2],
			}
		}},
		{`(a)`, func(t *test, tk Tokens) { // 2
			t.Output = ArgumentListOrComprehension{
				ArgumentList: &ArgumentList{
					PositionalArguments: []PositionalArgument{
						{
							AssignmentExpression: &AssignmentExpression{
								Expression: Expression{
									ConditionalExpression: WrapConditional(&Atom{
										Identifier: &tk[1],
										Tokens:     tk[1:2],
									}),
									Tokens: tk[1:2],
								},
								Tokens: tk[1:2],
							},
							Tokens: tk[1:2],
						},
					},
					Tokens: tk[1:2],
				},
				Tokens: tk[:3],
			}
		}},
		{`(a for b in c)`, func(t *test, tk Tokens) { // 3
			t.Output = ArgumentListOrComprehension{
				Comprehension: &Comprehension{
					AssignmentExpression: AssignmentExpression{
						Expression: Expression{
							ConditionalExpression: WrapConditional(&Atom{
								Identifier: &tk[1],
								Tokens:     tk[1:2],
							}),
							Tokens: tk[1:2],
						},
						Tokens: tk[1:2],
					},
					ComprehensionFor: ComprehensionFor{
						TargetList: TargetList{
							Targets: []Target{
								{
									PrimaryExpression: &PrimaryExpression{
										Atom: &Atom{
											Identifier: &tk[5],
											Tokens:     tk[5:6],
										},
										Tokens: tk[5:6],
									},
									Tokens: tk[5:6],
								},
							},
							Tokens: tk[5:6],
						},
						OrTest: WrapConditional(&Atom{
							Identifier: &tk[9],
							Tokens:     tk[9:10],
						}).OrTest,
						Tokens: tk[3:10],
					},
					Tokens: tk[1:10],
				},
				Tokens: tk[:11],
			}
		}},
		{`(nonlocal)`, func(t *test, tk Tokens) { // 4
			t.Err = Error{
				Err: Error{
					Err: Error{
						Err: Error{
							Err: Error{
								Err: wrapConditionalExpressionError(Error{
									Err:     ErrInvalidEnclosure,
									Parsing: "Enclosure",
									Token:   tk[1],
								}),
								Parsing: "Expression",
								Token:   tk[1],
							},
							Parsing: "AssignmentExpression",
							Token:   tk[1],
						},
						Parsing: "PositionalArgument",
						Token:   tk[1],
					},
					Parsing: "ArgumentList",
					Token:   tk[1],
				},
				Parsing: "ArgumentListOrComprehension",
				Token:   tk[1],
			}
		}},
		{`(nonlocal for a in b)`, func(t *test, tk Tokens) { // 5
			t.Err = Error{
				Err: Error{
					Err: Error{
						Err: Error{
							Err: wrapConditionalExpressionError(Error{
								Err:     ErrInvalidEnclosure,
								Parsing: "Enclosure",
								Token:   tk[1],
							}),
							Parsing: "Expression",
							Token:   tk[1],
						},
						Parsing: "AssignmentExpression",
						Token:   tk[1],
					},
					Parsing: "Comprehension",
					Token:   tk[1],
				},
				Parsing: "ArgumentListOrComprehension",
				Token:   tk[1],
			}
		}},
		{`(a for b in c d)`, func(t *test, tk Tokens) { // 6
			t.Err = Error{
				Err:     ErrMissingClosingParen,
				Parsing: "ArgumentListOrComprehension",
				Token:   tk[11],
			}
		}},
	}, func(t *test) (Type, error) {
		var a ArgumentListOrComprehension

		err := a.parse(t.Tokens)

		return a, err
	})
}

func TestExpressionList(t *testing.T) {
	doTests(t, []sourceFn{
		{`a`, func(t *test, tk Tokens) { // 1
			t.Output = ExpressionList{
				Expressions: []Expression{
					{
						ConditionalExpression: WrapConditional(&Atom{
							Identifier: &tk[0],
							Tokens:     tk[:1],
						}),
						Tokens: tk[:1],
					},
				},
				Tokens: tk[:1],
			}
		}},
		{"a\n", func(t *test, tk Tokens) { // 2
			t.Output = ExpressionList{
				Expressions: []Expression{
					{
						ConditionalExpression: WrapConditional(&Atom{
							Identifier: &tk[0],
							Tokens:     tk[:1],
						}),
						Tokens: tk[:1],
					},
				},
				Tokens: tk[:1],
			}
		}},
		{`a,b`, func(t *test, tk Tokens) { // 3
			t.Output = ExpressionList{
				Expressions: []Expression{
					{
						ConditionalExpression: WrapConditional(&Atom{
							Identifier: &tk[0],
							Tokens:     tk[:1],
						}),
						Tokens: tk[:1],
					},
					{
						ConditionalExpression: WrapConditional(&Atom{
							Identifier: &tk[2],
							Tokens:     tk[2:3],
						}),
						Tokens: tk[2:3],
					},
				},
				Tokens: tk[:3],
			}
		}},
		{`a, b`, func(t *test, tk Tokens) { // 4
			t.Output = ExpressionList{
				Expressions: []Expression{
					{
						ConditionalExpression: WrapConditional(&Atom{
							Identifier: &tk[0],
							Tokens:     tk[:1],
						}),
						Tokens: tk[:1],
					},
					{
						ConditionalExpression: WrapConditional(&Atom{
							Identifier: &tk[3],
							Tokens:     tk[3:4],
						}),
						Tokens: tk[3:4],
					},
				},
				Tokens: tk[:4],
			}
		}},
		{`a,b,`, func(t *test, tk Tokens) { // 5
			t.Output = ExpressionList{
				Expressions: []Expression{
					{
						ConditionalExpression: WrapConditional(&Atom{
							Identifier: &tk[0],
							Tokens:     tk[:1],
						}),
						Tokens: tk[:1],
					},
					{
						ConditionalExpression: WrapConditional(&Atom{
							Identifier: &tk[2],
							Tokens:     tk[2:3],
						}),
						Tokens: tk[2:3],
					},
				},
				Tokens: tk[:3],
			}
		}},
		{`nonlocal`, func(t *test, tk Tokens) { // 6
			t.Err = Error{
				Err: Error{
					Err: wrapConditionalExpressionError(Error{
						Err:     ErrInvalidEnclosure,
						Parsing: "Enclosure",
						Token:   tk[0],
					}),
					Parsing: "Expression",
					Token:   tk[0],
				},
				Parsing: "ExpressionList",
				Token:   tk[0],
			}
		}},
		{`a,nonlocal`, func(t *test, tk Tokens) { // 7
			t.Err = Error{
				Err: Error{
					Err: wrapConditionalExpressionError(Error{
						Err:     ErrInvalidEnclosure,
						Parsing: "Enclosure",
						Token:   tk[2],
					}),
					Parsing: "Expression",
					Token:   tk[2],
				},
				Parsing: "ExpressionList",
				Token:   tk[2],
			}
		}},
		{`a b`, func(t *test, tk Tokens) { // 8
			t.Err = Error{
				Err:     ErrMissingComma,
				Parsing: "ExpressionList",
				Token:   tk[1],
			}
		}},
	}, func(t *test) (Type, error) {
		var e ExpressionList

		err := e.parse(t.Tokens)

		return e, err
	})
}

func TestSliceList(t *testing.T) {
	doTests(t, []sourceFn{
		{`[]`, func(t *test, tk Tokens) { // 1
			t.Output = SliceList{
				Tokens: tk[:2],
			}
		}},
		{`[ ]`, func(t *test, tk Tokens) { // 2
			t.Output = SliceList{
				Tokens: tk[:3],
			}
		}},
		{`[a]`, func(t *test, tk Tokens) { // 3
			t.Output = SliceList{
				SliceItems: []SliceItem{
					{
						Expression: &Expression{
							ConditionalExpression: WrapConditional(&Atom{
								Identifier: &tk[1],
								Tokens:     tk[1:2],
							}),
							Tokens: tk[1:2],
						},
						Tokens: tk[1:2],
					},
				},
				Tokens: tk[:3],
			}
		}},
		{`[ a ]`, func(t *test, tk Tokens) { // 4
			t.Output = SliceList{
				SliceItems: []SliceItem{
					{
						Expression: &Expression{
							ConditionalExpression: WrapConditional(&Atom{
								Identifier: &tk[2],
								Tokens:     tk[2:3],
							}),
							Tokens: tk[2:3],
						},
						Tokens: tk[2:3],
					},
				},
				Tokens: tk[:5],
			}
		}},
		{`[a,]`, func(t *test, tk Tokens) { // 5
			t.Output = SliceList{
				SliceItems: []SliceItem{
					{
						Expression: &Expression{
							ConditionalExpression: WrapConditional(&Atom{
								Identifier: &tk[1],
								Tokens:     tk[1:2],
							}),
							Tokens: tk[1:2],
						},
						Tokens: tk[1:2],
					},
				},
				Tokens: tk[:4],
			}
		}},
		{`[a,b]`, func(t *test, tk Tokens) { // 6
			t.Output = SliceList{
				SliceItems: []SliceItem{
					{
						Expression: &Expression{
							ConditionalExpression: WrapConditional(&Atom{
								Identifier: &tk[1],
								Tokens:     tk[1:2],
							}),
							Tokens: tk[1:2],
						},
						Tokens: tk[1:2],
					},
					{
						Expression: &Expression{
							ConditionalExpression: WrapConditional(&Atom{
								Identifier: &tk[3],
								Tokens:     tk[3:4],
							}),
							Tokens: tk[3:4],
						},
						Tokens: tk[3:4],
					},
				},
				Tokens: tk[:5],
			}
		}},
		{`[a, b , ]`, func(t *test, tk Tokens) { // 7
			t.Output = SliceList{
				SliceItems: []SliceItem{
					{
						Expression: &Expression{
							ConditionalExpression: WrapConditional(&Atom{
								Identifier: &tk[1],
								Tokens:     tk[1:2],
							}),
							Tokens: tk[1:2],
						},
						Tokens: tk[1:2],
					},
					{
						Expression: &Expression{
							ConditionalExpression: WrapConditional(&Atom{
								Identifier: &tk[4],
								Tokens:     tk[4:5],
							}),
							Tokens: tk[4:5],
						},
						Tokens: tk[4:5],
					},
				},
				Tokens: tk[:9],
			}
		}},
		{"[ # A\na\n# B\n]", func(t *test, tk Tokens) { // 8
			t.Output = SliceList{
				SliceItems: []SliceItem{
					{
						Expression: &Expression{
							ConditionalExpression: WrapConditional(&Atom{
								Identifier: &tk[4],
								Tokens:     tk[4:5],
							}),
							Tokens: tk[4:5],
						},
						Tokens: tk[4:5],
					},
				},
				Comments: [2]Comments{{tk[2]}, {tk[6]}},
				Tokens:   tk[:9],
			}
		}},
		{"[ # A\na,\n# B\n]", func(t *test, tk Tokens) { // 9
			t.Output = SliceList{
				SliceItems: []SliceItem{
					{
						Expression: &Expression{
							ConditionalExpression: WrapConditional(&Atom{
								Identifier: &tk[4],
								Tokens:     tk[4:5],
							}),
							Tokens: tk[4:5],
						},
						Tokens: tk[4:5],
					},
				},
				Comments: [2]Comments{{tk[2]}, {tk[7]}},
				Tokens:   tk[:10],
			}
		}},
		{"[ # A\na,b\n# B\n]", func(t *test, tk Tokens) { // 10
			t.Output = SliceList{
				SliceItems: []SliceItem{
					{
						Expression: &Expression{
							ConditionalExpression: WrapConditional(&Atom{
								Identifier: &tk[4],
								Tokens:     tk[4:5],
							}),
							Tokens: tk[4:5],
						},
						Tokens: tk[4:5],
					},
					{
						Expression: &Expression{
							ConditionalExpression: WrapConditional(&Atom{
								Identifier: &tk[6],
								Tokens:     tk[6:7],
							}),
							Tokens: tk[6:7],
						},
						Tokens: tk[6:7],
					},
				},
				Comments: [2]Comments{{tk[2]}, {tk[8]}},
				Tokens:   tk[:11],
			}
		}},
		{`[nonlocal]`, func(t *test, tk Tokens) { // 11
			t.Err = Error{
				Err: Error{
					Err: Error{
						Err: wrapConditionalExpressionError(Error{
							Err:     ErrInvalidEnclosure,
							Parsing: "Enclosure",
							Token:   tk[1],
						}),
						Parsing: "Expression",
						Token:   tk[1],
					},
					Parsing: "SliceItem",
					Token:   tk[1],
				},
				Parsing: "SliceList",
				Token:   tk[1],
			}
		}},
		{`[a b]`, func(t *test, tk Tokens) { // 12
			t.Err = Error{
				Err:     ErrMissingComma,
				Parsing: "SliceList",
				Token:   tk[3],
			}
		}},
	}, func(t *test) (Type, error) {
		var sl SliceList

		err := sl.parse(t.Tokens)

		return sl, err
	})
}

func TestSliceItem(t *testing.T) {
	doTests(t, []sourceFn{
		{`a`, func(t *test, tk Tokens) { // 1
			t.Output = SliceItem{
				Expression: &Expression{
					ConditionalExpression: WrapConditional(&Atom{
						Identifier: &tk[0],
						Tokens:     tk[:1],
					}),
					Tokens: tk[:1],
				},
				Tokens: tk[:1],
			}
		}},
		{`a:b`, func(t *test, tk Tokens) { // 2
			t.Output = SliceItem{
				LowerBound: &Expression{
					ConditionalExpression: WrapConditional(&Atom{
						Identifier: &tk[0],
						Tokens:     tk[:1],
					}),
					Tokens: tk[:1],
				},
				UpperBound: &Expression{
					ConditionalExpression: WrapConditional(&Atom{
						Identifier: &tk[2],
						Tokens:     tk[2:3],
					}),
					Tokens: tk[2:3],
				},
				Tokens: tk[:3],
			}
		}},
		{`a: b`, func(t *test, tk Tokens) { // 3
			t.Output = SliceItem{
				LowerBound: &Expression{
					ConditionalExpression: WrapConditional(&Atom{
						Identifier: &tk[0],
						Tokens:     tk[:1],
					}),
					Tokens: tk[:1],
				},
				UpperBound: &Expression{
					ConditionalExpression: WrapConditional(&Atom{
						Identifier: &tk[3],
						Tokens:     tk[3:4],
					}),
					Tokens: tk[3:4],
				},
				Tokens: tk[:4],
			}
		}},
		{`a:b:c`, func(t *test, tk Tokens) { // 4
			t.Output = SliceItem{
				LowerBound: &Expression{
					ConditionalExpression: WrapConditional(&Atom{
						Identifier: &tk[0],
						Tokens:     tk[:1],
					}),
					Tokens: tk[:1],
				},
				UpperBound: &Expression{
					ConditionalExpression: WrapConditional(&Atom{
						Identifier: &tk[2],
						Tokens:     tk[2:3],
					}),
					Tokens: tk[2:3],
				},
				Stride: &Expression{
					ConditionalExpression: WrapConditional(&Atom{
						Identifier: &tk[4],
						Tokens:     tk[4:5],
					}),
					Tokens: tk[4:5],
				},
				Tokens: tk[:5],
			}
		}},
		{`a:b: c`, func(t *test, tk Tokens) { // 5
			t.Output = SliceItem{
				LowerBound: &Expression{
					ConditionalExpression: WrapConditional(&Atom{
						Identifier: &tk[0],
						Tokens:     tk[:1],
					}),
					Tokens: tk[:1],
				},
				UpperBound: &Expression{
					ConditionalExpression: WrapConditional(&Atom{
						Identifier: &tk[2],
						Tokens:     tk[2:3],
					}),
					Tokens: tk[2:3],
				},
				Stride: &Expression{
					ConditionalExpression: WrapConditional(&Atom{
						Identifier: &tk[5],
						Tokens:     tk[5:6],
					}),
					Tokens: tk[5:6],
				},
				Tokens: tk[:6],
			}
		}},
		{"[# A\na # B\n]", func(t *test, tk Tokens) { // 6
			t.Output = SliceItem{
				Expression: &Expression{
					ConditionalExpression: WrapConditional(&Atom{
						Identifier: &tk[3],
						Tokens:     tk[3:4],
					}),
					Tokens: tk[3:4],
				},
				Comments: [6]Comments{{tk[1]}, nil, nil, nil, nil, {tk[5]}},
				Tokens:   tk[1:6],
			}
		}},
		{"[# A\na # B\n: # C\nb # D\n]", func(t *test, tk Tokens) { // 7
			t.Output = SliceItem{
				LowerBound: &Expression{
					ConditionalExpression: WrapConditional(&Atom{
						Identifier: &tk[3],
						Tokens:     tk[3:4],
					}),
					Tokens: tk[3:4],
				},
				UpperBound: &Expression{
					ConditionalExpression: WrapConditional(&Atom{
						Identifier: &tk[11],
						Tokens:     tk[11:12],
					}),
					Tokens: tk[11:12],
				},
				Comments: [6]Comments{{tk[1]}, {tk[5]}, {tk[9]}, nil, nil, {tk[13]}},
				Tokens:   tk[1:14],
			}
		}},
		{"[# A\na # B\n: # C\nb # D\n: # E\nc # F\n]", func(t *test, tk Tokens) { // 8
			t.Output = SliceItem{
				LowerBound: &Expression{
					ConditionalExpression: WrapConditional(&Atom{
						Identifier: &tk[3],
						Tokens:     tk[3:4],
					}),
					Tokens: tk[3:4],
				},
				UpperBound: &Expression{
					ConditionalExpression: WrapConditional(&Atom{
						Identifier: &tk[11],
						Tokens:     tk[11:12],
					}),
					Tokens: tk[11:12],
				},
				Stride: &Expression{
					ConditionalExpression: WrapConditional(&Atom{
						Identifier: &tk[19],
						Tokens:     tk[19:20],
					}),
					Tokens: tk[19:20],
				},
				Comments: [6]Comments{{tk[1]}, {tk[5]}, {tk[9]}, {tk[13]}, {tk[17]}, {tk[21]}},
				Tokens:   tk[1:22],
			}
		}},
		{`nonlocal`, func(t *test, tk Tokens) { // 9
			t.Err = Error{
				Err: Error{
					Err: wrapConditionalExpressionError(Error{
						Err:     ErrInvalidEnclosure,
						Parsing: "Enclosure",
						Token:   tk[0],
					}),
					Parsing: "Expression",
					Token:   tk[0],
				},
				Parsing: "SliceItem",
				Token:   tk[0],
			}
		}},
		{`a:nonlocal`, func(t *test, tk Tokens) { // 10
			t.Err = Error{
				Err: Error{
					Err: wrapConditionalExpressionError(Error{
						Err:     ErrInvalidEnclosure,
						Parsing: "Enclosure",
						Token:   tk[2],
					}),
					Parsing: "Expression",
					Token:   tk[2],
				},
				Parsing: "SliceItem",
				Token:   tk[2],
			}
		}},
		{`a:b:nonlocal`, func(t *test, tk Tokens) { // 11
			t.Err = Error{
				Err: Error{
					Err: wrapConditionalExpressionError(Error{
						Err:     ErrInvalidEnclosure,
						Parsing: "Enclosure",
						Token:   tk[4],
					}),
					Parsing: "Expression",
					Token:   tk[4],
				},
				Parsing: "SliceItem",
				Token:   tk[4],
			}
		}},
	}, func(t *test) (Type, error) {
		var si SliceItem

		if t.Tokens.Peek().Data == "[" {
			t.Tokens.Tokens = t.Tokens.Tokens[1:1]
		}

		err := si.parse(t.Tokens)

		return si, err
	})
}

func TestOrExpression(t *testing.T) {
	doTests(t, []sourceFn{
		{`a`, func(t *test, tk Tokens) { // 1
			t.Output = OrExpression{
				XorExpression: XorExpression{
					AndExpression: AndExpression{
						ShiftExpression: ShiftExpression{
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
						},
						Tokens: tk[:1],
					},
					Tokens: tk[:1],
				},
				Tokens: tk[:1],
			}
		}},
		{`a|b`, func(t *test, tk Tokens) { // 2
			t.Output = OrExpression{
				XorExpression: XorExpression{
					AndExpression: AndExpression{
						ShiftExpression: ShiftExpression{
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
						},
						Tokens: tk[:1],
					},
					Tokens: tk[:1],
				},
				OrExpression: &OrExpression{
					XorExpression: XorExpression{
						AndExpression: AndExpression{
							ShiftExpression: ShiftExpression{
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
							Tokens: tk[2:3],
						},
						Tokens: tk[2:3],
					},
					Tokens: tk[2:3],
				},
				Tokens: tk[:3],
			}
		}},
		{`a | b`, func(t *test, tk Tokens) { // 3
			t.Output = OrExpression{
				XorExpression: XorExpression{
					AndExpression: AndExpression{
						ShiftExpression: ShiftExpression{
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
						},
						Tokens: tk[:1],
					},
					Tokens: tk[:1],
				},
				OrExpression: &OrExpression{
					XorExpression: XorExpression{
						AndExpression: AndExpression{
							ShiftExpression: ShiftExpression{
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
							Tokens: tk[4:5],
						},
						Tokens: tk[4:5],
					},
					Tokens: tk[4:5],
				},
				Tokens: tk[:5],
			}
		}},
		{"(a # A\n| # B\nb)", func(t *test, tk Tokens) { // 4
			t.Output = OrExpression{
				XorExpression: XorExpression{
					AndExpression: AndExpression{
						ShiftExpression: ShiftExpression{
							AddExpression: AddExpression{
								MultiplyExpression: MultiplyExpression{
									UnaryExpression: UnaryExpression{
										PowerExpression: &PowerExpression{
											PrimaryExpression: PrimaryExpression{
												Atom: &Atom{
													Identifier: &tk[1],
													Tokens:     tk[1:2],
												},
												Tokens: tk[1:2],
											},
											Tokens: tk[1:2],
										},
										Tokens: tk[1:2],
									},
									Tokens: tk[1:2],
								},
								Tokens: tk[1:2],
							},
							Tokens: tk[1:2],
						},
						Tokens: tk[1:2],
					},
					Tokens: tk[1:2],
				},
				OrExpression: &OrExpression{
					XorExpression: XorExpression{
						AndExpression: AndExpression{
							ShiftExpression: ShiftExpression{
								AddExpression: AddExpression{
									MultiplyExpression: MultiplyExpression{
										UnaryExpression: UnaryExpression{
											PowerExpression: &PowerExpression{
												PrimaryExpression: PrimaryExpression{
													Atom: &Atom{
														Identifier: &tk[9],
														Tokens:     tk[9:10],
													},
													Tokens: tk[9:10],
												},
												Tokens: tk[9:10],
											},
											Tokens: tk[9:10],
										},
										Tokens: tk[9:10],
									},
									Tokens: tk[9:10],
								},
								Tokens: tk[9:10],
							},
							Tokens: tk[9:10],
						},
						Tokens: tk[9:10],
					},
					Tokens: tk[9:10],
				},
				Comments: [2]Comments{{tk[3]}, {tk[7]}},
				Tokens:   tk[1:10],
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
							},
							Parsing: "ShiftExpression",
							Token:   tk[0],
						},
						Parsing: "AndExpression",
						Token:   tk[0],
					},
					Parsing: "XorExpression",
					Token:   tk[0],
				},
				Parsing: "OrExpression",
				Token:   tk[0],
			}
		}},
		{`1|nonlocal`, func(t *test, tk Tokens) { // 6
			t.Err = Error{
				Err: Error{
					Err: Error{
						Err: Error{
							Err: Error{
								Err: Error{
									Err: Error{
										Err: Error{
											Err: Error{
												Err: Error{
													Err: Error{
														Err: Error{
															Err:     ErrInvalidEnclosure,
															Parsing: "Enclosure",
															Token:   tk[2],
														},
														Parsing: "Atom",
														Token:   tk[2],
													},
													Parsing: "PrimaryExpression",
													Token:   tk[2],
												},
												Parsing: "PowerExpression",
												Token:   tk[2],
											},
											Parsing: "UnaryExpression",
											Token:   tk[2],
										},
										Parsing: "MultiplyExpression",
										Token:   tk[2],
									},
									Parsing: "AddExpression",
									Token:   tk[2],
								},
								Parsing: "ShiftExpression",
								Token:   tk[2],
							},
							Parsing: "AndExpression",
							Token:   tk[2],
						},
						Parsing: "XorExpression",
						Token:   tk[2],
					},
					Parsing: "OrExpression",
					Token:   tk[2],
				},
				Parsing: "OrExpression",
				Token:   tk[2],
			}
		}},
	}, func(t *test) (Type, error) {
		var oe OrExpression

		if t.Tokens.Peek().Data == "(" {
			t.Tokens.Tokens = t.Tokens.Tokens[1:1]
		}

		err := oe.parse(t.Tokens)

		return oe, err
	})
}

func TestXorExpression(t *testing.T) {
	doTests(t, []sourceFn{
		{`a`, func(t *test, tk Tokens) { // 1
			t.Output = XorExpression{
				AndExpression: AndExpression{
					ShiftExpression: ShiftExpression{
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
					},
					Tokens: tk[:1],
				},
				Tokens: tk[:1],
			}
		}},
		{`a^b`, func(t *test, tk Tokens) { // 2
			t.Output = XorExpression{
				AndExpression: AndExpression{
					ShiftExpression: ShiftExpression{
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
					},
					Tokens: tk[:1],
				},
				XorExpression: &XorExpression{
					AndExpression: AndExpression{
						ShiftExpression: ShiftExpression{
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
						Tokens: tk[2:3],
					},
					Tokens: tk[2:3],
				},
				Tokens: tk[:3],
			}
		}},
		{`a ^ b`, func(t *test, tk Tokens) { // 3
			t.Output = XorExpression{
				AndExpression: AndExpression{
					ShiftExpression: ShiftExpression{
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
					},
					Tokens: tk[:1],
				},
				XorExpression: &XorExpression{
					AndExpression: AndExpression{
						ShiftExpression: ShiftExpression{
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
						Tokens: tk[4:5],
					},
					Tokens: tk[4:5],
				},
				Tokens: tk[:5],
			}
		}},
		{"(a # A\n^ # B\nb)", func(t *test, tk Tokens) { // 4
			t.Output = XorExpression{
				AndExpression: AndExpression{
					ShiftExpression: ShiftExpression{
						AddExpression: AddExpression{
							MultiplyExpression: MultiplyExpression{
								UnaryExpression: UnaryExpression{
									PowerExpression: &PowerExpression{
										PrimaryExpression: PrimaryExpression{
											Atom: &Atom{
												Identifier: &tk[1],
												Tokens:     tk[1:2],
											},
											Tokens: tk[1:2],
										},
										Tokens: tk[1:2],
									},
									Tokens: tk[1:2],
								},
								Tokens: tk[1:2],
							},
							Tokens: tk[1:2],
						},
						Tokens: tk[1:2],
					},
					Tokens: tk[1:2],
				},
				XorExpression: &XorExpression{
					AndExpression: AndExpression{
						ShiftExpression: ShiftExpression{
							AddExpression: AddExpression{
								MultiplyExpression: MultiplyExpression{
									UnaryExpression: UnaryExpression{
										PowerExpression: &PowerExpression{
											PrimaryExpression: PrimaryExpression{
												Atom: &Atom{
													Identifier: &tk[9],
													Tokens:     tk[9:10],
												},
												Tokens: tk[9:10],
											},
											Tokens: tk[9:10],
										},
										Tokens: tk[9:10],
									},
									Tokens: tk[9:10],
								},
								Tokens: tk[9:10],
							},
							Tokens: tk[9:10],
						},
						Tokens: tk[9:10],
					},
					Tokens: tk[9:10],
				},
				Comments: [2]Comments{{tk[3]}, {tk[7]}},
				Tokens:   tk[1:10],
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
						},
						Parsing: "ShiftExpression",
						Token:   tk[0],
					},
					Parsing: "AndExpression",
					Token:   tk[0],
				},
				Parsing: "XorExpression",
				Token:   tk[0],
			}
		}},
		{`1^nonlocal`, func(t *test, tk Tokens) { // 6
			t.Err = Error{
				Err: Error{
					Err: Error{
						Err: Error{
							Err: Error{
								Err: Error{
									Err: Error{
										Err: Error{
											Err: Error{
												Err: Error{
													Err: Error{
														Err:     ErrInvalidEnclosure,
														Parsing: "Enclosure",
														Token:   tk[2],
													},
													Parsing: "Atom",
													Token:   tk[2],
												},
												Parsing: "PrimaryExpression",
												Token:   tk[2],
											},
											Parsing: "PowerExpression",
											Token:   tk[2],
										},
										Parsing: "UnaryExpression",
										Token:   tk[2],
									},
									Parsing: "MultiplyExpression",
									Token:   tk[2],
								},
								Parsing: "AddExpression",
								Token:   tk[2],
							},
							Parsing: "ShiftExpression",
							Token:   tk[2],
						},
						Parsing: "AndExpression",
						Token:   tk[2],
					},
					Parsing: "XorExpression",
					Token:   tk[2],
				},
				Parsing: "XorExpression",
				Token:   tk[2],
			}
		}},
	}, func(t *test) (Type, error) {
		var xe XorExpression

		if t.Tokens.Peek().Data == "(" {
			t.Tokens.Tokens = t.Tokens.Tokens[1:1]
		}

		err := xe.parse(t.Tokens)

		return xe, err
	})
}

func TestAndExpression(t *testing.T) {
	doTests(t, []sourceFn{
		{`a`, func(t *test, tk Tokens) { // 1
			t.Output = AndExpression{
				ShiftExpression: ShiftExpression{
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
				},
				Tokens: tk[:1],
			}
		}},
		{`a&b`, func(t *test, tk Tokens) { // 2
			t.Output = AndExpression{
				ShiftExpression: ShiftExpression{
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
				},
				AndExpression: &AndExpression{
					ShiftExpression: ShiftExpression{
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
					Tokens: tk[2:3],
				},
				Tokens: tk[:3],
			}
		}},
		{`a & b`, func(t *test, tk Tokens) { // 3
			t.Output = AndExpression{
				ShiftExpression: ShiftExpression{
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
				},
				AndExpression: &AndExpression{
					ShiftExpression: ShiftExpression{
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
					Tokens: tk[4:5],
				},
				Tokens: tk[:5],
			}
		}},
		{"(a # A\n& # B\nb)", func(t *test, tk Tokens) { // 4
			t.Output = AndExpression{
				ShiftExpression: ShiftExpression{
					AddExpression: AddExpression{
						MultiplyExpression: MultiplyExpression{
							UnaryExpression: UnaryExpression{
								PowerExpression: &PowerExpression{
									PrimaryExpression: PrimaryExpression{
										Atom: &Atom{
											Identifier: &tk[1],
											Tokens:     tk[1:2],
										},
										Tokens: tk[1:2],
									},
									Tokens: tk[1:2],
								},
								Tokens: tk[1:2],
							},
							Tokens: tk[1:2],
						},
						Tokens: tk[1:2],
					},
					Tokens: tk[1:2],
				},
				AndExpression: &AndExpression{
					ShiftExpression: ShiftExpression{
						AddExpression: AddExpression{
							MultiplyExpression: MultiplyExpression{
								UnaryExpression: UnaryExpression{
									PowerExpression: &PowerExpression{
										PrimaryExpression: PrimaryExpression{
											Atom: &Atom{
												Identifier: &tk[9],
												Tokens:     tk[9:10],
											},
											Tokens: tk[9:10],
										},
										Tokens: tk[9:10],
									},
									Tokens: tk[9:10],
								},
								Tokens: tk[9:10],
							},
							Tokens: tk[9:10],
						},
						Tokens: tk[9:10],
					},
					Tokens: tk[9:10],
				},
				Comments: [2]Comments{{tk[3]}, {tk[7]}},
				Tokens:   tk[1:10],
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
					},
					Parsing: "ShiftExpression",
					Token:   tk[0],
				},
				Parsing: "AndExpression",
				Token:   tk[0],
			}
		}},
		{`1&nonlocal`, func(t *test, tk Tokens) { // 6
			t.Err = Error{
				Err: Error{
					Err: Error{
						Err: Error{
							Err: Error{
								Err: Error{
									Err: Error{
										Err: Error{
											Err: Error{
												Err: Error{
													Err:     ErrInvalidEnclosure,
													Parsing: "Enclosure",
													Token:   tk[2],
												},
												Parsing: "Atom",
												Token:   tk[2],
											},
											Parsing: "PrimaryExpression",
											Token:   tk[2],
										},
										Parsing: "PowerExpression",
										Token:   tk[2],
									},
									Parsing: "UnaryExpression",
									Token:   tk[2],
								},
								Parsing: "MultiplyExpression",
								Token:   tk[2],
							},
							Parsing: "AddExpression",
							Token:   tk[2],
						},
						Parsing: "ShiftExpression",
						Token:   tk[2],
					},
					Parsing: "AndExpression",
					Token:   tk[2],
				},
				Parsing: "AndExpression",
				Token:   tk[2],
			}
		}},
	}, func(t *test) (Type, error) {
		var ae AndExpression

		if t.Tokens.Peek().Data == "(" {
			t.Tokens.Tokens = t.Tokens.Tokens[1:1]
		}

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
		{"(a # A\n<< # B\nb)", func(t *test, tk Tokens) { // 5
			t.Output = ShiftExpression{
				AddExpression: AddExpression{
					MultiplyExpression: MultiplyExpression{
						UnaryExpression: UnaryExpression{
							PowerExpression: &PowerExpression{
								PrimaryExpression: PrimaryExpression{
									Atom: &Atom{
										Identifier: &tk[1],
										Tokens:     tk[1:2],
									},
									Tokens: tk[1:2],
								},
								Tokens: tk[1:2],
							},
							Tokens: tk[1:2],
						},
						Tokens: tk[1:2],
					},
					Tokens: tk[1:2],
				},
				Shift: &tk[5],
				ShiftExpression: &ShiftExpression{
					AddExpression: AddExpression{
						MultiplyExpression: MultiplyExpression{
							UnaryExpression: UnaryExpression{
								PowerExpression: &PowerExpression{
									PrimaryExpression: PrimaryExpression{
										Atom: &Atom{
											Identifier: &tk[9],
											Tokens:     tk[9:10],
										},
										Tokens: tk[9:10],
									},
									Tokens: tk[9:10],
								},
								Tokens: tk[9:10],
							},
							Tokens: tk[9:10],
						},
						Tokens: tk[9:10],
					},
					Tokens: tk[9:10],
				},
				Comments: [2]Comments{{tk[3]}, {tk[7]}},
				Tokens:   tk[1:10],
			}
		}},
		{`nonlocal`, func(t *test, tk Tokens) { // 6
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
				},
				Parsing: "ShiftExpression",
				Token:   tk[0],
			}
		}},
		{`1<<nonlocal`, func(t *test, tk Tokens) { // 7
			t.Err = Error{
				Err: Error{
					Err: Error{
						Err: Error{
							Err: Error{
								Err: Error{
									Err: Error{
										Err: Error{
											Err: Error{
												Err:     ErrInvalidEnclosure,
												Parsing: "Enclosure",
												Token:   tk[2],
											},
											Parsing: "Atom",
											Token:   tk[2],
										},
										Parsing: "PrimaryExpression",
										Token:   tk[2],
									},
									Parsing: "PowerExpression",
									Token:   tk[2],
								},
								Parsing: "UnaryExpression",
								Token:   tk[2],
							},
							Parsing: "MultiplyExpression",
							Token:   tk[2],
						},
						Parsing: "AddExpression",
						Token:   tk[2],
					},
					Parsing: "ShiftExpression",
					Token:   tk[2],
				},
				Parsing: "ShiftExpression",
				Token:   tk[2],
			}
		}},
	}, func(t *test) (Type, error) {
		var se ShiftExpression

		if t.Tokens.Peek().Data == "(" {
			t.Tokens.Tokens = t.Tokens.Tokens[1:1]
		}

		err := se.parse(t.Tokens)

		return se, err
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
		{"(a # A\n+ # B\nb)", func(t *test, tk Tokens) { // 5
			t.Output = AddExpression{
				MultiplyExpression: MultiplyExpression{
					UnaryExpression: UnaryExpression{
						PowerExpression: &PowerExpression{
							PrimaryExpression: PrimaryExpression{
								Atom: &Atom{
									Identifier: &tk[1],
									Tokens:     tk[1:2],
								},
								Tokens: tk[1:2],
							},
							Tokens: tk[1:2],
						},
						Tokens: tk[1:2],
					},
					Tokens: tk[1:2],
				},
				Add: &tk[5],
				AddExpression: &AddExpression{
					MultiplyExpression: MultiplyExpression{
						UnaryExpression: UnaryExpression{
							PowerExpression: &PowerExpression{
								PrimaryExpression: PrimaryExpression{
									Atom: &Atom{
										Identifier: &tk[9],
										Tokens:     tk[9:10],
									},
									Tokens: tk[9:10],
								},
								Tokens: tk[9:10],
							},
							Tokens: tk[9:10],
						},
						Tokens: tk[9:10],
					},
					Tokens: tk[9:10],
				},
				Comments: [2]Comments{{tk[3]}, {tk[7]}},
				Tokens:   tk[1:10],
			}
		}},
		{`nonlocal`, func(t *test, tk Tokens) { // 6
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
		{`1 + nonlocal`, func(t *test, tk Tokens) { // 7
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

		if t.Tokens.Peek().Data == "(" {
			t.Tokens.Tokens = t.Tokens.Tokens[1:1]
		}

		err := ae.parse(t.Tokens)

		return ae, err
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
		{`a @ b`, func(t *test, tk Tokens) { // 3
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
		{`a//b`, func(t *test, tk Tokens) { // 4
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
		{`a / b`, func(t *test, tk Tokens) { // 5
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
		{`a % b`, func(t *test, tk Tokens) { // 6
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
		{`a / b @ c * d`, func(t *test, tk Tokens) { // 7
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
		{"(a # A\n* # B\nb)", func(t *test, tk Tokens) { // 8
			t.Output = MultiplyExpression{
				UnaryExpression: UnaryExpression{
					PowerExpression: &PowerExpression{
						PrimaryExpression: PrimaryExpression{
							Atom: &Atom{
								Identifier: &tk[1],
								Tokens:     tk[1:2],
							},
							Tokens: tk[1:2],
						},
						Tokens: tk[1:2],
					},
					Tokens: tk[1:2],
				},
				Multiply: &tk[5],
				MultiplyExpression: &MultiplyExpression{
					UnaryExpression: UnaryExpression{
						PowerExpression: &PowerExpression{
							PrimaryExpression: PrimaryExpression{
								Atom: &Atom{
									Identifier: &tk[9],
									Tokens:     tk[9:10],
								},
								Tokens: tk[9:10],
							},
							Tokens: tk[9:10],
						},
						Tokens: tk[9:10],
					},
					Tokens: tk[9:10],
				},
				Comments: [2]Comments{{tk[3]}, {tk[7]}},
				Tokens:   tk[1:10],
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

		if t.Tokens.Peek().Data == "(" {
			t.Tokens.Tokens = t.Tokens.Tokens[1:1]
		}

		err := me.parse(t.Tokens)

		return me, err
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
		{`~-a`, func(t *test, tk Tokens) { // 5
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
		{"(- # A\n1)", func(t *test, tk Tokens) { // 6
			t.Output = UnaryExpression{
				Unary: &tk[1],
				UnaryExpression: &UnaryExpression{
					PowerExpression: &PowerExpression{
						PrimaryExpression: PrimaryExpression{
							Atom: &Atom{
								Literal: &tk[5],
								Tokens:  tk[5:6],
							},
							Tokens: tk[5:6],
						},
						Tokens: tk[5:6],
					},
					Tokens: tk[5:6],
				},
				Comments: Comments{tk[3]},
				Tokens:   tk[1:6],
			}
		}},
		{`nonlocal`, func(t *test, tk Tokens) { // 7
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
		{`+nonlocal`, func(t *test, tk Tokens) { // 8
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

		if t.Tokens.Peek().Data == "(" {
			t.Tokens.Tokens = t.Tokens.Tokens[1:1]
		}

		err := ue.parse(t.Tokens)

		return ue, err
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
		{"(await # A\na)", func(t *test, tk Tokens) { // 6
			t.Output = PowerExpression{
				AwaitExpression: true,
				PrimaryExpression: PrimaryExpression{
					Atom: &Atom{
						Identifier: &tk[5],
						Tokens:     tk[5:6],
					},
					Tokens: tk[5:6],
				},
				Comments: [3]Comments{{tk[3]}},
				Tokens:   tk[1:6],
			}
		}},
		{"(a # A\n** # B\nb)", func(t *test, tk Tokens) { // 7
			t.Output = PowerExpression{
				PrimaryExpression: PrimaryExpression{
					Atom: &Atom{
						Identifier: &tk[1],
						Tokens:     tk[1:2],
					},
					Tokens: tk[1:2],
				},
				UnaryExpression: &UnaryExpression{
					PowerExpression: &PowerExpression{
						PrimaryExpression: PrimaryExpression{
							Atom: &Atom{
								Identifier: &tk[9],
								Tokens:     tk[9:10],
							},
							Tokens: tk[9:10],
						},
						Tokens: tk[9:10],
					},
					Tokens: tk[9:10],
				},
				Comments: [3]Comments{nil, {tk[3]}, {tk[7]}},
				Tokens:   tk[1:10],
			}
		}},
		{"(await # A\na # B\n** # C\nb)", func(t *test, tk Tokens) { // 8
			t.Output = PowerExpression{
				AwaitExpression: true,
				PrimaryExpression: PrimaryExpression{
					Atom: &Atom{
						Identifier: &tk[5],
						Tokens:     tk[5:6],
					},
					Tokens: tk[5:6],
				},
				UnaryExpression: &UnaryExpression{
					PowerExpression: &PowerExpression{
						PrimaryExpression: PrimaryExpression{
							Atom: &Atom{
								Identifier: &tk[13],
								Tokens:     tk[13:14],
							},
							Tokens: tk[13:14],
						},
						Tokens: tk[13:14],
					},
					Tokens: tk[13:14],
				},
				Comments: [3]Comments{{tk[3]}, {tk[7]}, {tk[11]}},
				Tokens:   tk[1:14],
			}
		}},
		{`nonlocal`, func(t *test, tk Tokens) { // 9
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
		{`1 ** nonlocal`, func(t *test, tk Tokens) { // 10
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

		if t.Tokens.Peek().Data == "(" {
			t.Tokens.Tokens = t.Tokens.Tokens[1:1]
		}

		err := pe.parse(t.Tokens)

		return pe, err
	})
}

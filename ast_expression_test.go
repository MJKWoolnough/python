package python

import (
	"testing"
)

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
					Parsing: "SliceList",
					Token:   tk[2],
				},
				Parsing: "PrimaryExpression",
				Token:   tk[1],
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
				Token:   tk[1],
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
					Tokens: tk[1:3],
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
				Token:   tk[1],
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
				Token:   tk[1],
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

		err := pe.parse(t.Tokens)

		return pe, err
	})
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
					Expression: &Expression{
						ConditionalExpression: WrapConditional(&Atom{
							Identifier: &tk[3],
							Tokens:     tk[3:4],
						}),
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
					Expression: &Expression{
						ConditionalExpression: WrapConditional(&Atom{
							Identifier: &tk[4],
							Tokens:     tk[4:5],
						}),
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
				ListDisplay: &StarredListOrComprehension{
					Tokens: tk[1:1],
				},
				Tokens: tk[:2],
			}
		}},
		{`[ ]`, func(t *test, tk Tokens) { // 10
			t.Output = Enclosure{
				ListDisplay: &StarredListOrComprehension{
					Tokens: tk[2:2],
				},
				Tokens: tk[:3],
			}
		}},
		{`[a]`, func(t *test, tk Tokens) { // 11
			t.Output = Enclosure{
				ListDisplay: &StarredListOrComprehension{
					StarredList: &StarredList{
						StarredItems: []StarredItem{
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
				ListDisplay: &StarredListOrComprehension{
					StarredList: &StarredList{
						StarredItems: []StarredItem{
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
				SetDisplay: &StarredListOrComprehension{
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
				Tokens: tk[:4],
			}
		}},
		{`{ *a }`, func(t *test, tk Tokens) { // 18
			t.Output = Enclosure{
				SetDisplay: &StarredListOrComprehension{
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
				Tokens: tk[:6],
			}
		}},
		{`{a:=b}`, func(t *test, tk Tokens) { // 19
			t.Output = Enclosure{
				SetDisplay: &StarredListOrComprehension{
					StarredList: &StarredList{
						StarredItems: []StarredItem{
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
				SetDisplay: &StarredListOrComprehension{
					StarredList: &StarredList{
						StarredItems: []StarredItem{
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
						Err: wrapConditionalExpressionError(Error{
							Err: Error{
								Err:     ErrInvalidEnclosure,
								Parsing: "Enclosure",
								Token:   tk[3],
							},
							Parsing: "Atom",
							Token:   tk[3],
						}),
						Parsing: "Expression",
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
							Err: Error{
								Err:     ErrInvalidEnclosure,
								Parsing: "Enclosure",
								Token:   tk[1],
							},
							Parsing: "Atom",
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
						Err: Error{
							Err: Error{
								Err: wrapConditionalExpressionError(Error{
									Err: Error{
										Err:     ErrInvalidEnclosure,
										Parsing: "Enclosure",
										Token:   tk[1],
									},
									Parsing: "Atom",
									Token:   tk[1],
								}),
								Parsing: "Expression",
								Token:   tk[1],
							},
							Parsing: "AssignmentExpression",
							Token:   tk[1],
						},
						Parsing: "StarredItem",
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
										Err: Error{
											Err:     ErrInvalidEnclosure,
											Parsing: "Enclosure",
											Token:   tk[1],
										},
										Parsing: "Atom",
										Token:   tk[1],
									}),
									Parsing: "Expression",
									Token:   tk[1],
								},
								Parsing: "AssignmentExpression",
								Token:   tk[1],
							},
							Parsing: "StarredItem",
							Token:   tk[1],
						},
						Parsing: "StarredList",
						Token:   tk[1],
					},
					Parsing: "StarredListOrComprehension",
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
							Err: Error{
								Err:     ErrInvalidEnclosure,
								Parsing: "Enclosure",
								Token:   tk[1],
							},
							Parsing: "Atom",
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
								Err: Error{
									Err:     ErrInvalidEnclosure,
									Parsing: "Enclosure",
									Token:   tk[3],
								},
								Parsing: "Atom",
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
					Parsing: "StarredListOrComprehension",
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
	}, func(t *test) (Type, error) {
		var e Enclosure

		err := e.parse(t.Tokens)

		return e, err
	})
}

func TestStarredListOrComprehension(t *testing.T) {
	doTests(t, []sourceFn{
		{`a`, func(t *test, tk Tokens) { // 1
			t.Output = StarredListOrComprehension{
				StarredList: &StarredList{
					StarredItems: []StarredItem{
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
			t.Output = StarredListOrComprehension{
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
			t.Output = StarredListOrComprehension{
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
			t.Output = StarredListOrComprehension{
				StarredList: &StarredList{
					StarredItems: []StarredItem{
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
			t.Output = StarredListOrComprehension{
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
			t.Output = StarredListOrComprehension{
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
									Err: Error{
										Err:     ErrInvalidEnclosure,
										Parsing: "Enclosure",
										Token:   tk[0],
									},
									Parsing: "Atom",
									Token:   tk[0],
								}),
								Parsing: "Expression",
								Token:   tk[0],
							},
							Parsing: "AssignmentExpression",
							Token:   tk[0],
						},
						Parsing: "StarredItem",
						Token:   tk[0],
					},
					Parsing: "StarredList",
					Token:   tk[0],
				},
				Parsing: "StarredListOrComprehension",
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
				Parsing: "StarredListOrComprehension",
				Token:   tk[0],
			}
		}},
	}, func(t *test) (Type, error) {
		var s StarredListOrComprehension

		err := s.parse(t.Tokens, t.AssignmentExpression)

		return s, err
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
							Err: Error{
								Err:     ErrInvalidEnclosure,
								Parsing: "Enclosure",
								Token:   tk[0],
							},
							Parsing: "Atom",
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
		{`a for nonlocal in c`, func(t *test, tk Tokens) { // 3
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
		{`async a in b if c`, func(t *test, tk Tokens) { // 6
			t.Err = Error{
				Err:     ErrMissingFor,
				Parsing: "ComprehensionFor",
				Token:   tk[2],
			}
		}},
		{`for nonlocal in a if b`, func(t *test, tk Tokens) { // 7
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
		{`for a in nonlocal if b`, func(t *test, tk Tokens) { // 7
			t.Err = Error{
				Err: wrapConditionalExpressionError(Error{
					Err: Error{
						Err:     ErrInvalidEnclosure,
						Parsing: "Enclosure",
						Token:   tk[6],
					},
					Parsing: "Atom",
					Token:   tk[6],
				}).Err,
				Parsing: "ComprehensionFor",
				Token:   tk[6],
			}
		}},
	}, func(t *test) (Type, error) {
		var c ComprehensionFor

		err := c.parse(t.Tokens)

		return c, err
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
									Err: Error{
										Err:     ErrInvalidEnclosure,
										Parsing: "Enclosure",
										Token:   tk[1],
									},
									Parsing: "Atom",
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
								Err: Error{
									Err:     ErrInvalidEnclosure,
									Parsing: "Enclosure",
									Token:   tk[1],
								},
								Parsing: "Atom",
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
		{`[nonlocal]`, func(t *test, tk Tokens) { // 8
			t.Err = Error{
				Err: Error{
					Err: Error{
						Err: wrapConditionalExpressionError(Error{
							Err: Error{
								Err:     ErrInvalidEnclosure,
								Parsing: "Enclosure",
								Token:   tk[1],
							},
							Parsing: "Atom",
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
		{`[a b]`, func(t *test, tk Tokens) { // 8
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
		{`nonlocal`, func(t *test, tk Tokens) { // 4
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
		{`1|nonlocal`, func(t *test, tk Tokens) { // 5
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
		{`nonlocal`, func(t *test, tk Tokens) { // 4
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
		{`1^nonlocal`, func(t *test, tk Tokens) { // 5
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
		{`nonlocal`, func(t *test, tk Tokens) { // 4
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
		{`1&nonlocal`, func(t *test, tk Tokens) { // 5
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
		{`nonlocal`, func(t *test, tk Tokens) { // 5
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
		{`1<<nonlocal`, func(t *test, tk Tokens) { // 6
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

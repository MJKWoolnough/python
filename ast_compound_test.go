package python

import "testing"

func TestStarredItem(t *testing.T) {
	doTests(t, []sourceFn{
		{`a`, func(t *test, tk Tokens) { // 1
			t.Output = StarredItem{
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
			t.Output = StarredItem{
				OrExpr: &WrapConditional(&Atom{
					Identifier: &tk[1],
					Tokens:     tk[1:2],
				}).OrTest.AndTest.NotTest.Comparison.OrExpression,
				Tokens: tk[:2],
			}
		}},
		{`* a`, func(t *test, tk Tokens) { // 3
			t.Output = StarredItem{
				OrExpr: &WrapConditional(&Atom{
					Identifier: &tk[2],
					Tokens:     tk[2:3],
				}).OrTest.AndTest.NotTest.Comparison.OrExpression,
				Tokens: tk[:3],
			}
		}},
		{`nonlocal`, func(t *test, tk Tokens) { // 4
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
				Parsing: "StarredItem",
				Token:   tk[0],
			}
		}},
		{`*nonlocal`, func(t *test, tk Tokens) { // 5
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
				Parsing: "StarredItem",
				Token:   tk[1],
			}
		}},
	}, func(t *test) (Type, error) {
		var s StarredItem

		err := s.parse(t.Tokens)

		return s, err
	})
}

func TestTypeParam(t *testing.T) {
	doTests(t, []sourceFn{
		{`a`, func(t *test, tk Tokens) { // 1
			t.Output = TypeParam{
				Identifier: &tk[0],
				Tokens:     tk[:1],
			}
		}},
		{`a:b`, func(t *test, tk Tokens) { // 2
			t.Output = TypeParam{
				Identifier: &tk[0],
				Expression: &Expression{
					ConditionalExpression: WrapConditional(&Atom{
						Identifier: &tk[2],
						Tokens:     tk[2:3],
					}),
					Tokens: tk[2:3],
				},
				Tokens: tk[:3],
			}
		}},
		{`a : b`, func(t *test, tk Tokens) { // 3
			t.Output = TypeParam{
				Identifier: &tk[0],
				Expression: &Expression{
					ConditionalExpression: WrapConditional(&Atom{
						Identifier: &tk[4],
						Tokens:     tk[4:5],
					}),
					Tokens: tk[4:5],
				},
				Tokens: tk[:5],
			}
		}},
		{`*a`, func(t *test, tk Tokens) { // 4
			t.Output = TypeParam{
				Type:       TypeParamVar,
				Identifier: &tk[1],
				Tokens:     tk[:2],
			}
		}},
		{`* a`, func(t *test, tk Tokens) { // 5
			t.Output = TypeParam{
				Type:       TypeParamVar,
				Identifier: &tk[2],
				Tokens:     tk[:3],
			}
		}},
		{`**a`, func(t *test, tk Tokens) { // 6
			t.Output = TypeParam{
				Type:       TypeParamVarTuple,
				Identifier: &tk[1],
				Tokens:     tk[:2],
			}
		}},
		{`** a`, func(t *test, tk Tokens) { // 7
			t.Output = TypeParam{
				Type:       TypeParamVarTuple,
				Identifier: &tk[2],
				Tokens:     tk[:3],
			}
		}},
		{`nonlocal`, func(t *test, tk Tokens) { // 8
			t.Err = Error{
				Err:     ErrMissingIdentifier,
				Parsing: "TypeParam",
				Token:   tk[0],
			}
		}},
		{`a:nonlocal`, func(t *test, tk Tokens) { // 9
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
				Parsing: "TypeParam",
				Token:   tk[2],
			}
		}},
		{`*nonlocal`, func(t *test, tk Tokens) { // 10
			t.Err = Error{
				Err:     ErrMissingIdentifier,
				Parsing: "TypeParam",
				Token:   tk[1],
			}
		}},
		{`**nonlocal`, func(t *test, tk Tokens) { // 11
			t.Err = Error{
				Err:     ErrMissingIdentifier,
				Parsing: "TypeParam",
				Token:   tk[1],
			}
		}},
	}, func(t *test) (Type, error) {
		var tp TypeParam

		err := tp.parse(t.Tokens)

		return tp, err
	})
}

func TestParameter(t *testing.T) {
	doTests(t, []sourceFn{
		{`a`, func(t *test, tk Tokens) { // 1
			t.Output = Parameter{
				Identifier: &tk[0],
				Tokens:     tk[:1],
			}
		}},
		{`a`, func(t *test, tk Tokens) { // 2
			t.AllowTypeAnnotations = true
			t.Output = Parameter{
				Identifier: &tk[0],
				Tokens:     tk[:1],
			}
		}},
		{`a:b`, func(t *test, tk Tokens) { // 3
			t.Output = Parameter{
				Identifier: &tk[0],
				Tokens:     tk[:1],
			}
		}},
		{`a:b`, func(t *test, tk Tokens) { // 4
			t.AllowTypeAnnotations = true
			t.Output = Parameter{
				Identifier: &tk[0],
				Type: &Expression{
					ConditionalExpression: WrapConditional(&Atom{
						Identifier: &tk[2],
						Tokens:     tk[2:3],
					}),
					Tokens: tk[2:3],
				},
				Tokens: tk[:3],
			}
		}},
		{`a : b`, func(t *test, tk Tokens) { // 5
			t.AllowTypeAnnotations = true
			t.Output = Parameter{
				Identifier: &tk[0],
				Type: &Expression{
					ConditionalExpression: WrapConditional(&Atom{
						Identifier: &tk[4],
						Tokens:     tk[4:5],
					}),
					Tokens: tk[4:5],
				},
				Tokens: tk[:5],
			}
		}},
	}, func(t *test) (Type, error) {
		var p Parameter

		err := p.parse(t.Tokens, t.AllowTypeAnnotations)

		return p, err
	})
}

func TestArgumentList(t *testing.T) {
	doTests(t, []sourceFn{
		{`a`, func(t *test, tk Tokens) { // 1
			t.Output = ArgumentList{
				PositionalArguments: []PositionalArgument{
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
		{`a,`, func(t *test, tk Tokens) { // 2
			t.Output = ArgumentList{
				PositionalArguments: []PositionalArgument{
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
		{`a,b`, func(t *test, tk Tokens) { // 3
			t.Output = ArgumentList{
				PositionalArguments: []PositionalArgument{
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
		{`a , b`, func(t *test, tk Tokens) { // 4
			t.Output = ArgumentList{
				PositionalArguments: []PositionalArgument{
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
									Identifier: &tk[4],
									Tokens:     tk[4:5],
								}),
								Tokens: tk[4:5],
							},
							Tokens: tk[4:5],
						},
						Tokens: tk[4:5],
					},
				},
				Tokens: tk[:5],
			}
		}},
		{`a,*b`, func(t *test, tk Tokens) { // 5
			t.Output = ArgumentList{
				PositionalArguments: []PositionalArgument{
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
						Expression: &Expression{
							ConditionalExpression: WrapConditional(&Atom{
								Identifier: &tk[3],
								Tokens:     tk[3:4],
							}),
							Tokens: tk[3:4],
						},
						Tokens: tk[2:4],
					},
				},
				Tokens: tk[:4],
			}
		}},
		{`a, *b`, func(t *test, tk Tokens) { // 6
			t.Output = ArgumentList{
				PositionalArguments: []PositionalArgument{
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
						Expression: &Expression{
							ConditionalExpression: WrapConditional(&Atom{
								Identifier: &tk[4],
								Tokens:     tk[4:5],
							}),
							Tokens: tk[4:5],
						},
						Tokens: tk[3:5],
					},
				},
				Tokens: tk[:5],
			}
		}},
		{`a,b=c`, func(t *test, tk Tokens) { // 7
			t.Output = ArgumentList{
				PositionalArguments: []PositionalArgument{
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
				StarredAndKeywordArguments: []StarredOrKeyword{
					{
						KeywordItem: &KeywordItem{
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
				Tokens: tk[:5],
			}
		}},
		{`a, b=c, *d`, func(t *test, tk Tokens) { // 8
			t.Output = ArgumentList{
				PositionalArguments: []PositionalArgument{
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
				StarredAndKeywordArguments: []StarredOrKeyword{
					{
						KeywordItem: &KeywordItem{
							Identifier: &tk[3],
							Expression: Expression{
								ConditionalExpression: WrapConditional(&Atom{
									Identifier: &tk[5],
									Tokens:     tk[5:6],
								}),
								Tokens: tk[5:6],
							},
							Tokens: tk[3:6],
						},
						Tokens: tk[3:6],
					},
					{
						Expression: &Expression{
							ConditionalExpression: WrapConditional(&Atom{
								Identifier: &tk[9],
								Tokens:     tk[9:10],
							}),
							Tokens: tk[9:10],
						},
						Tokens: tk[8:10],
					},
				},
				Tokens: tk[:10],
			}
		}},
		{`a=b,*c`, func(t *test, tk Tokens) { // 9
			t.Output = ArgumentList{
				StarredAndKeywordArguments: []StarredOrKeyword{
					{
						KeywordItem: &KeywordItem{
							Identifier: &tk[0],
							Expression: Expression{
								ConditionalExpression: WrapConditional(&Atom{
									Identifier: &tk[2],
									Tokens:     tk[2:3],
								}),
								Tokens: tk[2:3],
							},
							Tokens: tk[:3],
						},
						Tokens: tk[:3],
					},
					{
						Expression: &Expression{
							ConditionalExpression: WrapConditional(&Atom{
								Identifier: &tk[5],
								Tokens:     tk[5:6],
							}),
							Tokens: tk[5:6],
						},
						Tokens: tk[4:6],
					},
				},
				Tokens: tk[:6],
			}
		}},
		{`a,b=c,**d`, func(t *test, tk Tokens) { // 10
			t.Output = ArgumentList{
				PositionalArguments: []PositionalArgument{
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
				StarredAndKeywordArguments: []StarredOrKeyword{
					{
						KeywordItem: &KeywordItem{
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
				KeywordArguments: []KeywordArgument{
					{
						Expression: &Expression{
							ConditionalExpression: WrapConditional(&Atom{
								Identifier: &tk[7],
								Tokens:     tk[7:8],
							}),
							Tokens: tk[7:8],
						},
						Tokens: tk[6:8],
					},
				},
				Tokens: tk[:8],
			}
		}},
		{`a,**b,c=d`, func(t *test, tk Tokens) { // 11
			t.Output = ArgumentList{
				PositionalArguments: []PositionalArgument{
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
				KeywordArguments: []KeywordArgument{
					{
						Expression: &Expression{
							ConditionalExpression: WrapConditional(&Atom{
								Identifier: &tk[3],
								Tokens:     tk[3:4],
							}),
							Tokens: tk[3:4],
						},
						Tokens: tk[2:4],
					},
					{
						KeywordItem: &KeywordItem{
							Identifier: &tk[5],
							Expression: Expression{
								ConditionalExpression: WrapConditional(&Atom{
									Identifier: &tk[7],
									Tokens:     tk[7:8],
								}),
								Tokens: tk[7:8],
							},
							Tokens: tk[5:8],
						},
						Tokens: tk[5:8],
					},
				},
				Tokens: tk[:8],
			}
		}},
		{`**b`, func(t *test, tk Tokens) { // 12
			t.Output = ArgumentList{
				KeywordArguments: []KeywordArgument{
					{
						Expression: &Expression{
							ConditionalExpression: WrapConditional(&Atom{
								Identifier: &tk[1],
								Tokens:     tk[1:2],
							}),
							Tokens: tk[1:2],
						},
						Tokens: tk[:2],
					},
				},
				Tokens: tk[:2],
			}
		}},
		{`nonlocal`, func(t *test, tk Tokens) { // 13
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
					Parsing: "PositionalArgument",
					Token:   tk[0],
				},
				Parsing: "ArgumentList",
				Token:   tk[0],
			}
		}},
		{`a b`, func(t *test, tk Tokens) { // 14
			t.Err = Error{
				Err:     ErrMissingComma,
				Parsing: "ArgumentList",
				Token:   tk[1],
			}
		}},
		{`a=nonlocal`, func(t *test, tk Tokens) { // 15
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
						Parsing: "KeywordItem",
						Token:   tk[2],
					},
					Parsing: "StarredOrKeyword",
					Token:   tk[0],
				},
				Parsing: "ArgumentList",
				Token:   tk[0],
			}
		}},
		{`a=b c`, func(t *test, tk Tokens) { // 16
			t.Err = Error{
				Err:     ErrMissingComma,
				Parsing: "ArgumentList",
				Token:   tk[3],
			}
		}},
		{`**nonlocal`, func(t *test, tk Tokens) { // 17
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
					Parsing: "KeywordArgument",
					Token:   tk[1],
				},
				Parsing: "ArgumentList",
				Token:   tk[0],
			}
		}},
		{`**a b`, func(t *test, tk Tokens) { // 16
			t.Err = Error{
				Err:     ErrMissingComma,
				Parsing: "ArgumentList",
				Token:   tk[2],
			}
		}},
	}, func(t *test) (Type, error) {
		var a ArgumentList

		err := a.parse(t.Tokens)

		return a, err
	})
}

func TestPositionalArgument(t *testing.T) {
	doTests(t, []sourceFn{
		{`a`, func(t *test, tk Tokens) { // 1
			t.Output = PositionalArgument{
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
			t.Output = PositionalArgument{
				Expression: &Expression{
					ConditionalExpression: WrapConditional(&Atom{
						Identifier: &tk[1],
						Tokens:     tk[1:2],
					}),
					Tokens: tk[1:2],
				},
				Tokens: tk[:2],
			}
		}},
		{`* a`, func(t *test, tk Tokens) { // 3
			t.Output = PositionalArgument{
				Expression: &Expression{
					ConditionalExpression: WrapConditional(&Atom{
						Identifier: &tk[2],
						Tokens:     tk[2:3],
					}),
					Tokens: tk[2:3],
				},
				Tokens: tk[:3],
			}
		}},
		{`nonlocal`, func(t *test, tk Tokens) { // 4
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
				Parsing: "PositionalArgument",
				Token:   tk[0],
			}
		}},
		{`*nonlocal`, func(t *test, tk Tokens) { // 5
			t.Err = Error{
				Err: Error{
					Err: wrapConditionalExpressionError(Error{
						Err:     ErrInvalidEnclosure,
						Parsing: "Enclosure",
						Token:   tk[1],
					}),
					Parsing: "Expression",
					Token:   tk[1],
				},
				Parsing: "PositionalArgument",
				Token:   tk[1],
			}
		}},
	}, func(t *test) (Type, error) {
		var p PositionalArgument

		err := p.parse(t.Tokens)

		return p, err
	})
}

func TestStarredOrKeyword(t *testing.T) {
	doTests(t, []sourceFn{
		{`*a`, func(t *test, tk Tokens) { // 1
			t.Output = StarredOrKeyword{
				Expression: &Expression{
					ConditionalExpression: WrapConditional(&Atom{
						Identifier: &tk[1],
						Tokens:     tk[1:2],
					}),
					Tokens: tk[1:2],
				},
				Tokens: tk[:2],
			}
		}},
		{`* a`, func(t *test, tk Tokens) { // 2
			t.Output = StarredOrKeyword{
				Expression: &Expression{
					ConditionalExpression: WrapConditional(&Atom{
						Identifier: &tk[2],
						Tokens:     tk[2:3],
					}),
					Tokens: tk[2:3],
				},
				Tokens: tk[:3],
			}
		}},
		{`a=b`, func(t *test, tk Tokens) { // 3
			t.Output = StarredOrKeyword{
				KeywordItem: &KeywordItem{
					Identifier: &tk[0],
					Expression: Expression{
						ConditionalExpression: WrapConditional(&Atom{
							Identifier: &tk[2],
							Tokens:     tk[2:3],
						}),
						Tokens: tk[2:3],
					},
					Tokens: tk[:3],
				},
				Tokens: tk[:3],
			}
		}},
		{`*nonlocal`, func(t *test, tk Tokens) { // 4
			t.Err = Error{
				Err: Error{
					Err: wrapConditionalExpressionError(Error{
						Err:     ErrInvalidEnclosure,
						Parsing: "Enclosure",
						Token:   tk[1],
					}),
					Parsing: "Expression",
					Token:   tk[1],
				},
				Parsing: "StarredOrKeyword",
				Token:   tk[1],
			}
		}},
		{`nonlocal`, func(t *test, tk Tokens) { // 5
			t.Err = Error{
				Err: Error{
					Err:     ErrMissingIdentifier,
					Parsing: "KeywordItem",
					Token:   tk[0],
				},
				Parsing: "StarredOrKeyword",
				Token:   tk[0],
			}
		}},
	}, func(t *test) (Type, error) {
		var s StarredOrKeyword

		err := s.parse(t.Tokens)

		return s, err
	})
}

func TestKeywordItem(t *testing.T) {
	doTests(t, []sourceFn{
		{`a=b`, func(t *test, tk Tokens) { // 1
			t.Output = KeywordItem{
				Identifier: &tk[0],
				Expression: Expression{
					ConditionalExpression: WrapConditional(&Atom{
						Identifier: &tk[2],
						Tokens:     tk[2:3],
					}),
					Tokens: tk[2:3],
				},
				Tokens: tk[:3],
			}
		}},
		{`a = b`, func(t *test, tk Tokens) { // 2
			t.Output = KeywordItem{
				Identifier: &tk[0],
				Expression: Expression{
					ConditionalExpression: WrapConditional(&Atom{
						Identifier: &tk[4],
						Tokens:     tk[4:5],
					}),
					Tokens: tk[4:5],
				},
				Tokens: tk[:5],
			}
		}},
		{`nonlocal=a`, func(t *test, tk Tokens) { // 3
			t.Err = Error{
				Err:     ErrMissingIdentifier,
				Parsing: "KeywordItem",
				Token:   tk[0],
			}
		}},
		{`a:b`, func(t *test, tk Tokens) { // 4
			t.Err = Error{
				Err:     ErrMissingEquals,
				Parsing: "KeywordItem",
				Token:   tk[1],
			}
		}},
		{`a=nonlocal`, func(t *test, tk Tokens) { // 5
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
				Parsing: "KeywordItem",
				Token:   tk[2],
			}
		}},
	}, func(t *test) (Type, error) {
		var k KeywordItem

		err := k.parse(t.Tokens)

		return k, err
	})
}

func TestKeywordArgument(t *testing.T) {
	doTests(t, []sourceFn{
		{`a=b`, func(t *test, tk Tokens) { // 1
			t.Output = KeywordArgument{
				KeywordItem: &KeywordItem{
					Identifier: &tk[0],
					Expression: Expression{
						ConditionalExpression: WrapConditional(&Atom{
							Identifier: &tk[2],
							Tokens:     tk[2:3],
						}),
						Tokens: tk[2:3],
					},
					Tokens: tk[:3],
				},
				Tokens: tk[:3],
			}
		}},
		{`**a`, func(t *test, tk Tokens) { // 2
			t.Output = KeywordArgument{
				Expression: &Expression{
					ConditionalExpression: WrapConditional(&Atom{
						Identifier: &tk[1],
						Tokens:     tk[1:2],
					}),
					Tokens: tk[1:2],
				},
				Tokens: tk[:2],
			}
		}},
		{`** a`, func(t *test, tk Tokens) { // 3
			t.Output = KeywordArgument{
				Expression: &Expression{
					ConditionalExpression: WrapConditional(&Atom{
						Identifier: &tk[2],
						Tokens:     tk[2:3],
					}),
					Tokens: tk[2:3],
				},
				Tokens: tk[:3],
			}
		}},
		{`nonlocal`, func(t *test, tk Tokens) { // 4
			t.Err = Error{
				Err: Error{
					Err:     ErrMissingIdentifier,
					Parsing: "KeywordItem",
					Token:   tk[0],
				},
				Parsing: "KeywordArgument",
				Token:   tk[0],
			}
		}},
		{`**nonlocal`, func(t *test, tk Tokens) { // 5
			t.Err = Error{
				Err: Error{
					Err: wrapConditionalExpressionError(Error{
						Err:     ErrInvalidEnclosure,
						Parsing: "Enclosure",
						Token:   tk[1],
					}),
					Parsing: "Expression",
					Token:   tk[1],
				},
				Parsing: "KeywordArgument",
				Token:   tk[1],
			}
		}},
	}, func(t *test) (Type, error) {
		var k KeywordArgument

		err := k.parse(t.Tokens)

		return k, err
	})
}

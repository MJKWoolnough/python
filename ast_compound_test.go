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

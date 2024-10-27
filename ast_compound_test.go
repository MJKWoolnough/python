package python

import "testing"

func TestCompoundStatement(t *testing.T) {
	doTests(t, []sourceFn{
		{"@a\ndef b():c", func(t *test, tk Tokens) { // 1
			t.Output = CompoundStatement{
				Func: &FuncDefinition{
					Decorators: &Decorators{
						Decorators: []AssignmentExpression{
							{
								Expression: Expression{
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
					},
					FuncName: &tk[5],
					ParameterList: ParameterList{
						Tokens: tk[7:7],
					},
					Suite: Suite{
						StatementList: &StatementList{
							Statements: []SimpleStatement{
								{
									Type: StatementAssignment,
									AssignmentStatement: &AssignmentStatement{
										StarredExpression: &StarredExpression{
											OrExpr: WrapConditional(&Atom{
												Identifier: &tk[9],
												Tokens:     tk[9:10],
											}).OrTest.AndTest.NotTest.Comparison.OrExpression,
											Tokens: tk[9:10],
										},
										Tokens: tk[9:10],
									},
									Tokens: tk[9:10],
								},
							},
							Tokens: tk[9:10],
						},
						Tokens: tk[9:10],
					},
					Tokens: tk[:10],
				},
				Tokens: tk[:10],
			}
		}},
		{"@a\nclass b():c", func(t *test, tk Tokens) { // 2
			t.Output = CompoundStatement{
				Class: &ClassDefinition{
					Decorators: &Decorators{
						Decorators: []AssignmentExpression{
							{
								Expression: Expression{
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
					},
					ClassName: &tk[5],
					Inheritance: &ArgumentList{
						Tokens: tk[7:7],
					},
					Suite: Suite{
						StatementList: &StatementList{
							Statements: []SimpleStatement{
								{
									Type: StatementAssignment,
									AssignmentStatement: &AssignmentStatement{
										StarredExpression: &StarredExpression{
											OrExpr: WrapConditional(&Atom{
												Identifier: &tk[9],
												Tokens:     tk[9:10],
											}).OrTest.AndTest.NotTest.Comparison.OrExpression,
											Tokens: tk[9:10],
										},
										Tokens: tk[9:10],
									},
									Tokens: tk[9:10],
								},
							},
							Tokens: tk[9:10],
						},
						Tokens: tk[9:10],
					},
					Tokens: tk[:10],
				},
				Tokens: tk[:10],
			}
		}},
		{"@a\nasync def b():c", func(t *test, tk Tokens) { // 3
			t.Output = CompoundStatement{
				Func: &FuncDefinition{
					Decorators: &Decorators{
						Decorators: []AssignmentExpression{
							{
								Expression: Expression{
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
					},
					Async:    true,
					FuncName: &tk[7],
					ParameterList: ParameterList{
						Tokens: tk[9:9],
					},
					Suite: Suite{
						StatementList: &StatementList{
							Statements: []SimpleStatement{
								{
									Type: StatementAssignment,
									AssignmentStatement: &AssignmentStatement{
										StarredExpression: &StarredExpression{
											OrExpr: WrapConditional(&Atom{
												Identifier: &tk[11],
												Tokens:     tk[11:12],
											}).OrTest.AndTest.NotTest.Comparison.OrExpression,
											Tokens: tk[11:12],
										},
										Tokens: tk[11:12],
									},
									Tokens: tk[11:12],
								},
							},
							Tokens: tk[11:12],
						},
						Tokens: tk[11:12],
					},
					Tokens: tk[:12],
				},
				Tokens: tk[:12],
			}
		}},
		{"if a:b", func(t *test, tk Tokens) { // 4
			t.Output = CompoundStatement{
				If: &IfStatement{
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
					Suite: Suite{
						StatementList: &StatementList{
							Statements: []SimpleStatement{
								{
									Type: StatementAssignment,
									AssignmentStatement: &AssignmentStatement{
										StarredExpression: &StarredExpression{
											OrExpr: WrapConditional(&Atom{
												Identifier: &tk[4],
												Tokens:     tk[4:5],
											}).OrTest.AndTest.NotTest.Comparison.OrExpression,
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
					Tokens: tk[:5],
				},
				Tokens: tk[:5],
			}
		}},
		{"while a:b", func(t *test, tk Tokens) { // 5
			t.Output = CompoundStatement{
				While: &WhileStatement{
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
					Suite: Suite{
						StatementList: &StatementList{
							Statements: []SimpleStatement{
								{
									Type: StatementAssignment,
									AssignmentStatement: &AssignmentStatement{
										StarredExpression: &StarredExpression{
											OrExpr: WrapConditional(&Atom{
												Identifier: &tk[4],
												Tokens:     tk[4:5],
											}).OrTest.AndTest.NotTest.Comparison.OrExpression,
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
					Tokens: tk[:5],
				},
				Tokens: tk[:5],
			}
		}},
		{"for a in b:c", func(t *test, tk Tokens) { // 6
			t.Output = CompoundStatement{
				For: &ForStatement{
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
					StarredList: StarredList{
						StarredItems: []StarredItem{
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
					Suite: Suite{
						StatementList: &StatementList{
							Statements: []SimpleStatement{
								{
									Type: StatementAssignment,
									AssignmentStatement: &AssignmentStatement{
										StarredExpression: &StarredExpression{
											OrExpr: WrapConditional(&Atom{
												Identifier: &tk[8],
												Tokens:     tk[8:9],
											}).OrTest.AndTest.NotTest.Comparison.OrExpression,
											Tokens: tk[8:9],
										},
										Tokens: tk[8:9],
									},
									Tokens: tk[8:9],
								},
							},
							Tokens: tk[8:9],
						},
						Tokens: tk[8:9],
					},
					Tokens: tk[:9],
				},
				Tokens: tk[:9],
			}
		}},
		{"try:a\nexcept b:c", func(t *test, tk Tokens) { // 7
			t.Output = CompoundStatement{
				Try: &TryStatement{
					Try: Suite{
						StatementList: &StatementList{
							Statements: []SimpleStatement{
								{
									Type: StatementAssignment,
									AssignmentStatement: &AssignmentStatement{
										StarredExpression: &StarredExpression{
											OrExpr: WrapConditional(&Atom{
												Identifier: &tk[2],
												Tokens:     tk[2:3],
											}).OrTest.AndTest.NotTest.Comparison.OrExpression,
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
					Except: []Except{
						{
							Expression: Expression{
								ConditionalExpression: WrapConditional(&Atom{
									Identifier: &tk[6],
									Tokens:     tk[6:7],
								}),
								Tokens: tk[6:7],
							},
							Suite: Suite{
								StatementList: &StatementList{
									Statements: []SimpleStatement{
										{
											Type: StatementAssignment,
											AssignmentStatement: &AssignmentStatement{
												StarredExpression: &StarredExpression{
													OrExpr: WrapConditional(&Atom{
														Identifier: &tk[8],
														Tokens:     tk[8:9],
													}).OrTest.AndTest.NotTest.Comparison.OrExpression,
													Tokens: tk[8:9],
												},
												Tokens: tk[8:9],
											},
											Tokens: tk[8:9],
										},
									},
									Tokens: tk[8:9],
								},
								Tokens: tk[8:9],
							},
							Tokens: tk[6:9],
						},
					},
					Tokens: tk[:9],
				},
				Tokens: tk[:9],
			}
		}},
		{`with a:b`, func(t *test, tk Tokens) { // 8
			t.Output = CompoundStatement{
				With: &WithStatement{
					Contents: WithStatementContents{
						Items: []WithItem{
							{
								Expression: Expression{
									ConditionalExpression: WrapConditional(&Atom{
										Identifier: &tk[2],
										Tokens:     tk[2:3],
									}),
									Tokens: tk[2:3],
								},
								Tokens: tk[2:3],
							},
						},
						Tokens: tk[2:3],
					},
					Suite: Suite{
						StatementList: &StatementList{
							Statements: []SimpleStatement{
								{
									Type: StatementAssignment,
									AssignmentStatement: &AssignmentStatement{
										StarredExpression: &StarredExpression{
											OrExpr: WrapConditional(&Atom{
												Identifier: &tk[4],
												Tokens:     tk[4:5],
											}).OrTest.AndTest.NotTest.Comparison.OrExpression,
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
					Tokens: tk[:5],
				},
				Tokens: tk[:5],
			}
		}},
		{`def a():b`, func(t *test, tk Tokens) { // 9
			t.Output = CompoundStatement{
				Func: &FuncDefinition{
					FuncName: &tk[2],
					ParameterList: ParameterList{
						Tokens: tk[4:4],
					},
					Suite: Suite{
						StatementList: &StatementList{
							Statements: []SimpleStatement{
								{
									Type: StatementAssignment,
									AssignmentStatement: &AssignmentStatement{
										StarredExpression: &StarredExpression{
											OrExpr: WrapConditional(&Atom{
												Identifier: &tk[6],
												Tokens:     tk[6:7],
											}).OrTest.AndTest.NotTest.Comparison.OrExpression,
											Tokens: tk[6:7],
										},
										Tokens: tk[6:7],
									},
									Tokens: tk[6:7],
								},
							},
							Tokens: tk[6:7],
						},
						Tokens: tk[6:7],
					},
					Tokens: tk[:7],
				},
				Tokens: tk[:7],
			}
		}},
		{`class a:b`, func(t *test, tk Tokens) { // 10
			t.Output = CompoundStatement{
				Class: &ClassDefinition{
					ClassName: &tk[2],
					Suite: Suite{
						StatementList: &StatementList{
							Statements: []SimpleStatement{
								{
									Type: StatementAssignment,
									AssignmentStatement: &AssignmentStatement{
										StarredExpression: &StarredExpression{
											OrExpr: WrapConditional(&Atom{
												Identifier: &tk[4],
												Tokens:     tk[4:5],
											}).OrTest.AndTest.NotTest.Comparison.OrExpression,
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
					Tokens: tk[:5],
				},
				Tokens: tk[:5],
			}
		}},
		{`async with a:b`, func(t *test, tk Tokens) { // 11
			t.Output = CompoundStatement{
				With: &WithStatement{
					Async: true,
					Contents: WithStatementContents{
						Items: []WithItem{
							{
								Expression: Expression{
									ConditionalExpression: WrapConditional(&Atom{
										Identifier: &tk[4],
										Tokens:     tk[4:5],
									}),
									Tokens: tk[4:5],
								},
								Tokens: tk[4:5],
							},
						},
						Tokens: tk[4:5],
					},
					Suite: Suite{
						StatementList: &StatementList{
							Statements: []SimpleStatement{
								{
									Type: StatementAssignment,
									AssignmentStatement: &AssignmentStatement{
										StarredExpression: &StarredExpression{
											OrExpr: WrapConditional(&Atom{
												Identifier: &tk[6],
												Tokens:     tk[6:7],
											}).OrTest.AndTest.NotTest.Comparison.OrExpression,
											Tokens: tk[6:7],
										},
										Tokens: tk[6:7],
									},
									Tokens: tk[6:7],
								},
							},
							Tokens: tk[6:7],
						},
						Tokens: tk[6:7],
					},
					Tokens: tk[:7],
				},
				Tokens: tk[:7],
			}
		}},
		{`async def a():b`, func(t *test, tk Tokens) { // 12
			t.Output = CompoundStatement{
				Func: &FuncDefinition{
					Async:    true,
					FuncName: &tk[4],
					ParameterList: ParameterList{
						Tokens: tk[6:6],
					},
					Suite: Suite{
						StatementList: &StatementList{
							Statements: []SimpleStatement{
								{
									Type: StatementAssignment,
									AssignmentStatement: &AssignmentStatement{
										StarredExpression: &StarredExpression{
											OrExpr: WrapConditional(&Atom{
												Identifier: &tk[8],
												Tokens:     tk[8:9],
											}).OrTest.AndTest.NotTest.Comparison.OrExpression,
											Tokens: tk[8:9],
										},
										Tokens: tk[8:9],
									},
									Tokens: tk[8:9],
								},
							},
							Tokens: tk[8:9],
						},
						Tokens: tk[8:9],
					},
					Tokens: tk[:9],
				},
				Tokens: tk[:9],
			}
		}},
		{`async with a:b`, func(t *test, tk Tokens) { // 13
			t.Output = CompoundStatement{
				With: &WithStatement{
					Async: true,
					Contents: WithStatementContents{
						Items: []WithItem{
							{
								Expression: Expression{
									ConditionalExpression: WrapConditional(&Atom{
										Identifier: &tk[4],
										Tokens:     tk[4:5],
									}),
									Tokens: tk[4:5],
								},
								Tokens: tk[4:5],
							},
						},
						Tokens: tk[4:5],
					},
					Suite: Suite{
						StatementList: &StatementList{
							Statements: []SimpleStatement{
								{
									Type: StatementAssignment,
									AssignmentStatement: &AssignmentStatement{
										StarredExpression: &StarredExpression{
											OrExpr: WrapConditional(&Atom{
												Identifier: &tk[6],
												Tokens:     tk[6:7],
											}).OrTest.AndTest.NotTest.Comparison.OrExpression,
											Tokens: tk[6:7],
										},
										Tokens: tk[6:7],
									},
									Tokens: tk[6:7],
								},
							},
							Tokens: tk[6:7],
						},
						Tokens: tk[6:7],
					},
					Tokens: tk[:7],
				},
				Tokens: tk[:7],
			}
		}},
		{"@nonlocal\ndef a():b", func(t *test, tk Tokens) { // 14
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
					Parsing: "Decorators",
					Token:   tk[1],
				},
				Parsing: "CompoundStatement",
				Token:   tk[0],
			}
		}},
		{"@a\ndef nonlocal():b", func(t *test, tk Tokens) { // 15
			t.Err = Error{
				Err: Error{
					Err:     ErrMissingIdentifier,
					Parsing: "FuncDefinition",
					Token:   tk[5],
				},
				Parsing: "CompoundStatement",
				Token:   tk[0],
			}
		}},
		{"@a\nclass nonlocal():b", func(t *test, tk Tokens) { // 16
			t.Err = Error{
				Err: Error{
					Err:     ErrMissingIdentifier,
					Parsing: "ClassDefinition",
					Token:   tk[5],
				},
				Parsing: "CompoundStatement",
				Token:   tk[0],
			}
		}},
		{"@a\nasync def nonlocal():b", func(t *test, tk Tokens) { // 17
			t.Err = Error{
				Err: Error{
					Err:     ErrMissingIdentifier,
					Parsing: "FuncDefinition",
					Token:   tk[7],
				},
				Parsing: "CompoundStatement",
				Token:   tk[0],
			}
		}},
		{"@a\nasync with a:b", func(t *test, tk Tokens) { // 18
			t.Err = Error{
				Err:     ErrInvalidCompound,
				Parsing: "CompoundStatement",
				Token:   tk[0],
			}
		}},
		{"@a\nwith a:b", func(t *test, tk Tokens) { // 19
			t.Err = Error{
				Err:     ErrInvalidCompound,
				Parsing: "CompoundStatement",
				Token:   tk[0],
			}
		}},
		{"if nonlocal:a", func(t *test, tk Tokens) { // 20
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
					Parsing: "IfStatement",
					Token:   tk[2],
				},
				Parsing: "CompoundStatement",
				Token:   tk[0],
			}
		}},
		{"while nonlocal:a", func(t *test, tk Tokens) { // 21
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
					Parsing: "WhileStatement",
					Token:   tk[2],
				},
				Parsing: "CompoundStatement",
				Token:   tk[0],
			}
		}},
		{"for nonlocal in a:b", func(t *test, tk Tokens) { // 22
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
					Parsing: "ForStatement",
					Token:   tk[2],
				},
				Parsing: "CompoundStatement",
				Token:   tk[0],
			}
		}},
		{"try:nonlocal\nfinally:b", func(t *test, tk Tokens) { // 23
			t.Err = Error{
				Err: Error{
					Err: Error{
						Err: Error{
							Err: Error{
								Err: Error{
									Err:     ErrMissingIdentifier,
									Parsing: "NonLocalStatement",
									Token:   tk[3],
								},
								Parsing: "SimpleStatement",
								Token:   tk[2],
							},
							Parsing: "StatementList",
							Token:   tk[2],
						},
						Parsing: "Suite",
						Token:   tk[2],
					},
					Parsing: "TryStatement",
					Token:   tk[2],
				},
				Parsing: "CompoundStatement",
				Token:   tk[0],
			}
		}},
		{`with nonlocal:a`, func(t *test, tk Tokens) { // 24
			t.Err = Error{
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
							Parsing: "WithItem",
							Token:   tk[2],
						},
						Parsing: "WithStatementContents",
						Token:   tk[2],
					},
					Parsing: "WithStatement",
					Token:   tk[2],
				},
				Parsing: "CompoundStatement",
				Token:   tk[0],
			}
		}},
		{"def nonlocal():a", func(t *test, tk Tokens) { // 25
			t.Err = Error{
				Err: Error{
					Err:     ErrMissingIdentifier,
					Parsing: "FuncDefinition",
					Token:   tk[2],
				},
				Parsing: "CompoundStatement",
				Token:   tk[0],
			}
		}},
		{"class nonlocal():a", func(t *test, tk Tokens) { // 26
			t.Err = Error{
				Err: Error{
					Err:     ErrMissingIdentifier,
					Parsing: "ClassDefinition",
					Token:   tk[2],
				},
				Parsing: "CompoundStatement",
				Token:   tk[0],
			}
		}},
		{"async for nonlocal in a:b", func(t *test, tk Tokens) { // 27
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
					Parsing: "ForStatement",
					Token:   tk[4],
				},
				Parsing: "CompoundStatement",
				Token:   tk[0],
			}
		}},
		{`async with nonlocal:a`, func(t *test, tk Tokens) { // 28
			t.Err = Error{
				Err: Error{
					Err: Error{
						Err: Error{
							Err: Error{
								Err: wrapConditionalExpressionError(Error{
									Err:     ErrInvalidEnclosure,
									Parsing: "Enclosure",
									Token:   tk[4],
								}),
								Parsing: "Expression",
								Token:   tk[4],
							},
							Parsing: "WithItem",
							Token:   tk[4],
						},
						Parsing: "WithStatementContents",
						Token:   tk[4],
					},
					Parsing: "WithStatement",
					Token:   tk[4],
				},
				Parsing: "CompoundStatement",
				Token:   tk[0],
			}
		}},
		{"async def nonlocal():a", func(t *test, tk Tokens) { // 29
			t.Err = Error{
				Err: Error{
					Err:     ErrMissingIdentifier,
					Parsing: "FuncDefinition",
					Token:   tk[4],
				},
				Parsing: "CompoundStatement",
				Token:   tk[0],
			}
		}},
		{"async class a():b", func(t *test, tk Tokens) { // 30
			t.Err = Error{
				Err:     ErrInvalidCompound,
				Parsing: "CompoundStatement",
				Token:   tk[0],
			}
		}},
		{"a", func(t *test, tk Tokens) { // 31
			t.Err = Error{
				Err:     ErrInvalidCompound,
				Parsing: "CompoundStatement",
				Token:   tk[0],
			}
		}},
	}, func(t *test) (Type, error) {
		var c CompoundStatement

		err := c.parser(t.Tokens)

		return c, err
	})
}

func TestDecorators(t *testing.T) {
	doTests(t, []sourceFn{
		{"@a\n", func(t *test, tk Tokens) { // 1
			t.Output = Decorators{
				Decorators: []AssignmentExpression{
					{
						Expression: Expression{
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
		{"@a\n@b\n", func(t *test, tk Tokens) { // 2
			t.Output = Decorators{
				Decorators: []AssignmentExpression{
					{
						Expression: Expression{
							ConditionalExpression: WrapConditional(&Atom{
								Identifier: &tk[1],
								Tokens:     tk[1:2],
							}),
							Tokens: tk[1:2],
						},
						Tokens: tk[1:2],
					},
					{
						Expression: Expression{
							ConditionalExpression: WrapConditional(&Atom{
								Identifier: &tk[4],
								Tokens:     tk[4:5],
							}),
							Tokens: tk[4:5],
						},
						Tokens: tk[4:5],
					},
				},
				Tokens: tk[:6],
			}
		}},
		{"@a\n#test\n@b\n", func(t *test, tk Tokens) { // 3
			t.Output = Decorators{
				Decorators: []AssignmentExpression{
					{
						Expression: Expression{
							ConditionalExpression: WrapConditional(&Atom{
								Identifier: &tk[1],
								Tokens:     tk[1:2],
							}),
							Tokens: tk[1:2],
						},
						Tokens: tk[1:2],
					},
					{
						Expression: Expression{
							ConditionalExpression: WrapConditional(&Atom{
								Identifier: &tk[6],
								Tokens:     tk[6:7],
							}),
							Tokens: tk[6:7],
						},
						Tokens: tk[6:7],
					},
				},
				Tokens: tk[:8],
			}
		}},
		{"@nonlocal\n", func(t *test, tk Tokens) { // 4
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
				Parsing: "Decorators",
				Token:   tk[1],
			}
		}},
		{"@a", func(t *test, tk Tokens) { // 5
			t.Err = Error{
				Err:     ErrMissingNewline,
				Parsing: "Decorators",
				Token:   tk[2],
			}
		}},
	}, func(t *test) (Type, error) {
		var d Decorators

		err := d.parse(t.Tokens)

		return d, err
	})
}

func TestIfStatement(t *testing.T) {
	doTests(t, []sourceFn{
		{"if a:b", func(t *test, tk Tokens) { // 1
			t.Output = IfStatement{
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
				Suite: Suite{
					StatementList: &StatementList{
						Statements: []SimpleStatement{
							{
								Type: StatementAssignment,
								AssignmentStatement: &AssignmentStatement{
									StarredExpression: &StarredExpression{
										OrExpr: WrapConditional(&Atom{
											Identifier: &tk[4],
											Tokens:     tk[4:5],
										}).OrTest.AndTest.NotTest.Comparison.OrExpression,
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
				Tokens: tk[:5],
			}
		}},
		{"if a : b", func(t *test, tk Tokens) { // 2
			t.Output = IfStatement{
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
				Suite: Suite{
					StatementList: &StatementList{
						Statements: []SimpleStatement{
							{
								Type: StatementAssignment,
								AssignmentStatement: &AssignmentStatement{
									StarredExpression: &StarredExpression{
										OrExpr: WrapConditional(&Atom{
											Identifier: &tk[6],
											Tokens:     tk[6:7],
										}).OrTest.AndTest.NotTest.Comparison.OrExpression,
										Tokens: tk[6:7],
									},
									Tokens: tk[6:7],
								},
								Tokens: tk[6:7],
							},
						},
						Tokens: tk[6:7],
					},
					Tokens: tk[6:7],
				},
				Tokens: tk[:7],
			}
		}},
		{"if a:b\nelif c:d", func(t *test, tk Tokens) { // 3
			t.Output = IfStatement{
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
				Suite: Suite{
					StatementList: &StatementList{
						Statements: []SimpleStatement{
							{
								Type: StatementAssignment,
								AssignmentStatement: &AssignmentStatement{
									StarredExpression: &StarredExpression{
										OrExpr: WrapConditional(&Atom{
											Identifier: &tk[4],
											Tokens:     tk[4:5],
										}).OrTest.AndTest.NotTest.Comparison.OrExpression,
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
				Elif: []AssignmentExpressionAndSuite{
					{
						AssignmentExpression: AssignmentExpression{
							Expression: Expression{
								ConditionalExpression: WrapConditional(&Atom{
									Identifier: &tk[8],
									Tokens:     tk[8:9],
								}),
								Tokens: tk[8:9],
							},
							Tokens: tk[8:9],
						},
						Suite: Suite{
							StatementList: &StatementList{
								Statements: []SimpleStatement{
									{
										Type: StatementAssignment,
										AssignmentStatement: &AssignmentStatement{
											StarredExpression: &StarredExpression{
												OrExpr: WrapConditional(&Atom{
													Identifier: &tk[10],
													Tokens:     tk[10:11],
												}).OrTest.AndTest.NotTest.Comparison.OrExpression,
												Tokens: tk[10:11],
											},
											Tokens: tk[10:11],
										},
										Tokens: tk[10:11],
									},
								},
								Tokens: tk[10:11],
							},
							Tokens: tk[10:11],
						},
					},
				},
				Tokens: tk[:11],
			}
		}},
		{"if a:b\nelif c : d", func(t *test, tk Tokens) { // 4
			t.Output = IfStatement{
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
				Suite: Suite{
					StatementList: &StatementList{
						Statements: []SimpleStatement{
							{
								Type: StatementAssignment,
								AssignmentStatement: &AssignmentStatement{
									StarredExpression: &StarredExpression{
										OrExpr: WrapConditional(&Atom{
											Identifier: &tk[4],
											Tokens:     tk[4:5],
										}).OrTest.AndTest.NotTest.Comparison.OrExpression,
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
				Elif: []AssignmentExpressionAndSuite{
					{
						AssignmentExpression: AssignmentExpression{
							Expression: Expression{
								ConditionalExpression: WrapConditional(&Atom{
									Identifier: &tk[8],
									Tokens:     tk[8:9],
								}),
								Tokens: tk[8:9],
							},
							Tokens: tk[8:9],
						},
						Suite: Suite{
							StatementList: &StatementList{
								Statements: []SimpleStatement{
									{
										Type: StatementAssignment,
										AssignmentStatement: &AssignmentStatement{
											StarredExpression: &StarredExpression{
												OrExpr: WrapConditional(&Atom{
													Identifier: &tk[12],
													Tokens:     tk[12:13],
												}).OrTest.AndTest.NotTest.Comparison.OrExpression,
												Tokens: tk[12:13],
											},
											Tokens: tk[12:13],
										},
										Tokens: tk[12:13],
									},
								},
								Tokens: tk[12:13],
							},
							Tokens: tk[12:13],
						},
					},
				},
				Tokens: tk[:13],
			}
		}},
		{"if a:b\nelif c:d\nelif e:f", func(t *test, tk Tokens) { // 5
			t.Output = IfStatement{
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
				Suite: Suite{
					StatementList: &StatementList{
						Statements: []SimpleStatement{
							{
								Type: StatementAssignment,
								AssignmentStatement: &AssignmentStatement{
									StarredExpression: &StarredExpression{
										OrExpr: WrapConditional(&Atom{
											Identifier: &tk[4],
											Tokens:     tk[4:5],
										}).OrTest.AndTest.NotTest.Comparison.OrExpression,
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
				Elif: []AssignmentExpressionAndSuite{
					{
						AssignmentExpression: AssignmentExpression{
							Expression: Expression{
								ConditionalExpression: WrapConditional(&Atom{
									Identifier: &tk[8],
									Tokens:     tk[8:9],
								}),
								Tokens: tk[8:9],
							},
							Tokens: tk[8:9],
						},
						Suite: Suite{
							StatementList: &StatementList{
								Statements: []SimpleStatement{
									{
										Type: StatementAssignment,
										AssignmentStatement: &AssignmentStatement{
											StarredExpression: &StarredExpression{
												OrExpr: WrapConditional(&Atom{
													Identifier: &tk[10],
													Tokens:     tk[10:11],
												}).OrTest.AndTest.NotTest.Comparison.OrExpression,
												Tokens: tk[10:11],
											},
											Tokens: tk[10:11],
										},
										Tokens: tk[10:11],
									},
								},
								Tokens: tk[10:11],
							},
							Tokens: tk[10:11],
						},
					},
					{
						AssignmentExpression: AssignmentExpression{
							Expression: Expression{
								ConditionalExpression: WrapConditional(&Atom{
									Identifier: &tk[14],
									Tokens:     tk[14:15],
								}),
								Tokens: tk[14:15],
							},
							Tokens: tk[14:15],
						},
						Suite: Suite{
							StatementList: &StatementList{
								Statements: []SimpleStatement{
									{
										Type: StatementAssignment,
										AssignmentStatement: &AssignmentStatement{
											StarredExpression: &StarredExpression{
												OrExpr: WrapConditional(&Atom{
													Identifier: &tk[16],
													Tokens:     tk[16:17],
												}).OrTest.AndTest.NotTest.Comparison.OrExpression,
												Tokens: tk[16:17],
											},
											Tokens: tk[16:17],
										},
										Tokens: tk[16:17],
									},
								},
								Tokens: tk[16:17],
							},
							Tokens: tk[16:17],
						},
					},
				},
				Tokens: tk[:17],
			}
		}},
		{"if a:b\nelse:c", func(t *test, tk Tokens) { // 6
			t.Output = IfStatement{
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
				Suite: Suite{
					StatementList: &StatementList{
						Statements: []SimpleStatement{
							{
								Type: StatementAssignment,
								AssignmentStatement: &AssignmentStatement{
									StarredExpression: &StarredExpression{
										OrExpr: WrapConditional(&Atom{
											Identifier: &tk[4],
											Tokens:     tk[4:5],
										}).OrTest.AndTest.NotTest.Comparison.OrExpression,
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
				Else: &Suite{
					StatementList: &StatementList{
						Statements: []SimpleStatement{
							{
								Type: StatementAssignment,
								AssignmentStatement: &AssignmentStatement{
									StarredExpression: &StarredExpression{
										OrExpr: WrapConditional(&Atom{
											Identifier: &tk[8],
											Tokens:     tk[8:9],
										}).OrTest.AndTest.NotTest.Comparison.OrExpression,
										Tokens: tk[8:9],
									},
									Tokens: tk[8:9],
								},
								Tokens: tk[8:9],
							},
						},
						Tokens: tk[8:9],
					},
					Tokens: tk[8:9],
				},
				Tokens: tk[:9],
			}
		}},
		{"if a:b\nelse : c", func(t *test, tk Tokens) { // 7
			t.Output = IfStatement{
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
				Suite: Suite{
					StatementList: &StatementList{
						Statements: []SimpleStatement{
							{
								Type: StatementAssignment,
								AssignmentStatement: &AssignmentStatement{
									StarredExpression: &StarredExpression{
										OrExpr: WrapConditional(&Atom{
											Identifier: &tk[4],
											Tokens:     tk[4:5],
										}).OrTest.AndTest.NotTest.Comparison.OrExpression,
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
				Else: &Suite{
					StatementList: &StatementList{
						Statements: []SimpleStatement{
							{
								Type: StatementAssignment,
								AssignmentStatement: &AssignmentStatement{
									StarredExpression: &StarredExpression{
										OrExpr: WrapConditional(&Atom{
											Identifier: &tk[10],
											Tokens:     tk[10:11],
										}).OrTest.AndTest.NotTest.Comparison.OrExpression,
										Tokens: tk[10:11],
									},
									Tokens: tk[10:11],
								},
								Tokens: tk[10:11],
							},
						},
						Tokens: tk[10:11],
					},
					Tokens: tk[10:11],
				},
				Tokens: tk[:11],
			}
		}},
		{"if a:b\nelif c:d\nelse:e", func(t *test, tk Tokens) { // 8
			t.Output = IfStatement{
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
				Suite: Suite{
					StatementList: &StatementList{
						Statements: []SimpleStatement{
							{
								Type: StatementAssignment,
								AssignmentStatement: &AssignmentStatement{
									StarredExpression: &StarredExpression{
										OrExpr: WrapConditional(&Atom{
											Identifier: &tk[4],
											Tokens:     tk[4:5],
										}).OrTest.AndTest.NotTest.Comparison.OrExpression,
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
				Elif: []AssignmentExpressionAndSuite{
					{
						AssignmentExpression: AssignmentExpression{
							Expression: Expression{
								ConditionalExpression: WrapConditional(&Atom{
									Identifier: &tk[8],
									Tokens:     tk[8:9],
								}),
								Tokens: tk[8:9],
							},
							Tokens: tk[8:9],
						},
						Suite: Suite{
							StatementList: &StatementList{
								Statements: []SimpleStatement{
									{
										Type: StatementAssignment,
										AssignmentStatement: &AssignmentStatement{
											StarredExpression: &StarredExpression{
												OrExpr: WrapConditional(&Atom{
													Identifier: &tk[10],
													Tokens:     tk[10:11],
												}).OrTest.AndTest.NotTest.Comparison.OrExpression,
												Tokens: tk[10:11],
											},
											Tokens: tk[10:11],
										},
										Tokens: tk[10:11],
									},
								},
								Tokens: tk[10:11],
							},
							Tokens: tk[10:11],
						},
					},
				},
				Else: &Suite{
					StatementList: &StatementList{
						Statements: []SimpleStatement{
							{
								Type: StatementAssignment,
								AssignmentStatement: &AssignmentStatement{
									StarredExpression: &StarredExpression{
										OrExpr: WrapConditional(&Atom{
											Identifier: &tk[14],
											Tokens:     tk[14:15],
										}).OrTest.AndTest.NotTest.Comparison.OrExpression,
										Tokens: tk[14:15],
									},
									Tokens: tk[14:15],
								},
								Tokens: tk[14:15],
							},
						},
						Tokens: tk[14:15],
					},
					Tokens: tk[14:15],
				},
				Tokens: tk[:15],
			}
		}},
		{"if nonlocal:a", func(t *test, tk Tokens) { // 9
			t.Err = Error{
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
				Parsing: "IfStatement",
				Token:   tk[2],
			}
		}},
		{"if a b", func(t *test, tk Tokens) { // 10
			t.Err = Error{
				Err:     ErrMissingColon,
				Parsing: "IfStatement",
				Token:   tk[4],
			}
		}},
		{"if a:nonlocal", func(t *test, tk Tokens) { // 11
			t.Err = Error{
				Err: Error{
					Err: Error{
						Err: Error{
							Err: Error{
								Err:     ErrMissingIdentifier,
								Parsing: "NonLocalStatement",
								Token:   tk[5],
							},
							Parsing: "SimpleStatement",
							Token:   tk[4],
						},
						Parsing: "StatementList",
						Token:   tk[4],
					},
					Parsing: "Suite",
					Token:   tk[4],
				},
				Parsing: "IfStatement",
				Token:   tk[4],
			}
		}},
		{"if a:b\nelif nonlocal:c", func(t *test, tk Tokens) { // 12
			t.Err = Error{
				Err: Error{
					Err: Error{
						Err: wrapConditionalExpressionError(Error{
							Err:     ErrInvalidEnclosure,
							Parsing: "Enclosure",
							Token:   tk[8],
						}),
						Parsing: "Expression",
						Token:   tk[8],
					},
					Parsing: "AssignmentExpression",
					Token:   tk[8],
				},
				Parsing: "IfStatement",
				Token:   tk[8],
			}
		}},
		{"if a:b\nelif c d", func(t *test, tk Tokens) { // 13
			t.Err = Error{
				Err:     ErrMissingColon,
				Parsing: "IfStatement",
				Token:   tk[10],
			}
		}},
		{"if a:b\nelif c:nonlocal", func(t *test, tk Tokens) { // 14
			t.Err = Error{
				Err: Error{
					Err: Error{
						Err: Error{
							Err: Error{
								Err:     ErrMissingIdentifier,
								Parsing: "NonLocalStatement",
								Token:   tk[11],
							},
							Parsing: "SimpleStatement",
							Token:   tk[10],
						},
						Parsing: "StatementList",
						Token:   tk[10],
					},
					Parsing: "Suite",
					Token:   tk[10],
				},
				Parsing: "IfStatement",
				Token:   tk[10],
			}
		}},
		{"if a:b\nelse c", func(t *test, tk Tokens) { // 15
			t.Err = Error{
				Err:     ErrMissingColon,
				Parsing: "IfStatement",
				Token:   tk[8],
			}
		}},
		{"if a:b\nelse:nonlocal", func(t *test, tk Tokens) { // 16
			t.Err = Error{
				Err: Error{
					Err: Error{
						Err: Error{
							Err: Error{
								Err:     ErrMissingIdentifier,
								Parsing: "NonLocalStatement",
								Token:   tk[9],
							},
							Parsing: "SimpleStatement",
							Token:   tk[8],
						},
						Parsing: "StatementList",
						Token:   tk[8],
					},
					Parsing: "Suite",
					Token:   tk[8],
				},
				Parsing: "IfStatement",
				Token:   tk[8],
			}
		}},
	}, func(t *test) (Type, error) {
		var i IfStatement

		err := i.parse(t.Tokens)

		return i, err
	})
}

func TestWhileStatement(t *testing.T) {
	doTests(t, []sourceFn{
		{"while a:b", func(t *test, tk Tokens) { // 1
			t.Output = WhileStatement{
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
				Suite: Suite{
					StatementList: &StatementList{
						Statements: []SimpleStatement{
							{
								Type: StatementAssignment,
								AssignmentStatement: &AssignmentStatement{
									StarredExpression: &StarredExpression{
										OrExpr: WrapConditional(&Atom{
											Identifier: &tk[4],
											Tokens:     tk[4:5],
										}).OrTest.AndTest.NotTest.Comparison.OrExpression,
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
				Tokens: tk[:5],
			}
		}},
		{"while a : b", func(t *test, tk Tokens) { // 2
			t.Output = WhileStatement{
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
				Suite: Suite{
					StatementList: &StatementList{
						Statements: []SimpleStatement{
							{
								Type: StatementAssignment,
								AssignmentStatement: &AssignmentStatement{
									StarredExpression: &StarredExpression{
										OrExpr: WrapConditional(&Atom{
											Identifier: &tk[6],
											Tokens:     tk[6:7],
										}).OrTest.AndTest.NotTest.Comparison.OrExpression,
										Tokens: tk[6:7],
									},
									Tokens: tk[6:7],
								},
								Tokens: tk[6:7],
							},
						},
						Tokens: tk[6:7],
					},
					Tokens: tk[6:7],
				},
				Tokens: tk[:7],
			}
		}},
		{"while a:b\nelse:c", func(t *test, tk Tokens) { // 3
			t.Output = WhileStatement{
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
				Suite: Suite{
					StatementList: &StatementList{
						Statements: []SimpleStatement{
							{
								Type: StatementAssignment,
								AssignmentStatement: &AssignmentStatement{
									StarredExpression: &StarredExpression{
										OrExpr: WrapConditional(&Atom{
											Identifier: &tk[4],
											Tokens:     tk[4:5],
										}).OrTest.AndTest.NotTest.Comparison.OrExpression,
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
				Else: &Suite{
					StatementList: &StatementList{
						Statements: []SimpleStatement{
							{
								Type: StatementAssignment,
								AssignmentStatement: &AssignmentStatement{
									StarredExpression: &StarredExpression{
										OrExpr: WrapConditional(&Atom{
											Identifier: &tk[8],
											Tokens:     tk[8:9],
										}).OrTest.AndTest.NotTest.Comparison.OrExpression,
										Tokens: tk[8:9],
									},
									Tokens: tk[8:9],
								},
								Tokens: tk[8:9],
							},
						},
						Tokens: tk[8:9],
					},
					Tokens: tk[8:9],
				},
				Tokens: tk[:9],
			}
		}},
		{"while a:b\nelse : c", func(t *test, tk Tokens) { // 4
			t.Output = WhileStatement{
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
				Suite: Suite{
					StatementList: &StatementList{
						Statements: []SimpleStatement{
							{
								Type: StatementAssignment,
								AssignmentStatement: &AssignmentStatement{
									StarredExpression: &StarredExpression{
										OrExpr: WrapConditional(&Atom{
											Identifier: &tk[4],
											Tokens:     tk[4:5],
										}).OrTest.AndTest.NotTest.Comparison.OrExpression,
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
				Else: &Suite{
					StatementList: &StatementList{
						Statements: []SimpleStatement{
							{
								Type: StatementAssignment,
								AssignmentStatement: &AssignmentStatement{
									StarredExpression: &StarredExpression{
										OrExpr: WrapConditional(&Atom{
											Identifier: &tk[10],
											Tokens:     tk[10:11],
										}).OrTest.AndTest.NotTest.Comparison.OrExpression,
										Tokens: tk[10:11],
									},
									Tokens: tk[10:11],
								},
								Tokens: tk[10:11],
							},
						},
						Tokens: tk[10:11],
					},
					Tokens: tk[10:11],
				},
				Tokens: tk[:11],
			}
		}},
		{"while nonlocal:a", func(t *test, tk Tokens) { // 5
			t.Err = Error{
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
				Parsing: "WhileStatement",
				Token:   tk[2],
			}
		}},
		{"while a b", func(t *test, tk Tokens) { // 6
			t.Err = Error{
				Err:     ErrMissingColon,
				Parsing: "WhileStatement",
				Token:   tk[4],
			}
		}},
		{"while a:nonlocal", func(t *test, tk Tokens) { // 7
			t.Err = Error{
				Err: Error{
					Err: Error{
						Err: Error{
							Err: Error{
								Err:     ErrMissingIdentifier,
								Parsing: "NonLocalStatement",
								Token:   tk[5],
							},
							Parsing: "SimpleStatement",
							Token:   tk[4],
						},
						Parsing: "StatementList",
						Token:   tk[4],
					},
					Parsing: "Suite",
					Token:   tk[4],
				},
				Parsing: "WhileStatement",
				Token:   tk[4],
			}
		}},
		{"while a:b\nelse c", func(t *test, tk Tokens) { // 8
			t.Err = Error{
				Err:     ErrMissingColon,
				Parsing: "WhileStatement",
				Token:   tk[8],
			}
		}},
		{"while a:b\nelse:nonlocal", func(t *test, tk Tokens) { // 9
			t.Err = Error{
				Err: Error{
					Err: Error{
						Err: Error{
							Err: Error{
								Err:     ErrMissingIdentifier,
								Parsing: "NonLocalStatement",
								Token:   tk[9],
							},
							Parsing: "SimpleStatement",
							Token:   tk[8],
						},
						Parsing: "StatementList",
						Token:   tk[8],
					},
					Parsing: "Suite",
					Token:   tk[8],
				},
				Parsing: "WhileStatement",
				Token:   tk[8],
			}
		}},
	}, func(t *test) (Type, error) {
		var w WhileStatement

		err := w.parse(t.Tokens)

		return w, err
	})
}

func TestForStatement(t *testing.T) {
	doTests(t, []sourceFn{
		{"for a in b:c", func(t *test, tk Tokens) { // 1
			t.Output = ForStatement{
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
				StarredList: StarredList{
					StarredItems: []StarredItem{
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
				Suite: Suite{
					StatementList: &StatementList{
						Statements: []SimpleStatement{
							{
								Type: StatementAssignment,
								AssignmentStatement: &AssignmentStatement{
									StarredExpression: &StarredExpression{
										OrExpr: WrapConditional(&Atom{
											Identifier: &tk[8],
											Tokens:     tk[8:9],
										}).OrTest.AndTest.NotTest.Comparison.OrExpression,
										Tokens: tk[8:9],
									},
									Tokens: tk[8:9],
								},
								Tokens: tk[8:9],
							},
						},
						Tokens: tk[8:9],
					},
					Tokens: tk[8:9],
				},
				Tokens: tk[:9],
			}
		}},
		{"for a in b : c", func(t *test, tk Tokens) { // 2
			t.Output = ForStatement{
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
				StarredList: StarredList{
					StarredItems: []StarredItem{
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
				Suite: Suite{
					StatementList: &StatementList{
						Statements: []SimpleStatement{
							{
								Type: StatementAssignment,
								AssignmentStatement: &AssignmentStatement{
									StarredExpression: &StarredExpression{
										OrExpr: WrapConditional(&Atom{
											Identifier: &tk[10],
											Tokens:     tk[10:11],
										}).OrTest.AndTest.NotTest.Comparison.OrExpression,
										Tokens: tk[10:11],
									},
									Tokens: tk[10:11],
								},
								Tokens: tk[10:11],
							},
						},
						Tokens: tk[10:11],
					},
					Tokens: tk[10:11],
				},
				Tokens: tk[:11],
			}
		}},
		{"for a in b:c", func(t *test, tk Tokens) { // 3
			t.Async = true
			t.Output = ForStatement{
				Async: true,
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
				StarredList: StarredList{
					StarredItems: []StarredItem{
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
				Suite: Suite{
					StatementList: &StatementList{
						Statements: []SimpleStatement{
							{
								Type: StatementAssignment,
								AssignmentStatement: &AssignmentStatement{
									StarredExpression: &StarredExpression{
										OrExpr: WrapConditional(&Atom{
											Identifier: &tk[8],
											Tokens:     tk[8:9],
										}).OrTest.AndTest.NotTest.Comparison.OrExpression,
										Tokens: tk[8:9],
									},
									Tokens: tk[8:9],
								},
								Tokens: tk[8:9],
							},
						},
						Tokens: tk[8:9],
					},
					Tokens: tk[8:9],
				},
				Tokens: tk[:9],
			}
		}},
		{"for a in b:c\nelse:d", func(t *test, tk Tokens) { // 4
			t.Output = ForStatement{
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
				StarredList: StarredList{
					StarredItems: []StarredItem{
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
				Suite: Suite{
					StatementList: &StatementList{
						Statements: []SimpleStatement{
							{
								Type: StatementAssignment,
								AssignmentStatement: &AssignmentStatement{
									StarredExpression: &StarredExpression{
										OrExpr: WrapConditional(&Atom{
											Identifier: &tk[8],
											Tokens:     tk[8:9],
										}).OrTest.AndTest.NotTest.Comparison.OrExpression,
										Tokens: tk[8:9],
									},
									Tokens: tk[8:9],
								},
								Tokens: tk[8:9],
							},
						},
						Tokens: tk[8:9],
					},
					Tokens: tk[8:9],
				},
				Else: &Suite{
					StatementList: &StatementList{
						Statements: []SimpleStatement{
							{
								Type: StatementAssignment,
								AssignmentStatement: &AssignmentStatement{
									StarredExpression: &StarredExpression{
										OrExpr: WrapConditional(&Atom{
											Identifier: &tk[12],
											Tokens:     tk[12:13],
										}).OrTest.AndTest.NotTest.Comparison.OrExpression,
										Tokens: tk[12:13],
									},
									Tokens: tk[12:13],
								},
								Tokens: tk[12:13],
							},
						},
						Tokens: tk[12:13],
					},
					Tokens: tk[12:13],
				},
				Tokens: tk[:13],
			}
		}},
		{"for a in b:c\nelse : d", func(t *test, tk Tokens) { // 5
			t.Output = ForStatement{
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
				StarredList: StarredList{
					StarredItems: []StarredItem{
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
				Suite: Suite{
					StatementList: &StatementList{
						Statements: []SimpleStatement{
							{
								Type: StatementAssignment,
								AssignmentStatement: &AssignmentStatement{
									StarredExpression: &StarredExpression{
										OrExpr: WrapConditional(&Atom{
											Identifier: &tk[8],
											Tokens:     tk[8:9],
										}).OrTest.AndTest.NotTest.Comparison.OrExpression,
										Tokens: tk[8:9],
									},
									Tokens: tk[8:9],
								},
								Tokens: tk[8:9],
							},
						},
						Tokens: tk[8:9],
					},
					Tokens: tk[8:9],
				},
				Else: &Suite{
					StatementList: &StatementList{
						Statements: []SimpleStatement{
							{
								Type: StatementAssignment,
								AssignmentStatement: &AssignmentStatement{
									StarredExpression: &StarredExpression{
										OrExpr: WrapConditional(&Atom{
											Identifier: &tk[14],
											Tokens:     tk[14:15],
										}).OrTest.AndTest.NotTest.Comparison.OrExpression,
										Tokens: tk[14:15],
									},
									Tokens: tk[14:15],
								},
								Tokens: tk[14:15],
							},
						},
						Tokens: tk[14:15],
					},
					Tokens: tk[14:15],
				},
				Tokens: tk[:15],
			}
		}},
		{"for nonlocal in a:b\nelse:c", func(t *test, tk Tokens) { // 6
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
				Parsing: "ForStatement",
				Token:   tk[2],
			}
		}},
		{"for a in nonlocal:b\nelse:c", func(t *test, tk Tokens) { // 7
			t.Err = Error{
				Err: Error{
					Err: Error{
						Err: Error{
							Err: Error{
								Err: wrapConditionalExpressionError(Error{
									Err:     ErrInvalidEnclosure,
									Parsing: "Enclosure",
									Token:   tk[6],
								}),
								Parsing: "Expression",
								Token:   tk[6],
							},
							Parsing: "AssignmentExpression",
							Token:   tk[6],
						},
						Parsing: "StarredItem",
						Token:   tk[6],
					},
					Parsing: "StarredList",
					Token:   tk[6],
				},
				Parsing: "ForStatement",
				Token:   tk[6],
			}
		}},
		{"for a in b:nonlocal\nelse:c", func(t *test, tk Tokens) { // 7
			t.Err = Error{
				Err: Error{
					Err: Error{
						Err: Error{
							Err: Error{
								Err:     ErrMissingIdentifier,
								Parsing: "NonLocalStatement",
								Token:   tk[9],
							},
							Parsing: "SimpleStatement",
							Token:   tk[8],
						},
						Parsing: "StatementList",
						Token:   tk[8],
					},
					Parsing: "Suite",
					Token:   tk[8],
				},
				Parsing: "ForStatement",
				Token:   tk[8],
			}
		}},
		{"for a in b:c\nelse:nonlocal", func(t *test, tk Tokens) { // 8
			t.Err = Error{
				Err: Error{
					Err: Error{
						Err: Error{
							Err: Error{
								Err:     ErrMissingIdentifier,
								Parsing: "NonLocalStatement",
								Token:   tk[13],
							},
							Parsing: "SimpleStatement",
							Token:   tk[12],
						},
						Parsing: "StatementList",
						Token:   tk[12],
					},
					Parsing: "Suite",
					Token:   tk[12],
				},
				Parsing: "ForStatement",
				Token:   tk[12],
			}
		}},
		{"for a:b", func(t *test, tk Tokens) { // 9
			t.Err = Error{
				Err:     ErrMissingIn,
				Parsing: "ForStatement",
				Token:   tk[3],
			}
		}},
		{"for a in b for", func(t *test, tk Tokens) { // 10
			t.Err = Error{
				Err:     ErrMissingColon,
				Parsing: "ForStatement",
				Token:   tk[8],
			}
		}},
		{"for a in b:c\nelse d", func(t *test, tk Tokens) { // 10
			t.Err = Error{
				Err:     ErrMissingColon,
				Parsing: "ForStatement",
				Token:   tk[12],
			}
		}},
	}, func(t *test) (Type, error) {
		var f ForStatement

		err := f.parse(t.Tokens, t.Async)

		return f, err
	})
}

func TestTryStatement(t *testing.T) {
	doTests(t, []sourceFn{
		{"try:a\nexcept b:c", func(t *test, tk Tokens) { // 1
			t.Output = TryStatement{
				Try: Suite{
					StatementList: &StatementList{
						Statements: []SimpleStatement{
							{
								Type: StatementAssignment,
								AssignmentStatement: &AssignmentStatement{
									StarredExpression: &StarredExpression{
										OrExpr: WrapConditional(&Atom{
											Identifier: &tk[2],
											Tokens:     tk[2:3],
										}).OrTest.AndTest.NotTest.Comparison.OrExpression,
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
				Except: []Except{
					{
						Expression: Expression{
							ConditionalExpression: WrapConditional(&Atom{
								Identifier: &tk[6],
								Tokens:     tk[6:7],
							}),
							Tokens: tk[6:7],
						},
						Suite: Suite{
							StatementList: &StatementList{
								Statements: []SimpleStatement{
									{
										Type: StatementAssignment,
										AssignmentStatement: &AssignmentStatement{
											StarredExpression: &StarredExpression{
												OrExpr: WrapConditional(&Atom{
													Identifier: &tk[8],
													Tokens:     tk[8:9],
												}).OrTest.AndTest.NotTest.Comparison.OrExpression,
												Tokens: tk[8:9],
											},
											Tokens: tk[8:9],
										},
										Tokens: tk[8:9],
									},
								},
								Tokens: tk[8:9],
							},
							Tokens: tk[8:9],
						},
						Tokens: tk[6:9],
					},
				},
				Tokens: tk[:9],
			}
		}},
		{"try :a\nexcept b :c", func(t *test, tk Tokens) { // 2
			t.Output = TryStatement{
				Try: Suite{
					StatementList: &StatementList{
						Statements: []SimpleStatement{
							{
								Type: StatementAssignment,
								AssignmentStatement: &AssignmentStatement{
									StarredExpression: &StarredExpression{
										OrExpr: WrapConditional(&Atom{
											Identifier: &tk[3],
											Tokens:     tk[3:4],
										}).OrTest.AndTest.NotTest.Comparison.OrExpression,
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
				Except: []Except{
					{
						Expression: Expression{
							ConditionalExpression: WrapConditional(&Atom{
								Identifier: &tk[7],
								Tokens:     tk[7:8],
							}),
							Tokens: tk[7:8],
						},
						Suite: Suite{
							StatementList: &StatementList{
								Statements: []SimpleStatement{
									{
										Type: StatementAssignment,
										AssignmentStatement: &AssignmentStatement{
											StarredExpression: &StarredExpression{
												OrExpr: WrapConditional(&Atom{
													Identifier: &tk[10],
													Tokens:     tk[10:11],
												}).OrTest.AndTest.NotTest.Comparison.OrExpression,
												Tokens: tk[10:11],
											},
											Tokens: tk[10:11],
										},
										Tokens: tk[10:11],
									},
								},
								Tokens: tk[10:11],
							},
							Tokens: tk[10:11],
						},
						Tokens: tk[7:11],
					},
				},
				Tokens: tk[:11],
			}
		}},
		{"try:a\nexcept b:c\nexcept d:e", func(t *test, tk Tokens) { // 3
			t.Output = TryStatement{
				Try: Suite{
					StatementList: &StatementList{
						Statements: []SimpleStatement{
							{
								Type: StatementAssignment,
								AssignmentStatement: &AssignmentStatement{
									StarredExpression: &StarredExpression{
										OrExpr: WrapConditional(&Atom{
											Identifier: &tk[2],
											Tokens:     tk[2:3],
										}).OrTest.AndTest.NotTest.Comparison.OrExpression,
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
				Except: []Except{
					{
						Expression: Expression{
							ConditionalExpression: WrapConditional(&Atom{
								Identifier: &tk[6],
								Tokens:     tk[6:7],
							}),
							Tokens: tk[6:7],
						},
						Suite: Suite{
							StatementList: &StatementList{
								Statements: []SimpleStatement{
									{
										Type: StatementAssignment,
										AssignmentStatement: &AssignmentStatement{
											StarredExpression: &StarredExpression{
												OrExpr: WrapConditional(&Atom{
													Identifier: &tk[8],
													Tokens:     tk[8:9],
												}).OrTest.AndTest.NotTest.Comparison.OrExpression,
												Tokens: tk[8:9],
											},
											Tokens: tk[8:9],
										},
										Tokens: tk[8:9],
									},
								},
								Tokens: tk[8:9],
							},
							Tokens: tk[8:9],
						},
						Tokens: tk[6:9],
					},
					{
						Expression: Expression{
							ConditionalExpression: WrapConditional(&Atom{
								Identifier: &tk[12],
								Tokens:     tk[12:13],
							}),
							Tokens: tk[12:13],
						},
						Suite: Suite{
							StatementList: &StatementList{
								Statements: []SimpleStatement{
									{
										Type: StatementAssignment,
										AssignmentStatement: &AssignmentStatement{
											StarredExpression: &StarredExpression{
												OrExpr: WrapConditional(&Atom{
													Identifier: &tk[14],
													Tokens:     tk[14:15],
												}).OrTest.AndTest.NotTest.Comparison.OrExpression,
												Tokens: tk[14:15],
											},
											Tokens: tk[14:15],
										},
										Tokens: tk[14:15],
									},
								},
								Tokens: tk[14:15],
							},
							Tokens: tk[14:15],
						},
						Tokens: tk[12:15],
					},
				},
				Tokens: tk[:15],
			}
		}},
		{"try:a\nexcept *b:c", func(t *test, tk Tokens) { // 4
			t.Output = TryStatement{
				Try: Suite{
					StatementList: &StatementList{
						Statements: []SimpleStatement{
							{
								Type: StatementAssignment,
								AssignmentStatement: &AssignmentStatement{
									StarredExpression: &StarredExpression{
										OrExpr: WrapConditional(&Atom{
											Identifier: &tk[2],
											Tokens:     tk[2:3],
										}).OrTest.AndTest.NotTest.Comparison.OrExpression,
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
				Groups: true,
				Except: []Except{
					{
						Expression: Expression{
							ConditionalExpression: WrapConditional(&Atom{
								Identifier: &tk[7],
								Tokens:     tk[7:8],
							}),
							Tokens: tk[7:8],
						},
						Suite: Suite{
							StatementList: &StatementList{
								Statements: []SimpleStatement{
									{
										Type: StatementAssignment,
										AssignmentStatement: &AssignmentStatement{
											StarredExpression: &StarredExpression{
												OrExpr: WrapConditional(&Atom{
													Identifier: &tk[9],
													Tokens:     tk[9:10],
												}).OrTest.AndTest.NotTest.Comparison.OrExpression,
												Tokens: tk[9:10],
											},
											Tokens: tk[9:10],
										},
										Tokens: tk[9:10],
									},
								},
								Tokens: tk[9:10],
							},
							Tokens: tk[9:10],
						},
						Tokens: tk[7:10],
					},
				},
				Tokens: tk[:10],
			}
		}},
		{"try:a\nexcept * b:c", func(t *test, tk Tokens) { // 5
			t.Output = TryStatement{
				Try: Suite{
					StatementList: &StatementList{
						Statements: []SimpleStatement{
							{
								Type: StatementAssignment,
								AssignmentStatement: &AssignmentStatement{
									StarredExpression: &StarredExpression{
										OrExpr: WrapConditional(&Atom{
											Identifier: &tk[2],
											Tokens:     tk[2:3],
										}).OrTest.AndTest.NotTest.Comparison.OrExpression,
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
				Groups: true,
				Except: []Except{
					{
						Expression: Expression{
							ConditionalExpression: WrapConditional(&Atom{
								Identifier: &tk[8],
								Tokens:     tk[8:9],
							}),
							Tokens: tk[8:9],
						},
						Suite: Suite{
							StatementList: &StatementList{
								Statements: []SimpleStatement{
									{
										Type: StatementAssignment,
										AssignmentStatement: &AssignmentStatement{
											StarredExpression: &StarredExpression{
												OrExpr: WrapConditional(&Atom{
													Identifier: &tk[10],
													Tokens:     tk[10:11],
												}).OrTest.AndTest.NotTest.Comparison.OrExpression,
												Tokens: tk[10:11],
											},
											Tokens: tk[10:11],
										},
										Tokens: tk[10:11],
									},
								},
								Tokens: tk[10:11],
							},
							Tokens: tk[10:11],
						},
						Tokens: tk[8:11],
					},
				},
				Tokens: tk[:11],
			}
		}},
		{"try:a\nexcept *b:c\nexcept *d:e", func(t *test, tk Tokens) { // 6
			t.Output = TryStatement{
				Try: Suite{
					StatementList: &StatementList{
						Statements: []SimpleStatement{
							{
								Type: StatementAssignment,
								AssignmentStatement: &AssignmentStatement{
									StarredExpression: &StarredExpression{
										OrExpr: WrapConditional(&Atom{
											Identifier: &tk[2],
											Tokens:     tk[2:3],
										}).OrTest.AndTest.NotTest.Comparison.OrExpression,
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
				Groups: true,
				Except: []Except{
					{
						Expression: Expression{
							ConditionalExpression: WrapConditional(&Atom{
								Identifier: &tk[7],
								Tokens:     tk[7:8],
							}),
							Tokens: tk[7:8],
						},
						Suite: Suite{
							StatementList: &StatementList{
								Statements: []SimpleStatement{
									{
										Type: StatementAssignment,
										AssignmentStatement: &AssignmentStatement{
											StarredExpression: &StarredExpression{
												OrExpr: WrapConditional(&Atom{
													Identifier: &tk[9],
													Tokens:     tk[9:10],
												}).OrTest.AndTest.NotTest.Comparison.OrExpression,
												Tokens: tk[9:10],
											},
											Tokens: tk[9:10],
										},
										Tokens: tk[9:10],
									},
								},
								Tokens: tk[9:10],
							},
							Tokens: tk[9:10],
						},
						Tokens: tk[7:10],
					},
					{
						Expression: Expression{
							ConditionalExpression: WrapConditional(&Atom{
								Identifier: &tk[14],
								Tokens:     tk[14:15],
							}),
							Tokens: tk[14:15],
						},
						Suite: Suite{
							StatementList: &StatementList{
								Statements: []SimpleStatement{
									{
										Type: StatementAssignment,
										AssignmentStatement: &AssignmentStatement{
											StarredExpression: &StarredExpression{
												OrExpr: WrapConditional(&Atom{
													Identifier: &tk[16],
													Tokens:     tk[16:17],
												}).OrTest.AndTest.NotTest.Comparison.OrExpression,
												Tokens: tk[16:17],
											},
											Tokens: tk[16:17],
										},
										Tokens: tk[16:17],
									},
								},
								Tokens: tk[16:17],
							},
							Tokens: tk[16:17],
						},
						Tokens: tk[14:17],
					},
				},
				Tokens: tk[:17],
			}
		}},
		{"try:a\nexcept b:c\nelse:d", func(t *test, tk Tokens) { // 7
			t.Output = TryStatement{
				Try: Suite{
					StatementList: &StatementList{
						Statements: []SimpleStatement{
							{
								Type: StatementAssignment,
								AssignmentStatement: &AssignmentStatement{
									StarredExpression: &StarredExpression{
										OrExpr: WrapConditional(&Atom{
											Identifier: &tk[2],
											Tokens:     tk[2:3],
										}).OrTest.AndTest.NotTest.Comparison.OrExpression,
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
				Except: []Except{
					{
						Expression: Expression{
							ConditionalExpression: WrapConditional(&Atom{
								Identifier: &tk[6],
								Tokens:     tk[6:7],
							}),
							Tokens: tk[6:7],
						},
						Suite: Suite{
							StatementList: &StatementList{
								Statements: []SimpleStatement{
									{
										Type: StatementAssignment,
										AssignmentStatement: &AssignmentStatement{
											StarredExpression: &StarredExpression{
												OrExpr: WrapConditional(&Atom{
													Identifier: &tk[8],
													Tokens:     tk[8:9],
												}).OrTest.AndTest.NotTest.Comparison.OrExpression,
												Tokens: tk[8:9],
											},
											Tokens: tk[8:9],
										},
										Tokens: tk[8:9],
									},
								},
								Tokens: tk[8:9],
							},
							Tokens: tk[8:9],
						},
						Tokens: tk[6:9],
					},
				},
				Else: &Suite{
					StatementList: &StatementList{
						Statements: []SimpleStatement{
							{
								Type: StatementAssignment,
								AssignmentStatement: &AssignmentStatement{
									StarredExpression: &StarredExpression{
										OrExpr: WrapConditional(&Atom{
											Identifier: &tk[12],
											Tokens:     tk[12:13],
										}).OrTest.AndTest.NotTest.Comparison.OrExpression,
										Tokens: tk[12:13],
									},
									Tokens: tk[12:13],
								},
								Tokens: tk[12:13],
							},
						},
						Tokens: tk[12:13],
					},
					Tokens: tk[12:13],
				},
				Tokens: tk[:13],
			}
		}},
		{"try:a\nexcept b:c\nelse :d", func(t *test, tk Tokens) { // 8
			t.Output = TryStatement{
				Try: Suite{
					StatementList: &StatementList{
						Statements: []SimpleStatement{
							{
								Type: StatementAssignment,
								AssignmentStatement: &AssignmentStatement{
									StarredExpression: &StarredExpression{
										OrExpr: WrapConditional(&Atom{
											Identifier: &tk[2],
											Tokens:     tk[2:3],
										}).OrTest.AndTest.NotTest.Comparison.OrExpression,
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
				Except: []Except{
					{
						Expression: Expression{
							ConditionalExpression: WrapConditional(&Atom{
								Identifier: &tk[6],
								Tokens:     tk[6:7],
							}),
							Tokens: tk[6:7],
						},
						Suite: Suite{
							StatementList: &StatementList{
								Statements: []SimpleStatement{
									{
										Type: StatementAssignment,
										AssignmentStatement: &AssignmentStatement{
											StarredExpression: &StarredExpression{
												OrExpr: WrapConditional(&Atom{
													Identifier: &tk[8],
													Tokens:     tk[8:9],
												}).OrTest.AndTest.NotTest.Comparison.OrExpression,
												Tokens: tk[8:9],
											},
											Tokens: tk[8:9],
										},
										Tokens: tk[8:9],
									},
								},
								Tokens: tk[8:9],
							},
							Tokens: tk[8:9],
						},
						Tokens: tk[6:9],
					},
				},
				Else: &Suite{
					StatementList: &StatementList{
						Statements: []SimpleStatement{
							{
								Type: StatementAssignment,
								AssignmentStatement: &AssignmentStatement{
									StarredExpression: &StarredExpression{
										OrExpr: WrapConditional(&Atom{
											Identifier: &tk[13],
											Tokens:     tk[13:14],
										}).OrTest.AndTest.NotTest.Comparison.OrExpression,
										Tokens: tk[13:14],
									},
									Tokens: tk[13:14],
								},
								Tokens: tk[13:14],
							},
						},
						Tokens: tk[13:14],
					},
					Tokens: tk[13:14],
				},
				Tokens: tk[:14],
			}
		}},
		{"try:a\nexcept b:c\nelse:d\nfinally:e", func(t *test, tk Tokens) { // 9
			t.Output = TryStatement{
				Try: Suite{
					StatementList: &StatementList{
						Statements: []SimpleStatement{
							{
								Type: StatementAssignment,
								AssignmentStatement: &AssignmentStatement{
									StarredExpression: &StarredExpression{
										OrExpr: WrapConditional(&Atom{
											Identifier: &tk[2],
											Tokens:     tk[2:3],
										}).OrTest.AndTest.NotTest.Comparison.OrExpression,
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
				Except: []Except{
					{
						Expression: Expression{
							ConditionalExpression: WrapConditional(&Atom{
								Identifier: &tk[6],
								Tokens:     tk[6:7],
							}),
							Tokens: tk[6:7],
						},
						Suite: Suite{
							StatementList: &StatementList{
								Statements: []SimpleStatement{
									{
										Type: StatementAssignment,
										AssignmentStatement: &AssignmentStatement{
											StarredExpression: &StarredExpression{
												OrExpr: WrapConditional(&Atom{
													Identifier: &tk[8],
													Tokens:     tk[8:9],
												}).OrTest.AndTest.NotTest.Comparison.OrExpression,
												Tokens: tk[8:9],
											},
											Tokens: tk[8:9],
										},
										Tokens: tk[8:9],
									},
								},
								Tokens: tk[8:9],
							},
							Tokens: tk[8:9],
						},
						Tokens: tk[6:9],
					},
				},
				Else: &Suite{
					StatementList: &StatementList{
						Statements: []SimpleStatement{
							{
								Type: StatementAssignment,
								AssignmentStatement: &AssignmentStatement{
									StarredExpression: &StarredExpression{
										OrExpr: WrapConditional(&Atom{
											Identifier: &tk[12],
											Tokens:     tk[12:13],
										}).OrTest.AndTest.NotTest.Comparison.OrExpression,
										Tokens: tk[12:13],
									},
									Tokens: tk[12:13],
								},
								Tokens: tk[12:13],
							},
						},
						Tokens: tk[12:13],
					},
					Tokens: tk[12:13],
				},
				Finally: &Suite{
					StatementList: &StatementList{
						Statements: []SimpleStatement{
							{
								Type: StatementAssignment,
								AssignmentStatement: &AssignmentStatement{
									StarredExpression: &StarredExpression{
										OrExpr: WrapConditional(&Atom{
											Identifier: &tk[16],
											Tokens:     tk[16:17],
										}).OrTest.AndTest.NotTest.Comparison.OrExpression,
										Tokens: tk[16:17],
									},
									Tokens: tk[16:17],
								},
								Tokens: tk[16:17],
							},
						},
						Tokens: tk[16:17],
					},
					Tokens: tk[16:17],
				},
				Tokens: tk[:17],
			}
		}},
		{"try:a\nexcept b:c\nfinally:d", func(t *test, tk Tokens) { // 10
			t.Output = TryStatement{
				Try: Suite{
					StatementList: &StatementList{
						Statements: []SimpleStatement{
							{
								Type: StatementAssignment,
								AssignmentStatement: &AssignmentStatement{
									StarredExpression: &StarredExpression{
										OrExpr: WrapConditional(&Atom{
											Identifier: &tk[2],
											Tokens:     tk[2:3],
										}).OrTest.AndTest.NotTest.Comparison.OrExpression,
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
				Except: []Except{
					{
						Expression: Expression{
							ConditionalExpression: WrapConditional(&Atom{
								Identifier: &tk[6],
								Tokens:     tk[6:7],
							}),
							Tokens: tk[6:7],
						},
						Suite: Suite{
							StatementList: &StatementList{
								Statements: []SimpleStatement{
									{
										Type: StatementAssignment,
										AssignmentStatement: &AssignmentStatement{
											StarredExpression: &StarredExpression{
												OrExpr: WrapConditional(&Atom{
													Identifier: &tk[8],
													Tokens:     tk[8:9],
												}).OrTest.AndTest.NotTest.Comparison.OrExpression,
												Tokens: tk[8:9],
											},
											Tokens: tk[8:9],
										},
										Tokens: tk[8:9],
									},
								},
								Tokens: tk[8:9],
							},
							Tokens: tk[8:9],
						},
						Tokens: tk[6:9],
					},
				},
				Finally: &Suite{
					StatementList: &StatementList{
						Statements: []SimpleStatement{
							{
								Type: StatementAssignment,
								AssignmentStatement: &AssignmentStatement{
									StarredExpression: &StarredExpression{
										OrExpr: WrapConditional(&Atom{
											Identifier: &tk[12],
											Tokens:     tk[12:13],
										}).OrTest.AndTest.NotTest.Comparison.OrExpression,
										Tokens: tk[12:13],
									},
									Tokens: tk[12:13],
								},
								Tokens: tk[12:13],
							},
						},
						Tokens: tk[12:13],
					},
					Tokens: tk[12:13],
				},
				Tokens: tk[:13],
			}
		}},
		{"try:a\nfinally:b", func(t *test, tk Tokens) { // 11
			t.Output = TryStatement{
				Try: Suite{
					StatementList: &StatementList{
						Statements: []SimpleStatement{
							{
								Type: StatementAssignment,
								AssignmentStatement: &AssignmentStatement{
									StarredExpression: &StarredExpression{
										OrExpr: WrapConditional(&Atom{
											Identifier: &tk[2],
											Tokens:     tk[2:3],
										}).OrTest.AndTest.NotTest.Comparison.OrExpression,
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
				Finally: &Suite{
					StatementList: &StatementList{
						Statements: []SimpleStatement{
							{
								Type: StatementAssignment,
								AssignmentStatement: &AssignmentStatement{
									StarredExpression: &StarredExpression{
										OrExpr: WrapConditional(&Atom{
											Identifier: &tk[6],
											Tokens:     tk[6:7],
										}).OrTest.AndTest.NotTest.Comparison.OrExpression,
										Tokens: tk[6:7],
									},
									Tokens: tk[6:7],
								},
								Tokens: tk[6:7],
							},
						},
						Tokens: tk[6:7],
					},
					Tokens: tk[6:7],
				},
				Tokens: tk[:7],
			}
		}},
		{"try a\nfinally:b", func(t *test, tk Tokens) { // 12
			t.Err = Error{
				Err:     ErrMissingColon,
				Parsing: "TryStatement",
				Token:   tk[2],
			}
		}},
		{"try:nonlocal\nfinally:b", func(t *test, tk Tokens) { // 13
			t.Err = Error{
				Err: Error{
					Err: Error{
						Err: Error{
							Err: Error{
								Err:     ErrMissingIdentifier,
								Parsing: "NonLocalStatement",
								Token:   tk[3],
							},
							Parsing: "SimpleStatement",
							Token:   tk[2],
						},
						Parsing: "StatementList",
						Token:   tk[2],
					},
					Parsing: "Suite",
					Token:   tk[2],
				},
				Parsing: "TryStatement",
				Token:   tk[2],
			}
		}},
		{"try:a\nexcept nonlocal:b", func(t *test, tk Tokens) { // 14
			t.Err = Error{
				Err: Error{
					Err: Error{
						Err: wrapConditionalExpressionError(Error{
							Err:     ErrInvalidEnclosure,
							Parsing: "Enclosure",
							Token:   tk[6],
						}),
						Parsing: "Expression",
						Token:   tk[6],
					},
					Parsing: "Except",
					Token:   tk[6],
				},
				Parsing: "TryStatement",
				Token:   tk[6],
			}
		}},
		{"try:a\nexcept b:c\nexcept *d:e", func(t *test, tk Tokens) { // 15
			t.Err = Error{
				Err:     ErrMismatchedGroups,
				Parsing: "TryStatement",
				Token:   tk[13],
			}
		}},
		{"try:a\nexcept *b:c\nexcept d:e", func(t *test, tk Tokens) { // 16
			t.Err = Error{
				Err:     ErrMismatchedGroups,
				Parsing: "TryStatement",
				Token:   tk[13],
			}
		}},
		{"try:a\nexcept b:c\nelse d\nfinally:e", func(t *test, tk Tokens) { // 17
			t.Err = Error{
				Err:     ErrMissingColon,
				Parsing: "TryStatement",
				Token:   tk[12],
			}
		}},
		{"try:a\nexcept b:c\nelse:nonlocal\nfinally:e", func(t *test, tk Tokens) { // 18
			t.Err = Error{
				Err: Error{
					Err: Error{
						Err: Error{
							Err: Error{
								Err:     ErrMissingIdentifier,
								Parsing: "NonLocalStatement",
								Token:   tk[13],
							},
							Parsing: "SimpleStatement",
							Token:   tk[12],
						},
						Parsing: "StatementList",
						Token:   tk[12],
					},
					Parsing: "Suite",
					Token:   tk[12],
				},
				Parsing: "TryStatement",
				Token:   tk[12],
			}
		}},
		{"try:a\nexcept b:c\nfinally d", func(t *test, tk Tokens) { // 19
			t.Err = Error{
				Err:     ErrMissingColon,
				Parsing: "TryStatement",
				Token:   tk[12],
			}
		}},
		{"try:a\nexcept b:c\nfinally:nonlocal", func(t *test, tk Tokens) { // 20
			t.Err = Error{
				Err: Error{
					Err: Error{
						Err: Error{
							Err: Error{
								Err:     ErrMissingIdentifier,
								Parsing: "NonLocalStatement",
								Token:   tk[13],
							},
							Parsing: "SimpleStatement",
							Token:   tk[12],
						},
						Parsing: "StatementList",
						Token:   tk[12],
					},
					Parsing: "Suite",
					Token:   tk[12],
				},
				Parsing: "TryStatement",
				Token:   tk[12],
			}
		}},
		{"try:a", func(t *test, tk Tokens) { // 21
			t.Err = Error{
				Err:     ErrMissingFinally,
				Parsing: "TryStatement",
				Token:   tk[3],
			}
		}},
	}, func(t *test) (Type, error) {
		var ts TryStatement

		err := ts.parse(t.Tokens)

		return ts, err
	})
}

func TestExcept(t *testing.T) {
	doTests(t, []sourceFn{
		{`a:b`, func(t *test, tk Tokens) { // 1
			t.Output = Except{
				Expression: Expression{
					ConditionalExpression: WrapConditional(&Atom{
						Identifier: &tk[0],
						Tokens:     tk[:1],
					}),
					Tokens: tk[:1],
				},
				Suite: Suite{
					StatementList: &StatementList{
						Statements: []SimpleStatement{
							{
								Type: StatementAssignment,
								AssignmentStatement: &AssignmentStatement{
									StarredExpression: &StarredExpression{
										OrExpr: WrapConditional(&Atom{
											Identifier: &tk[2],
											Tokens:     tk[2:3],
										}).OrTest.AndTest.NotTest.Comparison.OrExpression,
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
				Tokens: tk[:3],
			}
		}},
		{`a :b`, func(t *test, tk Tokens) { // 2
			t.Output = Except{
				Expression: Expression{
					ConditionalExpression: WrapConditional(&Atom{
						Identifier: &tk[0],
						Tokens:     tk[:1],
					}),
					Tokens: tk[:1],
				},
				Suite: Suite{
					StatementList: &StatementList{
						Statements: []SimpleStatement{
							{
								Type: StatementAssignment,
								AssignmentStatement: &AssignmentStatement{
									StarredExpression: &StarredExpression{
										OrExpr: WrapConditional(&Atom{
											Identifier: &tk[3],
											Tokens:     tk[3:4],
										}).OrTest.AndTest.NotTest.Comparison.OrExpression,
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
				Tokens: tk[:4],
			}
		}},
		{`a as b:c`, func(t *test, tk Tokens) { // 3
			t.Output = Except{
				Expression: Expression{
					ConditionalExpression: WrapConditional(&Atom{
						Identifier: &tk[0],
						Tokens:     tk[:1],
					}),
					Tokens: tk[:1],
				},
				Identifier: &tk[4],
				Suite: Suite{
					StatementList: &StatementList{
						Statements: []SimpleStatement{
							{
								Type: StatementAssignment,
								AssignmentStatement: &AssignmentStatement{
									StarredExpression: &StarredExpression{
										OrExpr: WrapConditional(&Atom{
											Identifier: &tk[6],
											Tokens:     tk[6:7],
										}).OrTest.AndTest.NotTest.Comparison.OrExpression,
										Tokens: tk[6:7],
									},
									Tokens: tk[6:7],
								},
								Tokens: tk[6:7],
							},
						},
						Tokens: tk[6:7],
					},
					Tokens: tk[6:7],
				},
				Tokens: tk[:7],
			}
		}},
		{`nonlocal:a`, func(t *test, tk Tokens) { // 4
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
				Parsing: "Except",
				Token:   tk[0],
			}
		}},
		{`a as nonlocal:b`, func(t *test, tk Tokens) { // 5
			t.Err = Error{
				Err:     ErrMissingIdentifier,
				Parsing: "Except",
				Token:   tk[4],
			}
		}},
		{`a:nonlocal`, func(t *test, tk Tokens) { // 5
			t.Err = Error{
				Err: Error{
					Err: Error{
						Err: Error{
							Err: Error{
								Err:     ErrMissingIdentifier,
								Parsing: "NonLocalStatement",
								Token:   tk[3],
							},
							Parsing: "SimpleStatement",
							Token:   tk[2],
						},
						Parsing: "StatementList",
						Token:   tk[2],
					},
					Parsing: "Suite",
					Token:   tk[2],
				},
				Parsing: "Except",
				Token:   tk[2],
			}
		}},
		{`a b`, func(t *test, tk Tokens) { // 6
			t.Err = Error{
				Err:     ErrMissingColon,
				Parsing: "Except",
				Token:   tk[2],
			}
		}},
	}, func(t *test) (Type, error) {
		var w Except

		err := w.parse(t.Tokens)

		return w, err
	})
}

func TestWithStatement(t *testing.T) {
	doTests(t, []sourceFn{
		{`with a:b`, func(t *test, tk Tokens) { // 1
			t.Output = WithStatement{
				Contents: WithStatementContents{
					Items: []WithItem{
						{
							Expression: Expression{
								ConditionalExpression: WrapConditional(&Atom{
									Identifier: &tk[2],
									Tokens:     tk[2:3],
								}),
								Tokens: tk[2:3],
							},
							Tokens: tk[2:3],
						},
					},
					Tokens: tk[2:3],
				},
				Suite: Suite{
					StatementList: &StatementList{
						Statements: []SimpleStatement{
							{
								Type: StatementAssignment,
								AssignmentStatement: &AssignmentStatement{
									StarredExpression: &StarredExpression{
										OrExpr: WrapConditional(&Atom{
											Identifier: &tk[4],
											Tokens:     tk[4:5],
										}).OrTest.AndTest.NotTest.Comparison.OrExpression,
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
				Tokens: tk[:5],
			}
		}},
		{`with a : b`, func(t *test, tk Tokens) { // 2
			t.Output = WithStatement{
				Contents: WithStatementContents{
					Items: []WithItem{
						{
							Expression: Expression{
								ConditionalExpression: WrapConditional(&Atom{
									Identifier: &tk[2],
									Tokens:     tk[2:3],
								}),
								Tokens: tk[2:3],
							},
							Tokens: tk[2:3],
						},
					},
					Tokens: tk[2:3],
				},
				Suite: Suite{
					StatementList: &StatementList{
						Statements: []SimpleStatement{
							{
								Type: StatementAssignment,
								AssignmentStatement: &AssignmentStatement{
									StarredExpression: &StarredExpression{
										OrExpr: WrapConditional(&Atom{
											Identifier: &tk[6],
											Tokens:     tk[6:7],
										}).OrTest.AndTest.NotTest.Comparison.OrExpression,
										Tokens: tk[6:7],
									},
									Tokens: tk[6:7],
								},
								Tokens: tk[6:7],
							},
						},
						Tokens: tk[6:7],
					},
					Tokens: tk[6:7],
				},
				Tokens: tk[:7],
			}
		}},
		{`with a,b:c`, func(t *test, tk Tokens) { // 3
			t.Output = WithStatement{
				Contents: WithStatementContents{
					Items: []WithItem{
						{
							Expression: Expression{
								ConditionalExpression: WrapConditional(&Atom{
									Identifier: &tk[2],
									Tokens:     tk[2:3],
								}),
								Tokens: tk[2:3],
							},
							Tokens: tk[2:3],
						},
						{
							Expression: Expression{
								ConditionalExpression: WrapConditional(&Atom{
									Identifier: &tk[4],
									Tokens:     tk[4:5],
								}),
								Tokens: tk[4:5],
							},
							Tokens: tk[4:5],
						},
					},
					Tokens: tk[2:5],
				},
				Suite: Suite{
					StatementList: &StatementList{
						Statements: []SimpleStatement{
							{
								Type: StatementAssignment,
								AssignmentStatement: &AssignmentStatement{
									StarredExpression: &StarredExpression{
										OrExpr: WrapConditional(&Atom{
											Identifier: &tk[6],
											Tokens:     tk[6:7],
										}).OrTest.AndTest.NotTest.Comparison.OrExpression,
										Tokens: tk[6:7],
									},
									Tokens: tk[6:7],
								},
								Tokens: tk[6:7],
							},
						},
						Tokens: tk[6:7],
					},
					Tokens: tk[6:7],
				},
				Tokens: tk[:7],
			}
		}},
		{`with (a,b):c`, func(t *test, tk Tokens) { // 4
			t.Output = WithStatement{
				Contents: WithStatementContents{
					Items: []WithItem{
						{
							Expression: Expression{
								ConditionalExpression: WrapConditional(&Atom{
									Identifier: &tk[3],
									Tokens:     tk[3:4],
								}),
								Tokens: tk[3:4],
							},
							Tokens: tk[3:4],
						},
						{
							Expression: Expression{
								ConditionalExpression: WrapConditional(&Atom{
									Identifier: &tk[5],
									Tokens:     tk[5:6],
								}),
								Tokens: tk[5:6],
							},
							Tokens: tk[5:6],
						},
					},
					Tokens: tk[3:6],
				},
				Suite: Suite{
					StatementList: &StatementList{
						Statements: []SimpleStatement{
							{
								Type: StatementAssignment,
								AssignmentStatement: &AssignmentStatement{
									StarredExpression: &StarredExpression{
										OrExpr: WrapConditional(&Atom{
											Identifier: &tk[8],
											Tokens:     tk[8:9],
										}).OrTest.AndTest.NotTest.Comparison.OrExpression,
										Tokens: tk[8:9],
									},
									Tokens: tk[8:9],
								},
								Tokens: tk[8:9],
							},
						},
						Tokens: tk[8:9],
					},
					Tokens: tk[8:9],
				},
				Tokens: tk[:9],
			}
		}},
		{`with (a,b , ):c`, func(t *test, tk Tokens) { // 5
			t.Output = WithStatement{
				Contents: WithStatementContents{
					Items: []WithItem{
						{
							Expression: Expression{
								ConditionalExpression: WrapConditional(&Atom{
									Identifier: &tk[3],
									Tokens:     tk[3:4],
								}),
								Tokens: tk[3:4],
							},
							Tokens: tk[3:4],
						},
						{
							Expression: Expression{
								ConditionalExpression: WrapConditional(&Atom{
									Identifier: &tk[5],
									Tokens:     tk[5:6],
								}),
								Tokens: tk[5:6],
							},
							Tokens: tk[5:6],
						},
					},
					Tokens: tk[3:6],
				},
				Suite: Suite{
					StatementList: &StatementList{
						Statements: []SimpleStatement{
							{
								Type: StatementAssignment,
								AssignmentStatement: &AssignmentStatement{
									StarredExpression: &StarredExpression{
										OrExpr: WrapConditional(&Atom{
											Identifier: &tk[11],
											Tokens:     tk[11:12],
										}).OrTest.AndTest.NotTest.Comparison.OrExpression,
										Tokens: tk[11:12],
									},
									Tokens: tk[11:12],
								},
								Tokens: tk[11:12],
							},
						},
						Tokens: tk[11:12],
					},
					Tokens: tk[11:12],
				},
				Tokens: tk[:12],
			}
		}},
		{`with a:b`, func(t *test, tk Tokens) { // 6
			t.Async = true
			t.Output = WithStatement{
				Async: true,
				Contents: WithStatementContents{
					Items: []WithItem{
						{
							Expression: Expression{
								ConditionalExpression: WrapConditional(&Atom{
									Identifier: &tk[2],
									Tokens:     tk[2:3],
								}),
								Tokens: tk[2:3],
							},
							Tokens: tk[2:3],
						},
					},
					Tokens: tk[2:3],
				},
				Suite: Suite{
					StatementList: &StatementList{
						Statements: []SimpleStatement{
							{
								Type: StatementAssignment,
								AssignmentStatement: &AssignmentStatement{
									StarredExpression: &StarredExpression{
										OrExpr: WrapConditional(&Atom{
											Identifier: &tk[4],
											Tokens:     tk[4:5],
										}).OrTest.AndTest.NotTest.Comparison.OrExpression,
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
				Tokens: tk[:5],
			}
		}},
		{`with nonlocal:a`, func(t *test, tk Tokens) { // 7
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
						Parsing: "WithItem",
						Token:   tk[2],
					},
					Parsing: "WithStatementContents",
					Token:   tk[2],
				},
				Parsing: "WithStatement",
				Token:   tk[2],
			}
		}},
		{`with (a:):b`, func(t *test, tk Tokens) { // 8
			t.Err = Error{
				Err:     ErrMissingClosingParen,
				Parsing: "WithStatement",
				Token:   tk[4],
			}
		}},
		{`with a b`, func(t *test, tk Tokens) { // 9
			t.Err = Error{
				Err:     ErrMissingColon,
				Parsing: "WithStatement",
				Token:   tk[4],
			}
		}},
		{`with a:nonlocal`, func(t *test, tk Tokens) { // 10
			t.Err = Error{
				Err: Error{
					Err: Error{
						Err: Error{
							Err: Error{
								Err:     ErrMissingIdentifier,
								Parsing: "NonLocalStatement",
								Token:   tk[5],
							},
							Parsing: "SimpleStatement",
							Token:   tk[4],
						},
						Parsing: "StatementList",
						Token:   tk[4],
					},
					Parsing: "Suite",
					Token:   tk[4],
				},
				Parsing: "WithStatement",
				Token:   tk[4],
			}
		}},
	}, func(t *test) (Type, error) {
		var w WithStatement

		err := w.parse(t.Tokens, t.Async)

		return w, err
	})
}

func TestWithStatementContents(t *testing.T) {
	doTests(t, []sourceFn{
		{`a`, func(t *test, tk Tokens) { // 1
			t.Output = WithStatementContents{
				Items: []WithItem{
					{
						Expression: Expression{
							ConditionalExpression: WrapConditional(&Atom{
								Identifier: &tk[0],
								Tokens:     tk[:1],
							}),
							Tokens: tk[:1],
						},
						Tokens: tk[:1],
					},
				},
				Tokens: tk[:1],
			}
		}},
		{`a,b`, func(t *test, tk Tokens) { // 2
			t.Output = WithStatementContents{
				Items: []WithItem{
					{
						Expression: Expression{
							ConditionalExpression: WrapConditional(&Atom{
								Identifier: &tk[0],
								Tokens:     tk[:1],
							}),
							Tokens: tk[:1],
						},
						Tokens: tk[:1],
					},
					{
						Expression: Expression{
							ConditionalExpression: WrapConditional(&Atom{
								Identifier: &tk[2],
								Tokens:     tk[2:3],
							}),
							Tokens: tk[2:3],
						},
						Tokens: tk[2:3],
					},
				},
				Tokens: tk[:3],
			}
		}},
		{`a , b`, func(t *test, tk Tokens) { // 3
			t.Output = WithStatementContents{
				Items: []WithItem{
					{
						Expression: Expression{
							ConditionalExpression: WrapConditional(&Atom{
								Identifier: &tk[0],
								Tokens:     tk[:1],
							}),
							Tokens: tk[:1],
						},
						Tokens: tk[:1],
					},
					{
						Expression: Expression{
							ConditionalExpression: WrapConditional(&Atom{
								Identifier: &tk[4],
								Tokens:     tk[4:5],
							}),
							Tokens: tk[4:5],
						},
						Tokens: tk[4:5],
					},
				},
				Tokens: tk[:5],
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
					Parsing: "WithItem",
					Token:   tk[0],
				},
				Parsing: "WithStatementContents",
				Token:   tk[0],
			}
		}},
	}, func(t *test) (Type, error) {
		var w WithStatementContents

		err := w.parse(t.Tokens)

		return w, err
	})
}

func TestWithItem(t *testing.T) {
	doTests(t, []sourceFn{
		{`a`, func(t *test, tk Tokens) { // 1
			t.Output = WithItem{
				Expression: Expression{
					ConditionalExpression: WrapConditional(&Atom{
						Identifier: &tk[0],
						Tokens:     tk[:1],
					}),
					Tokens: tk[:1],
				},
				Tokens: tk[:1],
			}
		}},
		{`a as b`, func(t *test, tk Tokens) { // 2
			t.Output = WithItem{
				Expression: Expression{
					ConditionalExpression: WrapConditional(&Atom{
						Identifier: &tk[0],
						Tokens:     tk[:1],
					}),
					Tokens: tk[:1],
				},
				Target: &Target{
					PrimaryExpression: &PrimaryExpression{
						Atom: &Atom{
							Identifier: &tk[4],
							Tokens:     tk[4:5],
						},
						Tokens: tk[4:5],
					},
					Tokens: tk[4:5],
				},
				Tokens: tk[:5],
			}
		}},
		{`nonlocal`, func(t *test, tk Tokens) { // 3
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
				Parsing: "WithItem",
				Token:   tk[0],
			}
		}},
		{`a as nonlocal`, func(t *test, tk Tokens) { // 4
			t.Err = Error{
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
				Parsing: "WithItem",
				Token:   tk[4],
			}
		}},
	}, func(t *test) (Type, error) {
		var w WithItem

		err := w.parse(t.Tokens)

		return w, err
	})
}

func TestFuncDefinition(t *testing.T) {
	doTests(t, []sourceFn{
		{`def a():b`, func(t *test, tk Tokens) { // 1
			t.Output = FuncDefinition{
				FuncName: &tk[2],
				ParameterList: ParameterList{
					Tokens: tk[4:4],
				},
				Suite: Suite{
					StatementList: &StatementList{
						Statements: []SimpleStatement{
							{
								Type: StatementAssignment,
								AssignmentStatement: &AssignmentStatement{
									StarredExpression: &StarredExpression{
										OrExpr: WrapConditional(&Atom{
											Identifier: &tk[6],
											Tokens:     tk[6:7],
										}).OrTest.AndTest.NotTest.Comparison.OrExpression,
										Tokens: tk[6:7],
									},
									Tokens: tk[6:7],
								},
								Tokens: tk[6:7],
							},
						},
						Tokens: tk[6:7],
					},
					Tokens: tk[6:7],
				},
				Tokens: tk[:7],
			}
		}},
		{`def a () : b`, func(t *test, tk Tokens) { // 2
			t.Output = FuncDefinition{
				FuncName: &tk[2],
				ParameterList: ParameterList{
					Tokens: tk[5:5],
				},
				Suite: Suite{
					StatementList: &StatementList{
						Statements: []SimpleStatement{
							{
								Type: StatementAssignment,
								AssignmentStatement: &AssignmentStatement{
									StarredExpression: &StarredExpression{
										OrExpr: WrapConditional(&Atom{
											Identifier: &tk[9],
											Tokens:     tk[9:10],
										}).OrTest.AndTest.NotTest.Comparison.OrExpression,
										Tokens: tk[9:10],
									},
									Tokens: tk[9:10],
								},
								Tokens: tk[9:10],
							},
						},
						Tokens: tk[9:10],
					},
					Tokens: tk[9:10],
				},
				Tokens: tk[:10],
			}
		}},
		{`def a():b`, func(t *test, tk Tokens) { // 3
			t.Async = true
			t.Output = FuncDefinition{
				Async:    true,
				FuncName: &tk[2],
				ParameterList: ParameterList{
					Tokens: tk[4:4],
				},
				Suite: Suite{
					StatementList: &StatementList{
						Statements: []SimpleStatement{
							{
								Type: StatementAssignment,
								AssignmentStatement: &AssignmentStatement{
									StarredExpression: &StarredExpression{
										OrExpr: WrapConditional(&Atom{
											Identifier: &tk[6],
											Tokens:     tk[6:7],
										}).OrTest.AndTest.NotTest.Comparison.OrExpression,
										Tokens: tk[6:7],
									},
									Tokens: tk[6:7],
								},
								Tokens: tk[6:7],
							},
						},
						Tokens: tk[6:7],
					},
					Tokens: tk[6:7],
				},
				Tokens: tk[:7],
			}
		}},
		{`def a():b`, func(t *test, tk Tokens) { // 4
			t.Decorators = &Decorators{
				Decorators: []AssignmentExpression{},
			}
			t.Output = FuncDefinition{
				Decorators: t.Decorators,
				FuncName:   &tk[2],
				ParameterList: ParameterList{
					Tokens: tk[4:4],
				},
				Suite: Suite{
					StatementList: &StatementList{
						Statements: []SimpleStatement{
							{
								Type: StatementAssignment,
								AssignmentStatement: &AssignmentStatement{
									StarredExpression: &StarredExpression{
										OrExpr: WrapConditional(&Atom{
											Identifier: &tk[6],
											Tokens:     tk[6:7],
										}).OrTest.AndTest.NotTest.Comparison.OrExpression,
										Tokens: tk[6:7],
									},
									Tokens: tk[6:7],
								},
								Tokens: tk[6:7],
							},
						},
						Tokens: tk[6:7],
					},
					Tokens: tk[6:7],
				},
				Tokens: tk[:7],
			}
		}},
		{`def a[b]():c`, func(t *test, tk Tokens) { // 5
			t.Output = FuncDefinition{
				FuncName: &tk[2],
				TypeParams: &TypeParams{
					TypeParams: []TypeParam{
						{
							Identifier: &tk[4],
							Tokens:     tk[4:5],
						},
					},
					Tokens: tk[3:6],
				},
				ParameterList: ParameterList{
					Tokens: tk[7:7],
				},
				Suite: Suite{
					StatementList: &StatementList{
						Statements: []SimpleStatement{
							{
								Type: StatementAssignment,
								AssignmentStatement: &AssignmentStatement{
									StarredExpression: &StarredExpression{
										OrExpr: WrapConditional(&Atom{
											Identifier: &tk[9],
											Tokens:     tk[9:10],
										}).OrTest.AndTest.NotTest.Comparison.OrExpression,
										Tokens: tk[9:10],
									},
									Tokens: tk[9:10],
								},
								Tokens: tk[9:10],
							},
						},
						Tokens: tk[9:10],
					},
					Tokens: tk[9:10],
				},
				Tokens: tk[:10],
			}
		}},
		{`def a(b):c`, func(t *test, tk Tokens) { // 6
			t.Output = FuncDefinition{
				FuncName: &tk[2],
				ParameterList: ParameterList{
					NoPosOnly: []DefParameter{
						{
							Parameter: Parameter{
								Identifier: &tk[4],
								Tokens:     tk[4:5],
							},
							Tokens: tk[4:5],
						},
					},
					Tokens: tk[4:5],
				},
				Suite: Suite{
					StatementList: &StatementList{
						Statements: []SimpleStatement{
							{
								Type: StatementAssignment,
								AssignmentStatement: &AssignmentStatement{
									StarredExpression: &StarredExpression{
										OrExpr: WrapConditional(&Atom{
											Identifier: &tk[7],
											Tokens:     tk[7:8],
										}).OrTest.AndTest.NotTest.Comparison.OrExpression,
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
				Tokens: tk[:8],
			}
		}},
		{`def a()->b:c`, func(t *test, tk Tokens) { // 7
			t.Output = FuncDefinition{
				FuncName: &tk[2],
				ParameterList: ParameterList{
					Tokens: tk[4:4],
				},
				Expression: &Expression{
					ConditionalExpression: WrapConditional(&Atom{
						Identifier: &tk[6],
						Tokens:     tk[6:7],
					}),
					Tokens: tk[6:7],
				},
				Suite: Suite{
					StatementList: &StatementList{
						Statements: []SimpleStatement{
							{
								Type: StatementAssignment,
								AssignmentStatement: &AssignmentStatement{
									StarredExpression: &StarredExpression{
										OrExpr: WrapConditional(&Atom{
											Identifier: &tk[8],
											Tokens:     tk[8:9],
										}).OrTest.AndTest.NotTest.Comparison.OrExpression,
										Tokens: tk[8:9],
									},
									Tokens: tk[8:9],
								},
								Tokens: tk[8:9],
							},
						},
						Tokens: tk[8:9],
					},
					Tokens: tk[8:9],
				},
				Tokens: tk[:9],
			}
		}},
		{"def a [ b, c ] ( d, e ) -> f : \n\tg", func(t *test, tk Tokens) { // 8
			t.Async = true
			t.Decorators = &Decorators{
				Decorators: []AssignmentExpression{},
			}
			t.Output = FuncDefinition{
				Async:      true,
				Decorators: t.Decorators,
				FuncName:   &tk[2],
				TypeParams: &TypeParams{
					TypeParams: []TypeParam{
						{
							Identifier: &tk[6],
							Tokens:     tk[6:7],
						},
						{
							Identifier: &tk[9],
							Tokens:     tk[9:10],
						},
					},
					Tokens: tk[4:12],
				},
				ParameterList: ParameterList{
					NoPosOnly: []DefParameter{
						{
							Parameter: Parameter{
								Identifier: &tk[15],
								Tokens:     tk[15:16],
							},
							Tokens: tk[15:16],
						},
						{
							Parameter: Parameter{
								Identifier: &tk[18],
								Tokens:     tk[18:19],
							},
							Tokens: tk[18:19],
						},
					},
					Tokens: tk[15:19],
				},
				Expression: &Expression{
					ConditionalExpression: WrapConditional(&Atom{
						Identifier: &tk[24],
						Tokens:     tk[24:25],
					}),
					Tokens: tk[24:25],
				},
				Suite: Suite{
					Statements: []Statement{
						{
							StatementList: &StatementList{
								Statements: []SimpleStatement{
									{
										Type: StatementAssignment,
										AssignmentStatement: &AssignmentStatement{
											StarredExpression: &StarredExpression{
												OrExpr: WrapConditional(&Atom{
													Identifier: &tk[30],
													Tokens:     tk[30:31],
												}).OrTest.AndTest.NotTest.Comparison.OrExpression,
												Tokens: tk[30:31],
											},
											Tokens: tk[30:31],
										},
										Tokens: tk[30:31],
									},
								},
								Tokens: tk[30:31],
							},
							Tokens: tk[30:31],
						},
					},
					Tokens: tk[28:32],
				},
				Tokens: tk[:32],
			}
		}},
		{"def nonlocal():a", func(t *test, tk Tokens) { // 9
			t.Err = Error{
				Err:     ErrMissingIdentifier,
				Parsing: "FuncDefinition",
				Token:   tk[2],
			}
		}},
		{"def a[nonlocal]():b", func(t *test, tk Tokens) { // 10
			t.Err = Error{
				Err: Error{
					Err: Error{
						Err:     ErrMissingIdentifier,
						Parsing: "TypeParam",
						Token:   tk[4],
					},
					Parsing: "TypeParams",
					Token:   tk[4],
				},
				Parsing: "FuncDefinition",
				Token:   tk[3],
			}
		}},
		{"def a(nonlocal):b", func(t *test, tk Tokens) { // 11
			t.Err = Error{
				Err:     ErrMissingClosingParen,
				Parsing: "FuncDefinition",
				Token:   tk[4],
			}
		}},
		{"def a(b=nonlocal):c", func(t *test, tk Tokens) { // 12
			t.Err = Error{
				Err: Error{
					Err: Error{
						Err: Error{
							Err: wrapConditionalExpressionError(Error{
								Err:     ErrInvalidEnclosure,
								Parsing: "Enclosure",
								Token:   tk[6],
							}),
							Parsing: "Expression",
							Token:   tk[6],
						},
						Parsing: "DefParameter",
						Token:   tk[6],
					},
					Parsing: "ParameterList",
					Token:   tk[4],
				},
				Parsing: "FuncDefinition",
				Token:   tk[4],
			}
		}},
		{"def a()->nonlocal:b", func(t *test, tk Tokens) { // 13
			t.Err = Error{
				Err: Error{
					Err: wrapConditionalExpressionError(Error{
						Err:     ErrInvalidEnclosure,
						Parsing: "Enclosure",
						Token:   tk[6],
					}),
					Parsing: "Expression",
					Token:   tk[6],
				},
				Parsing: "FuncDefinition",
				Token:   tk[6],
			}
		}},
		{"def a():nonlocal", func(t *test, tk Tokens) { // 14
			t.Err = Error{
				Err: Error{
					Err: Error{
						Err: Error{
							Err: Error{
								Err:     ErrMissingIdentifier,
								Parsing: "NonLocalStatement",
								Token:   tk[7],
							},
							Parsing: "SimpleStatement",
							Token:   tk[6],
						},
						Parsing: "StatementList",
						Token:   tk[6],
					},
					Parsing: "Suite",
					Token:   tk[6],
				},
				Parsing: "FuncDefinition",
				Token:   tk[6],
			}
		}},
	}, func(t *test) (Type, error) {
		var f FuncDefinition

		err := f.parse(t.Tokens, t.Async, t.Decorators)

		return f, err
	})
}

func TestClassDefinition(t *testing.T) {
	doTests(t, []sourceFn{
		{`class a:b`, func(t *test, tk Tokens) { // 1
			t.Output = ClassDefinition{
				ClassName: &tk[2],
				Suite: Suite{
					StatementList: &StatementList{
						Statements: []SimpleStatement{
							{
								Type: StatementAssignment,
								AssignmentStatement: &AssignmentStatement{
									StarredExpression: &StarredExpression{
										OrExpr: WrapConditional(&Atom{
											Identifier: &tk[4],
											Tokens:     tk[4:5],
										}).OrTest.AndTest.NotTest.Comparison.OrExpression,
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
				Tokens: tk[:5],
			}
		}},
		{`class a:b`, func(t *test, tk Tokens) { // 2
			t.Decorators = &Decorators{
				Decorators: []AssignmentExpression{},
			}
			t.Output = ClassDefinition{
				Decorators: t.Decorators,
				ClassName:  &tk[2],
				Suite: Suite{
					StatementList: &StatementList{
						Statements: []SimpleStatement{
							{
								Type: StatementAssignment,
								AssignmentStatement: &AssignmentStatement{
									StarredExpression: &StarredExpression{
										OrExpr: WrapConditional(&Atom{
											Identifier: &tk[4],
											Tokens:     tk[4:5],
										}).OrTest.AndTest.NotTest.Comparison.OrExpression,
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
				Tokens: tk[:5],
			}
		}},
		{`class a: b`, func(t *test, tk Tokens) { // 3
			t.Output = ClassDefinition{
				ClassName: &tk[2],
				Suite: Suite{
					StatementList: &StatementList{
						Statements: []SimpleStatement{
							{
								Type: StatementAssignment,
								AssignmentStatement: &AssignmentStatement{
									StarredExpression: &StarredExpression{
										OrExpr: WrapConditional(&Atom{
											Identifier: &tk[5],
											Tokens:     tk[5:6],
										}).OrTest.AndTest.NotTest.Comparison.OrExpression,
										Tokens: tk[5:6],
									},
									Tokens: tk[5:6],
								},
								Tokens: tk[5:6],
							},
						},
						Tokens: tk[5:6],
					},
					Tokens: tk[5:6],
				},
				Tokens: tk[:6],
			}
		}},
		{`class a :b`, func(t *test, tk Tokens) { // 4
			t.Output = ClassDefinition{
				ClassName: &tk[2],
				Suite: Suite{
					StatementList: &StatementList{
						Statements: []SimpleStatement{
							{
								Type: StatementAssignment,
								AssignmentStatement: &AssignmentStatement{
									StarredExpression: &StarredExpression{
										OrExpr: WrapConditional(&Atom{
											Identifier: &tk[5],
											Tokens:     tk[5:6],
										}).OrTest.AndTest.NotTest.Comparison.OrExpression,
										Tokens: tk[5:6],
									},
									Tokens: tk[5:6],
								},
								Tokens: tk[5:6],
							},
						},
						Tokens: tk[5:6],
					},
					Tokens: tk[5:6],
				},
				Tokens: tk[:6],
			}
		}},
		{"class a:\n\tb", func(t *test, tk Tokens) { // 5
			t.Output = ClassDefinition{
				ClassName: &tk[2],
				Suite: Suite{
					Statements: []Statement{
						{
							StatementList: &StatementList{
								Statements: []SimpleStatement{
									{
										Type: StatementAssignment,
										AssignmentStatement: &AssignmentStatement{
											StarredExpression: &StarredExpression{
												OrExpr: WrapConditional(&Atom{
													Identifier: &tk[6],
													Tokens:     tk[6:7],
												}).OrTest.AndTest.NotTest.Comparison.OrExpression,
												Tokens: tk[6:7],
											},
											Tokens: tk[6:7],
										},
										Tokens: tk[6:7],
									},
								},
								Tokens: tk[6:7],
							},
							Tokens: tk[6:7],
						},
					},
					Tokens: tk[4:8],
				},
				Tokens: tk[:8],
			}
		}},
		{`class a():b`, func(t *test, tk Tokens) { // 6
			t.Output = ClassDefinition{
				ClassName: &tk[2],
				Inheritance: &ArgumentList{
					Tokens: tk[4:4],
				},
				Suite: Suite{
					StatementList: &StatementList{
						Statements: []SimpleStatement{
							{
								Type: StatementAssignment,
								AssignmentStatement: &AssignmentStatement{
									StarredExpression: &StarredExpression{
										OrExpr: WrapConditional(&Atom{
											Identifier: &tk[6],
											Tokens:     tk[6:7],
										}).OrTest.AndTest.NotTest.Comparison.OrExpression,
										Tokens: tk[6:7],
									},
									Tokens: tk[6:7],
								},
								Tokens: tk[6:7],
							},
						},
						Tokens: tk[6:7],
					},
					Tokens: tk[6:7],
				},
				Tokens: tk[:7],
			}
		}},
		{`class a ( b ) : c`, func(t *test, tk Tokens) { // 7
			t.Output = ClassDefinition{
				ClassName: &tk[2],
				Inheritance: &ArgumentList{
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
				Suite: Suite{
					StatementList: &StatementList{
						Statements: []SimpleStatement{
							{
								Type: StatementAssignment,
								AssignmentStatement: &AssignmentStatement{
									StarredExpression: &StarredExpression{
										OrExpr: WrapConditional(&Atom{
											Identifier: &tk[12],
											Tokens:     tk[12:13],
										}).OrTest.AndTest.NotTest.Comparison.OrExpression,
										Tokens: tk[12:13],
									},
									Tokens: tk[12:13],
								},
								Tokens: tk[12:13],
							},
						},
						Tokens: tk[12:13],
					},
					Tokens: tk[12:13],
				},
				Tokens: tk[:13],
			}
		}},
		{`class a [b:c]: d`, func(t *test, tk Tokens) { // 8
			t.Output = ClassDefinition{
				ClassName: &tk[2],
				TypeParams: &TypeParams{
					TypeParams: []TypeParam{
						{
							Identifier: &tk[5],
							Expression: &Expression{
								ConditionalExpression: WrapConditional(&Atom{
									Identifier: &tk[7],
									Tokens:     tk[7:8],
								}),
								Tokens: tk[7:8],
							},
							Tokens: tk[5:8],
						},
					},
					Tokens: tk[4:9],
				},
				Suite: Suite{
					StatementList: &StatementList{
						Statements: []SimpleStatement{
							{
								Type: StatementAssignment,
								AssignmentStatement: &AssignmentStatement{
									StarredExpression: &StarredExpression{
										OrExpr: WrapConditional(&Atom{
											Identifier: &tk[11],
											Tokens:     tk[11:12],
										}).OrTest.AndTest.NotTest.Comparison.OrExpression,
										Tokens: tk[11:12],
									},
									Tokens: tk[11:12],
								},
								Tokens: tk[11:12],
							},
						},
						Tokens: tk[11:12],
					},
					Tokens: tk[11:12],
				},
				Tokens: tk[:12],
			}
		}},
		{`class a [ *b ] : c`, func(t *test, tk Tokens) { // 9
			t.Output = ClassDefinition{
				ClassName: &tk[2],
				TypeParams: &TypeParams{
					TypeParams: []TypeParam{
						{
							Type:       TypeParamVar,
							Identifier: &tk[7],
							Tokens:     tk[6:8],
						},
					},
					Tokens: tk[4:10],
				},
				Suite: Suite{
					StatementList: &StatementList{
						Statements: []SimpleStatement{
							{
								Type: StatementAssignment,
								AssignmentStatement: &AssignmentStatement{
									StarredExpression: &StarredExpression{
										OrExpr: WrapConditional(&Atom{
											Identifier: &tk[13],
											Tokens:     tk[13:14],
										}).OrTest.AndTest.NotTest.Comparison.OrExpression,
										Tokens: tk[13:14],
									},
									Tokens: tk[13:14],
								},
								Tokens: tk[13:14],
							},
						},
						Tokens: tk[13:14],
					},
					Tokens: tk[13:14],
				},
				Tokens: tk[:14],
			}
		}},
		{`class a[*b](c):d`, func(t *test, tk Tokens) { // 10
			t.Output = ClassDefinition{
				ClassName: &tk[2],
				TypeParams: &TypeParams{
					TypeParams: []TypeParam{
						{
							Type:       TypeParamVar,
							Identifier: &tk[5],
							Tokens:     tk[4:6],
						},
					},
					Tokens: tk[3:7],
				},
				Inheritance: &ArgumentList{
					PositionalArguments: []PositionalArgument{
						{
							AssignmentExpression: &AssignmentExpression{
								Expression: Expression{
									ConditionalExpression: WrapConditional(&Atom{
										Identifier: &tk[8],
										Tokens:     tk[8:9],
									}),
									Tokens: tk[8:9],
								},
								Tokens: tk[8:9],
							},
							Tokens: tk[8:9],
						},
					},
					Tokens: tk[8:9],
				},
				Suite: Suite{
					StatementList: &StatementList{
						Statements: []SimpleStatement{
							{
								Type: StatementAssignment,
								AssignmentStatement: &AssignmentStatement{
									StarredExpression: &StarredExpression{
										OrExpr: WrapConditional(&Atom{
											Identifier: &tk[11],
											Tokens:     tk[11:12],
										}).OrTest.AndTest.NotTest.Comparison.OrExpression,
										Tokens: tk[11:12],
									},
									Tokens: tk[11:12],
								},
								Tokens: tk[11:12],
							},
						},
						Tokens: tk[11:12],
					},
					Tokens: tk[11:12],
				},
				Tokens: tk[:12],
			}
		}},
		{`class nonlocal:b`, func(t *test, tk Tokens) { // 11
			t.Err = Error{
				Err:     ErrMissingIdentifier,
				Parsing: "ClassDefinition",
				Token:   tk[2],
			}
		}},
		{`class a[nonlocal]`, func(t *test, tk Tokens) { // 12
			t.Err = Error{
				Err: Error{
					Err: Error{
						Err:     ErrMissingIdentifier,
						Parsing: "TypeParam",
						Token:   tk[4],
					},
					Parsing: "TypeParams",
					Token:   tk[4],
				},
				Parsing: "ClassDefinition",
				Token:   tk[3],
			}
		}},
		{`class a(nonlocal)`, func(t *test, tk Tokens) { // 13
			t.Err = Error{
				Err: Error{
					Err: Error{
						Err: Error{
							Err: Error{
								Err: wrapConditionalExpressionError(Error{
									Err:     ErrInvalidEnclosure,
									Parsing: "Enclosure",
									Token:   tk[4],
								}),
								Parsing: "Expression",
								Token:   tk[4],
							},
							Parsing: "AssignmentExpression",
							Token:   tk[4],
						},
						Parsing: "PositionalArgument",
						Token:   tk[4],
					},
					Parsing: "ArgumentList",
					Token:   tk[4],
				},
				Parsing: "ClassDefinition",
				Token:   tk[4],
			}
		}},
		{`class a(b)`, func(t *test, tk Tokens) { // 14
			t.Err = Error{
				Err:     ErrMissingColon,
				Parsing: "ClassDefinition",
				Token:   tk[6],
			}
		}},
		{`class a:nonlocal`, func(t *test, tk Tokens) { // 15
			t.Err = Error{
				Err: Error{
					Err: Error{
						Err: Error{
							Err: Error{
								Err:     ErrMissingIdentifier,
								Parsing: "NonLocalStatement",
								Token:   tk[5],
							},
							Parsing: "SimpleStatement",
							Token:   tk[4],
						},
						Parsing: "StatementList",
						Token:   tk[4],
					},
					Parsing: "Suite",
					Token:   tk[4],
				},
				Parsing: "ClassDefinition",
				Token:   tk[4],
			}
		}},
	}, func(t *test) (Type, error) {
		var c ClassDefinition

		err := c.parse(t.Tokens, t.Decorators)

		return c, err
	})
}

func TestSuite(t *testing.T) {
	doTests(t, []sourceFn{
		{`a`, func(t *test, tk Tokens) { // 1
			t.Output = Suite{
				StatementList: &StatementList{
					Statements: []SimpleStatement{
						{
							Type: StatementAssignment,
							AssignmentStatement: &AssignmentStatement{
								StarredExpression: &StarredExpression{
									OrExpr: WrapConditional(&Atom{
										Identifier: &tk[0],
										Tokens:     tk[:1],
									}).OrTest.AndTest.NotTest.Comparison.OrExpression,
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
		{"\n\ta", func(t *test, tk Tokens) { // 2
			t.Output = Suite{
				Statements: []Statement{
					{
						StatementList: &StatementList{
							Statements: []SimpleStatement{
								{
									Type: StatementAssignment,
									AssignmentStatement: &AssignmentStatement{
										StarredExpression: &StarredExpression{
											OrExpr: WrapConditional(&Atom{
												Identifier: &tk[2],
												Tokens:     tk[2:3],
											}).OrTest.AndTest.NotTest.Comparison.OrExpression,
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
				Tokens: tk[:4],
			}
		}},
		{"\n\ta\n\tb", func(t *test, tk Tokens) { // 3
			t.Output = Suite{
				Statements: []Statement{
					{
						StatementList: &StatementList{
							Statements: []SimpleStatement{
								{
									Type: StatementAssignment,
									AssignmentStatement: &AssignmentStatement{
										StarredExpression: &StarredExpression{
											OrExpr: WrapConditional(&Atom{
												Identifier: &tk[2],
												Tokens:     tk[2:3],
											}).OrTest.AndTest.NotTest.Comparison.OrExpression,
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
					{
						StatementList: &StatementList{
							Statements: []SimpleStatement{
								{
									Type: StatementAssignment,
									AssignmentStatement: &AssignmentStatement{
										StarredExpression: &StarredExpression{
											OrExpr: WrapConditional(&Atom{
												Identifier: &tk[5],
												Tokens:     tk[5:6],
											}).OrTest.AndTest.NotTest.Comparison.OrExpression,
											Tokens: tk[5:6],
										},
										Tokens: tk[5:6],
									},
									Tokens: tk[5:6],
								},
							},
							Tokens: tk[5:6],
						},
						Tokens: tk[5:6],
					},
				},
				Tokens: tk[:7],
			}
		}},
		{"#A comment\n\ta", func(t *test, tk Tokens) { // 4
			t.Output = Suite{
				Statements: []Statement{
					{
						StatementList: &StatementList{
							Statements: []SimpleStatement{
								{
									Type: StatementAssignment,
									AssignmentStatement: &AssignmentStatement{
										StarredExpression: &StarredExpression{
											OrExpr: WrapConditional(&Atom{
												Identifier: &tk[3],
												Tokens:     tk[3:4],
											}).OrTest.AndTest.NotTest.Comparison.OrExpression,
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
				Tokens: tk[:5],
			}
		}},
		{"\n\ta\nb", func(t *test, tk Tokens) { // 5
			t.Output = Suite{
				Statements: []Statement{
					{
						StatementList: &StatementList{
							Statements: []SimpleStatement{
								{
									Type: StatementAssignment,
									AssignmentStatement: &AssignmentStatement{
										StarredExpression: &StarredExpression{
											OrExpr: WrapConditional(&Atom{
												Identifier: &tk[2],
												Tokens:     tk[2:3],
											}).OrTest.AndTest.NotTest.Comparison.OrExpression,
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
				Tokens: tk[:5],
			}
		}},
		{`nonlocal`, func(t *test, tk Tokens) { // 6
			t.Err = Error{
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
				Parsing: "Suite",
				Token:   tk[0],
			}
		}},
		{"\n", func(t *test, tk Tokens) { // 7
			t.Err = Error{
				Err:     ErrMissingIndent,
				Parsing: "Suite",
				Token:   tk[1],
			}
		}},
		{"\n\tnonlocal", func(t *test, tk Tokens) { // 8
			t.Err = Error{
				Err: Error{
					Err: Error{
						Err: Error{
							Err: Error{
								Err:     ErrMissingIdentifier,
								Parsing: "NonLocalStatement",
								Token:   tk[3],
							},
							Parsing: "SimpleStatement",
							Token:   tk[2],
						},
						Parsing: "StatementList",
						Token:   tk[2],
					},
					Parsing: "Statement",
					Token:   tk[2],
				},
				Parsing: "Suite",
				Token:   tk[2],
			}
		}},
		{"\n\t@a\n\t@b\n\tdef c():\n\t\td", func(t *test, tk Tokens) { // 9
			t.Output = Suite{
				Statements: []Statement{
					{
						CompoundStatement: &CompoundStatement{
							Func: &FuncDefinition{
								Decorators: &Decorators{
									Decorators: []AssignmentExpression{
										{
											Expression: Expression{
												ConditionalExpression: WrapConditional(&Atom{
													Identifier: &tk[3],
													Tokens:     tk[3:4],
												}),
												Tokens: tk[3:4],
											},
											Tokens: tk[3:4],
										},
										{
											Expression: Expression{
												ConditionalExpression: WrapConditional(&Atom{
													Identifier: &tk[7],
													Tokens:     tk[7:8],
												}),
												Tokens: tk[7:8],
											},
											Tokens: tk[7:8],
										},
									},
									Tokens: tk[2:9],
								},
								FuncName: &tk[12],
								ParameterList: ParameterList{
									Tokens: tk[13:13],
								},
								Suite: Suite{
									Statements: []Statement{
										{
											StatementList: &StatementList{
												Statements: []SimpleStatement{
													{
														Type: StatementAssignment,
														AssignmentStatement: &AssignmentStatement{
															StarredExpression: &StarredExpression{
																OrExpr: WrapConditional(&Atom{
																	Identifier: &tk[18],
																	Tokens:     tk[18:19],
																}).OrTest.AndTest.NotTest.Comparison.OrExpression,
																Tokens: tk[18:19],
															},
															Tokens: tk[18:19],
														},
														Tokens: tk[18:19],
													},
												},
												Tokens: tk[18:19],
											},
											Tokens: tk[18:19],
										},
									},
									Tokens: tk[16:20],
								},
								Tokens: tk[2:20],
							},
							Tokens: tk[2:20],
						},
						Tokens: tk[2:20],
					},
				},
				Tokens: tk[:21],
			}
		}},
	}, func(t *test) (Type, error) {
		var s Suite

		err := s.parse(t.Tokens)

		return s, err
	})
}

func TestTargetList(t *testing.T) {
	doTests(t, []sourceFn{
		{`a`, func(t *test, tk Tokens) { // 1
			t.Output = TargetList{
				Targets: []Target{
					{
						PrimaryExpression: &PrimaryExpression{
							Atom: &Atom{
								Identifier: &tk[0],
								Tokens:     tk[:1],
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
			t.Output = TargetList{
				Targets: []Target{
					{
						PrimaryExpression: &PrimaryExpression{
							Atom: &Atom{
								Identifier: &tk[0],
								Tokens:     tk[:1],
							},
							Tokens: tk[:1],
						},
						Tokens: tk[:1],
					},
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
				Tokens: tk[:3],
			}
		}},
		{`a , b`, func(t *test, tk Tokens) { // 3
			t.Output = TargetList{
				Targets: []Target{
					{
						PrimaryExpression: &PrimaryExpression{
							Atom: &Atom{
								Identifier: &tk[0],
								Tokens:     tk[:1],
							},
							Tokens: tk[:1],
						},
						Tokens: tk[:1],
					},
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
				Tokens: tk[:5],
			}
		}},
		{`a, ;`, func(t *test, tk Tokens) { // 4
			t.Output = TargetList{
				Targets: []Target{
					{
						PrimaryExpression: &PrimaryExpression{
							Atom: &Atom{
								Identifier: &tk[0],
								Tokens:     tk[:1],
							},
							Tokens: tk[:1],
						},
						Tokens: tk[:1],
					},
				},
				Tokens: tk[:1],
			}
		}},
		{`a, in`, func(t *test, tk Tokens) { // 5
			t.Output = TargetList{
				Targets: []Target{
					{
						PrimaryExpression: &PrimaryExpression{
							Atom: &Atom{
								Identifier: &tk[0],
								Tokens:     tk[:1],
							},
							Tokens: tk[:1],
						},
						Tokens: tk[:1],
					},
				},
				Tokens: tk[:1],
			}
		}},
		{`a, =`, func(t *test, tk Tokens) { // 6
			t.Output = TargetList{
				Targets: []Target{
					{
						PrimaryExpression: &PrimaryExpression{
							Atom: &Atom{
								Identifier: &tk[0],
								Tokens:     tk[:1],
							},
							Tokens: tk[:1],
						},
						Tokens: tk[:1],
					},
				},
				Tokens: tk[:1],
			}
		}},
		{`a,`, func(t *test, tk Tokens) { // 7
			t.Output = TargetList{
				Targets: []Target{
					{
						PrimaryExpression: &PrimaryExpression{
							Atom: &Atom{
								Identifier: &tk[0],
								Tokens:     tk[:1],
							},
							Tokens: tk[:1],
						},
						Tokens: tk[:1],
					},
				},
				Tokens: tk[:1],
			}
		}},
		{`nonlocal`, func(t *test, tk Tokens) { // 8
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
					Parsing: "Target",
					Token:   tk[0],
				},
				Parsing: "TargetList",
				Token:   tk[0],
			}
		}},
	}, func(t *test) (Type, error) {
		var tl TargetList

		err := tl.parse(t.Tokens)

		return tl, err
	})
}

func TestTarget(t *testing.T) {
	doTests(t, []sourceFn{
		{`a`, func(t *test, tk Tokens) { // 1
			t.Output = Target{
				PrimaryExpression: &PrimaryExpression{
					Atom: &Atom{
						Identifier: &tk[0],
						Tokens:     tk[:1],
					},
					Tokens: tk[:1],
				},
				Tokens: tk[:1],
			}
		}},
		{`(a)`, func(t *test, tk Tokens) { // 2
			t.Output = Target{
				Tuple: &TargetList{
					Targets: []Target{
						{
							PrimaryExpression: &PrimaryExpression{
								Atom: &Atom{
									Identifier: &tk[1],
									Tokens:     tk[1:2],
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
		{`( a )`, func(t *test, tk Tokens) { // 3
			t.Output = Target{
				Tuple: &TargetList{
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
				Tokens: tk[:5],
			}
		}},
		{`[a]`, func(t *test, tk Tokens) { // 4
			t.Output = Target{
				Array: &TargetList{
					Targets: []Target{
						{
							PrimaryExpression: &PrimaryExpression{
								Atom: &Atom{
									Identifier: &tk[1],
									Tokens:     tk[1:2],
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
		{`[ a ]`, func(t *test, tk Tokens) { // 5
			t.Output = Target{
				Array: &TargetList{
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
				Tokens: tk[:5],
			}
		}},
		{`a.b`, func(t *test, tk Tokens) { // 6
			t.Output = Target{
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
				Tokens: tk[:3],
			}
		}},
		{`a[b]`, func(t *test, tk Tokens) { // 7
			t.Output = Target{
				PrimaryExpression: &PrimaryExpression{
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
				},
				Tokens: tk[:4],
			}
		}},
		{`*a`, func(t *test, tk Tokens) { // 8
			t.Output = Target{
				Star: &Target{
					PrimaryExpression: &PrimaryExpression{
						Atom: &Atom{
							Identifier: &tk[1],
							Tokens:     tk[1:2],
						},
						Tokens: tk[1:2],
					},
					Tokens: tk[1:2],
				},
				Tokens: tk[:2],
			}
		}},
		{`* a`, func(t *test, tk Tokens) { // 9
			t.Output = Target{
				Star: &Target{
					PrimaryExpression: &PrimaryExpression{
						Atom: &Atom{
							Identifier: &tk[2],
							Tokens:     tk[2:3],
						},
						Tokens: tk[2:3],
					},
					Tokens: tk[2:3],
				},
				Tokens: tk[:3],
			}
		}},
		{`nonlocal`, func(t *test, tk Tokens) { // 10
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
				Parsing: "Target",
				Token:   tk[0],
			}
		}},
		{`(nonlocal)`, func(t *test, tk Tokens) { // 11
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
						Parsing: "Target",
						Token:   tk[1],
					},
					Parsing: "TargetList",
					Token:   tk[1],
				},
				Parsing: "Target",
				Token:   tk[1],
			}
		}},
		{`[nonlocal]`, func(t *test, tk Tokens) { // 12
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
						Parsing: "Target",
						Token:   tk[1],
					},
					Parsing: "TargetList",
					Token:   tk[1],
				},
				Parsing: "Target",
				Token:   tk[1],
			}
		}},
		{`*nonlocal`, func(t *test, tk Tokens) { // 13
			t.Err = Error{
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
					Parsing: "Target",
					Token:   tk[1],
				},
				Parsing: "Target",
				Token:   tk[1],
			}
		}},
		{`a()`, func(t *test, tk Tokens) { // 14
			t.Err = Error{
				Err:     ErrMissingIdentifier,
				Parsing: "Target",
				Token:   tk[0],
			}
		}},
		{`{a}`, func(t *test, tk Tokens) { // 15
			t.Err = Error{
				Err:     ErrMissingIdentifier,
				Parsing: "Target",
				Token:   tk[0],
			}
		}},
	}, func(t *test) (Type, error) {
		var tt Target

		err := tt.parse(t.Tokens)

		return tt, err
	})
}

func TestStarredList(t *testing.T) {
	doTests(t, []sourceFn{
		{`a`, func(t *test, tk Tokens) { // 1
			t.Output = StarredList{
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
			}
		}},
		{`a,b`, func(t *test, tk Tokens) { // 2
			t.Output = StarredList{
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
		{`a , b`, func(t *test, tk Tokens) { // 3
			t.Output = StarredList{
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
		{`a : b, c`, func(t *test, tk Tokens) { // 4
			t.Output = StarredList{
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
			}
		}},
		{`a for b in c`, func(t *test, tk Tokens) { // 5
			t.Output = StarredList{
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
			}
		}},
		{`nonlocal`, func(t *test, tk Tokens) { // 6
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
					Parsing: "StarredItem",
					Token:   tk[0],
				},
				Parsing: "StarredList",
				Token:   tk[0],
			}
		}},
		{`a b`, func(t *test, tk Tokens) { // 6
			t.Err = Error{
				Err:     ErrMissingComma,
				Parsing: "StarredList",
				Token:   tk[2],
			}
		}},
	}, func(t *test) (Type, error) {
		var s StarredList

		err := s.parse(t.Tokens)

		return s, err
	})
}

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

func TestParameterList(t *testing.T) {
	doTests(t, []sourceFn{
		{`a`, func(t *test, tk Tokens) { // 1
			t.Output = ParameterList{
				NoPosOnly: []DefParameter{
					{
						Parameter: Parameter{
							Identifier: &tk[0],
							Tokens:     tk[:1],
						},
						Tokens: tk[:1],
					},
				},
				Tokens: tk[:1],
			}
		}},
		{`a,b`, func(t *test, tk Tokens) { // 2
			t.Output = ParameterList{
				NoPosOnly: []DefParameter{
					{
						Parameter: Parameter{
							Identifier: &tk[0],
							Tokens:     tk[:1],
						},
						Tokens: tk[:1],
					},
					{
						Parameter: Parameter{
							Identifier: &tk[2],
							Tokens:     tk[2:3],
						},
						Tokens: tk[2:3],
					},
				},
				Tokens: tk[:3],
			}
		}},
		{`a , b`, func(t *test, tk Tokens) { // 3
			t.Output = ParameterList{
				NoPosOnly: []DefParameter{
					{
						Parameter: Parameter{
							Identifier: &tk[0],
							Tokens:     tk[:1],
						},
						Tokens: tk[:1],
					},
					{
						Parameter: Parameter{
							Identifier: &tk[4],
							Tokens:     tk[4:5],
						},
						Tokens: tk[4:5],
					},
				},
				Tokens: tk[:5],
			}
		}},
		{`*a`, func(t *test, tk Tokens) { // 4
			t.Output = ParameterList{
				StarArg: &Parameter{
					Identifier: &tk[1],
					Tokens:     tk[1:2],
				},
				Tokens: tk[:2],
			}
		}},
		{`a,*b`, func(t *test, tk Tokens) { // 5
			t.Output = ParameterList{
				NoPosOnly: []DefParameter{
					{
						Parameter: Parameter{
							Identifier: &tk[0],
							Tokens:     tk[:1],
						},
						Tokens: tk[:1],
					},
				},
				StarArg: &Parameter{
					Identifier: &tk[3],
					Tokens:     tk[3:4],
				},
				Tokens: tk[:4],
			}
		}},
		{`**a`, func(t *test, tk Tokens) { // 6
			t.Output = ParameterList{
				StarStarArg: &Parameter{
					Identifier: &tk[1],
					Tokens:     tk[1:2],
				},
				Tokens: tk[:2],
			}
		}},
		{`a, ** b`, func(t *test, tk Tokens) { // 7
			t.Output = ParameterList{
				NoPosOnly: []DefParameter{
					{
						Parameter: Parameter{
							Identifier: &tk[0],
							Tokens:     tk[:1],
						},
						Tokens: tk[:1],
					},
				},
				StarStarArg: &Parameter{
					Identifier: &tk[5],
					Tokens:     tk[5:6],
				},
				Tokens: tk[:6],
			}
		}},
		{`a, ** b, c`, func(t *test, tk Tokens) { // 8
			t.Output = ParameterList{
				NoPosOnly: []DefParameter{
					{
						Parameter: Parameter{
							Identifier: &tk[0],
							Tokens:     tk[:1],
						},
						Tokens: tk[:1],
					},
				},
				StarStarArg: &Parameter{
					Identifier: &tk[5],
					Tokens:     tk[5:6],
				},
				Tokens: tk[:7],
			}
		}},
		{`**a`, func(t *test, tk Tokens) { // 9
			t.Output = ParameterList{
				StarStarArg: &Parameter{
					Identifier: &tk[1],
					Tokens:     tk[1:2],
				},
				Tokens: tk[:2],
			}
		}},
		{`*a, **b`, func(t *test, tk Tokens) { // 10
			t.Output = ParameterList{
				StarArg: &Parameter{
					Identifier: &tk[1],
					Tokens:     tk[1:2],
				},
				StarStarArg: &Parameter{
					Identifier: &tk[5],
					Tokens:     tk[5:6],
				},
				Tokens: tk[:6],
			}
		}},
		{`a, *b, **c`, func(t *test, tk Tokens) { // 11
			t.Output = ParameterList{
				NoPosOnly: []DefParameter{
					{
						Parameter: Parameter{
							Identifier: &tk[0],
							Tokens:     tk[:1],
						},
						Tokens: tk[:1],
					},
				},
				StarArg: &Parameter{
					Identifier: &tk[4],
					Tokens:     tk[4:5],
				},
				StarStarArg: &Parameter{
					Identifier: &tk[8],
					Tokens:     tk[8:9],
				},
				Tokens: tk[:9],
			}
		}},
		{`a, b, *c, d, e, **f`, func(t *test, tk Tokens) { // 12
			t.Output = ParameterList{
				NoPosOnly: []DefParameter{
					{
						Parameter: Parameter{
							Identifier: &tk[0],
							Tokens:     tk[:1],
						},
						Tokens: tk[:1],
					},
					{
						Parameter: Parameter{
							Identifier: &tk[3],
							Tokens:     tk[3:4],
						},
						Tokens: tk[3:4],
					},
				},
				StarArg: &Parameter{
					Identifier: &tk[7],
					Tokens:     tk[7:8],
				},
				StarArgs: []DefParameter{
					{
						Parameter: Parameter{
							Identifier: &tk[10],
							Tokens:     tk[10:11],
						},
						Tokens: tk[10:11],
					},
					{
						Parameter: Parameter{
							Identifier: &tk[13],
							Tokens:     tk[13:14],
						},
						Tokens: tk[13:14],
					},
				},
				StarStarArg: &Parameter{
					Identifier: &tk[17],
					Tokens:     tk[17:18],
				},
				Tokens: tk[:18],
			}
		}},
		{`a, /`, func(t *test, tk Tokens) { // 13
			t.Output = ParameterList{
				DefParameters: []DefParameter{
					{
						Parameter: Parameter{
							Identifier: &tk[0],
							Tokens:     tk[:1],
						},
						Tokens: tk[:1],
					},
				},
				Tokens: tk[:4],
			}
		}},
		{`a, b, /`, func(t *test, tk Tokens) { // 14
			t.Output = ParameterList{
				DefParameters: []DefParameter{
					{
						Parameter: Parameter{
							Identifier: &tk[0],
							Tokens:     tk[:1],
						},
						Tokens: tk[:1],
					},
					{
						Parameter: Parameter{
							Identifier: &tk[3],
							Tokens:     tk[3:4],
						},
						Tokens: tk[3:4],
					},
				},
				Tokens: tk[:7],
			}
		}},
		{`a, b, /, c`, func(t *test, tk Tokens) { // 15
			t.Output = ParameterList{
				DefParameters: []DefParameter{
					{
						Parameter: Parameter{
							Identifier: &tk[0],
							Tokens:     tk[:1],
						},
						Tokens: tk[:1],
					},
					{
						Parameter: Parameter{
							Identifier: &tk[3],
							Tokens:     tk[3:4],
						},
						Tokens: tk[3:4],
					},
				},
				NoPosOnly: []DefParameter{
					{
						Parameter: Parameter{
							Identifier: &tk[9],
							Tokens:     tk[9:10],
						},
						Tokens: tk[9:10],
					},
				},
				Tokens: tk[:10],
			}
		}},
		{`a, /, *b`, func(t *test, tk Tokens) { // 16
			t.Output = ParameterList{
				DefParameters: []DefParameter{
					{
						Parameter: Parameter{
							Identifier: &tk[0],
							Tokens:     tk[:1],
						},
						Tokens: tk[:1],
					},
				},
				StarArg: &Parameter{
					Identifier: &tk[7],
					Tokens:     tk[7:8],
				},
				Tokens: tk[:8],
			}
		}},
		{`a, /, **b`, func(t *test, tk Tokens) { // 17
			t.Output = ParameterList{
				DefParameters: []DefParameter{
					{
						Parameter: Parameter{
							Identifier: &tk[0],
							Tokens:     tk[:1],
						},
						Tokens: tk[:1],
					},
				},
				StarStarArg: &Parameter{
					Identifier: &tk[7],
					Tokens:     tk[7:8],
				},
				Tokens: tk[:8],
			}
		}},
		{`a, /, *b, **c`, func(t *test, tk Tokens) { // 18
			t.Output = ParameterList{
				DefParameters: []DefParameter{
					{
						Parameter: Parameter{
							Identifier: &tk[0],
							Tokens:     tk[:1],
						},
						Tokens: tk[:1],
					},
				},
				StarArg: &Parameter{
					Identifier: &tk[7],
					Tokens:     tk[7:8],
				},
				StarStarArg: &Parameter{
					Identifier: &tk[11],
					Tokens:     tk[11:12],
				},
				Tokens: tk[:12],
			}
		}},
		{`a, b, /, c, d, *e, f, g, **h`, func(t *test, tk Tokens) { // 19
			t.Output = ParameterList{
				DefParameters: []DefParameter{
					{
						Parameter: Parameter{
							Identifier: &tk[0],
							Tokens:     tk[:1],
						},
						Tokens: tk[:1],
					},
					{
						Parameter: Parameter{
							Identifier: &tk[3],
							Tokens:     tk[3:4],
						},
						Tokens: tk[3:4],
					},
				},
				NoPosOnly: []DefParameter{
					{
						Parameter: Parameter{
							Identifier: &tk[9],
							Tokens:     tk[9:10],
						},
						Tokens: tk[9:10],
					},
					{
						Parameter: Parameter{
							Identifier: &tk[12],
							Tokens:     tk[12:13],
						},
						Tokens: tk[12:13],
					},
				},
				StarArg: &Parameter{
					Identifier: &tk[16],
					Tokens:     tk[16:17],
				},
				StarArgs: []DefParameter{
					{
						Parameter: Parameter{
							Identifier: &tk[19],
							Tokens:     tk[19:20],
						},
						Tokens: tk[19:20],
					},
					{
						Parameter: Parameter{
							Identifier: &tk[22],
							Tokens:     tk[22:23],
						},
						Tokens: tk[22:23],
					},
				},
				StarStarArg: &Parameter{
					Identifier: &tk[26],
					Tokens:     tk[26:27],
				},
				Tokens: tk[:27],
			}
		}},
		{`a=nonlocal`, func(t *test, tk Tokens) { // 20
			t.Err = Error{
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
					Parsing: "DefParameter",
					Token:   tk[2],
				},
				Parsing: "ParameterList",
				Token:   tk[0],
			}
		}},
		{`*nonlocal`, func(t *test, tk Tokens) { // 21
			t.Err = Error{
				Err: Error{
					Err:     ErrMissingIdentifier,
					Parsing: "Parameter",
					Token:   tk[1],
				},
				Parsing: "ParameterList",
				Token:   tk[1],
			}
		}},
		{`**nonlocal`, func(t *test, tk Tokens) { // 22
			t.Err = Error{
				Err: Error{
					Err:     ErrMissingIdentifier,
					Parsing: "Parameter",
					Token:   tk[1],
				},
				Parsing: "ParameterList",
				Token:   tk[1],
			}
		}},
		{`a,*nonlocal`, func(t *test, tk Tokens) { // 23
			t.Err = Error{
				Err: Error{
					Err:     ErrMissingIdentifier,
					Parsing: "Parameter",
					Token:   tk[3],
				},
				Parsing: "ParameterList",
				Token:   tk[3],
			}
		}},
		{`a,**nonlocal`, func(t *test, tk Tokens) { // 24
			t.Err = Error{
				Err: Error{
					Err:     ErrMissingIdentifier,
					Parsing: "Parameter",
					Token:   tk[3],
				},
				Parsing: "ParameterList",
				Token:   tk[3],
			}
		}},
	}, func(t *test) (Type, error) {
		var p ParameterList

		err := p.parse(t.Tokens, t.AllowTypeAnnotations)

		return p, err
	})
}

func TestDefParameter(t *testing.T) {
	doTests(t, []sourceFn{
		{`a`, func(t *test, tk Tokens) { // 1
			t.Output = DefParameter{
				Parameter: Parameter{
					Identifier: &tk[0],
					Tokens:     tk[:1],
				},
				Tokens: tk[:1],
			}
		}},
		{`a=b`, func(t *test, tk Tokens) { // 2
			t.Output = DefParameter{
				Parameter: Parameter{
					Identifier: &tk[0],
					Tokens:     tk[:1],
				},
				Value: &Expression{
					ConditionalExpression: WrapConditional(&Atom{
						Identifier: &tk[2],
						Tokens:     tk[2:3],
					}),
					Tokens: tk[2:3],
				},
				Tokens: tk[:3],
			}
		}},
		{`a = b`, func(t *test, tk Tokens) { // 3
			t.Output = DefParameter{
				Parameter: Parameter{
					Identifier: &tk[0],
					Tokens:     tk[:1],
				},
				Value: &Expression{
					ConditionalExpression: WrapConditional(&Atom{
						Identifier: &tk[4],
						Tokens:     tk[4:5],
					}),
					Tokens: tk[4:5],
				},
				Tokens: tk[:5],
			}
		}},
		{`a: b = c`, func(t *test, tk Tokens) { // 4
			t.Output = DefParameter{
				Parameter: Parameter{
					Identifier: &tk[0],
					Tokens:     tk[:1],
				},
				Tokens: tk[:1],
			}
		}},
		{`a:b=c`, func(t *test, tk Tokens) { // 5
			t.AllowTypeAnnotations = true
			t.Output = DefParameter{
				Parameter: Parameter{
					Identifier: &tk[0],
					Type: &Expression{
						ConditionalExpression: WrapConditional(&Atom{
							Identifier: &tk[2],
							Tokens:     tk[2:3],
						}),
						Tokens: tk[2:3],
					},
					Tokens: tk[:3],
				},
				Value: &Expression{
					ConditionalExpression: WrapConditional(&Atom{
						Identifier: &tk[4],
						Tokens:     tk[4:5],
					}),
					Tokens: tk[4:5],
				},
				Tokens: tk[:5],
			}
		}},
		{`nonlocal`, func(t *test, tk Tokens) { // 6
			t.Err = Error{
				Err: Error{
					Err:     ErrMissingIdentifier,
					Parsing: "Parameter",
					Token:   tk[0],
				},
				Parsing: "DefParameter",
				Token:   tk[0],
			}
		}},
		{`a=nonlocal`, func(t *test, tk Tokens) { // 7
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
				Parsing: "DefParameter",
				Token:   tk[2],
			}
		}},
	}, func(t *test) (Type, error) {
		var d DefParameter

		err := d.parse(t.Tokens, t.AllowTypeAnnotations)

		return d, err
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
		{`nonlocal`, func(t *test, tk Tokens) { // 6
			t.Err = Error{
				Err:     ErrMissingIdentifier,
				Parsing: "Parameter",
				Token:   tk[0],
			}
		}},
		{`a:nonlocal`, func(t *test, tk Tokens) { // 7
			t.AllowTypeAnnotations = true
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
				Parsing: "Parameter",
				Token:   tk[2],
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

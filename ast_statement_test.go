package python

import "testing"

func TestStatement(t *testing.T) {
	doTests(t, []sourceFn{
		{`a`, func(t *test, tk Tokens) { // 1
			t.Output = Statement{
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
			}
		}},
		{`class a: pass`, func(t *test, tk Tokens) { // 2
			t.Output = Statement{
				CompoundStatement: &CompoundStatement{
					Class: &ClassDefinition{
						ClassName: &tk[2],
						Suite: Suite{
							StatementList: &StatementList{
								Statements: []SimpleStatement{
									{
										Type:   StatementPass,
										Tokens: tk[5:6],
									},
								},
								Tokens: tk[5:6],
							},
							Tokens: tk[5:6],
						},
						Tokens: tk[:6],
					},
					Tokens: tk[:6],
				},
				Tokens: tk[:6],
			}
		}},
		{"@decorator\nclass a: pass", func(t *test, tk Tokens) { // 3
			t.Output = Statement{
				CompoundStatement: &CompoundStatement{
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
						Suite: Suite{
							StatementList: &StatementList{
								Statements: []SimpleStatement{
									{
										Type:   StatementPass,
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
				},
				Tokens: tk[:9],
			}
		}},
		{`nonlocal`, func(t *test, tk Tokens) { // 4
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
				Parsing: "Statement",
				Token:   tk[0],
			}
		}},
		{`class nonlocal: a`, func(t *test, tk Tokens) { // 5
			t.Err = Error{
				Err: Error{
					Err: Error{
						Err:     ErrMissingIdentifier,
						Parsing: "ClassDefinition",
						Token:   tk[2],
					},
					Parsing: "CompoundStatement",
					Token:   tk[0],
				},
				Parsing: "Statement",
				Token:   tk[0],
			}
		}},
	}, func(t *test) (Type, error) {
		var s Statement

		err := s.parse(t.Tokens)

		return s, err
	})
}

func TestStatementList(t *testing.T) {
	doTests(t, []sourceFn{
		{`a`, func(t *test, tk Tokens) { // 1
			t.Output = StatementList{
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
			}
		}},
		{`a;`, func(t *test, tk Tokens) { // 2
			t.Output = StatementList{
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
				Tokens: tk[:2],
			}
		}},
		{`a ; `, func(t *test, tk Tokens) { // 3
			t.Output = StatementList{
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
				Tokens: tk[:3],
			}
		}},
		{`a #Comment`, func(t *test, tk Tokens) { // 4
			t.Output = StatementList{
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
			}
		}},
		{"a\n", func(t *test, tk Tokens) { // 5
			t.Output = StatementList{
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
			}
		}},
		{`a; #Comment`, func(t *test, tk Tokens) { // 6
			t.Output = StatementList{
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
				Tokens: tk[:2],
			}
		}},
		{"a;\n", func(t *test, tk Tokens) { // 7
			t.Output = StatementList{
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
				Tokens: tk[:2],
			}
		}},
		{`a;b`, func(t *test, tk Tokens) { // 8
			t.Output = StatementList{
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
				Tokens: tk[:3],
			}
		}},
		{`a ; b `, func(t *test, tk Tokens) { // 9
			t.Output = StatementList{
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
				Tokens: tk[:5],
			}
		}},
		{`nonlocal`, func(t *test, tk Tokens) { // 10
			t.Err = Error{
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
			}
		}},
		{`a;nonlocal`, func(t *test, tk Tokens) { // 11
			t.Err = Error{
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
			}
		}},
	}, func(t *test) (Type, error) {
		var s StatementList

		err := s.parse(t.Tokens)

		return s, err
	})
}

func TestSimpleStatement(t *testing.T) {
	doTests(t, []sourceFn{
		{`assert a`, func(t *test, tk Tokens) { // 1
			t.Output = SimpleStatement{
				Type: StatementAssert,
				AssertStatement: &AssertStatement{
					Expressions: []Expression{
						{
							ConditionalExpression: WrapConditional(&Atom{
								Identifier: &tk[2],
								Tokens:     tk[2:3],
							}),
							Tokens: tk[2:3],
						},
					},
					Tokens: tk[:3],
				},
				Tokens: tk[:3],
			}
		}},
		{`pass`, func(t *test, tk Tokens) { // 2
			t.Output = SimpleStatement{
				Type:   StatementPass,
				Tokens: tk[:1],
			}
		}},
		{`del a`, func(t *test, tk Tokens) { // 3
			t.Output = SimpleStatement{
				Type: StatementDel,
				DelStatement: &DelStatement{
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
					Tokens: tk[:3],
				},
				Tokens: tk[:3],
			}
		}},
		{`return a`, func(t *test, tk Tokens) { // 4
			t.Output = SimpleStatement{
				Type: StatementReturn,
				ReturnStatement: &ReturnStatement{
					Expression: &Expression{
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
		{`yield a`, func(t *test, tk Tokens) { // 5
			t.Output = SimpleStatement{
				Type: StatementYield,
				YieldStatement: &YieldExpression{
					ExpressionList: &ExpressionList{
						Expressions: []Expression{
							{
								ConditionalExpression: WrapConditional(&Atom{
									Identifier: &tk[2],
									Tokens:     tk[2:3],
								}),
								Tokens: tk[2:3],
							},
						},
						Tokens: tk[2:3],
					},
					Tokens: tk[:3],
				},
				Tokens: tk[:3],
			}
		}},
		{`raise a`, func(t *test, tk Tokens) { // 6
			t.Output = SimpleStatement{
				Type: StatementRaise,
				RaiseStatement: &RaiseStatement{
					Expression: &Expression{
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
		{`break`, func(t *test, tk Tokens) { // 7
			t.Output = SimpleStatement{
				Type:   StatementBreak,
				Tokens: tk[:1],
			}
		}},
		{`continue`, func(t *test, tk Tokens) { // 8
			t.Output = SimpleStatement{
				Type:   StatementContinue,
				Tokens: tk[:1],
			}
		}},
		{`import a`, func(t *test, tk Tokens) { // 9
			t.Output = SimpleStatement{
				Type: StatementImport,
				ImportStatement: &ImportStatement{
					Modules: []ModuleAs{
						{
							Module: Module{
								Identifiers: []*Token{
									&tk[2],
								},
								Tokens: tk[2:3],
							},
							Tokens: tk[2:3],
						},
					},
					Tokens: tk[:3],
				},
				Tokens: tk[:3],
			}
		}},
		{`from a import b`, func(t *test, tk Tokens) { // 10
			t.Output = SimpleStatement{
				Type: StatementImport,
				ImportStatement: &ImportStatement{
					RelativeModule: &RelativeModule{
						Module: &Module{
							Identifiers: []*Token{
								&tk[2],
							},
							Tokens: tk[2:3],
						},
						Tokens: tk[2:3],
					},
					Modules: []ModuleAs{
						{
							Module: Module{
								Identifiers: []*Token{
									&tk[6],
								},
								Tokens: tk[6:7],
							},
							Tokens: tk[6:7],
						},
					},
					Tokens: tk[:7],
				},
				Tokens: tk[:7],
			}
		}},
		{`global a`, func(t *test, tk Tokens) { // 11
			t.Output = SimpleStatement{
				Type: StatementGlobal,
				GlobalStatement: &GlobalStatement{
					Identifiers: []*Token{
						&tk[2],
					},
					Tokens: tk[:3],
				},
				Tokens: tk[:3],
			}
		}},
		{`nonlocal a`, func(t *test, tk Tokens) { // 12
			t.Output = SimpleStatement{
				Type: StatementNonLocal,
				NonLocalStatement: &NonLocalStatement{
					Identifiers: []*Token{
						&tk[2],
					},
					Tokens: tk[:3],
				},
				Tokens: tk[:3],
			}
		}},
		{`type a=b`, func(t *test, tk Tokens) { // 13
			t.Output = SimpleStatement{
				Type: StatementTyp,
				TypeStatement: &TypeStatement{
					Identifier: &tk[2],
					Expression: Expression{
						ConditionalExpression: WrapConditional(&Atom{
							Identifier: &tk[4],
							Tokens:     tk[4:5],
						}),
						Tokens: tk[4:5],
					},
					Tokens: tk[:5],
				},
				Tokens: tk[:5],
			}
		}},
		{`a`, func(t *test, tk Tokens) { // 14
			t.Output = SimpleStatement{
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
			}
		}},
		{`a:b`, func(t *test, tk Tokens) { // 15
			t.Output = SimpleStatement{
				Type: StatementAnnotatedAssignment,
				AnnotatedAssignmentStatement: &AnnotatedAssignmentStatement{
					AugTarget: AugTarget{
						PrimaryExpression: PrimaryExpression{
							Atom: &Atom{
								Identifier: &tk[0],
								Tokens:     tk[:1],
							},
							Tokens: tk[:1],
						},
						Tokens: tk[:1],
					},
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
		{`a+=b`, func(t *test, tk Tokens) { // 16
			t.Output = SimpleStatement{
				Type: StatementAugmentedAssignment,
				AugmentedAssignmentStatement: &AugmentedAssignmentStatement{
					AugTarget: AugTarget{
						PrimaryExpression: PrimaryExpression{
							Atom: &Atom{
								Identifier: &tk[0],
								Tokens:     tk[:1],
							},
							Tokens: tk[:1],
						},
						Tokens: tk[:1],
					},
					AugOp: &tk[1],
					ExpressionList: &ExpressionList{
						Expressions: []Expression{
							{
								ConditionalExpression: WrapConditional(&Atom{
									Identifier: &tk[2],
									Tokens:     tk[2:3],
								}),
								Tokens: tk[2:3],
							},
						},
						Tokens: tk[2:3],
					},
					Tokens: tk[:3],
				},
				Tokens: tk[:3],
			}
		}},
		{`assert nonlocal`, func(t *test, tk Tokens) { // 17
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
					Parsing: "AssertStatement",
					Token:   tk[2],
				},
				Parsing: "SimpleStatement",
				Token:   tk[0],
			}
		}},
		{`del nonlocal`, func(t *test, tk Tokens) { // 18
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
					Parsing: "DelStatement",
					Token:   tk[2],
				},
				Parsing: "SimpleStatement",
				Token:   tk[0],
			}
		}},
		{`return nonlocal`, func(t *test, tk Tokens) { // 19
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
					Parsing: "ReturnStatement",
					Token:   tk[2],
				},
				Parsing: "SimpleStatement",
				Token:   tk[0],
			}
		}},
		{`raise nonlocal from a`, func(t *test, tk Tokens) { // 20
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
					Parsing: "RaiseStatement",
					Token:   tk[2],
				},
				Parsing: "SimpleStatement",
				Token:   tk[0],
			}
		}},
		{`import nonlocal`, func(t *test, tk Tokens) { // 21
			t.Err = Error{
				Err: Error{
					Err: Error{
						Err: Error{
							Err:     ErrMissingIdentifier,
							Parsing: "Module",
							Token:   tk[2],
						},
						Parsing: "ModuleAs",
						Token:   tk[2],
					},
					Parsing: "ImportStatement",
					Token:   tk[2],
				},
				Parsing: "SimpleStatement",
				Token:   tk[0],
			}
		}},
		{`global nonlocal`, func(t *test, tk Tokens) { // 22
			t.Err = Error{
				Err: Error{
					Err:     ErrMissingIdentifier,
					Parsing: "GlobalStatement",
					Token:   tk[2],
				},
				Parsing: "SimpleStatement",
				Token:   tk[0],
			}
		}},
		{`nonlocal nonlocal`, func(t *test, tk Tokens) { // 23
			t.Err = Error{
				Err: Error{
					Err:     ErrMissingIdentifier,
					Parsing: "NonLocalStatement",
					Token:   tk[2],
				},
				Parsing: "SimpleStatement",
				Token:   tk[0],
			}
		}},
		{`type nonlocal[a] = b`, func(t *test, tk Tokens) { // 24
			t.Err = Error{
				Err: Error{
					Err:     ErrMissingIdentifier,
					Parsing: "TypeStatement",
					Token:   tk[2],
				},
				Parsing: "SimpleStatement",
				Token:   tk[0],
			}
		}},
		{`a=yield nonlocal`, func(t *test, tk Tokens) { // 25
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
							Parsing: "ExpressionList",
							Token:   tk[4],
						},
						Parsing: "YieldExpression",
						Token:   tk[4],
					},
					Parsing: "AssignmentStatement",
					Token:   tk[2],
				},
				Parsing: "SimpleStatement",
				Token:   tk[0],
			}
		}},
		{`a:nonlocal`, func(t *test, tk Tokens) { // 26
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
					Parsing: "AnnotatedAssignmentStatement",
					Token:   tk[2],
				},
				Parsing: "SimpleStatement",
				Token:   tk[0],
			}
		}},
		{`a/=nonlocal`, func(t *test, tk Tokens) { // 27
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
						Parsing: "ExpressionList",
						Token:   tk[2],
					},
					Parsing: "AugmentedAssignmentStatement",
					Token:   tk[2],
				},
				Parsing: "SimpleStatement",
				Token:   tk[0],
			}
		}},
	}, func(t *test) (Type, error) {
		var s SimpleStatement

		err := s.parse(t.Tokens)

		return s, err
	})
}

func TestAssertStatement(t *testing.T) {
	doTests(t, []sourceFn{
		{`assert a`, func(t *test, tk Tokens) { // 1
			t.Output = AssertStatement{
				Expressions: []Expression{
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
		{`assert a,b`, func(t *test, tk Tokens) { // 2
			t.Output = AssertStatement{
				Expressions: []Expression{
					{
						ConditionalExpression: WrapConditional(&Atom{
							Identifier: &tk[2],
							Tokens:     tk[2:3],
						}),
						Tokens: tk[2:3],
					},
					{
						ConditionalExpression: WrapConditional(&Atom{
							Identifier: &tk[4],
							Tokens:     tk[4:5],
						}),
						Tokens: tk[4:5],
					},
				},
				Tokens: tk[:5],
			}
		}},
		{`assert a , b`, func(t *test, tk Tokens) { // 3
			t.Output = AssertStatement{
				Expressions: []Expression{
					{
						ConditionalExpression: WrapConditional(&Atom{
							Identifier: &tk[2],
							Tokens:     tk[2:3],
						}),
						Tokens: tk[2:3],
					},
					{
						ConditionalExpression: WrapConditional(&Atom{
							Identifier: &tk[6],
							Tokens:     tk[6:7],
						}),
						Tokens: tk[6:7],
					},
				},
				Tokens: tk[:7],
			}
		}},
		{`assert nonlocal`, func(t *test, tk Tokens) { // 4
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
				Parsing: "AssertStatement",
				Token:   tk[2],
			}
		}},
	}, func(t *test) (Type, error) {
		var a AssertStatement

		err := a.parse(t.Tokens)

		return a, err
	})
}

func TestAssignmentStatement(t *testing.T) {
	doTests(t, []sourceFn{
		{`a`, func(t *test, tk Tokens) { // 1
			t.Output = AssignmentStatement{
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
			}
		}},
		{`yield a`, func(t *test, tk Tokens) { // 2
			t.Output = AssignmentStatement{
				YieldExpression: &YieldExpression{
					ExpressionList: &ExpressionList{
						Expressions: []Expression{
							{
								ConditionalExpression: WrapConditional(&Atom{
									Identifier: &tk[2],
									Tokens:     tk[2:3],
								}),
								Tokens: tk[2:3],
							},
						},
						Tokens: tk[2:3],
					},
					Tokens: tk[:3],
				},
				Tokens: tk[:3],
			}
		}},
		{`a=b`, func(t *test, tk Tokens) { // 3
			t.Output = AssignmentStatement{
				TargetLists: []TargetList{
					{
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
					},
				},
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
				Tokens: tk[:3],
			}
		}},
		{`a,=b`, func(t *test, tk Tokens) { // 4
			t.Output = AssignmentStatement{
				TargetLists: []TargetList{
					{
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
					},
				},
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
				Tokens: tk[:4],
			}
		}},
		{`a = b`, func(t *test, tk Tokens) { // 5
			t.Output = AssignmentStatement{
				TargetLists: []TargetList{
					{
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
					},
				},
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
				Tokens: tk[:5],
			}
		}},
		{`a=b=c`, func(t *test, tk Tokens) { // 6
			t.Output = AssignmentStatement{
				TargetLists: []TargetList{
					{
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
					},
					{
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
				},
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
				Tokens: tk[:5],
			}
		}},
		{`nonlocal`, func(t *test, tk Tokens) { // 7
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
					Parsing: "StarredExpression",
					Token:   tk[0],
				},
				Parsing: "AssignmentStatement",
				Token:   tk[0],
			}
		}},
		{`nonlocal=a`, func(t *test, tk Tokens) { // 8
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
						Parsing: "Target",
						Token:   tk[0],
					},
					Parsing: "TargetList",
					Token:   tk[0],
				},
				Parsing: "AssignmentStatement",
				Token:   tk[0],
			}
		}},
		{`a=yield nonlocal`, func(t *test, tk Tokens) { // 9
			t.Err = Error{
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
						Parsing: "ExpressionList",
						Token:   tk[4],
					},
					Parsing: "YieldExpression",
					Token:   tk[4],
				},
				Parsing: "AssignmentStatement",
				Token:   tk[2],
			}
		}},
	}, func(t *test) (Type, error) {
		var a AssignmentStatement

		err := a.parse(t.Tokens)

		return a, err
	})
}

func TestAugmentedAssignmentStatement(t *testing.T) {
	doTests(t, []sourceFn{
		{`a+=b`, func(t *test, tk Tokens) { // 1
			t.Output = AugmentedAssignmentStatement{
				AugTarget: AugTarget{
					PrimaryExpression: PrimaryExpression{
						Atom: &Atom{
							Identifier: &tk[0],
							Tokens:     tk[:1],
						},
						Tokens: tk[:1],
					},
					Tokens: tk[:1],
				},
				AugOp: &tk[1],
				ExpressionList: &ExpressionList{
					Expressions: []Expression{
						{
							ConditionalExpression: WrapConditional(&Atom{
								Identifier: &tk[2],
								Tokens:     tk[2:3],
							}),
							Tokens: tk[2:3],
						},
					},
					Tokens: tk[2:3],
				},
				Tokens: tk[:3],
			}
		}},
		{`a -= b`, func(t *test, tk Tokens) { // 2
			t.Output = AugmentedAssignmentStatement{
				AugTarget: AugTarget{
					PrimaryExpression: PrimaryExpression{
						Atom: &Atom{
							Identifier: &tk[0],
							Tokens:     tk[:1],
						},
						Tokens: tk[:1],
					},
					Tokens: tk[:1],
				},
				AugOp: &tk[2],
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
				Tokens: tk[:5],
			}
		}},
		{`a*=yield b`, func(t *test, tk Tokens) { // 3
			t.Output = AugmentedAssignmentStatement{
				AugTarget: AugTarget{
					PrimaryExpression: PrimaryExpression{
						Atom: &Atom{
							Identifier: &tk[0],
							Tokens:     tk[:1],
						},
						Tokens: tk[:1],
					},
					Tokens: tk[:1],
				},
				AugOp: &tk[1],
				YieldExpression: &YieldExpression{
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
				Tokens: tk[:5],
			}
		}},
		{`a @= yield b`, func(t *test, tk Tokens) { // 4
			t.Output = AugmentedAssignmentStatement{
				AugTarget: AugTarget{
					PrimaryExpression: PrimaryExpression{
						Atom: &Atom{
							Identifier: &tk[0],
							Tokens:     tk[:1],
						},
						Tokens: tk[:1],
					},
					Tokens: tk[:1],
				},
				AugOp: &tk[2],
				YieldExpression: &YieldExpression{
					ExpressionList: &ExpressionList{
						Expressions: []Expression{
							{
								ConditionalExpression: WrapConditional(&Atom{
									Identifier: &tk[6],
									Tokens:     tk[6:7],
								}),
								Tokens: tk[6:7],
							},
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
					Parsing: "AugTarget",
					Token:   tk[0],
				},
				Parsing: "AugmentedAssignmentStatement",
				Token:   tk[0],
			}
		}},
		{`a==b`, func(t *test, tk Tokens) { // 6
			t.Err = Error{
				Err:     ErrMissingOp,
				Parsing: "AugmentedAssignmentStatement",
				Token:   tk[1],
			}
		}},
		{`a/=nonlocal`, func(t *test, tk Tokens) { // 7
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
					Parsing: "ExpressionList",
					Token:   tk[2],
				},
				Parsing: "AugmentedAssignmentStatement",
				Token:   tk[2],
			}
		}},
		{`a/=yield nonlocal`, func(t *test, tk Tokens) { // 8
			t.Err = Error{
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
						Parsing: "ExpressionList",
						Token:   tk[4],
					},
					Parsing: "YieldExpression",
					Token:   tk[4],
				},
				Parsing: "AugmentedAssignmentStatement",
				Token:   tk[2],
			}
		}},
	}, func(t *test) (Type, error) {
		var a AugmentedAssignmentStatement

		err := a.parse(t.Tokens)

		return a, err
	})
}

func TestAugTarget(t *testing.T) {
	doTests(t, []sourceFn{
		{`a`, func(t *test, tk Tokens) { // 1
			t.Output = AugTarget{
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
		{`a[b]`, func(t *test, tk Tokens) { // 2
			t.Output = AugTarget{
				PrimaryExpression: PrimaryExpression{
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
		{`nonlocal`, func(t *test, tk Tokens) { // 3
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
				Parsing: "AugTarget",
				Token:   tk[0],
			}
		}},
		{`a()`, func(t *test, tk Tokens) { // 4
			t.Err = Error{
				Err:     ErrMissingIdentifier,
				Parsing: "AugTarget",
				Token:   tk[0],
			}
		}},
		{`(a)`, func(t *test, tk Tokens) { // 5
			t.Err = Error{
				Err:     ErrMissingIdentifier,
				Parsing: "AugTarget",
				Token:   tk[0],
			}
		}},
	}, func(t *test) (Type, error) {
		var a AugTarget

		err := a.parse(t.Tokens)

		return a, err
	})
}

func TestAnnotatedAssignmentStatement(t *testing.T) {
	doTests(t, []sourceFn{
		{`a:b`, func(t *test, tk Tokens) { // 1
			t.Output = AnnotatedAssignmentStatement{
				AugTarget: AugTarget{
					PrimaryExpression: PrimaryExpression{
						Atom: &Atom{
							Identifier: &tk[0],
							Tokens:     tk[:1],
						},
						Tokens: tk[:1],
					},
					Tokens: tk[:1],
				},
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
		{`a : b`, func(t *test, tk Tokens) { // 2
			t.Output = AnnotatedAssignmentStatement{
				AugTarget: AugTarget{
					PrimaryExpression: PrimaryExpression{
						Atom: &Atom{
							Identifier: &tk[0],
							Tokens:     tk[:1],
						},
						Tokens: tk[:1],
					},
					Tokens: tk[:1],
				},
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
		{`a:b=c`, func(t *test, tk Tokens) { // 3
			t.Output = AnnotatedAssignmentStatement{
				AugTarget: AugTarget{
					PrimaryExpression: PrimaryExpression{
						Atom: &Atom{
							Identifier: &tk[0],
							Tokens:     tk[:1],
						},
						Tokens: tk[:1],
					},
					Tokens: tk[:1],
				},
				Expression: Expression{
					ConditionalExpression: WrapConditional(&Atom{
						Identifier: &tk[2],
						Tokens:     tk[2:3],
					}),
					Tokens: tk[2:3],
				},
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
				Tokens: tk[:5],
			}
		}},
		{`a:b = c`, func(t *test, tk Tokens) { // 4
			t.Output = AnnotatedAssignmentStatement{
				AugTarget: AugTarget{
					PrimaryExpression: PrimaryExpression{
						Atom: &Atom{
							Identifier: &tk[0],
							Tokens:     tk[:1],
						},
						Tokens: tk[:1],
					},
					Tokens: tk[:1],
				},
				Expression: Expression{
					ConditionalExpression: WrapConditional(&Atom{
						Identifier: &tk[2],
						Tokens:     tk[2:3],
					}),
					Tokens: tk[2:3],
				},
				StarredExpression: &StarredExpression{
					Expression: &Expression{
						ConditionalExpression: WrapConditional(&Atom{
							Identifier: &tk[6],
							Tokens:     tk[6:7],
						}),
						Tokens: tk[6:7],
					},
					Tokens: tk[6:7],
				},
				Tokens: tk[:7],
			}
		}},
		{`a:b=yield c`, func(t *test, tk Tokens) { // 5
			t.Output = AnnotatedAssignmentStatement{
				AugTarget: AugTarget{
					PrimaryExpression: PrimaryExpression{
						Atom: &Atom{
							Identifier: &tk[0],
							Tokens:     tk[:1],
						},
						Tokens: tk[:1],
					},
					Tokens: tk[:1],
				},
				Expression: Expression{
					ConditionalExpression: WrapConditional(&Atom{
						Identifier: &tk[2],
						Tokens:     tk[2:3],
					}),
					Tokens: tk[2:3],
				},
				YieldExpression: &YieldExpression{
					ExpressionList: &ExpressionList{
						Expressions: []Expression{
							{
								ConditionalExpression: WrapConditional(&Atom{
									Identifier: &tk[6],
									Tokens:     tk[6:7],
								}),
								Tokens: tk[6:7],
							},
						},
						Tokens: tk[6:7],
					},
					Tokens: tk[4:7],
				},
				Tokens: tk[:7],
			}
		}},
		{`nonlocal`, func(t *test, tk Tokens) { // 6
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
					Parsing: "AugTarget",
					Token:   tk[0],
				},
				Parsing: "AnnotatedAssignmentStatement",
				Token:   tk[0],
			}
		}},
		{`a:nonlocal`, func(t *test, tk Tokens) { // 7
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
				Parsing: "AnnotatedAssignmentStatement",
				Token:   tk[2],
			}
		}},
		{`a:b=nonlocal`, func(t *test, tk Tokens) { // 8
			t.Err = Error{
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
					Parsing: "StarredExpression",
					Token:   tk[4],
				},
				Parsing: "AnnotatedAssignmentStatement",
				Token:   tk[4],
			}
		}},
		{`a:b=yield nonlocal`, func(t *test, tk Tokens) { // 9
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
						Parsing: "ExpressionList",
						Token:   tk[6],
					},
					Parsing: "YieldExpression",
					Token:   tk[6],
				},
				Parsing: "AnnotatedAssignmentStatement",
				Token:   tk[4],
			}
		}},
		{`a`, func(t *test, tk Tokens) { // 10
			t.Err = Error{
				Err:     ErrMissingColon,
				Parsing: "AnnotatedAssignmentStatement",
				Token:   tk[1],
			}
		}},
	}, func(t *test) (Type, error) {
		var a AnnotatedAssignmentStatement

		err := a.parse(t.Tokens)

		return a, err
	})
}

func TestStarredExpression(t *testing.T) {
	doTests(t, []sourceFn{
		{`a`, func(t *test, tk Tokens) { // 1
			t.Output = StarredExpression{
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
		{`*a`, func(t *test, tk Tokens) { // 2
			t.Output = StarredExpression{
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
			}
		}},
		{`* a`, func(t *test, tk Tokens) { // 3
			t.Output = StarredExpression{
				StarredList: &StarredList{
					StarredItems: []StarredItem{
						{
							OrExpr: &WrapConditional(&Atom{
								Identifier: &tk[2],
								Tokens:     tk[2:3],
							}).OrTest.AndTest.NotTest.Comparison.OrExpression,
							Tokens: tk[:3],
						},
					},
					Tokens: tk[:3],
				},
				Tokens: tk[:3],
			}
		}},
		{`nonlocal`, func(t *test, tk Tokens) { // 4
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
				Parsing: "StarredExpression",
				Token:   tk[0],
			}
		}},
	}, func(t *test) (Type, error) {
		var s StarredExpression

		err := s.parse(t.Tokens)

		return s, err
	})
}

func TestDelStatement(t *testing.T) {
	doTests(t, []sourceFn{
		{`del a`, func(t *test, tk Tokens) { // 1
			t.Output = DelStatement{
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
				Tokens: tk[:3],
			}
		}},
		{`del nonlocal`, func(t *test, tk Tokens) { // 2
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
				Parsing: "DelStatement",
				Token:   tk[2],
			}
		}},
	}, func(t *test) (Type, error) {
		var d DelStatement

		err := d.parse(t.Tokens)

		return d, err
	})
}

func TestReturnStatement(t *testing.T) {
	doTests(t, []sourceFn{
		{`return`, func(t *test, tk Tokens) { // 1
			t.Output = ReturnStatement{
				Tokens: tk[:1],
			}
		}},
		{`return a`, func(t *test, tk Tokens) { // 2
			t.Output = ReturnStatement{
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
		{`return nonlocal`, func(t *test, tk Tokens) { // 3
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
				Parsing: "ReturnStatement",
				Token:   tk[2],
			}
		}},
	}, func(t *test) (Type, error) {
		var r ReturnStatement

		err := r.parse(t.Tokens)

		return r, err
	})
}

func TestYieldExpression(t *testing.T) {
	doTests(t, []sourceFn{
		{`yield a`, func(t *test, tk Tokens) { // 1
			t.Output = YieldExpression{
				ExpressionList: &ExpressionList{
					Expressions: []Expression{
						{
							ConditionalExpression: WrapConditional(&Atom{
								Identifier: &tk[2],
								Tokens:     tk[2:3],
							}),
							Tokens: tk[2:3],
						},
					},
					Tokens: tk[2:3],
				},
				Tokens: tk[:3],
			}
		}},
		{`yield a,`, func(t *test, tk Tokens) { // 2
			t.Output = YieldExpression{
				ExpressionList: &ExpressionList{
					Expressions: []Expression{
						{
							ConditionalExpression: WrapConditional(&Atom{
								Identifier: &tk[2],
								Tokens:     tk[2:3],
							}),
							Tokens: tk[2:3],
						},
					},
					Tokens: tk[2:3],
				},
				Tokens: tk[:4],
			}
		}},
		{`yield a, b`, func(t *test, tk Tokens) { // 3
			t.Output = YieldExpression{
				ExpressionList: &ExpressionList{
					Expressions: []Expression{
						{
							ConditionalExpression: WrapConditional(&Atom{
								Identifier: &tk[2],
								Tokens:     tk[2:3],
							}),
							Tokens: tk[2:3],
						},
						{
							ConditionalExpression: WrapConditional(&Atom{
								Identifier: &tk[5],
								Tokens:     tk[5:6],
							}),
							Tokens: tk[5:6],
						},
					},
					Tokens: tk[2:6],
				},
				Tokens: tk[:6],
			}
		}},
		{`yield a, b or c`, func(t *test, tk Tokens) { // 4
			t.Output = YieldExpression{
				ExpressionList: &ExpressionList{
					Expressions: []Expression{
						{
							ConditionalExpression: WrapConditional(&Atom{
								Identifier: &tk[2],
								Tokens:     tk[2:3],
							}),
							Tokens: tk[2:3],
						},
						{
							ConditionalExpression: &ConditionalExpression{
								OrTest: OrTest{
									AndTest: WrapConditional(&Atom{
										Identifier: &tk[5],
										Tokens:     tk[5:6],
									}).OrTest.AndTest,
									OrTest: &WrapConditional(&Atom{
										Identifier: &tk[9],
										Tokens:     tk[9:10],
									}).OrTest,
									Tokens: tk[5:10],
								},
								Tokens: tk[5:10],
							},
							Tokens: tk[5:10],
						},
					},
					Tokens: tk[2:10],
				},
				Tokens: tk[:10],
			}
		}},
		{`yield a, b | c`, func(t *test, tk Tokens) { // 5
			t.Output = YieldExpression{
				ExpressionList: &ExpressionList{
					Expressions: []Expression{
						{
							ConditionalExpression: WrapConditional(&Atom{
								Identifier: &tk[2],
								Tokens:     tk[2:3],
							}),
							Tokens: tk[2:3],
						},
						{
							ConditionalExpression: WrapConditional(OrExpression{
								XorExpression: WrapConditional(&Atom{
									Identifier: &tk[5],
									Tokens:     tk[5:6],
								}).OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpression,
								OrExpression: &WrapConditional(&Atom{
									Identifier: &tk[9],
									Tokens:     tk[9:10],
								}).OrTest.AndTest.NotTest.Comparison.OrExpression,
								Tokens: tk[5:10],
							}),
							Tokens: tk[5:10],
						},
					},
					Tokens: tk[2:10],
				},
				Tokens: tk[:10],
			}
		}},
		{`yield from a`, func(t *test, tk Tokens) { // 6
			t.Output = YieldExpression{
				From: &Expression{
					ConditionalExpression: WrapConditional(&Atom{
						Identifier: &tk[4],
						Tokens:     tk[4:5],
					}),
					Tokens: tk[4:5],
				},
				Tokens: tk[:5],
			}
		}},
		{`yield nonlocal`, func(t *test, tk Tokens) { // 7
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
					Parsing: "ExpressionList",
					Token:   tk[2],
				},
				Parsing: "YieldExpression",
				Token:   tk[2],
			}
		}},
		{`yield from nonlocal`, func(t *test, tk Tokens) { // 8
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
				Parsing: "YieldExpression",
				Token:   tk[4],
			}
		}},
	}, func(t *test) (Type, error) {
		var y YieldExpression

		err := y.parse(t.Tokens)

		return y, err
	})
}

func TestRaiseStatement(t *testing.T) {
	doTests(t, []sourceFn{
		{`raise`, func(t *test, tk Tokens) { // 1
			t.Output = RaiseStatement{
				Tokens: tk[:1],
			}
		}},
		{`raise a`, func(t *test, tk Tokens) { // 2
			t.Output = RaiseStatement{
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
		{`raise a from b`, func(t *test, tk Tokens) { // 3
			t.Output = RaiseStatement{
				Expression: &Expression{
					ConditionalExpression: WrapConditional(&Atom{
						Identifier: &tk[2],
						Tokens:     tk[2:3],
					}),
					Tokens: tk[2:3],
				},
				From: &Expression{
					ConditionalExpression: WrapConditional(&Atom{
						Identifier: &tk[6],
						Tokens:     tk[6:7],
					}),
					Tokens: tk[6:7],
				},
				Tokens: tk[:7],
			}
		}},
		{`raise nonlocal from a`, func(t *test, tk Tokens) { // 4
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
				Parsing: "RaiseStatement",
				Token:   tk[2],
			}
		}},
		{`raise a from nonlocal`, func(t *test, tk Tokens) { // 5
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
				Parsing: "RaiseStatement",
				Token:   tk[6],
			}
		}},
	}, func(t *test) (Type, error) {
		var r RaiseStatement

		err := r.parse(t.Tokens)

		return r, err
	})
}

func TestImportStatement(t *testing.T) {
	doTests(t, []sourceFn{
		{`from a import *`, func(t *test, tk Tokens) { // 1
			t.Output = ImportStatement{
				RelativeModule: &RelativeModule{
					Module: &Module{
						Identifiers: []*Token{
							&tk[2],
						},
						Tokens: tk[2:3],
					},
					Tokens: tk[2:3],
				},
				Tokens: tk[:7],
			}
		}},
		{`from a import b`, func(t *test, tk Tokens) { // 2
			t.Output = ImportStatement{
				RelativeModule: &RelativeModule{
					Module: &Module{
						Identifiers: []*Token{
							&tk[2],
						},
						Tokens: tk[2:3],
					},
					Tokens: tk[2:3],
				},
				Modules: []ModuleAs{
					{
						Module: Module{
							Identifiers: []*Token{
								&tk[6],
							},
							Tokens: tk[6:7],
						},
						Tokens: tk[6:7],
					},
				},
				Tokens: tk[:7],
			}
		}},
		{`from a import b,c`, func(t *test, tk Tokens) { // 3
			t.Output = ImportStatement{
				RelativeModule: &RelativeModule{
					Module: &Module{
						Identifiers: []*Token{
							&tk[2],
						},
						Tokens: tk[2:3],
					},
					Tokens: tk[2:3],
				},
				Modules: []ModuleAs{
					{
						Module: Module{
							Identifiers: []*Token{
								&tk[6],
							},
							Tokens: tk[6:7],
						},
						Tokens: tk[6:7],
					},
					{
						Module: Module{
							Identifiers: []*Token{
								&tk[8],
							},
							Tokens: tk[8:9],
						},
						Tokens: tk[8:9],
					},
				},
				Tokens: tk[:9],
			}
		}},
		{`from a import b, c`, func(t *test, tk Tokens) { // 4
			t.Output = ImportStatement{
				RelativeModule: &RelativeModule{
					Module: &Module{
						Identifiers: []*Token{
							&tk[2],
						},
						Tokens: tk[2:3],
					},
					Tokens: tk[2:3],
				},
				Modules: []ModuleAs{
					{
						Module: Module{
							Identifiers: []*Token{
								&tk[6],
							},
							Tokens: tk[6:7],
						},
						Tokens: tk[6:7],
					},
					{
						Module: Module{
							Identifiers: []*Token{
								&tk[9],
							},
							Tokens: tk[9:10],
						},
						Tokens: tk[9:10],
					},
				},
				Tokens: tk[:10],
			}
		}},
		{`from a import (b)`, func(t *test, tk Tokens) { // 5
			t.Output = ImportStatement{
				RelativeModule: &RelativeModule{
					Module: &Module{
						Identifiers: []*Token{
							&tk[2],
						},
						Tokens: tk[2:3],
					},
					Tokens: tk[2:3],
				},
				Modules: []ModuleAs{
					{
						Module: Module{
							Identifiers: []*Token{
								&tk[7],
							},
							Tokens: tk[7:8],
						},
						Tokens: tk[7:8],
					},
				},
				Tokens: tk[:9],
			}
		}},
		{`from a import (b,c)`, func(t *test, tk Tokens) { // 6
			t.Output = ImportStatement{
				RelativeModule: &RelativeModule{
					Module: &Module{
						Identifiers: []*Token{
							&tk[2],
						},
						Tokens: tk[2:3],
					},
					Tokens: tk[2:3],
				},
				Modules: []ModuleAs{
					{
						Module: Module{
							Identifiers: []*Token{
								&tk[7],
							},
							Tokens: tk[7:8],
						},
						Tokens: tk[7:8],
					},
					{
						Module: Module{
							Identifiers: []*Token{
								&tk[9],
							},
							Tokens: tk[9:10],
						},
						Tokens: tk[9:10],
					},
				},
				Tokens: tk[:11],
			}
		}},
		{`from a import (b, c)`, func(t *test, tk Tokens) { // 7
			t.Output = ImportStatement{
				RelativeModule: &RelativeModule{
					Module: &Module{
						Identifiers: []*Token{
							&tk[2],
						},
						Tokens: tk[2:3],
					},
					Tokens: tk[2:3],
				},
				Modules: []ModuleAs{
					{
						Module: Module{
							Identifiers: []*Token{
								&tk[7],
							},
							Tokens: tk[7:8],
						},
						Tokens: tk[7:8],
					},
					{
						Module: Module{
							Identifiers: []*Token{
								&tk[10],
							},
							Tokens: tk[10:11],
						},
						Tokens: tk[10:11],
					},
				},
				Tokens: tk[:12],
			}
		}},
		{`import a`, func(t *test, tk Tokens) { // 8
			t.Output = ImportStatement{
				Modules: []ModuleAs{
					{
						Module: Module{
							Identifiers: []*Token{
								&tk[2],
							},
							Tokens: tk[2:3],
						},
						Tokens: tk[2:3],
					},
				},
				Tokens: tk[:3],
			}
		}},
		{`import a,b`, func(t *test, tk Tokens) { // 9
			t.Output = ImportStatement{
				Modules: []ModuleAs{
					{
						Module: Module{
							Identifiers: []*Token{
								&tk[2],
							},
							Tokens: tk[2:3],
						},
						Tokens: tk[2:3],
					},
					{
						Module: Module{
							Identifiers: []*Token{
								&tk[4],
							},
							Tokens: tk[4:5],
						},
						Tokens: tk[4:5],
					},
				},
				Tokens: tk[:5],
			}
		}},
		{`import a, b, c`, func(t *test, tk Tokens) { // 10
			t.Output = ImportStatement{
				Modules: []ModuleAs{
					{
						Module: Module{
							Identifiers: []*Token{
								&tk[2],
							},
							Tokens: tk[2:3],
						},
						Tokens: tk[2:3],
					},
					{
						Module: Module{
							Identifiers: []*Token{
								&tk[5],
							},
							Tokens: tk[5:6],
						},
						Tokens: tk[5:6],
					},
					{
						Module: Module{
							Identifiers: []*Token{
								&tk[8],
							},
							Tokens: tk[8:9],
						},
						Tokens: tk[8:9],
					},
				},
				Tokens: tk[:9],
			}
		}},
		{`from nonlocal import a`, func(t *test, tk Tokens) { // 11
			t.Err = Error{
				Err: Error{
					Err:     ErrMissingModule,
					Parsing: "RelativeModule",
					Token:   tk[2],
				},
				Parsing: "ImportStatement",
				Token:   tk[2],
			}
		}},
		{`from a b`, func(t *test, tk Tokens) { // 12
			t.Err = Error{
				Err:     ErrMissingImport,
				Parsing: "ImportStatement",
				Token:   tk[4],
			}
		}},
		{`from a import nonlocal`, func(t *test, tk Tokens) { // 13
			t.Err = Error{
				Err: Error{
					Err: Error{
						Err:     ErrMissingIdentifier,
						Parsing: "Module",
						Token:   tk[6],
					},
					Parsing: "ModuleAs",
					Token:   tk[6],
				},
				Parsing: "ImportStatement",
				Token:   tk[6],
			}
		}},
		{`from a import (b c)`, func(t *test, tk Tokens) { // 14
			t.Err = Error{
				Err:     ErrMissingComma,
				Parsing: "ImportStatement",
				Token:   tk[9],
			}
		}},
		{`import nonlocal`, func(t *test, tk Tokens) { // 15
			t.Err = Error{
				Err: Error{
					Err: Error{
						Err:     ErrMissingIdentifier,
						Parsing: "Module",
						Token:   tk[2],
					},
					Parsing: "ModuleAs",
					Token:   tk[2],
				},
				Parsing: "ImportStatement",
				Token:   tk[2],
			}
		}},
		{`import (b)`, func(t *test, tk Tokens) { // 16
			t.Err = Error{
				Err: Error{
					Err: Error{
						Err:     ErrMissingIdentifier,
						Parsing: "Module",
						Token:   tk[2],
					},
					Parsing: "ModuleAs",
					Token:   tk[2],
				},
				Parsing: "ImportStatement",
				Token:   tk[2],
			}
		}},
	}, func(t *test) (Type, error) {
		var i ImportStatement

		err := i.parse(t.Tokens)

		return i, err
	})
}

func TestRelativeModule(t *testing.T) {
	doTests(t, []sourceFn{
		{`a`, func(t *test, tk Tokens) { // 1
			t.Output = RelativeModule{
				Module: &Module{
					Identifiers: []*Token{
						&tk[0],
					},
					Tokens: tk[:1],
				},
				Tokens: tk[:1],
			}
		}},
		{`.a`, func(t *test, tk Tokens) { // 2
			t.Output = RelativeModule{
				Dots: 1,
				Module: &Module{
					Identifiers: []*Token{
						&tk[1],
					},
					Tokens: tk[1:2],
				},
				Tokens: tk[:2],
			}
		}},
		{`..a`, func(t *test, tk Tokens) { // 3
			t.Output = RelativeModule{
				Dots: 2,
				Module: &Module{
					Identifiers: []*Token{
						&tk[2],
					},
					Tokens: tk[2:3],
				},
				Tokens: tk[:3],
			}
		}},
		{`. . a`, func(t *test, tk Tokens) { // 4
			t.Output = RelativeModule{
				Dots: 2,
				Module: &Module{
					Identifiers: []*Token{
						&tk[4],
					},
					Tokens: tk[4:5],
				},
				Tokens: tk[:5],
			}
		}},
		{`.a.b`, func(t *test, tk Tokens) { // 5
			t.Output = RelativeModule{
				Dots: 1,
				Module: &Module{
					Identifiers: []*Token{
						&tk[1],
						&tk[3],
					},
					Tokens: tk[1:4],
				},
				Tokens: tk[:4],
			}
		}},
		{`.`, func(t *test, tk Tokens) { // 6
			t.Output = RelativeModule{
				Dots:   1,
				Tokens: tk[:1],
			}
		}},
		{`...`, func(t *test, tk Tokens) { // 7
			t.Output = RelativeModule{
				Dots:   3,
				Tokens: tk[:3],
			}
		}},
		{`nonlocal`, func(t *test, tk Tokens) { // 8
			t.Err = Error{
				Err:     ErrMissingModule,
				Parsing: "RelativeModule",
				Token:   tk[0],
			}
		}},
		{`.a.nonlocal`, func(t *test, tk Tokens) { // 9
			t.Err = Error{
				Err: Error{
					Err:     ErrMissingIdentifier,
					Parsing: "Module",
					Token:   tk[3],
				},
				Parsing: "RelativeModule",
				Token:   tk[1],
			}
		}},
	}, func(t *test) (Type, error) {
		var r RelativeModule

		err := r.parse(t.Tokens)

		return r, err
	})
}

func TestModuleAs(t *testing.T) {
	doTests(t, []sourceFn{
		{`a`, func(t *test, tk Tokens) { // 1
			t.Output = ModuleAs{
				Module: Module{
					Identifiers: []*Token{
						&tk[0],
					},
					Tokens: tk[:1],
				},
				Tokens: tk[:1],
			}
		}},
		{`a as b`, func(t *test, tk Tokens) { // 2
			t.Output = ModuleAs{
				Module: Module{
					Identifiers: []*Token{
						&tk[0],
					},
					Tokens: tk[:1],
				},
				As:     &tk[4],
				Tokens: tk[:5],
			}
		}},
		{`nonlocal`, func(t *test, tk Tokens) { // 3
			t.Err = Error{
				Err: Error{
					Err:     ErrMissingIdentifier,
					Parsing: "Module",
					Token:   tk[0],
				},
				Parsing: "ModuleAs",
				Token:   tk[0],
			}
		}},
		{`a as nonlocal`, func(t *test, tk Tokens) { // 4
			t.Err = Error{
				Err:     ErrMissingIdentifier,
				Parsing: "ModuleAs",
				Token:   tk[4],
			}
		}},
	}, func(t *test) (Type, error) {
		var m ModuleAs

		err := m.parse(t.Tokens)

		return m, err
	})
}

func TestModule(t *testing.T) {
	doTests(t, []sourceFn{
		{`a`, func(t *test, tk Tokens) { // 1
			t.Output = Module{
				Identifiers: []*Token{
					&tk[0],
				},
				Tokens: tk[:1],
			}
		}},
		{`a.b`, func(t *test, tk Tokens) { // 2
			t.Output = Module{
				Identifiers: []*Token{
					&tk[0],
					&tk[2],
				},
				Tokens: tk[:3],
			}
		}},
		{`a . b . c`, func(t *test, tk Tokens) { // 3
			t.Output = Module{
				Identifiers: []*Token{
					&tk[0],
					&tk[4],
					&tk[8],
				},
				Tokens: tk[:9],
			}
		}},
		{`nonlocal`, func(t *test, tk Tokens) { // 4
			t.Err = Error{
				Err:     ErrMissingIdentifier,
				Parsing: "Module",
				Token:   tk[0],
			}
		}},
		{`a.nonlocal`, func(t *test, tk Tokens) { // 5
			t.Err = Error{
				Err:     ErrMissingIdentifier,
				Parsing: "Module",
				Token:   tk[2],
			}
		}},
	}, func(t *test) (Type, error) {
		var m Module

		err := m.parse(t.Tokens)

		return m, err
	})
}

func TestGlobalStatement(t *testing.T) {
	doTests(t, []sourceFn{
		{`global a`, func(t *test, tk Tokens) { // 1
			t.Output = GlobalStatement{
				Identifiers: []*Token{
					&tk[2],
				},
				Tokens: tk[:3],
			}
		}},
		{`global a,b`, func(t *test, tk Tokens) { // 2
			t.Output = GlobalStatement{
				Identifiers: []*Token{
					&tk[2],
					&tk[4],
				},
				Tokens: tk[:5],
			}
		}},
		{`global a, b, c`, func(t *test, tk Tokens) { // 3
			t.Output = GlobalStatement{
				Identifiers: []*Token{
					&tk[2],
					&tk[5],
					&tk[8],
				},
				Tokens: tk[:9],
			}
		}},
		{`global nonlocal`, func(t *test, tk Tokens) { // 4
			t.Err = Error{
				Err:     ErrMissingIdentifier,
				Parsing: "GlobalStatement",
				Token:   tk[2],
			}
		}},
		{`global a, nonlocal`, func(t *test, tk Tokens) { // 5
			t.Err = Error{
				Err:     ErrMissingIdentifier,
				Parsing: "GlobalStatement",
				Token:   tk[5],
			}
		}},
	}, func(t *test) (Type, error) {
		var g GlobalStatement

		err := g.parse(t.Tokens)

		return g, err
	})
}

func TestNonLocalStatement(t *testing.T) {
	doTests(t, []sourceFn{
		{`nonlocal a`, func(t *test, tk Tokens) { // 1
			t.Output = NonLocalStatement{
				Identifiers: []*Token{
					&tk[2],
				},
				Tokens: tk[:3],
			}
		}},
		{`nonlocal a,b`, func(t *test, tk Tokens) { // 2
			t.Output = NonLocalStatement{
				Identifiers: []*Token{
					&tk[2],
					&tk[4],
				},
				Tokens: tk[:5],
			}
		}},
		{`nonlocal a, b, c`, func(t *test, tk Tokens) { // 3
			t.Output = NonLocalStatement{
				Identifiers: []*Token{
					&tk[2],
					&tk[5],
					&tk[8],
				},
				Tokens: tk[:9],
			}
		}},
		{`nonlocal nonlocal`, func(t *test, tk Tokens) { // 4
			t.Err = Error{
				Err:     ErrMissingIdentifier,
				Parsing: "NonLocalStatement",
				Token:   tk[2],
			}
		}},
		{`nonlocal a, nonlocal`, func(t *test, tk Tokens) { // 5
			t.Err = Error{
				Err:     ErrMissingIdentifier,
				Parsing: "NonLocalStatement",
				Token:   tk[5],
			}
		}},
	}, func(t *test) (Type, error) {
		var n NonLocalStatement

		err := n.parse(t.Tokens)

		return n, err
	})
}

func TestTypeStatement(t *testing.T) {
	doTests(t, []sourceFn{
		{`type a=b`, func(t *test, tk Tokens) { // 1
			t.Output = TypeStatement{
				Identifier: &tk[2],
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
		{`type a = b`, func(t *test, tk Tokens) { // 2
			t.Output = TypeStatement{
				Identifier: &tk[2],
				Expression: Expression{
					ConditionalExpression: WrapConditional(&Atom{
						Identifier: &tk[6],
						Tokens:     tk[6:7],
					}),
					Tokens: tk[6:7],
				},
				Tokens: tk[:7],
			}
		}},
		{`type a[b] = c`, func(t *test, tk Tokens) { // 3
			t.Output = TypeStatement{
				Identifier: &tk[2],
				TypeParams: &TypeParams{
					TypeParams: []TypeParam{
						{
							Identifier: &tk[4],
							Tokens:     tk[4:5],
						},
					},
					Tokens: tk[3:6],
				},
				Expression: Expression{
					ConditionalExpression: WrapConditional(&Atom{
						Identifier: &tk[9],
						Tokens:     tk[9:10],
					}),
					Tokens: tk[9:10],
				},
				Tokens: tk[:10],
			}
		}},
		{`type a [b] = c`, func(t *test, tk Tokens) { // 4
			t.Output = TypeStatement{
				Identifier: &tk[2],
				TypeParams: &TypeParams{
					TypeParams: []TypeParam{
						{
							Identifier: &tk[5],
							Tokens:     tk[5:6],
						},
					},
					Tokens: tk[4:7],
				},
				Expression: Expression{
					ConditionalExpression: WrapConditional(&Atom{
						Identifier: &tk[10],
						Tokens:     tk[10:11],
					}),
					Tokens: tk[10:11],
				},
				Tokens: tk[:11],
			}
		}},
		{`type nonlocal[a] = b`, func(t *test, tk Tokens) { // 5
			t.Err = Error{
				Err:     ErrMissingIdentifier,
				Parsing: "TypeStatement",
				Token:   tk[2],
			}
		}},
		{`type a[nonlocal] = b`, func(t *test, tk Tokens) { // 6
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
				Parsing: "TypeStatement",
				Token:   tk[3],
			}
		}},
		{`type a[b] = nonlocal`, func(t *test, tk Tokens) { // 7
			t.Err = Error{
				Err: Error{
					Err: wrapConditionalExpressionError(Error{
						Err:     ErrInvalidEnclosure,
						Parsing: "Enclosure",
						Token:   tk[9],
					}),
					Parsing: "Expression",
					Token:   tk[9],
				},
				Parsing: "TypeStatement",
				Token:   tk[9],
			}
		}},
		{`type a`, func(t *test, tk Tokens) { // 8
			t.Err = Error{
				Err:     ErrMissingEquals,
				Parsing: "TypeStatement",
				Token:   tk[3],
			}
		}},
	}, func(t *test) (Type, error) {
		var ts TypeStatement

		err := ts.parse(t.Tokens)

		return ts, err
	})
}

func TestTypeParams(t *testing.T) {
	doTests(t, []sourceFn{
		{`[a]`, func(t *test, tk Tokens) { // 1
			t.Output = TypeParams{
				TypeParams: []TypeParam{
					{
						Identifier: &tk[1],
						Tokens:     tk[1:2],
					},
				},
				Tokens: tk[:3],
			}
		}},
		{`[ a ]`, func(t *test, tk Tokens) { // 2
			t.Output = TypeParams{
				TypeParams: []TypeParam{
					{
						Identifier: &tk[2],
						Tokens:     tk[2:3],
					},
				},
				Tokens: tk[:5],
			}
		}},
		{`[a, b]`, func(t *test, tk Tokens) { // 3
			t.Output = TypeParams{
				TypeParams: []TypeParam{
					{
						Identifier: &tk[1],
						Tokens:     tk[1:2],
					},
					{
						Identifier: &tk[4],
						Tokens:     tk[4:5],
					},
				},
				Tokens: tk[:6],
			}
		}},
		{`[nonlocal]`, func(t *test, tk Tokens) { // 4
			t.Err = Error{
				Err: Error{
					Err:     ErrMissingIdentifier,
					Parsing: "TypeParam",
					Token:   tk[1],
				},
				Parsing: "TypeParams",
				Token:   tk[1],
			}
		}},
		{`[a b]`, func(t *test, tk Tokens) { // 5
			t.Err = Error{
				Err:     ErrMissingComma,
				Parsing: "TypeParams",
				Token:   tk[3],
			}
		}},
	}, func(t *test) (Type, error) {
		var tp TypeParams

		err := tp.parse(t.Tokens)

		return tp, err
	})
}

func TestAssignmentExpression(t *testing.T) {
	doTests(t, []sourceFn{
		{`a`, func(t *test, tk Tokens) { // 1
			t.Output = AssignmentExpression{
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
		{`a:=b`, func(t *test, tk Tokens) { // 2
			t.Output = AssignmentExpression{
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
		{`a := b`, func(t *test, tk Tokens) { // 3
			t.Output = AssignmentExpression{
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
		{`nonlocal := a`, func(t *test, tk Tokens) { // 4
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
				Parsing: "AssignmentExpression",
				Token:   tk[0],
			}
		}},
		{`a := nonlocal`, func(t *test, tk Tokens) { // 5
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
				Parsing: "AssignmentExpression",
				Token:   tk[4],
			}
		}},
	}, func(t *test) (Type, error) {
		var a AssignmentExpression

		err := a.parse(t.Tokens)

		return a, err
	})
}

func TestExpression(t *testing.T) {
	doTests(t, []sourceFn{
		{`a`, func(t *test, tk Tokens) { // 1
			t.Output = Expression{
				ConditionalExpression: WrapConditional(&Atom{
					Identifier: &tk[0],
					Tokens:     tk[:1],
				}),
				Tokens: tk[:1],
			}
		}},
		{`lambda:a`, func(t *test, tk Tokens) { // 2
			t.Output = Expression{
				LambdaExpression: &LambdaExpression{
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
		{`nonlocal`, func(t *test, tk Tokens) { // 3
			t.Err = Error{
				Err: wrapConditionalExpressionError(Error{
					Err:     ErrInvalidEnclosure,
					Parsing: "Enclosure",
					Token:   tk[0],
				}),
				Parsing: "Expression",
				Token:   tk[0],
			}
		}},
		{`lambda:nonlocal`, func(t *test, tk Tokens) { // 4
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
					Parsing: "LambdaExpression",
					Token:   tk[2],
				},
				Parsing: "Expression",
				Token:   tk[0],
			}
		}},
	}, func(t *test) (Type, error) {
		var e Expression

		err := e.parse(t.Tokens)

		return e, err
	})
}

func TestConditionalExpression(t *testing.T) {
	doTests(t, []sourceFn{
		{`a`, func(t *test, tk Tokens) { // 1
			t.Output = ConditionalExpression{
				OrTest: OrTest{
					AndTest: AndTest{
						NotTest: NotTest{
							Comparison: Comparison{
								OrExpression: OrExpression{
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
		{`a if b else c`, func(t *test, tk Tokens) { // 2
			t.Output = ConditionalExpression{
				OrTest: OrTest{
					AndTest: AndTest{
						NotTest: NotTest{
							Comparison: Comparison{
								OrExpression: OrExpression{
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
								},
								Tokens: tk[:1],
							},
							Tokens: tk[:1],
						},
						Tokens: tk[:1],
					},
					Tokens: tk[:1],
				},
				If: &OrTest{
					AndTest: AndTest{
						NotTest: NotTest{
							Comparison: Comparison{
								OrExpression: OrExpression{
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
								Tokens: tk[4:5],
							},
							Tokens: tk[4:5],
						},
						Tokens: tk[4:5],
					},
					Tokens: tk[4:5],
				},
				Else: &Expression{
					ConditionalExpression: &ConditionalExpression{
						OrTest: OrTest{
							AndTest: AndTest{
								NotTest: NotTest{
									Comparison: Comparison{
										OrExpression: OrExpression{
											XorExpression: XorExpression{
												AndExpression: AndExpression{
													ShiftExpression: ShiftExpression{
														AddExpression: AddExpression{
															MultiplyExpression: MultiplyExpression{
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
																Tokens: tk[8:9],
															},
															Tokens: tk[8:9],
														},
														Tokens: tk[8:9],
													},
													Tokens: tk[8:9],
												},
												Tokens: tk[8:9],
											},
											Tokens: tk[8:9],
										},
										Tokens: tk[8:9],
									},
									Tokens: tk[8:9],
								},
								Tokens: tk[8:9],
							},
							Tokens: tk[8:9],
						},
						Tokens: tk[8:9],
					},
					Tokens: tk[8:9],
				},
				Tokens: tk[:9],
			}
		}},
		{`nonlocal`, func(t *test, tk Tokens) { // 3
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
								},
								Parsing: "Comparison",
								Token:   tk[0],
							},
							Parsing: "NotTest",
							Token:   tk[0],
						},
						Parsing: "AndTest",
						Token:   tk[0],
					},
					Parsing: "OrTest",
					Token:   tk[0],
				},
				Parsing: "ConditionalExpression",
				Token:   tk[0],
			}
		}},
		{`a if nonlocal else c`, func(t *test, tk Tokens) { // 4
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
												Parsing: "ShiftExpression",
												Token:   tk[4],
											},
											Parsing: "AndExpression",
											Token:   tk[4],
										},
										Parsing: "XorExpression",
										Token:   tk[4],
									},
									Parsing: "OrExpression",
									Token:   tk[4],
								},
								Parsing: "Comparison",
								Token:   tk[4],
							},
							Parsing: "NotTest",
							Token:   tk[4],
						},
						Parsing: "AndTest",
						Token:   tk[4],
					},
					Parsing: "OrTest",
					Token:   tk[4],
				},
				Parsing: "ConditionalExpression",
				Token:   tk[4],
			}
		}},
		{`a if b els c`, func(t *test, tk Tokens) { // 5
			t.Err = Error{
				Err:     ErrMissingElse,
				Parsing: "ConditionalExpression",
				Token:   tk[6],
			}
		}},
		{`a if b else nonlocal`, func(t *test, tk Tokens) { // 6
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
																					Token:   tk[8],
																				},
																				Parsing: "Atom",
																				Token:   tk[8],
																			},
																			Parsing: "PrimaryExpression",
																			Token:   tk[8],
																		},
																		Parsing: "PowerExpression",
																		Token:   tk[8],
																	},
																	Parsing: "UnaryExpression",
																	Token:   tk[8],
																},
																Parsing: "MultiplyExpression",
																Token:   tk[8],
															},
															Parsing: "AddExpression",
															Token:   tk[8],
														},
														Parsing: "ShiftExpression",
														Token:   tk[8],
													},
													Parsing: "AndExpression",
													Token:   tk[8],
												},
												Parsing: "XorExpression",
												Token:   tk[8],
											},
											Parsing: "OrExpression",
											Token:   tk[8],
										},
										Parsing: "Comparison",
										Token:   tk[8],
									},
									Parsing: "NotTest",
									Token:   tk[8],
								},
								Parsing: "AndTest",
								Token:   tk[8],
							},
							Parsing: "OrTest",
							Token:   tk[8],
						},
						Parsing: "ConditionalExpression",
						Token:   tk[8],
					},
					Parsing: "Expression",
					Token:   tk[8],
				},
				Parsing: "ConditionalExpression",
				Token:   tk[8],
			}
		}},
	}, func(t *test) (Type, error) {
		var ce ConditionalExpression

		err := ce.parse(t.Tokens)

		return ce, err
	})
}

func TestLambdaExpression(t *testing.T) {
	doTests(t, []sourceFn{
		{`lambda:a`, func(t *test, tk Tokens) { // 1
			t.Output = LambdaExpression{
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
		{`lambda : a`, func(t *test, tk Tokens) { // 2
			t.Output = LambdaExpression{
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
		{`lambda a: b`, func(t *test, tk Tokens) { // 3
			t.Output = LambdaExpression{
				ParameterList: &ParameterList{
					NoPosOnly: []DefParameter{
						{
							Parameter: Parameter{
								Identifier: &tk[2],
								Tokens:     tk[2:3],
							},
							Tokens: tk[2:3],
						},
					},
					Tokens: tk[2:3],
				},
				Expression: Expression{
					ConditionalExpression: WrapConditional(&Atom{
						Identifier: &tk[5],
						Tokens:     tk[5:6],
					}),
					Tokens: tk[5:6],
				},
				Tokens: tk[:6],
			}
		}},
		{`lambda a, b : c`, func(t *test, tk Tokens) { // 4
			t.Output = LambdaExpression{
				ParameterList: &ParameterList{
					NoPosOnly: []DefParameter{
						{
							Parameter: Parameter{
								Identifier: &tk[2],
								Tokens:     tk[2:3],
							},
							Tokens: tk[2:3],
						},
						{
							Parameter: Parameter{
								Identifier: &tk[5],
								Tokens:     tk[5:6],
							},
							Tokens: tk[5:6],
						},
					},
					Tokens: tk[2:6],
				},
				Expression: Expression{
					ConditionalExpression: WrapConditional(&Atom{
						Identifier: &tk[9],
						Tokens:     tk[9:10],
					}),
					Tokens: tk[9:10],
				},
				Tokens: tk[:10],
			}
		}},
		{`lambda nonlocal: a`, func(t *test, tk Tokens) { // 5
			t.Err = Error{
				Err:     ErrMissingColon,
				Parsing: "LambdaExpression",
				Token:   tk[2],
			}
		}},
		{`lambda: nonlocal`, func(t *test, tk Tokens) { // 6
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
				Parsing: "LambdaExpression",
				Token:   tk[3],
			}
		}},
		{`lambda *nonlocal`, func(t *test, tk Tokens) { // 7
			t.Err = Error{
				Err: Error{
					Err: Error{
						Err:     ErrMissingIdentifier,
						Parsing: "Parameter",
						Token:   tk[3],
					},
					Parsing: "ParameterList",
					Token:   tk[3],
				},
				Parsing: "LambdaExpression",
				Token:   tk[2],
			}
		}},
	}, func(t *test) (Type, error) {
		var le LambdaExpression

		err := le.parse(t.Tokens)

		return le, err
	})
}

func TestOrTest(t *testing.T) {
	doTests(t, []sourceFn{
		{`a`, func(t *test, tk Tokens) { // 1
			t.Output = OrTest{
				AndTest: AndTest{
					NotTest: NotTest{
						Comparison: Comparison{
							OrExpression: OrExpression{
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
		{`a or b`, func(t *test, tk Tokens) { // 2
			t.Output = OrTest{
				AndTest: AndTest{
					NotTest: NotTest{
						Comparison: Comparison{
							OrExpression: OrExpression{
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
							},
							Tokens: tk[:1],
						},
						Tokens: tk[:1],
					},
					Tokens: tk[:1],
				},
				OrTest: &OrTest{
					AndTest: AndTest{
						NotTest: NotTest{
							Comparison: Comparison{
								OrExpression: OrExpression{
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
		{`nonlocal`, func(t *test, tk Tokens) { // 3
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
							},
							Parsing: "Comparison",
							Token:   tk[0],
						},
						Parsing: "NotTest",
						Token:   tk[0],
					},
					Parsing: "AndTest",
					Token:   tk[0],
				},
				Parsing: "OrTest",
				Token:   tk[0],
			}
		}},
		{`a or nonlocal`, func(t *test, tk Tokens) { // 4
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
												Parsing: "ShiftExpression",
												Token:   tk[4],
											},
											Parsing: "AndExpression",
											Token:   tk[4],
										},
										Parsing: "XorExpression",
										Token:   tk[4],
									},
									Parsing: "OrExpression",
									Token:   tk[4],
								},
								Parsing: "Comparison",
								Token:   tk[4],
							},
							Parsing: "NotTest",
							Token:   tk[4],
						},
						Parsing: "AndTest",
						Token:   tk[4],
					},
					Parsing: "OrTest",
					Token:   tk[4],
				},
				Parsing: "OrTest",
				Token:   tk[4],
			}
		}},
	}, func(t *test) (Type, error) {
		var ot OrTest

		err := ot.parse(t.Tokens)

		return ot, err
	})
}

func TestAndTest(t *testing.T) {
	doTests(t, []sourceFn{
		{`a`, func(t *test, tk Tokens) { // 1
			t.Output = AndTest{
				NotTest: NotTest{
					Comparison: Comparison{
						OrExpression: OrExpression{
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
						},
						Tokens: tk[:1],
					},
					Tokens: tk[:1],
				},
				Tokens: tk[:1],
			}
		}},
		{`a and b`, func(t *test, tk Tokens) { // 2
			t.Output = AndTest{
				NotTest: NotTest{
					Comparison: Comparison{
						OrExpression: OrExpression{
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
						},
						Tokens: tk[:1],
					},
					Tokens: tk[:1],
				},
				AndTest: &AndTest{
					NotTest: NotTest{
						Comparison: Comparison{
							OrExpression: OrExpression{
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
							Tokens: tk[4:5],
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
						},
						Parsing: "Comparison",
						Token:   tk[0],
					},
					Parsing: "NotTest",
					Token:   tk[0],
				},
				Parsing: "AndTest",
				Token:   tk[0],
			}
		}},
		{`a and nonlocal`, func(t *test, tk Tokens) { // 4
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
											Parsing: "ShiftExpression",
											Token:   tk[4],
										},
										Parsing: "AndExpression",
										Token:   tk[4],
									},
									Parsing: "XorExpression",
									Token:   tk[4],
								},
								Parsing: "OrExpression",
								Token:   tk[4],
							},
							Parsing: "Comparison",
							Token:   tk[4],
						},
						Parsing: "NotTest",
						Token:   tk[4],
					},
					Parsing: "AndTest",
					Token:   tk[4],
				},
				Parsing: "AndTest",
				Token:   tk[4],
			}
		}},
	}, func(t *test) (Type, error) {
		var at AndTest

		err := at.parse(t.Tokens)

		return at, err
	})
}

func TestNotTest(t *testing.T) {
	doTests(t, []sourceFn{
		{`a`, func(t *test, tk Tokens) { // 1
			t.Output = NotTest{
				Comparison: Comparison{
					OrExpression: OrExpression{
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
					},
					Tokens: tk[:1],
				},
				Tokens: tk[:1],
			}
		}},
		{`not a`, func(t *test, tk Tokens) { // 2
			t.Output = NotTest{
				Nots: 1,
				Comparison: Comparison{
					OrExpression: OrExpression{
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
					Tokens: tk[2:3],
				},
				Tokens: tk[:3],
			}
		}},
		{`not not not not a`, func(t *test, tk Tokens) { // 3
			t.Output = NotTest{
				Nots: 4,
				Comparison: Comparison{
					OrExpression: OrExpression{
						XorExpression: XorExpression{
							AndExpression: AndExpression{
								ShiftExpression: ShiftExpression{
									AddExpression: AddExpression{
										MultiplyExpression: MultiplyExpression{
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
											Tokens: tk[8:9],
										},
										Tokens: tk[8:9],
									},
									Tokens: tk[8:9],
								},
								Tokens: tk[8:9],
							},
							Tokens: tk[8:9],
						},
						Tokens: tk[8:9],
					},
					Tokens: tk[8:9],
				},
				Tokens: tk[:9],
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
					},
					Parsing: "Comparison",
					Token:   tk[0],
				},
				Parsing: "NotTest",
				Token:   tk[0],
			}
		}},
		{`not not not nonlocal`, func(t *test, tk Tokens) { // 5
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
																Err:     ErrInvalidEnclosure,
																Parsing: "Enclosure",
																Token:   tk[6],
															},
															Parsing: "Atom",
															Token:   tk[6],
														},
														Parsing: "PrimaryExpression",
														Token:   tk[6],
													},
													Parsing: "PowerExpression",
													Token:   tk[6],
												},
												Parsing: "UnaryExpression",
												Token:   tk[6],
											},
											Parsing: "MultiplyExpression",
											Token:   tk[6],
										},
										Parsing: "AddExpression",
										Token:   tk[6],
									},
									Parsing: "ShiftExpression",
									Token:   tk[6],
								},
								Parsing: "AndExpression",
								Token:   tk[6],
							},
							Parsing: "XorExpression",
							Token:   tk[6],
						},
						Parsing: "OrExpression",
						Token:   tk[6],
					},
					Parsing: "Comparison",
					Token:   tk[6],
				},
				Parsing: "NotTest",
				Token:   tk[6],
			}
		}},
	}, func(t *test) (Type, error) {
		var nt NotTest

		err := nt.parse(t.Tokens)

		return nt, err
	})
}

func TestComparison(t *testing.T) {
	doTests(t, []sourceFn{
		{`a`, func(t *test, tk Tokens) { // 1
			t.Output = Comparison{
				OrExpression: OrExpression{
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
				},
				Tokens: tk[:1],
			}
		}},
		{`a < b`, func(t *test, tk Tokens) { // 2
			t.Output = Comparison{
				OrExpression: OrExpression{
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
				},
				Comparisons: []ComparisonExpression{
					{
						ComparisonOperator: []Token{tk[2]},
						OrExpression: OrExpression{
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
					},
				},
				Tokens: tk[:5],
			}
		}},
		{`a>b`, func(t *test, tk Tokens) { // 3
			t.Output = Comparison{
				OrExpression: OrExpression{
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
				},
				Comparisons: []ComparisonExpression{
					{
						ComparisonOperator: []Token{tk[1]},
						OrExpression: OrExpression{
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
					},
				},
				Tokens: tk[:3],
			}
		}},
		{`a==b>=c <= d != e`, func(t *test, tk Tokens) { // 4
			t.Output = Comparison{
				OrExpression: OrExpression{
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
				},
				Comparisons: []ComparisonExpression{
					{
						ComparisonOperator: []Token{tk[1]},
						OrExpression: OrExpression{
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
					},
					{
						ComparisonOperator: []Token{tk[3]},
						OrExpression: OrExpression{
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
					},
					{
						ComparisonOperator: []Token{tk[6]},
						OrExpression: OrExpression{
							XorExpression: XorExpression{
								AndExpression: AndExpression{
									ShiftExpression: ShiftExpression{
										AddExpression: AddExpression{
											MultiplyExpression: MultiplyExpression{
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
												Tokens: tk[8:9],
											},
											Tokens: tk[8:9],
										},
										Tokens: tk[8:9],
									},
									Tokens: tk[8:9],
								},
								Tokens: tk[8:9],
							},
							Tokens: tk[8:9],
						},
					},
					{
						ComparisonOperator: []Token{tk[10]},
						OrExpression: OrExpression{
							XorExpression: XorExpression{
								AndExpression: AndExpression{
									ShiftExpression: ShiftExpression{
										AddExpression: AddExpression{
											MultiplyExpression: MultiplyExpression{
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
											Tokens: tk[12:13],
										},
										Tokens: tk[12:13],
									},
									Tokens: tk[12:13],
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
		{`a is b`, func(t *test, tk Tokens) { // 5
			t.Output = Comparison{
				OrExpression: OrExpression{
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
				},
				Comparisons: []ComparisonExpression{
					{
						ComparisonOperator: []Token{tk[2]},
						OrExpression: OrExpression{
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
					},
				},
				Tokens: tk[:5],
			}
		}},
		{`a is not b`, func(t *test, tk Tokens) { // 6
			t.Output = Comparison{
				OrExpression: OrExpression{
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
				},
				Comparisons: []ComparisonExpression{
					{
						ComparisonOperator: []Token{tk[2], tk[3], tk[4]},
						OrExpression: OrExpression{
							XorExpression: XorExpression{
								AndExpression: AndExpression{
									ShiftExpression: ShiftExpression{
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
									Tokens: tk[6:7],
								},
								Tokens: tk[6:7],
							},
							Tokens: tk[6:7],
						},
					},
				},
				Tokens: tk[:7],
			}
		}},
		{`a not in b`, func(t *test, tk Tokens) { // 7
			t.Output = Comparison{
				OrExpression: OrExpression{
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
				},
				Comparisons: []ComparisonExpression{
					{
						ComparisonOperator: []Token{tk[2], tk[3], tk[4]},
						OrExpression: OrExpression{
							XorExpression: XorExpression{
								AndExpression: AndExpression{
									ShiftExpression: ShiftExpression{
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
									Tokens: tk[6:7],
								},
								Tokens: tk[6:7],
							},
							Tokens: tk[6:7],
						},
					},
				},
				Tokens: tk[:7],
			}
		}},
		{`a in b`, func(t *test, tk Tokens) { // 8
			t.Output = Comparison{
				OrExpression: OrExpression{
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
				},
				Comparisons: []ComparisonExpression{
					{
						ComparisonOperator: []Token{tk[2]},
						OrExpression: OrExpression{
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
					},
				},
				Tokens: tk[:5],
			}
		}},
		{`nonlocal`, func(t *test, tk Tokens) { // 9
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
				},
				Parsing: "Comparison",
				Token:   tk[0],
			}
		}},
		{`a not a b`, func(t *test, tk Tokens) { // 10
			t.Err = Error{
				Err:     ErrMissingIn,
				Parsing: "Comparison",
				Token:   tk[4],
			}
		}},
		{`1<nonlocal`, func(t *test, tk Tokens) { // 11
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
				Parsing: "Comparison",
				Token:   tk[2],
			}
		}},
	}, func(t *test) (Type, error) {
		var c Comparison

		err := c.parse(t.Tokens)

		return c, err
	})
}

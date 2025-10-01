package python

import "testing"

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
		{"(#abc\na\n#def\n)", func(t *test, tk Tokens) { // 8
			t.Output = TargetList{
				Targets: []Target{
					{
						PrimaryExpression: &PrimaryExpression{
							Atom: &Atom{
								Identifier: &tk[3],
								Tokens:     tk[3:4],
							},
							Tokens: tk[3:4],
						},
						Tokens: tk[3:4],
					},
				},
				Comments: [2]Comments{{tk[1]}, {tk[5]}},
				Tokens:   tk[1:6],
			}
		}},
		{"(#abc\na,b\n#def\n)", func(t *test, tk Tokens) { // 9
			t.Output = TargetList{
				Targets: []Target{
					{
						PrimaryExpression: &PrimaryExpression{
							Atom: &Atom{
								Identifier: &tk[3],
								Tokens:     tk[3:4],
							},
							Tokens: tk[3:4],
						},
						Tokens: tk[3:4],
					},
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
				Comments: [2]Comments{{tk[1]}, {tk[7]}},
				Tokens:   tk[1:8],
			}
		}},
		{"(#abc\na,#def\n)", func(t *test, tk Tokens) { // 10
			t.Output = TargetList{
				Targets: []Target{
					{
						PrimaryExpression: &PrimaryExpression{
							Atom: &Atom{
								Identifier: &tk[3],
								Tokens:     tk[3:4],
							},
							Tokens: tk[3:4],
						},
						Tokens: tk[3:4],
					},
				},
				Comments: [2]Comments{{tk[1]}},
				Tokens:   tk[1:4],
			}
		}},
		{`nonlocal`, func(t *test, tk Tokens) { // 11
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

		if t.Tokens.Peek().Data == "(" {
			t.Tokens.Tokens = t.Tokens.Tokens[1:1]
		}

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
		{"( * # A\na )", func(t *test, tk Tokens) { // 10
			t.Output = Target{
				Tuple: &TargetList{
					Targets: []Target{
						{
							Star: &Target{
								PrimaryExpression: &PrimaryExpression{
									Atom: &Atom{
										Identifier: &tk[6],
										Tokens:     tk[6:7],
									},
									Tokens: tk[6:7],
								},
								Comments: [2]Comments{{tk[4]}},
								Tokens:   tk[4:7],
							},
							Tokens: tk[2:7],
						},
					},
					Tokens: tk[2:7],
				},
				Tokens: tk[:9],
			}
		}},
		{`nonlocal`, func(t *test, tk Tokens) { // 11
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
		{`(nonlocal)`, func(t *test, tk Tokens) { // 12
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
		{`[nonlocal]`, func(t *test, tk Tokens) { // 13
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
		{`*nonlocal`, func(t *test, tk Tokens) { // 14
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
		{`a()`, func(t *test, tk Tokens) { // 15
			t.Err = Error{
				Err:     ErrMissingIdentifier,
				Parsing: "Target",
				Token:   tk[0],
			}
		}},
		{`{a}`, func(t *test, tk Tokens) { // 16
			t.Err = Error{
				Err:     ErrMissingIdentifier,
				Parsing: "Target",
				Token:   tk[0],
			}
		}},
		{`(a b)`, func(t *test, tk Tokens) { // 17
			t.Err = Error{
				Err:     ErrMissingClosingParen,
				Parsing: "Target",
				Token:   tk[3],
			}
		}},
		{`[a b]`, func(t *test, tk Tokens) { // 18
			t.Err = Error{
				Err:     ErrMissingClosingBracket,
				Parsing: "Target",
				Token:   tk[3],
			}
		}},
	}, func(t *test) (Type, error) {
		var tt Target

		err := tt.parse(t.Tokens)

		return tt, err
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
								Identifiers: []IdentifierComments{
									{
										Identifier: &tk[2],
									},
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
							Identifiers: []IdentifierComments{
								{
									Identifier: &tk[2],
								},
							},
							Tokens: tk[2:3],
						},
						Tokens: tk[2:3],
					},
					Modules: []ModuleAs{
						{
							Module: Module{
								Identifiers: []IdentifierComments{
									{
										Identifier: &tk[6],
									},
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
		{`yield nonlocal`, func(t *test, tk Tokens) { // 20
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
					Parsing: "YieldExpression",
					Token:   tk[2],
				},
				Parsing: "SimpleStatement",
				Token:   tk[0],
			}
		}},
		{`raise nonlocal from a`, func(t *test, tk Tokens) { // 21
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
		{`import nonlocal`, func(t *test, tk Tokens) { // 22
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
		{`global nonlocal`, func(t *test, tk Tokens) { // 23
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
		{`nonlocal nonlocal`, func(t *test, tk Tokens) { // 24
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
		{`type nonlocal[a] = b`, func(t *test, tk Tokens) { // 25
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
		{`a=yield nonlocal`, func(t *test, tk Tokens) { // 26
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
		{`a:nonlocal`, func(t *test, tk Tokens) { // 27
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
		{`a/=nonlocal`, func(t *test, tk Tokens) { // 28
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
		{"(# A\nyield # B\na # C\n)", func(t *test, tk Tokens) { // 7
			t.Output = YieldExpression{
				ExpressionList: &ExpressionList{
					Expressions: []Expression{
						{
							ConditionalExpression: WrapConditional(&Atom{
								Identifier: &tk[7],
								Tokens:     tk[7:8],
							}),
							Tokens: tk[7:8],
						},
					},
					Tokens: tk[7:8],
				},
				Comments: [4]Comments{{tk[1]}, {tk[5]}, nil, {tk[9]}},
				Tokens:   tk[1:10],
			}
		}},
		{"(# A\nyield # B\na # C\n, # D\n)", func(t *test, tk Tokens) { // 8
			t.Output = YieldExpression{
				ExpressionList: &ExpressionList{
					Expressions: []Expression{
						{
							ConditionalExpression: WrapConditional(&Atom{
								Identifier: &tk[7],
								Tokens:     tk[7:8],
							}),
							Tokens: tk[7:8],
						},
					},
					Tokens: tk[7:8],
				},
				Comments: [4]Comments{{tk[1]}, {tk[5]}, {tk[9]}, {tk[13]}},
				Tokens:   tk[1:14],
			}
		}},
		{"(# A\nyield # B\nfrom # C\na # D\n)", func(t *test, tk Tokens) { // 9
			t.Output = YieldExpression{
				From: &Expression{
					ConditionalExpression: WrapConditional(&Atom{
						Identifier: &tk[11],
						Tokens:     tk[11:12],
					}),
					Tokens: tk[11:12],
				},
				Comments: [4]Comments{{tk[1]}, {tk[5]}, {tk[9]}, {tk[13]}},
				Tokens:   tk[1:14],
			}
		}},
		{`yield nonlocal`, func(t *test, tk Tokens) { // 10
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
		{`yield from nonlocal`, func(t *test, tk Tokens) { // 11
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

		if t.Tokens.Peek().Data == "(" {
			t.Tokens.Tokens = t.Tokens.Tokens[1:1]
		}

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
						Identifiers: []IdentifierComments{
							{
								Identifier: &tk[2],
							},
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
						Identifiers: []IdentifierComments{
							{
								Identifier: &tk[2],
							},
						},
						Tokens: tk[2:3],
					},
					Tokens: tk[2:3],
				},
				Modules: []ModuleAs{
					{
						Module: Module{
							Identifiers: []IdentifierComments{
								{
									Identifier: &tk[6],
								},
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
						Identifiers: []IdentifierComments{
							{
								Identifier: &tk[2],
							},
						},
						Tokens: tk[2:3],
					},
					Tokens: tk[2:3],
				},
				Modules: []ModuleAs{
					{
						Module: Module{
							Identifiers: []IdentifierComments{
								{
									Identifier: &tk[6],
								},
							},
							Tokens: tk[6:7],
						},
						Tokens: tk[6:7],
					},
					{
						Module: Module{
							Identifiers: []IdentifierComments{
								{
									Identifier: &tk[8],
								},
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
						Identifiers: []IdentifierComments{
							{
								Identifier: &tk[2],
							},
						},
						Tokens: tk[2:3],
					},
					Tokens: tk[2:3],
				},
				Modules: []ModuleAs{
					{
						Module: Module{
							Identifiers: []IdentifierComments{
								{
									Identifier: &tk[6],
								},
							},
							Tokens: tk[6:7],
						},
						Tokens: tk[6:7],
					},
					{
						Module: Module{
							Identifiers: []IdentifierComments{
								{
									Identifier: &tk[9],
								},
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
						Identifiers: []IdentifierComments{
							{
								Identifier: &tk[2],
							},
						},
						Tokens: tk[2:3],
					},
					Tokens: tk[2:3],
				},
				Modules: []ModuleAs{
					{
						Module: Module{
							Identifiers: []IdentifierComments{
								{
									Identifier: &tk[7],
								},
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
						Identifiers: []IdentifierComments{
							{
								Identifier: &tk[2],
							},
						},
						Tokens: tk[2:3],
					},
					Tokens: tk[2:3],
				},
				Modules: []ModuleAs{
					{
						Module: Module{
							Identifiers: []IdentifierComments{
								{
									Identifier: &tk[7],
								},
							},
							Tokens: tk[7:8],
						},
						Tokens: tk[7:8],
					},
					{
						Module: Module{
							Identifiers: []IdentifierComments{
								{
									Identifier: &tk[9],
								},
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
						Identifiers: []IdentifierComments{
							{
								Identifier: &tk[2],
							},
						},
						Tokens: tk[2:3],
					},
					Tokens: tk[2:3],
				},
				Modules: []ModuleAs{
					{
						Module: Module{
							Identifiers: []IdentifierComments{
								{
									Identifier: &tk[7],
								},
							},
							Tokens: tk[7:8],
						},
						Tokens: tk[7:8],
					},
					{
						Module: Module{
							Identifiers: []IdentifierComments{
								{
									Identifier: &tk[10],
								},
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
							Identifiers: []IdentifierComments{
								{
									Identifier: &tk[2],
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
		{`import a,b`, func(t *test, tk Tokens) { // 9
			t.Output = ImportStatement{
				Modules: []ModuleAs{
					{
						Module: Module{
							Identifiers: []IdentifierComments{
								{
									Identifier: &tk[2],
								},
							},
							Tokens: tk[2:3],
						},
						Tokens: tk[2:3],
					},
					{
						Module: Module{
							Identifiers: []IdentifierComments{
								{
									Identifier: &tk[4],
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
		{`import a, b, c`, func(t *test, tk Tokens) { // 10
			t.Output = ImportStatement{
				Modules: []ModuleAs{
					{
						Module: Module{
							Identifiers: []IdentifierComments{
								{
									Identifier: &tk[2],
								},
							},
							Tokens: tk[2:3],
						},
						Tokens: tk[2:3],
					},
					{
						Module: Module{
							Identifiers: []IdentifierComments{
								{
									Identifier: &tk[5],
								},
							},
							Tokens: tk[5:6],
						},
						Tokens: tk[5:6],
					},
					{
						Module: Module{
							Identifiers: []IdentifierComments{
								{
									Identifier: &tk[8],
								},
							},
							Tokens: tk[8:9],
						},
						Tokens: tk[8:9],
					},
				},
				Tokens: tk[:9],
			}
		}},
		{"from a import ( # A\nb,c\n# B\n)", func(t *test, tk Tokens) { // 11
			t.Output = ImportStatement{
				RelativeModule: &RelativeModule{
					Module: &Module{
						Identifiers: []IdentifierComments{
							{
								Identifier: &tk[2],
							},
						},
						Tokens: tk[2:3],
					},
					Tokens: tk[2:3],
				},
				Modules: []ModuleAs{
					{
						Module: Module{
							Identifiers: []IdentifierComments{
								{
									Identifier: &tk[10],
								},
							},
							Tokens: tk[10:11],
						},
						Tokens: tk[10:11],
					},
					{
						Module: Module{
							Identifiers: []IdentifierComments{
								{
									Identifier: &tk[12],
								},
							},
							Tokens: tk[12:13],
						},
						Tokens: tk[12:13],
					},
				},
				Comments: [2]Comments{{tk[8]}, {tk[14]}},
				Tokens:   tk[:17],
			}
		}},
		{`from nonlocal import a`, func(t *test, tk Tokens) { // 12
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
		{`from a b`, func(t *test, tk Tokens) { // 13
			t.Err = Error{
				Err:     ErrMissingImport,
				Parsing: "ImportStatement",
				Token:   tk[4],
			}
		}},
		{`from a import nonlocal`, func(t *test, tk Tokens) { // 14
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
		{`from a import (b c)`, func(t *test, tk Tokens) { // 15
			t.Err = Error{
				Err:     ErrMissingComma,
				Parsing: "ImportStatement",
				Token:   tk[8],
			}
		}},
		{`import nonlocal`, func(t *test, tk Tokens) { // 16
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
		{`import (b)`, func(t *test, tk Tokens) { // 17
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
					Identifiers: []IdentifierComments{
						{
							Identifier: &tk[0],
						},
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
					Identifiers: []IdentifierComments{
						{
							Identifier: &tk[1],
						},
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
					Identifiers: []IdentifierComments{
						{
							Identifier: &tk[2],
						},
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
					Identifiers: []IdentifierComments{
						{
							Identifier: &tk[4],
						},
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
					Identifiers: []IdentifierComments{
						{
							Identifier: &tk[1],
						},
						{
							Identifier: &tk[3],
						},
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
					Identifiers: []IdentifierComments{
						{
							Identifier: &tk[0],
						},
					},
					Tokens: tk[:1],
				},
				Tokens: tk[:1],
			}
		}},
		{`a as b`, func(t *test, tk Tokens) { // 2
			t.Output = ModuleAs{
				Module: Module{
					Identifiers: []IdentifierComments{
						{
							Identifier: &tk[0],
						},
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
				Identifiers: []IdentifierComments{
					{
						Identifier: &tk[0],
					},
				},
				Tokens: tk[:1],
			}
		}},
		{`a.b`, func(t *test, tk Tokens) { // 2
			t.Output = Module{
				Identifiers: []IdentifierComments{
					{
						Identifier: &tk[0],
					},
					{
						Identifier: &tk[2],
					},
				},
				Tokens: tk[:3],
			}
		}},
		{`a . b . c`, func(t *test, tk Tokens) { // 3
			t.Output = Module{
				Identifiers: []IdentifierComments{
					{
						Identifier: &tk[0],
					},
					{
						Identifier: &tk[4],
					},
					{
						Identifier: &tk[8],
					},
				},
				Tokens: tk[:9],
			}
		}},
		{"(a # A\n. b # B\n. c # C\n)", func(t *test, tk Tokens) { // 3
			t.Output = Module{
				Identifiers: []IdentifierComments{
					{
						Identifier: &tk[1],
						Comments:   Comments{tk[3]},
					},
					{
						Identifier: &tk[7],
						Comments:   Comments{tk[9]},
					},
					{
						Identifier: &tk[13],
						Comments:   Comments{tk[15]},
					},
				},
				Tokens: tk[1:16],
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

		if t.Tokens.Peek().Data == "(" {
			t.Tokens.Tokens = t.Tokens.Tokens[1:1]
		}

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

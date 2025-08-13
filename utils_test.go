package python

import (
	"errors"
	"reflect"
	"strconv"
	"testing"

	"vimagination.zapto.org/parser"
)

func TestUnquote(t *testing.T) {
	for n, test := range [...]struct {
		Input, Output string
		Err           error
	}{
		{ // 1
			Input:  "\"abc\"",
			Output: "abc",
		},
		{ // 2
			Input:  "\"ab\\\"c\"",
			Output: "ab\"c",
		},
		{ // 3
			Input:  "'ab\\\"c'",
			Output: "ab\"c",
		},
		{ // 4
			Input:  "'ab\\'c'",
			Output: "ab'c",
		},
		{ // 5
			Input: "\"ab\nc\"",
			Err:   strconv.ErrSyntax,
		},
		{ // 6
			Input:  "\"\"\"ab\nc\"\"\"",
			Output: "ab\nc",
		},
		{ // 7
			Input: "'ab\nc'",
			Err:   strconv.ErrSyntax,
		},
		{ // 8
			Input:  "'''ab\nc'''",
			Output: "ab\nc",
		},
		{ // 9
			Input: "\"abc\\\"",
			Err:   strconv.ErrSyntax,
		},
		{ // 10
			Input:  "r\"abc\\\"",
			Output: "abc\\",
		},
		{ // 11
			Input:  "R'abc\\'",
			Output: "abc\\",
		},
		{ // 12
			Input:  "\"\"",
			Output: "",
		},
		{ // 13
			Input:  "''",
			Output: "",
		},
		{ // 14
			Input: "\"\\09\"",
			Err:   strconv.ErrSyntax,
		},
		{ // 15
			Input:  "\"\\101\"",
			Output: "A",
		},
		{ // 16
			Input:  "\"\\a\"",
			Output: "\a",
		},
		{ // 17
			Input:  "\"\\b\"",
			Output: "\b",
		},
		{ // 18
			Input:  "\"\\f\"",
			Output: "\f",
		},
		{ // 19
			Input:  "\"\\n\"",
			Output: "\n",
		},
		{ // 20
			Input:  "\"\\r\"",
			Output: "\r",
		},
		{ // 21
			Input:  "\"\\t\"",
			Output: "\t",
		},
		{ // 22
			Input:  "\"\\v\"",
			Output: "\v",
		},
		{ // 23
			Input:  "\"\\x41\"",
			Output: "A",
		},
		{ // 24
			Input: "\"\\N\"",
			Err:   strconv.ErrSyntax,
		},
		{ // 25
			Input:  "\"\\u0041\"",
			Output: "A",
		},
		{ // 26
			Input:  "\"\\U00000041\"",
			Output: "A",
		},
		{ // 27
			Input: "\"\\B\"",
			Err:   strconv.ErrSyntax,
		},
		{ // 28
			Input: "B",
			Err:   strconv.ErrSyntax,
		},
		{ // 29
			Input:  "\"a\\\nb\"",
			Output: "ab",
		},
	} {
		output, err := Unquote(test.Input)

		if !errors.Is(test.Err, err) {
			t.Errorf("test %d: expecting error %q, got %q", n+1, test.Err, err)
		} else if test.Output != output {
			t.Errorf("test %d: expecting output %q, got %q", n+1, test.Output, output)
		}
	}
}

func TestWrapConditional(t *testing.T) {
	tks := Tokens{
		{
			Token: parser.Token{
				Type: TokenIdentifier,
				Data: "a",
			},
		},
	}
	ident := &tks[0]
	expectedOutput := ConditionalExpression{
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
																Identifier: ident,
																Tokens:     tks,
															},
															Tokens: tks,
														},
														Tokens: tks,
													},
													Tokens: tks,
												},
												Tokens: tks,
											},
											Tokens: tks,
										},
										Tokens: tks,
									},
									Tokens: tks,
								},
								Tokens: tks,
							},
							Tokens: tks,
						},
						Tokens: tks,
					},
					Tokens: tks,
				},
				Tokens: tks,
			},
			Tokens: tks,
		},
		Tokens: tks,
	}

	for n, test := range [...]ConditionalWrappable{
		&Atom{ // 1
			Identifier: ident,
			Tokens:     tks,
		},
		Atom{ // 2
			Identifier: ident,
			Tokens:     tks,
		},
		&PrimaryExpression{ // 3
			Atom: &Atom{
				Identifier: ident,
				Tokens:     tks,
			},
			Tokens: tks,
		},
		PrimaryExpression{ // 4
			Atom: &Atom{
				Identifier: ident,
				Tokens:     tks,
			},
			Tokens: tks,
		},
		&PowerExpression{ // 5
			PrimaryExpression: PrimaryExpression{
				Atom: &Atom{
					Identifier: ident,
					Tokens:     tks,
				},
				Tokens: tks,
			},
			Tokens: tks,
		},
		PowerExpression{ // 6
			PrimaryExpression: PrimaryExpression{
				Atom: &Atom{
					Identifier: ident,
					Tokens:     tks,
				},
				Tokens: tks,
			},
			Tokens: tks,
		},
		&UnaryExpression{ // 7
			PowerExpression: &PowerExpression{
				PrimaryExpression: PrimaryExpression{
					Atom: &Atom{
						Identifier: ident,
						Tokens:     tks,
					},
					Tokens: tks,
				},
				Tokens: tks,
			},
			Tokens: tks,
		},
		UnaryExpression{ // 8
			PowerExpression: &PowerExpression{
				PrimaryExpression: PrimaryExpression{
					Atom: &Atom{
						Identifier: ident,
						Tokens:     tks,
					},
					Tokens: tks,
				},
				Tokens: tks,
			},
			Tokens: tks,
		},
		&MultiplyExpression{ // 9
			UnaryExpression: UnaryExpression{
				PowerExpression: &PowerExpression{
					PrimaryExpression: PrimaryExpression{
						Atom: &Atom{
							Identifier: ident,
							Tokens:     tks,
						},
						Tokens: tks,
					},
					Tokens: tks,
				},
				Tokens: tks,
			},
			Tokens: tks,
		},
		MultiplyExpression{ // 10
			UnaryExpression: UnaryExpression{
				PowerExpression: &PowerExpression{
					PrimaryExpression: PrimaryExpression{
						Atom: &Atom{
							Identifier: ident,
							Tokens:     tks,
						},
						Tokens: tks,
					},
					Tokens: tks,
				},
				Tokens: tks,
			},
			Tokens: tks,
		},
		&AddExpression{ // 11
			MultiplyExpression: MultiplyExpression{
				UnaryExpression: UnaryExpression{
					PowerExpression: &PowerExpression{
						PrimaryExpression: PrimaryExpression{
							Atom: &Atom{
								Identifier: ident,
								Tokens:     tks,
							},
							Tokens: tks,
						},
						Tokens: tks,
					},
					Tokens: tks,
				},
				Tokens: tks,
			},
			Tokens: tks,
		},
		AddExpression{ // 12
			MultiplyExpression: MultiplyExpression{
				UnaryExpression: UnaryExpression{
					PowerExpression: &PowerExpression{
						PrimaryExpression: PrimaryExpression{
							Atom: &Atom{
								Identifier: ident,
								Tokens:     tks,
							},
							Tokens: tks,
						},
						Tokens: tks,
					},
					Tokens: tks,
				},
				Tokens: tks,
			},
			Tokens: tks,
		},
		&ShiftExpression{ // 13
			AddExpression: AddExpression{
				MultiplyExpression: MultiplyExpression{
					UnaryExpression: UnaryExpression{
						PowerExpression: &PowerExpression{
							PrimaryExpression: PrimaryExpression{
								Atom: &Atom{
									Identifier: ident,
									Tokens:     tks,
								},
								Tokens: tks,
							},
							Tokens: tks,
						},
						Tokens: tks,
					},
					Tokens: tks,
				},
				Tokens: tks,
			},
			Tokens: tks,
		},
		ShiftExpression{ // 14
			AddExpression: AddExpression{
				MultiplyExpression: MultiplyExpression{
					UnaryExpression: UnaryExpression{
						PowerExpression: &PowerExpression{
							PrimaryExpression: PrimaryExpression{
								Atom: &Atom{
									Identifier: ident,
									Tokens:     tks,
								},
								Tokens: tks,
							},
							Tokens: tks,
						},
						Tokens: tks,
					},
					Tokens: tks,
				},
				Tokens: tks,
			},
			Tokens: tks,
		},
		&AndExpression{ // 15
			ShiftExpression: ShiftExpression{
				AddExpression: AddExpression{
					MultiplyExpression: MultiplyExpression{
						UnaryExpression: UnaryExpression{
							PowerExpression: &PowerExpression{
								PrimaryExpression: PrimaryExpression{
									Atom: &Atom{
										Identifier: ident,
										Tokens:     tks,
									},
									Tokens: tks,
								},
								Tokens: tks,
							},
							Tokens: tks,
						},
						Tokens: tks,
					},
					Tokens: tks,
				},
				Tokens: tks,
			},
			Tokens: tks,
		},
		AndExpression{ // 16
			ShiftExpression: ShiftExpression{
				AddExpression: AddExpression{
					MultiplyExpression: MultiplyExpression{
						UnaryExpression: UnaryExpression{
							PowerExpression: &PowerExpression{
								PrimaryExpression: PrimaryExpression{
									Atom: &Atom{
										Identifier: ident,
										Tokens:     tks,
									},
									Tokens: tks,
								},
								Tokens: tks,
							},
							Tokens: tks,
						},
						Tokens: tks,
					},
					Tokens: tks,
				},
				Tokens: tks,
			},
			Tokens: tks,
		},
		&XorExpression{ // 17
			AndExpression: AndExpression{
				ShiftExpression: ShiftExpression{
					AddExpression: AddExpression{
						MultiplyExpression: MultiplyExpression{
							UnaryExpression: UnaryExpression{
								PowerExpression: &PowerExpression{
									PrimaryExpression: PrimaryExpression{
										Atom: &Atom{
											Identifier: ident,
											Tokens:     tks,
										},
										Tokens: tks,
									},
									Tokens: tks,
								},
								Tokens: tks,
							},
							Tokens: tks,
						},
						Tokens: tks,
					},
					Tokens: tks,
				},
				Tokens: tks,
			},
			Tokens: tks,
		},
		XorExpression{ // 18
			AndExpression: AndExpression{
				ShiftExpression: ShiftExpression{
					AddExpression: AddExpression{
						MultiplyExpression: MultiplyExpression{
							UnaryExpression: UnaryExpression{
								PowerExpression: &PowerExpression{
									PrimaryExpression: PrimaryExpression{
										Atom: &Atom{
											Identifier: ident,
											Tokens:     tks,
										},
										Tokens: tks,
									},
									Tokens: tks,
								},
								Tokens: tks,
							},
							Tokens: tks,
						},
						Tokens: tks,
					},
					Tokens: tks,
				},
				Tokens: tks,
			},
			Tokens: tks,
		},
		&OrExpression{ // 19
			XorExpression: XorExpression{
				AndExpression: AndExpression{
					ShiftExpression: ShiftExpression{
						AddExpression: AddExpression{
							MultiplyExpression: MultiplyExpression{
								UnaryExpression: UnaryExpression{
									PowerExpression: &PowerExpression{
										PrimaryExpression: PrimaryExpression{
											Atom: &Atom{
												Identifier: ident,
												Tokens:     tks,
											},
											Tokens: tks,
										},
										Tokens: tks,
									},
									Tokens: tks,
								},
								Tokens: tks,
							},
							Tokens: tks,
						},
						Tokens: tks,
					},
					Tokens: tks,
				},
				Tokens: tks,
			},
			Tokens: tks,
		},
		OrExpression{ // 20
			XorExpression: XorExpression{
				AndExpression: AndExpression{
					ShiftExpression: ShiftExpression{
						AddExpression: AddExpression{
							MultiplyExpression: MultiplyExpression{
								UnaryExpression: UnaryExpression{
									PowerExpression: &PowerExpression{
										PrimaryExpression: PrimaryExpression{
											Atom: &Atom{
												Identifier: ident,
												Tokens:     tks,
											},
											Tokens: tks,
										},
										Tokens: tks,
									},
									Tokens: tks,
								},
								Tokens: tks,
							},
							Tokens: tks,
						},
						Tokens: tks,
					},
					Tokens: tks,
				},
				Tokens: tks,
			},
			Tokens: tks,
		},
		&Comparison{ // 21
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
													Identifier: ident,
													Tokens:     tks,
												},
												Tokens: tks,
											},
											Tokens: tks,
										},
										Tokens: tks,
									},
									Tokens: tks,
								},
								Tokens: tks,
							},
							Tokens: tks,
						},
						Tokens: tks,
					},
					Tokens: tks,
				},
				Tokens: tks,
			},
			Tokens: tks,
		},
		Comparison{ // 22
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
													Identifier: ident,
													Tokens:     tks,
												},
												Tokens: tks,
											},
											Tokens: tks,
										},
										Tokens: tks,
									},
									Tokens: tks,
								},
								Tokens: tks,
							},
							Tokens: tks,
						},
						Tokens: tks,
					},
					Tokens: tks,
				},
				Tokens: tks,
			},
			Tokens: tks,
		},
		&NotTest{ // 23
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
														Identifier: ident,
														Tokens:     tks,
													},
													Tokens: tks,
												},
												Tokens: tks,
											},
											Tokens: tks,
										},
										Tokens: tks,
									},
									Tokens: tks,
								},
								Tokens: tks,
							},
							Tokens: tks,
						},
						Tokens: tks,
					},
					Tokens: tks,
				},
				Tokens: tks,
			},
			Tokens: tks,
		},
		NotTest{ // 24
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
														Identifier: ident,
														Tokens:     tks,
													},
													Tokens: tks,
												},
												Tokens: tks,
											},
											Tokens: tks,
										},
										Tokens: tks,
									},
									Tokens: tks,
								},
								Tokens: tks,
							},
							Tokens: tks,
						},
						Tokens: tks,
					},
					Tokens: tks,
				},
				Tokens: tks,
			},
			Tokens: tks,
		},
		&AndTest{ // 25
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
															Identifier: ident,
															Tokens:     tks,
														},
														Tokens: tks,
													},
													Tokens: tks,
												},
												Tokens: tks,
											},
											Tokens: tks,
										},
										Tokens: tks,
									},
									Tokens: tks,
								},
								Tokens: tks,
							},
							Tokens: tks,
						},
						Tokens: tks,
					},
					Tokens: tks,
				},
				Tokens: tks,
			},
			Tokens: tks,
		},
		AndTest{ // 26
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
															Identifier: ident,
															Tokens:     tks,
														},
														Tokens: tks,
													},
													Tokens: tks,
												},
												Tokens: tks,
											},
											Tokens: tks,
										},
										Tokens: tks,
									},
									Tokens: tks,
								},
								Tokens: tks,
							},
							Tokens: tks,
						},
						Tokens: tks,
					},
					Tokens: tks,
				},
				Tokens: tks,
			},
			Tokens: tks,
		},
		&OrTest{ // 27
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
																Identifier: ident,
																Tokens:     tks,
															},
															Tokens: tks,
														},
														Tokens: tks,
													},
													Tokens: tks,
												},
												Tokens: tks,
											},
											Tokens: tks,
										},
										Tokens: tks,
									},
									Tokens: tks,
								},
								Tokens: tks,
							},
							Tokens: tks,
						},
						Tokens: tks,
					},
					Tokens: tks,
				},
				Tokens: tks,
			},
			Tokens: tks,
		},
		OrTest{ // 28
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
																Identifier: ident,
																Tokens:     tks,
															},
															Tokens: tks,
														},
														Tokens: tks,
													},
													Tokens: tks,
												},
												Tokens: tks,
											},
											Tokens: tks,
										},
										Tokens: tks,
									},
									Tokens: tks,
								},
								Tokens: tks,
							},
							Tokens: tks,
						},
						Tokens: tks,
					},
					Tokens: tks,
				},
				Tokens: tks,
			},
			Tokens: tks,
		},
		&ConditionalExpression{ // 29
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
																	Identifier: ident,
																	Tokens:     tks,
																},
																Tokens: tks,
															},
															Tokens: tks,
														},
														Tokens: tks,
													},
													Tokens: tks,
												},
												Tokens: tks,
											},
											Tokens: tks,
										},
										Tokens: tks,
									},
									Tokens: tks,
								},
								Tokens: tks,
							},
							Tokens: tks,
						},
						Tokens: tks,
					},
					Tokens: tks,
				},
				Tokens: tks,
			},
			Tokens: tks,
		},
		ConditionalExpression{ // 30
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
																	Identifier: ident,
																	Tokens:     tks,
																},
																Tokens: tks,
															},
															Tokens: tks,
														},
														Tokens: tks,
													},
													Tokens: tks,
												},
												Tokens: tks,
											},
											Tokens: tks,
										},
										Tokens: tks,
									},
									Tokens: tks,
								},
								Tokens: tks,
							},
							Tokens: tks,
						},
						Tokens: tks,
					},
					Tokens: tks,
				},
				Tokens: tks,
			},
			Tokens: tks,
		},
	} {
		if output := WrapConditional(test); !reflect.DeepEqual(output, &expectedOutput) {
			t.Errorf("test %d: expecting\n%v\n...got...\n%v", n+1, expectedOutput, output)
		}
	}
}

func TestUnwrapConditional(t *testing.T) {
	tks := Tokens{
		{
			Token: parser.Token{
				Type: TokenIdentifier,
				Data: "a",
			},
		},
		{
			Token: parser.Token{
				Type: TokenIdentifier,
				Data: "b",
			},
		},
	}
	identA := &tks[0]
	identB := &tks[1]

	for n, test := range [...]ConditionalWrappable{
		&Atom{ // 1
			Identifier: identA,
			Tokens:     tks[:1],
		},
		&PrimaryExpression{ // 2
			PrimaryExpression: &PrimaryExpression{
				Atom: &Atom{
					Identifier: identA,
					Tokens:     tks[:1],
				},
			},
			AttributeRef: identB,
			Tokens:       tks[:2],
		},
		&PowerExpression{ // 3
			PrimaryExpression: PrimaryExpression{
				Atom: &Atom{
					Identifier: identA,
					Tokens:     tks[:1],
				},
				Tokens: tks[:1],
			},
			UnaryExpression: &UnaryExpression{
				PowerExpression: &PowerExpression{
					PrimaryExpression: PrimaryExpression{
						Atom: &Atom{
							Identifier: identB,
							Tokens:     tks[1:2],
						},
						Tokens: tks[1:2],
					},
					Tokens: tks[1:2],
				},
				Tokens: tks[1:2],
			},
			Tokens: tks[:2],
		},
		&UnaryExpression{ // 4
			UnaryExpression: &UnaryExpression{
				PowerExpression: &PowerExpression{
					PrimaryExpression: PrimaryExpression{
						Atom: &Atom{
							Identifier: identB,
							Tokens:     tks[1:2],
						},
						Tokens: tks[1:2],
					},
					Tokens: tks[1:2],
				},
				Tokens: tks[1:2],
			},
			Tokens: tks[1:2],
		},
		&MultiplyExpression{ // 5
			UnaryExpression: UnaryExpression{
				PowerExpression: &PowerExpression{
					PrimaryExpression: PrimaryExpression{
						Atom: &Atom{
							Identifier: identA,
							Tokens:     tks[:1],
						},
						Tokens: tks[:1],
					},
					Tokens: tks[:1],
				},
				Tokens: tks[:1],
			},
			MultiplyExpression: &MultiplyExpression{
				UnaryExpression: UnaryExpression{
					PowerExpression: &PowerExpression{
						PrimaryExpression: PrimaryExpression{
							Atom: &Atom{
								Identifier: identB,
								Tokens:     tks[1:2],
							},
							Tokens: tks[1:2],
						},
						Tokens: tks[1:2],
					},
					Tokens: tks[1:2],
				},
				Tokens: tks[1:2],
			},
			Tokens: tks[:2],
		},
		&AddExpression{ // 6
			MultiplyExpression: MultiplyExpression{
				UnaryExpression: UnaryExpression{
					PowerExpression: &PowerExpression{
						PrimaryExpression: PrimaryExpression{
							Atom: &Atom{
								Identifier: identA,
								Tokens:     tks[:1],
							},
							Tokens: tks[:1],
						},
						Tokens: tks[:1],
					},
					Tokens: tks[:1],
				},
				Tokens: tks[:1],
			},
			AddExpression: &AddExpression{
				MultiplyExpression: MultiplyExpression{
					UnaryExpression: UnaryExpression{
						PowerExpression: &PowerExpression{
							PrimaryExpression: PrimaryExpression{
								PrimaryExpression: &PrimaryExpression{
									Atom: &Atom{
										Identifier: identB,
										Tokens:     tks[1:2],
									},
									Tokens: tks[1:2],
								},
								AttributeRef: identB,
								Tokens:       tks[1:2],
							},
							Tokens: tks[1:2],
						},
						Tokens: tks[1:2],
					},
					Tokens: tks[1:2],
				},
				Tokens: tks[1:2],
			},
			Tokens: tks[:2],
		},
		&ShiftExpression{ // 7
			AddExpression: AddExpression{
				MultiplyExpression: MultiplyExpression{
					UnaryExpression: UnaryExpression{
						PowerExpression: &PowerExpression{
							PrimaryExpression: PrimaryExpression{
								Atom: &Atom{
									Identifier: identA,
									Tokens:     tks[:1],
								},
								Tokens: tks[:1],
							},
							Tokens: tks[:1],
						},
						Tokens: tks[:1],
					},
					Tokens: tks[:1],
				},
				Tokens: tks[:1],
			},
			ShiftExpression: &ShiftExpression{
				AddExpression: AddExpression{
					MultiplyExpression: MultiplyExpression{
						UnaryExpression: UnaryExpression{
							PowerExpression: &PowerExpression{
								PrimaryExpression: PrimaryExpression{
									PrimaryExpression: &PrimaryExpression{
										Atom: &Atom{
											Identifier: identB,
											Tokens:     tks[1:2],
										},
										Tokens: tks[1:2],
									},
									AttributeRef: identB,
									Tokens:       tks[1:2],
								},
								Tokens: tks[1:2],
							},
							Tokens: tks[1:2],
						},
						Tokens: tks[1:2],
					},
					Tokens: tks[1:2],
				},
				Tokens: tks[1:2],
			},
			Tokens: tks[:2],
		},
		&AndExpression{ // 8
			ShiftExpression: ShiftExpression{
				AddExpression: AddExpression{
					MultiplyExpression: MultiplyExpression{
						UnaryExpression: UnaryExpression{
							PowerExpression: &PowerExpression{
								PrimaryExpression: PrimaryExpression{
									Atom: &Atom{
										Identifier: identA,
										Tokens:     tks[:1],
									},
									Tokens: tks[:1],
								},
								Tokens: tks[:1],
							},
							Tokens: tks[:1],
						},
						Tokens: tks[:1],
					},
					Tokens: tks[:1],
				},
				Tokens: tks[:1],
			},
			AndExpression: &AndExpression{
				ShiftExpression: ShiftExpression{
					AddExpression: AddExpression{
						MultiplyExpression: MultiplyExpression{
							UnaryExpression: UnaryExpression{
								PowerExpression: &PowerExpression{
									PrimaryExpression: PrimaryExpression{
										PrimaryExpression: &PrimaryExpression{
											Atom: &Atom{
												Identifier: identB,
												Tokens:     tks[1:2],
											},
											Tokens: tks[1:2],
										},
										AttributeRef: identB,
										Tokens:       tks[1:2],
									},
									Tokens: tks[1:2],
								},
								Tokens: tks[1:2],
							},
							Tokens: tks[1:2],
						},
						Tokens: tks[1:2],
					},
					Tokens: tks[1:2],
				},
				Tokens: tks[1:2],
			},
			Tokens: tks[:2],
		},
		&XorExpression{ // 9
			AndExpression: AndExpression{
				ShiftExpression: ShiftExpression{
					AddExpression: AddExpression{
						MultiplyExpression: MultiplyExpression{
							UnaryExpression: UnaryExpression{
								PowerExpression: &PowerExpression{
									PrimaryExpression: PrimaryExpression{
										Atom: &Atom{
											Identifier: identA,
											Tokens:     tks[:1],
										},
										Tokens: tks[:1],
									},
									Tokens: tks[:1],
								},
								Tokens: tks[:1],
							},
							Tokens: tks[:1],
						},
						Tokens: tks[:1],
					},
					Tokens: tks[:1],
				},
				Tokens: tks[:1],
			},
			XorExpression: &XorExpression{
				AndExpression: AndExpression{
					ShiftExpression: ShiftExpression{
						AddExpression: AddExpression{
							MultiplyExpression: MultiplyExpression{
								UnaryExpression: UnaryExpression{
									PowerExpression: &PowerExpression{
										PrimaryExpression: PrimaryExpression{
											PrimaryExpression: &PrimaryExpression{
												Atom: &Atom{
													Identifier: identB,
													Tokens:     tks[1:2],
												},
												Tokens: tks[1:2],
											},
											AttributeRef: identB,
											Tokens:       tks[1:2],
										},
										Tokens: tks[1:2],
									},
									Tokens: tks[1:2],
								},
								Tokens: tks[1:2],
							},
							Tokens: tks[1:2],
						},
						Tokens: tks[1:2],
					},
					Tokens: tks[1:2],
				},
				Tokens: tks[1:2],
			},
			Tokens: tks[:2],
		},
		&OrExpression{ // 10
			XorExpression: XorExpression{
				AndExpression: AndExpression{
					ShiftExpression: ShiftExpression{
						AddExpression: AddExpression{
							MultiplyExpression: MultiplyExpression{
								UnaryExpression: UnaryExpression{
									PowerExpression: &PowerExpression{
										PrimaryExpression: PrimaryExpression{
											Atom: &Atom{
												Identifier: identA,
												Tokens:     tks[:1],
											},
											Tokens: tks[:1],
										},
										Tokens: tks[:1],
									},
									Tokens: tks[:1],
								},
								Tokens: tks[:1],
							},
							Tokens: tks[:1],
						},
						Tokens: tks[:1],
					},
					Tokens: tks[:1],
				},
				Tokens: tks[:1],
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
												PrimaryExpression: &PrimaryExpression{
													Atom: &Atom{
														Identifier: identB,
														Tokens:     tks[1:2],
													},
													Tokens: tks[1:2],
												},
												AttributeRef: identB,
												Tokens:       tks[1:2],
											},
											Tokens: tks[1:2],
										},
										Tokens: tks[1:2],
									},
									Tokens: tks[1:2],
								},
								Tokens: tks[1:2],
							},
							Tokens: tks[1:2],
						},
						Tokens: tks[1:2],
					},
					Tokens: tks[1:2],
				},
				Tokens: tks[1:2],
			},
			Tokens: tks[:2],
		},
		&Comparison{ // 11
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
													Identifier: identA,
													Tokens:     tks[:1],
												},
												Tokens: tks[:1],
											},
											Tokens: tks[:1],
										},
										Tokens: tks[:1],
									},
									Tokens: tks[:1],
								},
								Tokens: tks[:1],
							},
							Tokens: tks[:1],
						},
						Tokens: tks[:1],
					},
					Tokens: tks[:1],
				},
				Tokens: tks[:1],
			},
			Comparisons: []ComparisonExpression{{}},
		},
		&NotTest{ // 12
			Nots: make([]Comments, 1),
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
														Identifier: identA,
														Tokens:     tks[:1],
													},
													Tokens: tks[:1],
												},
												Tokens: tks[:1],
											},
											Tokens: tks[:1],
										},
										Tokens: tks[:1],
									},
									Tokens: tks[:1],
								},
								Tokens: tks[:1],
							},
							Tokens: tks[:1],
						},
						Tokens: tks[:1],
					},
					Tokens: tks[:1],
				},
				Tokens: tks[:1],
			},
			Tokens: tks[:1],
		},
		&AndTest{ // 13
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
															Identifier: identA,
															Tokens:     tks[:1],
														},
														Tokens: tks[:1],
													},
													Tokens: tks[:1],
												},
												Tokens: tks[:1],
											},
											Tokens: tks[:1],
										},
										Tokens: tks[:1],
									},
									Tokens: tks[:1],
								},
								Tokens: tks[:1],
							},
							Tokens: tks[:1],
						},
						Tokens: tks[:1],
					},
					Tokens: tks[:1],
				},
				Tokens: tks[:1],
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
																Identifier: identB,
																Tokens:     tks[1:2],
															},
															Tokens: tks[1:2],
														},
														Tokens: tks[1:2],
													},
													Tokens: tks[1:2],
												},
												Tokens: tks[1:2],
											},
											Tokens: tks[1:2],
										},
										Tokens: tks[1:2],
									},
									Tokens: tks[1:2],
								},
								Tokens: tks[1:2],
							},
							Tokens: tks[1:2],
						},
						Tokens: tks[1:2],
					},
					Tokens: tks[1:2],
				},
				Tokens: tks[1:2],
			},
			Tokens: tks[:2],
		},
		&OrTest{ // 14
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
																Identifier: identA,
																Tokens:     tks[:1],
															},
															Tokens: tks[:1],
														},
														Tokens: tks[:1],
													},
													Tokens: tks[:1],
												},
												Tokens: tks[:1],
											},
											Tokens: tks[:1],
										},
										Tokens: tks[:1],
									},
									Tokens: tks[:1],
								},
								Tokens: tks[:1],
							},
							Tokens: tks[:1],
						},
						Tokens: tks[:1],
					},
					Tokens: tks[:1],
				},
				Tokens: tks[:1],
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
																	Identifier: identB,
																	Tokens:     tks[1:2],
																},
																Tokens: tks[1:2],
															},
															Tokens: tks[1:2],
														},
														Tokens: tks[1:2],
													},
													Tokens: tks[1:2],
												},
												Tokens: tks[1:2],
											},
											Tokens: tks[1:2],
										},
										Tokens: tks[1:2],
									},
									Tokens: tks[1:2],
								},
								Tokens: tks[1:2],
							},
							Tokens: tks[1:2],
						},
						Tokens: tks[1:2],
					},
					Tokens: tks[1:2],
				},
				Tokens: tks[1:2],
			},
			Tokens: tks[:2],
		},
		&ConditionalExpression{ // 15
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
																	Identifier: identA,
																	Tokens:     tks[:1],
																},
																Tokens: tks[:1],
															},
															Tokens: tks[:1],
														},
														Tokens: tks[:1],
													},
													Tokens: tks[:1],
												},
												Tokens: tks[:1],
											},
											Tokens: tks[:1],
										},
										Tokens: tks[:1],
									},
									Tokens: tks[:1],
								},
								Tokens: tks[:1],
							},
							Tokens: tks[:1],
						},
						Tokens: tks[:1],
					},
					Tokens: tks[:1],
				},
				Tokens: tks[:1],
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
																	Identifier: identB,
																	Tokens:     tks[1:2],
																},
																Tokens: tks[1:2],
															},
															Tokens: tks[1:2],
														},
														Tokens: tks[1:2],
													},
													Tokens: tks[1:2],
												},
												Tokens: tks[1:2],
											},
											Tokens: tks[1:2],
										},
										Tokens: tks[1:2],
									},
									Tokens: tks[1:2],
								},
								Tokens: tks[1:2],
							},
							Tokens: tks[1:2],
						},
						Tokens: tks[1:2],
					},
					Tokens: tks[1:2],
				},
				Tokens: tks[1:2],
			},
			Tokens: tks[:2],
		},
	} {
		if output := UnwrapConditional(WrapConditional(test)); !reflect.DeepEqual(output, test) {
			t.Errorf("test %d: expecting\n%v\n...got...\n%v", n+1, test, output)
		}
	}
}

func TestUnwrapConditionalExtra(t *testing.T) {
	if res := UnwrapConditional(nil); res != nil {
		t.Errorf("test 1: expecting nil, got %v", res)
	}

	if res := UnwrapConditional(WrapConditional(&UnaryExpression{})); res != nil {
		t.Errorf("test 2: expecting nil, got %v", res)
	}
}

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
		{
			Input:  "\"abc\"",
			Output: "abc",
		},
		{
			Input:  "\"ab\\\"c\"",
			Output: "ab\"c",
		},
		{
			Input:  "'ab\\\"c'",
			Output: "ab\"c",
		},
		{
			Input:  "'ab\\'c'",
			Output: "ab'c",
		},
		{
			Input: "\"ab\nc\"",
			Err:   strconv.ErrSyntax,
		},
		{
			Input:  "\"\"\"ab\nc\"\"\"",
			Output: "ab\nc",
		},
		{
			Input: "'ab\nc'",
			Err:   strconv.ErrSyntax,
		},
		{
			Input:  "'''ab\nc'''",
			Output: "ab\nc",
		},
		{
			Input: "\"abc\\\"",
			Err:   strconv.ErrSyntax,
		},
		{
			Input:  "r\"abc\\\"",
			Output: "abc\\",
		},
		{
			Input:  "R'abc\\'",
			Output: "abc\\",
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
		&Atom{
			Identifier: ident,
			Tokens:     tks,
		},
		Atom{
			Identifier: ident,
			Tokens:     tks,
		},
		&PrimaryExpression{
			Atom: &Atom{
				Identifier: ident,
				Tokens:     tks,
			},
			Tokens: tks,
		},
		PrimaryExpression{
			Atom: &Atom{
				Identifier: ident,
				Tokens:     tks,
			},
			Tokens: tks,
		},
		&PowerExpression{
			PrimaryExpression: PrimaryExpression{
				Atom: &Atom{
					Identifier: ident,
					Tokens:     tks,
				},
				Tokens: tks,
			},
			Tokens: tks,
		},
		PowerExpression{
			PrimaryExpression: PrimaryExpression{
				Atom: &Atom{
					Identifier: ident,
					Tokens:     tks,
				},
				Tokens: tks,
			},
			Tokens: tks,
		},
		&UnaryExpression{
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
		UnaryExpression{
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
		&MultiplyExpression{
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
		MultiplyExpression{
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
		&AddExpression{
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
		AddExpression{
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
		&ShiftExpression{
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
		ShiftExpression{
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
		&AndExpression{
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
		AndExpression{
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
		&XorExpression{
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
		XorExpression{
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
		&OrExpression{
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
		OrExpression{
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
		&Comparison{
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
		Comparison{
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
		&NotTest{
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
		NotTest{
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
		&AndTest{
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
		AndTest{
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
		&OrTest{
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
		OrTest{
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
		&ConditionalExpression{
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
		ConditionalExpression{
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

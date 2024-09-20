package python

import "testing"

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
	}, func(t *test) (Type, error) {
		var ce ConditionalExpression

		err := ce.parse(t.Tokens)

		return ce, err
	})
}
package python

import "io"

func (a AddExpression) printSource(w io.Writer, v bool) {
	a.MultiplyExpression.printSource(w, v)

	if a.Add != nil && a.AddExpression != nil {
		io.WriteString(w, " ")
		io.WriteString(w, a.Add.Data)
		io.WriteString(w, " ")
		a.AddExpression.printSource(w, v)
	}
}

func (a AndExpression) printSource(w io.Writer, v bool) {
	a.ShiftExpression.printSource(w, v)

	if a.AndExpression != nil {
		if v {
			io.WriteString(w, " & ")
		} else {
			io.WriteString(w, "&")
		}

		a.AndExpression.printSource(w, v)
	}
}

func (a AndTest) printSource(w io.Writer, v bool) {
	a.NotTest.printSource(w, v)

	if a.AndTest != nil {
		io.WriteString(w, " and ")
		a.AndTest.printSource(w, v)
	}
}

func (f AnnotatedAssignmentStatement) printSource(w io.Writer, v bool) {
}

func (f ArgumentList) printSource(w io.Writer, v bool) {
}

func (f ArgumentListOrComprehension) printSource(w io.Writer, v bool) {
}

func (f AssertStatement) printSource(w io.Writer, v bool) {
}

func (f AssignmentExpressionAndSuite) printSource(w io.Writer, v bool) {
}

func (f AssignmentExpression) printSource(w io.Writer, v bool) {
}

func (f AssignmentStatement) printSource(w io.Writer, v bool) {
}

func (a Atom) printSource(w io.Writer, v bool) {
	if a.Identifier != nil {
		io.WriteString(w, a.Identifier.Data)
	} else if a.Literal != nil {
		io.WriteString(w, a.Literal.Data)
	} else if a.Enclosure != nil {
		a.Enclosure.printSource(w, v)
	}
}

func (f AugmentedAssignmentStatement) printSource(w io.Writer, v bool) {
}

func (f AugTarget) printSource(w io.Writer, v bool) {
}

func (f ClassDefinition) printSource(w io.Writer, v bool) {
}

func (c Comparison) printSource(w io.Writer, v bool) {
	c.OrExpression.printSource(w, v)

	for _, ce := range c.Comparisons {
		ce.printSource(w, v)
	}
}

func (c ComparisonExpression) printSource(w io.Writer, v bool) {
	if len(c.ComparisonOperator) == 0 {
		return
	}

	switch c.ComparisonOperator[0].Data {
	case "<", ">", "==", ">=", "<=", "!=":
		if v {
			io.WriteString(w, " ")
		}

		io.WriteString(w, c.ComparisonOperator[0].Data)

		if v {
			io.WriteString(w, " ")
		}
	case "is":
		if c.ComparisonOperator[len(c.ComparisonOperator)-1].Data == "not" {
			io.WriteString(w, " is not ")
		} else {
			io.WriteString(w, " is ")
		}
	case "not":
		if c.ComparisonOperator[len(c.ComparisonOperator)-1].Data == "in" {
			io.WriteString(w, " not in ")
		} else {
			return
		}
	default:
		return
	}

	c.OrExpression.printSource(w, v)
}

func (f CompoundStatement) printSource(w io.Writer, v bool) {
}

func (f Comprehension) printSource(w io.Writer, v bool) {
}

func (f ComprehensionFor) printSource(w io.Writer, v bool) {
}

func (f ComprehensionIf) printSource(w io.Writer, v bool) {
}

func (f ComprehensionIterator) printSource(w io.Writer, v bool) {
}

func (c ConditionalExpression) printSource(w io.Writer, v bool) {
	c.OrTest.printSource(w, v)

	if c.If != nil && c.Else != nil {
		io.WriteString(w, " if ")
		c.If.printSource(w, v)
		io.WriteString(w, " else ")
		c.Else.printSource(w, v)
	}
}

func (f Decorators) printSource(w io.Writer, v bool) {
}

func (f DefParameter) printSource(w io.Writer, v bool) {
}

func (f DelStatement) printSource(w io.Writer, v bool) {
}

func (f DictDisplay) printSource(w io.Writer, v bool) {
}

func (f DictItem) printSource(w io.Writer, v bool) {
}

func (e Enclosure) printSource(w io.Writer, v bool) {
	if e.ParenthForm != nil {
		e.ParenthForm.printSource(w, v)
	} else if e.ListDisplay != nil {
		e.ListDisplay.printSource(w, v)
	} else if e.DictDisplay != nil {
		e.DictDisplay.printSource(w, v)
	} else if e.SetDisplay != nil {
		e.SetDisplay.printSource(w, v)
	} else if e.GeneratorExpression != nil {
		e.GeneratorExpression.printSource(w, v)
	} else if e.YieldAtom != nil {
		e.YieldAtom.printSource(w, v)
	}
}

func (f Except) printSource(w io.Writer, v bool) {
}

func (e Expression) printSource(w io.Writer, v bool) {
	if e.LambdaExpression != nil {
		e.LambdaExpression.printSource(w, v)
	} else if e.ConditionalExpression != nil {
		e.ConditionalExpression.printSource(w, v)
	}
}

func (f ExpressionList) printSource(w io.Writer, v bool) {
}

func (f File) printSource(w io.Writer, v bool) {
}

func (f ForStatement) printSource(w io.Writer, v bool) {
}

func (f FuncDefinition) printSource(w io.Writer, v bool) {
}

func (f GeneratorExpression) printSource(w io.Writer, v bool) {
}

func (f GlobalStatement) printSource(w io.Writer, v bool) {
}

func (f IfStatement) printSource(w io.Writer, v bool) {
}

func (f ImportStatement) printSource(w io.Writer, v bool) {
}

func (f KeywordArgument) printSource(w io.Writer, v bool) {
}

func (f KeywordItem) printSource(w io.Writer, v bool) {
}

func (l LambdaExpression) printSource(w io.Writer, v bool) {
	io.WriteString(w, "lambda ")

	if l.ParameterList != nil {
		l.ParameterList.printSource(w, v)
	}

	io.WriteString(w, ": ")

	l.Expression.printSource(w, v)
}

func (f ModuleAs) printSource(w io.Writer, v bool) {
}

func (f Module) printSource(w io.Writer, v bool) {
}

func (m MultiplyExpression) printSource(w io.Writer, v bool) {
	m.UnaryExpression.printSource(w, v)

	if m.Multiply != nil && m.MultiplyExpression != nil {
		io.WriteString(w, " ")
		io.WriteString(w, m.Multiply.Data)
		io.WriteString(w, " ")
		m.MultiplyExpression.printSource(w, v)
	}
}

func (f NonLocalStatement) printSource(w io.Writer, v bool) {
}

func (n NotTest) printSource(w io.Writer, v bool) {
	for i := n.Nots; i >= 0; i-- {
		io.WriteString(w, "not ")
	}

	n.Comparison.printSource(w, v)
}

func (o OrExpression) printSource(w io.Writer, v bool) {
	o.XorExpression.printSource(w, v)

	if o.OrExpression != nil {
		io.WriteString(w, " | ")
		o.OrExpression.printSource(w, v)
	}
}

func (o OrTest) printSource(w io.Writer, v bool) {
	o.AndTest.printSource(w, v)

	if o.OrTest != nil {
		io.WriteString(w, " or ")
		o.OrTest.printSource(w, v)
	}
}

func (f Parameter) printSource(w io.Writer, v bool) {
}

func (f ParameterList) printSource(w io.Writer, v bool) {
}

func (f PositionalArgument) printSource(w io.Writer, v bool) {
}

func (p PowerExpression) printSource(w io.Writer, v bool) {
	if p.AwaitExpression {
		io.WriteString(w, "await ")
	}

	p.PrimaryExpression.printSource(w, v)

	if p.UnaryExpression != nil {
		io.WriteString(w, " ** ")
		p.UnaryExpression.printSource(w, v)
	}
}

func (p PrimaryExpression) printSource(w io.Writer, v bool) {
	if p.Atom != nil {
		p.Atom.printSource(w, v)
	} else if p.PrimaryExpression != nil {
		p.PrimaryExpression.printSource(w, v)

		if p.AttributeRef != nil {
			io.WriteString(w, ".")
			io.WriteString(w, p.AttributeRef.Data)
		} else if p.Slicing != nil {
			p.Slicing.printSource(w, v)
		} else if p.Call != nil {
			p.Call.printSource(w, v)
		}
	}
}

func (f RaiseStatement) printSource(w io.Writer, v bool) {
}

func (f RelativeModule) printSource(w io.Writer, v bool) {
}

func (f ReturnStatement) printSource(w io.Writer, v bool) {
}

func (s ShiftExpression) printSource(w io.Writer, v bool) {
	s.AddExpression.printSource(w, v)

	if s.Shift != nil && s.ShiftExpression != nil {
		if v {
			io.WriteString(w, " ")
		}

		io.WriteString(w, s.Shift.Data)

		if v {
			io.WriteString(w, " ")
		}

		s.ShiftExpression.printSource(w, v)
	}
}

func (f SimpleStatement) printSource(w io.Writer, v bool) {
}

func (f SliceItem) printSource(w io.Writer, v bool) {
}

func (f SliceList) printSource(w io.Writer, v bool) {
}

func (f StarredExpression) printSource(w io.Writer, v bool) {
}

func (f StarredItem) printSource(w io.Writer, v bool) {
}

func (f StarredList) printSource(w io.Writer, v bool) {
}

func (f StarredListOrComprehension) printSource(w io.Writer, v bool) {
}

func (f StarredOrKeywordArgument) printSource(w io.Writer, v bool) {
}

func (f Statement) printSource(w io.Writer, v bool) {
}

func (f StatementList) printSource(w io.Writer, v bool) {
}

func (f Suite) printSource(w io.Writer, v bool) {
}

func (f Target) printSource(w io.Writer, v bool) {
}

func (f TargetList) printSource(w io.Writer, v bool) {
}

func (f TryStatement) printSource(w io.Writer, v bool) {
}

func (f TypeParam) printSource(w io.Writer, v bool) {
}

func (f TypeParams) printSource(w io.Writer, v bool) {
}

func (f TypeStatement) printSource(w io.Writer, v bool) {
}

func (u UnaryExpression) printSource(w io.Writer, v bool) {
	if u.PowerExpression != nil {
		u.PowerExpression.printSource(w, v)
	} else if u.Unary != nil && u.UnaryExpression != nil {
		io.WriteString(w, u.Unary.Data)
		u.UnaryExpression.printSource(w, v)
	}
}

func (f WhileStatement) printSource(w io.Writer, v bool) {
}

func (f WithItem) printSource(w io.Writer, v bool) {
}

func (f WithStatement) printSource(w io.Writer, v bool) {
}

func (f WithStatementContents) printSource(w io.Writer, v bool) {
}

func (x XorExpression) printSource(w io.Writer, v bool) {
	x.AndExpression.printSource(w, v)

	if x.XorExpression != nil {
		io.WriteString(w, " ^ ")
		x.XorExpression.printSource(w, v)
	}
}

func (f YieldExpression) printSource(w io.Writer, v bool) {
}

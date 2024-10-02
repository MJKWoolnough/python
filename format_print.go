package python

import "io"

func (a AddExpression) printSource(w io.Writer, v bool) {
	a.MultiplyExpression.printSource(w, v)

	if a.Add != nil && a.AddExpression != nil {
		if v {
			io.WriteString(w, " ")
		}

		io.WriteString(w, a.Add.Data)

		if v {
			io.WriteString(w, " ")
		}

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

func (a ArgumentListOrComprehension) printSource(w io.Writer, v bool) {
	io.WriteString(w, "(")

	if a.ArgumentList != nil {
		a.ArgumentList.printSource(w, v)
	} else if a.Comprehension != nil {
		a.Comprehension.printSource(w, v)
	}

	io.WriteString(w, ")")
}

func (f AssertStatement) printSource(w io.Writer, v bool) {
}

func (f AssignmentExpressionAndSuite) printSource(w io.Writer, v bool) {
}

func (a AssignmentExpression) printSource(w io.Writer, v bool) {
	if a.Identifier != nil {
		io.WriteString(w, a.Identifier.Data)

		if v {
			io.WriteString(w, " := ")
		} else {
			io.WriteString(w, ":=")
		}
	}

	a.Expression.printSource(w, v)
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

func (c Comprehension) printSource(w io.Writer, v bool) {
	c.AssignmentExpression.printSource(w, v)
	io.WriteString(w, " ")
	c.ComprehensionFor.printSource(w, v)
}

func (c ComprehensionFor) printSource(w io.Writer, v bool) {
	if c.Async {
		io.WriteString(w, "async ")
	}

	io.WriteString(w, "for ")
	c.TargetList.printSource(w, v)
	io.WriteString(w, " in ")
	c.OrTest.printSource(w, v)

	if c.ComprehensionIterator != nil {
		c.ComprehensionIterator.printSource(w, v)
	}
}

func (c ComprehensionIf) printSource(w io.Writer, v bool) {
	io.WriteString(w, "if ")
	c.OrTest.printSource(w, v)

	if c.ComprehensionIterator != nil {
		io.WriteString(w, " ")
		c.ComprehensionIterator.printSource(w, v)
	}
}

func (c ComprehensionIterator) printSource(w io.Writer, v bool) {
	if c.ComprehensionIf != nil {
		c.ComprehensionIf.printSource(w, v)
	} else if c.ComprehensionFor != nil {
		c.ComprehensionFor.printSource(w, v)
	}
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

func (d DictDisplay) printSource(w io.Writer, v bool) {
	for n, di := range d.DictItems {
		if n > 0 {
			io.WriteString(w, ", ")
		}

		di.printSource(w, v)
	}

	if d.DictComprehension != nil {
		io.WriteString(w, " ")
		d.DictComprehension.printSource(w, v)
	}
}

func (d DictItem) printSource(w io.Writer, v bool) {
	if d.OrExpression != nil {
		io.WriteString(w, "**")
		d.OrExpression.printSource(w, v)
	} else if d.Key != nil && d.Value != nil {
		d.Key.printSource(w, v)

		if v {
			io.WriteString(w, ": ")
		} else {
			io.WriteString(w, ":")
		}

		d.Value.printSource(w, v)
	}
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

func (e ExpressionList) printSource(w io.Writer, v bool) {
	for n, ex := range e.Expressions {
		if n > 0 {
			io.WriteString(w, ", ")
		}

		ex.printSource(w, v)
	}
}

func (f File) printSource(w io.Writer, v bool) {
}

func (f ForStatement) printSource(w io.Writer, v bool) {
}

func (f FuncDefinition) printSource(w io.Writer, v bool) {
}

func (g GeneratorExpression) printSource(w io.Writer, v bool) {
	g.Expression.printSource(w, v)
	io.WriteString(w, " ")
	g.ComprehensionFor.printSource(w, v)
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

	if v {
		io.WriteString(w, ": ")
	} else {
		io.WriteString(w, ":")
	}

	l.Expression.printSource(w, v)
}

func (f ModuleAs) printSource(w io.Writer, v bool) {
}

func (f Module) printSource(w io.Writer, v bool) {
}

func (m MultiplyExpression) printSource(w io.Writer, v bool) {
	m.UnaryExpression.printSource(w, v)

	if m.Multiply != nil && m.MultiplyExpression != nil {
		if v {
			io.WriteString(w, " ")
		}

		io.WriteString(w, m.Multiply.Data)

		if v {
			io.WriteString(w, " ")
		}

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
		if v {
			io.WriteString(w, " | ")
		} else {
			io.WriteString(w, "|")
		}
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
		if v {
			io.WriteString(w, " ** ")
		} else {
			io.WriteString(w, "**")
		}

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

func (s SliceItem) printSource(w io.Writer, v bool) {
	if s.Expression != nil {
		s.Expression.printSource(w, v)
	} else if s.LowerBound != nil {
		s.LowerBound.printSource(w, v)

		if s.UpperBound != nil {
			if v {
				io.WriteString(w, " : ")
			} else {
				io.WriteString(w, ":")
			}

			s.UpperBound.printSource(w, v)

			if s.Stride != nil {
				if v {
					io.WriteString(w, " : ")
				} else {
					io.WriteString(w, ":")
				}

				s.Stride.printSource(w, v)
			}
		}
	}
}

func (s SliceList) printSource(w io.Writer, v bool) {
	io.WriteString(w, "[")

	for n, si := range s.SliceItems {
		if n > 0 {
			if v {
				io.WriteString(w, ", ")
			} else {
				io.WriteString(w, ",")
			}
		}

		si.printSource(w, v)
	}

	io.WriteString(w, "]")
}

func (f StarredExpression) printSource(w io.Writer, v bool) {
}

func (f StarredItem) printSource(w io.Writer, v bool) {
}

func (f StarredList) printSource(w io.Writer, v bool) {
}

func (s StarredListOrComprehension) printSource(w io.Writer, v bool) {
	if s.StarredList != nil {
		s.StarredList.printSource(w, v)
	} else if s.Comprehension != nil {
		s.Comprehension.printSource(w, v)
	}
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
		if v {
			io.WriteString(w, " ^ ")
		} else {
			io.WriteString(w, "^")
		}

		x.XorExpression.printSource(w, v)
	}
}

func (f YieldExpression) printSource(w io.Writer, v bool) {
}

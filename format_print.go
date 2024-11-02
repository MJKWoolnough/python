package python

import (
	"io"
)

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

func (a AnnotatedAssignmentStatement) printSource(w io.Writer, v bool) {
	a.AugTarget.printSource(w, v)
	io.WriteString(w, ": ")
	a.Expression.printSource(w, v)

	if a.StarredExpression != nil {
		io.WriteString(w, " = ")
		a.StarredExpression.printSource(w, v)
	} else if a.YieldExpression != nil {
		io.WriteString(w, " = ")
		a.YieldExpression.printSource(w, v)
	}
}

func (a ArgumentList) printSource(w io.Writer, v bool) {
	first := true

	for _, p := range a.PositionalArguments {
		if !first {
			io.WriteString(w, ",")

			if v {
				io.WriteString(w, " ")
			}
		}

		p.printSource(w, v)

		first = false
	}

	for _, s := range a.StarredAndKeywordArguments {
		if !first {
			io.WriteString(w, ",")

			if v {
				io.WriteString(w, " ")
			}
		}

		s.printSource(w, v)

		first = false
	}

	for _, k := range a.KeywordArguments {
		if !first {
			io.WriteString(w, ",")

			if v {
				io.WriteString(w, " ")
			}
		}

		k.printSource(w, v)

		first = false
	}
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

func (a AssertStatement) printSource(w io.Writer, v bool) {
	if len(a.Expressions) > 0 {
		io.WriteString(w, "assert ")
		a.Expressions[0].printSource(w, v)

		for _, e := range a.Expressions[1:] {
			io.WriteString(w, ",")

			if v {
				io.WriteString(w, " ")
			}

			e.printSource(w, v)
		}
	}
}

func (a AssignmentExpressionAndSuite) printSource(w io.Writer, v bool) {
	a.AssignmentExpression.printSource(w, v)
	io.WriteString(w, ":")
	a.Suite.printSource(w, v)
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

func (a AssignmentStatement) printSource(w io.Writer, v bool) {
	for _, t := range a.TargetLists {
		t.printSource(w, v)

		if v {
			io.WriteString(w, " = ")
		} else {
			io.WriteString(w, "=")
		}
	}

	if a.StarredExpression != nil {
		a.StarredExpression.printSource(w, v)
	} else if a.YieldExpression != nil {
		a.YieldExpression.printSource(w, v)
	}
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

func (a AugmentedAssignmentStatement) printSource(w io.Writer, v bool) {
	a.AugTarget.printSource(w, v)

	if v {
		io.WriteString(w, " ")
	}

	io.WriteString(w, a.AugOp.Data)

	if v {
		io.WriteString(w, " ")
	}

	if a.ExpressionList != nil {
		a.ExpressionList.printSource(w, v)
	} else if a.YieldExpression != nil {
		a.YieldExpression.printSource(w, v)
	}
}

func (a AugTarget) printSource(w io.Writer, v bool) {
	a.PrimaryExpression.printSource(w, v)
}

func (c ClassDefinition) printSource(w io.Writer, v bool) {
	if c.Decorators != nil {
		c.Decorators.printSource(w, v)
	}

	io.WriteString(w, "class ")
	io.WriteString(w, c.ClassName.Data)

	if c.TypeParams != nil {
		c.TypeParams.printSource(w, v)
	}

	if c.Inheritance != nil {
		io.WriteString(w, "(")
		c.Inheritance.printSource(w, v)
		io.WriteString(w, ")")
	}

	io.WriteString(w, ":")
	c.Suite.printSource(w, v)
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

func (c CompoundStatement) printSource(w io.Writer, v bool) {
	if c.If != nil {
		c.If.printSource(w, v)
	} else if c.While != nil {
		c.While.printSource(w, v)
	} else if c.For != nil {
		c.For.printSource(w, v)
	} else if c.Try != nil {
		c.Try.printSource(w, v)
	} else if c.With != nil {
		c.With.printSource(w, v)
	} else if c.Func != nil {
		c.Func.printSource(w, v)
	} else if c.Class != nil {
		c.Class.printSource(w, v)
	}
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

func (d Decorators) printSource(w io.Writer, v bool) {
	for _, dc := range d.Decorators {
		io.WriteString(w, "@")
		dc.printSource(w, v)
		io.WriteString(w, "\n")
	}
}

func (d DefParameter) printSource(w io.Writer, v bool) {
	d.Parameter.printSource(w, v)

	if d.Value != nil {
		io.WriteString(w, " = ")
		d.Value.printSource(w, v)
	}
}

func (d DelStatement) printSource(w io.Writer, v bool) {
	io.WriteString(w, "del ")
	d.TargetList.printSource(w, v)
}

func (d DictDisplay) printSource(w io.Writer, v bool) {
	if len(d.DictItems) > 0 {
		d.DictItems[0].printSource(w, v)

		for _, di := range d.DictItems[1:] {
			io.WriteString(w, ", ")
			di.printSource(w, v)
		}
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
		io.WriteString(w, "(")
		e.ParenthForm.printSource(w, v)
		io.WriteString(w, ")")
	} else if e.ListDisplay != nil {
		io.WriteString(w, "[")
		e.ListDisplay.printSource(w, v)
		io.WriteString(w, "]")
	} else if e.DictDisplay != nil {
		io.WriteString(w, "{")
		e.DictDisplay.printSource(w, v)
		io.WriteString(w, "}")
	} else if e.SetDisplay != nil {
		io.WriteString(w, "{")
		e.SetDisplay.printSource(w, v)
		io.WriteString(w, "}")
	} else if e.GeneratorExpression != nil {
		io.WriteString(w, "(")
		e.GeneratorExpression.printSource(w, v)
		io.WriteString(w, ")")
	} else if e.YieldAtom != nil {
		io.WriteString(w, "(")
		e.YieldAtom.printSource(w, v)
		io.WriteString(w, ")")
	}
}

func (e Except) printSource(w io.Writer, v bool) {
	e.Expression.printSource(w, v)

	if e.Identifier != nil {
		io.WriteString(w, " as ")
		io.WriteString(w, e.Identifier.Data)
	}

	io.WriteString(w, ":")
	e.Suite.printSource(w, v)
}

func (e Expression) printSource(w io.Writer, v bool) {
	if e.LambdaExpression != nil {
		e.LambdaExpression.printSource(w, v)
	} else if e.ConditionalExpression != nil {
		e.ConditionalExpression.printSource(w, v)
	}
}

func (e ExpressionList) printSource(w io.Writer, v bool) {
	if len(e.Expressions) > 0 {
		e.Expressions[0].printSource(w, v)

		for _, ex := range e.Expressions[1:] {
			if v {
				io.WriteString(w, ", ")
			} else {
				io.WriteString(w, ",")
			}

			ex.printSource(w, v)
		}
	}
}

func (f File) printSource(w io.Writer, v bool) {
	for _, s := range f.Statements {
		s.printSource(w, v)
		io.WriteString(w, "\n")
	}
}

func (f FlexibleExpressionListOrComprehension) printSource(w io.Writer, v bool) {
	if f.FlexibleExpressionList != nil {
		f.FlexibleExpressionList.printSource(w, v)
	} else if f.Comprehension != nil {
		f.Comprehension.printSource(w, v)
	}
}

func (f FlexibleExpressionList) printSource(w io.Writer, v bool) {
	if len(f.FlexibleExpressions) > 0 {
		f.FlexibleExpressions[0].printSource(w, v)
		for _, fe := range f.FlexibleExpressions[1:] {
			if v {
				io.WriteString(w, ", ")
			} else {
				io.WriteString(w, ",")
			}

			fe.printSource(w, v)
		}
	}
}

func (f FlexibleExpression) printSource(w io.Writer, v bool) {
	if f.AssignmentExpression != nil {
		f.AssignmentExpression.printSource(w, v)
	} else if f.StarredExpression != nil {
		f.StarredExpression.printSource(w, v)
	}
}

func (f ForStatement) printSource(w io.Writer, v bool) {
	if f.Async {
		io.WriteString(w, "async ")
	}

	io.WriteString(w, "for ")
	f.TargetList.printSource(w, v)
	io.WriteString(w, " in ")
	f.StarredList.printSource(w, v)
	io.WriteString(w, ":")
	f.Suite.printSource(w, v)

	if f.Else != nil {
		io.WriteString(w, "\nelse:")
		f.Else.printSource(w, v)
	}
}

func (f FuncDefinition) printSource(w io.Writer, v bool) {
	if f.Decorators != nil {
		f.Decorators.printSource(w, v)
	}

	if f.Async {
		io.WriteString(w, "async ")
	}

	io.WriteString(w, "def ")
	io.WriteString(w, f.FuncName.Data)

	if f.TypeParams != nil {
		f.TypeParams.printSource(w, v)
	}

	io.WriteString(w, "(")
	f.ParameterList.printSource(w, v)
	io.WriteString(w, ")")

	if f.Expression != nil {
		if v {
			io.WriteString(w, " -> ")
		} else {
			io.WriteString(w, "->")
		}
		f.Expression.printSource(w, v)
	}

	io.WriteString(w, ":")
	f.Suite.printSource(w, v)
}

func (g GeneratorExpression) printSource(w io.Writer, v bool) {
	g.Expression.printSource(w, v)
	io.WriteString(w, " ")
	g.ComprehensionFor.printSource(w, v)
}

func (g GlobalStatement) printSource(w io.Writer, v bool) {
	if len(g.Identifiers) > 0 {
		io.WriteString(w, "global ")
		io.WriteString(w, g.Identifiers[0].Data)

		for _, t := range g.Identifiers[1:] {
			if v {
				io.WriteString(w, ", ")
			} else {
				io.WriteString(w, ",")
			}

			io.WriteString(w, t.Data)
		}
	}
}

func (i IfStatement) printSource(w io.Writer, v bool) {
	io.WriteString(w, "if ")
	i.AssignmentExpression.printSource(w, v)
	io.WriteString(w, ":")
	i.Suite.printSource(w, v)

	for _, e := range i.Elif {
		io.WriteString(w, "\nelif ")
		e.printSource(w, v)
	}

	if i.Else != nil {
		io.WriteString(w, "\nelse:")
		i.Else.printSource(w, v)
	}
}

func (i ImportStatement) printSource(w io.Writer, v bool) {
	if i.RelativeModule != nil {
		io.WriteString(w, "from ")
		i.RelativeModule.printSource(w, v)
		io.WriteString(w, " import ")
	} else {
		io.WriteString(w, "import ")
	}

	if len(i.Modules) > 0 {
		i.Modules[0].printSource(w, v)

		for _, m := range i.Modules[1:] {
			if v {
				io.WriteString(w, ", ")
			} else {
				io.WriteString(w, ",")
			}

			m.printSource(w, v)
		}
	}
}

func (k KeywordArgument) printSource(w io.Writer, v bool) {
	if k.Expression != nil {
		io.WriteString(w, "**")
		k.Expression.printSource(w, v)
	} else if k.KeywordItem != nil {
		k.KeywordItem.printSource(w, v)
	}
}

func (k KeywordItem) printSource(w io.Writer, v bool) {
	if k.Identifier != nil {
		io.WriteString(w, k.Identifier.Data)

		if v {
			io.WriteString(w, " = ")
		} else {
			io.WriteString(w, "=")
		}

		k.Expression.printSource(w, v)
	}
}

func (l LambdaExpression) printSource(w io.Writer, v bool) {
	if l.ParameterList != nil {
		io.WriteString(w, "lambda ")
		l.ParameterList.printSource(w, v)
	} else {
		io.WriteString(w, "lambda")
	}

	if v {
		io.WriteString(w, ": ")
	} else {
		io.WriteString(w, ":")
	}

	l.Expression.printSource(w, v)
}

func (m ModuleAs) printSource(w io.Writer, v bool) {
	m.Module.printSource(w, v)

	if m.As != nil {
		io.WriteString(w, " as ")
		io.WriteString(w, m.As.Data)
	}
}

func (m Module) printSource(w io.Writer, v bool) {
	if len(m.Identifiers) > 0 {
		io.WriteString(w, m.Identifiers[0].Data)

		for _, i := range m.Identifiers[1:] {
			io.WriteString(w, ".")
			io.WriteString(w, i.Data)
		}
	}
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

func (n NonLocalStatement) printSource(w io.Writer, v bool) {
	if len(n.Identifiers) > 0 {
		io.WriteString(w, "nonlocal ")
		io.WriteString(w, n.Identifiers[0].Data)

		for _, t := range n.Identifiers[1:] {
			if v {
				io.WriteString(w, ", ")
			} else {
				io.WriteString(w, ",")
			}

			io.WriteString(w, t.Data)
		}
	}
}

func (n NotTest) printSource(w io.Writer, v bool) {
	for i := n.Nots; i > 0; i-- {
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

func (p Parameter) printSource(w io.Writer, v bool) {
	io.WriteString(w, p.Identifier.Data)

	if p.Type != nil {
		io.WriteString(w, ": ")
		p.Type.printSource(w, v)
	}
}

func (p ParameterList) printSource(w io.Writer, v bool) {
	first := len(p.DefParameters) == 0

	if !first {
		for _, d := range p.DefParameters {
			d.printSource(w, v)

			if v {
				io.WriteString(w, ", ")
			} else {
				io.WriteString(w, ",")
			}
		}

		io.WriteString(w, "/")
	}

	for _, n := range p.NoPosOnly {
		if first {
			first = false
		} else if v {
			io.WriteString(w, ", ")
		} else {
			io.WriteString(w, ",")
		}

		n.printSource(w, v)
	}

	if p.StarArg != nil {
		if first {
			first = false
		} else if v {
			io.WriteString(w, ", ")
		} else {
			io.WriteString(w, ",")
		}

		io.WriteString(w, "*")
		p.StarArg.printSource(w, v)

		for _, d := range p.StarArgs {
			if v {
				io.WriteString(w, ", ")
			} else {
				io.WriteString(w, ",")
			}

			d.printSource(w, v)
		}
	}

	if p.StarStarArg != nil {
		if first {
			first = false
		} else if v {
			io.WriteString(w, ", ")
		} else {
			io.WriteString(w, ",")
		}

		io.WriteString(w, "**")
		p.StarStarArg.printSource(w, v)
	}
}

func (p PositionalArgument) printSource(w io.Writer, v bool) {
	if p.AssignmentExpression != nil {
		p.AssignmentExpression.printSource(w, v)
	} else if p.Expression != nil {
		io.WriteString(w, "*")
		p.Expression.printSource(w, v)
	}
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

func (r RaiseStatement) printSource(w io.Writer, v bool) {
	if r.Expression == nil {
		io.WriteString(w, "raise")
	} else {
		io.WriteString(w, "raise")
		r.Expression.printSource(w, v)

		if r.From != nil {
			io.WriteString(w, " from ")
			r.From.printSource(w, v)
		}
	}
}

func (r RelativeModule) printSource(w io.Writer, v bool) {
	for range r.Dots {
		io.WriteString(w, ".")
	}

	if r.Module != nil {
		r.Module.printSource(w, v)
	}
}

func (r ReturnStatement) printSource(w io.Writer, v bool) {
	if r.Expression != nil {
		io.WriteString(w, "return ")
		r.Expression.printSource(w, v)
	} else {
		io.WriteString(w, "return")
	}
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

func (s SimpleStatement) printSource(w io.Writer, v bool) {
	if s.AssertStatement != nil {
		s.AssertStatement.printSource(w, v)
	} else if s.DelStatement != nil {
		s.DelStatement.printSource(w, v)
	} else if s.ReturnStatement != nil {
		s.ReturnStatement.printSource(w, v)
	} else if s.YieldStatement != nil {
		s.YieldStatement.printSource(w, v)
	} else if s.RaiseStatement != nil {
		s.RaiseStatement.printSource(w, v)
	} else if s.ImportStatement != nil {
		s.ImportStatement.printSource(w, v)
	} else if s.GlobalStatement != nil {
		s.GlobalStatement.printSource(w, v)
	} else if s.NonLocalStatement != nil {
		s.NonLocalStatement.printSource(w, v)
	} else if s.TypeStatement != nil {
		s.TypeStatement.printSource(w, v)
	} else if s.AssignmentStatement != nil {
		s.AssignmentStatement.printSource(w, v)
	} else if s.AnnotatedAssignmentStatement != nil {
		s.AnnotatedAssignmentStatement.printSource(w, v)
	} else if s.AugmentedAssignmentStatement != nil {
		s.AugmentedAssignmentStatement.printSource(w, v)
	} else if s.Type == StatementPass {
		io.WriteString(w, "pass")
	} else if s.Type == StatementBreak {
		io.WriteString(w, "break")
	} else if s.Type == StatementContinue {
		io.WriteString(w, "continue")
	}
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

	if len(s.SliceItems) > 0 {
		s.SliceItems[0].printSource(w, v)

		for _, si := range s.SliceItems[1:] {
			if v {
				io.WriteString(w, ", ")
			} else {
				io.WriteString(w, ",")
			}

			si.printSource(w, v)
		}
	}

	io.WriteString(w, "]")
}

func (s StarredExpression) printSource(w io.Writer, v bool) {
	if s.Expression != nil {
		s.Expression.printSource(w, v)
	} else if s.StarredList != nil {
		s.StarredList.printSource(w, v)
	}
}

func (s StarredItem) printSource(w io.Writer, v bool) {
	if s.AssignmentExpression != nil {
		s.AssignmentExpression.printSource(w, v)
	} else if s.OrExpr != nil {
		io.WriteString(w, "*")
		s.OrExpr.printSource(w, v)
	}
}

func (s StarredList) printSource(w io.Writer, v bool) {
	if len(s.StarredItems) > 0 {
		s.StarredItems[0].printSource(w, v)

		for _, si := range s.StarredItems[1:] {
			io.WriteString(w, ", ")
			si.printSource(w, v)
		}
	}
}

func (s StarredOrKeyword) printSource(w io.Writer, v bool) {
	if s.Expression != nil {
		io.WriteString(w, "*")
		s.Expression.printSource(w, v)
	} else if s.KeywordItem != nil {
		s.KeywordItem.printSource(w, v)
	}
}

func (s Statement) printSource(w io.Writer, v bool) {
	if s.StatementList != nil {
		s.StatementList.printSource(w, v)
	} else if s.CompoundStatement != nil {
		s.CompoundStatement.printSource(w, v)
	}
}

func (s StatementList) printSource(w io.Writer, v bool) {
	if len(s.Statements) > 0 {
		s.Statements[0].printSource(w, v)

		for _, ss := range s.Statements[1:] {
			io.WriteString(w, "; ")
			ss.printSource(w, v)
		}
	}
}

func (s Suite) printSource(w io.Writer, v bool) {
	if s.StatementList != nil {
		if v {
			io.WriteString(w, " ")
		}

		s.StatementList.printSource(w, v)
	} else {
		ip := indentPrinter{Writer: w}

		for _, stmt := range s.Statements {
			ip.WriteString("\n")
			stmt.printSource(&ip, v)
		}
	}
}

func (t Target) printSource(w io.Writer, v bool) {
	if t.PrimaryExpression != nil {
		t.PrimaryExpression.printSource(w, v)
	} else if t.Tuple != nil {
		io.WriteString(w, "(")
		t.Tuple.printSource(w, v)
		io.WriteString(w, ")")
	} else if t.Array != nil {
		io.WriteString(w, "[")
		t.Array.printSource(w, v)
		io.WriteString(w, "]")
	} else if t.Star != nil {
		io.WriteString(w, "*")
		t.Star.printSource(w, v)
	}
}

func (t TargetList) printSource(w io.Writer, v bool) {
	if len(t.Targets) > 0 {
		t.Targets[0].printSource(w, v)

		for _, tg := range t.Targets[1:] {
			io.WriteString(w, ", ")
			tg.printSource(w, v)
		}
	}
}

func (t TryStatement) printSource(w io.Writer, v bool) {
	io.WriteString(w, "try:")
	t.Try.printSource(w, v)

	if len(t.Except) > 0 {
		io.WriteString(w, "\nexcept ")

		if t.Groups {
			io.WriteString(w, "*")
		}

		t.Except[0].printSource(w, v)

		for _, e := range t.Except[1:] {
			io.WriteString(w, "\nexcept ")

			if t.Groups {
				io.WriteString(w, "*")
			}

			e.printSource(w, v)
		}
	}

	if t.Else != nil {
		io.WriteString(w, "\nelse:")
		t.Else.printSource(w, v)
	}

	if t.Finally != nil {
		io.WriteString(w, "\nfinally:")
		t.Finally.printSource(w, v)
	}
}

func (t TypeParam) printSource(w io.Writer, v bool) {
	if t.Type == TypeParamVar {
		io.WriteString(w, "*")
	} else if t.Type == TypeParamVarTuple {
		io.WriteString(w, "**")
	}

	io.WriteString(w, t.Identifier.Data)

	if t.Expression != nil {
		if v {
			io.WriteString(w, ": ")
		} else {
			io.WriteString(w, ":")
		}

		t.Expression.printSource(w, v)
	}
}

func (t TypeParams) printSource(w io.Writer, v bool) {
	io.WriteString(w, "[")

	if len(t.TypeParams) > 0 {
		t.TypeParams[0].printSource(w, v)

		for _, tp := range t.TypeParams[1:] {
			if v {
				io.WriteString(w, ", ")
			} else {
				io.WriteString(w, ",")
			}

			tp.printSource(w, v)
		}
	}

	io.WriteString(w, "]")
}

func (t TypeStatement) printSource(w io.Writer, v bool) {
	io.WriteString(w, "type ")
	io.WriteString(w, t.Identifier.Data)

	if t.TypeParams != nil {
		t.TypeParams.printSource(w, v)
	}

	io.WriteString(w, " = ")
	t.Expression.printSource(w, v)
}

func (u UnaryExpression) printSource(w io.Writer, v bool) {
	if u.PowerExpression != nil {
		u.PowerExpression.printSource(w, v)
	} else if u.Unary != nil && u.UnaryExpression != nil {
		io.WriteString(w, u.Unary.Data)
		u.UnaryExpression.printSource(w, v)
	}
}

func (ws WhileStatement) printSource(w io.Writer, v bool) {
	io.WriteString(w, "while ")
	ws.AssignmentExpression.printSource(w, v)
	io.WriteString(w, ":")
	ws.Suite.printSource(w, v)

	if ws.Else != nil {
		io.WriteString(w, "\nelse:")
		ws.Else.printSource(w, v)
	}
}

func (wi WithItem) printSource(w io.Writer, v bool) {
	wi.Expression.printSource(w, v)

	if wi.Target != nil {
		io.WriteString(w, " as ")
		wi.Target.printSource(w, v)
	}
}

func (ws WithStatement) printSource(w io.Writer, v bool) {
	io.WriteString(w, "with ")
	ws.Contents.printSource(w, v)
	io.WriteString(w, ":")
	ws.Suite.printSource(w, v)
}

func (wc WithStatementContents) printSource(w io.Writer, v bool) {
	if len(wc.Items) > 0 {
		wc.Items[0].printSource(w, v)

		for _, i := range wc.Items[1:] {
			i.printSource(w, v)
		}
	}
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

func (y YieldExpression) printSource(w io.Writer, v bool) {
	if y.From != nil {
		io.WriteString(w, "yield from")
		y.From.printSource(w, v)
	} else if y.ExpressionList != nil {
		io.WriteString(w, "yield ")
		y.ExpressionList.printSource(w, v)
	}
}

package python

func (a AddExpression) printSource(w writer, v bool) {
	a.MultiplyExpression.printSource(w, v)

	if a.Add != nil && a.AddExpression != nil {
		if v {
			w.WriteString(" ")

			if w.InMultiline() {
				a.Comments[0].printSource(w, true)
			}

			w.WriteString(a.Add.Data)
			w.WriteString(" ")

			if w.InMultiline() {
				a.Comments[1].printSource(w, true)
			}
		} else {
			w.WriteString(a.Add.Data)
		}

		a.AddExpression.printSource(w, v)
	}
}

func (a AndExpression) printSource(w writer, v bool) {
	a.ShiftExpression.printSource(w, v)

	if a.AndExpression != nil {
		if v {
			w.WriteString(" ")

			if w.InMultiline() {
				a.Comments[0].printSource(w, true)
			}

			w.WriteString("& ")

			if w.InMultiline() {
				a.Comments[1].printSource(w, true)
			}
		} else {
			w.WriteString("&")
		}

		a.AndExpression.printSource(w, v)
	}
}

func (a AndTest) printSource(w writer, v bool) {
	a.NotTest.printSource(w, v)

	if a.AndTest != nil {
		w.WriteString(" ")

		if v && w.InMultiline() {
			a.Comments[0].printSource(w, true)
		}

		w.WriteString("and ")

		if v && w.InMultiline() {
			a.Comments[1].printSource(w, true)
		}

		a.AndTest.printSource(w, v)
	}
}

func (a AnnotatedAssignmentStatement) printSource(w writer, v bool) {
	a.AugTarget.printSource(w, v)

	if v {
		w.WriteString(": ")
	} else {
		w.WriteString(":")
	}

	a.Expression.printSource(w, v)

	if a.StarredExpression != nil {
		if v {
			w.WriteString(" = ")
		} else {
			w.WriteString("=")
		}

		a.StarredExpression.printSource(w, v)
	} else if a.YieldExpression != nil {
		if v {
			w.WriteString(" = ")
		} else {
			w.WriteString("=")
		}

		a.YieldExpression.printSource(w, v)
	}
}

func (a ArgumentList) printSource(w writer, v bool) {
	first := true

	var last Comments

	for _, p := range a.PositionalArguments {
		if !first {
			w.WriteString(",")

			if v {
				w.WriteString(" ")
			}
		}

		p.printSource(w, v)

		first = false
		last = p.Comments[2]
	}

	for _, s := range a.StarredAndKeywordArguments {
		if !first {
			w.WriteString(",")

			if v {
				w.WriteString(" ")
			}
		}

		s.printSource(w, v)

		first = false
		last = s.Comments[1]
	}

	for _, k := range a.KeywordArguments {
		if !first {
			w.WriteString(",")

			if v {
				w.WriteString(" ")
			}
		}

		k.printSource(w, v)

		first = false
		last = k.Comments[2]
	}

	if v && len(last) > 0 {
		line := last[0].Line
		needComma := false

		for _, c := range last[1:] {
			line++

			if c.Line > line {
				needComma = true

				break
			}
		}

		if needComma {
			w.WriteString(",")
		}
	}
}

func (a ArgumentListOrComprehension) printSource(w writer, v bool) {
	w.WriteString("(")

	ip := w.Indent()

	if v && len(a.Comments[0]) > 0 {
		a.Comments[0].printSource(w, false)
		ip.WriteString("\n")
	}

	if a.ArgumentList != nil {
		a.ArgumentList.printSource(ip, v)
	} else if a.Comprehension != nil {
		a.Comprehension.printSource(ip, v)
	}

	if v && len(a.Comments[1]) > 0 {
		w.WriteString("\n")
		a.Comments[1].printSource(w, true)
	}

	w.WriteString(")")
}

func (a AssertStatement) printSource(w writer, v bool) {
	if len(a.Expressions) > 0 {
		w.WriteString("assert ")
		a.Expressions[0].printSource(w, v)

		for _, e := range a.Expressions[1:] {
			w.WriteString(",")

			if v {
				w.WriteString(" ")
			}

			e.printSource(w, v)
		}
	}
}

func (a AssignmentExpressionAndSuite) printSource(w writer, v bool) {
	a.AssignmentExpression.printSource(w, v)
	w.WriteString(":")
	a.Suite.printSource(w, v)
}

func (a AssignmentExpression) printSource(w writer, v bool) {
	if a.Identifier != nil {
		w.WriteString(a.Identifier.Data)

		if v {
			w.WriteString(" ")

			if w.InMultiline() {
				a.Comments[0].printSource(w, true)
			}

			w.WriteString(":= ")

			if w.InMultiline() {
				a.Comments[1].printSource(w, true)
			}
		} else {
			w.WriteString(":=")
		}
	}

	a.Expression.printSource(w, v)
}

func (a AssignmentStatement) printSource(w writer, v bool) {
	for _, t := range a.TargetLists {
		t.printSource(w, v)

		if v {
			w.WriteString(" = ")
		} else {
			w.WriteString("=")
		}
	}

	if a.StarredExpression != nil {
		a.StarredExpression.printSource(w, v)
	} else if a.YieldExpression != nil {
		a.YieldExpression.printSource(w, v)
	}
}

func (a Atom) printSource(w writer, v bool) {
	if a.Identifier != nil {
		w.WriteString(a.Identifier.Data)
	} else if a.Literal != nil && len(a.Literal.Data) > 0 {
		w.WriteString(a.Literal.Data[:1])
		w.Underlying().WriteString(a.Literal.Data[1:])
	} else if a.Enclosure != nil {
		a.Enclosure.printSource(w, v)
	}
}

func (a AugmentedAssignmentStatement) printSource(w writer, v bool) {
	a.AugTarget.printSource(w, v)

	if v {
		w.WriteString(" ")
	}

	w.WriteString(a.AugOp.Data)

	if v {
		w.WriteString(" ")
	}

	if a.ExpressionList != nil {
		a.ExpressionList.printSource(w, v)
	} else if a.YieldExpression != nil {
		a.YieldExpression.printSource(w, v)
	}
}

func (a AugTarget) printSource(w writer, v bool) {
	a.PrimaryExpression.printSource(w, v)
}

func (c ClassDefinition) printSource(w writer, v bool) {
	if c.Decorators != nil {
		c.Decorators.printSource(w, v)

		if v {
			c.Comments[0].printSource(w, true)
		}
	}

	w.WriteString("class ")
	w.WriteString(c.ClassName.Data)

	if c.TypeParams != nil {
		c.TypeParams.printSource(w, v)
	}

	if c.Inheritance != nil {

		w.WriteString("(")

		if v {
			c.Comments[1].printSource(w, true)
		}

		c.Inheritance.printSource(w.IndentMultiline(), v)

		if v && len(c.Comments[2]) > 0 {
			w.WriteString("\n")
			c.Comments[2].printSource(w, true)
		}

		w.WriteString(")")
	}

	w.WriteString(":")
	c.Suite.printSource(w, v)
}

func (c Comparison) printSource(w writer, v bool) {
	c.OrExpression.printSource(w, v)

	for _, ce := range c.Comparisons {
		ce.printSource(w, v)
	}
}

func (c ComparisonExpression) printSource(w writer, v bool) {
	var first string

	if len(c.ComparisonOperator) > 0 {
		first = c.ComparisonOperator[0].Data

		if v && w.InMultiline() {
			c.Comments[0].printSource(w, true)
		}
	}

	switch first {
	case "<", ">", "==", ">=", "<=", "!=":
		if v && (!w.InMultiline() || len(c.Comments[0]) == 0) {
			w.WriteString(" ")
		}

		w.WriteString(first)

		if v {
			w.WriteString(" ")

			if w.InMultiline() {
				c.Comments[2].printSource(w, true)
			}
		}

		c.OrExpression.printSource(w, v)
	case "in":
		if !v || !w.InMultiline() || len(c.Comments[0]) == 0 {
			w.WriteString(" ")
		}

		w.WriteString("in ")

		if v && w.InMultiline() {
			c.Comments[2].printSource(w, true)
		}

		c.OrExpression.printSource(w, v)
	case "is":
		if !v || !w.InMultiline() || len(c.Comments[0]) == 0 {
			w.WriteString(" ")
		}

		w.WriteString("is ")

		if v && w.InMultiline() {
			c.Comments[1].printSource(w, true)
		}

		if c.ComparisonOperator[len(c.ComparisonOperator)-1].Data == "not" {
			w.WriteString("not ")
		}

		if v && w.InMultiline() {
			c.Comments[2].printSource(w, true)
		}

		c.OrExpression.printSource(w, v)
	case "not":
		if !v || !w.InMultiline() || len(c.Comments[0]) == 0 {
			w.WriteString(" ")
		}

		w.WriteString("not ")

		if v && w.InMultiline() {
			c.Comments[1].printSource(w, true)
		}

		if c.ComparisonOperator[len(c.ComparisonOperator)-1].Data == "in" {
			w.WriteString("in ")

			if v && w.InMultiline() {
				c.Comments[2].printSource(w, true)
			}

			c.OrExpression.printSource(w, v)
		}
	}
}

func (c CompoundStatement) printSource(w writer, v bool) {
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

func (c Comprehension) printSource(w writer, v bool) {
	if v && len(c.Comments[0]) > 0 {
		w.WriteString("\n")
		c.Comments[0].printSource(w, true)
	}

	c.AssignmentExpression.printSource(w, v)
	w.WriteString(" ")

	if v {
		c.Comments[1].printSource(w, true)
	}

	c.ComprehensionFor.printSource(w, v)

	if v {
		c.Comments[2].printSource(w, true)
	}
}

func (c ComprehensionFor) printSource(w writer, v bool) {
	if c.Async {
		w.WriteString("async ")

		if v && w.InMultiline() {
			c.Comments[0].printSource(w, true)
		}
	}

	w.WriteString("for ")

	c.TargetList.printSource(w, v)

	if w.LastChar() == '\n' {
		w.WriteString("in ")
	} else {
		w.WriteString(" in ")
	}

	if v && w.InMultiline() {
		c.Comments[1].printSource(w, true)
	}

	c.OrTest.printSource(w, v)

	if c.ComprehensionIterator != nil {
		w.WriteString(" ")
		c.ComprehensionIterator.printSource(w, v)
	}
}

func (c ComprehensionIf) printSource(w writer, v bool) {
	w.WriteString("if ")

	if v && w.InMultiline() {
		c.Comments.printSource(w, true)
	}

	c.OrTest.printSource(w, v)

	if c.ComprehensionIterator != nil {
		w.WriteString(" ")
		c.ComprehensionIterator.printSource(w, v)
	}
}

func (c ComprehensionIterator) printSource(w writer, v bool) {
	if v {
		c.Comments[0].printSource(w, true)
	}

	if c.ComprehensionIf != nil {
		c.ComprehensionIf.printSource(w, v)
	} else if c.ComprehensionFor != nil {
		c.ComprehensionFor.printSource(w, v)
	}

	if v {
		c.Comments[1].printSource(w, true)
	}
}

func (c ConditionalExpression) printSource(w writer, v bool) {
	c.OrTest.printSource(w, v)

	if c.If != nil && c.Else != nil {
		w.WriteString(" ")

		if v && w.InMultiline() {
			c.Comments[0].printSource(w, true)
		}

		w.WriteString("if ")

		if v && w.InMultiline() {
			c.Comments[1].printSource(w, true)
		}

		c.If.printSource(w, v)
		w.WriteString(" ")

		if v && w.InMultiline() {
			c.Comments[2].printSource(w, true)
		}

		w.WriteString("else ")

		if v && w.InMultiline() {
			c.Comments[3].printSource(w, true)
		}

		c.Else.printSource(w, v)
	}
}

func (d Decorators) printSource(w writer, v bool) {
	for _, dc := range d.Decorators {
		dc.printSource(w, v)
	}
}

func (d Decorator) printSource(w writer, v bool) {
	if v {
		d.Comments[0].printSource(w, true)
	}

	w.WriteString("@")
	d.Decorator.printSource(w, v)

	if v {
		d.Comments[1].printSource(w, true)
	}

	w.WriteString("\n")
}

func (d DefParameter) printSource(w writer, v bool) {
	if v {
		d.Comments[0].printSource(w, true)
	}

	d.Parameter.printSource(w, v)

	if d.Value != nil {
		if v {
			w.WriteString(" = ")
		} else {
			w.WriteString("=")
		}

		d.Value.printSource(w, v)
	}

	if v {
		d.Comments[1].printSource(w, true)
	}
}

func (d DelStatement) printSource(w writer, v bool) {
	w.WriteString("del ")
	d.TargetList.printSource(w, v)
}

func (d DictDisplay) printSource(w writer, v bool) {
	if len(d.DictItems) > 0 {
		d.DictItems[0].printSource(w, v)

		for _, di := range d.DictItems[1:] {
			if v {
				w.WriteString(", ")
			} else {
				w.WriteString(",")
			}

			di.printSource(w, v)
		}
	}

	if d.DictComprehension != nil {
		w.WriteString(" ")
		d.DictComprehension.printSource(w, v)
	}
}

func (d DictItem) printSource(w writer, v bool) {
	if v && w.InMultiline() && len(d.Comments[0]) > 0 {
		w.WriteString("\n")
		d.Comments[0].printSource(w, true)
	}

	if d.OrExpression != nil {
		w.WriteString("**")

		if v && w.InMultiline() {
			d.Comments[1].printSource(w, true)
		}

		d.OrExpression.printSource(w, v)
	} else if d.Key != nil && d.Value != nil {
		d.Key.printSource(w, v)

		if v {
			if w.InMultiline() {
				d.Comments[1].printSource(w, true)
			}

			w.WriteString(": ")

			if w.InMultiline() {
				d.Comments[2].printSource(w, true)
			}
		} else {
			w.WriteString(":")
		}

		d.Value.printSource(w, v)
	}

	if v && w.InMultiline() {
		d.Comments[3].printSource(w, true)
	}
}

func (e Enclosure) printSource(w writer, v bool) {
	if e.ParenthForm != nil || e.ListDisplay != nil || e.DictDisplay != nil || e.SetDisplay != nil || e.GeneratorExpression != nil || e.YieldAtom != nil {
		var (
			t  formatter
			oc string
		)

		if e.ParenthForm != nil {
			t = e.ParenthForm
			oc = "()"
		} else if e.ListDisplay != nil {
			t = e.ListDisplay
			oc = "[]"
		} else if e.DictDisplay != nil {
			t = e.DictDisplay
			oc = "{}"
		} else if e.SetDisplay != nil {
			t = e.SetDisplay
			oc = "{}"
		} else if e.GeneratorExpression != nil {
			t = e.GeneratorExpression
			oc = "()"
		} else if e.YieldAtom != nil {
			t = e.YieldAtom
			oc = "()"
		}

		ip := w.IndentMultiline()

		w.WriteString(oc[:1])

		if v && len(e.Comments[0]) > 0 {
			e.Comments[0].printSource(w, false)
			ip.WriteString("\n")
		}

		t.printSource(ip, v)

		if v && len(e.Comments[1]) > 0 {
			w.WriteString("\n")
			e.Comments[1].printSource(w, true)
		}

		w.WriteString(oc[1:])
	}
}

func (e Except) printSource(w writer, v bool) {
	e.Expression.printSource(w, v)

	if e.Identifier != nil {
		w.WriteString(" as ")
		w.WriteString(e.Identifier.Data)
	}

	w.WriteString(":")
	e.Suite.printSource(w, v)
}

func (e Expression) printSource(w writer, v bool) {
	if e.LambdaExpression != nil {
		e.LambdaExpression.printSource(w, v)
	} else if e.ConditionalExpression != nil {
		e.ConditionalExpression.printSource(w, v)
	}
}

func (e ExpressionList) printSource(w writer, v bool) {
	if len(e.Expressions) > 0 {
		e.Expressions[0].printSource(w, v)

		for _, ex := range e.Expressions[1:] {
			if v {
				w.WriteString(", ")
			} else {
				w.WriteString(",")
			}

			ex.printSource(w, v)
		}
	}
}

func (f File) printSource(w writer, v bool) {
	if v {
		f.Comments[0].printSource(w, true)
	}

	printStatements(w, v, f.Statements)
	w.WriteString("\n")

	if v && len(f.Comments[1]) > 0 {
		w.WriteString("\n")
		f.Comments[1].printSource(w, true)
	}
}

func printStatements(w writer, v bool, s []Statement) {
	if len(s) > 0 {
		s[0].printSource(w, v)

		lastLine := lastTokenPos(s[0].Tokens)

		for _, s := range s[1:] {
			if v && firstTokenPos(s.Tokens) > lastLine+1 {
				w.WriteString("\n")
			}

			w.WriteString("\n")
			s.printSource(w, v)

			lastLine = lastTokenPos(s.Tokens)
		}
	}
}

func firstTokenPos(tk Tokens) (pos uint64) {
	if len(tk) > 0 {
		pos = tk[0].Line
	}

	return pos
}

func lastTokenPos(tk Tokens) (pos uint64) {
	if len(tk) > 0 {
		pos = tk[len(tk)-1].Line
	}

	return pos
}

func (f FlexibleExpressionListOrComprehension) printSource(w writer, v bool) {
	if f.FlexibleExpressionList != nil {
		f.FlexibleExpressionList.printSource(w, v)
	} else if f.Comprehension != nil {
		f.Comprehension.printSource(w, v)
	}
}

func (f FlexibleExpressionList) printSource(w writer, v bool) {
	if len(f.FlexibleExpressions) > 0 {
		f.FlexibleExpressions[0].printSource(w, v)
		for _, fe := range f.FlexibleExpressions[1:] {
			if v {
				w.WriteString(", ")
			} else {
				w.WriteString(",")
			}

			fe.printSource(w, v)
		}
	}
}

func (f FlexibleExpression) printSource(w writer, v bool) {
	if v && len(f.Comments[0]) > 0 {
		w.WriteString("\n")
		f.Comments[0].printSource(w, true)
	}

	if f.AssignmentExpression != nil {
		f.AssignmentExpression.printSource(w, v)
	} else if f.StarredExpression != nil {
		w.WriteString("*")
		f.StarredExpression.printSource(w, v)
	}

	if v {
		f.Comments[1].printSource(w, true)
	}
}

func (f ForStatement) printSource(w writer, v bool) {
	if f.Async {
		w.WriteString("async ")
	}

	w.WriteString("for ")
	f.TargetList.printSource(w, v)
	w.WriteString(" in ")
	f.StarredList.printSource(w, v)
	w.WriteString(":")
	f.Suite.printSource(w, v)

	if f.Else != nil {
		w.WriteString("\nelse:")
		f.Else.printSource(w, v)
	}
}

func (f FuncDefinition) printSource(w writer, v bool) {
	if f.Decorators != nil && len(f.Decorators.Decorators) > 0 {
		f.Decorators.printSource(w, v)

		if v {
			f.Comments[0].printSource(w, true)
		}
	}

	if f.Async {
		w.WriteString("async ")
	}

	w.WriteString("def ")
	w.WriteString(f.FuncName.Data)

	if f.TypeParams != nil {
		f.TypeParams.printSource(w, v)
	}

	w.WriteString("(")

	if v {
		f.Comments[1].printSource(w, true)

		if len(f.Comments[2]) > 0 {
			ip := w.Indent()

			ip.WriteString("\n")
			f.Comments[2].printSource(ip, true)
		}
	}

	f.ParameterList.printSource(w.IndentMultiline(), v)
	w.WriteString(")")

	if f.Expression != nil {
		if v {
			w.WriteString(" -> ")
		} else {
			w.WriteString("->")
		}
		f.Expression.printSource(w, v)
	}

	w.WriteString(":")
	f.Suite.printSource(w, v)
}

func (g GeneratorExpression) printSource(w writer, v bool) {
	if v && len(g.Comments[0]) > 0 {
		w.WriteString("\n")
		g.Comments[0].printSource(w, true)
	}

	g.Expression.printSource(w, v)
	w.WriteString(" ")

	if v && len(g.Comments[1]) > 0 {
		g.Comments[1].printSource(w, true)
	}

	g.ComprehensionFor.printSource(w, v)

	if v {
		g.Comments[2].printSource(w, true)
	}
}

func (g GlobalStatement) printSource(w writer, v bool) {
	if len(g.Identifiers) > 0 {
		w.WriteString("global ")
		w.WriteString(g.Identifiers[0].Data)

		for _, t := range g.Identifiers[1:] {
			if v {
				w.WriteString(", ")
			} else {
				w.WriteString(",")
			}

			w.WriteString(t.Data)
		}
	}
}

func (i IfStatement) printSource(w writer, v bool) {
	w.WriteString("if ")
	i.AssignmentExpression.printSource(w, v)
	w.WriteString(":")
	i.Suite.printSource(w, v)

	for _, e := range i.Elif {
		w.WriteString("\nelif ")
		e.printSource(w, v)
	}

	if i.Else != nil {
		w.WriteString("\nelse:")
		i.Else.printSource(w, v)
	}
}

func (i ImportStatement) printSource(w writer, v bool) {
	if i.RelativeModule != nil {
		w.WriteString("from ")
		i.RelativeModule.printSource(w, v)
		w.WriteString(" import ")
	} else {
		w.WriteString("import ")
	}

	if len(i.Modules) > 0 {
		i.Modules[0].printSource(w, v)

		for _, m := range i.Modules[1:] {
			if v {
				w.WriteString(", ")
			} else {
				w.WriteString(",")
			}

			m.printSource(w, v)
		}
	}
}

func (k KeywordArgument) printSource(w writer, v bool) {
	if v && len(k.Comments[0]) > 0 {
		w.WriteString("\n")
		k.Comments[0].printSource(w, true)
	}

	if k.Expression != nil {
		w.WriteString("**")

		if v {
			k.Comments[1].printSource(w, true)
		}

		k.Expression.printSource(w, v)
	} else if k.KeywordItem != nil {
		k.KeywordItem.printSource(w, v)
	}

	if v {
		k.Comments[2].printSource(w, true)
	}
}

func (k KeywordItem) printSource(w writer, v bool) {
	if k.Identifier != nil {
		w.WriteString(k.Identifier.Data)

		if v {
			w.WriteString(" = ")
		} else {
			w.WriteString("=")
		}

		k.Expression.printSource(w, v)
	}
}

func (l LambdaExpression) printSource(w writer, v bool) {
	if l.ParameterList != nil {
		w.WriteString("lambda ")

		if v && w.InMultiline() {
			l.Comments[0].printSource(w, true)
		}

		l.ParameterList.printSource(w, v)

		if v && w.InMultiline() && len(l.Comments[2]) > 0 {
			w.WriteString("\n")
			l.Comments[1].printSource(w, true)
		}
	} else {
		w.WriteString("lambda")

		if v && w.InMultiline() {
			l.Comments[0].printSource(w, true)
		}
	}

	if v {
		w.WriteString(": ")

		if w.InMultiline() {
			l.Comments[2].printSource(w, true)
		}
	} else {
		w.WriteString(":")
	}

	l.Expression.printSource(w, v)
}

func (m ModuleAs) printSource(w writer, v bool) {
	m.Module.printSource(w, v)

	if m.As != nil {
		w.WriteString(" as ")
		w.WriteString(m.As.Data)
	}
}

func (m Module) printSource(w writer, v bool) {
	if len(m.Identifiers) > 0 {
		w.WriteString(m.Identifiers[0].Data)

		for _, i := range m.Identifiers[1:] {
			w.WriteString(".")
			w.WriteString(i.Data)
		}
	}
}

func (m MultiplyExpression) printSource(w writer, v bool) {
	m.UnaryExpression.printSource(w, v)

	if m.Multiply != nil && m.MultiplyExpression != nil {
		if v {
			w.WriteString(" ")

			if w.InMultiline() {
				m.Comments[0].printSource(w, true)
			}

			w.WriteString(m.Multiply.Data)
			w.WriteString(" ")

			if w.InMultiline() {
				m.Comments[1].printSource(w, true)
			}
		} else {
			w.WriteString(m.Multiply.Data)
		}

		m.MultiplyExpression.printSource(w, v)
	}
}

func (n NonLocalStatement) printSource(w writer, v bool) {
	if len(n.Identifiers) > 0 {
		w.WriteString("nonlocal ")
		w.WriteString(n.Identifiers[0].Data)

		for _, t := range n.Identifiers[1:] {
			if v {
				w.WriteString(", ")
			} else {
				w.WriteString(",")
			}

			w.WriteString(t.Data)
		}
	}
}

func (n NotTest) printSource(w writer, v bool) {
	for _, comments := range n.Nots {
		w.WriteString("not ")

		if v && w.InMultiline() {
			comments.printSource(w, true)
		}
	}

	n.Comparison.printSource(w, v)
}

func (o OrExpression) printSource(w writer, v bool) {
	o.XorExpression.printSource(w, v)

	if o.OrExpression != nil {
		if v {
			w.WriteString(" ")

			if w.InMultiline() {
				o.Comments[0].printSource(w, true)
			}

			w.WriteString("| ")

			if w.InMultiline() {
				o.Comments[1].printSource(w, true)
			}
		} else {
			w.WriteString("|")
		}

		o.OrExpression.printSource(w, v)
	}
}

func (o OrTest) printSource(w writer, v bool) {
	o.AndTest.printSource(w, v)

	if o.OrTest != nil {
		w.WriteString(" ")

		if v && w.InMultiline() {
			o.Comments[0].printSource(w, true)
		}

		w.WriteString("or ")

		if v && w.InMultiline() {
			o.Comments[1].printSource(w, true)
		}

		o.OrTest.printSource(w, v)
	}
}

func (p Parameter) printSource(w writer, v bool) {
	w.WriteString(p.Identifier.Data)

	if p.Type != nil {
		if v {
			p.Comments[0].printSource(w, true)
			w.WriteString(": ")
			p.Comments[1].printSource(w, true)
		} else {
			w.WriteString(":")
		}

		p.Type.printSource(w, v)
	}
}

func (p ParameterList) printSource(w writer, v bool) {
	first := len(p.DefParameters) == 0

	if v && w.InMultiline() {
		p.Comments[0].printSource(w, true)
	}

	if !first {
		if v && w.InMultiline() && len(p.DefParameters[0].Comments[0]) > 0 {
			w.WriteString("\n")
		}

		for _, d := range p.DefParameters {
			d.printSource(w, v)

			if v {
				w.WriteString(", ")
			} else {
				w.WriteString(",")
			}
		}

		if v {
			p.Comments[1].printSource(w, true)
		}

		w.WriteString("/")

		if v && w.InMultiline() {
			p.Comments[2].printSource(w, true)
		}
	}

	for _, n := range p.NoPosOnly {
		if first {
			first = false

			if v && w.InMultiline() && len(p.NoPosOnly[0].Comments[0]) > 0 {
				w.WriteString("\n")
			}
		} else if v {
			w.WriteString(", ")
		} else {
			w.WriteString(",")
		}

		n.printSource(w, v)
	}

	if p.StarArg != nil {
		if first {
			first = false

			if v && w.InMultiline() && len(p.Comments[3]) > 0 {
				w.WriteString("\n")
			}
		} else if v {
			w.WriteString(", ")
		} else {
			w.WriteString(",")
		}

		if v && len(p.Comments[3]) > 0 {
			p.Comments[3].printSource(w, true)
		}

		w.WriteString("*")

		if v && len(p.Comments[4]) > 0 {
			w.WriteString(" ")
			p.Comments[4].printSource(w, true)
		}

		p.StarArg.printSource(w, v)

		if v && w.InMultiline() {
			p.Comments[5].printSource(w, true)
		}

		for _, d := range p.StarArgs {
			if v {
				w.WriteString(", ")
			} else {
				w.WriteString(",")
			}

			d.printSource(w, v)
		}
	}

	if p.StarStarArg != nil {
		if !first {
			if v {
				w.WriteString(", ")
			} else {
				w.WriteString(",")
			}
		} else if v && w.InMultiline() && len(p.Comments[6]) > 0 {
			w.WriteString("\n")
		}

		if v && len(p.Comments[6]) > 0 {
			p.Comments[6].printSource(w, true)
		}

		w.WriteString("**")

		if v {
			p.Comments[7].printSource(w, true)
		}

		p.StarStarArg.printSource(w, v)

		if v && w.InMultiline() {
			p.Comments[8].printSource(w, true)
		}
	}

	if v && len(p.Comments[9]) > 0 {
		w.WriteString("\n")
		p.Comments[9].printSource(w, true)
	}
}

func (p PositionalArgument) printSource(w writer, v bool) {
	if v {
		p.Comments[0].printSource(w, true)
	}

	if p.AssignmentExpression != nil {
		p.AssignmentExpression.printSource(w, v)
	} else if p.Expression != nil {
		w.WriteString("*")

		if v {
			p.Comments[1].printSource(w, true)
		}

		p.Expression.printSource(w, v)
	}

	if v {
		p.Comments[2].printSource(w, true)
	}
}

func (p PowerExpression) printSource(w writer, v bool) {
	if p.AwaitExpression {
		w.WriteString("await ")

		if v && w.InMultiline() {
			p.Comments[0].printSource(w, true)
		}
	}

	p.PrimaryExpression.printSource(w, v)

	if p.UnaryExpression != nil {
		if v {
			w.WriteString(" ")

			if w.InMultiline() {
				p.Comments[1].printSource(w, true)
			}

			w.WriteString("** ")

			if w.InMultiline() {
				p.Comments[2].printSource(w, true)
			}
		} else {
			w.WriteString("**")
		}

		p.UnaryExpression.printSource(w, v)
	}
}

func (p PrimaryExpression) printSource(w writer, v bool) {
	if p.Atom != nil {
		p.Atom.printSource(w, v)
	} else if p.PrimaryExpression != nil {
		p.PrimaryExpression.printSource(w, v)

		if p.AttributeRef != nil {
			if v && w.InMultiline() {
				p.Comments[0].printSource(w, true)
			}

			w.WriteString(".")

			if v && w.InMultiline() {
				p.Comments[1].printSource(w, true)
			}

			w.WriteString(p.AttributeRef.Data)
		} else if p.Slicing != nil {
			if v && w.InMultiline() {
				p.Comments[0].printSource(w, true)
			}

			p.Slicing.printSource(w, v)
		} else if p.Call != nil {
			if v && w.InMultiline() {
				p.Comments[0].printSource(w, true)
			}

			p.Call.printSource(w, v)
		}
	}
}

func (r RaiseStatement) printSource(w writer, v bool) {
	if r.Expression == nil {
		w.WriteString("raise")
	} else {
		w.WriteString("raise ")
		r.Expression.printSource(w, v)

		if r.From != nil {
			w.WriteString(" from ")
			r.From.printSource(w, v)
		}
	}
}

func (r RelativeModule) printSource(w writer, v bool) {
	for range r.Dots {
		w.WriteString(".")
	}

	if r.Module != nil {
		r.Module.printSource(w, v)
	}
}

func (r ReturnStatement) printSource(w writer, v bool) {
	if r.Expression != nil {
		w.WriteString("return ")
		r.Expression.printSource(w, v)
	} else {
		w.WriteString("return")
	}
}

func (s ShiftExpression) printSource(w writer, v bool) {
	s.AddExpression.printSource(w, v)

	if s.Shift != nil && s.ShiftExpression != nil {
		if v {
			w.WriteString(" ")

			if w.InMultiline() {
				s.Comments[0].printSource(w, true)
			}

			w.WriteString(s.Shift.Data)
			w.WriteString(" ")

			if w.InMultiline() {
				s.Comments[1].printSource(w, true)
			}
		} else {
			w.WriteString(s.Shift.Data)
		}

		s.ShiftExpression.printSource(w, v)
	}
}

func (s SimpleStatement) printSource(w writer, v bool) {
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
		w.WriteString("pass")
	} else if s.Type == StatementBreak {
		w.WriteString("break")
	} else if s.Type == StatementContinue {
		w.WriteString("continue")
	}
}

func (s SliceItem) printSource(w writer, v bool) {
	if s.Expression != nil {
		if v && w.InMultiline() && len(s.Comments[0]) > 0 {
			w.WriteString("\n")
			s.Comments[0].printSource(w, true)
		}

		s.Expression.printSource(w, v)

		if v && w.InMultiline() {
			s.Comments[5].printSource(w, true)
		}
	} else if s.LowerBound != nil {
		if v && w.InMultiline() && len(s.Comments[0]) > 0 {
			w.WriteString("\n")
			s.Comments[0].printSource(w, true)
		}

		s.LowerBound.printSource(w, v)

		if s.UpperBound != nil {
			if v {
				w.WriteString(" ")

				if w.InMultiline() {
					s.Comments[1].printSource(w, true)
				}

				w.WriteString(": ")

				if w.InMultiline() {
					s.Comments[2].printSource(w, true)
				}
			} else {
				w.WriteString(":")
			}

			s.UpperBound.printSource(w, v)

			if s.Stride != nil {
				if v {
					w.WriteString(" ")

					if w.InMultiline() {
						s.Comments[3].printSource(w, true)
					}

					w.WriteString(": ")

					if w.InMultiline() {
						s.Comments[4].printSource(w, true)
					}
				} else {
					w.WriteString(":")
				}

				s.Stride.printSource(w, v)
			}
		}

		if v && w.InMultiline() {
			s.Comments[5].printSource(w, true)
		}
	}
}

func (s SliceList) printSource(w writer, v bool) {
	w.WriteString("[")

	ip := w.IndentMultiline()

	if v && len(s.Comments[0]) > 0 {
		s.Comments[0].printSource(w, false)
		ip.WriteString("\n")
	}

	if len(s.SliceItems) > 0 {
		s.SliceItems[0].printSource(ip, v)

		for _, si := range s.SliceItems[1:] {
			if v {
				ip.WriteString(", ")
			} else {
				ip.WriteString(",")
			}

			si.printSource(ip, v)
		}
	}

	if v && len(s.Comments[1]) > 0 {
		w.WriteString("\n")
		s.Comments[1].printSource(w, true)
	}

	w.WriteString("]")
}

func (s StarredExpression) printSource(w writer, v bool) {
	if v && w.InMultiline() && len(s.Comments[0]) > 0 {
		w.WriteString("\n")
		s.Comments[0].printSource(w, true)
	}

	if s.Expression != nil {
		s.Expression.printSource(w, v)
	} else if s.StarredList != nil {
		s.StarredList.printSource(w, v)
	}

	if v && w.InMultiline() {
		s.Comments[1].printSource(w, true)
	}
}

func (s StarredItem) printSource(w writer, v bool) {
	if v && w.InMultiline() {
		s.Comments[0].printSource(w, true)
	}

	if s.AssignmentExpression != nil {
		s.AssignmentExpression.printSource(w, v)
	} else if s.OrExpr != nil {
		w.WriteString("*")

		if v && w.InMultiline() {
			s.Comments[1].printSource(w, true)
		}

		s.OrExpr.printSource(w, v)
	}

	if v && w.InMultiline() {
		s.Comments[2].printSource(w, true)
	}
}

func (s StarredList) printSource(w writer, v bool) {
	if len(s.StarredItems) > 0 {
		s.StarredItems[0].printSource(w, v)

		for _, si := range s.StarredItems[1:] {
			if v {
				w.WriteString(", ")
			} else {
				w.WriteString(",")
			}

			si.printSource(w, v)
		}

		if s.TrailingComma {
			w.WriteString(",")
		}
	}
}

func (s StarredOrKeyword) printSource(w writer, v bool) {
	if v && len(s.Comments[0]) > 0 {
		w.WriteString("\n")
		s.Comments[0].printSource(w, true)
	}

	if s.Expression != nil {
		w.WriteString("*")
		s.Expression.printSource(w, v)
	} else if s.KeywordItem != nil {
		s.KeywordItem.printSource(w, v)
	}

	if v {
		s.Comments[1].printSource(w, true)
	}
}

func (s Statement) printSource(w writer, v bool) {
	if v {
		s.Comments.printSource(w, true)
	}

	if s.StatementList != nil {
		s.StatementList.printSource(w, v)
	} else if s.CompoundStatement != nil {
		s.CompoundStatement.printSource(w, v)
	}
}

func (s StatementList) printSource(w writer, v bool) {
	if len(s.Statements) > 0 {
		s.Statements[0].printSource(w, v)

		for _, ss := range s.Statements[1:] {
			if v {
				w.WriteString("; ")
			} else {
				w.WriteString(";")
			}

			ss.printSource(w, v)
		}

		if v {
			s.Comments.printSource(w, false)
		}
	}
}

func (s Suite) printSource(w writer, v bool) {
	if s.StatementList != nil {
		if v {
			w.WriteString(" ")
		}

		s.StatementList.printSource(w, v)
	} else {
		ip := w.Indent()
		if v && len(s.Comments[0]) > 0 {
			if len(s.Tokens) > 0 && len(s.Comments[0]) > 0 && s.Comments[0][0].Line > s.Tokens[0].Line {
				ip.WriteString("\n")
				s.Comments[0].printSource(ip, false)
			} else {
				s.Comments[0].printSource(w, false)
			}
		}

		ip.WriteString("\n")
		printStatements(ip, v, s.Statements)

		if v && len(s.Comments[1]) > 0 {
			ip.WriteString("\n\n")
			s.Comments[1].printSource(ip, false)
		}
	}
}

func (t Target) printSource(w writer, v bool) {
	if v && w.InMultiline() && len(t.Comments[0]) > 0 {
		t.Comments[0].printSource(w, true)
	}

	if t.PrimaryExpression != nil {
		t.PrimaryExpression.printSource(w, v)
	} else if t.Tuple != nil {
		w.WriteString("(")
		t.Tuple.printSource(w.IndentMultiline(), v)
		w.WriteString(")")
	} else if t.Array != nil {
		w.WriteString("[")
		t.Array.printSource(w.IndentMultiline(), v)
		w.WriteString("]")
	} else if t.Star != nil {
		w.WriteString("*")
		t.Star.printSource(w, v)
	}

	if v && w.InMultiline() {
		t.Comments[1].printSource(w, true)
	}
}

func (t TargetList) printSource(w writer, v bool) {
	if v && len(t.Comments[0]) > 0 {
		t.Comments[0].printSource(w, true)
	}

	if len(t.Targets) > 0 {
		t.Targets[0].printSource(w, v)

		for _, tg := range t.Targets[1:] {
			if v {
				w.WriteString(", ")
			} else {
				w.WriteString(",")
			}

			tg.printSource(w, v)
		}

		if v {
			t.Comments[1].printSource(w, true)
		}
	}
}

func (t TryStatement) printSource(w writer, v bool) {
	w.WriteString("try:")
	t.Try.printSource(w, v)

	if len(t.Except) > 0 {
		w.WriteString("\nexcept ")

		if t.Groups {
			w.WriteString("*")
		}

		t.Except[0].printSource(w, v)

		for _, e := range t.Except[1:] {
			w.WriteString("\nexcept ")

			if t.Groups {
				w.WriteString("*")
			}

			e.printSource(w, v)
		}
	}

	if t.Else != nil {
		w.WriteString("\nelse:")
		t.Else.printSource(w, v)
	}

	if t.Finally != nil {
		w.WriteString("\nfinally:")
		t.Finally.printSource(w, v)
	}
}

func (t TypeParam) printSource(w writer, v bool) {
	if v {
		t.Comments[0].printSource(w, true)
	}

	if t.Type == TypeParamVar {
		w.WriteString("*")

		if v {
			t.Comments[1].printSource(w, true)
		}
	} else if t.Type == TypeParamVarTuple {
		w.WriteString("**")

		if v {
			t.Comments[1].printSource(w, true)
		}
	}

	w.WriteString(t.Identifier.Data)

	if t.Expression != nil {
		if v {
			t.Comments[1].printSource(w, true)
			w.WriteString(": ")
			t.Comments[2].printSource(w, true)
		} else {
			w.WriteString(":")
		}

		t.Expression.printSource(w, v)
	}

	if v {
		t.Comments[3].printSource(w, true)
	}
}

func (t TypeParams) printSource(w writer, v bool) {
	ip := w.Indent()

	w.WriteString("[")

	if v && len(t.Comments[0]) > 0 {
		w.WriteString(" ")
		t.Comments[0].printSource(w, true)

		if len(t.TypeParams) > 0 && len(t.TypeParams[0].Comments) > 0 {
			ip.WriteString("\n")
		}
	}

	if len(t.TypeParams) > 0 {
		t.TypeParams[0].printSource(ip, v)

		for _, tp := range t.TypeParams[1:] {
			if v {
				ip.WriteString(", ")
			} else {
				ip.WriteString(",")
			}

			tp.printSource(ip, v)
		}
	}

	if v && len(t.Comments[1]) > 0 {
		w.WriteString("\n")
		t.Comments[1].printSource(w, true)
	}

	w.WriteString("]")
}

func (t TypeStatement) printSource(w writer, v bool) {
	w.WriteString("type ")
	w.WriteString(t.Identifier.Data)

	if t.TypeParams != nil {
		t.TypeParams.printSource(w, v)
	}

	if v {
		w.WriteString(" = ")
	} else {
		w.WriteString("=")
	}

	t.Expression.printSource(w, v)
}

func (u UnaryExpression) printSource(w writer, v bool) {
	if u.PowerExpression != nil {
		u.PowerExpression.printSource(w, v)
	} else if u.Unary != nil && u.UnaryExpression != nil {
		w.WriteString(u.Unary.Data)

		if v && w.InMultiline() {
			u.Comments.printSource(w, true)
		}

		u.UnaryExpression.printSource(w, v)
	}
}

func (ws WhileStatement) printSource(w writer, v bool) {
	w.WriteString("while ")
	ws.AssignmentExpression.printSource(w, v)
	w.WriteString(":")
	ws.Suite.printSource(w, v)

	if ws.Else != nil {
		w.WriteString("\nelse:")
		ws.Else.printSource(w, v)
	}
}

func (wi WithItem) printSource(w writer, v bool) {
	if v && w.InMultiline() && len(wi.Comments[0]) > 0 {
		w.WriteString("\n")
		wi.Comments[0].printSource(w, true)
	}

	wi.Expression.printSource(w, v)

	if wi.Target != nil {
		w.WriteString(" ")

		if v && w.InMultiline() {
			wi.Comments[1].printSource(w, true)
		}

		w.WriteString("as ")

		wi.Target.printSource(w, v)
	} else if v && w.InMultiline() {
		wi.Comments[1].printSource(w, true)
	}
}

func (ws WithStatement) printSource(w writer, v bool) {
	w.WriteString("with ")
	ws.Contents.printSource(w, v)
	w.WriteString(":")
	ws.Suite.printSource(w, v)
}

func (wc WithStatementContents) printSource(w writer, v bool) {
	if len(wc.Items) > 0 {
		parens := v && (len(wc.Comments[0]) > 0 || len(wc.Comments[1]) > 0)

		ip := w

		if parens {
			w.WriteString("(")

			if v {
				wc.Comments[0].printSource(w, true)
			}

			ip = w.IndentMultiline()
		}

		wc.Items[0].printSource(ip, v)

		for _, i := range wc.Items[1:] {
			if v {
				ip.WriteString(", ")
			} else {
				ip.WriteString(",")
			}

			i.printSource(ip, v)
		}

		if parens {
			if len(wc.Comments[1]) > 0 {
				w.WriteString("\n")
				wc.Comments[1].printSource(w, true)
			}

			w.WriteString(")")
		}
	}
}

func (x XorExpression) printSource(w writer, v bool) {
	x.AndExpression.printSource(w, v)

	if x.XorExpression != nil {
		if v {
			w.WriteString(" ")

			if w.InMultiline() {
				x.Comments[0].printSource(w, true)
			}

			w.WriteString("^ ")

			if w.InMultiline() {
				x.Comments[1].printSource(w, true)
			}
		} else {
			w.WriteString("^")
		}

		x.XorExpression.printSource(w, v)
	}
}

func (y YieldExpression) printSource(w writer, v bool) {
	if y.From != nil || y.ExpressionList != nil {
		if v && w.InMultiline() && len(y.Comments[0]) > 0 {
			w.WriteString("\n")
			y.Comments[0].printSource(w, true)
		}

		w.WriteString("yield ")

		if v && w.InMultiline() {
			y.Comments[1].printSource(w, true)
		}

		if y.From != nil {
			w.WriteString("from ")

			if v && w.InMultiline() {
				y.Comments[2].printSource(w, true)
			}

			y.From.printSource(w, v)
		} else if y.ExpressionList != nil {
			y.ExpressionList.printSource(w, v)

			if v && w.InMultiline() && len(y.Comments[2]) > 0 {
				y.Comments[2].printSource(w, true)
				w.WriteString(",")
			}
		}

		if v && w.InMultiline() {
			y.Comments[3].printSource(w, true)
		}
	}
}

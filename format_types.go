package python

// File automatically generated with format.sh.

func (f *AddExpression) printType(w writer, v bool) {
	pp := w.Indent()

	pp.WriteString("AddExpression {")

	pp.WriteString("\nMultiplyExpression: ")
	f.MultiplyExpression.printType(pp, v)

	if f.Add != nil {
		pp.WriteString("\nAdd: ")
		f.Add.printType(pp, v)
	} else if v {
		pp.WriteString("\nAdd: nil")
	}

	if f.AddExpression != nil {
		pp.WriteString("\nAddExpression: ")
		f.AddExpression.printType(pp, v)
	} else if v {
		pp.WriteString("\nAddExpression: nil")
	}

	pp.WriteString("\nTokens: ")
	f.Tokens.printType(pp, v)

	w.WriteString("\n}")
}

func (f *AndExpression) printType(w writer, v bool) {
	pp := w.Indent()

	pp.WriteString("AndExpression {")

	pp.WriteString("\nShiftExpression: ")
	f.ShiftExpression.printType(pp, v)

	if f.AndExpression != nil {
		pp.WriteString("\nAndExpression: ")
		f.AndExpression.printType(pp, v)
	} else if v {
		pp.WriteString("\nAndExpression: nil")
	}

	pp.WriteString("\nTokens: ")
	f.Tokens.printType(pp, v)

	w.WriteString("\n}")
}

func (f *AndTest) printType(w writer, v bool) {
	pp := w.Indent()

	pp.WriteString("AndTest {")

	pp.WriteString("\nNotTest: ")
	f.NotTest.printType(pp, v)

	if f.AndTest != nil {
		pp.WriteString("\nAndTest: ")
		f.AndTest.printType(pp, v)
	} else if v {
		pp.WriteString("\nAndTest: nil")
	}

	pp.WriteString("\nTokens: ")
	f.Tokens.printType(pp, v)

	w.WriteString("\n}")
}

func (f *AnnotatedAssignmentStatement) printType(w writer, v bool) {
	pp := w.Indent()

	pp.WriteString("AnnotatedAssignmentStatement {")

	pp.WriteString("\nAugTarget: ")
	f.AugTarget.printType(pp, v)

	pp.WriteString("\nExpression: ")
	f.Expression.printType(pp, v)

	if f.StarredExpression != nil {
		pp.WriteString("\nStarredExpression: ")
		f.StarredExpression.printType(pp, v)
	} else if v {
		pp.WriteString("\nStarredExpression: nil")
	}

	if f.YieldExpression != nil {
		pp.WriteString("\nYieldExpression: ")
		f.YieldExpression.printType(pp, v)
	} else if v {
		pp.WriteString("\nYieldExpression: nil")
	}

	pp.WriteString("\nTokens: ")
	f.Tokens.printType(pp, v)

	w.WriteString("\n}")
}

func (f *ArgumentList) printType(w writer, v bool) {
	pp := w.Indent()

	pp.WriteString("ArgumentList {")

	if f.PositionalArguments == nil {
		pp.WriteString("\nPositionalArguments: nil")
	} else if len(f.PositionalArguments) > 0 {
		pp.WriteString("\nPositionalArguments: [")

		ipp := pp.Indent()

		for n, e := range f.PositionalArguments {
			ipp.Printf("\n%d: ", n)
			e.printType(ipp, v)
		}

		pp.WriteString("\n]")
	} else if v {
		pp.WriteString("\nPositionalArguments: []")
	}

	if f.StarredAndKeywordArguments == nil {
		pp.WriteString("\nStarredAndKeywordArguments: nil")
	} else if len(f.StarredAndKeywordArguments) > 0 {
		pp.WriteString("\nStarredAndKeywordArguments: [")

		ipp := pp.Indent()

		for n, e := range f.StarredAndKeywordArguments {
			ipp.Printf("\n%d: ", n)
			e.printType(ipp, v)
		}

		pp.WriteString("\n]")
	} else if v {
		pp.WriteString("\nStarredAndKeywordArguments: []")
	}

	if f.KeywordArguments == nil {
		pp.WriteString("\nKeywordArguments: nil")
	} else if len(f.KeywordArguments) > 0 {
		pp.WriteString("\nKeywordArguments: [")

		ipp := pp.Indent()

		for n, e := range f.KeywordArguments {
			ipp.Printf("\n%d: ", n)
			e.printType(ipp, v)
		}

		pp.WriteString("\n]")
	} else if v {
		pp.WriteString("\nKeywordArguments: []")
	}

	pp.WriteString("\nTokens: ")
	f.Tokens.printType(pp, v)

	w.WriteString("\n}")
}

func (f *ArgumentListOrComprehension) printType(w writer, v bool) {
	pp := w.Indent()

	pp.WriteString("ArgumentListOrComprehension {")

	if f.ArgumentList != nil {
		pp.WriteString("\nArgumentList: ")
		f.ArgumentList.printType(pp, v)
	} else if v {
		pp.WriteString("\nArgumentList: nil")
	}

	if f.Comprehension != nil {
		pp.WriteString("\nComprehension: ")
		f.Comprehension.printType(pp, v)
	} else if v {
		pp.WriteString("\nComprehension: nil")
	}

	pp.WriteString("\nTokens: ")
	f.Tokens.printType(pp, v)

	w.WriteString("\n}")
}

func (f *AssertStatement) printType(w writer, v bool) {
	pp := w.Indent()

	pp.WriteString("AssertStatement {")

	if f.Expressions == nil {
		pp.WriteString("\nExpressions: nil")
	} else if len(f.Expressions) > 0 {
		pp.WriteString("\nExpressions: [")

		ipp := pp.Indent()

		for n, e := range f.Expressions {
			ipp.Printf("\n%d: ", n)
			e.printType(ipp, v)
		}

		pp.WriteString("\n]")
	} else if v {
		pp.WriteString("\nExpressions: []")
	}

	pp.WriteString("\nTokens: ")
	f.Tokens.printType(pp, v)

	w.WriteString("\n}")
}

func (f *AssignmentExpressionAndSuite) printType(w writer, v bool) {
	pp := w.Indent()

	pp.WriteString("AssignmentExpressionAndSuite {")

	pp.WriteString("\nAssignmentExpression: ")
	f.AssignmentExpression.printType(pp, v)

	pp.WriteString("\nSuite: ")
	f.Suite.printType(pp, v)

	w.WriteString("\n}")
}

func (f *AssignmentExpression) printType(w writer, v bool) {
	pp := w.Indent()

	pp.WriteString("AssignmentExpression {")

	if f.Identifier != nil {
		pp.WriteString("\nIdentifier: ")
		f.Identifier.printType(pp, v)
	} else if v {
		pp.WriteString("\nIdentifier: nil")
	}

	pp.WriteString("\nExpression: ")
	f.Expression.printType(pp, v)

	pp.WriteString("\nTokens: ")
	f.Tokens.printType(pp, v)

	w.WriteString("\n}")
}

func (f *AssignmentStatement) printType(w writer, v bool) {
	pp := w.Indent()

	pp.WriteString("AssignmentStatement {")

	if f.TargetLists == nil {
		pp.WriteString("\nTargetLists: nil")
	} else if len(f.TargetLists) > 0 {
		pp.WriteString("\nTargetLists: [")

		ipp := pp.Indent()

		for n, e := range f.TargetLists {
			ipp.Printf("\n%d: ", n)
			e.printType(ipp, v)
		}

		pp.WriteString("\n]")
	} else if v {
		pp.WriteString("\nTargetLists: []")
	}

	if f.StarredExpression != nil {
		pp.WriteString("\nStarredExpression: ")
		f.StarredExpression.printType(pp, v)
	} else if v {
		pp.WriteString("\nStarredExpression: nil")
	}

	if f.YieldExpression != nil {
		pp.WriteString("\nYieldExpression: ")
		f.YieldExpression.printType(pp, v)
	} else if v {
		pp.WriteString("\nYieldExpression: nil")
	}

	pp.WriteString("\nTokens: ")
	f.Tokens.printType(pp, v)

	w.WriteString("\n}")
}

func (f *Atom) printType(w writer, v bool) {
	pp := w.Indent()

	pp.WriteString("Atom {")

	if f.Identifier != nil {
		pp.WriteString("\nIdentifier: ")
		f.Identifier.printType(pp, v)
	} else if v {
		pp.WriteString("\nIdentifier: nil")
	}

	if f.Literal != nil {
		pp.WriteString("\nLiteral: ")
		f.Literal.printType(pp, v)
	} else if v {
		pp.WriteString("\nLiteral: nil")
	}

	if f.Enclosure != nil {
		pp.WriteString("\nEnclosure: ")
		f.Enclosure.printType(pp, v)
	} else if v {
		pp.WriteString("\nEnclosure: nil")
	}

	pp.WriteString("\nTokens: ")
	f.Tokens.printType(pp, v)

	w.WriteString("\n}")
}

func (f *AugmentedAssignmentStatement) printType(w writer, v bool) {
	pp := w.Indent()

	pp.WriteString("AugmentedAssignmentStatement {")

	pp.WriteString("\nAugTarget: ")
	f.AugTarget.printType(pp, v)

	if f.AugOp != nil {
		pp.WriteString("\nAugOp: ")
		f.AugOp.printType(pp, v)
	} else if v {
		pp.WriteString("\nAugOp: nil")
	}

	if f.ExpressionList != nil {
		pp.WriteString("\nExpressionList: ")
		f.ExpressionList.printType(pp, v)
	} else if v {
		pp.WriteString("\nExpressionList: nil")
	}

	if f.YieldExpression != nil {
		pp.WriteString("\nYieldExpression: ")
		f.YieldExpression.printType(pp, v)
	} else if v {
		pp.WriteString("\nYieldExpression: nil")
	}

	pp.WriteString("\nTokens: ")
	f.Tokens.printType(pp, v)

	w.WriteString("\n}")
}

func (f *AugTarget) printType(w writer, v bool) {
	pp := w.Indent()

	pp.WriteString("AugTarget {")

	pp.WriteString("\nPrimaryExpression: ")
	f.PrimaryExpression.printType(pp, v)

	pp.WriteString("\nTokens: ")
	f.Tokens.printType(pp, v)

	w.WriteString("\n}")
}

func (f *ClassDefinition) printType(w writer, v bool) {
	pp := w.Indent()

	pp.WriteString("ClassDefinition {")

	if f.Decorators != nil {
		pp.WriteString("\nDecorators: ")
		f.Decorators.printType(pp, v)
	} else if v {
		pp.WriteString("\nDecorators: nil")
	}

	if f.ClassName != nil {
		pp.WriteString("\nClassName: ")
		f.ClassName.printType(pp, v)
	} else if v {
		pp.WriteString("\nClassName: nil")
	}

	if f.TypeParams != nil {
		pp.WriteString("\nTypeParams: ")
		f.TypeParams.printType(pp, v)
	} else if v {
		pp.WriteString("\nTypeParams: nil")
	}

	if f.Inheritance != nil {
		pp.WriteString("\nInheritance: ")
		f.Inheritance.printType(pp, v)
	} else if v {
		pp.WriteString("\nInheritance: nil")
	}

	pp.WriteString("\nSuite: ")
	f.Suite.printType(pp, v)

	pp.WriteString("\nTokens: ")
	f.Tokens.printType(pp, v)

	w.WriteString("\n}")
}

func (f *Comparison) printType(w writer, v bool) {
	pp := w.Indent()

	pp.WriteString("Comparison {")

	pp.WriteString("\nOrExpression: ")
	f.OrExpression.printType(pp, v)

	if f.Comparisons == nil {
		pp.WriteString("\nComparisons: nil")
	} else if len(f.Comparisons) > 0 {
		pp.WriteString("\nComparisons: [")

		ipp := pp.Indent()

		for n, e := range f.Comparisons {
			ipp.Printf("\n%d: ", n)
			e.printType(ipp, v)
		}

		pp.WriteString("\n]")
	} else if v {
		pp.WriteString("\nComparisons: []")
	}

	pp.WriteString("\nTokens: ")
	f.Tokens.printType(pp, v)

	w.WriteString("\n}")
}

func (f *ComparisonExpression) printType(w writer, v bool) {
	pp := w.Indent()

	pp.WriteString("ComparisonExpression {")

	if f.ComparisonOperator == nil {
		pp.WriteString("\nComparisonOperator: nil")
	} else if len(f.ComparisonOperator) > 0 {
		pp.WriteString("\nComparisonOperator: [")

		ipp := pp.Indent()

		for n, e := range f.ComparisonOperator {
			ipp.Printf("\n%d: ", n)
			e.printType(ipp, v)
		}

		pp.WriteString("\n]")
	} else if v {
		pp.WriteString("\nComparisonOperator: []")
	}

	pp.WriteString("\nOrExpression: ")
	f.OrExpression.printType(pp, v)

	w.WriteString("\n}")
}

func (f *CompoundStatement) printType(w writer, v bool) {
	pp := w.Indent()

	pp.WriteString("CompoundStatement {")

	if f.If != nil {
		pp.WriteString("\nIf: ")
		f.If.printType(pp, v)
	} else if v {
		pp.WriteString("\nIf: nil")
	}

	if f.While != nil {
		pp.WriteString("\nWhile: ")
		f.While.printType(pp, v)
	} else if v {
		pp.WriteString("\nWhile: nil")
	}

	if f.For != nil {
		pp.WriteString("\nFor: ")
		f.For.printType(pp, v)
	} else if v {
		pp.WriteString("\nFor: nil")
	}

	if f.Try != nil {
		pp.WriteString("\nTry: ")
		f.Try.printType(pp, v)
	} else if v {
		pp.WriteString("\nTry: nil")
	}

	if f.With != nil {
		pp.WriteString("\nWith: ")
		f.With.printType(pp, v)
	} else if v {
		pp.WriteString("\nWith: nil")
	}

	if f.Func != nil {
		pp.WriteString("\nFunc: ")
		f.Func.printType(pp, v)
	} else if v {
		pp.WriteString("\nFunc: nil")
	}

	if f.Class != nil {
		pp.WriteString("\nClass: ")
		f.Class.printType(pp, v)
	} else if v {
		pp.WriteString("\nClass: nil")
	}

	pp.WriteString("\nTokens: ")
	f.Tokens.printType(pp, v)

	w.WriteString("\n}")
}

func (f *Comprehension) printType(w writer, v bool) {
	pp := w.Indent()

	pp.WriteString("Comprehension {")

	pp.WriteString("\nAssignmentExpression: ")
	f.AssignmentExpression.printType(pp, v)

	pp.WriteString("\nComprehensionFor: ")
	f.ComprehensionFor.printType(pp, v)

	pp.WriteString("\nTokens: ")
	f.Tokens.printType(pp, v)

	w.WriteString("\n}")
}

func (f *ComprehensionFor) printType(w writer, v bool) {
	pp := w.Indent()

	pp.WriteString("ComprehensionFor {")

	if f.Async || v {
		pp.Printf("\nAsync: %v", f.Async)
	}

	pp.WriteString("\nTargetList: ")
	f.TargetList.printType(pp, v)

	pp.WriteString("\nOrTest: ")
	f.OrTest.printType(pp, v)

	if f.ComprehensionIterator != nil {
		pp.WriteString("\nComprehensionIterator: ")
		f.ComprehensionIterator.printType(pp, v)
	} else if v {
		pp.WriteString("\nComprehensionIterator: nil")
	}

	pp.WriteString("\nTokens: ")
	f.Tokens.printType(pp, v)

	w.WriteString("\n}")
}

func (f *ComprehensionIf) printType(w writer, v bool) {
	pp := w.Indent()

	pp.WriteString("ComprehensionIf {")

	pp.WriteString("\nOrTest: ")
	f.OrTest.printType(pp, v)

	if f.ComprehensionIterator != nil {
		pp.WriteString("\nComprehensionIterator: ")
		f.ComprehensionIterator.printType(pp, v)
	} else if v {
		pp.WriteString("\nComprehensionIterator: nil")
	}

	pp.WriteString("\nTokens: ")
	f.Tokens.printType(pp, v)

	w.WriteString("\n}")
}

func (f *ComprehensionIterator) printType(w writer, v bool) {
	pp := w.Indent()

	pp.WriteString("ComprehensionIterator {")

	if f.ComprehensionFor != nil {
		pp.WriteString("\nComprehensionFor: ")
		f.ComprehensionFor.printType(pp, v)
	} else if v {
		pp.WriteString("\nComprehensionFor: nil")
	}

	if f.ComprehensionIf != nil {
		pp.WriteString("\nComprehensionIf: ")
		f.ComprehensionIf.printType(pp, v)
	} else if v {
		pp.WriteString("\nComprehensionIf: nil")
	}

	pp.WriteString("\nTokens: ")
	f.Tokens.printType(pp, v)

	w.WriteString("\n}")
}

func (f *ConditionalExpression) printType(w writer, v bool) {
	pp := w.Indent()

	pp.WriteString("ConditionalExpression {")

	pp.WriteString("\nOrTest: ")
	f.OrTest.printType(pp, v)

	if f.If != nil {
		pp.WriteString("\nIf: ")
		f.If.printType(pp, v)
	} else if v {
		pp.WriteString("\nIf: nil")
	}

	if f.Else != nil {
		pp.WriteString("\nElse: ")
		f.Else.printType(pp, v)
	} else if v {
		pp.WriteString("\nElse: nil")
	}

	pp.WriteString("\nTokens: ")
	f.Tokens.printType(pp, v)

	w.WriteString("\n}")
}

func (f *Decorators) printType(w writer, v bool) {
	pp := w.Indent()

	pp.WriteString("Decorators {")

	if f.Decorators == nil {
		pp.WriteString("\nDecorators: nil")
	} else if len(f.Decorators) > 0 {
		pp.WriteString("\nDecorators: [")

		ipp := pp.Indent()

		for n, e := range f.Decorators {
			ipp.Printf("\n%d: ", n)
			e.printType(ipp, v)
		}

		pp.WriteString("\n]")
	} else if v {
		pp.WriteString("\nDecorators: []")
	}

	pp.WriteString("\nTokens: ")
	f.Tokens.printType(pp, v)

	w.WriteString("\n}")
}

func (f *DefParameter) printType(w writer, v bool) {
	pp := w.Indent()

	pp.WriteString("DefParameter {")

	pp.WriteString("\nParameter: ")
	f.Parameter.printType(pp, v)

	if f.Value != nil {
		pp.WriteString("\nValue: ")
		f.Value.printType(pp, v)
	} else if v {
		pp.WriteString("\nValue: nil")
	}

	pp.WriteString("\nComments: [")

	ipp := pp.Indent()

	for n, e := range f.Comments {
		ipp.Printf("\n%d: ", n)
		e.printType(ipp, v)
	}

	pp.WriteString("\n]")

	pp.WriteString("\nTokens: ")
	f.Tokens.printType(pp, v)

	w.WriteString("\n}")
}

func (f *DelStatement) printType(w writer, v bool) {
	pp := w.Indent()

	pp.WriteString("DelStatement {")

	pp.WriteString("\nTargetList: ")
	f.TargetList.printType(pp, v)

	pp.WriteString("\nTokens: ")
	f.Tokens.printType(pp, v)

	w.WriteString("\n}")
}

func (f *DictDisplay) printType(w writer, v bool) {
	pp := w.Indent()

	pp.WriteString("DictDisplay {")

	if f.DictItems == nil {
		pp.WriteString("\nDictItems: nil")
	} else if len(f.DictItems) > 0 {
		pp.WriteString("\nDictItems: [")

		ipp := pp.Indent()

		for n, e := range f.DictItems {
			ipp.Printf("\n%d: ", n)
			e.printType(ipp, v)
		}

		pp.WriteString("\n]")
	} else if v {
		pp.WriteString("\nDictItems: []")
	}

	if f.DictComprehension != nil {
		pp.WriteString("\nDictComprehension: ")
		f.DictComprehension.printType(pp, v)
	} else if v {
		pp.WriteString("\nDictComprehension: nil")
	}

	pp.WriteString("\nTokens: ")
	f.Tokens.printType(pp, v)

	w.WriteString("\n}")
}

func (f *DictItem) printType(w writer, v bool) {
	pp := w.Indent()

	pp.WriteString("DictItem {")

	if f.Key != nil {
		pp.WriteString("\nKey: ")
		f.Key.printType(pp, v)
	} else if v {
		pp.WriteString("\nKey: nil")
	}

	if f.Value != nil {
		pp.WriteString("\nValue: ")
		f.Value.printType(pp, v)
	} else if v {
		pp.WriteString("\nValue: nil")
	}

	if f.OrExpression != nil {
		pp.WriteString("\nOrExpression: ")
		f.OrExpression.printType(pp, v)
	} else if v {
		pp.WriteString("\nOrExpression: nil")
	}

	pp.WriteString("\nTokens: ")
	f.Tokens.printType(pp, v)

	w.WriteString("\n}")
}

func (f *Enclosure) printType(w writer, v bool) {
	pp := w.Indent()

	pp.WriteString("Enclosure {")

	if f.ParenthForm != nil {
		pp.WriteString("\nParenthForm: ")
		f.ParenthForm.printType(pp, v)
	} else if v {
		pp.WriteString("\nParenthForm: nil")
	}

	if f.ListDisplay != nil {
		pp.WriteString("\nListDisplay: ")
		f.ListDisplay.printType(pp, v)
	} else if v {
		pp.WriteString("\nListDisplay: nil")
	}

	if f.DictDisplay != nil {
		pp.WriteString("\nDictDisplay: ")
		f.DictDisplay.printType(pp, v)
	} else if v {
		pp.WriteString("\nDictDisplay: nil")
	}

	if f.SetDisplay != nil {
		pp.WriteString("\nSetDisplay: ")
		f.SetDisplay.printType(pp, v)
	} else if v {
		pp.WriteString("\nSetDisplay: nil")
	}

	if f.GeneratorExpression != nil {
		pp.WriteString("\nGeneratorExpression: ")
		f.GeneratorExpression.printType(pp, v)
	} else if v {
		pp.WriteString("\nGeneratorExpression: nil")
	}

	if f.YieldAtom != nil {
		pp.WriteString("\nYieldAtom: ")
		f.YieldAtom.printType(pp, v)
	} else if v {
		pp.WriteString("\nYieldAtom: nil")
	}

	pp.WriteString("\nComments: [")

	ipp := pp.Indent()

	for n, e := range f.Comments {
		ipp.Printf("\n%d: ", n)
		e.printType(ipp, v)
	}

	pp.WriteString("\n]")

	pp.WriteString("\nTokens: ")
	f.Tokens.printType(pp, v)

	w.WriteString("\n}")
}

func (f *Except) printType(w writer, v bool) {
	pp := w.Indent()

	pp.WriteString("Except {")

	pp.WriteString("\nExpression: ")
	f.Expression.printType(pp, v)

	if f.Identifier != nil {
		pp.WriteString("\nIdentifier: ")
		f.Identifier.printType(pp, v)
	} else if v {
		pp.WriteString("\nIdentifier: nil")
	}

	pp.WriteString("\nSuite: ")
	f.Suite.printType(pp, v)

	pp.WriteString("\nTokens: ")
	f.Tokens.printType(pp, v)

	w.WriteString("\n}")
}

func (f *Expression) printType(w writer, v bool) {
	pp := w.Indent()

	pp.WriteString("Expression {")

	if f.ConditionalExpression != nil {
		pp.WriteString("\nConditionalExpression: ")
		f.ConditionalExpression.printType(pp, v)
	} else if v {
		pp.WriteString("\nConditionalExpression: nil")
	}

	if f.LambdaExpression != nil {
		pp.WriteString("\nLambdaExpression: ")
		f.LambdaExpression.printType(pp, v)
	} else if v {
		pp.WriteString("\nLambdaExpression: nil")
	}

	pp.WriteString("\nTokens: ")
	f.Tokens.printType(pp, v)

	w.WriteString("\n}")
}

func (f *ExpressionList) printType(w writer, v bool) {
	pp := w.Indent()

	pp.WriteString("ExpressionList {")

	if f.Expressions == nil {
		pp.WriteString("\nExpressions: nil")
	} else if len(f.Expressions) > 0 {
		pp.WriteString("\nExpressions: [")

		ipp := pp.Indent()

		for n, e := range f.Expressions {
			ipp.Printf("\n%d: ", n)
			e.printType(ipp, v)
		}

		pp.WriteString("\n]")
	} else if v {
		pp.WriteString("\nExpressions: []")
	}

	pp.WriteString("\nTokens: ")
	f.Tokens.printType(pp, v)

	w.WriteString("\n}")
}

func (f *File) printType(w writer, v bool) {
	pp := w.Indent()

	pp.WriteString("File {")

	if f.Statements == nil {
		pp.WriteString("\nStatements: nil")
	} else if len(f.Statements) > 0 {
		pp.WriteString("\nStatements: [")

		ipp := pp.Indent()

		for n, e := range f.Statements {
			ipp.Printf("\n%d: ", n)
			e.printType(ipp, v)
		}

		pp.WriteString("\n]")
	} else if v {
		pp.WriteString("\nStatements: []")
	}

	pp.WriteString("\nComments: [")

	ipp := pp.Indent()

	for n, e := range f.Comments {
		ipp.Printf("\n%d: ", n)
		e.printType(ipp, v)
	}

	pp.WriteString("\n]")

	pp.WriteString("\nTokens: ")
	f.Tokens.printType(pp, v)

	w.WriteString("\n}")
}

func (f *FlexibleExpression) printType(w writer, v bool) {
	pp := w.Indent()

	pp.WriteString("FlexibleExpression {")

	if f.AssignmentExpression != nil {
		pp.WriteString("\nAssignmentExpression: ")
		f.AssignmentExpression.printType(pp, v)
	} else if v {
		pp.WriteString("\nAssignmentExpression: nil")
	}

	if f.StarredExpression != nil {
		pp.WriteString("\nStarredExpression: ")
		f.StarredExpression.printType(pp, v)
	} else if v {
		pp.WriteString("\nStarredExpression: nil")
	}

	pp.WriteString("\nTokens: ")
	f.Tokens.printType(pp, v)

	w.WriteString("\n}")
}

func (f *FlexibleExpressionList) printType(w writer, v bool) {
	pp := w.Indent()

	pp.WriteString("FlexibleExpressionList {")

	if f.FlexibleExpressions == nil {
		pp.WriteString("\nFlexibleExpressions: nil")
	} else if len(f.FlexibleExpressions) > 0 {
		pp.WriteString("\nFlexibleExpressions: [")

		ipp := pp.Indent()

		for n, e := range f.FlexibleExpressions {
			ipp.Printf("\n%d: ", n)
			e.printType(ipp, v)
		}

		pp.WriteString("\n]")
	} else if v {
		pp.WriteString("\nFlexibleExpressions: []")
	}

	pp.WriteString("\nTokens: ")
	f.Tokens.printType(pp, v)

	w.WriteString("\n}")
}

func (f *FlexibleExpressionListOrComprehension) printType(w writer, v bool) {
	pp := w.Indent()

	pp.WriteString("FlexibleExpressionListOrComprehension {")

	if f.FlexibleExpressionList != nil {
		pp.WriteString("\nFlexibleExpressionList: ")
		f.FlexibleExpressionList.printType(pp, v)
	} else if v {
		pp.WriteString("\nFlexibleExpressionList: nil")
	}

	if f.Comprehension != nil {
		pp.WriteString("\nComprehension: ")
		f.Comprehension.printType(pp, v)
	} else if v {
		pp.WriteString("\nComprehension: nil")
	}

	pp.WriteString("\nTokens: ")
	f.Tokens.printType(pp, v)

	w.WriteString("\n}")
}

func (f *ForStatement) printType(w writer, v bool) {
	pp := w.Indent()

	pp.WriteString("ForStatement {")

	if f.Async || v {
		pp.Printf("\nAsync: %v", f.Async)
	}

	pp.WriteString("\nTargetList: ")
	f.TargetList.printType(pp, v)

	pp.WriteString("\nStarredList: ")
	f.StarredList.printType(pp, v)

	pp.WriteString("\nSuite: ")
	f.Suite.printType(pp, v)

	if f.Else != nil {
		pp.WriteString("\nElse: ")
		f.Else.printType(pp, v)
	} else if v {
		pp.WriteString("\nElse: nil")
	}

	pp.WriteString("\nTokens: ")
	f.Tokens.printType(pp, v)

	w.WriteString("\n}")
}

func (f *FuncDefinition) printType(w writer, v bool) {
	pp := w.Indent()

	pp.WriteString("FuncDefinition {")

	if f.Decorators != nil {
		pp.WriteString("\nDecorators: ")
		f.Decorators.printType(pp, v)
	} else if v {
		pp.WriteString("\nDecorators: nil")
	}

	if f.Async || v {
		pp.Printf("\nAsync: %v", f.Async)
	}

	if f.FuncName != nil {
		pp.WriteString("\nFuncName: ")
		f.FuncName.printType(pp, v)
	} else if v {
		pp.WriteString("\nFuncName: nil")
	}

	if f.TypeParams != nil {
		pp.WriteString("\nTypeParams: ")
		f.TypeParams.printType(pp, v)
	} else if v {
		pp.WriteString("\nTypeParams: nil")
	}

	pp.WriteString("\nParameterList: ")
	f.ParameterList.printType(pp, v)

	if f.Expression != nil {
		pp.WriteString("\nExpression: ")
		f.Expression.printType(pp, v)
	} else if v {
		pp.WriteString("\nExpression: nil")
	}

	pp.WriteString("\nSuite: ")
	f.Suite.printType(pp, v)

	pp.WriteString("\nComments: ")
	f.Comments.printType(pp, v)

	pp.WriteString("\nTokens: ")
	f.Tokens.printType(pp, v)

	w.WriteString("\n}")
}

func (f *GeneratorExpression) printType(w writer, v bool) {
	pp := w.Indent()

	pp.WriteString("GeneratorExpression {")

	pp.WriteString("\nExpression: ")
	f.Expression.printType(pp, v)

	pp.WriteString("\nComprehensionFor: ")
	f.ComprehensionFor.printType(pp, v)

	pp.WriteString("\nComments: [")

	ipp := pp.Indent()

	for n, e := range f.Comments {
		ipp.Printf("\n%d: ", n)
		e.printType(ipp, v)
	}

	pp.WriteString("\n]")

	pp.WriteString("\nTokens: ")
	f.Tokens.printType(pp, v)

	w.WriteString("\n}")
}

func (f *GlobalStatement) printType(w writer, v bool) {
	pp := w.Indent()

	pp.WriteString("GlobalStatement {")

	if f.Identifiers == nil {
		pp.WriteString("\nIdentifiers: nil")
	} else if len(f.Identifiers) > 0 {
		pp.WriteString("\nIdentifiers: [")

		ipp := pp.Indent()

		for n, e := range f.Identifiers {
			ipp.Printf("\n%d: ", n)
			e.printType(ipp, v)
		}

		pp.WriteString("\n]")
	} else if v {
		pp.WriteString("\nIdentifiers: []")
	}

	pp.WriteString("\nTokens: ")
	f.Tokens.printType(pp, v)

	w.WriteString("\n}")
}

func (f *IfStatement) printType(w writer, v bool) {
	pp := w.Indent()

	pp.WriteString("IfStatement {")

	pp.WriteString("\nAssignmentExpression: ")
	f.AssignmentExpression.printType(pp, v)

	pp.WriteString("\nSuite: ")
	f.Suite.printType(pp, v)

	if f.Elif == nil {
		pp.WriteString("\nElif: nil")
	} else if len(f.Elif) > 0 {
		pp.WriteString("\nElif: [")

		ipp := pp.Indent()

		for n, e := range f.Elif {
			ipp.Printf("\n%d: ", n)
			e.printType(ipp, v)
		}

		pp.WriteString("\n]")
	} else if v {
		pp.WriteString("\nElif: []")
	}

	if f.Else != nil {
		pp.WriteString("\nElse: ")
		f.Else.printType(pp, v)
	} else if v {
		pp.WriteString("\nElse: nil")
	}

	pp.WriteString("\nTokens: ")
	f.Tokens.printType(pp, v)

	w.WriteString("\n}")
}

func (f *ImportStatement) printType(w writer, v bool) {
	pp := w.Indent()

	pp.WriteString("ImportStatement {")

	if f.RelativeModule != nil {
		pp.WriteString("\nRelativeModule: ")
		f.RelativeModule.printType(pp, v)
	} else if v {
		pp.WriteString("\nRelativeModule: nil")
	}

	if f.Modules == nil {
		pp.WriteString("\nModules: nil")
	} else if len(f.Modules) > 0 {
		pp.WriteString("\nModules: [")

		ipp := pp.Indent()

		for n, e := range f.Modules {
			ipp.Printf("\n%d: ", n)
			e.printType(ipp, v)
		}

		pp.WriteString("\n]")
	} else if v {
		pp.WriteString("\nModules: []")
	}

	pp.WriteString("\nTokens: ")
	f.Tokens.printType(pp, v)

	w.WriteString("\n}")
}

func (f *KeywordArgument) printType(w writer, v bool) {
	pp := w.Indent()

	pp.WriteString("KeywordArgument {")

	if f.KeywordItem != nil {
		pp.WriteString("\nKeywordItem: ")
		f.KeywordItem.printType(pp, v)
	} else if v {
		pp.WriteString("\nKeywordItem: nil")
	}

	if f.Expression != nil {
		pp.WriteString("\nExpression: ")
		f.Expression.printType(pp, v)
	} else if v {
		pp.WriteString("\nExpression: nil")
	}

	pp.WriteString("\nTokens: ")
	f.Tokens.printType(pp, v)

	w.WriteString("\n}")
}

func (f *KeywordItem) printType(w writer, v bool) {
	pp := w.Indent()

	pp.WriteString("KeywordItem {")

	if f.Identifier != nil {
		pp.WriteString("\nIdentifier: ")
		f.Identifier.printType(pp, v)
	} else if v {
		pp.WriteString("\nIdentifier: nil")
	}

	pp.WriteString("\nExpression: ")
	f.Expression.printType(pp, v)

	pp.WriteString("\nTokens: ")
	f.Tokens.printType(pp, v)

	w.WriteString("\n}")
}

func (f *LambdaExpression) printType(w writer, v bool) {
	pp := w.Indent()

	pp.WriteString("LambdaExpression {")

	if f.ParameterList != nil {
		pp.WriteString("\nParameterList: ")
		f.ParameterList.printType(pp, v)
	} else if v {
		pp.WriteString("\nParameterList: nil")
	}

	pp.WriteString("\nExpression: ")
	f.Expression.printType(pp, v)

	pp.WriteString("\nTokens: ")
	f.Tokens.printType(pp, v)

	w.WriteString("\n}")
}

func (f *ModuleAs) printType(w writer, v bool) {
	pp := w.Indent()

	pp.WriteString("ModuleAs {")

	pp.WriteString("\nModule: ")
	f.Module.printType(pp, v)

	if f.As != nil {
		pp.WriteString("\nAs: ")
		f.As.printType(pp, v)
	} else if v {
		pp.WriteString("\nAs: nil")
	}

	pp.WriteString("\nTokens: ")
	f.Tokens.printType(pp, v)

	w.WriteString("\n}")
}

func (f *Module) printType(w writer, v bool) {
	pp := w.Indent()

	pp.WriteString("Module {")

	if f.Identifiers == nil {
		pp.WriteString("\nIdentifiers: nil")
	} else if len(f.Identifiers) > 0 {
		pp.WriteString("\nIdentifiers: [")

		ipp := pp.Indent()

		for n, e := range f.Identifiers {
			ipp.Printf("\n%d: ", n)
			e.printType(ipp, v)
		}

		pp.WriteString("\n]")
	} else if v {
		pp.WriteString("\nIdentifiers: []")
	}

	pp.WriteString("\nTokens: ")
	f.Tokens.printType(pp, v)

	w.WriteString("\n}")
}

func (f *MultiplyExpression) printType(w writer, v bool) {
	pp := w.Indent()

	pp.WriteString("MultiplyExpression {")

	pp.WriteString("\nUnaryExpression: ")
	f.UnaryExpression.printType(pp, v)

	if f.Multiply != nil {
		pp.WriteString("\nMultiply: ")
		f.Multiply.printType(pp, v)
	} else if v {
		pp.WriteString("\nMultiply: nil")
	}

	if f.MultiplyExpression != nil {
		pp.WriteString("\nMultiplyExpression: ")
		f.MultiplyExpression.printType(pp, v)
	} else if v {
		pp.WriteString("\nMultiplyExpression: nil")
	}

	pp.WriteString("\nTokens: ")
	f.Tokens.printType(pp, v)

	w.WriteString("\n}")
}

func (f *NonLocalStatement) printType(w writer, v bool) {
	pp := w.Indent()

	pp.WriteString("NonLocalStatement {")

	if f.Identifiers == nil {
		pp.WriteString("\nIdentifiers: nil")
	} else if len(f.Identifiers) > 0 {
		pp.WriteString("\nIdentifiers: [")

		ipp := pp.Indent()

		for n, e := range f.Identifiers {
			ipp.Printf("\n%d: ", n)
			e.printType(ipp, v)
		}

		pp.WriteString("\n]")
	} else if v {
		pp.WriteString("\nIdentifiers: []")
	}

	pp.WriteString("\nTokens: ")
	f.Tokens.printType(pp, v)

	w.WriteString("\n}")
}

func (f *NotTest) printType(w writer, v bool) {
	pp := w.Indent()

	pp.WriteString("NotTest {")

	if f.Nots != 0 || v {
		pp.Printf("\nNots: %v", f.Nots)
	}

	pp.WriteString("\nComparison: ")
	f.Comparison.printType(pp, v)

	pp.WriteString("\nTokens: ")
	f.Tokens.printType(pp, v)

	w.WriteString("\n}")
}

func (f *OrExpression) printType(w writer, v bool) {
	pp := w.Indent()

	pp.WriteString("OrExpression {")

	pp.WriteString("\nXorExpression: ")
	f.XorExpression.printType(pp, v)

	if f.OrExpression != nil {
		pp.WriteString("\nOrExpression: ")
		f.OrExpression.printType(pp, v)
	} else if v {
		pp.WriteString("\nOrExpression: nil")
	}

	pp.WriteString("\nTokens: ")
	f.Tokens.printType(pp, v)

	w.WriteString("\n}")
}

func (f *OrTest) printType(w writer, v bool) {
	pp := w.Indent()

	pp.WriteString("OrTest {")

	pp.WriteString("\nAndTest: ")
	f.AndTest.printType(pp, v)

	if f.OrTest != nil {
		pp.WriteString("\nOrTest: ")
		f.OrTest.printType(pp, v)
	} else if v {
		pp.WriteString("\nOrTest: nil")
	}

	pp.WriteString("\nTokens: ")
	f.Tokens.printType(pp, v)

	w.WriteString("\n}")
}

func (f *Parameter) printType(w writer, v bool) {
	pp := w.Indent()

	pp.WriteString("Parameter {")

	if f.Identifier != nil {
		pp.WriteString("\nIdentifier: ")
		f.Identifier.printType(pp, v)
	} else if v {
		pp.WriteString("\nIdentifier: nil")
	}

	if f.Type != nil {
		pp.WriteString("\nType: ")
		f.Type.printType(pp, v)
	} else if v {
		pp.WriteString("\nType: nil")
	}

	pp.WriteString("\nTokens: ")
	f.Tokens.printType(pp, v)

	w.WriteString("\n}")
}

func (f *ParameterList) printType(w writer, v bool) {
	pp := w.Indent()

	pp.WriteString("ParameterList {")

	if f.DefParameters == nil {
		pp.WriteString("\nDefParameters: nil")
	} else if len(f.DefParameters) > 0 {
		pp.WriteString("\nDefParameters: [")

		ipp := pp.Indent()

		for n, e := range f.DefParameters {
			ipp.Printf("\n%d: ", n)
			e.printType(ipp, v)
		}

		pp.WriteString("\n]")
	} else if v {
		pp.WriteString("\nDefParameters: []")
	}

	if f.NoPosOnly == nil {
		pp.WriteString("\nNoPosOnly: nil")
	} else if len(f.NoPosOnly) > 0 {
		pp.WriteString("\nNoPosOnly: [")

		ipp := pp.Indent()

		for n, e := range f.NoPosOnly {
			ipp.Printf("\n%d: ", n)
			e.printType(ipp, v)
		}

		pp.WriteString("\n]")
	} else if v {
		pp.WriteString("\nNoPosOnly: []")
	}

	if f.StarArg != nil {
		pp.WriteString("\nStarArg: ")
		f.StarArg.printType(pp, v)
	} else if v {
		pp.WriteString("\nStarArg: nil")
	}

	if f.StarArgs == nil {
		pp.WriteString("\nStarArgs: nil")
	} else if len(f.StarArgs) > 0 {
		pp.WriteString("\nStarArgs: [")

		ipp := pp.Indent()

		for n, e := range f.StarArgs {
			ipp.Printf("\n%d: ", n)
			e.printType(ipp, v)
		}

		pp.WriteString("\n]")
	} else if v {
		pp.WriteString("\nStarArgs: []")
	}

	if f.StarStarArg != nil {
		pp.WriteString("\nStarStarArg: ")
		f.StarStarArg.printType(pp, v)
	} else if v {
		pp.WriteString("\nStarStarArg: nil")
	}

	pp.WriteString("\nComments: [")

	ipp := pp.Indent()

	for n, e := range f.Comments {
		ipp.Printf("\n%d: ", n)
		e.printType(ipp, v)
	}

	pp.WriteString("\n]")

	pp.WriteString("\nTokens: ")
	f.Tokens.printType(pp, v)

	w.WriteString("\n}")
}

func (f *PositionalArgument) printType(w writer, v bool) {
	pp := w.Indent()

	pp.WriteString("PositionalArgument {")

	if f.AssignmentExpression != nil {
		pp.WriteString("\nAssignmentExpression: ")
		f.AssignmentExpression.printType(pp, v)
	} else if v {
		pp.WriteString("\nAssignmentExpression: nil")
	}

	if f.Expression != nil {
		pp.WriteString("\nExpression: ")
		f.Expression.printType(pp, v)
	} else if v {
		pp.WriteString("\nExpression: nil")
	}

	pp.WriteString("\nTokens: ")
	f.Tokens.printType(pp, v)

	w.WriteString("\n}")
}

func (f *PowerExpression) printType(w writer, v bool) {
	pp := w.Indent()

	pp.WriteString("PowerExpression {")

	if f.AwaitExpression || v {
		pp.Printf("\nAwaitExpression: %v", f.AwaitExpression)
	}

	pp.WriteString("\nPrimaryExpression: ")
	f.PrimaryExpression.printType(pp, v)

	if f.UnaryExpression != nil {
		pp.WriteString("\nUnaryExpression: ")
		f.UnaryExpression.printType(pp, v)
	} else if v {
		pp.WriteString("\nUnaryExpression: nil")
	}

	pp.WriteString("\nTokens: ")
	f.Tokens.printType(pp, v)

	w.WriteString("\n}")
}

func (f *PrimaryExpression) printType(w writer, v bool) {
	pp := w.Indent()

	pp.WriteString("PrimaryExpression {")

	if f.PrimaryExpression != nil {
		pp.WriteString("\nPrimaryExpression: ")
		f.PrimaryExpression.printType(pp, v)
	} else if v {
		pp.WriteString("\nPrimaryExpression: nil")
	}

	if f.Atom != nil {
		pp.WriteString("\nAtom: ")
		f.Atom.printType(pp, v)
	} else if v {
		pp.WriteString("\nAtom: nil")
	}

	if f.AttributeRef != nil {
		pp.WriteString("\nAttributeRef: ")
		f.AttributeRef.printType(pp, v)
	} else if v {
		pp.WriteString("\nAttributeRef: nil")
	}

	if f.Slicing != nil {
		pp.WriteString("\nSlicing: ")
		f.Slicing.printType(pp, v)
	} else if v {
		pp.WriteString("\nSlicing: nil")
	}

	if f.Call != nil {
		pp.WriteString("\nCall: ")
		f.Call.printType(pp, v)
	} else if v {
		pp.WriteString("\nCall: nil")
	}

	pp.WriteString("\nTokens: ")
	f.Tokens.printType(pp, v)

	w.WriteString("\n}")
}

func (f *RaiseStatement) printType(w writer, v bool) {
	pp := w.Indent()

	pp.WriteString("RaiseStatement {")

	if f.Expression != nil {
		pp.WriteString("\nExpression: ")
		f.Expression.printType(pp, v)
	} else if v {
		pp.WriteString("\nExpression: nil")
	}

	if f.From != nil {
		pp.WriteString("\nFrom: ")
		f.From.printType(pp, v)
	} else if v {
		pp.WriteString("\nFrom: nil")
	}

	pp.WriteString("\nTokens: ")
	f.Tokens.printType(pp, v)

	w.WriteString("\n}")
}

func (f *RelativeModule) printType(w writer, v bool) {
	pp := w.Indent()

	pp.WriteString("RelativeModule {")

	if f.Dots != 0 || v {
		pp.Printf("\nDots: %v", f.Dots)
	}

	if f.Module != nil {
		pp.WriteString("\nModule: ")
		f.Module.printType(pp, v)
	} else if v {
		pp.WriteString("\nModule: nil")
	}

	pp.WriteString("\nTokens: ")
	f.Tokens.printType(pp, v)

	w.WriteString("\n}")
}

func (f *ReturnStatement) printType(w writer, v bool) {
	pp := w.Indent()

	pp.WriteString("ReturnStatement {")

	if f.Expression != nil {
		pp.WriteString("\nExpression: ")
		f.Expression.printType(pp, v)
	} else if v {
		pp.WriteString("\nExpression: nil")
	}

	pp.WriteString("\nTokens: ")
	f.Tokens.printType(pp, v)

	w.WriteString("\n}")
}

func (f *ShiftExpression) printType(w writer, v bool) {
	pp := w.Indent()

	pp.WriteString("ShiftExpression {")

	pp.WriteString("\nAddExpression: ")
	f.AddExpression.printType(pp, v)

	if f.Shift != nil {
		pp.WriteString("\nShift: ")
		f.Shift.printType(pp, v)
	} else if v {
		pp.WriteString("\nShift: nil")
	}

	if f.ShiftExpression != nil {
		pp.WriteString("\nShiftExpression: ")
		f.ShiftExpression.printType(pp, v)
	} else if v {
		pp.WriteString("\nShiftExpression: nil")
	}

	pp.WriteString("\nTokens: ")
	f.Tokens.printType(pp, v)

	w.WriteString("\n}")
}

func (f *SimpleStatement) printType(w writer, v bool) {
	pp := w.Indent()

	pp.WriteString("SimpleStatement {")

	pp.WriteString("\nType: ")
	f.Type.printType(pp, v)

	if f.AssertStatement != nil {
		pp.WriteString("\nAssertStatement: ")
		f.AssertStatement.printType(pp, v)
	} else if v {
		pp.WriteString("\nAssertStatement: nil")
	}

	if f.AssignmentStatement != nil {
		pp.WriteString("\nAssignmentStatement: ")
		f.AssignmentStatement.printType(pp, v)
	} else if v {
		pp.WriteString("\nAssignmentStatement: nil")
	}

	if f.AugmentedAssignmentStatement != nil {
		pp.WriteString("\nAugmentedAssignmentStatement: ")
		f.AugmentedAssignmentStatement.printType(pp, v)
	} else if v {
		pp.WriteString("\nAugmentedAssignmentStatement: nil")
	}

	if f.AnnotatedAssignmentStatement != nil {
		pp.WriteString("\nAnnotatedAssignmentStatement: ")
		f.AnnotatedAssignmentStatement.printType(pp, v)
	} else if v {
		pp.WriteString("\nAnnotatedAssignmentStatement: nil")
	}

	if f.DelStatement != nil {
		pp.WriteString("\nDelStatement: ")
		f.DelStatement.printType(pp, v)
	} else if v {
		pp.WriteString("\nDelStatement: nil")
	}

	if f.ReturnStatement != nil {
		pp.WriteString("\nReturnStatement: ")
		f.ReturnStatement.printType(pp, v)
	} else if v {
		pp.WriteString("\nReturnStatement: nil")
	}

	if f.YieldStatement != nil {
		pp.WriteString("\nYieldStatement: ")
		f.YieldStatement.printType(pp, v)
	} else if v {
		pp.WriteString("\nYieldStatement: nil")
	}

	if f.RaiseStatement != nil {
		pp.WriteString("\nRaiseStatement: ")
		f.RaiseStatement.printType(pp, v)
	} else if v {
		pp.WriteString("\nRaiseStatement: nil")
	}

	if f.ImportStatement != nil {
		pp.WriteString("\nImportStatement: ")
		f.ImportStatement.printType(pp, v)
	} else if v {
		pp.WriteString("\nImportStatement: nil")
	}

	if f.GlobalStatement != nil {
		pp.WriteString("\nGlobalStatement: ")
		f.GlobalStatement.printType(pp, v)
	} else if v {
		pp.WriteString("\nGlobalStatement: nil")
	}

	if f.NonLocalStatement != nil {
		pp.WriteString("\nNonLocalStatement: ")
		f.NonLocalStatement.printType(pp, v)
	} else if v {
		pp.WriteString("\nNonLocalStatement: nil")
	}

	if f.TypeStatement != nil {
		pp.WriteString("\nTypeStatement: ")
		f.TypeStatement.printType(pp, v)
	} else if v {
		pp.WriteString("\nTypeStatement: nil")
	}

	pp.WriteString("\nTokens: ")
	f.Tokens.printType(pp, v)

	w.WriteString("\n}")
}

func (f *SliceItem) printType(w writer, v bool) {
	pp := w.Indent()

	pp.WriteString("SliceItem {")

	if f.Expression != nil {
		pp.WriteString("\nExpression: ")
		f.Expression.printType(pp, v)
	} else if v {
		pp.WriteString("\nExpression: nil")
	}

	if f.LowerBound != nil {
		pp.WriteString("\nLowerBound: ")
		f.LowerBound.printType(pp, v)
	} else if v {
		pp.WriteString("\nLowerBound: nil")
	}

	if f.UpperBound != nil {
		pp.WriteString("\nUpperBound: ")
		f.UpperBound.printType(pp, v)
	} else if v {
		pp.WriteString("\nUpperBound: nil")
	}

	if f.Stride != nil {
		pp.WriteString("\nStride: ")
		f.Stride.printType(pp, v)
	} else if v {
		pp.WriteString("\nStride: nil")
	}

	pp.WriteString("\nTokens: ")
	f.Tokens.printType(pp, v)

	w.WriteString("\n}")
}

func (f *SliceList) printType(w writer, v bool) {
	pp := w.Indent()

	pp.WriteString("SliceList {")

	if f.SliceItems == nil {
		pp.WriteString("\nSliceItems: nil")
	} else if len(f.SliceItems) > 0 {
		pp.WriteString("\nSliceItems: [")

		ipp := pp.Indent()

		for n, e := range f.SliceItems {
			ipp.Printf("\n%d: ", n)
			e.printType(ipp, v)
		}

		pp.WriteString("\n]")
	} else if v {
		pp.WriteString("\nSliceItems: []")
	}

	pp.WriteString("\nTokens: ")
	f.Tokens.printType(pp, v)

	w.WriteString("\n}")
}

func (f *StarredExpression) printType(w writer, v bool) {
	pp := w.Indent()

	pp.WriteString("StarredExpression {")

	if f.Expression != nil {
		pp.WriteString("\nExpression: ")
		f.Expression.printType(pp, v)
	} else if v {
		pp.WriteString("\nExpression: nil")
	}

	if f.StarredList != nil {
		pp.WriteString("\nStarredList: ")
		f.StarredList.printType(pp, v)
	} else if v {
		pp.WriteString("\nStarredList: nil")
	}

	pp.WriteString("\nTokens: ")
	f.Tokens.printType(pp, v)

	w.WriteString("\n}")
}

func (f *StarredItem) printType(w writer, v bool) {
	pp := w.Indent()

	pp.WriteString("StarredItem {")

	if f.AssignmentExpression != nil {
		pp.WriteString("\nAssignmentExpression: ")
		f.AssignmentExpression.printType(pp, v)
	} else if v {
		pp.WriteString("\nAssignmentExpression: nil")
	}

	if f.OrExpr != nil {
		pp.WriteString("\nOrExpr: ")
		f.OrExpr.printType(pp, v)
	} else if v {
		pp.WriteString("\nOrExpr: nil")
	}

	pp.WriteString("\nTokens: ")
	f.Tokens.printType(pp, v)

	w.WriteString("\n}")
}

func (f *StarredList) printType(w writer, v bool) {
	pp := w.Indent()

	pp.WriteString("StarredList {")

	if f.StarredItems == nil {
		pp.WriteString("\nStarredItems: nil")
	} else if len(f.StarredItems) > 0 {
		pp.WriteString("\nStarredItems: [")

		ipp := pp.Indent()

		for n, e := range f.StarredItems {
			ipp.Printf("\n%d: ", n)
			e.printType(ipp, v)
		}

		pp.WriteString("\n]")
	} else if v {
		pp.WriteString("\nStarredItems: []")
	}

	if f.TrailingComma || v {
		pp.Printf("\nTrailingComma: %v", f.TrailingComma)
	}

	pp.WriteString("\nTokens: ")
	f.Tokens.printType(pp, v)

	w.WriteString("\n}")
}

func (f *StarredOrKeyword) printType(w writer, v bool) {
	pp := w.Indent()

	pp.WriteString("StarredOrKeyword {")

	if f.Expression != nil {
		pp.WriteString("\nExpression: ")
		f.Expression.printType(pp, v)
	} else if v {
		pp.WriteString("\nExpression: nil")
	}

	if f.KeywordItem != nil {
		pp.WriteString("\nKeywordItem: ")
		f.KeywordItem.printType(pp, v)
	} else if v {
		pp.WriteString("\nKeywordItem: nil")
	}

	pp.WriteString("\nTokens: ")
	f.Tokens.printType(pp, v)

	w.WriteString("\n}")
}

func (f *Statement) printType(w writer, v bool) {
	pp := w.Indent()

	pp.WriteString("Statement {")

	if f.StatementList != nil {
		pp.WriteString("\nStatementList: ")
		f.StatementList.printType(pp, v)
	} else if v {
		pp.WriteString("\nStatementList: nil")
	}

	if f.CompoundStatement != nil {
		pp.WriteString("\nCompoundStatement: ")
		f.CompoundStatement.printType(pp, v)
	} else if v {
		pp.WriteString("\nCompoundStatement: nil")
	}

	pp.WriteString("\nComments: ")
	f.Comments.printType(pp, v)

	pp.WriteString("\nTokens: ")
	f.Tokens.printType(pp, v)

	w.WriteString("\n}")
}

func (f *StatementList) printType(w writer, v bool) {
	pp := w.Indent()

	pp.WriteString("StatementList {")

	if f.Statements == nil {
		pp.WriteString("\nStatements: nil")
	} else if len(f.Statements) > 0 {
		pp.WriteString("\nStatements: [")

		ipp := pp.Indent()

		for n, e := range f.Statements {
			ipp.Printf("\n%d: ", n)
			e.printType(ipp, v)
		}

		pp.WriteString("\n]")
	} else if v {
		pp.WriteString("\nStatements: []")
	}

	pp.WriteString("\nComments: ")
	f.Comments.printType(pp, v)

	pp.WriteString("\nTokens: ")
	f.Tokens.printType(pp, v)

	w.WriteString("\n}")
}

func (f *Suite) printType(w writer, v bool) {
	pp := w.Indent()

	pp.WriteString("Suite {")

	if f.StatementList != nil {
		pp.WriteString("\nStatementList: ")
		f.StatementList.printType(pp, v)
	} else if v {
		pp.WriteString("\nStatementList: nil")
	}

	if f.Statements == nil {
		pp.WriteString("\nStatements: nil")
	} else if len(f.Statements) > 0 {
		pp.WriteString("\nStatements: [")

		ipp := pp.Indent()

		for n, e := range f.Statements {
			ipp.Printf("\n%d: ", n)
			e.printType(ipp, v)
		}

		pp.WriteString("\n]")
	} else if v {
		pp.WriteString("\nStatements: []")
	}

	pp.WriteString("\nComments: [")

	ipp := pp.Indent()

	for n, e := range f.Comments {
		ipp.Printf("\n%d: ", n)
		e.printType(ipp, v)
	}

	pp.WriteString("\n]")

	pp.WriteString("\nTokens: ")
	f.Tokens.printType(pp, v)

	w.WriteString("\n}")
}

func (f *Target) printType(w writer, v bool) {
	pp := w.Indent()

	pp.WriteString("Target {")

	if f.PrimaryExpression != nil {
		pp.WriteString("\nPrimaryExpression: ")
		f.PrimaryExpression.printType(pp, v)
	} else if v {
		pp.WriteString("\nPrimaryExpression: nil")
	}

	if f.Tuple != nil {
		pp.WriteString("\nTuple: ")
		f.Tuple.printType(pp, v)
	} else if v {
		pp.WriteString("\nTuple: nil")
	}

	if f.Array != nil {
		pp.WriteString("\nArray: ")
		f.Array.printType(pp, v)
	} else if v {
		pp.WriteString("\nArray: nil")
	}

	if f.Star != nil {
		pp.WriteString("\nStar: ")
		f.Star.printType(pp, v)
	} else if v {
		pp.WriteString("\nStar: nil")
	}

	pp.WriteString("\nTokens: ")
	f.Tokens.printType(pp, v)

	w.WriteString("\n}")
}

func (f *TargetList) printType(w writer, v bool) {
	pp := w.Indent()

	pp.WriteString("TargetList {")

	if f.Targets == nil {
		pp.WriteString("\nTargets: nil")
	} else if len(f.Targets) > 0 {
		pp.WriteString("\nTargets: [")

		ipp := pp.Indent()

		for n, e := range f.Targets {
			ipp.Printf("\n%d: ", n)
			e.printType(ipp, v)
		}

		pp.WriteString("\n]")
	} else if v {
		pp.WriteString("\nTargets: []")
	}

	pp.WriteString("\nComments: [")

	ipp := pp.Indent()

	for n, e := range f.Comments {
		ipp.Printf("\n%d: ", n)
		e.printType(ipp, v)
	}

	pp.WriteString("\n]")

	pp.WriteString("\nTokens: ")
	f.Tokens.printType(pp, v)

	w.WriteString("\n}")
}

func (f *TryStatement) printType(w writer, v bool) {
	pp := w.Indent()

	pp.WriteString("TryStatement {")

	pp.WriteString("\nTry: ")
	f.Try.printType(pp, v)

	if f.Groups || v {
		pp.Printf("\nGroups: %v", f.Groups)
	}

	if f.Except == nil {
		pp.WriteString("\nExcept: nil")
	} else if len(f.Except) > 0 {
		pp.WriteString("\nExcept: [")

		ipp := pp.Indent()

		for n, e := range f.Except {
			ipp.Printf("\n%d: ", n)
			e.printType(ipp, v)
		}

		pp.WriteString("\n]")
	} else if v {
		pp.WriteString("\nExcept: []")
	}

	if f.Else != nil {
		pp.WriteString("\nElse: ")
		f.Else.printType(pp, v)
	} else if v {
		pp.WriteString("\nElse: nil")
	}

	if f.Finally != nil {
		pp.WriteString("\nFinally: ")
		f.Finally.printType(pp, v)
	} else if v {
		pp.WriteString("\nFinally: nil")
	}

	pp.WriteString("\nTokens: ")
	f.Tokens.printType(pp, v)

	w.WriteString("\n}")
}

func (f *TypeParam) printType(w writer, v bool) {
	pp := w.Indent()

	pp.WriteString("TypeParam {")

	pp.WriteString("\nType: ")
	f.Type.printType(pp, v)

	if f.Identifier != nil {
		pp.WriteString("\nIdentifier: ")
		f.Identifier.printType(pp, v)
	} else if v {
		pp.WriteString("\nIdentifier: nil")
	}

	if f.Expression != nil {
		pp.WriteString("\nExpression: ")
		f.Expression.printType(pp, v)
	} else if v {
		pp.WriteString("\nExpression: nil")
	}

	pp.WriteString("\nComments: [")

	ipp := pp.Indent()

	for n, e := range f.Comments {
		ipp.Printf("\n%d: ", n)
		e.printType(ipp, v)
	}

	pp.WriteString("\n]")

	pp.WriteString("\nTokens: ")
	f.Tokens.printType(pp, v)

	w.WriteString("\n}")
}

func (f *TypeParams) printType(w writer, v bool) {
	pp := w.Indent()

	pp.WriteString("TypeParams {")

	if f.TypeParams == nil {
		pp.WriteString("\nTypeParams: nil")
	} else if len(f.TypeParams) > 0 {
		pp.WriteString("\nTypeParams: [")

		ipp := pp.Indent()

		for n, e := range f.TypeParams {
			ipp.Printf("\n%d: ", n)
			e.printType(ipp, v)
		}

		pp.WriteString("\n]")
	} else if v {
		pp.WriteString("\nTypeParams: []")
	}

	pp.WriteString("\nComments: [")

	ipp := pp.Indent()

	for n, e := range f.Comments {
		ipp.Printf("\n%d: ", n)
		e.printType(ipp, v)
	}

	pp.WriteString("\n]")

	pp.WriteString("\nTokens: ")
	f.Tokens.printType(pp, v)

	w.WriteString("\n}")
}

func (f *TypeStatement) printType(w writer, v bool) {
	pp := w.Indent()

	pp.WriteString("TypeStatement {")

	if f.Identifier != nil {
		pp.WriteString("\nIdentifier: ")
		f.Identifier.printType(pp, v)
	} else if v {
		pp.WriteString("\nIdentifier: nil")
	}

	if f.TypeParams != nil {
		pp.WriteString("\nTypeParams: ")
		f.TypeParams.printType(pp, v)
	} else if v {
		pp.WriteString("\nTypeParams: nil")
	}

	pp.WriteString("\nExpression: ")
	f.Expression.printType(pp, v)

	pp.WriteString("\nTokens: ")
	f.Tokens.printType(pp, v)

	w.WriteString("\n}")
}

func (f *UnaryExpression) printType(w writer, v bool) {
	pp := w.Indent()

	pp.WriteString("UnaryExpression {")

	if f.PowerExpression != nil {
		pp.WriteString("\nPowerExpression: ")
		f.PowerExpression.printType(pp, v)
	} else if v {
		pp.WriteString("\nPowerExpression: nil")
	}

	if f.Unary != nil {
		pp.WriteString("\nUnary: ")
		f.Unary.printType(pp, v)
	} else if v {
		pp.WriteString("\nUnary: nil")
	}

	if f.UnaryExpression != nil {
		pp.WriteString("\nUnaryExpression: ")
		f.UnaryExpression.printType(pp, v)
	} else if v {
		pp.WriteString("\nUnaryExpression: nil")
	}

	pp.WriteString("\nTokens: ")
	f.Tokens.printType(pp, v)

	w.WriteString("\n}")
}

func (f *WhileStatement) printType(w writer, v bool) {
	pp := w.Indent()

	pp.WriteString("WhileStatement {")

	pp.WriteString("\nAssignmentExpression: ")
	f.AssignmentExpression.printType(pp, v)

	pp.WriteString("\nSuite: ")
	f.Suite.printType(pp, v)

	if f.Else != nil {
		pp.WriteString("\nElse: ")
		f.Else.printType(pp, v)
	} else if v {
		pp.WriteString("\nElse: nil")
	}

	pp.WriteString("\nTokens: ")
	f.Tokens.printType(pp, v)

	w.WriteString("\n}")
}

func (f *WithItem) printType(w writer, v bool) {
	pp := w.Indent()

	pp.WriteString("WithItem {")

	pp.WriteString("\nExpression: ")
	f.Expression.printType(pp, v)

	if f.Target != nil {
		pp.WriteString("\nTarget: ")
		f.Target.printType(pp, v)
	} else if v {
		pp.WriteString("\nTarget: nil")
	}

	pp.WriteString("\nTokens: ")
	f.Tokens.printType(pp, v)

	w.WriteString("\n}")
}

func (f *WithStatement) printType(w writer, v bool) {
	pp := w.Indent()

	pp.WriteString("WithStatement {")

	if f.Async || v {
		pp.Printf("\nAsync: %v", f.Async)
	}

	pp.WriteString("\nContents: ")
	f.Contents.printType(pp, v)

	pp.WriteString("\nSuite: ")
	f.Suite.printType(pp, v)

	pp.WriteString("\nTokens: ")
	f.Tokens.printType(pp, v)

	w.WriteString("\n}")
}

func (f *WithStatementContents) printType(w writer, v bool) {
	pp := w.Indent()

	pp.WriteString("WithStatementContents {")

	if f.Items == nil {
		pp.WriteString("\nItems: nil")
	} else if len(f.Items) > 0 {
		pp.WriteString("\nItems: [")

		ipp := pp.Indent()

		for n, e := range f.Items {
			ipp.Printf("\n%d: ", n)
			e.printType(ipp, v)
		}

		pp.WriteString("\n]")
	} else if v {
		pp.WriteString("\nItems: []")
	}

	pp.WriteString("\nTokens: ")
	f.Tokens.printType(pp, v)

	w.WriteString("\n}")
}

func (f *XorExpression) printType(w writer, v bool) {
	pp := w.Indent()

	pp.WriteString("XorExpression {")

	pp.WriteString("\nAndExpression: ")
	f.AndExpression.printType(pp, v)

	if f.XorExpression != nil {
		pp.WriteString("\nXorExpression: ")
		f.XorExpression.printType(pp, v)
	} else if v {
		pp.WriteString("\nXorExpression: nil")
	}

	pp.WriteString("\nTokens: ")
	f.Tokens.printType(pp, v)

	w.WriteString("\n}")
}

func (f *YieldExpression) printType(w writer, v bool) {
	pp := w.Indent()

	pp.WriteString("YieldExpression {")

	if f.ExpressionList != nil {
		pp.WriteString("\nExpressionList: ")
		f.ExpressionList.printType(pp, v)
	} else if v {
		pp.WriteString("\nExpressionList: nil")
	}

	if f.From != nil {
		pp.WriteString("\nFrom: ")
		f.From.printType(pp, v)
	} else if v {
		pp.WriteString("\nFrom: nil")
	}

	pp.WriteString("\nComments: [")

	ipp := pp.Indent()

	for n, e := range f.Comments {
		ipp.Printf("\n%d: ", n)
		e.printType(ipp, v)
	}

	pp.WriteString("\n]")

	pp.WriteString("\nTokens: ")
	f.Tokens.printType(pp, v)

	w.WriteString("\n}")
}

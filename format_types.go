package python

// File automatically generated with format.sh.

import "io"

func (f *AddExpression) printType(w io.Writer, v bool) {
	pp := indentPrinter{w}

	pp.Print("AddExpression {")

	pp.Print("\nMultiplyExpression: ")
	f.MultiplyExpression.printType(&pp, v)

	if f.Add != nil {
		pp.Print("\nAdd: ")
		f.Add.printType(&pp, v)
	} else if v {
		pp.Print("\nAdd: nil")
	}

	if f.AddExpression != nil {
		pp.Print("\nAddExpression: ")
		f.AddExpression.printType(&pp, v)
	} else if v {
		pp.Print("\nAddExpression: nil")
	}

	pp.Print("\nTokens: ")
	f.Tokens.printType(&pp, v)

	io.WriteString(w, "\n}")
}

func (f *AndExpression) printType(w io.Writer, v bool) {
	pp := indentPrinter{w}

	pp.Print("AndExpression {")

	pp.Print("\nShiftExpression: ")
	f.ShiftExpression.printType(&pp, v)

	if f.AndExpression != nil {
		pp.Print("\nAndExpression: ")
		f.AndExpression.printType(&pp, v)
	} else if v {
		pp.Print("\nAndExpression: nil")
	}

	pp.Print("\nTokens: ")
	f.Tokens.printType(&pp, v)

	io.WriteString(w, "\n}")
}

func (f *AndTest) printType(w io.Writer, v bool) {
	pp := indentPrinter{w}

	pp.Print("AndTest {")

	pp.Print("\nNotTest: ")
	f.NotTest.printType(&pp, v)

	if f.AndTest != nil {
		pp.Print("\nAndTest: ")
		f.AndTest.printType(&pp, v)
	} else if v {
		pp.Print("\nAndTest: nil")
	}

	pp.Print("\nTokens: ")
	f.Tokens.printType(&pp, v)

	io.WriteString(w, "\n}")
}

func (f *AnnotatedAssignmentStatement) printType(w io.Writer, v bool) {
	pp := indentPrinter{w}

	pp.Print("AnnotatedAssignmentStatement {")

	pp.Print("\nAugTarget: ")
	f.AugTarget.printType(&pp, v)

	pp.Print("\nExpression: ")
	f.Expression.printType(&pp, v)

	if f.StarredExpression != nil {
		pp.Print("\nStarredExpression: ")
		f.StarredExpression.printType(&pp, v)
	} else if v {
		pp.Print("\nStarredExpression: nil")
	}

	if f.YieldExpression != nil {
		pp.Print("\nYieldExpression: ")
		f.YieldExpression.printType(&pp, v)
	} else if v {
		pp.Print("\nYieldExpression: nil")
	}

	pp.Print("\nTokens: ")
	f.Tokens.printType(&pp, v)

	io.WriteString(w, "\n}")
}

func (f *ArgumentList) printType(w io.Writer, v bool) {
	pp := indentPrinter{w}

	pp.Print("ArgumentList {")

	if f.PositionalArguments == nil {
		pp.Print("\nPositionalArguments: nil")
	} else if len(f.PositionalArguments) > 0 {
		pp.Print("\nPositionalArguments: [")

		ipp := indentPrinter{&pp}

		for n, e := range f.PositionalArguments {
			ipp.Printf("\n%d: ", n)
			e.printType(&ipp, v)
		}

		pp.Print("\n]")
	} else if v {
		pp.Print("\nPositionalArguments: []")
	}

	if f.StarredAndKeywordArguments == nil {
		pp.Print("\nStarredAndKeywordArguments: nil")
	} else if len(f.StarredAndKeywordArguments) > 0 {
		pp.Print("\nStarredAndKeywordArguments: [")

		ipp := indentPrinter{&pp}

		for n, e := range f.StarredAndKeywordArguments {
			ipp.Printf("\n%d: ", n)
			e.printType(&ipp, v)
		}

		pp.Print("\n]")
	} else if v {
		pp.Print("\nStarredAndKeywordArguments: []")
	}

	if f.KeywordArguments == nil {
		pp.Print("\nKeywordArguments: nil")
	} else if len(f.KeywordArguments) > 0 {
		pp.Print("\nKeywordArguments: [")

		ipp := indentPrinter{&pp}

		for n, e := range f.KeywordArguments {
			ipp.Printf("\n%d: ", n)
			e.printType(&ipp, v)
		}

		pp.Print("\n]")
	} else if v {
		pp.Print("\nKeywordArguments: []")
	}

	pp.Print("\nTokens: ")
	f.Tokens.printType(&pp, v)

	io.WriteString(w, "\n}")
}

func (f *ArgumentListOrComprehension) printType(w io.Writer, v bool) {
	pp := indentPrinter{w}

	pp.Print("ArgumentListOrComprehension {")

	if f.ArgumentList != nil {
		pp.Print("\nArgumentList: ")
		f.ArgumentList.printType(&pp, v)
	} else if v {
		pp.Print("\nArgumentList: nil")
	}

	if f.Comprehension != nil {
		pp.Print("\nComprehension: ")
		f.Comprehension.printType(&pp, v)
	} else if v {
		pp.Print("\nComprehension: nil")
	}

	pp.Print("\nTokens: ")
	f.Tokens.printType(&pp, v)

	io.WriteString(w, "\n}")
}

func (f *AssertStatement) printType(w io.Writer, v bool) {
	pp := indentPrinter{w}

	pp.Print("AssertStatement {")

	if f.Expressions == nil {
		pp.Print("\nExpressions: nil")
	} else if len(f.Expressions) > 0 {
		pp.Print("\nExpressions: [")

		ipp := indentPrinter{&pp}

		for n, e := range f.Expressions {
			ipp.Printf("\n%d: ", n)
			e.printType(&ipp, v)
		}

		pp.Print("\n]")
	} else if v {
		pp.Print("\nExpressions: []")
	}

	pp.Print("\nTokens: ")
	f.Tokens.printType(&pp, v)

	io.WriteString(w, "\n}")
}

func (f *AssignmentExpressionAndSuite) printType(w io.Writer, v bool) {
	pp := indentPrinter{w}

	pp.Print("AssignmentExpressionAndSuite {")

	pp.Print("\nAssignmentExpression: ")
	f.AssignmentExpression.printType(&pp, v)

	pp.Print("\nSuite: ")
	f.Suite.printType(&pp, v)

	pp.Print("\nTokens: ")
	f.Tokens.printType(&pp, v)

	io.WriteString(w, "\n}")
}

func (f *AssignmentExpression) printType(w io.Writer, v bool) {
	pp := indentPrinter{w}

	pp.Print("AssignmentExpression {")

	if f.Identifier != nil {
		pp.Print("\nIdentifier: ")
		f.Identifier.printType(&pp, v)
	} else if v {
		pp.Print("\nIdentifier: nil")
	}

	pp.Print("\nExpression: ")
	f.Expression.printType(&pp, v)

	pp.Print("\nTokens: ")
	f.Tokens.printType(&pp, v)

	io.WriteString(w, "\n}")
}

func (f *AssignmentStatement) printType(w io.Writer, v bool) {
	pp := indentPrinter{w}

	pp.Print("AssignmentStatement {")

	if f.TargetLists == nil {
		pp.Print("\nTargetLists: nil")
	} else if len(f.TargetLists) > 0 {
		pp.Print("\nTargetLists: [")

		ipp := indentPrinter{&pp}

		for n, e := range f.TargetLists {
			ipp.Printf("\n%d: ", n)
			e.printType(&ipp, v)
		}

		pp.Print("\n]")
	} else if v {
		pp.Print("\nTargetLists: []")
	}

	if f.StarredExpression != nil {
		pp.Print("\nStarredExpression: ")
		f.StarredExpression.printType(&pp, v)
	} else if v {
		pp.Print("\nStarredExpression: nil")
	}

	if f.YieldExpression != nil {
		pp.Print("\nYieldExpression: ")
		f.YieldExpression.printType(&pp, v)
	} else if v {
		pp.Print("\nYieldExpression: nil")
	}

	pp.Print("\nTokens: ")
	f.Tokens.printType(&pp, v)

	io.WriteString(w, "\n}")
}

func (f *Atom) printType(w io.Writer, v bool) {
	pp := indentPrinter{w}

	pp.Print("Atom {")

	if f.Identifier != nil {
		pp.Print("\nIdentifier: ")
		f.Identifier.printType(&pp, v)
	} else if v {
		pp.Print("\nIdentifier: nil")
	}

	if f.Literal != nil {
		pp.Print("\nLiteral: ")
		f.Literal.printType(&pp, v)
	} else if v {
		pp.Print("\nLiteral: nil")
	}

	if f.Enclosure != nil {
		pp.Print("\nEnclosure: ")
		f.Enclosure.printType(&pp, v)
	} else if v {
		pp.Print("\nEnclosure: nil")
	}

	pp.Print("\nTokens: ")
	f.Tokens.printType(&pp, v)

	io.WriteString(w, "\n}")
}

func (f *AugmentedAssignmentStatement) printType(w io.Writer, v bool) {
	pp := indentPrinter{w}

	pp.Print("AugmentedAssignmentStatement {")

	pp.Print("\nAugTarget: ")
	f.AugTarget.printType(&pp, v)

	if f.AugOp != nil {
		pp.Print("\nAugOp: ")
		f.AugOp.printType(&pp, v)
	} else if v {
		pp.Print("\nAugOp: nil")
	}

	if f.ExpressionList != nil {
		pp.Print("\nExpressionList: ")
		f.ExpressionList.printType(&pp, v)
	} else if v {
		pp.Print("\nExpressionList: nil")
	}

	if f.YieldExpression != nil {
		pp.Print("\nYieldExpression: ")
		f.YieldExpression.printType(&pp, v)
	} else if v {
		pp.Print("\nYieldExpression: nil")
	}

	pp.Print("\nTokens: ")
	f.Tokens.printType(&pp, v)

	io.WriteString(w, "\n}")
}

func (f *AugTarget) printType(w io.Writer, v bool) {
	pp := indentPrinter{w}

	pp.Print("AugTarget {")

	pp.Print("\nPrimaryExpression: ")
	f.PrimaryExpression.printType(&pp, v)

	pp.Print("\nTokens: ")
	f.Tokens.printType(&pp, v)

	io.WriteString(w, "\n}")
}

func (f *ClassDefinition) printType(w io.Writer, v bool) {
	pp := indentPrinter{w}

	pp.Print("ClassDefinition {")

	if f.Decorators != nil {
		pp.Print("\nDecorators: ")
		f.Decorators.printType(&pp, v)
	} else if v {
		pp.Print("\nDecorators: nil")
	}

	if f.ClassName != nil {
		pp.Print("\nClassName: ")
		f.ClassName.printType(&pp, v)
	} else if v {
		pp.Print("\nClassName: nil")
	}

	if f.TypeParams != nil {
		pp.Print("\nTypeParams: ")
		f.TypeParams.printType(&pp, v)
	} else if v {
		pp.Print("\nTypeParams: nil")
	}

	pp.Print("\nInheritance: ")
	f.Inheritance.printType(&pp, v)

	pp.Print("\nSuite: ")
	f.Suite.printType(&pp, v)

	pp.Print("\nTokens: ")
	f.Tokens.printType(&pp, v)

	io.WriteString(w, "\n}")
}

func (f *Comparison) printType(w io.Writer, v bool) {
	pp := indentPrinter{w}

	pp.Print("Comparison {")

	pp.Print("\nOrExpression: ")
	f.OrExpression.printType(&pp, v)

	if f.Comparisons == nil {
		pp.Print("\nComparisons: nil")
	} else if len(f.Comparisons) > 0 {
		pp.Print("\nComparisons: [")

		ipp := indentPrinter{&pp}

		for n, e := range f.Comparisons {
			ipp.Printf("\n%d: ", n)
			e.printType(&ipp, v)
		}

		pp.Print("\n]")
	} else if v {
		pp.Print("\nComparisons: []")
	}

	pp.Print("\nTokens: ")
	f.Tokens.printType(&pp, v)

	io.WriteString(w, "\n}")
}

func (f *ComparisonExpression) printType(w io.Writer, v bool) {
	pp := indentPrinter{w}

	pp.Print("ComparisonExpression {")

	if f.ComparisonOperator == nil {
		pp.Print("\nComparisonOperator: nil")
	} else if len(f.ComparisonOperator) > 0 {
		pp.Print("\nComparisonOperator: [")

		ipp := indentPrinter{&pp}

		for n, e := range f.ComparisonOperator {
			ipp.Printf("\n%d: ", n)
			e.printType(&ipp, v)
		}

		pp.Print("\n]")
	} else if v {
		pp.Print("\nComparisonOperator: []")
	}

	pp.Print("\nOrExpression: ")
	f.OrExpression.printType(&pp, v)

	io.WriteString(w, "\n}")
}

func (f *CompoundStatement) printType(w io.Writer, v bool) {
	pp := indentPrinter{w}

	pp.Print("CompoundStatement {")

	if f.If != nil {
		pp.Print("\nIf: ")
		f.If.printType(&pp, v)
	} else if v {
		pp.Print("\nIf: nil")
	}

	if f.While != nil {
		pp.Print("\nWhile: ")
		f.While.printType(&pp, v)
	} else if v {
		pp.Print("\nWhile: nil")
	}

	if f.For != nil {
		pp.Print("\nFor: ")
		f.For.printType(&pp, v)
	} else if v {
		pp.Print("\nFor: nil")
	}

	if f.Try != nil {
		pp.Print("\nTry: ")
		f.Try.printType(&pp, v)
	} else if v {
		pp.Print("\nTry: nil")
	}

	if f.With != nil {
		pp.Print("\nWith: ")
		f.With.printType(&pp, v)
	} else if v {
		pp.Print("\nWith: nil")
	}

	if f.Func != nil {
		pp.Print("\nFunc: ")
		f.Func.printType(&pp, v)
	} else if v {
		pp.Print("\nFunc: nil")
	}

	if f.Class != nil {
		pp.Print("\nClass: ")
		f.Class.printType(&pp, v)
	} else if v {
		pp.Print("\nClass: nil")
	}

	pp.Print("\nTokens: ")
	f.Tokens.printType(&pp, v)

	io.WriteString(w, "\n}")
}

func (f *Comprehension) printType(w io.Writer, v bool) {
	pp := indentPrinter{w}

	pp.Print("Comprehension {")

	pp.Print("\nAssignmentExpression: ")
	f.AssignmentExpression.printType(&pp, v)

	pp.Print("\nComprehensionFor: ")
	f.ComprehensionFor.printType(&pp, v)

	pp.Print("\nTokens: ")
	f.Tokens.printType(&pp, v)

	io.WriteString(w, "\n}")
}

func (f *ComprehensionFor) printType(w io.Writer, v bool) {
	pp := indentPrinter{w}

	pp.Print("ComprehensionFor {")

	if f.Async || v {
		pp.Printf("\nAsync: %v", f.Async)
	}

	pp.Print("\nTargetList: ")
	f.TargetList.printType(&pp, v)

	pp.Print("\nOrTest: ")
	f.OrTest.printType(&pp, v)

	if f.ComprehensionIterator != nil {
		pp.Print("\nComprehensionIterator: ")
		f.ComprehensionIterator.printType(&pp, v)
	} else if v {
		pp.Print("\nComprehensionIterator: nil")
	}

	pp.Print("\nTokens: ")
	f.Tokens.printType(&pp, v)

	io.WriteString(w, "\n}")
}

func (f *ComprehensionIf) printType(w io.Writer, v bool) {
	pp := indentPrinter{w}

	pp.Print("ComprehensionIf {")

	pp.Print("\nOrTest: ")
	f.OrTest.printType(&pp, v)

	if f.ComprehensionIterator != nil {
		pp.Print("\nComprehensionIterator: ")
		f.ComprehensionIterator.printType(&pp, v)
	} else if v {
		pp.Print("\nComprehensionIterator: nil")
	}

	pp.Print("\nTokens: ")
	f.Tokens.printType(&pp, v)

	io.WriteString(w, "\n}")
}

func (f *ComprehensionIterator) printType(w io.Writer, v bool) {
	pp := indentPrinter{w}

	pp.Print("ComprehensionIterator {")

	if f.ComprehensionFor != nil {
		pp.Print("\nComprehensionFor: ")
		f.ComprehensionFor.printType(&pp, v)
	} else if v {
		pp.Print("\nComprehensionFor: nil")
	}

	if f.ComprehensionIf != nil {
		pp.Print("\nComprehensionIf: ")
		f.ComprehensionIf.printType(&pp, v)
	} else if v {
		pp.Print("\nComprehensionIf: nil")
	}

	pp.Print("\nTokens: ")
	f.Tokens.printType(&pp, v)

	io.WriteString(w, "\n}")
}

func (f *ConditionalExpression) printType(w io.Writer, v bool) {
	pp := indentPrinter{w}

	pp.Print("ConditionalExpression {")

	pp.Print("\nOrTest: ")
	f.OrTest.printType(&pp, v)

	if f.If != nil {
		pp.Print("\nIf: ")
		f.If.printType(&pp, v)
	} else if v {
		pp.Print("\nIf: nil")
	}

	if f.Else != nil {
		pp.Print("\nElse: ")
		f.Else.printType(&pp, v)
	} else if v {
		pp.Print("\nElse: nil")
	}

	pp.Print("\nTokens: ")
	f.Tokens.printType(&pp, v)

	io.WriteString(w, "\n}")
}

func (f *Decorators) printType(w io.Writer, v bool) {
	pp := indentPrinter{w}

	pp.Print("Decorators {")

	if f.Decorators == nil {
		pp.Print("\nDecorators: nil")
	} else if len(f.Decorators) > 0 {
		pp.Print("\nDecorators: [")

		ipp := indentPrinter{&pp}

		for n, e := range f.Decorators {
			ipp.Printf("\n%d: ", n)
			e.printType(&ipp, v)
		}

		pp.Print("\n]")
	} else if v {
		pp.Print("\nDecorators: []")
	}

	pp.Print("\nTokens: ")
	f.Tokens.printType(&pp, v)

	io.WriteString(w, "\n}")
}

func (f *DefParameter) printType(w io.Writer, v bool) {
	pp := indentPrinter{w}

	pp.Print("DefParameter {")

	pp.Print("\nParameter: ")
	f.Parameter.printType(&pp, v)

	if f.Value != nil {
		pp.Print("\nValue: ")
		f.Value.printType(&pp, v)
	} else if v {
		pp.Print("\nValue: nil")
	}

	pp.Print("\nTokens: ")
	f.Tokens.printType(&pp, v)

	io.WriteString(w, "\n}")
}

func (f *DelStatement) printType(w io.Writer, v bool) {
	pp := indentPrinter{w}

	pp.Print("DelStatement {")

	pp.Print("\nTargetList: ")
	f.TargetList.printType(&pp, v)

	pp.Print("\nTokens: ")
	f.Tokens.printType(&pp, v)

	io.WriteString(w, "\n}")
}

func (f *DictDisplay) printType(w io.Writer, v bool) {
	pp := indentPrinter{w}

	pp.Print("DictDisplay {")

	if f.DictItems == nil {
		pp.Print("\nDictItems: nil")
	} else if len(f.DictItems) > 0 {
		pp.Print("\nDictItems: [")

		ipp := indentPrinter{&pp}

		for n, e := range f.DictItems {
			ipp.Printf("\n%d: ", n)
			e.printType(&ipp, v)
		}

		pp.Print("\n]")
	} else if v {
		pp.Print("\nDictItems: []")
	}

	if f.DictComprehension != nil {
		pp.Print("\nDictComprehension: ")
		f.DictComprehension.printType(&pp, v)
	} else if v {
		pp.Print("\nDictComprehension: nil")
	}

	pp.Print("\nTokens: ")
	f.Tokens.printType(&pp, v)

	io.WriteString(w, "\n}")
}

func (f *DictItem) printType(w io.Writer, v bool) {
	pp := indentPrinter{w}

	pp.Print("DictItem {")

	if f.Key != nil {
		pp.Print("\nKey: ")
		f.Key.printType(&pp, v)
	} else if v {
		pp.Print("\nKey: nil")
	}

	if f.Value != nil {
		pp.Print("\nValue: ")
		f.Value.printType(&pp, v)
	} else if v {
		pp.Print("\nValue: nil")
	}

	if f.OrExpression != nil {
		pp.Print("\nOrExpression: ")
		f.OrExpression.printType(&pp, v)
	} else if v {
		pp.Print("\nOrExpression: nil")
	}

	pp.Print("\nTokens: ")
	f.Tokens.printType(&pp, v)

	io.WriteString(w, "\n}")
}

func (f *Enclosure) printType(w io.Writer, v bool) {
	pp := indentPrinter{w}

	pp.Print("Enclosure {")

	if f.ParenthForm != nil {
		pp.Print("\nParenthForm: ")
		f.ParenthForm.printType(&pp, v)
	} else if v {
		pp.Print("\nParenthForm: nil")
	}

	if f.ListDisplay != nil {
		pp.Print("\nListDisplay: ")
		f.ListDisplay.printType(&pp, v)
	} else if v {
		pp.Print("\nListDisplay: nil")
	}

	if f.DictDisplay != nil {
		pp.Print("\nDictDisplay: ")
		f.DictDisplay.printType(&pp, v)
	} else if v {
		pp.Print("\nDictDisplay: nil")
	}

	if f.SetDisplay != nil {
		pp.Print("\nSetDisplay: ")
		f.SetDisplay.printType(&pp, v)
	} else if v {
		pp.Print("\nSetDisplay: nil")
	}

	if f.GeneratorExpression != nil {
		pp.Print("\nGeneratorExpression: ")
		f.GeneratorExpression.printType(&pp, v)
	} else if v {
		pp.Print("\nGeneratorExpression: nil")
	}

	if f.YieldAtom != nil {
		pp.Print("\nYieldAtom: ")
		f.YieldAtom.printType(&pp, v)
	} else if v {
		pp.Print("\nYieldAtom: nil")
	}

	pp.Print("\nTokens: ")
	f.Tokens.printType(&pp, v)

	io.WriteString(w, "\n}")
}

func (f *Except) printType(w io.Writer, v bool) {
	pp := indentPrinter{w}

	pp.Print("Except {")

	pp.Print("\nExpression: ")
	f.Expression.printType(&pp, v)

	if f.Identifier != nil {
		pp.Print("\nIdentifier: ")
		f.Identifier.printType(&pp, v)
	} else if v {
		pp.Print("\nIdentifier: nil")
	}

	pp.Print("\nSuite: ")
	f.Suite.printType(&pp, v)

	pp.Print("\nTokens: ")
	f.Tokens.printType(&pp, v)

	io.WriteString(w, "\n}")
}

func (f *Expression) printType(w io.Writer, v bool) {
	pp := indentPrinter{w}

	pp.Print("Expression {")

	if f.ConditionalExpression != nil {
		pp.Print("\nConditionalExpression: ")
		f.ConditionalExpression.printType(&pp, v)
	} else if v {
		pp.Print("\nConditionalExpression: nil")
	}

	if f.LambdaExpression != nil {
		pp.Print("\nLambdaExpression: ")
		f.LambdaExpression.printType(&pp, v)
	} else if v {
		pp.Print("\nLambdaExpression: nil")
	}

	pp.Print("\nTokens: ")
	f.Tokens.printType(&pp, v)

	io.WriteString(w, "\n}")
}

func (f *ExpressionList) printType(w io.Writer, v bool) {
	pp := indentPrinter{w}

	pp.Print("ExpressionList {")

	if f.Expressions == nil {
		pp.Print("\nExpressions: nil")
	} else if len(f.Expressions) > 0 {
		pp.Print("\nExpressions: [")

		ipp := indentPrinter{&pp}

		for n, e := range f.Expressions {
			ipp.Printf("\n%d: ", n)
			e.printType(&ipp, v)
		}

		pp.Print("\n]")
	} else if v {
		pp.Print("\nExpressions: []")
	}

	pp.Print("\nTokens: ")
	f.Tokens.printType(&pp, v)

	io.WriteString(w, "\n}")
}

func (f *File) printType(w io.Writer, v bool) {
	pp := indentPrinter{w}

	pp.Print("File {")

	if f.Statements == nil {
		pp.Print("\nStatements: nil")
	} else if len(f.Statements) > 0 {
		pp.Print("\nStatements: [")

		ipp := indentPrinter{&pp}

		for n, e := range f.Statements {
			ipp.Printf("\n%d: ", n)
			e.printType(&ipp, v)
		}

		pp.Print("\n]")
	} else if v {
		pp.Print("\nStatements: []")
	}

	pp.Print("\nTokens: ")
	f.Tokens.printType(&pp, v)

	io.WriteString(w, "\n}")
}

func (f *FlexibleExpression) printType(w io.Writer, v bool) {
	pp := indentPrinter{w}

	pp.Print("FlexibleExpression {")

	if f.AssignmentExpression != nil {
		pp.Print("\nAssignmentExpression: ")
		f.AssignmentExpression.printType(&pp, v)
	} else if v {
		pp.Print("\nAssignmentExpression: nil")
	}

	if f.StarredExpression != nil {
		pp.Print("\nStarredExpression: ")
		f.StarredExpression.printType(&pp, v)
	} else if v {
		pp.Print("\nStarredExpression: nil")
	}

	pp.Print("\nTokens: ")
	f.Tokens.printType(&pp, v)

	io.WriteString(w, "\n}")
}

func (f *FlexibleExpressionList) printType(w io.Writer, v bool) {
	pp := indentPrinter{w}

	pp.Print("FlexibleExpressionList {")

	if f.FlexibleExpressions == nil {
		pp.Print("\nFlexibleExpressions: nil")
	} else if len(f.FlexibleExpressions) > 0 {
		pp.Print("\nFlexibleExpressions: [")

		ipp := indentPrinter{&pp}

		for n, e := range f.FlexibleExpressions {
			ipp.Printf("\n%d: ", n)
			e.printType(&ipp, v)
		}

		pp.Print("\n]")
	} else if v {
		pp.Print("\nFlexibleExpressions: []")
	}

	pp.Print("\nTokens: ")
	f.Tokens.printType(&pp, v)

	io.WriteString(w, "\n}")
}

func (f *FlexibleExpressionListOrComprehension) printType(w io.Writer, v bool) {
	pp := indentPrinter{w}

	pp.Print("FlexibleExpressionListOrComprehension {")

	if f.FlexibleExpressionList != nil {
		pp.Print("\nFlexibleExpressionList: ")
		f.FlexibleExpressionList.printType(&pp, v)
	} else if v {
		pp.Print("\nFlexibleExpressionList: nil")
	}

	if f.Comprehension != nil {
		pp.Print("\nComprehension: ")
		f.Comprehension.printType(&pp, v)
	} else if v {
		pp.Print("\nComprehension: nil")
	}

	pp.Print("\nTokens: ")
	f.Tokens.printType(&pp, v)

	io.WriteString(w, "\n}")
}

func (f *ForStatement) printType(w io.Writer, v bool) {
	pp := indentPrinter{w}

	pp.Print("ForStatement {")

	if f.Async || v {
		pp.Printf("\nAsync: %v", f.Async)
	}

	pp.Print("\nTargetList: ")
	f.TargetList.printType(&pp, v)

	pp.Print("\nStarredList: ")
	f.StarredList.printType(&pp, v)

	pp.Print("\nSuite: ")
	f.Suite.printType(&pp, v)

	if f.Else != nil {
		pp.Print("\nElse: ")
		f.Else.printType(&pp, v)
	} else if v {
		pp.Print("\nElse: nil")
	}

	pp.Print("\nTokens: ")
	f.Tokens.printType(&pp, v)

	io.WriteString(w, "\n}")
}

func (f *FuncDefinition) printType(w io.Writer, v bool) {
	pp := indentPrinter{w}

	pp.Print("FuncDefinition {")

	if f.Decorators != nil {
		pp.Print("\nDecorators: ")
		f.Decorators.printType(&pp, v)
	} else if v {
		pp.Print("\nDecorators: nil")
	}

	if f.Async || v {
		pp.Printf("\nAsync: %v", f.Async)
	}

	if f.FuncName != nil {
		pp.Print("\nFuncName: ")
		f.FuncName.printType(&pp, v)
	} else if v {
		pp.Print("\nFuncName: nil")
	}

	if f.TypeParams != nil {
		pp.Print("\nTypeParams: ")
		f.TypeParams.printType(&pp, v)
	} else if v {
		pp.Print("\nTypeParams: nil")
	}

	pp.Print("\nParameterList: ")
	f.ParameterList.printType(&pp, v)

	if f.Expression != nil {
		pp.Print("\nExpression: ")
		f.Expression.printType(&pp, v)
	} else if v {
		pp.Print("\nExpression: nil")
	}

	pp.Print("\nSuite: ")
	f.Suite.printType(&pp, v)

	pp.Print("\nTokens: ")
	f.Tokens.printType(&pp, v)

	io.WriteString(w, "\n}")
}

func (f *GeneratorExpression) printType(w io.Writer, v bool) {
	pp := indentPrinter{w}

	pp.Print("GeneratorExpression {")

	pp.Print("\nExpression: ")
	f.Expression.printType(&pp, v)

	pp.Print("\nComprehensionFor: ")
	f.ComprehensionFor.printType(&pp, v)

	pp.Print("\nTokens: ")
	f.Tokens.printType(&pp, v)

	io.WriteString(w, "\n}")
}

func (f *GlobalStatement) printType(w io.Writer, v bool) {
	pp := indentPrinter{w}

	pp.Print("GlobalStatement {")

	if f.Identifiers == nil {
		pp.Print("\nIdentifiers: nil")
	} else if len(f.Identifiers) > 0 {
		pp.Print("\nIdentifiers: [")

		ipp := indentPrinter{&pp}

		for n, e := range f.Identifiers {
			ipp.Printf("\n%d: ", n)
			e.printType(&ipp, v)
		}

		pp.Print("\n]")
	} else if v {
		pp.Print("\nIdentifiers: []")
	}

	pp.Print("\nTokens: ")
	f.Tokens.printType(&pp, v)

	io.WriteString(w, "\n}")
}

func (f *IfStatement) printType(w io.Writer, v bool) {
	pp := indentPrinter{w}

	pp.Print("IfStatement {")

	pp.Print("\nIf: ")
	f.If.printType(&pp, v)

	if f.Elif == nil {
		pp.Print("\nElif: nil")
	} else if len(f.Elif) > 0 {
		pp.Print("\nElif: [")

		ipp := indentPrinter{&pp}

		for n, e := range f.Elif {
			ipp.Printf("\n%d: ", n)
			e.printType(&ipp, v)
		}

		pp.Print("\n]")
	} else if v {
		pp.Print("\nElif: []")
	}

	if f.Else != nil {
		pp.Print("\nElse: ")
		f.Else.printType(&pp, v)
	} else if v {
		pp.Print("\nElse: nil")
	}

	pp.Print("\nTokens: ")
	f.Tokens.printType(&pp, v)

	io.WriteString(w, "\n}")
}

func (f *ImportStatement) printType(w io.Writer, v bool) {
	pp := indentPrinter{w}

	pp.Print("ImportStatement {")

	if f.RelativeModule != nil {
		pp.Print("\nRelativeModule: ")
		f.RelativeModule.printType(&pp, v)
	} else if v {
		pp.Print("\nRelativeModule: nil")
	}

	if f.Modules == nil {
		pp.Print("\nModules: nil")
	} else if len(f.Modules) > 0 {
		pp.Print("\nModules: [")

		ipp := indentPrinter{&pp}

		for n, e := range f.Modules {
			ipp.Printf("\n%d: ", n)
			e.printType(&ipp, v)
		}

		pp.Print("\n]")
	} else if v {
		pp.Print("\nModules: []")
	}

	pp.Print("\nTokens: ")
	f.Tokens.printType(&pp, v)

	io.WriteString(w, "\n}")
}

func (f *KeywordArgument) printType(w io.Writer, v bool) {
	pp := indentPrinter{w}

	pp.Print("KeywordArgument {")

	if f.KeywordItem != nil {
		pp.Print("\nKeywordItem: ")
		f.KeywordItem.printType(&pp, v)
	} else if v {
		pp.Print("\nKeywordItem: nil")
	}

	if f.Expression != nil {
		pp.Print("\nExpression: ")
		f.Expression.printType(&pp, v)
	} else if v {
		pp.Print("\nExpression: nil")
	}

	pp.Print("\nTokens: ")
	f.Tokens.printType(&pp, v)

	io.WriteString(w, "\n}")
}

func (f *KeywordItem) printType(w io.Writer, v bool) {
	pp := indentPrinter{w}

	pp.Print("KeywordItem {")

	if f.Identifier != nil {
		pp.Print("\nIdentifier: ")
		f.Identifier.printType(&pp, v)
	} else if v {
		pp.Print("\nIdentifier: nil")
	}

	pp.Print("\nExpression: ")
	f.Expression.printType(&pp, v)

	pp.Print("\nTokens: ")
	f.Tokens.printType(&pp, v)

	io.WriteString(w, "\n}")
}

func (f *LambdaExpression) printType(w io.Writer, v bool) {
	pp := indentPrinter{w}

	pp.Print("LambdaExpression {")

	if f.ParameterList != nil {
		pp.Print("\nParameterList: ")
		f.ParameterList.printType(&pp, v)
	} else if v {
		pp.Print("\nParameterList: nil")
	}

	pp.Print("\nExpression: ")
	f.Expression.printType(&pp, v)

	pp.Print("\nTokens: ")
	f.Tokens.printType(&pp, v)

	io.WriteString(w, "\n}")
}

func (f *ModuleAs) printType(w io.Writer, v bool) {
	pp := indentPrinter{w}

	pp.Print("ModuleAs {")

	pp.Print("\nModule: ")
	f.Module.printType(&pp, v)

	if f.As != nil {
		pp.Print("\nAs: ")
		f.As.printType(&pp, v)
	} else if v {
		pp.Print("\nAs: nil")
	}

	pp.Print("\nTokens: ")
	f.Tokens.printType(&pp, v)

	io.WriteString(w, "\n}")
}

func (f *Module) printType(w io.Writer, v bool) {
	pp := indentPrinter{w}

	pp.Print("Module {")

	if f.Identifiers == nil {
		pp.Print("\nIdentifiers: nil")
	} else if len(f.Identifiers) > 0 {
		pp.Print("\nIdentifiers: [")

		ipp := indentPrinter{&pp}

		for n, e := range f.Identifiers {
			ipp.Printf("\n%d: ", n)
			e.printType(&ipp, v)
		}

		pp.Print("\n]")
	} else if v {
		pp.Print("\nIdentifiers: []")
	}

	pp.Print("\nTokens: ")
	f.Tokens.printType(&pp, v)

	io.WriteString(w, "\n}")
}

func (f *MultiplyExpression) printType(w io.Writer, v bool) {
	pp := indentPrinter{w}

	pp.Print("MultiplyExpression {")

	pp.Print("\nUnaryExpression: ")
	f.UnaryExpression.printType(&pp, v)

	if f.Multiply != nil {
		pp.Print("\nMultiply: ")
		f.Multiply.printType(&pp, v)
	} else if v {
		pp.Print("\nMultiply: nil")
	}

	if f.MultiplyExpression != nil {
		pp.Print("\nMultiplyExpression: ")
		f.MultiplyExpression.printType(&pp, v)
	} else if v {
		pp.Print("\nMultiplyExpression: nil")
	}

	pp.Print("\nTokens: ")
	f.Tokens.printType(&pp, v)

	io.WriteString(w, "\n}")
}

func (f *NonLocalStatement) printType(w io.Writer, v bool) {
	pp := indentPrinter{w}

	pp.Print("NonLocalStatement {")

	if f.Identifiers == nil {
		pp.Print("\nIdentifiers: nil")
	} else if len(f.Identifiers) > 0 {
		pp.Print("\nIdentifiers: [")

		ipp := indentPrinter{&pp}

		for n, e := range f.Identifiers {
			ipp.Printf("\n%d: ", n)
			e.printType(&ipp, v)
		}

		pp.Print("\n]")
	} else if v {
		pp.Print("\nIdentifiers: []")
	}

	pp.Print("\nTokens: ")
	f.Tokens.printType(&pp, v)

	io.WriteString(w, "\n}")
}

func (f *NotTest) printType(w io.Writer, v bool) {
	pp := indentPrinter{w}

	pp.Print("NotTest {")

	if f.Nots != 0 || v {
		pp.Printf("\nNots: %v", f.Nots)
	}

	pp.Print("\nComparison: ")
	f.Comparison.printType(&pp, v)

	pp.Print("\nTokens: ")
	f.Tokens.printType(&pp, v)

	io.WriteString(w, "\n}")
}

func (f *OrExpression) printType(w io.Writer, v bool) {
	pp := indentPrinter{w}

	pp.Print("OrExpression {")

	pp.Print("\nXorExpression: ")
	f.XorExpression.printType(&pp, v)

	if f.OrExpression != nil {
		pp.Print("\nOrExpression: ")
		f.OrExpression.printType(&pp, v)
	} else if v {
		pp.Print("\nOrExpression: nil")
	}

	pp.Print("\nTokens: ")
	f.Tokens.printType(&pp, v)

	io.WriteString(w, "\n}")
}

func (f *OrTest) printType(w io.Writer, v bool) {
	pp := indentPrinter{w}

	pp.Print("OrTest {")

	pp.Print("\nAndTest: ")
	f.AndTest.printType(&pp, v)

	if f.OrTest != nil {
		pp.Print("\nOrTest: ")
		f.OrTest.printType(&pp, v)
	} else if v {
		pp.Print("\nOrTest: nil")
	}

	pp.Print("\nTokens: ")
	f.Tokens.printType(&pp, v)

	io.WriteString(w, "\n}")
}

func (f *Parameter) printType(w io.Writer, v bool) {
	pp := indentPrinter{w}

	pp.Print("Parameter {")

	if f.Identifier != nil {
		pp.Print("\nIdentifier: ")
		f.Identifier.printType(&pp, v)
	} else if v {
		pp.Print("\nIdentifier: nil")
	}

	if f.Type != nil {
		pp.Print("\nType: ")
		f.Type.printType(&pp, v)
	} else if v {
		pp.Print("\nType: nil")
	}

	pp.Print("\nTokens: ")
	f.Tokens.printType(&pp, v)

	io.WriteString(w, "\n}")
}

func (f *ParameterList) printType(w io.Writer, v bool) {
	pp := indentPrinter{w}

	pp.Print("ParameterList {")

	if f.DefParameters == nil {
		pp.Print("\nDefParameters: nil")
	} else if len(f.DefParameters) > 0 {
		pp.Print("\nDefParameters: [")

		ipp := indentPrinter{&pp}

		for n, e := range f.DefParameters {
			ipp.Printf("\n%d: ", n)
			e.printType(&ipp, v)
		}

		pp.Print("\n]")
	} else if v {
		pp.Print("\nDefParameters: []")
	}

	if f.NoPosOnly == nil {
		pp.Print("\nNoPosOnly: nil")
	} else if len(f.NoPosOnly) > 0 {
		pp.Print("\nNoPosOnly: [")

		ipp := indentPrinter{&pp}

		for n, e := range f.NoPosOnly {
			ipp.Printf("\n%d: ", n)
			e.printType(&ipp, v)
		}

		pp.Print("\n]")
	} else if v {
		pp.Print("\nNoPosOnly: []")
	}

	if f.StarArg != nil {
		pp.Print("\nStarArg: ")
		f.StarArg.printType(&pp, v)
	} else if v {
		pp.Print("\nStarArg: nil")
	}

	if f.StarArgs == nil {
		pp.Print("\nStarArgs: nil")
	} else if len(f.StarArgs) > 0 {
		pp.Print("\nStarArgs: [")

		ipp := indentPrinter{&pp}

		for n, e := range f.StarArgs {
			ipp.Printf("\n%d: ", n)
			e.printType(&ipp, v)
		}

		pp.Print("\n]")
	} else if v {
		pp.Print("\nStarArgs: []")
	}

	if f.StarStarArg != nil {
		pp.Print("\nStarStarArg: ")
		f.StarStarArg.printType(&pp, v)
	} else if v {
		pp.Print("\nStarStarArg: nil")
	}

	pp.Print("\nTokens: ")
	f.Tokens.printType(&pp, v)

	io.WriteString(w, "\n}")
}

func (f *PositionalArgument) printType(w io.Writer, v bool) {
	pp := indentPrinter{w}

	pp.Print("PositionalArgument {")

	if f.AssignmentExpression != nil {
		pp.Print("\nAssignmentExpression: ")
		f.AssignmentExpression.printType(&pp, v)
	} else if v {
		pp.Print("\nAssignmentExpression: nil")
	}

	if f.Expression != nil {
		pp.Print("\nExpression: ")
		f.Expression.printType(&pp, v)
	} else if v {
		pp.Print("\nExpression: nil")
	}

	pp.Print("\nTokens: ")
	f.Tokens.printType(&pp, v)

	io.WriteString(w, "\n}")
}

func (f *PowerExpression) printType(w io.Writer, v bool) {
	pp := indentPrinter{w}

	pp.Print("PowerExpression {")

	if f.AwaitExpression || v {
		pp.Printf("\nAwaitExpression: %v", f.AwaitExpression)
	}

	pp.Print("\nPrimaryExpression: ")
	f.PrimaryExpression.printType(&pp, v)

	if f.UnaryExpression != nil {
		pp.Print("\nUnaryExpression: ")
		f.UnaryExpression.printType(&pp, v)
	} else if v {
		pp.Print("\nUnaryExpression: nil")
	}

	pp.Print("\nTokens: ")
	f.Tokens.printType(&pp, v)

	io.WriteString(w, "\n}")
}

func (f *PrimaryExpression) printType(w io.Writer, v bool) {
	pp := indentPrinter{w}

	pp.Print("PrimaryExpression {")

	if f.PrimaryExpression != nil {
		pp.Print("\nPrimaryExpression: ")
		f.PrimaryExpression.printType(&pp, v)
	} else if v {
		pp.Print("\nPrimaryExpression: nil")
	}

	if f.Atom != nil {
		pp.Print("\nAtom: ")
		f.Atom.printType(&pp, v)
	} else if v {
		pp.Print("\nAtom: nil")
	}

	if f.AttributeRef != nil {
		pp.Print("\nAttributeRef: ")
		f.AttributeRef.printType(&pp, v)
	} else if v {
		pp.Print("\nAttributeRef: nil")
	}

	if f.Slicing != nil {
		pp.Print("\nSlicing: ")
		f.Slicing.printType(&pp, v)
	} else if v {
		pp.Print("\nSlicing: nil")
	}

	if f.Call != nil {
		pp.Print("\nCall: ")
		f.Call.printType(&pp, v)
	} else if v {
		pp.Print("\nCall: nil")
	}

	pp.Print("\nTokens: ")
	f.Tokens.printType(&pp, v)

	io.WriteString(w, "\n}")
}

func (f *RaiseStatement) printType(w io.Writer, v bool) {
	pp := indentPrinter{w}

	pp.Print("RaiseStatement {")

	if f.Expression != nil {
		pp.Print("\nExpression: ")
		f.Expression.printType(&pp, v)
	} else if v {
		pp.Print("\nExpression: nil")
	}

	if f.From != nil {
		pp.Print("\nFrom: ")
		f.From.printType(&pp, v)
	} else if v {
		pp.Print("\nFrom: nil")
	}

	pp.Print("\nTokens: ")
	f.Tokens.printType(&pp, v)

	io.WriteString(w, "\n}")
}

func (f *RelativeModule) printType(w io.Writer, v bool) {
	pp := indentPrinter{w}

	pp.Print("RelativeModule {")

	if f.Dots != 0 || v {
		pp.Printf("\nDots: %v", f.Dots)
	}

	if f.Module != nil {
		pp.Print("\nModule: ")
		f.Module.printType(&pp, v)
	} else if v {
		pp.Print("\nModule: nil")
	}

	pp.Print("\nTokens: ")
	f.Tokens.printType(&pp, v)

	io.WriteString(w, "\n}")
}

func (f *ReturnStatement) printType(w io.Writer, v bool) {
	pp := indentPrinter{w}

	pp.Print("ReturnStatement {")

	pp.Print("\nExpression: ")
	f.Expression.printType(&pp, v)

	if f.From != nil {
		pp.Print("\nFrom: ")
		f.From.printType(&pp, v)
	} else if v {
		pp.Print("\nFrom: nil")
	}

	pp.Print("\nTokens: ")
	f.Tokens.printType(&pp, v)

	io.WriteString(w, "\n}")
}

func (f *ShiftExpression) printType(w io.Writer, v bool) {
	pp := indentPrinter{w}

	pp.Print("ShiftExpression {")

	pp.Print("\nAddExpression: ")
	f.AddExpression.printType(&pp, v)

	if f.Shift != nil {
		pp.Print("\nShift: ")
		f.Shift.printType(&pp, v)
	} else if v {
		pp.Print("\nShift: nil")
	}

	if f.ShiftExpression != nil {
		pp.Print("\nShiftExpression: ")
		f.ShiftExpression.printType(&pp, v)
	} else if v {
		pp.Print("\nShiftExpression: nil")
	}

	pp.Print("\nTokens: ")
	f.Tokens.printType(&pp, v)

	io.WriteString(w, "\n}")
}

func (f *SimpleStatement) printType(w io.Writer, v bool) {
	pp := indentPrinter{w}

	pp.Print("SimpleStatement {")

	pp.Print("\nType: ")
	f.Type.printType(&pp, v)

	if f.AssertStatement != nil {
		pp.Print("\nAssertStatement: ")
		f.AssertStatement.printType(&pp, v)
	} else if v {
		pp.Print("\nAssertStatement: nil")
	}

	if f.ExpressionStatement != nil {
		pp.Print("\nExpressionStatement: ")
		f.ExpressionStatement.printType(&pp, v)
	} else if v {
		pp.Print("\nExpressionStatement: nil")
	}

	if f.AssignmentStatement != nil {
		pp.Print("\nAssignmentStatement: ")
		f.AssignmentStatement.printType(&pp, v)
	} else if v {
		pp.Print("\nAssignmentStatement: nil")
	}

	if f.AugmentedAssignmentStatement != nil {
		pp.Print("\nAugmentedAssignmentStatement: ")
		f.AugmentedAssignmentStatement.printType(&pp, v)
	} else if v {
		pp.Print("\nAugmentedAssignmentStatement: nil")
	}

	if f.AnnotatedAssignmentStatement != nil {
		pp.Print("\nAnnotatedAssignmentStatement: ")
		f.AnnotatedAssignmentStatement.printType(&pp, v)
	} else if v {
		pp.Print("\nAnnotatedAssignmentStatement: nil")
	}

	if f.DelStatement != nil {
		pp.Print("\nDelStatement: ")
		f.DelStatement.printType(&pp, v)
	} else if v {
		pp.Print("\nDelStatement: nil")
	}

	if f.ReturnStatement != nil {
		pp.Print("\nReturnStatement: ")
		f.ReturnStatement.printType(&pp, v)
	} else if v {
		pp.Print("\nReturnStatement: nil")
	}

	if f.YieldStatement != nil {
		pp.Print("\nYieldStatement: ")
		f.YieldStatement.printType(&pp, v)
	} else if v {
		pp.Print("\nYieldStatement: nil")
	}

	if f.RaiseStatement != nil {
		pp.Print("\nRaiseStatement: ")
		f.RaiseStatement.printType(&pp, v)
	} else if v {
		pp.Print("\nRaiseStatement: nil")
	}

	if f.ImportStatement != nil {
		pp.Print("\nImportStatement: ")
		f.ImportStatement.printType(&pp, v)
	} else if v {
		pp.Print("\nImportStatement: nil")
	}

	if f.GlobalStatement != nil {
		pp.Print("\nGlobalStatement: ")
		f.GlobalStatement.printType(&pp, v)
	} else if v {
		pp.Print("\nGlobalStatement: nil")
	}

	if f.NonLocalStatement != nil {
		pp.Print("\nNonLocalStatement: ")
		f.NonLocalStatement.printType(&pp, v)
	} else if v {
		pp.Print("\nNonLocalStatement: nil")
	}

	if f.TypeStatement != nil {
		pp.Print("\nTypeStatement: ")
		f.TypeStatement.printType(&pp, v)
	} else if v {
		pp.Print("\nTypeStatement: nil")
	}

	pp.Print("\nTokens: ")
	f.Tokens.printType(&pp, v)

	io.WriteString(w, "\n}")
}

func (f *SliceItem) printType(w io.Writer, v bool) {
	pp := indentPrinter{w}

	pp.Print("SliceItem {")

	if f.Expression != nil {
		pp.Print("\nExpression: ")
		f.Expression.printType(&pp, v)
	} else if v {
		pp.Print("\nExpression: nil")
	}

	if f.LowerBound != nil {
		pp.Print("\nLowerBound: ")
		f.LowerBound.printType(&pp, v)
	} else if v {
		pp.Print("\nLowerBound: nil")
	}

	if f.UpperBound != nil {
		pp.Print("\nUpperBound: ")
		f.UpperBound.printType(&pp, v)
	} else if v {
		pp.Print("\nUpperBound: nil")
	}

	if f.Stride != nil {
		pp.Print("\nStride: ")
		f.Stride.printType(&pp, v)
	} else if v {
		pp.Print("\nStride: nil")
	}

	pp.Print("\nTokens: ")
	f.Tokens.printType(&pp, v)

	io.WriteString(w, "\n}")
}

func (f *SliceList) printType(w io.Writer, v bool) {
	pp := indentPrinter{w}

	pp.Print("SliceList {")

	if f.SliceItems == nil {
		pp.Print("\nSliceItems: nil")
	} else if len(f.SliceItems) > 0 {
		pp.Print("\nSliceItems: [")

		ipp := indentPrinter{&pp}

		for n, e := range f.SliceItems {
			ipp.Printf("\n%d: ", n)
			e.printType(&ipp, v)
		}

		pp.Print("\n]")
	} else if v {
		pp.Print("\nSliceItems: []")
	}

	pp.Print("\nTokens: ")
	f.Tokens.printType(&pp, v)

	io.WriteString(w, "\n}")
}

func (f *StarredExpression) printType(w io.Writer, v bool) {
	pp := indentPrinter{w}

	pp.Print("StarredExpression {")

	if f.Expression != nil {
		pp.Print("\nExpression: ")
		f.Expression.printType(&pp, v)
	} else if v {
		pp.Print("\nExpression: nil")
	}

	if f.StarredItems == nil {
		pp.Print("\nStarredItems: nil")
	} else if len(f.StarredItems) > 0 {
		pp.Print("\nStarredItems: [")

		ipp := indentPrinter{&pp}

		for n, e := range f.StarredItems {
			ipp.Printf("\n%d: ", n)
			e.printType(&ipp, v)
		}

		pp.Print("\n]")
	} else if v {
		pp.Print("\nStarredItems: []")
	}

	pp.Print("\nTokens: ")
	f.Tokens.printType(&pp, v)

	io.WriteString(w, "\n}")
}

func (f *StarredExpressionList) printType(w io.Writer, v bool) {
	pp := indentPrinter{w}

	pp.Print("StarredExpressionList {")

	if f.StarredExpressions == nil {
		pp.Print("\nStarredExpressions: nil")
	} else if len(f.StarredExpressions) > 0 {
		pp.Print("\nStarredExpressions: [")

		ipp := indentPrinter{&pp}

		for n, e := range f.StarredExpressions {
			ipp.Printf("\n%d: ", n)
			e.printType(&ipp, v)
		}

		pp.Print("\n]")
	} else if v {
		pp.Print("\nStarredExpressions: []")
	}

	pp.Print("\nTokens: ")
	f.Tokens.printType(&pp, v)

	io.WriteString(w, "\n}")
}

func (f *StarredItem) printType(w io.Writer, v bool) {
	pp := indentPrinter{w}

	pp.Print("StarredItem {")

	if f.AssignmentExpression != nil {
		pp.Print("\nAssignmentExpression: ")
		f.AssignmentExpression.printType(&pp, v)
	} else if v {
		pp.Print("\nAssignmentExpression: nil")
	}

	if f.OrExpr != nil {
		pp.Print("\nOrExpr: ")
		f.OrExpr.printType(&pp, v)
	} else if v {
		pp.Print("\nOrExpr: nil")
	}

	pp.Print("\nTokens: ")
	f.Tokens.printType(&pp, v)

	io.WriteString(w, "\n}")
}

func (f *StarredList) printType(w io.Writer, v bool) {
	pp := indentPrinter{w}

	pp.Print("StarredList {")

	if f.StarredItems == nil {
		pp.Print("\nStarredItems: nil")
	} else if len(f.StarredItems) > 0 {
		pp.Print("\nStarredItems: [")

		ipp := indentPrinter{&pp}

		for n, e := range f.StarredItems {
			ipp.Printf("\n%d: ", n)
			e.printType(&ipp, v)
		}

		pp.Print("\n]")
	} else if v {
		pp.Print("\nStarredItems: []")
	}

	pp.Print("\nTokens: ")
	f.Tokens.printType(&pp, v)

	io.WriteString(w, "\n}")
}

func (f *StarredOrKeywordArgument) printType(w io.Writer, v bool) {
	pp := indentPrinter{w}

	pp.Print("StarredOrKeywordArgument {")

	if f.Expression != nil {
		pp.Print("\nExpression: ")
		f.Expression.printType(&pp, v)
	} else if v {
		pp.Print("\nExpression: nil")
	}

	if f.KeywordItem != nil {
		pp.Print("\nKeywordItem: ")
		f.KeywordItem.printType(&pp, v)
	} else if v {
		pp.Print("\nKeywordItem: nil")
	}

	pp.Print("\nTokens: ")
	f.Tokens.printType(&pp, v)

	io.WriteString(w, "\n}")
}

func (f *Statement) printType(w io.Writer, v bool) {
	pp := indentPrinter{w}

	pp.Print("Statement {")

	if f.StatementList != nil {
		pp.Print("\nStatementList: ")
		f.StatementList.printType(&pp, v)
	} else if v {
		pp.Print("\nStatementList: nil")
	}

	if f.CompoundStatement != nil {
		pp.Print("\nCompoundStatement: ")
		f.CompoundStatement.printType(&pp, v)
	} else if v {
		pp.Print("\nCompoundStatement: nil")
	}

	pp.Print("\nTokens: ")
	f.Tokens.printType(&pp, v)

	io.WriteString(w, "\n}")
}

func (f *StatementList) printType(w io.Writer, v bool) {
	pp := indentPrinter{w}

	pp.Print("StatementList {")

	if f.Statements == nil {
		pp.Print("\nStatements: nil")
	} else if len(f.Statements) > 0 {
		pp.Print("\nStatements: [")

		ipp := indentPrinter{&pp}

		for n, e := range f.Statements {
			ipp.Printf("\n%d: ", n)
			e.printType(&ipp, v)
		}

		pp.Print("\n]")
	} else if v {
		pp.Print("\nStatements: []")
	}

	pp.Print("\nTokens: ")
	f.Tokens.printType(&pp, v)

	io.WriteString(w, "\n}")
}

func (f *Suite) printType(w io.Writer, v bool) {
	pp := indentPrinter{w}

	pp.Print("Suite {")

	if f.StatementList != nil {
		pp.Print("\nStatementList: ")
		f.StatementList.printType(&pp, v)
	} else if v {
		pp.Print("\nStatementList: nil")
	}

	if f.Statements == nil {
		pp.Print("\nStatements: nil")
	} else if len(f.Statements) > 0 {
		pp.Print("\nStatements: [")

		ipp := indentPrinter{&pp}

		for n, e := range f.Statements {
			ipp.Printf("\n%d: ", n)
			e.printType(&ipp, v)
		}

		pp.Print("\n]")
	} else if v {
		pp.Print("\nStatements: []")
	}

	pp.Print("\nTokens: ")
	f.Tokens.printType(&pp, v)

	io.WriteString(w, "\n}")
}

func (f *Target) printType(w io.Writer, v bool) {
	pp := indentPrinter{w}

	pp.Print("Target {")

	if f.PrimaryExpression != nil {
		pp.Print("\nPrimaryExpression: ")
		f.PrimaryExpression.printType(&pp, v)
	} else if v {
		pp.Print("\nPrimaryExpression: nil")
	}

	if f.Tuple != nil {
		pp.Print("\nTuple: ")
		f.Tuple.printType(&pp, v)
	} else if v {
		pp.Print("\nTuple: nil")
	}

	if f.Array != nil {
		pp.Print("\nArray: ")
		f.Array.printType(&pp, v)
	} else if v {
		pp.Print("\nArray: nil")
	}

	if f.Star != nil {
		pp.Print("\nStar: ")
		f.Star.printType(&pp, v)
	} else if v {
		pp.Print("\nStar: nil")
	}

	pp.Print("\nTokens: ")
	f.Tokens.printType(&pp, v)

	io.WriteString(w, "\n}")
}

func (f *TargetList) printType(w io.Writer, v bool) {
	pp := indentPrinter{w}

	pp.Print("TargetList {")

	if f.Targets == nil {
		pp.Print("\nTargets: nil")
	} else if len(f.Targets) > 0 {
		pp.Print("\nTargets: [")

		ipp := indentPrinter{&pp}

		for n, e := range f.Targets {
			ipp.Printf("\n%d: ", n)
			e.printType(&ipp, v)
		}

		pp.Print("\n]")
	} else if v {
		pp.Print("\nTargets: []")
	}

	pp.Print("\nTokens: ")
	f.Tokens.printType(&pp, v)

	io.WriteString(w, "\n}")
}

func (f *TryStatement) printType(w io.Writer, v bool) {
	pp := indentPrinter{w}

	pp.Print("TryStatement {")

	pp.Print("\nTry: ")
	f.Try.printType(&pp, v)

	if f.Groups || v {
		pp.Printf("\nGroups: %v", f.Groups)
	}

	if f.Except == nil {
		pp.Print("\nExcept: nil")
	} else if len(f.Except) > 0 {
		pp.Print("\nExcept: [")

		ipp := indentPrinter{&pp}

		for n, e := range f.Except {
			ipp.Printf("\n%d: ", n)
			e.printType(&ipp, v)
		}

		pp.Print("\n]")
	} else if v {
		pp.Print("\nExcept: []")
	}

	if f.Else != nil {
		pp.Print("\nElse: ")
		f.Else.printType(&pp, v)
	} else if v {
		pp.Print("\nElse: nil")
	}

	if f.Finally != nil {
		pp.Print("\nFinally: ")
		f.Finally.printType(&pp, v)
	} else if v {
		pp.Print("\nFinally: nil")
	}

	pp.Print("\nTokens: ")
	f.Tokens.printType(&pp, v)

	io.WriteString(w, "\n}")
}

func (f *TypeParam) printType(w io.Writer, v bool) {
	pp := indentPrinter{w}

	pp.Print("TypeParam {")

	pp.Print("\nType: ")
	f.Type.printType(&pp, v)

	if f.Identifier != nil {
		pp.Print("\nIdentifier: ")
		f.Identifier.printType(&pp, v)
	} else if v {
		pp.Print("\nIdentifier: nil")
	}

	if f.Expression != nil {
		pp.Print("\nExpression: ")
		f.Expression.printType(&pp, v)
	} else if v {
		pp.Print("\nExpression: nil")
	}

	pp.Print("\nTokens: ")
	f.Tokens.printType(&pp, v)

	io.WriteString(w, "\n}")
}

func (f *TypeParams) printType(w io.Writer, v bool) {
	pp := indentPrinter{w}

	pp.Print("TypeParams {")

	if f.TypeParams == nil {
		pp.Print("\nTypeParams: nil")
	} else if len(f.TypeParams) > 0 {
		pp.Print("\nTypeParams: [")

		ipp := indentPrinter{&pp}

		for n, e := range f.TypeParams {
			ipp.Printf("\n%d: ", n)
			e.printType(&ipp, v)
		}

		pp.Print("\n]")
	} else if v {
		pp.Print("\nTypeParams: []")
	}

	pp.Print("\nTokens: ")
	f.Tokens.printType(&pp, v)

	io.WriteString(w, "\n}")
}

func (f *TypeStatement) printType(w io.Writer, v bool) {
	pp := indentPrinter{w}

	pp.Print("TypeStatement {")

	if f.Identifier != nil {
		pp.Print("\nIdentifier: ")
		f.Identifier.printType(&pp, v)
	} else if v {
		pp.Print("\nIdentifier: nil")
	}

	if f.TypeParams != nil {
		pp.Print("\nTypeParams: ")
		f.TypeParams.printType(&pp, v)
	} else if v {
		pp.Print("\nTypeParams: nil")
	}

	pp.Print("\nExpression: ")
	f.Expression.printType(&pp, v)

	pp.Print("\nTokens: ")
	f.Tokens.printType(&pp, v)

	io.WriteString(w, "\n}")
}

func (f *UnaryExpression) printType(w io.Writer, v bool) {
	pp := indentPrinter{w}

	pp.Print("UnaryExpression {")

	if f.PowerExpression != nil {
		pp.Print("\nPowerExpression: ")
		f.PowerExpression.printType(&pp, v)
	} else if v {
		pp.Print("\nPowerExpression: nil")
	}

	if f.Unary != nil {
		pp.Print("\nUnary: ")
		f.Unary.printType(&pp, v)
	} else if v {
		pp.Print("\nUnary: nil")
	}

	if f.UnaryExpression != nil {
		pp.Print("\nUnaryExpression: ")
		f.UnaryExpression.printType(&pp, v)
	} else if v {
		pp.Print("\nUnaryExpression: nil")
	}

	pp.Print("\nTokens: ")
	f.Tokens.printType(&pp, v)

	io.WriteString(w, "\n}")
}

func (f *WhileStatement) printType(w io.Writer, v bool) {
	pp := indentPrinter{w}

	pp.Print("WhileStatement {")

	pp.Print("\nWhile: ")
	f.While.printType(&pp, v)

	if f.Else != nil {
		pp.Print("\nElse: ")
		f.Else.printType(&pp, v)
	} else if v {
		pp.Print("\nElse: nil")
	}

	pp.Print("\nTokens: ")
	f.Tokens.printType(&pp, v)

	io.WriteString(w, "\n}")
}

func (f *WithItem) printType(w io.Writer, v bool) {
	pp := indentPrinter{w}

	pp.Print("WithItem {")

	pp.Print("\nExpression: ")
	f.Expression.printType(&pp, v)

	if f.Target != nil {
		pp.Print("\nTarget: ")
		f.Target.printType(&pp, v)
	} else if v {
		pp.Print("\nTarget: nil")
	}

	pp.Print("\nTokens: ")
	f.Tokens.printType(&pp, v)

	io.WriteString(w, "\n}")
}

func (f *WithStatement) printType(w io.Writer, v bool) {
	pp := indentPrinter{w}

	pp.Print("WithStatement {")

	if f.Async || v {
		pp.Printf("\nAsync: %v", f.Async)
	}

	pp.Print("\nContents: ")
	f.Contents.printType(&pp, v)

	pp.Print("\nSuite: ")
	f.Suite.printType(&pp, v)

	pp.Print("\nTokens: ")
	f.Tokens.printType(&pp, v)

	io.WriteString(w, "\n}")
}

func (f *WithStatementContents) printType(w io.Writer, v bool) {
	pp := indentPrinter{w}

	pp.Print("WithStatementContents {")

	if f.Items == nil {
		pp.Print("\nItems: nil")
	} else if len(f.Items) > 0 {
		pp.Print("\nItems: [")

		ipp := indentPrinter{&pp}

		for n, e := range f.Items {
			ipp.Printf("\n%d: ", n)
			e.printType(&ipp, v)
		}

		pp.Print("\n]")
	} else if v {
		pp.Print("\nItems: []")
	}

	pp.Print("\nTokens: ")
	f.Tokens.printType(&pp, v)

	io.WriteString(w, "\n}")
}

func (f *XorExpression) printType(w io.Writer, v bool) {
	pp := indentPrinter{w}

	pp.Print("XorExpression {")

	pp.Print("\nAndExpression: ")
	f.AndExpression.printType(&pp, v)

	if f.XorExpression != nil {
		pp.Print("\nXorExpression: ")
		f.XorExpression.printType(&pp, v)
	} else if v {
		pp.Print("\nXorExpression: nil")
	}

	pp.Print("\nTokens: ")
	f.Tokens.printType(&pp, v)

	io.WriteString(w, "\n}")
}

func (f *YieldExpression) printType(w io.Writer, v bool) {
	pp := indentPrinter{w}

	pp.Print("YieldExpression {")

	if f.ExpressionList != nil {
		pp.Print("\nExpressionList: ")
		f.ExpressionList.printType(&pp, v)
	} else if v {
		pp.Print("\nExpressionList: nil")
	}

	if f.StarredList != nil {
		pp.Print("\nStarredList: ")
		f.StarredList.printType(&pp, v)
	} else if v {
		pp.Print("\nStarredList: nil")
	}

	if f.From != nil {
		pp.Print("\nFrom: ")
		f.From.printType(&pp, v)
	} else if v {
		pp.Print("\nFrom: nil")
	}

	pp.Print("\nTokens: ")
	f.Tokens.printType(&pp, v)

	io.WriteString(w, "\n}")
}

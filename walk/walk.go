// Package walk provides a python type walker.
package walk

import "vimagination.zapto.org/python"

// Handler is used to process python types.
type Handler interface {
	Handle(python.Type) error
}

// HandlerFunc wraps a func to implement Handler interface.
type HandlerFunc func(python.Type) error

// Handle implements the Handler interface.
func (h HandlerFunc) Handle(t python.Type) error {
	return h(t)
}

// Walk calls the Handle function on the given interface for each non-nil, non-Token field of the given R type.
func Walk(t python.Type, fn Handler) error {
	switch t := t.(type) {
	case python.AddExpression:
		return walkAddExpression(&t, fn)
	case *python.AddExpression:
		return walkAddExpression(t, fn)
	case python.AndExpression:
		return walkAndExpression(&t, fn)
	case *python.AndExpression:
		return walkAndExpression(t, fn)
	case python.AndTest:
		return walkAndTest(&t, fn)
	case *python.AndTest:
		return walkAndTest(t, fn)
	case python.AnnotatedAssignmentStatement:
		return walkAnnotatedAssignmentStatement(&t, fn)
	case *python.AnnotatedAssignmentStatement:
		return walkAnnotatedAssignmentStatement(t, fn)
	case python.ArgumentList:
		return walkArgumentList(&t, fn)
	case *python.ArgumentList:
		return walkArgumentList(t, fn)
	case python.ArgumentListOrComprehension:
		return walkArgumentListOrComprehension(&t, fn)
	case *python.ArgumentListOrComprehension:
		return walkArgumentListOrComprehension(t, fn)
	case python.AssertStatement:
		return walkAssertStatement(&t, fn)
	case *python.AssertStatement:
		return walkAssertStatement(t, fn)
	case python.AssignmentExpressionAndSuite:
		return walkAssignmentExpressionAndSuite(&t, fn)
	case *python.AssignmentExpressionAndSuite:
		return walkAssignmentExpressionAndSuite(t, fn)
	case python.AssignmentExpression:
		return walkAssignmentExpression(&t, fn)
	case *python.AssignmentExpression:
		return walkAssignmentExpression(t, fn)
	case python.AssignmentStatement:
		return walkAssignmentStatement(&t, fn)
	case *python.AssignmentStatement:
		return walkAssignmentStatement(t, fn)
	case python.Atom:
		return walkAtom(&t, fn)
	case *python.Atom:
		return walkAtom(t, fn)
	case python.AugmentedAssignmentStatement:
		return walkAugmentedAssignmentStatement(&t, fn)
	case *python.AugmentedAssignmentStatement:
		return walkAugmentedAssignmentStatement(t, fn)
	case python.AugTarget:
		return walkAugTarget(&t, fn)
	case *python.AugTarget:
		return walkAugTarget(t, fn)
	case python.ClassDefinition:
		return walkClassDefinition(&t, fn)
	case *python.ClassDefinition:
		return walkClassDefinition(t, fn)
	case python.Comparison:
		return walkComparison(&t, fn)
	case *python.Comparison:
		return walkComparison(t, fn)
	case python.ComparisonExpression:
		return walkComparisonExpression(&t, fn)
	case *python.ComparisonExpression:
		return walkComparisonExpression(t, fn)
	case python.CompoundStatement:
		return walkCompoundStatement(&t, fn)
	case *python.CompoundStatement:
		return walkCompoundStatement(t, fn)
	case python.Comprehension:
	case *python.Comprehension:
	case python.ComprehensionFor:
	case *python.ComprehensionFor:
	case python.ComprehensionIf:
	case *python.ComprehensionIf:
	case python.ComprehensionIterator:
	case *python.ComprehensionIterator:
	case python.ConditionalExpression:
	case *python.ConditionalExpression:
	case python.Decorators:
	case *python.Decorators:
	case python.DefParameter:
	case *python.DefParameter:
	case python.DelStatement:
	case *python.DelStatement:
	case python.DictDisplay:
	case *python.DictDisplay:
	case python.DictItem:
	case *python.DictItem:
	case python.Enclosure:
	case *python.Enclosure:
	case python.Except:
	case *python.Except:
	case python.Expression:
	case *python.Expression:
	case python.ExpressionList:
	case *python.ExpressionList:
	case python.File:
	case *python.File:
	case python.FlexibleExpression:
	case *python.FlexibleExpression:
	case python.FlexibleExpressionList:
	case *python.FlexibleExpressionList:
	case python.FlexibleExpressionListOrComprehension:
	case *python.FlexibleExpressionListOrComprehension:
	case python.ForStatement:
	case *python.ForStatement:
	case python.FuncDefinition:
	case *python.FuncDefinition:
	case python.GeneratorExpression:
	case *python.GeneratorExpression:
	case python.GlobalStatement:
	case *python.GlobalStatement:
	case python.IfStatement:
	case *python.IfStatement:
	case python.ImportStatement:
	case *python.ImportStatement:
	case python.KeywordArgument:
	case *python.KeywordArgument:
	case python.KeywordItem:
	case *python.KeywordItem:
	case python.LambdaExpression:
	case *python.LambdaExpression:
	case python.ModuleAs:
	case *python.ModuleAs:
	case python.Module:
	case *python.Module:
	case python.MultiplyExpression:
	case *python.MultiplyExpression:
	case python.NonLocalStatement:
	case *python.NonLocalStatement:
	case python.NotTest:
	case *python.NotTest:
	case python.OrExpression:
	case *python.OrExpression:
	case python.OrTest:
	case *python.OrTest:
	case python.Parameter:
	case *python.Parameter:
	case python.ParameterList:
	case *python.ParameterList:
	case python.PositionalArgument:
	case *python.PositionalArgument:
	case python.PowerExpression:
	case *python.PowerExpression:
	case python.PrimaryExpression:
	case *python.PrimaryExpression:
	case python.RaiseStatement:
	case *python.RaiseStatement:
	case python.RelativeModule:
	case *python.RelativeModule:
	case python.ReturnStatement:
	case *python.ReturnStatement:
	case python.ShiftExpression:
	case *python.ShiftExpression:
	case python.SimpleStatement:
	case *python.SimpleStatement:
	case python.SliceItem:
	case *python.SliceItem:
	case python.SliceList:
	case *python.SliceList:
	case python.StarredExpression:
	case *python.StarredExpression:
	case python.StarredItem:
	case *python.StarredItem:
	case python.StarredList:
	case *python.StarredList:
	case python.StarredOrKeyword:
	case *python.StarredOrKeyword:
	case python.Statement:
	case *python.Statement:
	case python.StatementList:
	case *python.StatementList:
	case python.Suite:
	case *python.Suite:
	case python.Target:
	case *python.Target:
	case python.TargetList:
	case *python.TargetList:
	case python.TryStatement:
	case *python.TryStatement:
	case python.TypeParam:
	case *python.TypeParam:
	case python.TypeParams:
	case *python.TypeParams:
	case python.TypeStatement:
	case *python.TypeStatement:
	case python.UnaryExpression:
	case *python.UnaryExpression:
	case python.WhileStatement:
	case *python.WhileStatement:
	case python.WithItem:
	case *python.WithItem:
	case python.WithStatement:
	case *python.WithStatement:
	case python.WithStatementContents:
	case *python.WithStatementContents:
	case python.XorExpression:
	case *python.XorExpression:
	case python.YieldExpression:
	case *python.YieldExpression:
	}

	return nil
}

func walkAddExpression(t *python.AddExpression, fn Handler) error {
	if err := fn.Handle(&t.MultiplyExpression); err != nil {
		return err
	}

	if t.AddExpression != nil {
		return fn.Handle(t.AddExpression)
	}

	return nil
}

func walkAndExpression(t *python.AndExpression, fn Handler) error {
	if err := fn.Handle(&t.ShiftExpression); err != nil {
		return err
	}

	if t.AndExpression != nil {
		return fn.Handle(t.AndExpression)
	}

	return nil
}

func walkAndTest(t *python.AndTest, fn Handler) error {
	if err := fn.Handle(&t.NotTest); err != nil {
		return err
	}

	if t.AndTest != nil {
		return fn.Handle(t.AndTest)
	}

	return nil
}

func walkAnnotatedAssignmentStatement(t *python.AnnotatedAssignmentStatement, fn Handler) error {
	if err := fn.Handle(&t.AugTarget); err != nil {
		return err
	}

	if err := fn.Handle(&t.Expression); err != nil {
		return err
	}

	if t.StarredExpression != nil {
		return fn.Handle(t.StarredExpression)
	} else if t.YieldExpression != nil {
		return fn.Handle(t.YieldExpression)
	}

	return nil
}

func walkArgumentList(t *python.ArgumentList, fn Handler) error {
	for n := range t.PositionalArguments {
		if err := fn.Handle(&t.PositionalArguments[n]); err != nil {
			return err
		}
	}

	for n := range t.StarredAndKeywordArguments {
		if err := fn.Handle(&t.StarredAndKeywordArguments[n]); err != nil {
			return err
		}
	}

	for n := range t.KeywordArguments {
		if err := fn.Handle(&t.KeywordArguments[n]); err != nil {
			return err
		}
	}

	return nil
}

func walkArgumentListOrComprehension(t *python.ArgumentListOrComprehension, fn Handler) error {
	if t.ArgumentList != nil {
		return fn.Handle(t.ArgumentList)
	} else if t.Comprehension != nil {
		return fn.Handle(t.Comprehension)
	}

	return nil
}

func walkAssertStatement(t *python.AssertStatement, fn Handler) error {
	for n := range t.Expressions {
		if err := fn.Handle(&t.Expressions[n]); err != nil {
			return err
		}
	}

	return nil
}

func walkAssignmentExpressionAndSuite(t *python.AssignmentExpressionAndSuite, fn Handler) error {
	if err := fn.Handle(&t.AssignmentExpression); err != nil {
		return err
	}

	return fn.Handle(&t.Suite)
}

func walkAssignmentExpression(t *python.AssignmentExpression, fn Handler) error {
	return fn.Handle(&t.Expression)
}

func walkAssignmentStatement(t *python.AssignmentStatement, fn Handler) error {
	for n := range t.TargetLists {
		if err := fn.Handle(&t.TargetLists[n]); err != nil {
			return err
		}
	}

	if t.StarredExpression != nil {
		return fn.Handle(t.StarredExpression)
	} else if t.YieldExpression != nil {
		return fn.Handle(t.YieldExpression)
	}

	return nil
}

func walkAtom(t *python.Atom, fn Handler) error {
	if t.Enclosure != nil {
		return fn.Handle(t.Enclosure)
	}

	return nil
}

func walkAugmentedAssignmentStatement(t *python.AugmentedAssignmentStatement, fn Handler) error {
	if err := fn.Handle(&t.AugTarget); err != nil {
		return err
	}

	if t.ExpressionList != nil {
		return fn.Handle(t.ExpressionList)
	} else if t.YieldExpression != nil {
		return fn.Handle(t.YieldExpression)
	}

	return nil
}

func walkAugTarget(t *python.AugTarget, fn Handler) error {
	return fn.Handle(&t.PrimaryExpression)
}

func walkClassDefinition(t *python.ClassDefinition, fn Handler) error {
	if t.Decorators != nil {
		if err := fn.Handle(t.Decorators); err != nil {
			return err
		}
	}

	if t.TypeParams != nil {
		if err := fn.Handle(t.TypeParams); err != nil {
			return err
		}
	}

	if t.Inheritance != nil {
		if err := fn.Handle(t.Inheritance); err != nil {
			return err
		}
	}

	return fn.Handle(&t.Suite)
}

func walkComparison(t *python.Comparison, fn Handler) error {
	if err := fn.Handle(&t.OrExpression); err != nil {
		return err
	}

	for n := range t.Comparisons {
		if err := fn.Handle(&t.Comparisons[n]); err != nil {
			return err
		}
	}

	return nil
}

func walkComparisonExpression(t *python.ComparisonExpression, fn Handler) error {
	return fn.Handle(&t.OrExpression)
}

func walkCompoundStatement(t *python.CompoundStatement, fn Handler) error {
	if t.If != nil {
		return fn.Handle(t.If)
	} else if t.While != nil {
		return fn.Handle(t.While)
	} else if t.For != nil {
		return fn.Handle(t.For)
	} else if t.Try != nil {
		return fn.Handle(t.Try)
	} else if t.With != nil {
		return fn.Handle(t.With)
	} else if t.Func != nil {
		return fn.Handle(t.Func)
	} else if t.Class != nil {
		return fn.Handle(t.Class)
	}

	return nil
}

func walkComprehension(t *python.Comprehension, fn Handler) error { return nil }

func walkComprehensionFor(t *python.ComprehensionFor, fn Handler) error { return nil }

func walkComprehensionIf(t *python.ComprehensionIf, fn Handler) error { return nil }

func walkComprehensionIterator(t *python.ComprehensionIterator, fn Handler) error { return nil }

func walkConditionalExpression(t *python.ConditionalExpression, fn Handler) error { return nil }

func walkDecorators(t *python.Decorators, fn Handler) error { return nil }

func walkDefParameter(t *python.DefParameter, fn Handler) error { return nil }

func walkDelStatement(t *python.DelStatement, fn Handler) error { return nil }

func walkDictDisplay(t *python.DictDisplay, fn Handler) error { return nil }

func walkDictItem(t *python.DictItem, fn Handler) error { return nil }

func walkEnclosure(t *python.Enclosure, fn Handler) error { return nil }

func walkExcept(t *python.Except, fn Handler) error { return nil }

func walkExpression(t *python.Expression, fn Handler) error { return nil }

func walkExpressionList(t *python.ExpressionList, fn Handler) error { return nil }

func walkFile(t *python.File, fn Handler) error { return nil }

func walkFlexibleExpression(t *python.FlexibleExpression, fn Handler) error { return nil }

func walkFlexibleExpressionList(t *python.FlexibleExpressionList, fn Handler) error { return nil }

func walkFlexibleExpressionListOrComprehension(t *python.FlexibleExpressionListOrComprehension, fn Handler) error {
	return nil
}

func walkForStatement(t *python.ForStatement, fn Handler) error { return nil }

func walkFuncDefinition(t *python.FuncDefinition, fn Handler) error { return nil }

func walkGeneratorExpression(t *python.GeneratorExpression, fn Handler) error { return nil }

func walkGlobalStatement(t *python.GlobalStatement, fn Handler) error { return nil }

func walkIfStatement(t *python.IfStatement, fn Handler) error { return nil }

func walkImportStatement(t *python.ImportStatement, fn Handler) error { return nil }

func walkKeywordArgument(t *python.KeywordArgument, fn Handler) error { return nil }

func walkKeywordItem(t *python.KeywordItem, fn Handler) error { return nil }

func walkLambdaExpression(t *python.LambdaExpression, fn Handler) error { return nil }

func walkModuleAs(t *python.ModuleAs, fn Handler) error { return nil }

func walkModule(t *python.Module, fn Handler) error { return nil }

func walkMultiplyExpression(t *python.MultiplyExpression, fn Handler) error { return nil }

func walkNonLocalStatement(t *python.NonLocalStatement, fn Handler) error { return nil }

func walkNotTest(t *python.NotTest, fn Handler) error { return nil }

func walkOrExpression(t *python.OrExpression, fn Handler) error { return nil }

func walkOrTest(t *python.OrTest, fn Handler) error { return nil }

func walkParameter(t *python.Parameter, fn Handler) error { return nil }

func walkParameterList(t *python.ParameterList, fn Handler) error { return nil }

func walkPositionalArgument(t *python.PositionalArgument, fn Handler) error { return nil }

func walkPowerExpression(t *python.PowerExpression, fn Handler) error { return nil }

func walkPrimaryExpression(t *python.PrimaryExpression, fn Handler) error { return nil }

func walkRaiseStatement(t *python.RaiseStatement, fn Handler) error { return nil }

func walkRelativeModule(t *python.RelativeModule, fn Handler) error { return nil }

func walkReturnStatement(t *python.ReturnStatement, fn Handler) error { return nil }

func walkShiftExpression(t *python.ShiftExpression, fn Handler) error { return nil }

func walkSimpleStatement(t *python.SimpleStatement, fn Handler) error { return nil }

func walkSliceItem(t *python.SliceItem, fn Handler) error { return nil }

func walkSliceList(t *python.SliceList, fn Handler) error { return nil }

func walkStarredExpression(t *python.StarredExpression, fn Handler) error { return nil }

func walkStarredItem(t *python.StarredItem, fn Handler) error { return nil }

func walkStarredList(t *python.StarredList, fn Handler) error { return nil }

func walkStarredOrKeyword(t *python.StarredOrKeyword, fn Handler) error { return nil }

func walkStatement(t *python.Statement, fn Handler) error { return nil }

func walkStatementList(t *python.StatementList, fn Handler) error { return nil }

func walkSuite(t *python.Suite, fn Handler) error { return nil }

func walkTarget(t *python.Target, fn Handler) error { return nil }

func walkTargetList(t *python.TargetList, fn Handler) error { return nil }

func walkTryStatement(t *python.TryStatement, fn Handler) error { return nil }

func walkTypeParam(t *python.TypeParam, fn Handler) error { return nil }

func walkTypeParams(t *python.TypeParams, fn Handler) error { return nil }

func walkTypeStatement(t *python.TypeStatement, fn Handler) error { return nil }

func walkUnaryExpression(t *python.UnaryExpression, fn Handler) error { return nil }

func walkWhileStatement(t *python.WhileStatement, fn Handler) error { return nil }

func walkWithItem(t *python.WithItem, fn Handler) error { return nil }

func walkWithStatement(t *python.WithStatement, fn Handler) error { return nil }

func walkWithStatementContents(t *python.WithStatementContents, fn Handler) error { return nil }

func walkXorExpression(t *python.XorExpression, fn Handler) error { return nil }

func walkYieldExpression(t *python.YieldExpression, fn Handler) error { return nil }

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
		return walkComprehension(&t, fn)
	case *python.Comprehension:
		return walkComprehension(t, fn)
	case python.ComprehensionFor:
		return walkComprehensionFor(&t, fn)
	case *python.ComprehensionFor:
		return walkComprehensionFor(t, fn)
	case python.ComprehensionIf:
		return walkComprehensionIf(&t, fn)
	case *python.ComprehensionIf:
		return walkComprehensionIf(t, fn)
	case python.ComprehensionIterator:
		return walkComprehensionIterator(&t, fn)
	case *python.ComprehensionIterator:
		return walkComprehensionIterator(t, fn)
	case python.ConditionalExpression:
		return walkConditionalExpression(&t, fn)
	case *python.ConditionalExpression:
		return walkConditionalExpression(t, fn)
	case python.Decorators:
		return walkDecorators(&t, fn)
	case *python.Decorators:
		return walkDecorators(t, fn)
	case python.DefParameter:
		return walkDefParameter(&t, fn)
	case *python.DefParameter:
		return walkDefParameter(t, fn)
	case python.DelStatement:
		return walkDelStatement(&t, fn)
	case *python.DelStatement:
		return walkDelStatement(t, fn)
	case python.DictDisplay:
		return walkDictDisplay(&t, fn)
	case *python.DictDisplay:
		return walkDictDisplay(t, fn)
	case python.DictItem:
		return walkDictItem(&t, fn)
	case *python.DictItem:
		return walkDictItem(t, fn)
	case python.Enclosure:
		return walkEnclosure(&t, fn)
	case *python.Enclosure:
		return walkEnclosure(t, fn)
	case python.Except:
		return walkExcept(&t, fn)
	case *python.Except:
		return walkExcept(t, fn)
	case python.Expression:
		return walkExpression(&t, fn)
	case *python.Expression:
		return walkExpression(t, fn)
	case python.ExpressionList:
		return walkExpressionList(&t, fn)
	case *python.ExpressionList:
		return walkExpressionList(t, fn)
	case python.File:
		return walkFile(&t, fn)
	case *python.File:
		return walkFile(t, fn)
	case python.FlexibleExpression:
		return walkFlexibleExpression(&t, fn)
	case *python.FlexibleExpression:
		return walkFlexibleExpression(t, fn)
	case python.FlexibleExpressionList:
		return walkFlexibleExpressionList(&t, fn)
	case *python.FlexibleExpressionList:
		return walkFlexibleExpressionList(t, fn)
	case python.FlexibleExpressionListOrComprehension:
		return walkFlexibleExpressionListOrComprehension(&t, fn)
	case *python.FlexibleExpressionListOrComprehension:
		return walkFlexibleExpressionListOrComprehension(t, fn)
	case python.ForStatement:
		return walkForStatement(&t, fn)
	case *python.ForStatement:
		return walkForStatement(t, fn)
	case python.FuncDefinition:
		return walkFuncDefinition(&t, fn)
	case *python.FuncDefinition:
		return walkFuncDefinition(t, fn)
	case python.GeneratorExpression:
		return walkGeneratorExpression(&t, fn)
	case *python.GeneratorExpression:
		return walkGeneratorExpression(t, fn)
	case python.GlobalStatement:
		return walkGlobalStatement(&t, fn)
	case *python.GlobalStatement:
		return walkGlobalStatement(t, fn)
	case python.IfStatement:
		return walkIfStatement(&t, fn)
	case *python.IfStatement:
		return walkIfStatement(t, fn)
	case python.ImportStatement:
		return walkImportStatement(&t, fn)
	case *python.ImportStatement:
		return walkImportStatement(t, fn)
	case python.KeywordArgument:
		return walkKeywordArgument(&t, fn)
	case *python.KeywordArgument:
		return walkKeywordArgument(t, fn)
	case python.KeywordItem:
		return walkKeywordItem(&t, fn)
	case *python.KeywordItem:
		return walkKeywordItem(t, fn)
	case python.LambdaExpression:
		return walkLambdaExpression(&t, fn)
	case *python.LambdaExpression:
		return walkLambdaExpression(t, fn)
	case python.ModuleAs:
		return walkModuleAs(&t, fn)
	case *python.ModuleAs:
		return walkModuleAs(t, fn)
	case python.Module:
		return walkModule(&t, fn)
	case *python.Module:
		return walkModule(t, fn)
	case python.MultiplyExpression:
		return walkMultiplyExpression(&t, fn)
	case *python.MultiplyExpression:
		return walkMultiplyExpression(t, fn)
	case python.NonLocalStatement:
		return walkNonLocalStatement(&t, fn)
	case *python.NonLocalStatement:
		return walkNonLocalStatement(t, fn)
	case python.NotTest:
		return walkNotTest(&t, fn)
	case *python.NotTest:
		return walkNotTest(t, fn)
	case python.OrExpression:
		return walkOrExpression(&t, fn)
	case *python.OrExpression:
		return walkOrExpression(t, fn)
	case python.OrTest:
		return walkOrTest(&t, fn)
	case *python.OrTest:
		return walkOrTest(t, fn)
	case python.Parameter:
		return walkParameter(&t, fn)
	case *python.Parameter:
		return walkParameter(t, fn)
	case python.ParameterList:
		return walkParameterList(&t, fn)
	case *python.ParameterList:
		return walkParameterList(t, fn)
	case python.PositionalArgument:
		return walkPositionalArgument(&t, fn)
	case *python.PositionalArgument:
		return walkPositionalArgument(t, fn)
	case python.PowerExpression:
		return walkPowerExpression(&t, fn)
	case *python.PowerExpression:
		return walkPowerExpression(t, fn)
	case python.PrimaryExpression:
		return walkPrimaryExpression(&t, fn)
	case *python.PrimaryExpression:
		return walkPrimaryExpression(t, fn)
	case python.RaiseStatement:
		return walkRaiseStatement(&t, fn)
	case *python.RaiseStatement:
		return walkRaiseStatement(t, fn)
	case python.RelativeModule:
		return walkRelativeModule(&t, fn)
	case *python.RelativeModule:
		return walkRelativeModule(t, fn)
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

func walkComprehension(t *python.Comprehension, fn Handler) error {
	if err := fn.Handle(&t.AssignmentExpression); err != nil {
		return err
	}

	return fn.Handle(&t.ComprehensionFor)
}

func walkComprehensionFor(t *python.ComprehensionFor, fn Handler) error {
	if err := fn.Handle(&t.TargetList); err != nil {
		return err
	}

	if err := fn.Handle(&t.OrTest); err != nil {
		return err
	}

	if t.ComprehensionIterator != nil {
		return fn.Handle(t.ComprehensionIterator)
	}

	return nil
}

func walkComprehensionIf(t *python.ComprehensionIf, fn Handler) error {
	if err := fn.Handle(&t.OrTest); err != nil {
		return err
	}

	if t.ComprehensionIterator != nil {
		return fn.Handle(t.ComprehensionIterator)
	}

	return nil
}

func walkComprehensionIterator(t *python.ComprehensionIterator, fn Handler) error {
	if t.ComprehensionFor != nil {
		return fn.Handle(t.ComprehensionFor)
	}

	if t.ComprehensionIf != nil {
		return fn.Handle(t.ComprehensionIf)
	}

	return nil
}

func walkConditionalExpression(t *python.ConditionalExpression, fn Handler) error {
	if err := fn.Handle(&t.OrTest); err != nil {
		return err
	}

	if t.If != nil {
		if err := fn.Handle(t.If); err != nil {
			return err
		}

		if t.Else != nil {
			return fn.Handle(t.Else)
		}
	}

	return nil
}

func walkDecorators(t *python.Decorators, fn Handler) error {
	for n := range t.Decorators {
		if err := fn.Handle(&t.Decorators[n]); err != nil {
			return err
		}
	}

	return nil
}

func walkDefParameter(t *python.DefParameter, fn Handler) error {
	if err := fn.Handle(&t.Parameter); err != nil {
		return err
	}

	if t.Value != nil {
		return fn.Handle(t.Value)
	}

	return nil
}

func walkDelStatement(t *python.DelStatement, fn Handler) error {
	return fn.Handle(&t.TargetList)
}

func walkDictDisplay(t *python.DictDisplay, fn Handler) error {
	for n := range t.DictItems {
		if err := fn.Handle(&t.DictItems[n]); err != nil {
			return err
		}
	}

	if t.DictComprehension != nil {
		return fn.Handle(t.DictComprehension)
	}

	return nil
}

func walkDictItem(t *python.DictItem, fn Handler) error {
	if t.OrExpression != nil {
		if err := fn.Handle(t.OrExpression); err != nil {
			return err
		}
	} else if t.Key != nil && t.Value != nil {
		if err := fn.Handle(t.Key); err != nil {
			return err
		}

		return fn.Handle(t.Value)
	}

	return nil
}

func walkEnclosure(t *python.Enclosure, fn Handler) error {
	if t.ParenthForm != nil {
		return fn.Handle(t.ParenthForm)
	} else if t.ListDisplay != nil {
		return fn.Handle(t.ListDisplay)
	} else if t.DictDisplay != nil {
		return fn.Handle(t.DictDisplay)
	} else if t.SetDisplay != nil {
		return fn.Handle(t.SetDisplay)
	} else if t.GeneratorExpression != nil {
		return fn.Handle(t.GeneratorExpression)
	} else if t.YieldAtom != nil {
		return fn.Handle(t.YieldAtom)
	}

	return nil
}

func walkExcept(t *python.Except, fn Handler) error {
	if err := fn.Handle(&t.Expression); err != nil {
		return err
	}

	return fn.Handle(&t.Suite)
}

func walkExpression(t *python.Expression, fn Handler) error {
	if t.ConditionalExpression != nil {
		return fn.Handle(t.ConditionalExpression)
	} else if t.LambdaExpression != nil {
		return fn.Handle(t.LambdaExpression)
	}

	return nil
}

func walkExpressionList(t *python.ExpressionList, fn Handler) error {
	for n := range t.Expressions {
		if err := fn.Handle(&t.Expressions[n]); err != nil {
			return err
		}
	}

	return nil
}

func walkFile(t *python.File, fn Handler) error {
	for n := range t.Statements {
		if err := fn.Handle(&t.Statements[n]); err != nil {
			return err
		}
	}

	return nil
}

func walkFlexibleExpression(t *python.FlexibleExpression, fn Handler) error {
	if t.AssignmentExpression != nil {
		return fn.Handle(t.AssignmentExpression)
	} else if t.StarredExpression != nil {
		return fn.Handle(t.StarredExpression)
	}

	return nil
}

func walkFlexibleExpressionList(t *python.FlexibleExpressionList, fn Handler) error {
	for n := range t.FlexibleExpressions {
		if err := fn.Handle(&t.FlexibleExpressions[n]); err != nil {
			return err
		}
	}

	return nil
}

func walkFlexibleExpressionListOrComprehension(t *python.FlexibleExpressionListOrComprehension, fn Handler) error {
	if t.FlexibleExpressionList != nil {
		return fn.Handle(t.FlexibleExpressionList)
	} else if t.Comprehension != nil {
		return fn.Handle(t.Comprehension)
	}

	return nil
}

func walkForStatement(t *python.ForStatement, fn Handler) error {
	if err := fn.Handle(&t.TargetList); err != nil {
		return err
	}

	if err := fn.Handle(&t.StarredList); err != nil {
		return err
	}

	if err := fn.Handle(&t.Suite); err != nil {
		return err
	}

	if t.Else != nil {
		return fn.Handle(t.Else)
	}

	return nil
}

func walkFuncDefinition(t *python.FuncDefinition, fn Handler) error {
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

	if err := fn.Handle(&t.ParameterList); err != nil {
		return err
	}

	if t.Expression != nil {
		if err := fn.Handle(t.Expression); err != nil {
			return err
		}
	}

	return fn.Handle(&t.Suite)
}

func walkGeneratorExpression(t *python.GeneratorExpression, fn Handler) error {
	if err := fn.Handle(&t.Expression); err != nil {
	}

	return fn.Handle(&t.ComprehensionFor)
}

func walkGlobalStatement(t *python.GlobalStatement, fn Handler) error {
	return nil
}

func walkIfStatement(t *python.IfStatement, fn Handler) error {
	if err := fn.Handle(&t.AssignmentExpression); err != nil {
		return err
	}

	if err := fn.Handle(&t.Suite); err != nil {
		return err
	}

	for n := range t.Elif {
		if err := fn.Handle(&t.Elif[n]); err != nil {
			return err
		}
	}

	if t.Else != nil {
		return fn.Handle(t.Else)
	}

	return nil
}

func walkImportStatement(t *python.ImportStatement, fn Handler) error {
	if t.RelativeModule != nil {
		if err := fn.Handle(t.RelativeModule); err != nil {
			return err
		}
	}

	for n := range t.Modules {
		if err := fn.Handle(&t.Modules[n]); err != nil {
			return err
		}
	}

	return nil
}

func walkKeywordArgument(t *python.KeywordArgument, fn Handler) error {
	if t.KeywordItem != nil {
		return fn.Handle(t.KeywordItem)
	} else if t.Expression != nil {
		return fn.Handle(t.Expression)
	}

	return nil
}

func walkKeywordItem(t *python.KeywordItem, fn Handler) error {
	return fn.Handle(&t.Expression)
}

func walkLambdaExpression(t *python.LambdaExpression, fn Handler) error {
	if t.ParameterList != nil {
		if err := fn.Handle(t.ParameterList); err != nil {
			return err
		}
	}

	return fn.Handle(&t.Expression)
}

func walkModuleAs(t *python.ModuleAs, fn Handler) error {
	return fn.Handle(&t.Module)
}

func walkModule(t *python.Module, fn Handler) error {
	return nil
}

func walkMultiplyExpression(t *python.MultiplyExpression, fn Handler) error {
	if err := fn.Handle(&t.UnaryExpression); err != nil {
		return err
	}

	if t.MultiplyExpression != nil {
		return fn.Handle(t.MultiplyExpression)
	}

	return nil
}

func walkNonLocalStatement(t *python.NonLocalStatement, fn Handler) error {
	return nil
}

func walkNotTest(t *python.NotTest, fn Handler) error {
	return fn.Handle(&t.Comparison)
}

func walkOrExpression(t *python.OrExpression, fn Handler) error {
	if err := fn.Handle(&t.XorExpression); err != nil {
		return err
	}

	if t.OrExpression != nil {
		return fn.Handle(t.OrExpression)
	}

	return nil
}

func walkOrTest(t *python.OrTest, fn Handler) error {
	if err := fn.Handle(&t.AndTest); err != nil {
		return err
	}

	if t.OrTest != nil {
		return fn.Handle(t.OrTest)
	}

	return nil
}

func walkParameter(t *python.Parameter, fn Handler) error {
	if t.Type != nil {
		return fn.Handle(t.Type)
	}

	return nil
}

func walkParameterList(t *python.ParameterList, fn Handler) error {
	for n := range t.DefParameters {
		if err := fn.Handle(&t.DefParameters[n]); err != nil {
			return err
		}
	}

	for n := range t.NoPosOnly {
		if err := fn.Handle(&t.NoPosOnly[n]); err != nil {
			return err
		}
	}

	if t.StarArg != nil {
		if err := fn.Handle(t.StarArg); err != nil {
			return err
		}
	}

	for n := range t.StarArgs {
		if err := fn.Handle(&t.StarArgs[n]); err != nil {
			return err
		}
	}

	if t.StarStarArg != nil {
		return fn.Handle(t.StarStarArg)
	}

	return nil
}

func walkPositionalArgument(t *python.PositionalArgument, fn Handler) error {
	if t.AssignmentExpression != nil {
		return fn.Handle(t.AssignmentExpression)
	} else if t.Expression != nil {
		return fn.Handle(t.Expression)
	}

	return nil
}

func walkPowerExpression(t *python.PowerExpression, fn Handler) error {
	if err := fn.Handle(&t.PrimaryExpression); err != nil {
		return err
	}

	if t.UnaryExpression != nil {
		return fn.Handle(t.UnaryExpression)
	}

	return nil
}

func walkPrimaryExpression(t *python.PrimaryExpression, fn Handler) error {
	if t.Atom != nil {
		return fn.Handle(t.Atom)
	} else if t.PrimaryExpression != nil {
		if err := fn.Handle(t.PrimaryExpression); err != nil {
			return err
		}

		if t.Slicing != nil {
			return fn.Handle(t.Slicing)
		} else if t.Call != nil {
			return fn.Handle(t.Call)
		}
	}

	return nil
}

func walkRaiseStatement(t *python.RaiseStatement, fn Handler) error {
	if t.Expression != nil {
		if err := fn.Handle(t.Expression); err != nil {
			return err
		}

		if t.From != nil {
			return fn.Handle(t.From)
		}
	}

	return nil
}

func walkRelativeModule(t *python.RelativeModule, fn Handler) error {
	if t.Module != nil {
		return fn.Handle(t.Module)
	}

	return nil
}

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

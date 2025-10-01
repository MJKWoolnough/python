package python

// File automatically generated with format.sh.

import "fmt"

// Type is an interface satisfied by all python structural types.
type Type interface {
	fmt.Formatter
	pythonType()
}

func (Tokens) pythonType() {}

func (AddExpression) pythonType() {}

func (AndExpression) pythonType() {}

func (AndTest) pythonType() {}

func (AnnotatedAssignmentStatement) pythonType() {}

func (ArgumentList) pythonType() {}

func (ArgumentListOrComprehension) pythonType() {}

func (AssertStatement) pythonType() {}

func (AssignmentExpressionAndSuite) pythonType() {}

func (AssignmentExpression) pythonType() {}

func (AssignmentStatement) pythonType() {}

func (Atom) pythonType() {}

func (AugmentedAssignmentStatement) pythonType() {}

func (AugTarget) pythonType() {}

func (ClassDefinition) pythonType() {}

func (Comparison) pythonType() {}

func (ComparisonExpression) pythonType() {}

func (CompoundStatement) pythonType() {}

func (Comprehension) pythonType() {}

func (ComprehensionFor) pythonType() {}

func (ComprehensionIf) pythonType() {}

func (ComprehensionIterator) pythonType() {}

func (ConditionalExpression) pythonType() {}

func (Decorator) pythonType() {}

func (Decorators) pythonType() {}

func (DefParameter) pythonType() {}

func (DelStatement) pythonType() {}

func (DictDisplay) pythonType() {}

func (DictItem) pythonType() {}

func (Enclosure) pythonType() {}

func (Except) pythonType() {}

func (Expression) pythonType() {}

func (ExpressionList) pythonType() {}

func (File) pythonType() {}

func (FlexibleExpression) pythonType() {}

func (FlexibleExpressionList) pythonType() {}

func (FlexibleExpressionListOrComprehension) pythonType() {}

func (ForStatement) pythonType() {}

func (FuncDefinition) pythonType() {}

func (GeneratorExpression) pythonType() {}

func (GlobalStatement) pythonType() {}

func (IdentifierComments) pythonType() {}

func (IfStatement) pythonType() {}

func (ImportStatement) pythonType() {}

func (KeywordArgument) pythonType() {}

func (KeywordItem) pythonType() {}

func (LambdaExpression) pythonType() {}

func (ModuleAs) pythonType() {}

func (Module) pythonType() {}

func (MultiplyExpression) pythonType() {}

func (NonLocalStatement) pythonType() {}

func (NotTest) pythonType() {}

func (OrExpression) pythonType() {}

func (OrTest) pythonType() {}

func (Parameter) pythonType() {}

func (ParameterList) pythonType() {}

func (PositionalArgument) pythonType() {}

func (PowerExpression) pythonType() {}

func (PrimaryExpression) pythonType() {}

func (RaiseStatement) pythonType() {}

func (RelativeModule) pythonType() {}

func (ReturnStatement) pythonType() {}

func (ShiftExpression) pythonType() {}

func (SimpleStatement) pythonType() {}

func (SliceItem) pythonType() {}

func (SliceList) pythonType() {}

func (StarredExpression) pythonType() {}

func (StarredItem) pythonType() {}

func (StarredList) pythonType() {}

func (StarredOrKeyword) pythonType() {}

func (Statement) pythonType() {}

func (StatementList) pythonType() {}

func (Suite) pythonType() {}

func (Target) pythonType() {}

func (TargetList) pythonType() {}

func (TryStatement) pythonType() {}

func (TypeParam) pythonType() {}

func (TypeParams) pythonType() {}

func (TypeStatement) pythonType() {}

func (UnaryExpression) pythonType() {}

func (WhileStatement) pythonType() {}

func (WithItem) pythonType() {}

func (WithStatement) pythonType() {}

func (WithStatementContents) pythonType() {}

func (XorExpression) pythonType() {}

func (YieldExpression) pythonType() {}

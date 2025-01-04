# python
--
    import "vimagination.zapto.org/python"

Package python implements a python tokeniser and parser.

## Usage

```go
const (
	TokenWhitespace parser.TokenType = iota
	TokenLineTerminator
	TokenComment
	TokenIdentifier
	TokenKeyword
	TokenOperator
	TokenDelimiter
	TokenBooleanLiteral
	TokenNumericLiteral
	TokenStringLiteral
	TokenNullLiteral
	TokenIndent
	TokenDedent
)
```
TokenType IDs.

```go
var (
	ErrInvalidCharacter      = errors.New("invalid character")
	ErrInvalidCompound       = errors.New("invalid compound statement")
	ErrInvalidEnclosure      = errors.New("invalid enclosure")
	ErrInvalidIndent         = errors.New("invalid indent")
	ErrInvalidKeyword        = errors.New("unexpected keyword")
	ErrInvalidNumber         = errors.New("invalid number")
	ErrMismatchedGroups      = errors.New("mismatched groups in except")
	ErrMissingClosingBrace   = errors.New("missing closing brace")
	ErrMissingClosingBracket = errors.New("missing closing bracket")
	ErrMissingClosingParen   = errors.New("missing closing paren")
	ErrMissingColon          = errors.New("missing colon")
	ErrMissingComma          = errors.New("missing comma")
	ErrMissingElse           = errors.New("missing else")
	ErrMissingEquals         = errors.New("missing equals")
	ErrMissingFinally        = errors.New("missing finally")
	ErrMissingFor            = errors.New("missing for keyword")
	ErrMissingIdentifier     = errors.New("missing identifier")
	ErrMissingIf             = errors.New("missing if keyword")
	ErrMissingImport         = errors.New("missing import keyword")
	ErrMissingIn             = errors.New("missing in")
	ErrMissingIndent         = errors.New("missing indent")
	ErrMissingModule         = errors.New("missing module")
	ErrMissingNewline        = errors.New("missing newline")
	ErrMissingOp             = errors.New("missing operator")
	ErrMissingOpeningParen   = errors.New("missing opening paren")
)
```
Errors

#### func  SetTokeniser

```go
func SetTokeniser(t *parser.Tokeniser) *parser.Tokeniser
```
SetTokeniser sets the initial tokeniser state of a parser.Tokeniser.

Used if you want to manually tokeniser python source code.

#### func  Unquote

```go
func Unquote(str string) (string, error)
```
Unquote takes a python quoted string and returns the unquoted string.

#### type AddExpression

```go
type AddExpression struct {
	MultiplyExpression MultiplyExpression
	Add                *Token
	AddExpression      *AddExpression
	Tokens             Tokens
}
```

AddExpression as defined in python@3.13.0:
https://docs.python.org/release/3.13.0/reference/expressions.html#grammar-token-python-grammar-a_expr

#### func (AddExpression) Format

```go
func (f AddExpression) Format(s fmt.State, v rune)
```
Format implements the fmt.Formatter interface

#### type AndExpression

```go
type AndExpression struct {
	ShiftExpression ShiftExpression
	AndExpression   *AndExpression
	Tokens          Tokens
}
```

AndExpression as defined in python@3.13.0:
https://docs.python.org/release/3.13.0/reference/expressions.html#grammar-token-python-grammar-and_expr

#### func (AndExpression) Format

```go
func (f AndExpression) Format(s fmt.State, v rune)
```
Format implements the fmt.Formatter interface

#### type AndTest

```go
type AndTest struct {
	NotTest NotTest
	AndTest *AndTest
	Tokens  Tokens
}
```

AndTest as defined in python@3.13.0:
https://docs.python.org/release/3.13.0/reference/expressions.html#grammar-token-python-grammar-and_test

#### func (AndTest) Format

```go
func (f AndTest) Format(s fmt.State, v rune)
```
Format implements the fmt.Formatter interface

#### type AnnotatedAssignmentStatement

```go
type AnnotatedAssignmentStatement struct {
	AugTarget         AugTarget
	Expression        Expression
	StarredExpression *StarredExpression
	YieldExpression   *YieldExpression
	Tokens            Tokens
}
```

AnnotatedAssignmentStatement as defined in python@3.13.0:
https://docs.python.org/release/3.13.0/reference/simple_stmts.html#grammar-token-python-grammar-annotated_assignment_stmt

#### func (AnnotatedAssignmentStatement) Format

```go
func (f AnnotatedAssignmentStatement) Format(s fmt.State, v rune)
```
Format implements the fmt.Formatter interface

#### type ArgumentList

```go
type ArgumentList struct {
	PositionalArguments        []PositionalArgument
	StarredAndKeywordArguments []StarredOrKeyword
	KeywordArguments           []KeywordArgument
	Tokens                     Tokens
}
```


#### func (ArgumentList) Format

```go
func (f ArgumentList) Format(s fmt.State, v rune)
```
Format implements the fmt.Formatter interface

#### type ArgumentListOrComprehension

```go
type ArgumentListOrComprehension struct {
	ArgumentList  *ArgumentList
	Comprehension *Comprehension
	Tokens        Tokens
}
```

ArgumentListOrComprehension as defined in python@3.13.0:
https://docs.python.org/release/3.13.0/reference/expressions.html#grammar-token-python-grammar-call

#### func (ArgumentListOrComprehension) Format

```go
func (f ArgumentListOrComprehension) Format(s fmt.State, v rune)
```
Format implements the fmt.Formatter interface

#### type AssertStatement

```go
type AssertStatement struct {
	Expressions []Expression
	Tokens      Tokens
}
```

AssertStatement as defined in python@3.13.0:
https://docs.python.org/release/3.13.0/reference/simple_stmts.html#grammar-token-python-grammar-assert_stmt

#### func (AssertStatement) Format

```go
func (f AssertStatement) Format(s fmt.State, v rune)
```
Format implements the fmt.Formatter interface

#### type AssignmentExpression

```go
type AssignmentExpression struct {
	Identifier *Token
	Expression Expression
	Tokens
}
```

AssignmentExpression as defined in python@3.13.0:
https://docs.python.org/release/3.13.0/reference/expressions.html#grammar-token-python-grammar-assignment_expression

#### func (AssignmentExpression) Format

```go
func (f AssignmentExpression) Format(s fmt.State, v rune)
```
Format implements the fmt.Formatter interface

#### type AssignmentExpressionAndSuite

```go
type AssignmentExpressionAndSuite struct {
	AssignmentExpression AssignmentExpression
	Suite                Suite
}
```

AssignmentExpressionAndSuite is a combination of the AssignmentExpression and
Suite types.

#### func (AssignmentExpressionAndSuite) Format

```go
func (f AssignmentExpressionAndSuite) Format(s fmt.State, v rune)
```
Format implements the fmt.Formatter interface

#### type AssignmentStatement

```go
type AssignmentStatement struct {
	TargetLists       []TargetList
	StarredExpression *StarredExpression
	YieldExpression   *YieldExpression
	Tokens            Tokens
}
```

AssignmentStatement as defined in python:3.13.0:
https://docs.python.org/release/3.13.0/reference/simple_stmts.html#grammar-token-python-grammar-assignment_stmt

#### func (AssignmentStatement) Format

```go
func (f AssignmentStatement) Format(s fmt.State, v rune)
```
Format implements the fmt.Formatter interface

#### type Atom

```go
type Atom struct {
	Identifier *Token
	Literal    *Token
	Enclosure  *Enclosure
	Tokens     Tokens
}
```

Atom as defined in python@3.13.0:
https://docs.python.org/release/3.13.0/reference/expressions.html#grammar-token-python-grammar-atom

#### func (Atom) Format

```go
func (f Atom) Format(s fmt.State, v rune)
```
Format implements the fmt.Formatter interface

#### func (*Atom) IsIdentifier

```go
func (a *Atom) IsIdentifier() bool
```
IsIdentifier returns true if the Atom contains an Idneitifer.

#### type AugTarget

```go
type AugTarget struct {
	PrimaryExpression PrimaryExpression
	Tokens            Tokens
}
```

AugTarget as defined in python@3.13.0:
https://docs.python.org/release/3.13.0/reference/simple_stmts.html#grammar-token-python-grammar-augtarget

#### func (AugTarget) Format

```go
func (f AugTarget) Format(s fmt.State, v rune)
```
Format implements the fmt.Formatter interface

#### type AugmentedAssignmentStatement

```go
type AugmentedAssignmentStatement struct {
	AugTarget       AugTarget
	AugOp           *Token
	ExpressionList  *ExpressionList
	YieldExpression *YieldExpression
	Tokens          Tokens
}
```

AugmentedAssignmentStatement as defined in python@3.13.0:
https://docs.python.org/release/3.13.0/reference/simple_stmts.html#grammar-token-python-grammar-augmented_assignment_stmt

#### func (AugmentedAssignmentStatement) Format

```go
func (f AugmentedAssignmentStatement) Format(s fmt.State, v rune)
```
Format implements the fmt.Formatter interface

#### type ClassDefinition

```go
type ClassDefinition struct {
	Decorators  *Decorators
	ClassName   *Token
	TypeParams  *TypeParams
	Inheritance *ArgumentList
	Suite       Suite
	Tokens      Tokens
}
```

ClassDefinition as defined in python@3.13.0:
https://docs.python.org/3.13.0/reference/compound_stmts.html#grammar-token-python-grammar-classdef

#### func (ClassDefinition) Format

```go
func (f ClassDefinition) Format(s fmt.State, v rune)
```
Format implements the fmt.Formatter interface

#### type Comparison

```go
type Comparison struct {
	OrExpression OrExpression
	Comparisons  []ComparisonExpression
	Tokens       Tokens
}
```

Comparison as defined in python@3.13.0:
https://docs.python.org/release/3.13.0/reference/expressions.html#grammar-token-python-grammar-comparison

#### func (Comparison) Format

```go
func (f Comparison) Format(s fmt.State, v rune)
```
Format implements the fmt.Formatter interface

#### type ComparisonExpression

```go
type ComparisonExpression struct {
	ComparisonOperator []Token
	OrExpression       OrExpression
}
```

ComparisonExpression combines combines the operators with an OrExpression.

#### func (ComparisonExpression) Format

```go
func (f ComparisonExpression) Format(s fmt.State, v rune)
```
Format implements the fmt.Formatter interface

#### type CompoundStatement

```go
type CompoundStatement struct {
	If     *IfStatement
	While  *WhileStatement
	For    *ForStatement
	Try    *TryStatement
	With   *WithStatement
	Func   *FuncDefinition
	Class  *ClassDefinition
	Tokens Tokens
}
```

CompoundStatement as defined in python@3.13.0:
https://docs.python.org/3.13.0/reference/compound_stmts.html#grammar-token-python-grammar-compound_stmt

#### func (CompoundStatement) Format

```go
func (f CompoundStatement) Format(s fmt.State, v rune)
```
Format implements the fmt.Formatter interface

#### type Comprehension

```go
type Comprehension struct {
	AssignmentExpression AssignmentExpression
	ComprehensionFor     ComprehensionFor
	Tokens               Tokens
}
```

Comprehension as defined in python@3.13.0:
https://docs.python.org/release/3.13.0/reference/expressions.html#grammar-token-python-grammar-comprehension

#### func (Comprehension) Format

```go
func (f Comprehension) Format(s fmt.State, v rune)
```
Format implements the fmt.Formatter interface

#### type ComprehensionFor

```go
type ComprehensionFor struct {
	Async                 bool
	TargetList            TargetList
	OrTest                OrTest
	ComprehensionIterator *ComprehensionIterator
	Tokens                Tokens
}
```

ComprehensionFor as defined in python@3.13.0:
https://docs.python.org/release/3.13.0/reference/expressions.html#grammar-token-python-grammar-comp_for

#### func (ComprehensionFor) Format

```go
func (f ComprehensionFor) Format(s fmt.State, v rune)
```
Format implements the fmt.Formatter interface

#### type ComprehensionIf

```go
type ComprehensionIf struct {
	OrTest                OrTest
	ComprehensionIterator *ComprehensionIterator
	Tokens                Tokens
}
```

ComprehensionIf as defined in python@3.13.0:
https://docs.python.org/release/3.13.0/reference/expressions.html#grammar-token-python-grammar-comp_if

#### func (ComprehensionIf) Format

```go
func (f ComprehensionIf) Format(s fmt.State, v rune)
```
Format implements the fmt.Formatter interface

#### type ComprehensionIterator

```go
type ComprehensionIterator struct {
	ComprehensionFor *ComprehensionFor
	ComprehensionIf  *ComprehensionIf
	Tokens           Tokens
}
```

ComprehensionIterator as defined in python@3.13.0:
https://docs.python.org/release/3.13.0/reference/expressions.html#grammar-token-python-grammar-comp_iter

#### func (ComprehensionIterator) Format

```go
func (f ComprehensionIterator) Format(s fmt.State, v rune)
```
Format implements the fmt.Formatter interface

#### type ConditionalExpression

```go
type ConditionalExpression struct {
	OrTest OrTest
	If     *OrTest
	Else   *Expression
	Tokens Tokens
}
```

ConditionalExpression as defined in python@3.13.0:
https://docs.python.org/release/3.13.0/reference/expressions.html#grammar-token-python-grammar-conditional_expression

#### func  WrapConditional

```go
func WrapConditional(p ConditionalWrappable) *ConditionalExpression
```
WrapConditional takes one of many types and wraps it in a
*ConditionalExpression.

The accepted types/pointers are as follows:

    ConditionalExpression
    *ConditionalExpression
    OrTest
    *OrTest
    AndTest
    *AndTest
    NotTest
    *NotTest
    Comparison
    *Comparison
    OrExpression
    *OrExpression
    XorExpression
    *XorExpression
    AndExpression
    *AndExpression
    ShiftExpression
    *ShiftExpression
    AddExpression
    *AddExpression
    MultiplyExpression
    *MultiplyExpression
    UnaryExpression
    *UnaryExpression
    PowerExpression
    *PowerExpression
    PrimaryExpression
    *PrimaryExpression
    Atom
    *Atom

#### func (ConditionalExpression) Format

```go
func (f ConditionalExpression) Format(s fmt.State, v rune)
```
Format implements the fmt.Formatter interface

#### type ConditionalWrappable

```go
type ConditionalWrappable interface {
	Type
	// contains filtered or unexported methods
}
```

ConditionalWrappable represents the types that can be wrapped with
WrapConditional and unwrapped with UnwrapConditional.

#### func  UnwrapConditional

```go
func UnwrapConditional(c *ConditionalExpression) ConditionalWrappable
```
Possible returns types are as follows:

    *ConditionalExpression
    *OrTest
    *AndTest
    *NotTest
    *Comparison
    *OrExpression
    *XorExpression
    *AndExpression
    *ShiftExpression
    *AddExpression
    *MultiplyExpression
    *UnaryExpression
    *PowerExpression
    *PrimaryExpression
    *Atom

#### type Decorators

```go
type Decorators struct {
	Decorators []AssignmentExpression
	Tokens
}
```

Decorators as defined in python@3.13.0:
https://docs.python.org/3.13.0/reference/compound_stmts.html#grammar-token-python-grammar-decorators

#### func (Decorators) Format

```go
func (f Decorators) Format(s fmt.State, v rune)
```
Format implements the fmt.Formatter interface

#### type DefParameter

```go
type DefParameter struct {
	Parameter Parameter
	Value     *Expression
	Tokens    Tokens
}
```

DefParameter as defined in python@3.13.0:
https://docs.python.org/release/3.13.0/reference/compound_stmts.html#grammar-token-python-grammar-defparameter

#### func (DefParameter) Format

```go
func (f DefParameter) Format(s fmt.State, v rune)
```
Format implements the fmt.Formatter interface

#### type DelStatement

```go
type DelStatement struct {
	TargetList TargetList
	Tokens     Tokens
}
```

DelStatement as defined in python@3.13.0:
https://docs.python.org/release/3.13.0/reference/simple_stmts.html#grammar-token-python-grammar-del_stmt

#### func (DelStatement) Format

```go
func (f DelStatement) Format(s fmt.State, v rune)
```
Format implements the fmt.Formatter interface

#### type DictDisplay

```go
type DictDisplay struct {
	DictItems         []DictItem
	DictComprehension *ComprehensionFor
	Tokens            Tokens
}
```

DictDisplay as defined in python@3.13.0:
https://docs.python.org/release/3.13.0/reference/expressions.html#grammar-token-python-grammar-dict_display

#### func (DictDisplay) Format

```go
func (f DictDisplay) Format(s fmt.State, v rune)
```
Format implements the fmt.Formatter interface

#### type DictItem

```go
type DictItem struct {
	Key          *Expression
	Value        *Expression
	OrExpression *OrExpression
	Tokens       Tokens
}
```

DictItem as defined in python@3.13.0:
https://docs.python.org/release/3.13.0/reference/expressions.html#grammar-token-python-grammar-dict_item

#### func (DictItem) Format

```go
func (f DictItem) Format(s fmt.State, v rune)
```
Format implements the fmt.Formatter interface

#### type Enclosure

```go
type Enclosure struct {
	ParenthForm         *StarredExpression
	ListDisplay         *FlexibleExpressionListOrComprehension
	DictDisplay         *DictDisplay
	SetDisplay          *FlexibleExpressionListOrComprehension
	GeneratorExpression *GeneratorExpression
	YieldAtom           *YieldExpression
	Tokens              Tokens
}
```

Enclosure as defined in python@3.13.0:
https://docs.python.org/release/3.13.0/reference/expressions.html#grammar-token-python-grammar-enclosure

#### func (Enclosure) Format

```go
func (f Enclosure) Format(s fmt.State, v rune)
```
Format implements the fmt.Formatter interface

#### type Error

```go
type Error struct {
	Err     error
	Parsing string
	Token   Token
}
```

Error represents a Python parsing error.

#### func (Error) Error

```go
func (e Error) Error() string
```
Error implements the error interface.

#### func (Error) Unwrap

```go
func (e Error) Unwrap() error
```
Unwrap returns the underlying error.

#### type Except

```go
type Except struct {
	Expression Expression
	Identifier *Token
	Suite      Suite
	Tokens     Tokens
}
```

Except as defined in python@3.13.0:
https://docs.python.org/3.13.0/reference/compound_stmts.html#grammar-token-python-grammar-try1_stmt

#### func (Except) Format

```go
func (f Except) Format(s fmt.State, v rune)
```
Format implements the fmt.Formatter interface

#### type Expression

```go
type Expression struct {
	ConditionalExpression *ConditionalExpression
	LambdaExpression      *LambdaExpression
	Tokens                Tokens
}
```

Expression as defined in python@3.13.0:
https://docs.python.org/release/3.13.0/reference/expressions.html#grammar-token-python-grammar-expression

#### func (Expression) Format

```go
func (f Expression) Format(s fmt.State, v rune)
```
Format implements the fmt.Formatter interface

#### type ExpressionList

```go
type ExpressionList struct {
	Expressions []Expression
	Tokens      Tokens
}
```

ExpressionList as defined in python@3.13.0:
https://docs.python.org/release/3.13.0/reference/expressions.html#grammar-token-python-grammar-expression_list

#### func (ExpressionList) Format

```go
func (f ExpressionList) Format(s fmt.State, v rune)
```
Format implements the fmt.Formatter interface

#### type File

```go
type File struct {
	Statements []Statement
	Tokens     Tokens
}
```

File represents a parsed Python file.

#### func  Parse

```go
func Parse(t Tokeniser) (*File, error)
```
Parse parses Python input into AST.

#### func (File) Format

```go
func (f File) Format(s fmt.State, v rune)
```
Format implements the fmt.Formatter interface

#### type FlexibleExpression

```go
type FlexibleExpression struct {
	AssignmentExpression *AssignmentExpression
	StarredExpression    *StarredExpression
	Tokens
}
```

FlexibleExpression as defined in python@3.13.0:
https://docs.python.org/release/3.13.0/reference/expressions.html#grammar-token-python-grammar-flexible_expression

#### func (FlexibleExpression) Format

```go
func (f FlexibleExpression) Format(s fmt.State, v rune)
```
Format implements the fmt.Formatter interface

#### type FlexibleExpressionList

```go
type FlexibleExpressionList struct {
	FlexibleExpressions []FlexibleExpression
	Tokens
}
```

FlexibleExpressionList as defined in python@3.13.0:
https://docs.python.org/release/3.13.0/reference/expressions.html#grammar-token-python-grammar-flexible_expression_list

#### func (FlexibleExpressionList) Format

```go
func (f FlexibleExpressionList) Format(s fmt.State, v rune)
```
Format implements the fmt.Formatter interface

#### type FlexibleExpressionListOrComprehension

```go
type FlexibleExpressionListOrComprehension struct {
	FlexibleExpressionList *FlexibleExpressionList
	Comprehension          *Comprehension
	Tokens                 Tokens
}
```

FlexibleExpressionListOrComprehension as defined in python@3.13.0:
https://docs.python.org/release/3.13.0/reference/expressions.html#grammar-token-python-grammar-list_display

#### func (FlexibleExpressionListOrComprehension) Format

```go
func (f FlexibleExpressionListOrComprehension) Format(s fmt.State, v rune)
```
Format implements the fmt.Formatter interface

#### type ForStatement

```go
type ForStatement struct {
	Async       bool
	TargetList  TargetList
	StarredList StarredList
	Suite       Suite
	Else        *Suite
	Tokens
}
```

ForStatement as defined in python@3.13.0:
https://docs.python.org/3.13.0/reference/compound_stmts.html#grammar-token-python-grammar-for_stmt

#### func (ForStatement) Format

```go
func (f ForStatement) Format(s fmt.State, v rune)
```
Format implements the fmt.Formatter interface

#### type FuncDefinition

```go
type FuncDefinition struct {
	Decorators    *Decorators
	Async         bool
	FuncName      *Token
	TypeParams    *TypeParams
	ParameterList ParameterList
	Expression    *Expression
	Suite         Suite
	Tokens        Tokens
}
```

FuncDefinition as defined in python@3.13.0:
https://docs.python.org/3.13.0/reference/compound_stmts.html#grammar-token-python-grammar-funcdef

#### func (FuncDefinition) Format

```go
func (f FuncDefinition) Format(s fmt.State, v rune)
```
Format implements the fmt.Formatter interface

#### type GeneratorExpression

```go
type GeneratorExpression struct {
	Expression       Expression
	ComprehensionFor ComprehensionFor
	Tokens           Tokens
}
```

GeneratorExpression as defined in python@3.13.0:
https://docs.python.org/release/3.13.0/reference/expressions.html#grammar-token-python-grammar-generator_expression

#### func (GeneratorExpression) Format

```go
func (f GeneratorExpression) Format(s fmt.State, v rune)
```
Format implements the fmt.Formatter interface

#### type GlobalStatement

```go
type GlobalStatement struct {
	Identifiers []*Token
	Tokens      Tokens
}
```

GlobalStatement as defined in python@3.13.0:
https://docs.python.org/release/3.13.0/reference/simple_stmts.html#grammar-token-python-grammar-global_stmt

#### func (GlobalStatement) Format

```go
func (f GlobalStatement) Format(s fmt.State, v rune)
```
Format implements the fmt.Formatter interface

#### type IfStatement

```go
type IfStatement struct {
	AssignmentExpression AssignmentExpression
	Suite                Suite
	Elif                 []AssignmentExpressionAndSuite
	Else                 *Suite
	Tokens               Tokens
}
```

IfStatement as defined in python@3.13.0:
https://docs.python.org/3.13.0/reference/compound_stmts.html#grammar-token-python-grammar-if_stmt

#### func (IfStatement) Format

```go
func (f IfStatement) Format(s fmt.State, v rune)
```
Format implements the fmt.Formatter interface

#### type ImportStatement

```go
type ImportStatement struct {
	RelativeModule *RelativeModule
	Modules        []ModuleAs
	Tokens         Tokens
}
```

ImportStatement as defined in python@3.13.0:
https://docs.python.org/release/3.13.0/reference/simple_stmts.html#grammar-token-python-grammar-import_stmt

#### func (ImportStatement) Format

```go
func (f ImportStatement) Format(s fmt.State, v rune)
```
Format implements the fmt.Formatter interface

#### type KeywordArgument

```go
type KeywordArgument struct {
	KeywordItem *KeywordItem
	Expression  *Expression
	Tokens      Tokens
}
```

KeywordArgument as defined in python@3.13.0:
https://docs.python.org/release/3.13.0/reference/expressions.html#grammar-token-python-grammar-keywords_arguments

#### func (KeywordArgument) Format

```go
func (f KeywordArgument) Format(s fmt.State, v rune)
```
Format implements the fmt.Formatter interface

#### type KeywordItem

```go
type KeywordItem struct {
	Identifier *Token
	Expression Expression
	Tokens     Tokens
}
```

KeywordItem as defined in python@3.13.0:
https://docs.python.org/release/3.13.0/reference/expressions.html#grammar-token-python-grammar-keyword_item

#### func (KeywordItem) Format

```go
func (f KeywordItem) Format(s fmt.State, v rune)
```
Format implements the fmt.Formatter interface

#### type LambdaExpression

```go
type LambdaExpression struct {
	ParameterList *ParameterList
	Expression    Expression
	Tokens        Tokens
}
```

LambdaExpression as defined in python@3.13.0:
https://docs.python.org/release/3.13.0/reference/expressions.html#grammar-token-python-grammar-lambda_expr

#### func (LambdaExpression) Format

```go
func (f LambdaExpression) Format(s fmt.State, v rune)
```
Format implements the fmt.Formatter interface

#### type Module

```go
type Module struct {
	Identifiers []*Token
	Tokens      Tokens
}
```

Module as defined in python@3.13.0:
https://docs.python.org/release/3.13.0/reference/simple_stmts.html#grammar-token-python-grammar-module

#### func (Module) Format

```go
func (f Module) Format(s fmt.State, v rune)
```
Format implements the fmt.Formatter interface

#### type ModuleAs

```go
type ModuleAs struct {
	Module Module
	As     *Token
	Tokens Tokens
}
```

ModuleAs as defined in python@3.13.0:
https://docs.python.org/release/3.13.0/reference/simple_stmts.html#grammar-token-python-grammar-import_stmt

#### func (ModuleAs) Format

```go
func (f ModuleAs) Format(s fmt.State, v rune)
```
Format implements the fmt.Formatter interface

#### type MultiplyExpression

```go
type MultiplyExpression struct {
	UnaryExpression    UnaryExpression
	Multiply           *Token
	MultiplyExpression *MultiplyExpression
	Tokens             Tokens
}
```

MultiplyExpression as defined in python@3.13.0:
https://docs.python.org/release/3.13.0/reference/expressions.html#grammar-token-python-grammar-m_expr

#### func (MultiplyExpression) Format

```go
func (f MultiplyExpression) Format(s fmt.State, v rune)
```
Format implements the fmt.Formatter interface

#### type NonLocalStatement

```go
type NonLocalStatement struct {
	Identifiers []*Token
	Tokens      Tokens
}
```

NonLocalStatement as defined in python@3.13.0:
https://docs.python.org/release/3.13.0/reference/simple_stmts.html#grammar-token-python-grammar-nonlocal_stmt

#### func (NonLocalStatement) Format

```go
func (f NonLocalStatement) Format(s fmt.State, v rune)
```
Format implements the fmt.Formatter interface

#### type NotTest

```go
type NotTest struct {
	Nots       uint
	Comparison Comparison
	Tokens     Tokens
}
```

NotTest as defined in python@3.13.0:
https://docs.python.org/release/3.13.0/reference/expressions.html#grammar-token-python-grammar-not_test

#### func (NotTest) Format

```go
func (f NotTest) Format(s fmt.State, v rune)
```
Format implements the fmt.Formatter interface

#### type OrExpression

```go
type OrExpression struct {
	XorExpression XorExpression
	OrExpression  *OrExpression
	Tokens        Tokens
}
```

OrExpression as defined in python@3.13.0:
https://docs.python.org/release/3.13.0/reference/expressions.html#grammar-token-python-grammar-or_expr

#### func (OrExpression) Format

```go
func (f OrExpression) Format(s fmt.State, v rune)
```
Format implements the fmt.Formatter interface

#### type OrTest

```go
type OrTest struct {
	AndTest AndTest
	OrTest  *OrTest
	Tokens  Tokens
}
```

OrTest as defined in python@3.13.0:
https://docs.python.org/release/3.13.0/reference/expressions.html#grammar-token-python-grammar-or_test

#### func (OrTest) Format

```go
func (f OrTest) Format(s fmt.State, v rune)
```
Format implements the fmt.Formatter interface

#### type Parameter

```go
type Parameter struct {
	Identifier *Token
	Type       *Expression
	Tokens     Tokens
}
```

Parameter as defined in python@3.13.0:
https://docs.python.org/release/3.13.0/reference/compound_stmts.html#grammar-token-python-grammar-parameter

#### func (Parameter) Format

```go
func (f Parameter) Format(s fmt.State, v rune)
```
Format implements the fmt.Formatter interface

#### type ParameterList

```go
type ParameterList struct {
	DefParameters []DefParameter
	NoPosOnly     []DefParameter
	StarArg       *Parameter
	StarArgs      []DefParameter
	StarStarArg   *Parameter
	Tokens        Tokens
}
```

ParameterList as defined in python@3.13.0:
https://docs.python.org/release/3.13.0/reference/compound_stmts.html#grammar-token-python-grammar-parameter_list

#### func (ParameterList) Format

```go
func (f ParameterList) Format(s fmt.State, v rune)
```
Format implements the fmt.Formatter interface

#### type PositionalArgument

```go
type PositionalArgument struct {
	AssignmentExpression *AssignmentExpression
	Expression           *Expression
	Tokens               Tokens
}
```

PositionalArgument as defined in python@3.13.0:
https://docs.python.org/release/3.13.0/reference/expressions.html#grammar-token-python-grammar-positional_arguments

#### func (PositionalArgument) Format

```go
func (f PositionalArgument) Format(s fmt.State, v rune)
```
Format implements the fmt.Formatter interface

#### type PowerExpression

```go
type PowerExpression struct {
	AwaitExpression   bool
	PrimaryExpression PrimaryExpression
	UnaryExpression   *UnaryExpression
	Tokens            Tokens
}
```

PowerExpression as defined in python@3.13.0:
https://docs.python.org/release/3.13.0/reference/expressions.html#grammar-token-python-grammar-power

#### func (PowerExpression) Format

```go
func (f PowerExpression) Format(s fmt.State, v rune)
```
Format implements the fmt.Formatter interface

#### type PrimaryExpression

```go
type PrimaryExpression struct {
	PrimaryExpression *PrimaryExpression
	Atom              *Atom
	AttributeRef      *Token
	Slicing           *SliceList
	Call              *ArgumentListOrComprehension
	Tokens            Tokens
}
```

PrimaryExpression as defined in python@3.13.0:
https://docs.python.org/release/3.13.0/reference/expressions.html#grammar-token-python-grammar-primary

#### func (PrimaryExpression) Format

```go
func (f PrimaryExpression) Format(s fmt.State, v rune)
```
Format implements the fmt.Formatter interface

#### func (*PrimaryExpression) IsIdentifier

```go
func (pr *PrimaryExpression) IsIdentifier() bool
```
IsIdentifier returns true if the Primary expression is based on an Identifier.

#### type RaiseStatement

```go
type RaiseStatement struct {
	Expression *Expression
	From       *Expression
	Tokens     Tokens
}
```

RaiseStatement as defined in python@3.13.0:
https://docs.python.org/release/3.13.0/reference/simple_stmts.html#grammar-token-python-grammar-raise_stmt

#### func (RaiseStatement) Format

```go
func (f RaiseStatement) Format(s fmt.State, v rune)
```
Format implements the fmt.Formatter interface

#### type RelativeModule

```go
type RelativeModule struct {
	Dots   int
	Module *Module
	Tokens Tokens
}
```

RelativeModule as defined in python@3.13.0:
https://docs.python.org/release/3.13.0/reference/simple_stmts.html#grammar-token-python-grammar-relative_module

#### func (RelativeModule) Format

```go
func (f RelativeModule) Format(s fmt.State, v rune)
```
Format implements the fmt.Formatter interface

#### type ReturnStatement

```go
type ReturnStatement struct {
	Expression *Expression
	Tokens     Tokens
}
```

ReturnStatement as defined in python@3.13.0:
https://docs.python.org/release/3.13.0/reference/simple_stmts.html#grammar-token-python-grammar-return_stmt

#### func (ReturnStatement) Format

```go
func (f ReturnStatement) Format(s fmt.State, v rune)
```
Format implements the fmt.Formatter interface

#### type ShiftExpression

```go
type ShiftExpression struct {
	AddExpression   AddExpression
	Shift           *Token
	ShiftExpression *ShiftExpression
	Tokens          Tokens
}
```

ShiftExpression as defined in python@3.13.0:
https://docs.python.org/release/3.13.0/reference/expressions.html#grammar-token-python-grammar-shift_expr

#### func (ShiftExpression) Format

```go
func (f ShiftExpression) Format(s fmt.State, v rune)
```
Format implements the fmt.Formatter interface

#### type SimpleStatement

```go
type SimpleStatement struct {
	Type                         StatementType
	AssertStatement              *AssertStatement
	AssignmentStatement          *AssignmentStatement
	AugmentedAssignmentStatement *AugmentedAssignmentStatement
	AnnotatedAssignmentStatement *AnnotatedAssignmentStatement
	DelStatement                 *DelStatement
	ReturnStatement              *ReturnStatement
	YieldStatement               *YieldExpression
	RaiseStatement               *RaiseStatement
	ImportStatement              *ImportStatement
	GlobalStatement              *GlobalStatement
	NonLocalStatement            *NonLocalStatement
	TypeStatement                *TypeStatement
	Tokens                       Tokens
}
```

SimpleStatement as defined in python@3.13.0:
https://docs.python.org/release/3.13.0/reference/simple_stmts.html#grammar-token-python-grammar-simple_stmt

#### func (SimpleStatement) Format

```go
func (f SimpleStatement) Format(s fmt.State, v rune)
```
Format implements the fmt.Formatter interface

#### type SliceItem

```go
type SliceItem struct {
	Expression *Expression
	LowerBound *Expression
	UpperBound *Expression
	Stride     *Expression
	Tokens     Tokens
}
```

SliceItem as defined in python@3.13.0:
https://docs.python.org/release/3.13.0/reference/expressions.html#grammar-token-python-grammar-slice_item

#### func (SliceItem) Format

```go
func (f SliceItem) Format(s fmt.State, v rune)
```
Format implements the fmt.Formatter interface

#### type SliceList

```go
type SliceList struct {
	SliceItems []SliceItem
	Tokens     Tokens
}
```

SliceList as defined in python@3.13.0:
https://docs.python.org/release/3.13.0/reference/expressions.html#grammar-token-python-grammar-slice_list

#### func (SliceList) Format

```go
func (f SliceList) Format(s fmt.State, v rune)
```
Format implements the fmt.Formatter interface

#### type StarredExpression

```go
type StarredExpression struct {
	Expression  *Expression
	StarredList *StarredList
	Tokens      Tokens
}
```

StarredExpression as defined in python@3.12.6:
https://docs.python.org/release/3.12.6/reference/expressions.html#grammar-token-python-grammar-starred_expression

#### func (StarredExpression) Format

```go
func (f StarredExpression) Format(s fmt.State, v rune)
```
Format implements the fmt.Formatter interface

#### type StarredItem

```go
type StarredItem struct {
	AssignmentExpression *AssignmentExpression
	OrExpr               *OrExpression
	Tokens               Tokens
}
```

StarredItem as defined in python@3.12.6:
https://docs.python.org/release/3.12.6/reference/expressions.html#grammar-token-python-grammar-starred_item

#### func (StarredItem) Format

```go
func (f StarredItem) Format(s fmt.State, v rune)
```
Format implements the fmt.Formatter interface

#### type StarredList

```go
type StarredList struct {
	StarredItems  []StarredItem
	TrailingComma bool
	Tokens        Tokens
}
```

StarredList as defined in python@3.12.6:
https://docs.python.org/release/3.12.6/reference/expressions.html#grammar-token-python-grammar-starred_list

#### func (StarredList) Format

```go
func (f StarredList) Format(s fmt.State, v rune)
```
Format implements the fmt.Formatter interface

#### type StarredOrKeyword

```go
type StarredOrKeyword struct {
	Expression  *Expression
	KeywordItem *KeywordItem
	Tokens      Tokens
}
```

StarredOrKeyword as defined in python@3.13.0:
https://docs.python.org/release/3.13.0/reference/expressions.html#grammar-token-python-grammar-starred_and_keywords

#### func (StarredOrKeyword) Format

```go
func (f StarredOrKeyword) Format(s fmt.State, v rune)
```
Format implements the fmt.Formatter interface

#### type Statement

```go
type Statement struct {
	StatementList     *StatementList
	CompoundStatement *CompoundStatement
	Tokens            Tokens
}
```

Statement as defined in python@3.13.0:
https://docs.python.org/release/3.13.0/reference/compound_stmts.html#grammar-token-python-grammar-statement

#### func (Statement) Format

```go
func (f Statement) Format(s fmt.State, v rune)
```
Format implements the fmt.Formatter interface

#### type StatementList

```go
type StatementList struct {
	Statements []SimpleStatement
	Tokens
}
```

StatementList as defined in python@3.13.0:
https://docs.python.org/release/3.13.0/reference/compound_stmts.html#grammar-token-python-grammar-stmt_list

#### func (StatementList) Format

```go
func (f StatementList) Format(s fmt.State, v rune)
```
Format implements the fmt.Formatter interface

#### type StatementType

```go
type StatementType uint8
```

StatementType specifies the type of a SimpleStatment.

```go
const (
	StatementAssert StatementType = iota
	StatementAssignment
	StatementAugmentedAssignment
	StatementAnnotatedAssignment
	StatementPass
	StatementDel
	StatementReturn
	StatementYield
	StatementRaise
	StatementBreak
	StatementContinue
	StatementImport
	StatementGlobal
	StatementNonLocal
	StatementTyp
)
```

#### func (StatementType) String

```go
func (s StatementType) String() string
```
String implements the fmt.Stringer interface.

#### type Suite

```go
type Suite struct {
	StatementList *StatementList
	Statements    []Statement
	Tokens        Tokens
}
```

Suite as defined in python@3.13.0:
https://docs.python.org/3.13.0/reference/compound_stmts.html#grammar-token-python-grammar-suite

#### func (Suite) Format

```go
func (f Suite) Format(s fmt.State, v rune)
```
Format implements the fmt.Formatter interface

#### type Target

```go
type Target struct {
	PrimaryExpression *PrimaryExpression
	Tuple             *TargetList
	Array             *TargetList
	Star              *Target
	Tokens            Tokens
}
```

Target as defined in python@3.13.0:
https://docs.python.org/3.13.0/reference/simple_stmts.html#grammar-token-python-grammar-target

#### func (Target) Format

```go
func (f Target) Format(s fmt.State, v rune)
```
Format implements the fmt.Formatter interface

#### type TargetList

```go
type TargetList struct {
	Targets []Target
	Tokens  Tokens
}
```

TargetList as defined in python@3.13.0:
https://docs.python.org/3.13.0/reference/simple_stmts.html#grammar-token-python-grammar-target_list

#### func (TargetList) Format

```go
func (f TargetList) Format(s fmt.State, v rune)
```
Format implements the fmt.Formatter interface

#### type Token

```go
type Token struct {
	parser.Token
	Pos, Line, LinePos uint64
}
```

Token represents a parser.Token combined with positioning information.

#### type Tokeniser

```go
type Tokeniser interface {
	Iter(func(parser.Token) bool)
	TokeniserState(parser.TokenFunc)
	GetError() error
}
```

Tokeniser represents the methods required by the python tokeniser.

#### type Tokens

```go
type Tokens []Token
```

Tokens represents a list ok tokens that have been parsed.

#### type TryStatement

```go
type TryStatement struct {
	Try     Suite
	Groups  bool
	Except  []Except
	Else    *Suite
	Finally *Suite
	Tokens  Tokens
}
```

TryStatement as defined in python@3.13.0:
https://docs.python.org/3.13.0/reference/compound_stmts.html#grammar-token-python-grammar-try_stmt

#### func (TryStatement) Format

```go
func (f TryStatement) Format(s fmt.State, v rune)
```
Format implements the fmt.Formatter interface

#### type Type

```go
type Type interface {
	fmt.Formatter
	// contains filtered or unexported methods
}
```

Type is an interface satisfied by all python structural types.

#### type TypeParam

```go
type TypeParam struct {
	Type       TypeParamType
	Identifier *Token
	Expression *Expression
	Tokens     Tokens
}
```

TypeParam as defined in python@3.13.0:
https://docs.python.org/release/3.13.0/reference/compound_stmts.html#grammar-token-python-grammar-type_param

#### func (TypeParam) Format

```go
func (f TypeParam) Format(s fmt.State, v rune)
```
Format implements the fmt.Formatter interface

#### type TypeParamType

```go
type TypeParamType byte
```

TypeParamType determines the type of a TypeParam.

```go
const (
	TypeParamIdentifer TypeParamType = iota
	TypeParamVar
	TypeParamVarTuple
)
```

#### func (TypeParamType) String

```go
func (t TypeParamType) String() string
```
String implements the fmt.Stringer interface.

#### type TypeParams

```go
type TypeParams struct {
	TypeParams []TypeParam
	Tokens     Tokens
}
```

TypeParams as defined in python@3.13.0:
https://docs.python.org/release/3.13.0/reference/compound_stmts.html#grammar-token-python-grammar-type_params

#### func (TypeParams) Format

```go
func (f TypeParams) Format(s fmt.State, v rune)
```
Format implements the fmt.Formatter interface

#### type TypeStatement

```go
type TypeStatement struct {
	Identifier *Token
	TypeParams *TypeParams
	Expression Expression
	Tokens     Tokens
}
```

TypeStatement as defined in python@3.13.0:
https://docs.python.org/release/3.13.0/reference/simple_stmts.html#grammar-token-python-grammar-type_stmt

#### func (TypeStatement) Format

```go
func (f TypeStatement) Format(s fmt.State, v rune)
```
Format implements the fmt.Formatter interface

#### type UnaryExpression

```go
type UnaryExpression struct {
	PowerExpression *PowerExpression
	Unary           *Token
	UnaryExpression *UnaryExpression
	Tokens          Tokens
}
```

UnaryExpression as defined in python@3.13.0:
https://docs.python.org/release/3.13.0/reference/expressions.html#grammar-token-python-grammar-u_expr

#### func (UnaryExpression) Format

```go
func (f UnaryExpression) Format(s fmt.State, v rune)
```
Format implements the fmt.Formatter interface

#### type WhileStatement

```go
type WhileStatement struct {
	AssignmentExpression AssignmentExpression
	Suite                Suite
	Else                 *Suite
	Tokens               Tokens
}
```

WhileStatement as defined in python@3.13.0:
https://docs.python.org/3.13.0/reference/compound_stmts.html#grammar-token-python-grammar-while_stmt

#### func (WhileStatement) Format

```go
func (f WhileStatement) Format(s fmt.State, v rune)
```
Format implements the fmt.Formatter interface

#### type WithItem

```go
type WithItem struct {
	Expression Expression
	Target     *Target
	Tokens     Tokens
}
```

WithItem as defined in python@3.13:
https://docs.python.org/3.13/reference/compound_stmts.html#grammar-token-python-grammar-with_item

#### func (WithItem) Format

```go
func (f WithItem) Format(s fmt.State, v rune)
```
Format implements the fmt.Formatter interface

#### type WithStatement

```go
type WithStatement struct {
	Async    bool
	Contents WithStatementContents
	Suite    Suite
	Tokens   Tokens
}
```

WithStatement as defined in python@3.13.0:
https://docs.python.org/3.13.0/reference/compound_stmts.html#grammar-token-python-grammar-with_stmt

#### func (WithStatement) Format

```go
func (f WithStatement) Format(s fmt.State, v rune)
```
Format implements the fmt.Formatter interface

#### type WithStatementContents

```go
type WithStatementContents struct {
	Items  []WithItem
	Tokens Tokens
}
```

WithStatementContents as defined in python:3.13.0:
https://docs.python.org/3.13.0/reference/compound_stmts.html#grammar-token-python-grammar-with_stmt_contents

#### func (WithStatementContents) Format

```go
func (f WithStatementContents) Format(s fmt.State, v rune)
```
Format implements the fmt.Formatter interface

#### type XorExpression

```go
type XorExpression struct {
	AndExpression AndExpression
	XorExpression *XorExpression
	Tokens        Tokens
}
```

XorExpression as defined in python@3.13.0:
https://docs.python.org/release/3.13.0/reference/expressions.html#grammar-token-python-grammar-xor_expr

#### func (XorExpression) Format

```go
func (f XorExpression) Format(s fmt.State, v rune)
```
Format implements the fmt.Formatter interface

#### type YieldExpression

```go
type YieldExpression struct {
	ExpressionList *ExpressionList
	From           *Expression
	Tokens         Tokens
}
```

YieldExpression as defined in python@3.13.0:
https://docs.python.org/release/3.13.0/reference/simple_stmts.html#grammar-token-python-grammar-yield_stmt

#### func (YieldExpression) Format

```go
func (f YieldExpression) Format(s fmt.State, v rune)
```
Format implements the fmt.Formatter interface

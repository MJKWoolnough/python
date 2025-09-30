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

Used if you want to manually tokenise python source code.

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
	Comments           [2]Comments
	Tokens             Tokens
}
```

AddExpression as defined in python@3.13.0:
https://docs.python.org/release/3.13.0/reference/expressions.html#grammar-token-python-grammar-a_expr

When in a multiline structure, comments are parsed on either side of the '+' or
'-' tokens.

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
	Comments        [2]Comments
	Tokens          Tokens
}
```

AndExpression as defined in python@3.13.0:
https://docs.python.org/release/3.13.0/reference/expressions.html#grammar-token-python-grammar-and_expr

When in a multiline structure, comments are parsed on either side of the '&'
token.

#### func (AndExpression) Format

```go
func (f AndExpression) Format(s fmt.State, v rune)
```
Format implements the fmt.Formatter interface

#### type AndTest

```go
type AndTest struct {
	NotTest  NotTest
	AndTest  *AndTest
	Comments [2]Comments
	Tokens   Tokens
}
```

AndTest as defined in python@3.13.0:
https://docs.python.org/release/3.13.0/reference/expressions.html#grammar-token-python-grammar-and_test

When in a multiline structure, comments are parsed on either side of the 'and'
operator.

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
	Comments      [2]Comments
	Tokens        Tokens
}
```

ArgumentListOrComprehension as defined in python@3.13.0:
https://docs.python.org/release/3.13.0/reference/expressions.html#grammar-token-python-grammar-call

The first set of comments are parsed from directly after the opening paren.

The second set of comments are parsed from directly before the closing paren.

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
	Comments   [2]Comments
	Tokens
}
```

AssignmentExpression as defined in python@3.13.0:
https://docs.python.org/release/3.13.0/reference/expressions.html#grammar-token-python-grammar-assignment_expression

When in a multiline structure, comments are parsed on either side of the ':='
operator.

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
IsIdentifier returns true if the Atom contains an Identifier.

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
	Comments    [3]Comments
	Tokens      Tokens
}
```

ClassDefinition as defined in python@3.13.0:
https://docs.python.org/release/3.13.0/reference/compound_stmts.html#grammar-token-python-grammar-classdef

The first set of comments are parsed after any Decorators, before the 'class'
keyword.

The second and third set of comments are parsed inside of an empty Inheritance
list.

#### func (ClassDefinition) Format

```go
func (f ClassDefinition) Format(s fmt.State, v rune)
```
Format implements the fmt.Formatter interface

#### type Comments

```go
type Comments []Token
```


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
	Comments           [3]Comments
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
https://docs.python.org/release/3.13.0/reference/compound_stmts.html#grammar-token-python-grammar-compound_stmt

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
	Comments             [3]Comments
	Tokens               Tokens
}
```

Comprehension as defined in python@3.13.0:
https://docs.python.org/release/3.13.0/reference/expressions.html#grammar-token-python-grammar-comprehension

The first set of comments are parsed from before the Comprehension.

The second set of comments are parsed from between the AssignmentExpression and
he ComprehensionFor.

The final set of comments are parsed from after the Comprehension.

NB: Comments are only parsed when in a multiline structure.

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
	Comments              [2]Comments
	Tokens                Tokens
}
```

ComprehensionFor as defined in python@3.13.0:
https://docs.python.org/release/3.13.0/reference/expressions.html#grammar-token-python-grammar-comp_for

The first set of comments are parsed after an 'async' token.

The second set of comments are parsed from after the 'in' token.

NB: Comments are only parsed when in a multiline structure.

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
	Comments              Comments
	Tokens                Tokens
}
```

ComprehensionIf as defined in python@3.13.0:
https://docs.python.org/release/3.13.0/reference/expressions.html#grammar-token-python-grammar-comp_if

When in a multiline structure, the comments are parsed from after the 'if'
keyword.

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
	Comments         [2]Comments
	Tokens           Tokens
}
```

ComprehensionIterator as defined in python@3.13.0:
https://docs.python.org/release/3.13.0/reference/expressions.html#grammar-token-python-grammar-comp_iter

When in a multiline structure, the comments are parsed from before and after the
ComprehensionIterator.

#### func (ComprehensionIterator) Format

```go
func (f ComprehensionIterator) Format(s fmt.State, v rune)
```
Format implements the fmt.Formatter interface

#### type ConditionalExpression

```go
type ConditionalExpression struct {
	OrTest   OrTest
	If       *OrTest
	Else     *Expression
	Comments [4]Comments
	Tokens   Tokens
}
```

ConditionalExpression as defined in python@3.13.0:
https://docs.python.org/release/3.13.0/reference/expressions.html#grammar-token-python-grammar-conditional_expression

The first and second sets of comments are parsed from before and after an 'if'
keyword.

The third and fourth sets of comments are parsed from before and after an 'else'
keyword.

NB: Comments are only parsed when in a multiline structure.

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

#### type Decorator

```go
type Decorator struct {
	Decorator AssignmentExpression
	Comments  [2]Comments
	Tokens    Tokens
}
```

Decorator as defined in python@3.13.0:
https://docs.python.org/release/3.13.0/reference/compound_stmts.html#grammar-token-python-grammar-decorator

The first set of comments are parsed from before the decorator, and the second
set are parsed from after, on the same line.

#### func (Decorator) Format

```go
func (f Decorator) Format(s fmt.State, v rune)
```
Format implements the fmt.Formatter interface

#### type Decorators

```go
type Decorators struct {
	Decorators []Decorator
	Tokens
}
```

Decorators as defined in python@3.13.0:
https://docs.python.org/release/3.13.0/reference/compound_stmts.html#grammar-token-python-grammar-decorators

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
	Comments  [4]Comments
	Tokens    Tokens
}
```

DefParameter as defined in python@3.13.0:
https://docs.python.org/release/3.13.0/reference/compound_stmts.html#grammar-token-python-grammar-defparameter

The first set of comments are parsed from before the DefParameter.

The second and thrid set of comments are parsed from either side of a '=' token.

The final set of comments are parsed from after the DefParameter.

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
	Comments     [4]Comments
	Tokens       Tokens
}
```

DictItem as defined in python@3.13.0:
https://docs.python.org/release/3.13.0/reference/expressions.html#grammar-token-python-grammar-dict_item

The first set of comments are parsed from before the DictItem.

In a key/value DictItem, the second and third comments are parsed from before
and after the ':' token; otherwise, the second comments are parsed from after
the '**' token.

The final set of comments are parsed from after the DictItem.

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
	Comments            [2]Comments
	Tokens              Tokens
}
```

Enclosure as defined in python@3.13.0:
https://docs.python.org/release/3.13.0/reference/expressions.html#grammar-token-python-grammar-enclosure

The first set of comments are parsed from directly after the opening paren,
brace, or bracket; the second set of comments are parsed from directly before
the closing paren, brace, or bracket.

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
https://docs.python.org/release/3.13.0/reference/compound_stmts.html#grammar-token-python-grammar-try1_stmt

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
	Comments   [2]Comments
	Tokens     Tokens
}
```

File represents a parsed Python file.

The first set of comments are parsed and printed at the top of the file, the
second set are parsed and printed at the bottom of the file.

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
	StarredExpression    *OrExpression
	Comments             [2]Comments
	Tokens
}
```

FlexibleExpression as defined in python@3.13.0:
https://docs.python.org/release/3.13.0/reference/expressions.html#grammar-token-python-grammar-flexible_expression

The first set of comments are parsed from before the FlexibleExpression; the
second set of comments are parsed from after the FlexibleExpression.

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
https://docs.python.org/release/3.13.0/reference/compound_stmts.html#grammar-token-python-grammar-for_stmt

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
	Comments      [3]Comments
	Tokens        Tokens
}
```

FuncDefinition as defined in python@3.13.0:
https://docs.python.org/release/3.13.0/reference/compound_stmts.html#grammar-token-python-grammar-funcdef

The first set of comments are parsed after any Decorators, before the 'async' of
'def' keywords.

The second and third set of comments are parsed inside of an empty parameter
list.

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
	Comments         [3]Comments
	Tokens           Tokens
}
```

GeneratorExpression as defined in python@3.13.0:
https://docs.python.org/release/3.13.0/reference/expressions.html#grammar-token-python-grammar-generator_expression

The first set of comments are parsed from before the GeneratorExpression.

The second set of comments are parsed after the expression.

The third set of comments are parsed from after the GeneratorExpression.

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
https://docs.python.org/release/3.13.0/reference/compound_stmts.html#grammar-token-python-grammar-if_stmt

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
	Comments    [3]Comments
	Tokens      Tokens
}
```

KeywordArgument as defined in python@3.13.0:
https://docs.python.org/release/3.13.0/reference/expressions.html#grammar-token-python-grammar-keywords_arguments

The first set of comments are parsed from before the KeywordArgument.

The second set of comments are parsed from after any '**' token.

The final set of comments are parsed from after the KeywordArgument.

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
	Comments      [3]Comments
	Tokens        Tokens
}
```

LambdaExpression as defined in python@3.13.0:
https://docs.python.org/release/3.13.0/reference/expressions.html#grammar-token-python-grammar-lambda_expr

The first set of comments are parsed after the 'lambda' keyword.

If there are params, the second set of comments are parsed before the ':' token.

The third set of comments are parsed after the ':' token.

NB: Comments are only parsed when in a multiline structure.

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
	Comments           [2]Comments
	Tokens             Tokens
}
```

MultiplyExpression as defined in python@3.13.0:
https://docs.python.org/release/3.13.0/reference/expressions.html#grammar-token-python-grammar-m_expr

When in a multiline structure, comments are parsed on either side of the
operator token.

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
	Nots       []Comments
	Comparison Comparison
	Tokens     Tokens
}
```

NotTest as defined in python@3.13.0:
https://docs.python.org/release/3.13.0/reference/expressions.html#grammar-token-python-grammar-not_test

When in a multiline structure, comments are parsed after every 'not' operator

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
	Comments      [2]Comments
	Tokens        Tokens
}
```

OrExpression as defined in python@3.13.0:
https://docs.python.org/release/3.13.0/reference/expressions.html#grammar-token-python-grammar-or_expr

When in a multiline structure, comments are parsed on either side of the '|'
token.

#### func (OrExpression) Format

```go
func (f OrExpression) Format(s fmt.State, v rune)
```
Format implements the fmt.Formatter interface

#### type OrTest

```go
type OrTest struct {
	AndTest  AndTest
	OrTest   *OrTest
	Comments [2]Comments
	Tokens   Tokens
}
```

OrTest as defined in python@3.13.0:
https://docs.python.org/release/3.13.0/reference/expressions.html#grammar-token-python-grammar-or_test

When in a multiline structure, comments are parsed on either side of the 'or'
operator.

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
	Comments   [2]Comments
	Tokens     Tokens
}
```

Parameter as defined in python@3.13.0:
https://docs.python.org/release/3.13.0/reference/compound_stmts.html#grammar-token-python-grammar-parameter

The comments are parsed on either side of the ':' token.

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
	Comments      [10]Comments
	Tokens        Tokens
}
```

ParameterList as defined in python@3.13.0:
https://docs.python.org/release/3.13.0/reference/compound_stmts.html#grammar-token-python-grammar-parameter_list

The first set of comments are parsed from the beginning of the list.

The second and third sets of comments are parsed from before and after the '/'
token.

The fourth and fifth sets of comments are parsed from before and after the '*'
token; the sixth set of comments is parsed from after the StarArg.

The seventh and eighth set of comments are parsed from before and after the '**'
token; the ninth set of comments are parse from after the StarStarArg.

The final set of comments are parsed from after the ParameterList.

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
	Comments             [3]Comments
	Tokens               Tokens
}
```

PositionalArgument as defined in python@3.13.0:
https://docs.python.org/release/3.13.0/reference/expressions.html#grammar-token-python-grammar-positional_arguments

The first set of comments are parsed from before the PositionalArgument.

The second set of comments are parsed from after any '*' token.

The final set of comments are parsed from after the PositionalArgument.

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
	Comments          [3]Comments
	Tokens            Tokens
}
```

PowerExpression as defined in python@3.13.0:
https://docs.python.org/release/3.13.0/reference/expressions.html#grammar-token-python-grammar-power

The first set of comments are parsed after an 'await' keyword.

The second and third set of comments are parsed from before and after the '**'
token.

NB: Comments are only parsed when in a multiline structure.

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
	Comments          [2]Comments
	Tokens            Tokens
}
```

PrimaryExpression as defined in python@3.13.0:
https://docs.python.org/release/3.13.0/reference/expressions.html#grammar-token-python-grammar-primary

For an AttributeRef, the comments are parsed from before and after the '.'
token.

For a Slice or a Call, the comments are parsed from before the opening '[' or
'('.

NB: Comments are only parsed when in a multiline structure.

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
	Comments        [2]Comments
	Tokens          Tokens
}
```

ShiftExpression as defined in python@3.13.0:
https://docs.python.org/release/3.13.0/reference/expressions.html#grammar-token-python-grammar-shift_expr

When in a multiline structure, comments are parsed on either side of the '<<' or
'>>' tokens.

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
	Comments   [6]Comments
	Tokens     Tokens
}
```

SliceItem as defined in python@3.13.0:
https://docs.python.org/release/3.13.0/reference/expressions.html#grammar-token-python-grammar-slice_item

The first set of comments are parsed from directly after the opening bracket.

The second and third set of comments are parsed from either side of the first
':' token.

The fourth and fifth set of comments are parsed from either side of the second
':' token, if it exists.

The final set of comments are parsed from directly before the closing bracket.

#### func (SliceItem) Format

```go
func (f SliceItem) Format(s fmt.State, v rune)
```
Format implements the fmt.Formatter interface

#### type SliceList

```go
type SliceList struct {
	SliceItems []SliceItem
	Comments   [2]Comments
	Tokens     Tokens
}
```

SliceList as defined in python@3.13.0:
https://docs.python.org/release/3.13.0/reference/expressions.html#grammar-token-python-grammar-slice_list

The first set of comments are parsed from directly after the opening brace.

The second set of comments are parsed from directly before the closing brace.

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
	Comments    [2]Comments
	Tokens      Tokens
}
```

StarredExpression as defined in python@3.12.6:
https://docs.python.org/release/3.12.6/reference/expressions.html#grammar-token-python-grammar-starred_expression

When in a multiline structure, the comments are parsed before and after the
StarredExpression.

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
	Comments             [3]Comments
	Tokens               Tokens
}
```

StarredItem as defined in python@3.12.6:
https://docs.python.org/release/3.12.6/reference/expressions.html#grammar-token-python-grammar-starred_item

The first set of comments are parsed from before the StarredItem.

The second set of comments are parsed from after any '*' token.

The final set of comments are parsed from after the StarredItem.

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
	Comments    [3]Comments
	Tokens      Tokens
}
```

StarredOrKeyword as defined in python@3.13.0:
https://docs.python.org/release/3.13.0/reference/expressions.html#grammar-token-python-grammar-starred_and_keywords

The first set of comments are parsed from before the StarredOrKeyword item.

The second set of comments are parsed from after any '*' token.

The final set of comments are parsed from after the StarredOrKeyword item.

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
	Comments          Comments
	Tokens            Tokens
}
```

Statement as defined in python@3.13.0:
https://docs.python.org/release/3.13.0/reference/compound_stmts.html#grammar-token-python-grammar-statement

The comments are parsed from before the Statement.

#### func (Statement) Format

```go
func (f Statement) Format(s fmt.State, v rune)
```
Format implements the fmt.Formatter interface

#### type StatementList

```go
type StatementList struct {
	Statements []SimpleStatement
	Comments   Comments
	Tokens
}
```

StatementList as defined in python@3.13.0:
https://docs.python.org/release/3.13.0/reference/compound_stmts.html#grammar-token-python-grammar-stmt_list

The comments are parsed from after the StatementList.

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
	Comments      [2]Comments
	Tokens        Tokens
}
```

Suite as defined in python@3.13.0:
https://docs.python.org/release/3.13.0/reference/compound_stmts.html#grammar-token-python-grammar-suite

In a multiline Suite, the comments are parsed at the top and bottom.

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
	Comments          [2]Comments
	Tokens            Tokens
}
```

Target as defined in python@3.13.0:
https://docs.python.org/release/3.13.0/reference/simple_stmts.html#grammar-token-python-grammar-target

When in a multiline stucture, the comments are parsed from before and after the
Target.

#### func (Target) Format

```go
func (f Target) Format(s fmt.State, v rune)
```
Format implements the fmt.Formatter interface

#### type TargetList

```go
type TargetList struct {
	Targets  []Target
	Comments [2]Comments
	Tokens   Tokens
}
```

TargetList as defined in python@3.13.0:
https://docs.python.org/release/3.13.0/reference/simple_stmts.html#grammar-token-python-grammar-target_list

When in a multiline stucture, the comments are parsed from before and after the
TargetList.

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

Tokens represents a list of tokens that have been parsed.

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
https://docs.python.org/release/3.13.0/reference/compound_stmts.html#grammar-token-python-grammar-try_stmt

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
	Comments   [4]Comments
	Tokens     Tokens
}
```

TypeParam as defined in python@3.13.0:
https://docs.python.org/release/3.13.0/reference/compound_stmts.html#grammar-token-python-grammar-type_param

The first set of comments are parsed at the start of the type param.

When the type is TypeParamVar or TypeParamVarTuple, the second set of comments
are parsed after the '*' or '**' token.

When the type is TypeParamIdentifer, the second and third set of comments are
parsed on either side of the ':' token.

The last set of comments are parsed after the TypeParam.

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
	Comments   [2]Comments
	Tokens     Tokens
}
```

TypeParams as defined in python@3.13.0:
https://docs.python.org/release/3.13.0/reference/compound_stmts.html#grammar-token-python-grammar-type_params

The first set of comments are parsed from directly after the opening bracket;
the second set of comments are parsed from directly before the closing bracket.

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
	Comments
	Tokens Tokens
}
```

UnaryExpression as defined in python@3.13.0:
https://docs.python.org/release/3.13.0/reference/expressions.html#grammar-token-python-grammar-u_expr

When in a multiline structure, comments are parsed on either side of the
operator token.

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
https://docs.python.org/release/3.13.0/reference/compound_stmts.html#grammar-token-python-grammar-while_stmt

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
	Comments   [2]Comments
	Tokens     Tokens
}
```

WithItem as defined in python@3.13:
https://docs.python.org/release/3.13/reference/compound_stmts.html#grammar-token-python-grammar-with_item

If in a multiline context, the comments are parsed from either side of the
Expression.

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
https://docs.python.org/release/3.13.0/reference/compound_stmts.html#grammar-token-python-grammar-with_stmt

#### func (WithStatement) Format

```go
func (f WithStatement) Format(s fmt.State, v rune)
```
Format implements the fmt.Formatter interface

#### type WithStatementContents

```go
type WithStatementContents struct {
	Items    []WithItem
	Comments [2]Comments
	Tokens   Tokens
}
```

WithStatementContents as defined in python:3.13.0:
https://docs.python.org/release/3.13.0/reference/compound_stmts.html#grammar-token-python-grammar-with_stmt_contents

If in a multiline construct, the first set of comments are parsed from before
the contents, the second set are parsed from after.

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
	Comments      [2]Comments
	Tokens        Tokens
}
```

XorExpression as defined in python@3.13.0:
https://docs.python.org/release/3.13.0/reference/expressions.html#grammar-token-python-grammar-xor_expr

When in a multiline structure, comments are parsed on either side of the '^'
token.

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
	Comments       [4]Comments
	Tokens         Tokens
}
```

YieldExpression as defined in python@3.13.0:
https://docs.python.org/release/3.13.0/reference/simple_stmts.html#grammar-token-python-grammar-yield_stmt

The first and second sets of comments are parsed from before and after the
'yield' keyword.

The third set of comments are parsed from after the 'from' keyword, or after the
ExpressionList if it's followed by a trailing comma.

The final set of comments are parsed from directly after the YieldExpression.

#### func (YieldExpression) Format

```go
func (f YieldExpression) Format(s fmt.State, v rune)
```
Format implements the fmt.Formatter interface

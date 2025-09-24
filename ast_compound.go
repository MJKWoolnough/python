package python

import (
	"slices"

	"vimagination.zapto.org/parser"
)

var compounds = [...]string{"if", "while", "for", "try", "with", "func", "class", "async", "def"}

// CompoundStatement as defined in python@3.13.0:
// https://docs.python.org/release/3.13.0/reference/compound_stmts.html#grammar-token-python-grammar-compound_stmt
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

func (c *CompoundStatement) parse(p *pyParser) error {
	var err error

	q := p.NewGoal()

	if tk := q.Peek(); tk == (parser.Token{Type: TokenOperator, Data: "@"}) {
		r := q.NewGoal()

		skipDecorators(r)

		switch tk := r.Peek(); tk.Data {
		case "def":
			c.Func = new(FuncDefinition)
			err = c.Func.parse(q)
		case "class":
			c.Class = new(ClassDefinition)
			err = c.Class.parse(q)
		case "async":
			r.Next()
			r.AcceptRunWhitespace()

			switch tk := r.Peek(); tk.Data {
			case "def":
				c.Func = new(FuncDefinition)
				err = c.Func.parse(q)
			default:
				err = ErrInvalidCompound
			}
		default:
			err = ErrInvalidCompound
		}
	} else {
		switch tk := q.Peek(); tk.Data {
		case "if":
			c.If = new(IfStatement)
			err = c.If.parse(q)
		case "while":
			c.While = new(WhileStatement)
			err = c.While.parse(q)
		case "for":
			c.For = new(ForStatement)
			err = c.For.parse(q)
		case "try":
			c.Try = new(TryStatement)
			err = c.Try.parse(q)
		case "with":
			c.With = new(WithStatement)
			err = c.With.parse(q)
		case "def":
			c.Func = new(FuncDefinition)
			err = c.Func.parse(q)
		case "class":
			c.Class = new(ClassDefinition)
			err = c.Class.parse(q)
		case "async":
			r := q.NewGoal()

			r.Next()
			r.AcceptRunWhitespace()

			switch tk := r.Peek(); tk.Data {
			case "for":
				c.For = new(ForStatement)
				err = c.For.parse(q)
			case "with":
				c.With = new(WithStatement)
				err = c.With.parse(q)
			case "def":
				c.Func = new(FuncDefinition)
				err = c.Func.parse(q)
			default:
				err = ErrInvalidCompound
			}
		default:
			err = ErrInvalidCompound
		}
	}

	if err != nil {
		return p.Error("CompoundStatement", err)
	}

	p.Score(q)

	c.Tokens = p.ToTokens()

	return nil
}

// Decorators as defined in python@3.13.0:
// https://docs.python.org/release/3.13.0/reference/compound_stmts.html#grammar-token-python-grammar-decorators
type Decorators struct {
	Decorators []Decorator
	Tokens
}

func (d *Decorators) parse(p *pyParser) error {
	q := p.NewGoal()

	q.AcceptRunWhitespace()

	for q.AcceptToken(parser.Token{Type: TokenOperator, Data: "@"}) {
		p.AcceptRunAllWhitespaceNoComment()

		q = p.NewGoal()

		var dc Decorator

		if err := dc.parse(q); err != nil {
			return p.Error("Decorators", err)
		}

		d.Decorators = append(d.Decorators, dc)

		p.Score(q)

		q = p.NewGoal()

		q.AcceptRunWhitespace()

		if !q.Accept(TokenLineTerminator) {
			return p.Error("Decorators", ErrMissingNewline)
		}

		q.AcceptRunAllWhitespace()
	}

	d.Tokens = p.ToTokens()

	return nil
}

func skipDecorators(p *pyParser) {
	for p.AcceptToken(parser.Token{Type: TokenOperator, Data: "@"}) {
		p.AcceptRunWhitespace()
		skipAssignmentExpression(p)
		p.AcceptRunAllWhitespace()
	}
}

// Decorator as defined in python@3.13.0:
// https://docs.python.org/release/3.13.0/reference/compound_stmts.html#grammar-token-python-grammar-decorator
//
// The first set of comments are parsed from before the decorator, and the
// second set are parsed from after, on the same line.
type Decorator struct {
	Decorator AssignmentExpression
	Comments  [2]Comments
	Tokens    Tokens
}

func (d *Decorator) parse(p *pyParser) error {
	d.Comments[0] = p.AcceptRunWhitespaceComments()

	p.AcceptRunAllWhitespace()
	p.Next()
	p.AcceptRunWhitespace()

	q := p.NewGoal()

	if err := d.Decorator.parse(q); err != nil {
		return p.Error("Decorator", err)
	}

	p.Score(q)

	d.Comments[1] = p.AcceptRunWhitespaceCommentsNoNewline()
	d.Tokens = p.ToTokens()

	return nil
}

// IfStatement as defined in python@3.13.0:
// https://docs.python.org/release/3.13.0/reference/compound_stmts.html#grammar-token-python-grammar-if_stmt
type IfStatement struct {
	AssignmentExpression AssignmentExpression
	Suite                Suite
	Elif                 []AssignmentExpressionAndSuite
	Else                 *Suite
	Tokens               Tokens
}

func (i *IfStatement) parse(p *pyParser) error {
	p.AcceptToken(parser.Token{Type: TokenKeyword, Data: "if"})
	p.AcceptRunWhitespace()

	q := p.NewGoal()

	if err := i.AssignmentExpression.parse(q); err != nil {
		return p.Error("IfStatement", err)
	}

	p.Score(q)
	p.AcceptRunWhitespace()

	if !p.AcceptToken(parser.Token{Type: TokenDelimiter, Data: ":"}) {
		return p.Error("IfStatement", ErrMissingColon)
	}

	p.AcceptRunWhitespaceNoComment()

	q = p.NewGoal()

	if err := i.Suite.parse(q); err != nil {
		return p.Error("IfStatement", err)
	}

	p.Score(q)

	q = p.NewGoal()

	q.AcceptRunAllWhitespace()

	for q.AcceptToken(parser.Token{Type: TokenKeyword, Data: "elif"}) {
		q.AcceptRunWhitespace()
		p.Score(q)

		q = p.NewGoal()

		var as AssignmentExpressionAndSuite

		if err := as.AssignmentExpression.parse(q); err != nil {
			return p.Error("IfStatement", err)
		}

		p.Score(q)
		p.AcceptRunWhitespace()

		if !p.AcceptToken(parser.Token{Type: TokenDelimiter, Data: ":"}) {
			return p.Error("IfStatement", ErrMissingColon)
		}

		p.AcceptRunWhitespaceNoComment()

		q = p.NewGoal()

		if err := as.Suite.parse(q); err != nil {
			return p.Error("IfStatement", err)
		}

		p.Score(q)

		i.Elif = append(i.Elif, as)
		q = p.NewGoal()

		q.AcceptRunAllWhitespace()
	}

	q = p.NewGoal()

	q.AcceptRun(TokenLineTerminator)

	if q.AcceptToken(parser.Token{Type: TokenKeyword, Data: "else"}) {
		p.Score(q)
		p.AcceptRunWhitespace()

		if !p.AcceptToken(parser.Token{Type: TokenDelimiter, Data: ":"}) {
			return p.Error("IfStatement", ErrMissingColon)
		}

		p.AcceptRunWhitespaceNoComment()

		i.Else = new(Suite)
		q = p.NewGoal()

		if err := i.Else.parse(q); err != nil {
			return p.Error("IfStatement", err)
		}

		p.Score(q)
	}

	i.Tokens = p.ToTokens()

	return nil
}

// AssignmentExpressionAndSuite is a combination of the AssignmentExpression and Suite types.
type AssignmentExpressionAndSuite struct {
	AssignmentExpression AssignmentExpression
	Suite                Suite
}

// WhileStatement as defined in python@3.13.0:
// https://docs.python.org/release/3.13.0/reference/compound_stmts.html#grammar-token-python-grammar-while_stmt
type WhileStatement struct {
	AssignmentExpression AssignmentExpression
	Suite                Suite
	Else                 *Suite
	Tokens               Tokens
}

func (w *WhileStatement) parse(p *pyParser) error {
	p.AcceptToken(parser.Token{Type: TokenKeyword, Data: "while"})
	p.AcceptRunWhitespace()

	q := p.NewGoal()

	if err := w.AssignmentExpression.parse(q); err != nil {
		return p.Error("WhileStatement", err)
	}

	p.Score(q)
	p.AcceptRunWhitespace()

	if !p.AcceptToken(parser.Token{Type: TokenDelimiter, Data: ":"}) {
		return p.Error("WhileStatement", ErrMissingColon)
	}

	p.AcceptRunWhitespaceNoComment()

	q = p.NewGoal()

	if err := w.Suite.parse(q); err != nil {
		return p.Error("WhileStatement", err)
	}

	p.Score(q)

	q = p.NewGoal()

	q.AcceptRunAllWhitespace()

	if q.AcceptToken(parser.Token{Type: TokenKeyword, Data: "else"}) {
		p.Score(q)
		p.AcceptRunWhitespace()

		if !p.AcceptToken(parser.Token{Type: TokenDelimiter, Data: ":"}) {
			return p.Error("WhileStatement", ErrMissingColon)
		}

		p.AcceptRunWhitespaceNoComment()

		w.Else = new(Suite)
		q = p.NewGoal()

		if err := w.Else.parse(q); err != nil {
			return p.Error("WhileStatement", err)
		}

		p.Score(q)
	}

	w.Tokens = p.ToTokens()

	return nil
}

// ForStatement as defined in python@3.13.0:
// https://docs.python.org/release/3.13.0/reference/compound_stmts.html#grammar-token-python-grammar-for_stmt
type ForStatement struct {
	Async       bool
	TargetList  TargetList
	StarredList StarredList
	Suite       Suite
	Else        *Suite
	Tokens
}

func (f *ForStatement) parse(p *pyParser) error {
	if f.Async = p.AcceptToken(parser.Token{Type: TokenKeyword, Data: "async"}); f.Async {
		p.AcceptRunWhitespace()
	}

	if p.AcceptToken(parser.Token{Type: TokenKeyword, Data: "for"}) {
		p.AcceptRunWhitespace()
	}

	q := p.NewGoal()

	if err := f.TargetList.parse(q); err != nil {
		return p.Error("ForStatement", err)
	}

	p.Score(q)
	p.AcceptRunWhitespace()

	if !p.AcceptToken(parser.Token{Type: TokenKeyword, Data: "in"}) {
		return p.Error("ForStatement", ErrMissingIn)
	}

	p.AcceptRunWhitespace()

	q = p.NewGoal()

	if err := f.StarredList.parse(q); err != nil {
		return p.Error("ForStatement", err)
	}

	p.Score(q)
	p.AcceptRunWhitespace()

	if !p.AcceptToken(parser.Token{Type: TokenDelimiter, Data: ":"}) {
		return p.Error("ForStatement", ErrMissingColon)
	}

	p.AcceptRunWhitespaceNoComment()

	q = p.NewGoal()

	if err := f.Suite.parse(q); err != nil {
		return p.Error("ForStatement", err)
	}

	p.Score(q)

	q = p.NewGoal()

	q.AcceptRunAllWhitespace()

	if q.AcceptToken(parser.Token{Type: TokenKeyword, Data: "else"}) {
		p.Score(q)
		p.AcceptRunWhitespace()

		if !p.AcceptToken(parser.Token{Type: TokenDelimiter, Data: ":"}) {
			return p.Error("ForStatement", ErrMissingColon)
		}

		p.AcceptRunWhitespaceNoComment()

		f.Else = new(Suite)
		q = p.NewGoal()

		if err := f.Else.parse(q); err != nil {
			return p.Error("ForStatement", err)
		}

		p.Score(q)
	}

	f.Tokens = p.ToTokens()

	return nil
}

// TryStatement as defined in python@3.13.0:
// https://docs.python.org/release/3.13.0/reference/compound_stmts.html#grammar-token-python-grammar-try_stmt
type TryStatement struct {
	Try     Suite
	Groups  bool
	Except  []Except
	Else    *Suite
	Finally *Suite
	Tokens  Tokens
}

func (t *TryStatement) parse(p *pyParser) error {
	p.AcceptToken(parser.Token{Type: TokenKeyword, Data: "try"})
	p.AcceptRunWhitespace()

	if !p.AcceptToken(parser.Token{Type: TokenDelimiter, Data: ":"}) {
		return p.Error("TryStatement", ErrMissingColon)
	}

	p.AcceptRunWhitespaceNoComment()

	q := p.NewGoal()

	if err := t.Try.parse(q); err != nil {
		return p.Error("TryStatement", err)
	}

	p.Score(q)

	q = p.NewGoal()

	q.AcceptRunAllWhitespace()

	for q.AcceptToken(parser.Token{Type: TokenKeyword, Data: "except"}) {
		p.Score(q)
		p.AcceptRunWhitespace()

		group := p.AcceptToken(parser.Token{Type: TokenOperator, Data: "*"})

		if len(t.Except) > 0 && t.Groups != group {
			return p.Error("TryStatement", ErrMismatchedGroups)
		}

		p.AcceptRunWhitespace()

		t.Groups = group
		q = p.NewGoal()

		var except Except

		if err := except.parse(q); err != nil {
			return p.Error("TryStatement", err)
		}

		t.Except = append(t.Except, except)

		p.Score(q)

		q = p.NewGoal()

		q.AcceptRunAllWhitespace()
	}

	q = p.NewGoal()

	q.AcceptRunAllWhitespace()

	if len(t.Except) > 0 && q.AcceptToken(parser.Token{Type: TokenKeyword, Data: "else"}) {
		p.Score(q)
		p.AcceptRunWhitespace()

		if !p.AcceptToken(parser.Token{Type: TokenDelimiter, Data: ":"}) {
			return p.Error("TryStatement", ErrMissingColon)
		}

		p.AcceptRunWhitespaceNoComment()

		t.Else = new(Suite)
		q := p.NewGoal()

		if err := t.Else.parse(q); err != nil {
			return p.Error("TryStatement", err)
		}

		p.Score(q)
	}

	q = p.NewGoal()

	q.AcceptRunAllWhitespace()

	if q.AcceptToken(parser.Token{Type: TokenKeyword, Data: "finally"}) {
		p.Score(q)
		p.AcceptRunWhitespace()

		if !p.AcceptToken(parser.Token{Type: TokenDelimiter, Data: ":"}) {
			return p.Error("TryStatement", ErrMissingColon)
		}

		p.AcceptRunWhitespaceNoComment()

		t.Finally = new(Suite)
		q := p.NewGoal()

		if err := t.Finally.parse(q); err != nil {
			return p.Error("TryStatement", err)
		}

		p.Score(q)
	} else if len(t.Except) == 0 {
		return p.Error("TryStatement", ErrMissingFinally)
	}

	t.Tokens = p.ToTokens()

	return nil
}

// Except as defined in python@3.13.0:
// https://docs.python.org/release/3.13.0/reference/compound_stmts.html#grammar-token-python-grammar-try1_stmt
type Except struct {
	Expression Expression
	Identifier *Token
	Suite      Suite
	Tokens     Tokens
}

func (e *Except) parse(p *pyParser) error {
	q := p.NewGoal()

	if err := e.Expression.parse(q); err != nil {
		return p.Error("Except", err)
	}

	p.Score(q)
	p.AcceptRunWhitespace()

	if p.AcceptToken(parser.Token{Type: TokenKeyword, Data: "as"}) {
		p.AcceptRunWhitespace()

		if !p.Accept(TokenIdentifier) {
			return p.Error("Except", ErrMissingIdentifier)
		}

		e.Identifier = p.GetLastToken()

		p.AcceptRunWhitespace()
	}

	if !p.AcceptToken(parser.Token{Type: TokenDelimiter, Data: ":"}) {
		return p.Error("Except", ErrMissingColon)
	}

	p.AcceptRunWhitespaceNoComment()

	q = p.NewGoal()

	if err := e.Suite.parse(q); err != nil {
		return p.Error("Except", err)
	}

	p.Score(q)

	e.Tokens = p.ToTokens()

	return nil
}

// WithStatement as defined in python@3.13.0:
// https://docs.python.org/release/3.13.0/reference/compound_stmts.html#grammar-token-python-grammar-with_stmt
type WithStatement struct {
	Async    bool
	Contents WithStatementContents
	Suite    Suite
	Tokens   Tokens
}

func (w *WithStatement) parse(p *pyParser) error {
	if w.Async = p.AcceptToken(parser.Token{Type: TokenKeyword, Data: "async"}); w.Async {
		p.AcceptRunWhitespace()
	}

	p.AcceptToken(parser.Token{Type: TokenKeyword, Data: "with"})
	p.AcceptRunWhitespace()

	parens := p.AcceptToken(parser.Token{Type: TokenDelimiter, Data: "("})

	if parens {
		p.AcceptRunWhitespaceNoComment()
	}

	q := p.NewGoal()

	if err := w.Contents.parse(q); err != nil {
		return p.Error("WithStatement", err)
	}

	p.Score(q)
	p.AcceptRunWhitespace()

	if parens {
		if p.AcceptToken(parser.Token{Type: TokenDelimiter, Data: ","}) {
			p.AcceptRunWhitespace()
		}

		if !p.AcceptToken(parser.Token{Type: TokenDelimiter, Data: ")"}) {
			return p.Error("WithStatement", ErrMissingClosingParen)
		}

		p.AcceptRunWhitespace()
	}

	if !p.AcceptToken(parser.Token{Type: TokenDelimiter, Data: ":"}) {
		return p.Error("WithStatement", ErrMissingColon)
	}

	p.AcceptRunWhitespaceNoComment()

	q = p.NewGoal()

	if err := w.Suite.parse(q); err != nil {
		return p.Error("WithStatement", err)
	}

	p.Score(q)

	w.Tokens = p.ToTokens()

	return nil
}

// WithStatementContents as defined in python:3.13.0:
// https://docs.python.org/release/3.13.0/reference/compound_stmts.html#grammar-token-python-grammar-with_stmt_contents
//
// If in a multiline construct, the first set of comments are parsed from
// before the contents, the second set are parsed from after.
type WithStatementContents struct {
	Items    []WithItem
	Comments [2]Comments
	Tokens   Tokens
}

func (w *WithStatementContents) parse(p *pyParser) error {
	w.Comments[0] = p.AcceptRunWhitespaceCommentsNoNewlineIfMultiline()

	another := true

	for another {
		p.AcceptRunWhitespaceNoComment()

		q := p.NewGoal()

		var wi WithItem

		if err := wi.parse(q); err != nil {
			return p.Error("WithStatementContents", err)
		}

		p.Score(q)

		w.Items = append(w.Items, wi)
		q = p.NewGoal()

		q.AcceptRunWhitespace()

		another = q.AcceptToken(parser.Token{Type: TokenDelimiter, Data: ","})

		r := q.NewGoal()

		r.AcceptRunWhitespace()

		if r.Peek() == (parser.Token{Type: TokenDelimiter, Data: ")"}) {
			break
		}

		if another {
			p.Score(q)
		}
	}

	w.Comments[1] = p.AcceptRunWhitespaceCommentsIfMultiline()
	w.Tokens = p.ToTokens()

	return nil
}

// WithItem as defined in python@3.13:
// https://docs.python.org/release/3.13/reference/compound_stmts.html#grammar-token-python-grammar-with_item
//
// If in a multiline context, the comments are parsed from either side of the Expression.
type WithItem struct {
	Expression Expression
	Target     *Target
	Comments   [2]Comments
	Tokens     Tokens
}

func (w *WithItem) parse(p *pyParser) error {
	w.Comments[0] = p.AcceptRunWhitespaceCommentsIfMultiline()

	p.AcceptRunWhitespace()

	q := p.NewGoal()

	if err := w.Expression.parse(q); err != nil {
		return p.Error("WithItem", err)
	}

	p.Score(q)

	q = p.NewGoal()

	q.AcceptRunWhitespace()

	if q.AcceptToken(parser.Token{Type: TokenKeyword, Data: "as"}) {
		w.Comments[1] = p.AcceptRunWhitespaceCommentsIfMultiline()

		p.AcceptRunWhitespace()
		p.Next()

		p.AcceptRunWhitespaceNoComment()

		q = p.NewGoal()
		w.Target = new(Target)

		if err := w.Target.parse(q); err != nil {
			return p.Error("WithItem", err)
		}

		p.Score(q)
	} else if q.AcceptToken(parser.Token{Type: TokenDelimiter, Data: ","}) {
		w.Comments[1] = p.AcceptRunWhitespaceCommentsIfMultiline()
	} else {
		w.Comments[1] = p.AcceptRunWhitespaceCommentsNoNewlineIfMultiline()
	}

	w.Tokens = p.ToTokens()

	return nil
}

// FuncDefinition as defined in python@3.13.0:
// https://docs.python.org/release/3.13.0/reference/compound_stmts.html#grammar-token-python-grammar-funcdef
//
// The first set of comments are parsed after any Decorators, before the 'async' of 'def' keywords.
//
// The second and third set of comments are parsed inside of an empty parameter list.
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

func (f *FuncDefinition) parse(p *pyParser) error {
	if p.Peek() == (parser.Token{Type: TokenOperator, Data: "@"}) {
		q := p.NewGoal()
		f.Decorators = new(Decorators)

		if err := f.Decorators.parse(q); err != nil {
			return p.Error("FuncDefinition", err)
		}

		p.Score(q)

		f.Comments[0] = p.AcceptRunWhitespaceComments()

		p.AcceptRunAllWhitespace()
	}

	if f.Async = p.AcceptToken(parser.Token{Type: TokenKeyword, Data: "async"}); f.Async {
		p.AcceptRunWhitespace()
	}

	p.AcceptToken(parser.Token{Type: TokenKeyword, Data: "def"})
	p.AcceptRunWhitespace()

	if !p.Accept(TokenIdentifier) {
		return p.Error("FuncDefinition", ErrMissingIdentifier)
	}

	f.FuncName = p.GetLastToken()

	p.AcceptRunWhitespace()

	if p.Peek() == (parser.Token{Type: TokenDelimiter, Data: "["}) {
		q := p.NewGoal()
		f.TypeParams = new(TypeParams)

		if err := f.TypeParams.parse(q); err != nil {
			return p.Error("FuncDefinition", err)
		}

		p.Score(q)
		p.AcceptRunWhitespace()
	}

	if !p.AcceptToken(parser.Token{Type: TokenDelimiter, Data: "("}) {
		return p.Error("FuncDefinition", ErrMissingOpeningParen)
	}

	q := p.NewGoal()

	q.AcceptRunAllWhitespace()

	if !q.AcceptToken(parser.Token{Type: TokenDelimiter, Data: ")"}) {
		q = p.NewGoal()

		if q.AcceptRunAllWhitespaceNoComment() == TokenComment {
			p.AcceptRunWhitespaceNoNewline()
		} else {
			p.Score(q)
		}

		q = p.NewGoal()

		if err := f.ParameterList.parse(q, true); err != nil {
			return p.Error("FuncDefinition", err)
		}

		p.Score(q)
		p.AcceptRunWhitespace()

		if !p.AcceptToken(parser.Token{Type: TokenDelimiter, Data: ")"}) {
			return p.Error("FuncDefinition", ErrMissingClosingParen)
		}

	} else {
		f.Comments[1] = p.AcceptRunWhitespaceCommentsNoNewline()
		f.Comments[2] = p.AcceptRunWhitespaceComments()

		p.AcceptRunAllWhitespace()
		p.AcceptToken(parser.Token{Type: TokenDelimiter, Data: ")"})

		f.ParameterList.Tokens = p.NewGoal().ToTokens()
	}

	p.AcceptRunWhitespace()

	if p.AcceptToken(parser.Token{Type: TokenDelimiter, Data: "->"}) {
		p.AcceptRunWhitespace()

		q := p.NewGoal()
		f.Expression = new(Expression)

		if err := f.Expression.parse(q); err != nil {
			return p.Error("FuncDefinition", err)
		}

		p.Score(q)
		p.AcceptRunWhitespace()
	}

	if !p.AcceptToken(parser.Token{Type: TokenDelimiter, Data: ":"}) {
		return p.Error("FuncDefinition", ErrMissingColon)
	}

	p.AcceptRunWhitespaceNoComment()

	q = p.NewGoal()

	if err := f.Suite.parse(q); err != nil {
		return p.Error("FuncDefinition", err)
	}

	p.Score(q)

	f.Tokens = p.ToTokens()

	return nil
}

// ClassDefinition as defined in python@3.13.0:
// https://docs.python.org/release/3.13.0/reference/compound_stmts.html#grammar-token-python-grammar-classdef
//
// The first set of comments are parsed after any Decorators, before the 'class' keyword.
//
// The second and third set of comments are parsed inside of an empty Inheritance list.
type ClassDefinition struct {
	Decorators  *Decorators
	ClassName   *Token
	TypeParams  *TypeParams
	Inheritance *ArgumentList
	Suite       Suite
	Comments    [3]Comments
	Tokens      Tokens
}

func (c *ClassDefinition) parse(p *pyParser) error {
	if p.Peek() == (parser.Token{Type: TokenOperator, Data: "@"}) {
		q := p.NewGoal()
		c.Decorators = new(Decorators)

		if err := c.Decorators.parse(q); err != nil {
			return p.Error("ClassDefinition", err)
		}

		p.Score(q)

		c.Comments[0] = p.AcceptRunWhitespaceComments()

		p.AcceptRunAllWhitespace()
	}

	p.AcceptToken(parser.Token{Type: TokenKeyword, Data: "class"})
	p.AcceptRunWhitespace()

	if !p.Accept(TokenIdentifier) {
		return p.Error("ClassDefinition", ErrMissingIdentifier)
	}

	c.ClassName = p.GetLastToken()

	p.AcceptRunWhitespace()

	if p.Peek() == (parser.Token{Type: TokenDelimiter, Data: "["}) {
		q := p.NewGoal()
		c.TypeParams = new(TypeParams)

		if err := c.TypeParams.parse(q); err != nil {
			return p.Error("ClassDefinition", err)
		}

		p.Score(q)
		p.AcceptRunWhitespace()
	}

	if p.AcceptToken(parser.Token{Type: TokenDelimiter, Data: "("}) {
		c.Comments[1] = p.AcceptRunWhitespaceCommentsNoNewline()

		p.AcceptRunAllWhitespaceNoComment()

		c.Inheritance = new(ArgumentList)

		q := p.NewGoal()

		q.AcceptRunWhitespace()

		if !q.AcceptToken(parser.Token{Type: TokenDelimiter, Data: ")"}) {
			q := p.NewGoal()

			if err := c.Inheritance.parse(q); err != nil {
				return p.Error("ClassDefinition", err)
			}

			p.Score(q)

		} else {
			c.Inheritance.Tokens = p.NewGoal().ToTokens()
		}

		c.Comments[2] = p.AcceptRunWhitespaceComments()

		p.AcceptRunWhitespace()
		p.AcceptToken(parser.Token{Type: TokenDelimiter, Data: ")"})
		p.AcceptRunWhitespace()
	}

	if !p.AcceptToken(parser.Token{Type: TokenDelimiter, Data: ":"}) {
		return p.Error("ClassDefinition", ErrMissingColon)
	}

	p.AcceptRunWhitespaceNoComment()

	q := p.NewGoal()

	if err := c.Suite.parse(q); err != nil {
		return p.Error("ClassDefinition", err)
	}

	p.Score(q)

	c.Tokens = p.ToTokens()

	return nil
}

// Suite as defined in python@3.13.0:
// https://docs.python.org/release/3.13.0/reference/compound_stmts.html#grammar-token-python-grammar-suite
//
// In a multiline Suite, the comments are parsed at the top and bottom.
type Suite struct {
	StatementList *StatementList
	Statements    []Statement
	Comments      [2]Comments
	Tokens        Tokens
}

func (s *Suite) parse(p *pyParser) error {
	if tk := p.Peek(); tk.Type == TokenLineTerminator || tk.Type == TokenComment {
		s.Comments[0] = p.AcceptRunWhitespaceComments()

		p.AcceptRunAllWhitespace()

		if !p.Accept(TokenIndent) {
			return p.Error("Suite", ErrMissingIndent)
		}

		for {
			p.AcceptRunAllWhitespaceNoComment()

			q := p.NewGoal()

			var stmt Statement

			if err := stmt.parse(q); err != nil {
				return p.Error("Suite", err)
			}

			s.Statements = append(s.Statements, stmt)

			p.Score(q)

			q = p.NewGoal()

			q.AcceptRunAllWhitespace()

			if q.Accept(TokenDedent) {
				s.Comments[1] = p.AcceptRunWhitespaceComments()

				p.AcceptRunAllWhitespace()
				p.Accept(TokenDedent)

				break
			}
		}
	} else {
		s.StatementList = new(StatementList)

		if err := s.StatementList.parse(p); err != nil {
			return p.Error("Suite", err)
		}
	}

	s.Tokens = p.ToTokens()

	return nil
}

// TypeParamType determines the type of a TypeParam.
type TypeParamType byte

const (
	TypeParamIdentifer TypeParamType = iota
	TypeParamVar
	TypeParamVarTuple
)

// TypeParam as defined in python@3.13.0:
// https://docs.python.org/release/3.13.0/reference/compound_stmts.html#grammar-token-python-grammar-type_param
//
// The first set of comments are parsed at the start of the type param.
//
// When the type is TypeParamVar or TypeParamVarTuple, the second set of
// comments are parsed after the '*' or '**' token.
//
// When the type is TypeParamIdentifer, the second and third set of comments
// are parsed on either side of the ':' token.
//
// The last set of comments are parsed after the TypeParam.
type TypeParam struct {
	Type       TypeParamType
	Identifier *Token
	Expression *Expression
	Comments   [4]Comments
	Tokens     Tokens
}

func (t *TypeParam) parse(p *pyParser) error {
	t.Comments[0] = p.AcceptRunWhitespaceComments()

	p.AcceptRunAllWhitespace()

	if p.AcceptToken(parser.Token{Type: TokenOperator, Data: "*"}) {
		t.Type = TypeParamVar
		t.Comments[1] = p.AcceptRunWhitespaceComments()

		p.AcceptRunWhitespace()
	} else if p.AcceptToken(parser.Token{Type: TokenOperator, Data: "**"}) {
		t.Type = TypeParamVarTuple
		t.Comments[1] = p.AcceptRunWhitespaceComments()

		p.AcceptRunWhitespace()
	}

	if !p.Accept(TokenIdentifier) {
		return p.Error("TypeParam", ErrMissingIdentifier)
	}

	t.Identifier = p.GetLastToken()

	if t.Type == TypeParamIdentifer {
		q := p.NewGoal()

		q.AcceptRunWhitespace()

		if q.AcceptToken(parser.Token{Type: TokenDelimiter, Data: ":"}) {
			q.AcceptRunWhitespace()

			t.Comments[1] = p.AcceptRunWhitespaceComments()

			p.AcceptRunWhitespace()
			p.AcceptToken(parser.Token{Type: TokenDelimiter, Data: ":"})

			t.Comments[2] = p.AcceptRunWhitespaceComments()

			p.AcceptRunWhitespace()

			q = p.NewGoal()
			t.Expression = new(Expression)

			if err := t.Expression.parse(q); err != nil {
				return p.Error("TypeParam", err)
			}

			p.Score(q)
		}
	}

	q := p.NewGoal()

	q.AcceptRunWhitespace()

	if q.Peek() == (parser.Token{Type: TokenDelimiter, Data: "]"}) {
		t.Comments[3] = p.AcceptRunWhitespaceCommentsNoNewline()
	} else {
		t.Comments[3] = p.AcceptRunWhitespaceComments()
	}

	t.Tokens = p.ToTokens()

	return nil
}

// ParameterList as defined in python@3.13.0:
// https://docs.python.org/release/3.13.0/reference/compound_stmts.html#grammar-token-python-grammar-parameter_list
//
// The first set of comments are parsed from the beginning of the list.
//
// The second and third sets of comments are parsed from before and after the '/' token.
//
// The fourth and fifth sets of comments are parsed from before and after the
// '*' token; the sixth set of comments is parsed from after the StarArg.
//
// The seventh and eighth set of comments are parsed from before and after the
// '**' token; the ninth set of comments are parse from after the StarStarArg.
//
// The final set of comments are parsed from after the ParameterList.
type ParameterList struct {
	DefParameters []DefParameter
	NoPosOnly     []DefParameter
	StarArg       *Parameter
	StarArgs      []DefParameter
	StarStarArg   *Parameter
	Comments      [10]Comments
	Tokens        Tokens
}

func (l *ParameterList) parse(p *pyParser, allowAnnotations bool) error {
	l.Comments[0] = p.AcceptRunWhitespaceCommentsNoNewlineIfMultiline()

	hasSlash := false
	q := p.NewGoal()

	q.AcceptRunWhitespace()

	target, err := l.parseStars(p, q, &l.DefParameters, allowAnnotations)
	if err != nil {
		return err
	}

	for target != nil && q.Peek().Type == TokenIdentifier {
		p.AcceptRunWhitespaceNoComment()

		q = p.NewGoal()

		var df DefParameter

		if err := df.parse(q, allowAnnotations); err != nil {
			return p.Error("ParameterList", err)
		}

		p.Score(q)

		*target = append(*target, df)
		q = p.NewGoal()

		q.AcceptRunWhitespace()

		if !q.AcceptToken(parser.Token{Type: TokenDelimiter, Data: ","}) {
			break
		}

		p.Score(q)

		q = p.NewGoal()

		q.AcceptRunWhitespace()

		switch target {
		case &l.DefParameters:
			if q.AcceptToken(parser.Token{Type: TokenOperator, Data: "/"}) {
				l.Comments[1] = p.AcceptRunWhitespaceCommentsIfMultiline()

				p.AcceptRunWhitespace()
				p.Next()

				hasSlash = true
				q = p.NewGoal()

				q.AcceptRunWhitespace()

				if !q.AcceptToken(parser.Token{Type: TokenDelimiter, Data: ","}) {
					l.Comments[2] = p.AcceptRunWhitespaceCommentsNoNewlineIfMultiline()

					target = nil

					break
				}

				l.Comments[2] = p.AcceptRunWhitespaceCommentsIfMultiline()

				p.AcceptRunWhitespace()
				p.Next()

				q = p.NewGoal()
				q.AcceptRunAllWhitespace()

				target = &l.NoPosOnly
			}

			fallthrough
		case &l.NoPosOnly:
			if target, err = l.parseStars(p, q, target, allowAnnotations); err != nil {
				return err
			}
		default:
			if q.AcceptToken(parser.Token{Type: TokenOperator, Data: "**"}) {
				if err = l.parseStarStar(p, q, allowAnnotations); err != nil {
					return err
				}

				target = nil
			}
		}
	}

	if !hasSlash {
		l.NoPosOnly = l.DefParameters
		l.DefParameters = nil
	}

	if allowAnnotations {
		l.Comments[9] = p.AcceptRunWhitespaceCommentsIfMultiline()
	} else {
		l.Comments[9] = p.AcceptRunWhitespaceCommentsNoNewlineIfMultiline()
	}

	l.Tokens = p.ToTokens()

	return nil
}

func (l *ParameterList) parseStars(p, q *pyParser, target *[]DefParameter, allowAnnotations bool) (*[]DefParameter, error) {
	if q.AcceptToken(parser.Token{Type: TokenOperator, Data: "*"}) {
		r := p.NewGoal()

		l.Comments[3] = r.AcceptRunWhitespaceCommentsIfMultiline()

		r.AcceptRunWhitespace()
		r.Next()

		l.Comments[4] = r.AcceptRunWhitespaceCommentsIfMultiline()

		r.AcceptRunWhitespace()
		p.Score(r)

		r = p.NewGoal()
		l.StarArg = new(Parameter)

		if err := l.StarArg.parse(r, allowAnnotations); err != nil {
			return nil, p.Error("ParameterList", err)
		}

		p.Score(r)

		target = &l.StarArgs
		r = p.NewGoal()

		r.AcceptRunWhitespace()

		if r.AcceptToken(parser.Token{Type: TokenDelimiter, Data: ","}) {
			l.Comments[5] = p.AcceptRunWhitespaceCommentsIfMultiline()

			p.AcceptRunAllWhitespace()
			p.Next()
		} else {
			l.Comments[5] = p.AcceptRunWhitespaceCommentsNoNewlineIfMultiline()
		}

		*q = *p.NewGoal()

		q.AcceptRunWhitespace()
	}

	if q.AcceptToken(parser.Token{Type: TokenOperator, Data: "**"}) {
		return nil, l.parseStarStar(p, q, allowAnnotations)
	}

	return target, nil
}

func (l *ParameterList) parseStarStar(p, q *pyParser, allowAnnotations bool) error {
	r := p.NewGoal()

	l.Comments[6] = r.AcceptRunWhitespaceCommentsIfMultiline()

	r.AcceptRunWhitespace()
	r.Next()

	l.Comments[7] = r.AcceptRunWhitespaceCommentsIfMultiline()

	r.AcceptRunWhitespace()
	p.Score(r)

	r = p.NewGoal()
	l.StarStarArg = new(Parameter)

	if err := l.StarStarArg.parse(r, allowAnnotations); err != nil {
		return p.Error("ParameterList", err)
	}

	p.Score(r)

	r = p.NewGoal()

	r.AcceptRunWhitespace()

	if r.AcceptToken(parser.Token{Type: TokenDelimiter, Data: ","}) {
		l.Comments[8] = p.AcceptRunWhitespaceCommentsIfMultiline()

		p.AcceptRunAllWhitespace()
		p.Next()
	} else {
		l.Comments[8] = p.AcceptRunWhitespaceCommentsNoNewlineIfMultiline()
	}

	*q = *p.NewGoal()

	return nil
}

// DefParameter as defined in python@3.13.0:
// https://docs.python.org/release/3.13.0/reference/compound_stmts.html#grammar-token-python-grammar-defparameter
//
// The first set of comments are parsed from before the DefParameter.
//
// The second and thrid set of comments are parsed from either side of a '='
// token.
//
// The final set of comments are parsed from after the DefParameter.
type DefParameter struct {
	Parameter Parameter
	Value     *Expression
	Comments  [4]Comments
	Tokens    Tokens
}

func (d *DefParameter) parse(p *pyParser, allowAnnotations bool) error {
	d.Comments[0] = p.AcceptRunWhitespaceCommentsIfMultiline()

	p.AcceptRunAllWhitespace()

	q := p.NewGoal()

	if err := d.Parameter.parse(q, allowAnnotations); err != nil {
		return p.Error("DefParameter", err)
	}

	p.Score(q)

	q = p.NewGoal()

	q.AcceptRunWhitespace()

	if q.AcceptToken(parser.Token{Type: TokenDelimiter, Data: "="}) {
		d.Comments[1] = p.AcceptRunWhitespaceCommentsIfMultiline()

		p.AcceptRunWhitespace()
		p.Next()

		d.Comments[2] = p.AcceptRunWhitespaceCommentsIfMultiline()

		p.AcceptRunWhitespace()

		q = p.NewGoal()
		d.Value = new(Expression)

		if err := d.Value.parse(q); err != nil {
			return p.Error("DefParameter", err)
		}

		p.Score(q)
	}

	q = p.NewGoal()

	q.AcceptRunWhitespace()

	if q.Peek() == (parser.Token{Type: TokenDelimiter, Data: ","}) {
		d.Comments[3] = p.AcceptRunWhitespaceCommentsIfMultiline()
	} else {
		d.Comments[3] = p.AcceptRunWhitespaceCommentsNoNewlineIfMultiline()
	}

	d.Tokens = p.ToTokens()

	return nil
}

// Parameter as defined in python@3.13.0:
// https://docs.python.org/release/3.13.0/reference/compound_stmts.html#grammar-token-python-grammar-parameter
//
// The comments are parsed on either side of the ':' token.
type Parameter struct {
	Identifier *Token
	Type       *Expression
	Comments   [2]Comments
	Tokens     Tokens
}

func (pr *Parameter) parse(p *pyParser, allowAnnotations bool) error {
	if !p.Accept(TokenIdentifier) {
		return p.Error("Parameter", ErrMissingIdentifier)
	}

	pr.Identifier = p.GetLastToken()
	q := p.NewGoal()

	q.AcceptRunWhitespace()

	if allowAnnotations && q.AcceptToken(parser.Token{Type: TokenDelimiter, Data: ":"}) {
		pr.Comments[0] = p.AcceptRunWhitespaceCommentsIfMultiline()

		p.AcceptRunWhitespace()
		p.Next()

		pr.Comments[1] = p.AcceptRunWhitespaceCommentsIfMultiline()

		p.AcceptRunWhitespace()

		q = p.NewGoal()
		pr.Type = new(Expression)

		if err := pr.Type.parse(q); err != nil {
			return p.Error("Parameter", err)
		}

		p.Score(q)
	}

	pr.Tokens = p.ToTokens()

	return nil
}

// Statement as defined in python@3.13.0:
// https://docs.python.org/release/3.13.0/reference/compound_stmts.html#grammar-token-python-grammar-statement
//
// The comments are parsed from before the Statement.
type Statement struct {
	StatementList     *StatementList
	CompoundStatement *CompoundStatement
	Comments          Comments
	Tokens            Tokens
}

func (s *Statement) parse(p *pyParser) error {
	var isCompound, isSoftCompound bool

	s.Comments = p.AcceptRunWhitespaceComments()

	p.AcceptRunAllWhitespace()

	q := p.NewGoal()

	switch tk := p.Peek(); tk.Type {
	case TokenOperator:
		isCompound = tk.Data == "@"
	case TokenKeyword:
		isCompound = slices.Contains(compounds[:], tk.Data)
	case TokenIdentifier:
		isCompound = tk.Data == "match"
		isSoftCompound = isCompound
	}

	if isCompound {
		c := new(CompoundStatement)

		if err := c.parse(q); err != nil {
			if !isSoftCompound {
				return p.Error("Statement", err)
			}
		} else {
			p.Score(q)

			s.CompoundStatement = c
		}
	}

	if s.CompoundStatement == nil {
		q = p.NewGoal()
		s.StatementList = new(StatementList)

		if err := s.StatementList.parse(q); err != nil {
			return p.Error("Statement", err)
		}

		p.Score(q)
	}

	s.Tokens = p.ToTokens()

	return nil
}

// StatementList as defined in python@3.13.0:
// https://docs.python.org/release/3.13.0/reference/compound_stmts.html#grammar-token-python-grammar-stmt_list
//
// The comments are parsed from after the StatementList.
type StatementList struct {
	Statements []SimpleStatement
	Comments   Comments
	Tokens
}

func (s *StatementList) parse(p *pyParser) error {
	for {
		q := p.NewGoal()

		var ss SimpleStatement

		if err := ss.parse(q); err != nil {
			return p.Error("StatementList", err)
		}

		p.Score(q)

		s.Statements = append(s.Statements, ss)
		q = p.NewGoal()

		q.AcceptRunWhitespace()

		if !q.AcceptToken(parser.Token{Type: TokenDelimiter, Data: ";"}) {
			break
		}

		p.Score(q)

		q = p.NewGoal()

		if tk := q.AcceptRunWhitespace(); tk == TokenComment || tk == TokenLineTerminator || tk == TokenDedent || tk == parser.TokenDone {
			break
		}

		p.Score(q)
	}

	s.Comments = p.AcceptRunWhitespaceCommentsNoNewline()
	s.Tokens = p.ToTokens()

	return nil
}

// TypeParams as defined in python@3.13.0:
// https://docs.python.org/release/3.13.0/reference/compound_stmts.html#grammar-token-python-grammar-type_params
//
// The first set of comments are parsed from directly after the opening bracket;
// the second set of comments are parsed from directly before the closing
// bracket.
type TypeParams struct {
	TypeParams []TypeParam
	Comments   [2]Comments
	Tokens     Tokens
}

func (t *TypeParams) parse(p *pyParser) error {
	p.AcceptToken(parser.Token{Type: TokenDelimiter, Data: "["})

	t.Comments[0] = p.AcceptRunWhitespaceCommentsNoNewline()
	q := p.NewGoal()

	q.AcceptRunAllWhitespace()

	if !q.AcceptToken(parser.Token{Type: TokenDelimiter, Data: "]"}) {
		for {
			p.AcceptRunWhitespaceNoComment()

			q := p.NewGoal()

			var tp TypeParam

			if err := tp.parse(q); err != nil {
				return p.Error("TypeParams", err)
			}

			t.TypeParams = append(t.TypeParams, tp)

			p.Score(q)

			q = p.NewGoal()

			q.AcceptRunAllWhitespace()

			if q.AcceptToken(parser.Token{Type: TokenDelimiter, Data: "]"}) {
				break
			} else if !q.AcceptToken(parser.Token{Type: TokenDelimiter, Data: ","}) {
				return q.Error("TypeParams", ErrMissingComma)
			}

			p.AcceptRunWhitespace()
			p.AcceptToken(parser.Token{Type: TokenDelimiter, Data: ","})
		}
	}

	t.Comments[1] = p.AcceptRunWhitespaceComments()

	p.AcceptRunAllWhitespace()

	p.AcceptToken(parser.Token{Type: TokenDelimiter, Data: "]"})

	t.Tokens = p.ToTokens()

	return nil
}

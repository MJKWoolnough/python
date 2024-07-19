package python

import (
	"errors"

	"vimagination.zapto.org/parser"
)

var compounds = [...]string{"if", "while", "for", "try", "with", "func", "class", "async", "def"}

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

func (c *CompoundStatement) parser(p *pyParser) error {
	var decorators *Decorators

	if tk := p.Peek(); tk.Type == TokenDelimiter && tk.Data == "@" {
		decorators = new(Decorators)

		q := p.NewGoal()

		if err := decorators.parse(q); err != nil {
			return p.Error("CompoundStatement", err)
		}

		p.Score(q)

		if tk := p.Peek(); tk.Type != TokenKeyword {
			return p.Error("CompoundStatement", ErrInvalidCompound)
		}
	}

	var err error

	q := p.NewGoal()

	switch tk := p.Peek(); tk.Data {
	case "if":
		c.If = new(IfStatement)
		err = c.If.parse(q)
	case "while":
		c.While = new(WhileStatement)
		err = c.While.parse(q)
	case "for":
		c.For = new(ForStatement)
		err = c.For.parse(q, false)
	case "try":
		c.Try = new(TryStatement)
		err = c.Try.parse(q)
	case "with":
		c.With = new(WithStatement)
		err = c.With.parse(q, false)
	case "def":
		c.Func = new(FuncDefinition)
		err = c.Func.parse(q, false, decorators)
	case "class":
		c.Class = new(ClassDefinition)
		err = c.Class.parse(q, decorators)
	case "async":
		p.next()
		p.AcceptRun(TokenWhitespace)

		switch tk := p.Peek(); tk.Data {
		case "for":
			c.For = new(ForStatement)
			err = c.For.parse(q, true)
		case "with":
			c.With = new(WithStatement)
			err = c.With.parse(q, true)
		case "def":
			c.Func = new(FuncDefinition)
			err = c.Func.parse(q, true, decorators)
		default:
			err = ErrInvalidCompound
		}
	default:
		err = ErrInvalidCompound
	}

	if err != nil {
		return p.Error("CompoundStatement", err)
	}

	p.Score(q)

	c.Tokens = p.ToTokens()

	return nil
}

type Decorators struct {
	Decorators []AssignmentExpression
	Tokens
}

func (d *Decorators) parse(p *pyParser) error {
	for p.AcceptToken(parser.Token{Type: TokenDelimiter, Data: "@"}) {
		var ae AssignmentExpression

		q := p.NewGoal()

		if err := ae.parse(q); err != nil {
			return p.Error("Decorator", err)
		}

		q.AcceptRun(TokenWhitespace)
		p.Score(q)

		if !q.Accept(TokenLineTerminator) {
			return p.Error("Decorator", ErrMissingNewline)
		}
	}

	d.Tokens = p.ToTokens()

	return nil
}

type IfStatement struct {
	If     AssignmentExpressionAndSuite
	Elif   []AssignmentExpressionAndSuite
	Else   *AssignmentExpression
	Tokens Tokens
}

func (i *IfStatement) parse(p *pyParser) error {
	p.AcceptToken(parser.Token{Type: TokenKeyword, Data: "if"})
	p.AcceptRun(TokenWhitespace)

	q := p.NewGoal()

	if err := i.If.parse(q); err != nil {
		return p.Error("IfStatement", err)
	}

	p.Score(q)

	q = p.NewGoal()

	q.AcceptRun(TokenLineTerminator)

	for q.AcceptToken(parser.Token{Type: TokenKeyword, Data: "elif"}) {
		q.AcceptRun(TokenWhitespace)
		p.Score(q)

		q := p.NewGoal()

		var as AssignmentExpressionAndSuite

		if err := as.parse(q); err != nil {
			return p.Error("IfStatement", err)
		}

		i.Elif = append(i.Elif, as)

		p.Score(q)
		p.AcceptRun(TokenWhitespace)

		q = p.NewGoal()

		q.AcceptRun(TokenLineTerminator)
	}

	if q.AcceptToken(parser.Token{Type: TokenKeyword, Data: "else"}) {
		p.Score(q)
		p.AcceptRun(TokenWhitespace)

		if !p.AcceptToken(parser.Token{Type: TokenDelimiter, Data: ":"}) {
			return p.Error("IfStatement", ErrMissingColon)
		}

		p.AcceptRun(TokenWhitespace)

		q = p.NewGoal()

		i.Else = new(AssignmentExpression)

		if err := i.Else.parse(q); err != nil {
			return p.Error("IfStatement", err)
		}
	}

	p.Score(q)

	i.Tokens = p.ToTokens()

	return nil
}

type WhileStatement struct{}

func (w *WhileStatement) parse(_ *pyParser) error {
	return nil
}

type ForStatement struct{}

func (f *ForStatement) parse(_ *pyParser, _ bool) error {
	return nil
}

type TryStatement struct{}

func (t *TryStatement) parse(_ *pyParser) error {
	return nil
}

type WithStatement struct{}

func (w *WithStatement) parse(_ *pyParser, _ bool) error {
	return nil
}

type FuncDefinition struct{}

func (f *FuncDefinition) parse(_ *pyParser, _ bool, _ *Decorators) error {
	return nil
}

type ClassDefinition struct{}

func (c *ClassDefinition) parse(_ *pyParser, _ *Decorators) error {
	return nil
}

type AssignmentExpressionAndSuite struct {
	AssignmentExpression AssignmentExpression
	Suite                Suite
	Tokens               Tokens
}

func (a *AssignmentExpressionAndSuite) parse(_ *pyParser) error {
	return nil
}

type Suite struct{}

func (s *Suite) parse(_ *pyParser) error {
	return nil
}

var (
	ErrInvalidCompound = errors.New("invalid compound statement")
	ErrMissingNewline  = errors.New("missing newline")
	ErrMissingColon    = errors.New("missing colon")
)

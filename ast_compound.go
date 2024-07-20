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
	Else   *Suite
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

		q = p.NewGoal()

		q.AcceptRun(TokenLineTerminator)
	}

	q = p.NewGoal()

	q.AcceptRun(TokenLineTerminator)

	if q.AcceptToken(parser.Token{Type: TokenKeyword, Data: "else"}) {
		p.Score(q)
		p.AcceptRun(TokenWhitespace)

		if !p.AcceptToken(parser.Token{Type: TokenDelimiter, Data: ":"}) {
			return p.Error("IfStatement", ErrMissingColon)
		}

		p.AcceptRun(TokenWhitespace)

		q = p.NewGoal()

		i.Else = new(Suite)

		if err := i.Else.parse(q); err != nil {
			return p.Error("IfStatement", err)
		}

		p.Score(q)
	}

	i.Tokens = p.ToTokens()

	return nil
}

type WhileStatement struct {
	While  AssignmentExpressionAndSuite
	Else   *Suite
	Tokens Tokens
}

func (w *WhileStatement) parse(p *pyParser) error {
	p.AcceptToken(parser.Token{Type: TokenKeyword, Data: "while"})
	p.AcceptRun(TokenWhitespace)

	q := p.NewGoal()

	if err := w.While.parse(q); err != nil {
		return p.Error("WhileStatement", err)
	}

	p.Score(q)

	q = p.NewGoal()

	q.AcceptRun(TokenLineTerminator)

	if q.AcceptToken(parser.Token{Type: TokenKeyword, Data: "else"}) {
		p.Score(q)
		p.AcceptRun(TokenWhitespace)

		if !p.AcceptToken(parser.Token{Type: TokenDelimiter, Data: ":"}) {
			return p.Error("WhileStatement", ErrMissingColon)
		}

		p.AcceptRun(TokenWhitespace)

		q = p.NewGoal()

		w.Else = new(Suite)

		if err := w.Else.parse(q); err != nil {
			return p.Error("WhileStatement", err)
		}

		p.Score(q)
	}

	w.Tokens = p.ToTokens()

	return nil
}

type ForStatement struct {
	Async       bool
	TargetList  TargetList
	StarredList StarredList
	Suite       Suite
	Else        *Suite
	Tokens
}

func (f *ForStatement) parse(p *pyParser, async bool) error {
	f.Async = async

	p.AcceptToken(parser.Token{Type: TokenKeyword, Data: "while"})
	p.AcceptRun(TokenWhitespace)

	q := p.NewGoal()

	if err := f.TargetList.parse(p); err != nil {
		return p.Error("ForStatement", err)
	}

	p.Score(q)

	p.AcceptRun(TokenWhitespace)

	if !p.AcceptToken(parser.Token{Type: TokenKeyword, Data: "in"}) {
		return p.Error("ForStatement", ErrMissingIn)
	}

	p.AcceptRun(TokenWhitespace)

	q = p.NewGoal()

	if err := f.StarredList.parse(p); err != nil {
		return p.Error("ForStatement", err)
	}

	p.Score(q)

	p.AcceptRun(TokenWhitespace)

	if !p.AcceptToken(parser.Token{Type: TokenDelimiter, Data: ":"}) {
		return p.Error("ForStatement", ErrMissingIn)
	}

	p.AcceptRun(TokenWhitespace)

	q = p.NewGoal()

	if err := f.Suite.parse(q); err != nil {
		return p.Error("ForStatement", err)
	}

	p.Score(q)

	q = p.NewGoal()

	q.AcceptRun(TokenLineTerminator)

	if q.AcceptToken(parser.Token{Type: TokenKeyword, Data: "else"}) {
		p.Score(q)
		p.AcceptRun(TokenWhitespace)

		if !p.AcceptToken(parser.Token{Type: TokenDelimiter, Data: ":"}) {
			return p.Error("ForStatement", ErrMissingColon)
		}

		p.AcceptRun(TokenWhitespace)

		q = p.NewGoal()

		f.Else = new(Suite)

		if err := f.Else.parse(q); err != nil {
			return p.Error("ForStatement", err)
		}

		p.Score(q)
	}

	f.Tokens = p.ToTokens()

	return nil
}

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
	p.AcceptRun(TokenWhitespace)

	if !p.AcceptToken(parser.Token{Type: TokenDelimiter, Data: ":"}) {
		return p.Error("TryStatement", ErrMissingColon)
	}

	q := p.NewGoal()

	q.AcceptRun(TokenLineTerminator)

	for p.AcceptToken(parser.Token{Type: TokenKeyword, Data: "except"}) {
		p.Score(q)
		p.AcceptRun(TokenWhitespace)

		group := p.AcceptToken(parser.Token{Type: TokenOperator, Data: "*"})

		if len(t.Except) > 0 && t.Groups != group {
			return p.Error("TryStatement", ErrMismatchedGroups)
		}

		t.Groups = group

		p.AcceptRun(TokenWhitespace)

		q = p.NewGoal()

		var except Except

		if err := except.parse(q); err != nil {
			return p.Error("TryStatement", err)
		}

		t.Except = append(t.Except, except)

		q.Score(p)
	}

	q = p.NewGoal()

	q.AcceptRun(TokenLineTerminator)

	if len(t.Except) > 0 && p.AcceptToken(parser.Token{Type: TokenKeyword, Data: "else"}) {
		p.Score(q)
		p.AcceptRun(TokenWhitespace)

		if !p.AcceptToken(parser.Token{Type: TokenDelimiter, Data: ":"}) {
			return p.Error("TryStatement", ErrMissingColon)
		}

		p.AcceptRun(TokenWhitespace)

		q := p.NewGoal()

		t.Else = new(Suite)

		if err := t.Else.parse(q); err != nil {
			return p.Error("TryStatement", err)
		}

		p.Score(q)
	}

	q = p.NewGoal()

	q.AcceptRun(TokenLineTerminator)

	if q.AcceptToken(parser.Token{Type: TokenKeyword, Data: "finally"}) {
		p.Score(q)
		p.AcceptRun(TokenWhitespace)

		if !p.AcceptToken(parser.Token{Type: TokenDelimiter, Data: ":"}) {
			return p.Error("TryStatement", ErrMissingColon)
		}

		p.AcceptRun(TokenWhitespace)

		q := p.NewGoal()

		t.Finally = new(Suite)

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

type Except struct {
	Expression Expression
	Identifier *string
	Suite      Suite
	Tokens     Tokens
}

func (e *Except) parse(p *pyParser) error {
	q := p.NewGoal()

	if err := e.Expression.parse(q); err != nil {
		return p.Error("Except", err)
	}

	p.Score(q)

	p.AcceptRun(TokenWhitespace)

	if p.AcceptToken(parser.Token{Type: TokenKeyword, Data: "as"}) {
		p.AcceptRun(TokenWhitespace)

		token := p.next()

		if token.Type != TokenIdentifier {
			return p.Error("Except", ErrMissingIdentifier)
		}

		e.Identifier = &token.Data

		p.AcceptRun(TokenWhitespace)
	}

	if !p.AcceptToken(parser.Token{Type: TokenDelimiter, Data: ":"}) {
		return p.Error("Except", ErrMissingColon)
	}

	p.AcceptRun(TokenWhitespace)

	q = p.NewGoal()

	if err := e.Suite.parse(q); err != nil {
		return p.Error("Except", err)
	}

	p.Score(q)

	e.Tokens = p.ToTokens()

	return nil
}

type WithStatement struct {
	Async    bool
	Contents WithStatementContents
	Suite    Suite
	Tokens   Tokens
}

func (w *WithStatement) parse(p *pyParser, async bool) error {
	w.Async = async

	p.AcceptToken(parser.Token{Type: TokenKeyword, Data: "while"})
	p.AcceptRun(TokenWhitespace)

	parens := p.AcceptToken(parser.Token{Type: TokenDelimiter, Data: "("})

	p.AcceptRun(TokenWhitespace)

	q := p.NewGoal()

	if err := w.Contents.parse(q); err != nil {
		return p.Error("WithStatement", err)
	}

	p.Score(q)

	p.AcceptRun(TokenWhitespace)

	if parens {
		if p.AcceptToken(parser.Token{Type: TokenDelimiter, Data: ","}) {
			p.AcceptRun(TokenWhitespace)
		}

		if !p.AcceptToken(parser.Token{Type: TokenDelimiter, Data: ")"}) {
			return p.Error("WithStatement", ErrMissingClosingParen)
		}

		p.AcceptRun(TokenWhitespace)
	}

	q = p.NewGoal()

	if err := w.Suite.parse(q); err != nil {
		return p.Error("WithStatement", err)
	}

	w.Tokens = p.ToTokens()

	return nil
}

type WithStatementContents struct {
	Items  []WithItem
	Tokens Tokens
}

func (w *WithStatementContents) parse(p *pyParser) error {
	q := p.NewGoal()

	another := true

	for another {
		var wi WithItem

		r := q.NewGoal()

		if err := wi.parse(r); err != nil {
			return p.Error("WithStatementContents", err)
		}

		q.Score(r)
		p.Score(q)

		q = p.NewGoal()

		q.AcceptRun(TokenWhitespace)

		another = q.AcceptToken(parser.Token{Type: TokenDelimiter, Data: ","})

		q.AcceptRun(TokenWhitespace)

		if q.Peek() == (parser.Token{Type: TokenDelimiter, Data: ")"}) {
			break
		}
	}

	w.Tokens = p.ToTokens()

	return nil
}

type WithItem struct {
	Expression Expression
	Target     *Target
	Tokens     Tokens
}

func (w *WithItem) parse(p *pyParser) error {
	q := p.NewGoal()

	if err := w.Expression.parse(q); err != nil {
		return p.Error("WithItem", err)
	}

	p.Score(q)

	q = p.NewGoal()

	q.AcceptRun(TokenWhitespace)

	if q.AcceptToken(parser.Token{Type: TokenKeyword, Data: "as"}) {
		q.AcceptRun(TokenWhitespace)
		p.Score(q)

		q = p.NewGoal()
		w.Target = new(Target)

		if err := w.Target.parse(q); err != nil {
			return p.Error("WithItem", err)
		}

		p.Score(q)
	}

	w.Tokens = p.ToTokens()

	return nil
}

type FuncDefinition struct {
	Decorators    *Decorators
	Async         bool
	FuncName      *Token
	TypeParams    []TypeParam
	ParameterList ParameterList
	Expression    *Expression
	Suite         Suite
	Tokens        Tokens
}

func (f *FuncDefinition) parse(p *pyParser, async bool, decorators *Decorators) error {
	f.Decorators = decorators
	f.Async = async

	p.AcceptToken(parser.Token{Type: TokenKeyword, Data: "def"})
	p.AcceptRun(TokenWhitespace)

	if !p.Accept(TokenIdentifier) {
		return p.Error("FuncDefinition", ErrMissingIdentifier)
	}

	f.FuncName = p.GetLastToken()

	if p.AcceptToken(parser.Token{Type: TokenDelimiter, Data: "["}) {
		for {
			p.AcceptRun(TokenWhitespace)

			q := p.NewGoal()

			var t TypeParam

			if err := t.parse(q); err != nil {
				return p.Error("FuncDefinition", err)
			}

			p.Score(q)

			f.TypeParams = append(f.TypeParams, t)

			p.AcceptRun(TokenWhitespace)

			if !p.AcceptToken(parser.Token{Type: TokenDelimiter, Data: ","}) {
				if !p.AcceptToken(parser.Token{Type: TokenDelimiter, Data: "]"}) {
					return p.Error("FuncDefinition", ErrMissingClosingBracket)
				}

				break
			}
		}

		p.AcceptRun(TokenWhitespace)
	}

	if !p.AcceptToken(parser.Token{Type: TokenDelimiter, Data: "("}) {
		return p.Error("FuncDefinition", ErrMissingOpeningParen)
	}

	p.AcceptRun(TokenWhitespace)

	if !p.AcceptToken(parser.Token{Type: TokenDelimiter, Data: ")"}) {
		q := p.NewGoal()

		if err := f.ParameterList.parse(q); err != nil {
			return p.Error("FuncDefinition", err)
		}

		p.AcceptRun(TokenWhitespace)

		if !p.AcceptToken(parser.Token{Type: TokenDelimiter, Data: ")"}) {
			return p.Error("FuncDefinition", ErrMissingClosingParen)
		}

		p.AcceptRun(TokenWhitespace)
	}

	if p.AcceptToken(parser.Token{Type: TokenDelimiter, Data: "->"}) {
		p.AcceptRun(TokenWhitespace)

		q := p.NewGoal()

		f.Expression = new(Expression)

		if err := f.Expression.parse(q); err != nil {
			return p.Error("FuncDefinition", err)
		}

		p.AcceptRun(TokenWhitespace)
	}

	if !p.AcceptToken(parser.Token{Type: TokenDelimiter, Data: ":"}) {
		return p.Error("FuncDefinition", ErrMissingColon)
	}

	p.AcceptRun(TokenWhitespace)

	q := p.NewGoal()

	if err := f.Suite.parse(q); err != nil {
		return p.Error("FuncDefinition", err)
	}

	p.Score(q)

	f.Tokens = p.ToTokens()

	return nil
}

type ClassDefinition struct {
	Decorators  *Decorators
	ClassName   *Token
	TypeParams  []TypeParam
	Inheritance ArgumentList
	Suite       Suite
	Tokens      Tokens
}

func (c *ClassDefinition) parse(p *pyParser, decorators *Decorators) error {
	c.Decorators = decorators

	p.AcceptToken(parser.Token{Type: TokenKeyword, Data: "class"})
	p.AcceptRun(TokenWhitespace)

	if !p.Accept(TokenIdentifier) {
		return p.Error("ClassDefinition", ErrMissingIdentifier)
	}

	c.ClassName = p.GetLastToken()

	if p.AcceptToken(parser.Token{Type: TokenDelimiter, Data: "["}) {
		for {
			p.AcceptRun(TokenWhitespace)

			q := p.NewGoal()

			var t TypeParam

			if err := t.parse(q); err != nil {
				return p.Error("ClassDefinition", err)
			}

			p.Score(q)

			c.TypeParams = append(c.TypeParams, t)

			p.AcceptRun(TokenWhitespace)

			if !p.AcceptToken(parser.Token{Type: TokenDelimiter, Data: ","}) {
				if !p.AcceptToken(parser.Token{Type: TokenDelimiter, Data: "]"}) {
					return p.Error("ClassDefinition", ErrMissingClosingBracket)
				}

				break
			}
		}

		p.AcceptRun(TokenWhitespace)
	}

	if p.AcceptToken(parser.Token{Type: TokenDelimiter, Data: "("}) {
		p.AcceptRun(TokenWhitespace)

		if !p.AcceptToken(parser.Token{Type: TokenDelimiter, Data: ")"}) {
			q := p.NewGoal()

			if err := c.Inheritance.parse(q); err != nil {
				return p.Error("ClassDefinition", err)
			}

			p.AcceptRun(TokenWhitespace)

			if !p.AcceptToken(parser.Token{Type: TokenDelimiter, Data: ")"}) {
				return p.Error("ClassDefinition", ErrMissingClosingParen)
			}

			p.AcceptRun(TokenWhitespace)
		}
	}

	if !p.AcceptToken(parser.Token{Type: TokenDelimiter, Data: ":"}) {
		return p.Error("ClassDefinition", ErrMissingColon)
	}

	p.AcceptRun(TokenWhitespace)

	q := p.NewGoal()

	if err := c.Suite.parse(q); err != nil {
		return p.Error("ClassDefinition", err)
	}

	p.Score(q)

	c.Tokens = p.ToTokens()

	return nil
}

type AssignmentExpressionAndSuite struct {
	AssignmentExpression AssignmentExpression
	Suite                Suite
	Tokens               Tokens
}

func (a *AssignmentExpressionAndSuite) parse(p *pyParser) error {
	q := p.NewGoal()

	if err := a.AssignmentExpression.parse(q); err != nil {
		return p.Error("AssignmentExpressionAndSuite", err)
	}

	p.Score(q)

	p.AcceptRun(TokenWhitespace)

	if !p.AcceptToken(parser.Token{Type: TokenDelimiter, Data: ":"}) {
		return p.Error("AssignmentExpressionAndSuite", ErrMissingColon)
	}

	p.AcceptRun(TokenWhitespace)

	q = p.NewGoal()

	if err := a.Suite.parse(q); err != nil {
		return p.Error("AssignmentExpressionAndSuite", err)
	}

	p.Score(q)

	a.Tokens = p.ToTokens()

	return nil
}

type Suite struct{}

func (s *Suite) parse(_ *pyParser) error {
	return nil
}

type TargetList struct{}

func (t *TargetList) parse(_ *pyParser) error {
	return nil
}

type StarredList struct{}

func (s *StarredList) parse(_ *pyParser) error {
	return nil
}

type Target struct{}

func (t *Target) parse(_ *pyParser) error {
	return nil
}

type TypeParam struct{}

func (t *TypeParam) parse(_ *pyParser) error {
	return nil
}

type ParameterList struct{}

func (l *ParameterList) parse(_ *pyParser) error {
	return nil
}

type ArgumentList struct{}

func (a *ArgumentList) parse(_ *pyParser) error {
	return nil
}

var (
	ErrInvalidCompound       = errors.New("invalid compound statement")
	ErrMissingNewline        = errors.New("missing newline")
	ErrMissingColon          = errors.New("missing colon")
	ErrMissingIn             = errors.New("missing in")
	ErrMissingFinally        = errors.New("missing finally")
	ErrMissingIdentifier     = errors.New("missing identifier")
	ErrMismatchedGroups      = errors.New("mismatched groups in except")
	ErrMissingOpeningParen   = errors.New("missing opening paren")
	ErrMissingClosingParen   = errors.New("missing closing paren")
	ErrMissingClosingBracket = errors.New("missing closing bracket")
)

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

	if err := f.TargetList.parse(p, whitespaceToken); err != nil {
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

	if err := e.Expression.parse(q, whitespaceToken); err != nil {
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

	p.AcceptRun(TokenWhitespace, TokenComment)

	q := p.NewGoal()

	if err := w.Contents.parse(q, parens); err != nil {
		return p.Error("WithStatement", err)
	}

	p.Score(q)

	p.AcceptRun(TokenWhitespace, TokenComment)

	if parens {
		if p.AcceptToken(parser.Token{Type: TokenDelimiter, Data: ","}) {
			p.AcceptRun(TokenWhitespace, TokenComment)
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

var (
	whitespaceCommentTokens = []parser.TokenType{TokenWhitespace, TokenComment}
	whitespaceToken         = whitespaceCommentTokens[:1]
)

func (w *WithStatementContents) parse(p *pyParser, inParen bool) error {
	q := p.NewGoal()

	another := true
	ws := whitespaceToken

	if inParen {
		ws = whitespaceCommentTokens
	}

	for another {
		var wi WithItem

		r := q.NewGoal()

		if err := wi.parse(r, ws); err != nil {
			return p.Error("WithStatementContents", err)
		}

		q.Score(r)
		p.Score(q)

		q = p.NewGoal()

		q.AcceptRun(ws...)

		another = q.AcceptToken(parser.Token{Type: TokenDelimiter, Data: ","})

		q.AcceptRun(ws...)

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

func (w *WithItem) parse(p *pyParser, ws []parser.TokenType) error {
	q := p.NewGoal()

	if err := w.Expression.parse(q, ws); err != nil {
		return p.Error("WithItem", err)
	}

	p.Score(q)

	q = p.NewGoal()

	q.AcceptRun(ws...)

	if q.AcceptToken(parser.Token{Type: TokenKeyword, Data: "as"}) {
		q.AcceptRun(ws...)
		p.Score(q)

		q = p.NewGoal()
		w.Target = new(Target)

		if err := w.Target.parse(q, ws); err != nil {
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
			p.AcceptRun(TokenWhitespace, TokenComment)

			q := p.NewGoal()

			var t TypeParam

			if err := t.parse(q); err != nil {
				return p.Error("FuncDefinition", err)
			}

			p.Score(q)

			f.TypeParams = append(f.TypeParams, t)

			p.AcceptRun(TokenWhitespace, TokenComment)

			if !p.AcceptToken(parser.Token{Type: TokenDelimiter, Data: ","}) {
				if !p.AcceptToken(parser.Token{Type: TokenDelimiter, Data: "]"}) {
					return p.Error("FuncDefinition", ErrMissingClosingBracket)
				}

				break
			}
		}

		p.AcceptRun(TokenWhitespace, TokenComment)
	}

	if !p.AcceptToken(parser.Token{Type: TokenDelimiter, Data: "("}) {
		return p.Error("FuncDefinition", ErrMissingOpeningParen)
	}

	p.AcceptRun(TokenWhitespace, TokenComment)

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

		if err := f.Expression.parse(q, whitespaceToken); err != nil {
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
			p.AcceptRun(TokenWhitespace, TokenComment)

			q := p.NewGoal()

			var t TypeParam

			if err := t.parse(q); err != nil {
				return p.Error("ClassDefinition", err)
			}

			p.Score(q)

			c.TypeParams = append(c.TypeParams, t)

			p.AcceptRun(TokenWhitespace, TokenComment)

			if !p.AcceptToken(parser.Token{Type: TokenDelimiter, Data: ","}) {
				if !p.AcceptToken(parser.Token{Type: TokenDelimiter, Data: "]"}) {
					return p.Error("ClassDefinition", ErrMissingClosingBracket)
				}

				break
			}
		}

		p.AcceptRun(TokenWhitespace, TokenComment)
	}

	if p.AcceptToken(parser.Token{Type: TokenDelimiter, Data: "("}) {
		p.AcceptRun(TokenWhitespace, TokenComment)

		if !p.AcceptToken(parser.Token{Type: TokenDelimiter, Data: ")"}) {
			q := p.NewGoal()

			if err := c.Inheritance.parse(q); err != nil {
				return p.Error("ClassDefinition", err)
			}

			p.AcceptRun(TokenWhitespace, TokenComment)

			if !p.AcceptToken(parser.Token{Type: TokenDelimiter, Data: ")"}) {
				return p.Error("ClassDefinition", ErrMissingClosingParen)
			}

			p.AcceptRun(TokenWhitespace, TokenComment)
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

type Suite struct {
	StatementList *StatementList
	Statements    []Statement
	Tokens        Tokens
}

func (s *Suite) parse(p *pyParser) error {
	if p.Accept(TokenLineTerminator) {
		p.AcceptRun(TokenLineTerminator, TokenWhitespace, TokenComment)

		if !p.Accept(TokenIndent) {
			return p.Error("Suite", ErrMissingIndent)
		}

		p.AcceptRun(TokenLineTerminator, TokenWhitespace, TokenComment)

		for {
			q := p.NewGoal()

			var stmt Statement

			if err := stmt.parse(q); err != nil {
				return p.Error("Suite", err)
			}

			s.Statements = append(s.Statements, stmt)

			p.Score(q)

			p.AcceptRun(TokenLineTerminator, TokenWhitespace, TokenComment)

			if p.Accept(TokenDedent) {
				break
			}

			p.AcceptRun(TokenLineTerminator, TokenWhitespace, TokenComment)
		}
	} else {
		s.StatementList = new(StatementList)

		q := p.NewGoal()

		if err := s.StatementList.parse(p); err != nil {
			return p.Error("Suite", err)
		}

		p.Score(q)
	}

	s.Tokens = p.ToTokens()

	return nil
}

type TargetList struct {
	Targets []Target
	Tokens  Tokens
}

func (t *TargetList) parse(p *pyParser, ws []parser.TokenType) error {
Loop:
	for {
		q := p.NewGoal()

		var tg Target

		if err := tg.parse(q, ws); err != nil {
			return p.Error("TargetList", err)
		}

		t.Targets = append(t.Targets, tg)

		p.Score(q)

		q = p.NewGoal()

		q.AcceptRun(ws...)

		if !q.AcceptToken(parser.Token{Type: TokenDelimiter, Data: ","}) {
			break
		}

		switch tk := q.Peek(); tk {
		case parser.Token{Type: TokenDelimiter, Data: ";"}:
		case parser.Token{Type: TokenDelimiter, Data: "="}:
		case parser.Token{Type: TokenDelimiter, Data: "]"}:
		case parser.Token{Type: TokenDelimiter, Data: ")"}:
		case parser.Token{Type: TokenKeyword, Data: "in"}:
		default:
			if tk.Type != TokenLineTerminator {
				break Loop
			}
		}

		q.AcceptRun(ws...)
		p.Score(q)
	}

	t.Tokens = p.ToTokens()

	return nil
}

type Target struct {
	Primary      *Primary
	Tuple        *TargetList
	Array        *TargetList
	AttributeRef *Token
	Slicing      *SliceList
	Star         *Target
	Tokens       Tokens
}

func (t *Target) parse(p *pyParser, ws []parser.TokenType) error {
	if p.AcceptToken(parser.Token{Type: TokenDelimiter, Data: "("}) {
		p.AcceptRun(whitespaceCommentTokens...)

		q := p.NewGoal()

		t.Tuple = new(TargetList)

		if err := t.Tuple.parse(p, whitespaceCommentTokens); err != nil {
			return p.Error("Target", err)
		}

		p.Score(q)

		p.AcceptRun(whitespaceCommentTokens...)

		if !p.AcceptToken(parser.Token{Type: TokenDelimiter, Data: ")"}) {
			return p.Error("Target", ErrMissingClosingParen)
		}
	} else if p.AcceptToken(parser.Token{Type: TokenDelimiter, Data: "["}) {
		p.AcceptRun(whitespaceCommentTokens...)

		q := p.NewGoal()

		t.Array = new(TargetList)

		if err := t.Array.parse(p, whitespaceCommentTokens); err != nil {
			return p.Error("Target", err)
		}

		p.Score(q)

		p.AcceptRun(whitespaceCommentTokens...)

		if !p.AcceptToken(parser.Token{Type: TokenDelimiter, Data: "]"}) {
			return p.Error("Target", ErrMissingClosingBracket)
		}
	} else if p.AcceptToken(parser.Token{Type: TokenOperator, Data: "*"}) {
		p.AcceptRun(ws...)

		q := p.NewGoal()

		t.Star = new(Target)

		if err := t.Star.parse(q, ws); err != nil {
			return p.Error("Target", err)
		}

		p.Score(q)
	} else {
		t.Primary = new(Primary)

		q := p.NewGoal()

		if err := t.Primary.parse(q); err != nil {
			return err
		}

		r := q.NewGoal()

		r.AcceptRun(ws...)

		switch r.Peek() {
		case parser.Token{Type: TokenDelimiter, Data: "."}:
			q.Score(r)
			p.Score(q)

			p.AcceptRun(ws...)

			if !p.Accept(TokenIdentifier) {
				return p.Error("Target", ErrMissingIdentifier)
			}

			t.AttributeRef = p.GetLastToken()
		case parser.Token{Type: TokenDelimiter, Data: "["}:
			p.AcceptRun(whitespaceCommentTokens...)

			t.Slicing = new(SliceList)

			q := p.NewGoal()

			if err := t.Slicing.parse(q); err != nil {
				return p.Error("Target", err)
			}

			p.Score(q)

			p.AcceptRun(whitespaceCommentTokens...)

			if !p.AcceptToken(parser.Token{Type: TokenDelimiter, Data: "]"}) {
				return p.Error("Target", ErrMissingClosingBracket)
			}
		default:
			if !t.Primary.IsIdentifier() {
				return p.Error("Target", ErrMissingIdentifier)
			}
		}
	}

	t.Tokens = p.ToTokens()

	return nil
}

type StarredList struct {
	StarredItems []StarredItem
	Tokens       Tokens
}

func (s *StarredList) parse(p *pyParser) error {
Loop:
	for {
		q := p.NewGoal()

		var si StarredItem

		if err := si.parse(q); err != nil {
			return p.Error("StarredList", err)
		}

		s.StarredItems = append(s.StarredItems, si)

		p.Score(q)

		q = p.NewGoal()

		q.AcceptRun(TokenWhitespace)

		if !q.AcceptToken(parser.Token{Type: TokenDelimiter, Data: "]"}) {
			break
		}

		switch q.Peek() {
		case parser.Token{Type: TokenDelimiter, Data: "]"}:
		case parser.Token{Type: TokenDelimiter, Data: "}"}:
		case parser.Token{Type: TokenDelimiter, Data: ":"}:
			break Loop
		}

		q.AcceptRun(TokenWhitespace)

		p.Score(q)
	}

	s.Tokens = p.ToTokens()

	return nil
}

type StarredItem struct{}

func (s *StarredItem) parse(_ *pyParser) error {
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
	ErrMissingIndent         = errors.New("missing indent")
)

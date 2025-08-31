package python

import "vimagination.zapto.org/parser"

var compounds = [...]string{"if", "while", "for", "try", "with", "func", "class", "async", "def"}

// CompoundStatement as defined in python@3.13.0:
// https://docs.python.org/3.13.0/reference/compound_stmts.html#grammar-token-python-grammar-compound_stmt
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
	var (
		decorators *Decorators
		err        error
	)

	q := p.NewGoal()

	if tk := q.Peek(); tk == (parser.Token{Type: TokenOperator, Data: "@"}) {
		decorators = new(Decorators)

		if err = decorators.parse(q); err != nil {
			return p.Error("CompoundStatement", err)
		}

		q.AcceptRunWhitespace()

		switch tk := q.Peek(); tk.Data {
		case "def":
			c.Func = new(FuncDefinition)
			err = c.Func.parse(q, false, decorators)
		case "class":
			c.Class = new(ClassDefinition)
			err = c.Class.parse(q, decorators)
		case "async":
			q.Next()
			q.AcceptRunWhitespace()

			switch tk := q.Peek(); tk.Data {
			case "def":
				c.Func = new(FuncDefinition)
				err = c.Func.parse(q, true, decorators)
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
			q.Next()
			q.AcceptRunWhitespace()

			switch tk := q.Peek(); tk.Data {
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
	}

	if err != nil {
		return p.Error("CompoundStatement", err)
	}

	p.Score(q)

	c.Tokens = p.ToTokens()

	return nil
}

// Decorators as defined in python@3.13.0:
// https://docs.python.org/3.13.0/reference/compound_stmts.html#grammar-token-python-grammar-decorators
type Decorators struct {
	Decorators []AssignmentExpression
	Tokens
}

func (d *Decorators) parse(p *pyParser) error {
	q := p.NewGoal()

	for q.AcceptToken(parser.Token{Type: TokenOperator, Data: "@"}) {
		q.AcceptRunWhitespace()

		var ae AssignmentExpression

		p.Score(q)

		q = p.NewGoal()

		if err := ae.parse(q); err != nil {
			return p.Error("Decorators", err)
		}

		d.Decorators = append(d.Decorators, ae)

		p.Score(q)
		p.AcceptRunWhitespace()

		if !p.Accept(TokenLineTerminator) {
			return p.Error("Decorators", ErrMissingNewline)
		}

		q = p.NewGoal()

		q.AcceptRunAllWhitespace()
	}

	d.Tokens = p.ToTokens()

	return nil
}

// IfStatement as defined in python@3.13.0:
// https://docs.python.org/3.13.0/reference/compound_stmts.html#grammar-token-python-grammar-if_stmt
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
// https://docs.python.org/3.13.0/reference/compound_stmts.html#grammar-token-python-grammar-while_stmt
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
// https://docs.python.org/3.13.0/reference/compound_stmts.html#grammar-token-python-grammar-for_stmt
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
// https://docs.python.org/3.13.0/reference/compound_stmts.html#grammar-token-python-grammar-try_stmt
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
// https://docs.python.org/3.13.0/reference/compound_stmts.html#grammar-token-python-grammar-try1_stmt
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
// https://docs.python.org/3.13.0/reference/compound_stmts.html#grammar-token-python-grammar-with_stmt
type WithStatement struct {
	Async    bool
	Contents WithStatementContents
	Suite    Suite
	Tokens   Tokens
}

func (w *WithStatement) parse(p *pyParser, async bool) error {
	p.AcceptToken(parser.Token{Type: TokenKeyword, Data: "with"})
	p.AcceptRunWhitespace()

	w.Async = async
	parens := p.AcceptToken(parser.Token{Type: TokenDelimiter, Data: "("})

	if parens {
		p.AcceptRunWhitespace()
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
// https://docs.python.org/3.13.0/reference/compound_stmts.html#grammar-token-python-grammar-with_stmt_contents
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

		w.Items = append(w.Items, wi)
		q = p.NewGoal()

		q.AcceptRunWhitespace()

		another = q.AcceptToken(parser.Token{Type: TokenDelimiter, Data: ","})

		q.AcceptRunWhitespace()

		if q.Peek() == (parser.Token{Type: TokenDelimiter, Data: ")"}) {
			break
		}
	}

	w.Tokens = p.ToTokens()

	return nil
}

// WithItem as defined in python@3.13:
// https://docs.python.org/3.13/reference/compound_stmts.html#grammar-token-python-grammar-with_item
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

	q.AcceptRunWhitespace()

	if q.AcceptToken(parser.Token{Type: TokenKeyword, Data: "as"}) {
		q.AcceptRunWhitespace()
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

// FuncDefinition as defined in python@3.13.0:
// https://docs.python.org/3.13.0/reference/compound_stmts.html#grammar-token-python-grammar-funcdef
type FuncDefinition struct {
	Decorators    *Decorators
	Async         bool
	FuncName      *Token
	TypeParams    *TypeParams
	ParameterList ParameterList
	Expression    *Expression
	Suite         Suite
	Comments      Comments
	Tokens        Tokens
}

func (f *FuncDefinition) parse(p *pyParser, async bool, decorators *Decorators) error {
	f.Decorators = decorators
	f.Async = async

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
		f.Comments = p.AcceptRunWhitespaceComments()

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
// https://docs.python.org/3.13.0/reference/compound_stmts.html#grammar-token-python-grammar-classdef
type ClassDefinition struct {
	Decorators  *Decorators
	ClassName   *Token
	TypeParams  *TypeParams
	Inheritance *ArgumentList
	Suite       Suite
	Tokens      Tokens
}

func (c *ClassDefinition) parse(p *pyParser, decorators *Decorators) error {
	c.Decorators = decorators

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
		p.AcceptRunWhitespace()

		c.Inheritance = new(ArgumentList)

		if !p.AcceptToken(parser.Token{Type: TokenDelimiter, Data: ")"}) {
			q := p.NewGoal()

			if err := c.Inheritance.parse(q); err != nil {
				return p.Error("ClassDefinition", err)
			}

			p.Score(q)
			p.AcceptRunWhitespace()
			p.AcceptToken(parser.Token{Type: TokenDelimiter, Data: ")"})
		} else {
			c.Inheritance.Tokens = p.NewGoal().ToTokens()
		}

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
// https://docs.python.org/3.13.0/reference/compound_stmts.html#grammar-token-python-grammar-suite
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

// TargetList as defined in python@3.13.0:
// https://docs.python.org/3.13.0/reference/simple_stmts.html#grammar-token-python-grammar-target_list
type TargetList struct {
	Targets  []Target
	Comments [2]Comments
	Tokens   Tokens
}

func (t *TargetList) parse(p *pyParser) error {
	t.Comments[0] = p.AcceptRunWhitespaceComments()

	p.AcceptRunAllWhitespace()

Loop:
	for {
		q := p.NewGoal()

		var tg Target

		if err := tg.parse(q); err != nil {
			return p.Error("TargetList", err)
		}

		t.Targets = append(t.Targets, tg)

		p.Score(q)

		q = p.NewGoal()

		q.AcceptRunWhitespace()

		if !q.AcceptToken(parser.Token{Type: TokenDelimiter, Data: ","}) {
			break
		}

		q.AcceptRunWhitespace()

		switch tk := q.Peek(); tk {
		case parser.Token{Type: TokenDelimiter, Data: ";"},
			parser.Token{Type: TokenDelimiter, Data: "="},
			parser.Token{Type: TokenDelimiter, Data: "]"},
			parser.Token{Type: TokenDelimiter, Data: ")"},
			parser.Token{Type: TokenKeyword, Data: "in"}:
			break Loop
		default:
			if tk.Type == TokenLineTerminator || tk.Type == parser.TokenDone {
				break Loop
			}
		}

		q.AcceptRunWhitespace()
		p.Score(q)
	}

	t.Comments[1] = p.AcceptRunWhitespaceComments()
	t.Tokens = p.ToTokens()

	return nil
}

// Target as defined in python@3.13.0:
// https://docs.python.org/3.13.0/reference/simple_stmts.html#grammar-token-python-grammar-target
type Target struct {
	PrimaryExpression *PrimaryExpression
	Tuple             *TargetList
	Array             *TargetList
	Star              *Target
	Tokens            Tokens
}

func (t *Target) parse(p *pyParser) error {
	if p.AcceptToken(parser.Token{Type: TokenDelimiter, Data: "("}) {
		p.AcceptRunWhitespaceNoComment()

		q := p.NewGoal()

		t.Tuple = new(TargetList)

		if err := t.Tuple.parse(q); err != nil {
			return p.Error("Target", err)
		}

		p.Score(q)
		p.AcceptRunWhitespace()

		if !p.AcceptToken(parser.Token{Type: TokenDelimiter, Data: ")"}) {
			return p.Error("Target", ErrMissingClosingParen)
		}

	} else if p.AcceptToken(parser.Token{Type: TokenDelimiter, Data: "["}) {
		p.AcceptRunWhitespaceNoComment()

		q := p.NewGoal()
		t.Array = new(TargetList)

		if err := t.Array.parse(q); err != nil {
			return p.Error("Target", err)
		}

		p.Score(q)
		p.AcceptRunWhitespace()

		if !p.AcceptToken(parser.Token{Type: TokenDelimiter, Data: "]"}) {
			return p.Error("Target", ErrMissingClosingBracket)
		}

	} else if p.AcceptToken(parser.Token{Type: TokenOperator, Data: "*"}) {
		p.AcceptRunWhitespace()

		q := p.NewGoal()
		t.Star = new(Target)

		if err := t.Star.parse(q); err != nil {
			return p.Error("Target", err)
		}

		p.Score(q)
	} else {
		t.PrimaryExpression = new(PrimaryExpression)
		q := p.NewGoal()

		if err := t.PrimaryExpression.parse(q); err != nil {
			return p.Error("Target", err)
		} else if t.PrimaryExpression.Call != nil || t.PrimaryExpression.Atom != nil && !t.PrimaryExpression.IsIdentifier() {
			return p.Error("Target", ErrMissingIdentifier)
		}

		p.Score(q)
	}

	t.Tokens = p.ToTokens()

	return nil
}

// StarredList as defined in python@3.12.6:
// https://docs.python.org/release/3.12.6/reference/expressions.html#grammar-token-python-grammar-starred_list
type StarredList struct {
	StarredItems  []StarredItem
	TrailingComma bool
	Tokens        Tokens
}

func (s *StarredList) parse(p *pyParser) error {
Loop:
	for {
		q := p.NewGoal()

		var si StarredItem

		if err := si.parse(q); err != nil {
			return p.Error("StarredList", err)
		}

		p.Score(q)

		s.StarredItems = append(s.StarredItems, si)
		q = p.NewGoal()

		q.AcceptRunWhitespace()

		hasComma := q.AcceptToken(parser.Token{Type: TokenDelimiter, Data: ","})

		r := q.NewGoal()

		if hasComma {
			r.AcceptRunWhitespace()
		}

		switch tk := r.Peek(); tk {
		case parser.Token{Type: TokenDelimiter, Data: "]"}, parser.Token{Type: TokenDelimiter, Data: "}"}, parser.Token{Type: TokenDelimiter, Data: ")"}, parser.Token{Type: TokenDelimiter, Data: ":"}, parser.Token{Type: TokenKeyword, Data: "for"}, parser.Token{Type: TokenKeyword, Data: "async"}, parser.Token{Type: parser.TokenDone}:
			if hasComma {
				if len(s.StarredItems) == 1 {
					s.TrailingComma = true
				}

				p.Score(q)
			}

			break Loop
		default:
			if tk.Type == TokenLineTerminator || tk.Type == TokenDedent {
				if hasComma {
					if len(s.StarredItems) == 1 {
						s.TrailingComma = true
					}

					p.Score(q)
				}

				break Loop
			}
		}

		q = p.NewGoal()

		q.AcceptRunWhitespace()

		if !q.AcceptToken(parser.Token{Type: TokenDelimiter, Data: ","}) {
			return q.Error("StarredList", ErrMissingComma)
		}

		q.AcceptRunWhitespaceNoComment()
		p.Score(q)
	}

	s.Tokens = p.ToTokens()

	return nil
}

// StarredItem as defined in python@3.12.6:
// https://docs.python.org/release/3.12.6/reference/expressions.html#grammar-token-python-grammar-starred_item
type StarredItem struct {
	AssignmentExpression *AssignmentExpression
	OrExpr               *OrExpression
	Comments             [3]Comments
	Tokens               Tokens
}

func (s *StarredItem) parse(p *pyParser) error {
	s.Comments[0] = p.AcceptRunWhitespaceCommentsIfMultiline()
	p.AcceptRunWhitespace()

	if p.AcceptToken(parser.Token{Type: TokenOperator, Data: "*"}) {
		s.Comments[1] = p.AcceptRunWhitespaceCommentsIfMultiline()
		p.AcceptRunWhitespace()

		q := p.NewGoal()
		s.OrExpr = new(OrExpression)

		if err := s.OrExpr.parse(q); err != nil {
			return p.Error("StarredItem", err)
		}

		p.Score(q)
	} else {
		q := p.NewGoal()
		s.AssignmentExpression = new(AssignmentExpression)

		if err := s.AssignmentExpression.parse(q); err != nil {
			return p.Error("StarredItem", err)
		}

		p.Score(q)
	}

	q := p.NewGoal()

	q.AcceptRunAllWhitespace()

	if q.Peek() == (parser.Token{Type: TokenDelimiter, Data: ","}) {
		s.Comments[2] = p.AcceptRunWhitespaceCommentsIfMultiline()
	} else {
		s.Comments[2] = p.AcceptRunWhitespaceCommentsNoNewlineIfMultiline()
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
type TypeParam struct {
	Type       TypeParamType
	Identifier *Token
	Expression *Expression
	Comments   [2]Comments
	Tokens     Tokens
}

func (t *TypeParam) parse(p *pyParser) error {
	t.Comments[0] = p.AcceptRunWhitespaceComments()

	p.AcceptRunAllWhitespace()

	if p.AcceptToken(parser.Token{Type: TokenOperator, Data: "*"}) {
		t.Type = TypeParamVar

		p.AcceptRunWhitespace()
	} else if p.AcceptToken(parser.Token{Type: TokenOperator, Data: "**"}) {
		t.Type = TypeParamVarTuple

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
			p.Score(q)

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
		t.Comments[1] = p.AcceptRunWhitespaceCommentsNoNewline()
	} else {
		t.Comments[1] = p.AcceptRunWhitespaceComments()
	}

	t.Tokens = p.ToTokens()

	return nil
}

// ParameterList as defined in python@3.13.0:
// https://docs.python.org/release/3.13.0/reference/compound_stmts.html#grammar-token-python-grammar-parameter_list
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
type DefParameter struct {
	Parameter Parameter
	Value     *Expression
	Comments  [2]Comments
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
		q.AcceptRunWhitespace()
		p.Score(q)

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
		d.Comments[1] = p.AcceptRunWhitespaceCommentsIfMultiline()
	} else {
		d.Comments[1] = p.AcceptRunWhitespaceCommentsNoNewlineIfMultiline()
	}

	d.Tokens = p.ToTokens()

	return nil
}

// Parameter as defined in python@3.13.0:
// https://docs.python.org/release/3.13.0/reference/compound_stmts.html#grammar-token-python-grammar-parameter
type Parameter struct {
	Identifier *Token
	Type       *Expression
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
		q.AcceptRunWhitespace()
		p.Score(q)

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

type ArgumentList struct {
	PositionalArguments        []PositionalArgument
	StarredAndKeywordArguments []StarredOrKeyword
	KeywordArguments           []KeywordArgument
	Tokens                     Tokens
}

// ArgumentList as defined in python@3.13.0:
// https://docs.python.org/release/3.13.0/reference/expressions.html#grammar-token-python-grammar-argument_list
func (a *ArgumentList) parse(p *pyParser) error {
	var nextIsKeywordItem, nextIsDoubleStarred bool

	for {
		q := p.NewGoal()

		q.AcceptRunWhitespace()

		if q.AcceptToken(parser.Token{Type: TokenDelimiter, Data: ")"}) || q.Peek().Type == parser.TokenDone {
			break
		} else if q.AcceptToken(parser.Token{Type: TokenOperator, Data: "**"}) {
			nextIsDoubleStarred = true

			break
		} else if q.Accept(TokenIdentifier) {
			q.AcceptRunWhitespace()

			if q.AcceptToken(parser.Token{Type: TokenDelimiter, Data: "="}) {
				nextIsKeywordItem = true

				break
			}
		}

		p.AcceptRunWhitespaceNoComment()

		q = p.NewGoal()

		var pa PositionalArgument

		if err := pa.parse(q); err != nil {
			return p.Error("ArgumentList", err)
		}

		p.Score(q)

		a.PositionalArguments = append(a.PositionalArguments, pa)
		q = p.NewGoal()

		q.AcceptRunWhitespace()

		if tk := q.Peek(); tk == (parser.Token{Type: TokenDelimiter, Data: ")"}) || tk.Type == parser.TokenDone {
			break
		} else if !q.AcceptToken(parser.Token{Type: TokenDelimiter, Data: ","}) {
			return p.Error("ArgumentList", ErrMissingComma)
		}

		p.Score(q)
	}

	if nextIsKeywordItem {
		for {
			q := p.NewGoal()

			q.AcceptRunWhitespace()

			if q.AcceptToken(parser.Token{Type: TokenOperator, Data: "**"}) {
				nextIsDoubleStarred = true

				break
			}

			p.AcceptRunAllWhitespaceNoComment()

			q = p.NewGoal()

			var sk StarredOrKeyword

			if err := sk.parse(q); err != nil {
				return p.Error("ArgumentList", err)
			}

			p.Score(q)

			a.StarredAndKeywordArguments = append(a.StarredAndKeywordArguments, sk)
			q = p.NewGoal()

			q.AcceptRunWhitespace()

			if tk := q.Peek(); tk == (parser.Token{Type: TokenDelimiter, Data: ")"}) || tk.Type == parser.TokenDone {
				break
			} else if !q.AcceptToken(parser.Token{Type: TokenDelimiter, Data: ","}) {
				return p.Error("ArgumentList", ErrMissingComma)
			}

			p.Score(q)

			q = p.NewGoal()

			q.AcceptRunWhitespace()

			if tk := q.Peek(); tk == (parser.Token{Type: TokenDelimiter, Data: ")"}) || tk.Type == parser.TokenDone {
				break
			}
		}
	}

	if nextIsDoubleStarred {
		for {
			p.AcceptRunAllWhitespaceNoComment()

			q := p.NewGoal()

			var ka KeywordArgument

			if err := ka.parse(q); err != nil {
				return p.Error("ArgumentList", err)
			}

			p.Score(q)

			a.KeywordArguments = append(a.KeywordArguments, ka)
			q = p.NewGoal()

			q.AcceptRunWhitespace()

			if tk := q.Peek(); tk == (parser.Token{Type: TokenDelimiter, Data: ")"}) || tk.Type == parser.TokenDone {
				break
			} else if !q.AcceptToken(parser.Token{Type: TokenDelimiter, Data: ","}) {
				return p.Error("ArgumentList", ErrMissingComma)
			}

			q.AcceptRunWhitespace()

			if tk := q.Peek(); tk == (parser.Token{Type: TokenDelimiter, Data: ")"}) || tk.Type == parser.TokenDone {
				break
			}

			p.Score(q)
		}
	}

	a.Tokens = p.ToTokens()

	return nil
}

// PositionalArgument as defined in python@3.13.0:
// https://docs.python.org/release/3.13.0/reference/expressions.html#grammar-token-python-grammar-positional_arguments
type PositionalArgument struct {
	AssignmentExpression *AssignmentExpression
	Expression           *Expression
	Comments             [3]Comments
	Tokens               Tokens
}

func (pa *PositionalArgument) parse(p *pyParser) error {
	pa.Comments[0] = p.AcceptRunWhitespaceComments()

	p.AcceptRunWhitespace()

	if p.AcceptToken(parser.Token{Type: TokenOperator, Data: "*"}) {
		pa.Comments[1] = p.AcceptRunWhitespaceComments()

		p.AcceptRunWhitespace()

		q := p.NewGoal()
		pa.Expression = new(Expression)

		if err := pa.Expression.parse(q); err != nil {
			return p.Error("PositionalArgument", err)
		}

		p.Score(q)
	} else {
		q := p.NewGoal()
		pa.AssignmentExpression = new(AssignmentExpression)

		if err := pa.AssignmentExpression.parse(q); err != nil {
			return p.Error("PositionalArgument", err)
		}

		p.Score(q)
	}

	q := p.NewGoal()

	q.AcceptRunWhitespace()

	if q.AcceptToken(parser.Token{Type: TokenDelimiter, Data: ","}) {
		pa.Comments[2] = p.AcceptRunWhitespaceComments()
	} else {
		pa.Comments[2] = p.AcceptRunWhitespaceCommentsNoNewline()
	}

	pa.Tokens = p.ToTokens()

	return nil
}

// StarredOrKeyword as defined in python@3.13.0:
// https://docs.python.org/release/3.13.0/reference/expressions.html#grammar-token-python-grammar-starred_and_keywords
type StarredOrKeyword struct {
	Expression  *Expression
	KeywordItem *KeywordItem
	Comments    [2]Comments
	Tokens      Tokens
}

func (s *StarredOrKeyword) parse(p *pyParser) error {
	s.Comments[0] = p.AcceptRunWhitespaceComments()

	p.AcceptRunWhitespace()

	if p.AcceptToken(parser.Token{Type: TokenOperator, Data: "*"}) {
		p.AcceptRunWhitespaceNoComment()

		q := p.NewGoal()
		s.Expression = new(Expression)

		if err := s.Expression.parse(q); err != nil {
			return p.Error("StarredOrKeyword", err)
		}

		p.Score(q)
	} else {
		q := p.NewGoal()
		s.KeywordItem = new(KeywordItem)

		if err := s.KeywordItem.parse(q); err != nil {
			return p.Error("StarredOrKeyword", err)
		}

		p.Score(q)
	}

	q := p.NewGoal()

	q.AcceptRunWhitespace()

	if q.AcceptToken(parser.Token{Type: TokenDelimiter, Data: ","}) {
		s.Comments[1] = p.AcceptRunWhitespaceComments()
	} else {
		s.Comments[1] = p.AcceptRunWhitespaceCommentsNoNewline()
	}

	s.Tokens = p.ToTokens()

	return nil
}

// KeywordItem as defined in python@3.13.0:
// https://docs.python.org/release/3.13.0/reference/expressions.html#grammar-token-python-grammar-keyword_item
type KeywordItem struct {
	Identifier *Token
	Expression Expression
	Tokens     Tokens
}

func (k *KeywordItem) parse(p *pyParser) error {
	if !p.Accept(TokenIdentifier) {
		return p.Error("KeywordItem", ErrMissingIdentifier)
	}

	k.Identifier = p.GetLastToken()

	p.AcceptRunWhitespace()

	if !p.AcceptToken(parser.Token{Type: TokenDelimiter, Data: "="}) {
		return p.Error("KeywordItem", ErrMissingEquals)
	}

	p.AcceptRunWhitespace()

	q := p.NewGoal()

	if err := k.Expression.parse(q); err != nil {
		return p.Error("KeywordItem", err)
	}

	p.Score(q)

	k.Tokens = p.ToTokens()

	return nil
}

// KeywordArgument as defined in python@3.13.0:
// https://docs.python.org/release/3.13.0/reference/expressions.html#grammar-token-python-grammar-keywords_arguments
type KeywordArgument struct {
	KeywordItem *KeywordItem
	Expression  *Expression
	Comments    [3]Comments
	Tokens      Tokens
}

func (k *KeywordArgument) parse(p *pyParser) error {
	k.Comments[0] = p.AcceptRunWhitespaceComments()

	p.AcceptRunWhitespace()

	if p.AcceptToken(parser.Token{Type: TokenOperator, Data: "**"}) {
		k.Comments[1] = p.AcceptRunWhitespaceComments()

		p.AcceptRunWhitespace()

		q := p.NewGoal()
		k.Expression = new(Expression)

		if err := k.Expression.parse(q); err != nil {
			return p.Error("KeywordArgument", err)
		}

		p.Score(q)
	} else {
		q := p.NewGoal()
		k.KeywordItem = new(KeywordItem)

		if err := k.KeywordItem.parse(q); err != nil {
			return p.Error("KeywordArgument", err)
		}

		p.Score(q)
	}

	q := p.NewGoal()

	q.AcceptRunWhitespace()

	if q.Peek() == (parser.Token{Type: TokenDelimiter, Data: ","}) {
		k.Comments[2] = p.AcceptRunWhitespaceCommentsIfMultiline()
	} else {
		k.Comments[2] = p.AcceptRunWhitespaceCommentsNoNewlineIfMultiline()
	}

	k.Tokens = p.ToTokens()

	return nil
}

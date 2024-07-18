package python

import "errors"

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

type IfStatement struct{}

func (i *IfStatement) parse(_ *pyParser) error {
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

type Decorators struct{}

func (d *Decorators) parse(_ *pyParser) error {
	return nil
}

var ErrInvalidCompound = errors.New("invalid compound statement")

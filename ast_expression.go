package python

import "vimagination.zapto.org/parser"

type PrimaryExpression struct {
	PrimaryExpression *PrimaryExpression
	Atom              *Atom
	AttributeRef      *Token
	Slicing           *SliceList
	Call              *ArgumentListOrComprehension
	Tokens            Tokens
}

func (pr *PrimaryExpression) parse(p *pyParser) error {
	pr.Atom = new(Atom)

	q := p.NewGoal()

	if err := pr.Atom.parse(q); err != nil {
		return p.Error("Primary", err)
	}

	p.Score(q)

	q = p.NewGoal()

	for {
		q.AcceptRunWhitespace()

		if q.AcceptToken(parser.Token{Type: TokenDelimiter, Data: "."}) {
			if !q.Accept(TokenIdentifier) {
				return q.Error("Primary", ErrMissingIdentifier)
			}

			pr.Tokens = p.ToTokens()
			pr = &PrimaryExpression{
				PrimaryExpression: pr,
				AttributeRef:      q.GetLastToken(),
			}
		} else if q.AcceptToken(parser.Token{Type: TokenDelimiter, Data: "["}) {
			q.OpenBrackets()
			var sl SliceList

			q.AcceptRunWhitespace()

			r := q.NewGoal()

			if err := sl.parse(r); err != nil {
				return q.Error("Primary", err)
			}

			q.Score(r)

			pr.Tokens = p.ToTokens()
			pr = &PrimaryExpression{
				PrimaryExpression: pr,
				Slicing:           &sl,
			}

			q.AcceptRunWhitespace()

			if !q.AcceptToken(parser.Token{Type: TokenDelimiter, Data: "]"}) {
				return q.Error("Primary", ErrMissingClosingBracket)
			}

			q.CloseBrackets()
		} else if q.AcceptToken(parser.Token{Type: TokenDelimiter, Data: "("}) {
			var call ArgumentListOrComprehension

			q.OpenBrackets()
			q.AcceptRunWhitespace()

			r := q.NewGoal()

			if err := call.parse(r); err != nil {
				return q.Error("Primary", err)
			}

			q.Score(r)

			pr.Tokens = p.ToTokens()
			pr = &PrimaryExpression{
				PrimaryExpression: pr,
				Call:              &call,
			}

			q.AcceptRunWhitespace()

			if !q.AcceptToken(parser.Token{Type: TokenDelimiter, Data: ")"}) {
				return q.Error("Primary", ErrMissingClosingParen)
			}

			q.CloseBrackets()
		} else {
			break
		}

		p.Score(q)

		q = p.NewGoal()
	}

	pr.Tokens = p.ToTokens()

	return nil
}

func (pr *PrimaryExpression) IsIdentifier() bool {
	if pr.Atom != nil {
		return pr.Atom.IsIdentifier()
	}

	if pr.PrimaryExpression != nil {
		return pr.PrimaryExpression.IsIdentifier()
	}

	return false
}

type Atom struct {
	Identifier *Token
	Literal    *Token
	Enclosure  *Enclosure
	Tokens     Tokens
}

func (a *Atom) parse(p *pyParser) error {
	if p.Accept(TokenIdentifier) {
		a.Identifier = p.GetLastToken()
	} else if p.Accept(TokenNumericLiteral, TokenStringLiteral) {
		a.Literal = p.GetLastToken()
	} else {
		a.Enclosure = new(Enclosure)

		q := p.NewGoal()

		if err := a.Enclosure.parse(q); err != nil {
			return p.Error("Atom", err)
		}

		p.Score(q)
	}

	a.Tokens = p.ToTokens()

	return nil
}

func (a *Atom) IsIdentifier() bool {
	return a.Identifier != nil
}

type Enclosure struct{}

func (e *Enclosure) parse(_ *pyParser) error {
	return nil
}

type ArgumentListOrComprehension struct{}

func (a *ArgumentListOrComprehension) parse(_ *pyParser) error {
	return nil
}

type ExpressionList struct {
	Expressions []Expression
	Tokens      Tokens
}

func (e *ExpressionList) parse(p *pyParser) error {
	for {
		q := p.NewGoal()

		var ex Expression

		if err := ex.parse(q); err != nil {
			return p.Error("ExpressionList", err)
		}

		p.Score(q)

		e.Expressions = append(e.Expressions, ex)
		q = p.NewGoal()

		q.AcceptRunWhitespace()

		if q.Peek() == (parser.Token{Type: TokenDelimiter, Data: "]"}) {
			break
		} else if !q.AcceptToken(parser.Token{Type: TokenDelimiter, Data: ","}) {
			return p.Error("ExpressionList", ErrMissingComma)
		}

		q.AcceptRunWhitespace()

		if q.Peek() == (parser.Token{Type: TokenDelimiter, Data: "]"}) {
			break
		}

		p.Score(q)
	}

	e.Tokens = p.ToTokens()

	return nil
}

type SliceList struct {
	SliceItems []SliceItem
	Tokens     Tokens
}

func (s *SliceList) parse(p *pyParser) error {
	for {
		q := p.NewGoal()

		var si SliceItem

		if err := si.parse(q); err != nil {
			return p.Error("ExpressionList", err)
		}

		p.Score(q)

		s.SliceItems = append(s.SliceItems, si)
		q = p.NewGoal()

		q.AcceptRunWhitespace()

		if q.Peek() == (parser.Token{Type: TokenDelimiter, Data: "]"}) {
			break
		} else if !q.AcceptToken(parser.Token{Type: TokenDelimiter, Data: ","}) {
			return p.Error("SliceList", ErrMissingComma)
		}

		q.AcceptRunWhitespace()

		if q.Peek() == (parser.Token{Type: TokenDelimiter, Data: "]"}) {
			break
		}

		p.Score(q)
	}

	s.Tokens = p.ToTokens()

	return nil
}

type SliceItem struct {
	Expression *Expression
	LowerBound *Expression
	UpperBound *Expression
	Stride     *Expression
	Tokens     Tokens
}

func (s *SliceItem) parse(p *pyParser) error {
	q := p.NewGoal()

	s.Expression = new(Expression)

	if err := s.Expression.parse(q); err != nil {
		return p.Error("SliceItem", err)
	}

	p.Score(q)

	q = p.NewGoal()

	q.AcceptRunWhitespace()

	if q.AcceptToken(parser.Token{Type: TokenDelimiter, Data: ":"}) {
		q.AcceptRunWhitespace()
		p.Score(q)

		q = p.NewGoal()
		s.LowerBound = s.Expression
		s.Expression = nil
		s.UpperBound = new(Expression)

		if err := s.UpperBound.parse(q); err != nil {
			return p.Error("SliceItem", err)
		}

		p.Score(q)

		q = p.NewGoal()

		q.AcceptRunWhitespace()

		if q.AcceptToken(parser.Token{Type: TokenDelimiter, Data: ":"}) {
			q.AcceptRunWhitespace()
			p.Score(q)

			q = p.NewGoal()
			s.Stride = new(Expression)

			if err := s.Stride.parse(q); err != nil {
				return p.Error("SliceItem", err)
			}

			p.Score(q)
		}

	}

	s.Tokens = p.ToTokens()

	return nil
}

type OrExpression struct {
	XorExpressions XorExpression
	OrExpression   *OrExpression
	Tokens         Tokens
}

func (o *OrExpression) parse(p *pyParser) error {
	q := p.NewGoal()

	if err := o.XorExpressions.parse(p); err != nil {
		return p.Error("OrExpression", err)
	}

	p.Score(q)

	q = p.NewGoal()

	q.AcceptRunWhitespace()

	if q.AcceptToken(parser.Token{Type: TokenOperator, Data: "|"}) {
		q.AcceptRunWhitespace()

		p.Score(q)

		q = p.NewGoal()
		o.OrExpression = new(OrExpression)

		if err := o.OrExpression.parse(q); err != nil {
			return p.Error("OrExpression", err)
		}

		p.Score(q)
	}

	o.Tokens = p.ToTokens()

	return nil
}

type XorExpression struct {
	AndExpressions AndExpression
	XorExpression  *XorExpression
	Tokens         Tokens
}

func (x *XorExpression) parse(p *pyParser) error {
	q := p.NewGoal()

	if err := x.AndExpressions.parse(p); err != nil {
		return p.Error("XorExpression", err)
	}

	p.Score(q)

	q = p.NewGoal()

	q.AcceptRunWhitespace()

	if q.AcceptToken(parser.Token{Type: TokenOperator, Data: "^"}) {
		q.AcceptRunWhitespace()

		p.Score(q)

		q = p.NewGoal()
		x.XorExpression = new(XorExpression)

		if err := x.XorExpression.parse(q); err != nil {
			return p.Error("XorExpression", err)
		}

		p.Score(q)
	}

	x.Tokens = p.ToTokens()

	return nil
}

type AndExpression struct {
	ShiftExpression ShiftExpression
	AndExpression   *AndExpression
	Tokens          Tokens
}

func (a *AndExpression) parse(p *pyParser) error {
	q := p.NewGoal()

	if err := a.ShiftExpression.parse(p); err != nil {
		return p.Error("AndExpression", err)
	}

	p.Score(q)

	q = p.NewGoal()

	q.AcceptRunWhitespace()

	if q.AcceptToken(parser.Token{Type: TokenOperator, Data: "&"}) {
		q.AcceptRunWhitespace()

		p.Score(q)

		q = p.NewGoal()
		a.AndExpression = new(AndExpression)

		if err := a.AndExpression.parse(q); err != nil {
			return p.Error("AndExpression", err)
		}

		p.Score(q)
	}

	a.Tokens = p.ToTokens()

	return nil
}

type ShiftExpression struct {
	AddExpression   AddExpression
	Shift           *Token
	ShiftExpression *ShiftExpression
	Tokens          Tokens
}

func (s *ShiftExpression) parse(p *pyParser) error {
	q := p.NewGoal()

	if err := s.AddExpression.parse(p); err != nil {
		return p.Error("ShiftExpression", err)
	}

	p.Score(q)

	q = p.NewGoal()

	q.AcceptRunWhitespace()

	if q.AcceptToken(parser.Token{Type: TokenOperator, Data: "<<"}) || q.AcceptToken(parser.Token{Type: TokenOperator, Data: ">>"}) {
		s.Shift = q.GetLastToken()

		q.AcceptRunWhitespace()

		p.Score(q)

		q = p.NewGoal()
		s.ShiftExpression = new(ShiftExpression)

		if err := s.ShiftExpression.parse(q); err != nil {
			return p.Error("ShiftExpression", err)
		}

		p.Score(q)
	}

	s.Tokens = p.ToTokens()

	return nil
}

type AddExpression struct {
	MultiplyExpression MultiplyExpression
	Add                *Token
	AddExpression      *AddExpression
	Tokens             Tokens
}

func (a *AddExpression) parse(p *pyParser) error {
	q := p.NewGoal()

	if err := a.MultiplyExpression.parse(p); err != nil {
		return p.Error("AddExpression", err)
	}

	p.Score(q)

	q = p.NewGoal()

	q.AcceptRunWhitespace()

	if q.AcceptToken(parser.Token{Type: TokenOperator, Data: "+"}) || q.AcceptToken(parser.Token{Type: TokenOperator, Data: "-"}) {
		a.Add = q.GetLastToken()

		q.AcceptRunWhitespace()

		p.Score(q)

		q = p.NewGoal()
		a.AddExpression = new(AddExpression)

		if err := a.AddExpression.parse(q); err != nil {
			return p.Error("AddExpression", err)
		}

		p.Score(q)
	}

	a.Tokens = p.ToTokens()

	return nil
}

type MultiplyExpression struct {
	UnaryExpression    UnaryExpression
	Multiply           *Token
	MultiplyExpression *MultiplyExpression
	Tokens             Tokens
}

func (m *MultiplyExpression) parse(p *pyParser) error {
	q := p.NewGoal()

	if err := m.UnaryExpression.parse(p); err != nil {
		return p.Error("MultiplyExpression", err)
	}

	p.Score(q)

	q = p.NewGoal()

	q.AcceptRunWhitespace()

	if q.AcceptToken(parser.Token{Type: TokenOperator, Data: "*"}) || q.AcceptToken(parser.Token{Type: TokenOperator, Data: "@"}) ||
		q.AcceptToken(parser.Token{Type: TokenOperator, Data: "//"}) || q.AcceptToken(parser.Token{Type: TokenOperator, Data: "/"}) ||
		q.AcceptToken(parser.Token{Type: TokenOperator, Data: "%"}) {
		m.Multiply = q.GetLastToken()

		q.AcceptRunWhitespace()

		p.Score(q)

		q = p.NewGoal()
		m.MultiplyExpression = new(MultiplyExpression)

		if err := m.MultiplyExpression.parse(q); err != nil {
			return p.Error("MultiplyExpression", err)
		}

		p.Score(q)
	}

	m.Tokens = p.ToTokens()

	return nil
}

type UnaryExpression struct {
	PowerExpression *PowerExpression
	Unary           *Token
	UnaryExpression *UnaryExpression
	Tokens          Tokens
}

func (u *UnaryExpression) parse(p *pyParser) error {
	if p.AcceptToken(parser.Token{Type: TokenOperator, Data: "-"}) ||
		p.AcceptToken(parser.Token{Type: TokenOperator, Data: "+"}) ||
		p.AcceptToken(parser.Token{Type: TokenOperator, Data: "~"}) {
		u.Unary = p.GetLastToken()

		p.AcceptRunWhitespace()

		q := p.NewGoal()
		u.UnaryExpression = new(UnaryExpression)

		if err := u.UnaryExpression.parse(q); err != nil {
			return p.Error("UnaryExpression", err)
		}

		p.Score(q)
	} else {
		q := p.NewGoal()
		u.PowerExpression = new(PowerExpression)

		if err := u.PowerExpression.parse(q); err != nil {
			return p.Error("UnaryExpression", err)
		}

		p.Score(q)
	}

	u.Tokens = p.ToTokens()

	return nil
}

type PowerExpression struct {
	AwaitExpression   bool
	PrimaryExpression PrimaryExpression
	UnaryExpression   *UnaryExpression
	Tokens            Tokens
}

func (pe *PowerExpression) parse(p *pyParser) error {
	if p.AcceptToken(parser.Token{Type: TokenKeyword, Data: "await"}) {
		pe.AwaitExpression = true

		p.AcceptRunWhitespace()
	}

	q := p.NewGoal()

	if err := pe.PrimaryExpression.parse(q); err != nil {
		return p.Error("PowerExpression", err)
	}

	p.Score(q)

	q = p.NewGoal()

	q.AcceptRunWhitespace()

	if q.AcceptToken(parser.Token{Type: TokenOperator, Data: "**"}) {
		q.AcceptRunWhitespace()
		p.Score(q)

		q = p.NewGoal()
		pe.UnaryExpression = new(UnaryExpression)

		if err := pe.UnaryExpression.parse(q); err != nil {
			return p.Error("PowerExpression", err)
		}

		p.Score(q)
	}

	pe.Tokens = p.ToTokens()

	return nil
}

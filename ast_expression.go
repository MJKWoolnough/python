package python

import "vimagination.zapto.org/parser"

type Primary struct {
	Primary      *Primary
	Atom         *Atom
	AttributeRef *Token
	Slicing      *SliceList
	Call         *ArgumentListOrComprehension
	Tokens       Tokens
}

func (pr *Primary) parse(p *pyParser) error {
	pr.Atom = new(Atom)

	q := p.NewGoal()

	if err := pr.Atom.parse(q); err != nil {
		return p.Error("Primary", err)
	}

	p.Score(q)

	q = p.NewGoal()

	for {
		q.AcceptRun(TokenWhitespace)

		if q.AcceptToken(parser.Token{Type: TokenDelimiter, Data: "."}) {
			if !q.Accept(TokenIdentifier) {
				return q.Error("Primary", ErrMissingIdentifier)
			}

			pr.Tokens = p.ToTokens()
			pr = &Primary{
				Primary:      pr,
				AttributeRef: q.GetLastToken(),
			}
		} else if q.AcceptToken(parser.Token{Type: TokenDelimiter, Data: "["}) {
			var sl SliceList

			q.AcceptRun(TokenWhitespace, TokenComment)

			r := q.NewGoal()

			if err := sl.parse(r); err != nil {
				return q.Error("Primary", err)
			}

			q.Score(r)

			pr.Tokens = p.ToTokens()
			pr = &Primary{
				Primary: pr,
				Slicing: &sl,
			}

			q.AcceptRun(TokenWhitespace, TokenComment)

			if !q.AcceptToken(parser.Token{Type: TokenDelimiter, Data: "]"}) {
				return q.Error("Primary", ErrMissingClosingBracket)
			}
		} else if q.AcceptToken(parser.Token{Type: TokenDelimiter, Data: "("}) {
			var call ArgumentListOrComprehension

			q.AcceptRun(TokenWhitespace, TokenComment)

			r := q.NewGoal()

			if err := call.parse(r); err != nil {
				return q.Error("Primary", err)
			}

			q.Score(r)

			pr.Tokens = p.ToTokens()
			pr = &Primary{
				Primary: pr,
				Call:    &call,
			}

			q.AcceptRun(TokenWhitespace, TokenComment)

			if !q.AcceptToken(parser.Token{Type: TokenDelimiter, Data: ")"}) {
				return q.Error("Primary", ErrMissingClosingParen)
			}
		} else {
			break
		}

		p.Score(q)

		q = p.NewGoal()
	}

	pr.Tokens = p.ToTokens()

	return nil
}

func (pr *Primary) IsIdentifier() bool {
	if pr.Atom != nil {
		return pr.Atom.IsIdentifier()
	}

	if pr.Primary != nil {
		return pr.Primary.IsIdentifier()
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

type ExpressionList struct{}

func (e *ExpressionList) parse(_ *pyParser) error {
	return nil
}

type SliceList struct{}

func (s *SliceList) parse(_ *pyParser) error {
	return nil
}

type OrExpression struct {
	XorExpressions XorExpression
	OrExpression   *OrExpression
	Tokens         Tokens
}

func (o *OrExpression) parse(p *pyParser, ws []parser.TokenType) error {
	q := p.NewGoal()

	if err := o.XorExpressions.parse(p, ws); err != nil {
		return p.Error("OrExpression", err)
	}

	p.Score(q)

	q = p.NewGoal()

	q.AcceptRun(ws...)

	if q.AcceptToken(parser.Token{Type: TokenOperator, Data: "|"}) {
		q.AcceptRun(ws...)

		p.Score(q)

		q = p.NewGoal()
		o.OrExpression = new(OrExpression)

		if err := o.OrExpression.parse(q, ws); err != nil {
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

func (x *XorExpression) parse(p *pyParser, ws []parser.TokenType) error {
	q := p.NewGoal()

	if err := x.AndExpressions.parse(p, ws); err != nil {
		return p.Error("XorExpression", err)
	}

	p.Score(q)

	q = p.NewGoal()

	q.AcceptRun(ws...)

	if q.AcceptToken(parser.Token{Type: TokenOperator, Data: "^"}) {
		q.AcceptRun(ws...)

		p.Score(q)

		q = p.NewGoal()
		x.XorExpression = new(XorExpression)

		if err := x.XorExpression.parse(q, ws); err != nil {
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

func (a *AndExpression) parse(p *pyParser, ws []parser.TokenType) error {
	q := p.NewGoal()

	if err := a.ShiftExpression.parse(p, ws); err != nil {
		return p.Error("AndExpression", err)
	}

	p.Score(q)

	q = p.NewGoal()

	q.AcceptRun(ws...)

	if q.AcceptToken(parser.Token{Type: TokenOperator, Data: "&"}) {
		q.AcceptRun(ws...)

		p.Score(q)

		q = p.NewGoal()
		a.AndExpression = new(AndExpression)

		if err := a.AndExpression.parse(q, ws); err != nil {
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

func (s *ShiftExpression) parse(p *pyParser, ws []parser.TokenType) error {
	q := p.NewGoal()

	if err := s.AddExpression.parse(p, ws); err != nil {
		return p.Error("ShiftExpression", err)
	}

	p.Score(q)

	q = p.NewGoal()

	q.AcceptRun(ws...)

	if q.AcceptToken(parser.Token{Type: TokenOperator, Data: "<<"}) || q.AcceptToken(parser.Token{Type: TokenOperator, Data: ">>"}) {
		s.Shift = q.GetLastToken()

		q.AcceptRun(ws...)

		p.Score(q)

		q = p.NewGoal()
		s.ShiftExpression = new(ShiftExpression)

		if err := s.ShiftExpression.parse(q, ws); err != nil {
			return p.Error("AndExpression", err)
		}

		p.Score(q)
	}

	s.Tokens = p.ToTokens()

	return nil
}

type AddExpression struct{}

func (a *AddExpression) parse(_ *pyParser, _ []parser.TokenType) error {
	return nil
}

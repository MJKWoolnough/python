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
	return false
}

type Atom struct{}

func (a *Atom) parse(_ *pyParser) error {
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

type OrExpr struct{}

func (o *OrExpr) parse(_ *pyParser) error {
	return nil
}

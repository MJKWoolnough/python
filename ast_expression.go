package python

import (
	"vimagination.zapto.org/parser"
)

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

	for {
		q := p.NewGoal()

		q.AcceptRunWhitespace()

		if q.AcceptToken(parser.Token{Type: TokenDelimiter, Data: "."}) {
			if !q.Accept(TokenIdentifier) {
				return q.Error("Primary", ErrMissingIdentifier)
			}

			pr.Tokens = p.ToTokens()
			ipr := *pr
			*pr = PrimaryExpression{
				PrimaryExpression: &ipr,
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
			ipr := *pr
			*pr = PrimaryExpression{
				PrimaryExpression: &ipr,
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
			ipr := *pr
			*pr = PrimaryExpression{
				PrimaryExpression: &ipr,
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
	}

	pr.Tokens = p.ToTokens()

	return nil
}

func (pr *PrimaryExpression) IsIdentifier() bool {
	if pr.Atom != nil {
		return pr.Atom.IsIdentifier()
	} else if pr.PrimaryExpression != nil {
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

type Enclosure struct {
	ParenthForm         *StarredExpression
	ListDisplay         *StarredListOrComprehension
	DictDisplay         *DictDisplay
	SetDisplay          *StarredListOrComprehension
	GeneratorExpression *GeneratorExpression
	YieldAtom           *YieldExpression
	Tokens              Tokens
}

func (e *Enclosure) parse(p *pyParser) error {
	if p.AcceptToken(parser.Token{Type: TokenDelimiter, Data: "("}) {
		p.AcceptRunWhitespace()

		q := p.NewGoal()

		if q.Peek() == (parser.Token{Type: TokenKeyword, Data: "yield"}) {
			e.YieldAtom = new(YieldExpression)

			if err := e.YieldAtom.parse(q); err != nil {
				return p.Error("Enclosure", err)
			}
		} else if q.LookaheadLine(parser.Token{Type: TokenKeyword, Data: "for"}) == 0 {
			e.GeneratorExpression = new(GeneratorExpression)

			if err := e.GeneratorExpression.parse(q); err != nil {
				return p.Error("Enclosure", err)
			}
		} else {
			e.ParenthForm = new(StarredExpression)

			if err := e.ParenthForm.parse(q); err != nil {
				return p.Error("Enclosure", err)
			}
		}

		p.Score(q)
		p.AcceptRunWhitespace()

		if !p.AcceptToken(parser.Token{Type: TokenDelimiter, Data: ")"}) {
			return p.Error("Enclosure", ErrMissingClosingParen)
		}
	} else if p.AcceptToken(parser.Token{Type: TokenDelimiter, Data: "["}) {
		p.AcceptRunWhitespace()

		q := p.NewGoal()
		e.ListDisplay = new(StarredListOrComprehension)

		if err := e.ListDisplay.parse(q, nil); err != nil {
			return p.Error("Enclosure", err)
		}

		p.Score(q)
		p.AcceptRunWhitespace()

		if !p.AcceptToken(parser.Token{Type: TokenDelimiter, Data: "]"}) {
			return p.Error("Enclosure", ErrMissingClosingBracket)
		}
	} else if p.AcceptToken(parser.Token{Type: TokenDelimiter, Data: "{"}) {
		p.AcceptRunWhitespace()

		q := p.NewGoal()

		var isDict bool
		var ae *AssignmentExpression

		switch q.Peek() {
		case parser.Token{Type: TokenDelimiter, Data: "**"}:
			isDict = true
		case parser.Token{Type: TokenOperator, Data: "*"}:
		default:
			ae = new(AssignmentExpression)

			if err := ae.parse(q); err != nil {
				return p.Error("Enclosure", err)
			}

			if ae.Identifier == nil {
				r := q.NewGoal()

				r.AcceptRunWhitespace()

				isDict = r.AcceptToken(parser.Token{Type: TokenDelimiter, Data: ":"})
			}
		}

		if isDict {
			var ex *Expression

			if ae != nil {
				ex = &ae.Expression
			}

			e.DictDisplay = new(DictDisplay)

			if err := e.DictDisplay.parse(q, ex); err != nil {
				return p.Error("Enclosure", err)
			}
		} else {
			e.SetDisplay = new(StarredListOrComprehension)

			if err := e.SetDisplay.parse(q, ae); err != nil {
				return p.Error("Enclosure", err)
			}
		}

		p.Score(q)
		p.AcceptRunWhitespace()

		if !p.AcceptToken(parser.Token{Type: TokenDelimiter, Data: "}"}) {
			return p.Error("Enclosure", ErrMissingClosingBrace)
		}
	} else {
		return p.Error("Enclosure", ErrInvalidEnclosure)
	}

	e.Tokens = p.ToTokens()

	return nil
}

type StarredListOrComprehension struct {
	StarredList   *StarredList
	Comprehension *Comprehension
	Tokens        Tokens
}

func (s *StarredListOrComprehension) parse(p *pyParser, ae *AssignmentExpression) error {
	if ae == nil {
		q := p.NewGoal()

		s.StarredList = new(StarredList)

		if err := s.StarredList.parse(q); err != nil {
			return p.Error("StarredListOrComprehension", err)
		}

		p.Score(q)
	} else {
		s.StarredList = &StarredList{
			StarredItems: []StarredItem{
				{
					AssignmentExpression: ae,
					Tokens:               ae.Tokens,
				},
			},
			Tokens: ae.Tokens,
		}
	}

	if len(s.StarredList.StarredItems) == 1 && s.StarredList.StarredItems[0].AssignmentExpression != nil {
		q := p.NewGoal()

		q.AcceptRunWhitespace()

		if tk := q.Peek(); tk == (parser.Token{Type: TokenKeyword, Data: "async"}) || tk == (parser.Token{Type: TokenKeyword, Data: "for"}) {
			p.Score(q)

			q = p.NewGoal()
			s.Comprehension = new(Comprehension)

			if err := s.Comprehension.parse(q, s.StarredList.StarredItems[0].AssignmentExpression); err != nil {
				return p.Error("StarredListOrComprehension", err)
			}

			p.Score(q)

			s.StarredList = nil
		}
	}

	s.Tokens = p.ToTokens()

	return nil
}

type Comprehension struct {
	AssignmentExpression AssignmentExpression
	ComprehensionFor     ComprehensionFor
	Tokens               Tokens
}

func (c *Comprehension) parse(p *pyParser, ae *AssignmentExpression) error {
	if ae != nil {
		c.AssignmentExpression = *ae
	} else {
		q := p.NewGoal()

		if err := c.AssignmentExpression.parse(q); err != nil {
			return p.Error("Comprehension", err)
		}

		p.Score(q)
	}

	p.AcceptRunWhitespace()

	q := p.NewGoal()

	if err := c.ComprehensionFor.parse(q); err != nil {
		return p.Error("Comprehension", err)
	}

	p.Score(q)

	c.Tokens = p.ToTokens()

	return nil
}

type ComprehensionFor struct {
	Async                 bool
	TargetList            TargetList
	OrTest                OrTest
	ComprehensionIterator *ComprehensionIterator
	Tokens                Tokens
}

func (c *ComprehensionFor) parse(p *pyParser) error {
	c.Async = p.AcceptToken(parser.Token{Type: TokenKeyword, Data: "async"})

	p.AcceptRunWhitespace()

	if !p.AcceptToken(parser.Token{Type: TokenKeyword, Data: "for"}) {
		return p.Error("ComprehensionFor", ErrMissingFor)
	}

	p.AcceptRunWhitespace()

	q := p.NewGoal()

	if err := c.TargetList.parse(q); err != nil {
		return p.Error("ComprehensionFor", err)
	}

	p.Score(q)
	p.AcceptRunWhitespace()

	if !p.AcceptToken(parser.Token{Type: TokenKeyword, Data: "in"}) {
		return p.Error("ComprehensionFor", ErrMissingIn)
	}

	p.Score(q)
	p.AcceptRunWhitespace()

	q = p.NewGoal()

	if err := c.OrTest.parse(q); err != nil {
		return p.Error("ComprehensionFor", err)
	}

	p.Score(q)
	p.AcceptRunWhitespace()

	q = p.NewGoal()

	q.AcceptRunWhitespace()

	switch q.Peek() {
	case parser.Token{Type: TokenKeyword, Data: "if"}, parser.Token{Type: TokenKeyword, Data: "async"}, parser.Token{Type: TokenKeyword, Data: "for"}:
		p.Score(q)

		q = p.NewGoal()
		c.ComprehensionIterator = new(ComprehensionIterator)

		if err := c.ComprehensionIterator.parse(q); err != nil {
			return p.Error("ComprehensionFor", err)
		}

		p.Score(q)
	}

	c.Tokens = p.ToTokens()

	return nil
}

type ComprehensionIterator struct {
	ComprehensionFor *ComprehensionFor
	ComprehensionIf  *ComprehensionIf
	Tokens           Tokens
}

func (c *ComprehensionIterator) parse(p *pyParser) error {
	q := p.NewGoal()

	if q.Peek() == (parser.Token{Type: TokenKeyword, Data: "if"}) {
		c.ComprehensionIf = new(ComprehensionIf)

		if err := c.ComprehensionIf.parse(q); err != nil {
			return p.Error("ComprehensionIterator", err)
		}
	} else {
		c.ComprehensionFor = new(ComprehensionFor)

		if err := c.ComprehensionFor.parse(q); err != nil {
			return p.Error("ComprehensionIterator", err)
		}
	}

	p.Score(q)

	c.Tokens = p.ToTokens()

	return nil
}

type ComprehensionIf struct {
	OrTest                OrTest
	ComprehensionIterator *ComprehensionIterator
	Tokens                Tokens
}

func (c *ComprehensionIf) parse(p *pyParser) error {
	if !p.AcceptToken(parser.Token{Type: TokenKeyword, Data: "ff"}) {
		return p.Error("ComprehensionIf", ErrMissingIf)
	}

	p.AcceptRunWhitespace()

	q := p.NewGoal()

	if err := c.OrTest.parse(q); err != nil {
		return p.Error("ComprehensionIf", err)
	}

	p.Score(q)
	p.AcceptRunWhitespace()

	q = p.NewGoal()

	q.AcceptRunWhitespace()

	switch q.Peek() {
	case parser.Token{Type: TokenKeyword, Data: "if"}, parser.Token{Type: TokenKeyword, Data: "async"}, parser.Token{Type: TokenKeyword, Data: "for"}:
		p.Score(q)

		q = p.NewGoal()
		c.ComprehensionIterator = new(ComprehensionIterator)

		if err := c.ComprehensionIterator.parse(q); err != nil {
			return p.Error("ComprehensionFor", err)
		}

		p.Score(q)
	}

	c.Tokens = p.ToTokens()

	return nil
}

type DictDisplay struct {
	DictItems         []DictItem
	DictComprehension *ComprehensionFor
	Tokens            Tokens
}

func (d *DictDisplay) parse(p *pyParser, e *Expression) error {
Loop:
	for {
		q := p.NewGoal()

		var di DictItem

		if err := di.parse(q, e); err != nil {
			return p.Error("DictDisplay", err)
		}

		p.Score(q)

		e = nil
		d.DictItems = append(d.DictItems, di)
		q = p.NewGoal()

		q.AcceptRunWhitespace()

		switch q.Peek() {
		case parser.Token{Type: TokenKeyword, Data: "async"}, parser.Token{Type: TokenKeyword, Data: "for"}:
			if len(d.DictItems) > 1 || d.DictItems[0].OrExpression != nil {
				return p.Error("DictDisplay", ErrInvalidKeyword)
			}

			p.Score(q)

			q = p.NewGoal()
			d.DictComprehension = new(ComprehensionFor)

			if err := d.DictComprehension.parse(q); err != nil {
				return p.Error("DictDisplay", err)
			}

			p.Score(q)

			break Loop
		}

		if !q.AcceptToken(parser.Token{Type: TokenDelimiter, Data: ","}) {
			break
		}

		q.AcceptRunWhitespace()

		if q.Peek() == (parser.Token{Type: TokenDelimiter, Data: "}"}) {
			break
		}

		p.Score(q)
	}

	d.Tokens = p.ToTokens()

	return nil
}

type DictItem struct {
	Key          *Expression
	Value        *Expression
	OrExpression *OrExpression
	Tokens       Tokens
}

func (d *DictItem) parse(p *pyParser, e *Expression) error {
	if e == nil && p.AcceptToken(parser.Token{Type: TokenOperator, Data: "**"}) {
		p.AcceptRunWhitespace()

		q := p.NewGoal()
		d.OrExpression = new(OrExpression)

		if err := d.OrExpression.parse(q); err != nil {
			return p.Error("DictItem", err)
		}

		p.Score(q)
	} else {
		if e != nil {
			d.Key = e
		} else {
			q := p.NewGoal()
			d.Key = new(Expression)

			if err := d.Key.parse(q); err != nil {
				return p.Error("DictItem", err)
			}

			p.Score(q)
		}

		p.AcceptRunWhitespace()

		if !p.AcceptToken(parser.Token{Type: TokenDelimiter, Data: ":"}) {
			return p.Error("DictItem", ErrMissingColon)
		}

		p.AcceptRunWhitespace()

		q := p.NewGoal()
		d.Value = new(Expression)

		if err := d.Value.parse(q); err != nil {
			return p.Error("DictItem", err)
		}

		p.Score(q)
	}

	return nil
}

type GeneratorExpression struct {
	Expression       Expression
	ComprehensionFor ComprehensionFor
	Tokens           Tokens
}

func (g *GeneratorExpression) parse(p *pyParser) error {
	q := p.NewGoal()

	if err := g.Expression.parse(q); err != nil {
		return p.Error("GeneratorExpression", err)
	}

	p.Score(q)

	p.AcceptRunWhitespace()

	q = p.NewGoal()

	if err := g.ComprehensionFor.parse(q); err != nil {
		return p.Error("GeneratorExpression", err)
	}

	p.Score(q)

	g.Tokens = p.ToTokens()

	return nil
}

type ArgumentListOrComprehension struct {
	ArgumentList  *ArgumentList
	Comprehension *Comprehension
	Tokens        Tokens
}

func (a *ArgumentListOrComprehension) parse(p *pyParser) error {
	q := p.NewGoal()

	if q.LookaheadLine(parser.Token{Type: TokenKeyword, Data: "for"}) == 0 {
		a.Comprehension = new(Comprehension)

		if err := a.Comprehension.parse(q, nil); err != nil {
			return p.Error("ArgumentListOrComprehension", err)
		}
	} else {
		a.Comprehension = new(Comprehension)

		if err := a.ArgumentList.parse(q); err != nil {
			return p.Error("ArgumentListOrComprehension", err)
		}
	}

	p.Score(q)

	a.Tokens = p.ToTokens()

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
	AndExpression AndExpression
	XorExpression *XorExpression
	Tokens        Tokens
}

func (x *XorExpression) parse(p *pyParser) error {
	q := p.NewGoal()

	if err := x.AndExpression.parse(p); err != nil {
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

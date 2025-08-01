package python

import (
	"vimagination.zapto.org/parser"
)

// PrimaryExpression as defined in python@3.13.0:
// https://docs.python.org/release/3.13.0/reference/expressions.html#grammar-token-python-grammar-primary
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
		return p.Error("PrimaryExpression", err)
	}

	p.Score(q)

	for {
		q := p.NewGoal()

		q.AcceptRunWhitespace()

		if q.AcceptToken(parser.Token{Type: TokenDelimiter, Data: "."}) {
			q.AcceptRunWhitespace()

			if !q.Accept(TokenIdentifier) {
				return q.Error("PrimaryExpression", ErrMissingIdentifier)
			}

			pr.Tokens = p.ToTokens()
			ipr := *pr
			*pr = PrimaryExpression{
				PrimaryExpression: &ipr,
				AttributeRef:      q.GetLastToken(),
			}
		} else if tk := q.Peek(); tk == (parser.Token{Type: TokenDelimiter, Data: "["}) {
			r := q.NewGoal()

			var sl SliceList

			if err := sl.parse(r); err != nil {
				return q.Error("PrimaryExpression", err)
			}

			q.Score(r)

			pr.Tokens = p.ToTokens()
			ipr := *pr
			*pr = PrimaryExpression{
				PrimaryExpression: &ipr,
				Slicing:           &sl,
			}
		} else if tk == (parser.Token{Type: TokenDelimiter, Data: "("}) {
			r := q.NewGoal()

			var call ArgumentListOrComprehension

			if err := call.parse(r); err != nil {
				return q.Error("PrimaryExpression", err)
			}

			q.Score(r)

			pr.Tokens = p.ToTokens()
			ipr := *pr
			*pr = PrimaryExpression{
				PrimaryExpression: &ipr,
				Call:              &call,
			}
		} else {
			break
		}

		p.Score(q)
	}

	pr.Tokens = p.ToTokens()

	return nil
}

// IsIdentifier returns true if the Primary expression is based on an Identifier.
func (pr *PrimaryExpression) IsIdentifier() bool {
	if pr.Atom != nil {
		return pr.Atom.IsIdentifier()
	} else if pr.PrimaryExpression != nil {
		return pr.PrimaryExpression.IsIdentifier()
	}

	return false
}

// Atom as defined in python@3.13.0:
// https://docs.python.org/release/3.13.0/reference/expressions.html#grammar-token-python-grammar-atom
type Atom struct {
	Identifier *Token
	Literal    *Token
	Enclosure  *Enclosure
	Tokens     Tokens
}

func (a *Atom) parse(p *pyParser) error {
	if p.Accept(TokenIdentifier) {
		a.Identifier = p.GetLastToken()
	} else if p.Accept(TokenNumericLiteral, TokenStringLiteral, TokenBooleanLiteral) {
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

// IsIdentifier returns true if the Atom contains an Idneitifer.
func (a *Atom) IsIdentifier() bool {
	return a.Identifier != nil
}

// Enclosure as defined in python@3.13.0:
// https://docs.python.org/release/3.13.0/reference/expressions.html#grammar-token-python-grammar-enclosure
type Enclosure struct {
	ParenthForm         *StarredExpression
	ListDisplay         *FlexibleExpressionListOrComprehension
	DictDisplay         *DictDisplay
	SetDisplay          *FlexibleExpressionListOrComprehension
	GeneratorExpression *GeneratorExpression
	YieldAtom           *YieldExpression
	Comments            [2]Comments
	Tokens              Tokens
}

func (e *Enclosure) parse(p *pyParser) error {
	if p.AcceptToken(parser.Token{Type: TokenDelimiter, Data: "("}) {
		p.OpenBrackets()

		q := p.NewGoal()

		if q.AcceptRunAllWhitespace() == TokenDelimiter {
			e.Comments[0] = p.AcceptRunWhitespaceComments()
		} else {
			e.Comments[0] = p.AcceptRunWhitespaceCommentsNoNewline()
		}

		p.AcceptRunAllWhitespace()

		if p.AcceptToken(parser.Token{Type: TokenDelimiter, Data: ")"}) {
			e.ParenthForm = &StarredExpression{
				Tokens: p.NewGoal().ToTokens(),
			}
		} else {
			q := p.NewGoal()

			q.AcceptRunAllWhitespace()

			if q.Peek() == (parser.Token{Type: TokenKeyword, Data: "yield"}) {
				p.AcceptRunWhitespaceNoComment()

				q = p.NewGoal()
				e.YieldAtom = new(YieldExpression)

				if err := e.YieldAtom.parse(q); err != nil {
					return p.Error("Enclosure", err)
				}
			} else if q.LookaheadLine(parser.Token{Type: TokenKeyword, Data: "for"}) == 0 {
				p.AcceptRunWhitespaceNoComment()

				q = p.NewGoal()
				e.GeneratorExpression = new(GeneratorExpression)

				if err := e.GeneratorExpression.parse(q); err != nil {
					return p.Error("Enclosure", err)
				}
			} else {
				p.AcceptRunWhitespaceNoComment()

				q = p.NewGoal()
				e.ParenthForm = new(StarredExpression)

				if err := e.ParenthForm.parse(q); err != nil {
					return p.Error("Enclosure", err)
				}
			}

			p.Score(q)

			e.Comments[1] = p.AcceptRunWhitespaceComments()

			p.AcceptRunWhitespace()

			if !p.AcceptToken(parser.Token{Type: TokenDelimiter, Data: ")"}) {
				return p.Error("Enclosure", ErrMissingClosingParen)
			}
		}

		p.CloseBrackets()
	} else if p.AcceptToken(parser.Token{Type: TokenDelimiter, Data: "["}) {
		p.OpenBrackets()

		q := p.NewGoal()

		if q.AcceptRunAllWhitespace() == TokenDelimiter {
			e.Comments[0] = p.AcceptRunWhitespaceComments()
		} else {
			e.Comments[0] = p.AcceptRunWhitespaceCommentsNoNewline()
		}

		p.AcceptRunAllWhitespace()

		if p.AcceptToken(parser.Token{Type: TokenDelimiter, Data: "]"}) {
			e.ListDisplay = &FlexibleExpressionListOrComprehension{
				Tokens: p.NewGoal().ToTokens(),
			}
		} else {
			p.AcceptRunWhitespaceNoComment()

			q := p.NewGoal()
			e.ListDisplay = new(FlexibleExpressionListOrComprehension)

			if err := e.ListDisplay.parse(q, nil); err != nil {
				return p.Error("Enclosure", err)
			}

			p.Score(q)

			e.Comments[1] = p.AcceptRunWhitespaceComments()

			p.AcceptRunAllWhitespace()

			if !p.AcceptToken(parser.Token{Type: TokenDelimiter, Data: "]"}) {
				return p.Error("Enclosure", ErrMissingClosingBracket)
			}
		}

		p.CloseBrackets()
	} else if p.AcceptToken(parser.Token{Type: TokenDelimiter, Data: "{"}) {
		p.OpenBrackets()

		q := p.NewGoal()

		if q.AcceptRunAllWhitespace() == TokenDelimiter {
			e.Comments[0] = p.AcceptRunWhitespaceComments()
		} else {
			e.Comments[0] = p.AcceptRunWhitespaceCommentsNoNewline()
		}

		p.AcceptRunAllWhitespace()

		if p.AcceptToken(parser.Token{Type: TokenDelimiter, Data: "}"}) {
			e.DictDisplay = &DictDisplay{
				Tokens: p.NewGoal().ToTokens(),
			}
		} else {
			q := p.NewGoal()

			q.AcceptRunAllWhitespace()

			var isDict bool
			var ae *AssignmentExpression

			switch q.Peek() {
			case parser.Token{Type: TokenOperator, Data: "**"}:
				isDict = true

				fallthrough
			case parser.Token{Type: TokenOperator, Data: "*"}:
				p.AcceptRunWhitespaceNoComment()

				q = p.NewGoal()
			default:
				p.AcceptRunWhitespaceNoComment()

				q = p.NewGoal()
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

				if err := e.DictDisplay.parse(q, p, ex); err != nil {
					return p.Error("Enclosure", err)
				}
			} else {
				e.SetDisplay = new(FlexibleExpressionListOrComprehension)

				if err := e.SetDisplay.parse(q, ae); err != nil {
					return p.Error("Enclosure", err)
				}
			}

			p.Score(q)

			e.Comments[1] = p.AcceptRunWhitespaceComments()

			p.AcceptRunAllWhitespace()

			if !p.AcceptToken(parser.Token{Type: TokenDelimiter, Data: "}"}) {
				return p.Error("Enclosure", ErrMissingClosingBrace)
			}
		}
		p.CloseBrackets()
	} else {
		return p.Error("Enclosure", ErrInvalidEnclosure)
	}

	e.Tokens = p.ToTokens()

	return nil
}

// FlexibleExpressionListOrComprehension as defined in python@3.13.0:
// https://docs.python.org/release/3.13.0/reference/expressions.html#grammar-token-python-grammar-list_display
type FlexibleExpressionListOrComprehension struct {
	FlexibleExpressionList *FlexibleExpressionList
	Comprehension          *Comprehension
	Tokens                 Tokens
}

func (f *FlexibleExpressionListOrComprehension) parse(p *pyParser, ae *AssignmentExpression) error {
	o := p.NewGoal()
	if ae == nil {
		q := p.NewGoal()
		f.FlexibleExpressionList = new(FlexibleExpressionList)

		if err := f.FlexibleExpressionList.parse(q); err != nil {
			return p.Error("FlexibleExpressionListOrComprehension", err)
		}

		p.Score(q)
	} else {
		f.FlexibleExpressionList = &FlexibleExpressionList{
			FlexibleExpressions: []FlexibleExpression{
				{
					AssignmentExpression: ae,
					Tokens:               ae.Tokens,
				},
			},
			Tokens: ae.Tokens,
		}
	}

	if len(f.FlexibleExpressionList.FlexibleExpressions) == 1 && f.FlexibleExpressionList.FlexibleExpressions[0].AssignmentExpression != nil {
		q := p.NewGoal()

		q.AcceptRunWhitespace()

		if tk := q.Peek(); tk == (parser.Token{Type: TokenKeyword, Data: "async"}) || tk == (parser.Token{Type: TokenKeyword, Data: "for"}) {
			p.Score(q)

			f.Comprehension = new(Comprehension)

			if err := f.Comprehension.parse(p, f.FlexibleExpressionList.FlexibleExpressions[0].AssignmentExpression); err != nil {
				return o.Error("FlexibleExpressionListOrComprehension", err)
			}

			f.FlexibleExpressionList = nil
		}
	}

	f.Tokens = p.ToTokens()

	return nil
}

// FlexibleExpressionList as defined in python@3.13.0:
// https://docs.python.org/release/3.13.0/reference/expressions.html#grammar-token-python-grammar-flexible_expression_list
type FlexibleExpressionList struct {
	FlexibleExpressions []FlexibleExpression
	Tokens
}

func (f *FlexibleExpressionList) parse(p *pyParser) error {
	for {
		q := p.NewGoal()

		var fe FlexibleExpression

		if err := fe.parse(q); err != nil {
			return p.Error("FlexibleExpressionList", err)
		}

		f.FlexibleExpressions = append(f.FlexibleExpressions, fe)
		p.Score(q)

		q = p.NewGoal()

		q.AcceptRunWhitespace()

		if !q.AcceptToken(parser.Token{Type: TokenDelimiter, Data: ","}) {
			break
		}

		p.Score(q)

		q = p.NewGoal()

		q.AcceptRunWhitespace()

		if tk := q.Peek(); tk == (parser.Token{Type: TokenDelimiter, Data: "}"}) || tk == (parser.Token{Type: TokenDelimiter, Data: "]"}) || tk == (parser.Token{Type: TokenDelimiter, Data: ")"}) || tk.Type == TokenLineTerminator || tk.Type == parser.TokenDone {
			break
		}

		p.Score(q)
	}

	f.Tokens = p.ToTokens()

	return nil
}

// FlexibleExpression as defined in python@3.13.0:
// https://docs.python.org/release/3.13.0/reference/expressions.html#grammar-token-python-grammar-flexible_expression
type FlexibleExpression struct {
	AssignmentExpression *AssignmentExpression
	StarredExpression    *StarredExpression
	Tokens
}

func (f *FlexibleExpression) parse(p *pyParser) error {
	q := p.NewGoal()
	if q.Peek() == (parser.Token{Type: TokenOperator, Data: "*"}) {
		f.StarredExpression = new(StarredExpression)

		if err := f.StarredExpression.parse(q); err != nil {
			return p.Error("FlexibleExpression", err)
		}
	} else {
		f.AssignmentExpression = new(AssignmentExpression)

		if err := f.AssignmentExpression.parse(q); err != nil {
			return p.Error("FlexibleExpression", err)
		}
	}

	p.Score(q)

	f.Tokens = p.ToTokens()

	return nil
}

// Comprehension as defined in python@3.13.0:
// https://docs.python.org/release/3.13.0/reference/expressions.html#grammar-token-python-grammar-comprehension
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

// ComprehensionFor as defined in python@3.13.0:
// https://docs.python.org/release/3.13.0/reference/expressions.html#grammar-token-python-grammar-comp_for
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

// ComprehensionIterator as defined in python@3.13.0:
// https://docs.python.org/release/3.13.0/reference/expressions.html#grammar-token-python-grammar-comp_iter
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

// ComprehensionIf as defined in python@3.13.0:
// https://docs.python.org/release/3.13.0/reference/expressions.html#grammar-token-python-grammar-comp_if
type ComprehensionIf struct {
	OrTest                OrTest
	ComprehensionIterator *ComprehensionIterator
	Tokens                Tokens
}

func (c *ComprehensionIf) parse(p *pyParser) error {
	if !p.AcceptToken(parser.Token{Type: TokenKeyword, Data: "if"}) {
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
			return p.Error("ComprehensionIf", err)
		}

		p.Score(q)
	}

	c.Tokens = p.ToTokens()

	return nil
}

// DictDisplay as defined in python@3.13.0:
// https://docs.python.org/release/3.13.0/reference/expressions.html#grammar-token-python-grammar-dict_display
type DictDisplay struct {
	DictItems         []DictItem
	DictComprehension *ComprehensionFor
	Tokens            Tokens
}

func (d *DictDisplay) parse(p, o *pyParser, e *Expression) error {
Loop:
	for {
		q := p

		if e == nil {
			q = p.NewGoal()
		}

		var di DictItem

		if err := di.parse(q, e); err != nil {
			if e != nil {
				p = o
			}

			return p.Error("DictDisplay", err)
		}

		if e == nil {
			p.Score(q)
		} else {
			e = nil
		}

		d.DictItems = append(d.DictItems, di)
		q = p.NewGoal()

		q.AcceptRunWhitespace()

		switch q.Peek() {
		case parser.Token{Type: TokenKeyword, Data: "async"}, parser.Token{Type: TokenKeyword, Data: "for"}:
			if len(d.DictItems) > 1 || d.DictItems[0].OrExpression != nil {
				return q.Error("DictDisplay", ErrInvalidKeyword)
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

		if tk := q.Peek(); tk == (parser.Token{Type: TokenDelimiter, Data: "}"}) || tk.Type == parser.TokenDone {
			break
		}

		p.Score(q)
	}

	d.Tokens = p.ToTokens()

	return nil
}

// DictItem as defined in python@3.13.0:
// https://docs.python.org/release/3.13.0/reference/expressions.html#grammar-token-python-grammar-dict_item
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

	d.Tokens = p.ToTokens()

	return nil
}

// GeneratorExpression as defined in python@3.13.0:
// https://docs.python.org/release/3.13.0/reference/expressions.html#grammar-token-python-grammar-generator_expression
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

// ArgumentListOrComprehension as defined in python@3.13.0:
// https://docs.python.org/release/3.13.0/reference/expressions.html#grammar-token-python-grammar-call
type ArgumentListOrComprehension struct {
	ArgumentList  *ArgumentList
	Comprehension *Comprehension
	Tokens        Tokens
}

func (a *ArgumentListOrComprehension) parse(p *pyParser) error {
	p.AcceptToken(parser.Token{Type: TokenDelimiter, Data: "("})
	p.OpenBrackets()
	p.AcceptRunWhitespace()

	q := p.NewGoal()

	if q.LookaheadLine(parser.Token{Type: TokenKeyword, Data: "for"}) == 0 {
		a.Comprehension = new(Comprehension)

		if err := a.Comprehension.parse(q, nil); err != nil {
			return p.Error("ArgumentListOrComprehension", err)
		}
	} else {
		a.ArgumentList = new(ArgumentList)

		if err := a.ArgumentList.parse(q); err != nil {
			return p.Error("ArgumentListOrComprehension", err)
		}
	}

	p.Score(q)
	p.AcceptRunWhitespace()

	if !p.AcceptToken(parser.Token{Type: TokenDelimiter, Data: ")"}) {
		return p.Error("ArgumentListOrComprehension", ErrMissingClosingParen)
	}

	p.CloseBrackets()

	a.Tokens = p.ToTokens()

	return nil
}

// ExpressionList as defined in python@3.13.0:
// https://docs.python.org/release/3.13.0/reference/expressions.html#grammar-token-python-grammar-expression_list
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

		if tk := q.Peek(); tk == (parser.Token{Type: TokenDelimiter, Data: "]"}) || tk == (parser.Token{Type: TokenDelimiter, Data: ")"}) || tk.Type == parser.TokenDone || tk.Type == TokenLineTerminator || tk.Type == TokenDedent {
			break
		} else if !q.AcceptToken(parser.Token{Type: TokenDelimiter, Data: ","}) {
			return p.Error("ExpressionList", ErrMissingComma)
		}

		q.AcceptRunWhitespace()

		if tk := q.Peek(); tk == (parser.Token{Type: TokenDelimiter, Data: "]"}) || tk == (parser.Token{Type: TokenDelimiter, Data: ")"}) || tk.Type == parser.TokenDone || tk.Type == TokenLineTerminator || tk.Type == TokenDedent {
			break
		}

		p.Score(q)
	}

	e.Tokens = p.ToTokens()

	return nil
}

// SliceList as defined in python@3.13.0:
// https://docs.python.org/release/3.13.0/reference/expressions.html#grammar-token-python-grammar-slice_list
type SliceList struct {
	SliceItems []SliceItem
	Tokens     Tokens
}

func (s *SliceList) parse(p *pyParser) error {
	p.AcceptToken(parser.Token{Type: TokenDelimiter, Data: "["})
	p.OpenBrackets()
	p.AcceptRunWhitespace()

	for !p.AcceptToken(parser.Token{Type: TokenDelimiter, Data: "]"}) {
		q := p.NewGoal()

		var si SliceItem

		if err := si.parse(q); err != nil {
			return p.Error("SliceList", err)
		}

		p.Score(q)

		s.SliceItems = append(s.SliceItems, si)

		p.AcceptRunWhitespace()

		if p.Peek() != (parser.Token{Type: TokenDelimiter, Data: "]"}) {
			if !p.AcceptToken(parser.Token{Type: TokenDelimiter, Data: ","}) {
				return p.Error("SliceList", ErrMissingComma)
			}

			p.AcceptRunWhitespace()
		}
	}

	p.CloseBrackets()

	s.Tokens = p.ToTokens()

	return nil
}

// SliceItem as defined in python@3.13.0:
// https://docs.python.org/release/3.13.0/reference/expressions.html#grammar-token-python-grammar-slice_item
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

// OrExpression as defined in python@3.13.0:
// https://docs.python.org/release/3.13.0/reference/expressions.html#grammar-token-python-grammar-or_expr
type OrExpression struct {
	XorExpression XorExpression
	OrExpression  *OrExpression
	Tokens        Tokens
}

func (o *OrExpression) parse(p *pyParser) error {
	q := p.NewGoal()

	if err := o.XorExpression.parse(p); err != nil {
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

// XorExpression as defined in python@3.13.0:
// https://docs.python.org/release/3.13.0/reference/expressions.html#grammar-token-python-grammar-xor_expr
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

// AndExpression as defined in python@3.13.0:
// https://docs.python.org/release/3.13.0/reference/expressions.html#grammar-token-python-grammar-and_expr
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

// ShiftExpression as defined in python@3.13.0:
// https://docs.python.org/release/3.13.0/reference/expressions.html#grammar-token-python-grammar-shift_expr
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

// AddExpression as defined in python@3.13.0:
// https://docs.python.org/release/3.13.0/reference/expressions.html#grammar-token-python-grammar-a_expr
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

// MultiplyExpression as defined in python@3.13.0:
// https://docs.python.org/release/3.13.0/reference/expressions.html#grammar-token-python-grammar-m_expr
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

// UnaryExpression as defined in python@3.13.0:
// https://docs.python.org/release/3.13.0/reference/expressions.html#grammar-token-python-grammar-u_expr
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

// PowerExpression as defined in python@3.13.0:
// https://docs.python.org/release/3.13.0/reference/expressions.html#grammar-token-python-grammar-power
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

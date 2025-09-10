package python

import "vimagination.zapto.org/parser"

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
//
// The first set of comments are parsed from before the StarredItem.
//
// The second set of comments are parsed from after any '*' token.
//
// The final set of comments are parsed from after the StarredItem.
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
//
// The first set of comments are parsed from before the PositionalArgument.
//
// The second set of comments are parsed from after any '*' token.
//
// The final set of comments are parsed from after the PositionalArgument.
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
//
// The first set of comments are parsed from before the StarredOrKeyword item.
//
// The final set of comments are parsed from after the StarredOrKeyword item.
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
//
// The first set of comments are parsed from before the KeywordArgument.
//
// The second set of comments are parsed from after any '**' token.
//
// The final set of comments are parsed from after the KeywordArgument.
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

// PrimaryExpression as defined in python@3.13.0:
// https://docs.python.org/release/3.13.0/reference/expressions.html#grammar-token-python-grammar-primary
//
// For an AttributeRef, the comments are parsed from before and after the '.'
// token.
//
// For a Slice or a Call, the comments are parsed from before the opening '[' or '('.
//
// NB: Comments are only parsed when in a multiline structure.
type PrimaryExpression struct {
	PrimaryExpression *PrimaryExpression
	Atom              *Atom
	AttributeRef      *Token
	Slicing           *SliceList
	Call              *ArgumentListOrComprehension
	Comments          [2]Comments
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

		aComments := q.AcceptRunWhitespaceCommentsIfMultiline()

		q.AcceptRunWhitespace()

		if q.AcceptToken(parser.Token{Type: TokenDelimiter, Data: "."}) {
			bComments := q.AcceptRunWhitespaceCommentsIfMultiline()

			q.AcceptRunWhitespace()

			if !q.Accept(TokenIdentifier) {
				return q.Error("PrimaryExpression", ErrMissingIdentifier)
			}

			pr.Tokens = p.ToTokens()
			ipr := *pr
			*pr = PrimaryExpression{
				PrimaryExpression: &ipr,
				AttributeRef:      q.GetLastToken(),
				Comments:          [2]Comments{aComments, bComments},
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
				Comments:          [2]Comments{aComments},
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
				Comments:          [2]Comments{aComments},
			}
		} else {
			break
		}

		p.Score(q)
	}

	pr.Tokens = p.ToTokens()

	return nil
}

func skipPrimaryExpression(p *pyParser) {
	skipAtom(p)
	p.AcceptRunWhitespace()

	if p.AcceptToken(parser.Token{Type: TokenDelimiter, Data: "."}) {
		p.AcceptRunWhitespace()
		skipPrimaryExpression(p)
	} else if p.AcceptToken(parser.Token{Type: TokenDelimiter, Data: "["}) {
		skipDepth(p, parser.Token{Type: TokenDelimiter, Data: "["}, parser.Token{Type: TokenDelimiter, Data: "]"})
	} else if p.AcceptToken(parser.Token{Type: TokenDelimiter, Data: "("}) {
		skipDepth(p, parser.Token{Type: TokenDelimiter, Data: "("}, parser.Token{Type: TokenDelimiter, Data: ")"})
	}
}

func skipDepth(p *pyParser, opener, closer parser.Token) {
	depth := 1

	for {
		switch p.Next().Token {
		case opener:
			depth++
		case closer:
			if depth--; depth == 0 {
				return
			}
		}
	}
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

func skipAtom(p *pyParser) {
	if !p.Accept(TokenIdentifier, TokenNumericLiteral, TokenStringLiteral, TokenBooleanLiteral) {
		skipEnclosure(p)
	}
}

// IsIdentifier returns true if the Atom contains an Identifier.
func (a *Atom) IsIdentifier() bool {
	return a.Identifier != nil
}

// Enclosure as defined in python@3.13.0:
// https://docs.python.org/release/3.13.0/reference/expressions.html#grammar-token-python-grammar-enclosure
//
// The first set of comments are parsed from directly after the opening paren,
// brace, or bracket; the second set of comments are parsed from directly
// before the closing paren, brace, or bracket.
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
		q := p.NewGoal()

		if q.AcceptRunAllWhitespace() == TokenDelimiter {
			e.Comments[0] = p.AcceptRunWhitespaceComments()

			p.AcceptRunAllWhitespace()
		} else {
			e.Comments[0] = p.AcceptRunWhitespaceCommentsNoNewline()
		}

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
	} else if p.AcceptToken(parser.Token{Type: TokenDelimiter, Data: "["}) {
		q := p.NewGoal()

		if q.AcceptRunAllWhitespace() == TokenDelimiter {
			e.Comments[0] = p.AcceptRunWhitespaceComments()
		} else {
			e.Comments[0] = p.AcceptRunWhitespaceCommentsNoNewline()
		}

		q = p.NewGoal()

		q.AcceptRunAllWhitespace()

		if q.AcceptToken(parser.Token{Type: TokenDelimiter, Data: "]"}) {
			p.Score(q)

			e.ListDisplay = &FlexibleExpressionListOrComprehension{
				Tokens: p.NewGoal().ToTokens(),
			}
		} else {
			p.AcceptRunWhitespaceNoComment()

			q := p.NewGoal()
			e.ListDisplay = new(FlexibleExpressionListOrComprehension)

			if err := e.ListDisplay.parse(q); err != nil {
				return p.Error("Enclosure", err)
			}

			p.Score(q)

			e.Comments[1] = p.AcceptRunWhitespaceComments()

			p.AcceptRunAllWhitespace()

			if !p.AcceptToken(parser.Token{Type: TokenDelimiter, Data: "]"}) {
				return p.Error("Enclosure", ErrMissingClosingBracket)
			}
		}
	} else if p.AcceptToken(parser.Token{Type: TokenDelimiter, Data: "{"}) {
		q := p.NewGoal()

		if q.AcceptRunAllWhitespace() == TokenDelimiter {
			e.Comments[0] = p.AcceptRunWhitespaceComments()
		} else {
			e.Comments[0] = p.AcceptRunWhitespaceCommentsNoNewline()
		}

		q = p.NewGoal()

		q.AcceptRunAllWhitespace()

		if q.AcceptToken(parser.Token{Type: TokenDelimiter, Data: "}"}) {
			p.Score(q)

			e.DictDisplay = &DictDisplay{
				Tokens: p.NewGoal().ToTokens(),
			}
		} else {
			q := p.NewGoal()

			q.AcceptRunAllWhitespace()

			var isDict bool

			switch q.Peek() {
			case parser.Token{Type: TokenOperator, Data: "*"}:
			case parser.Token{Type: TokenOperator, Data: "**"}:
				isDict = true
			default:
				skipExpression(q)
				q.AcceptRunWhitespace()

				isDict = q.AcceptToken(parser.Token{Type: TokenDelimiter, Data: ":"})
			}

			p.AcceptRunWhitespaceNoComment()

			if isDict {
				q = p.NewGoal()
				e.DictDisplay = new(DictDisplay)

				if err := e.DictDisplay.parse(q); err != nil {
					return p.Error("Enclosure", err)
				}
			} else {
				q = p.NewGoal()
				e.SetDisplay = new(FlexibleExpressionListOrComprehension)

				if err := e.SetDisplay.parse(q); err != nil {
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
	} else {
		return p.Error("Enclosure", ErrInvalidEnclosure)
	}

	e.Tokens = p.ToTokens()

	return nil
}

func skipEnclosure(p *pyParser) {
	if p.AcceptToken(parser.Token{Type: TokenDelimiter, Data: "("}) {
		skipDepth(p, parser.Token{Type: TokenDelimiter, Data: "("}, parser.Token{Type: TokenDelimiter, Data: ")"})
	} else if p.AcceptToken(parser.Token{Type: TokenDelimiter, Data: "["}) {
		skipDepth(p, parser.Token{Type: TokenDelimiter, Data: "["}, parser.Token{Type: TokenDelimiter, Data: "]"})
	} else if p.AcceptToken(parser.Token{Type: TokenDelimiter, Data: "{"}) {
		skipDepth(p, parser.Token{Type: TokenDelimiter, Data: "{"}, parser.Token{Type: TokenDelimiter, Data: "}"})
	}
}

// FlexibleExpressionListOrComprehension as defined in python@3.13.0:
// https://docs.python.org/release/3.13.0/reference/expressions.html#grammar-token-python-grammar-list_display
type FlexibleExpressionListOrComprehension struct {
	FlexibleExpressionList *FlexibleExpressionList
	Comprehension          *Comprehension
	Tokens                 Tokens
}

func (f *FlexibleExpressionListOrComprehension) parse(p *pyParser) error {
	q := p.NewGoal()

	skipAssignmentExpression(q)

	q.AcceptRunWhitespace()

	if q.AcceptToken(parser.Token{Type: TokenKeyword, Data: "async"}) || q.AcceptToken(parser.Token{Type: TokenKeyword, Data: "for"}) {
		q = p.NewGoal()
		f.Comprehension = new(Comprehension)

		if err := f.Comprehension.parse(q); err != nil {
			return p.Error("FlexibleExpressionListOrComprehension", err)
		}
	} else {
		q = p.NewGoal()
		f.FlexibleExpressionList = new(FlexibleExpressionList)

		if err := f.FlexibleExpressionList.parse(q); err != nil {
			return p.Error("FlexibleExpressionListOrComprehension", err)
		}
	}

	p.Score(q)

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
		p.AcceptRunAllWhitespaceNoComment()

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
	}

	f.Tokens = p.ToTokens()

	return nil
}

// FlexibleExpression as defined in python@3.13.0:
// https://docs.python.org/release/3.13.0/reference/expressions.html#grammar-token-python-grammar-flexible_expression
//
// The first set of comments are parsed from before the FlexibleExpression; the
// second set of comments are parsed from after the FlexibleExpression.
type FlexibleExpression struct {
	AssignmentExpression *AssignmentExpression
	StarredExpression    *OrExpression
	Comments             [2]Comments
	Tokens
}

func (f *FlexibleExpression) parse(p *pyParser) error {
	f.Comments[0] = p.AcceptRunWhitespaceComments()

	p.AcceptRunWhitespace()

	if p.AcceptToken(parser.Token{Type: TokenOperator, Data: "*"}) {
		p.AcceptRunWhitespace()

		q := p.NewGoal()
		f.StarredExpression = new(OrExpression)

		if err := f.StarredExpression.parse(q); err != nil {
			return p.Error("FlexibleExpression", err)
		}

		p.Score(q)
	} else {
		q := p.NewGoal()
		f.AssignmentExpression = new(AssignmentExpression)

		if err := f.AssignmentExpression.parse(q); err != nil {
			return p.Error("FlexibleExpression", err)
		}

		p.Score(q)
	}

	q := p.NewGoal()

	q.AcceptRunWhitespace()

	if q.Peek() == (parser.Token{Type: TokenDelimiter, Data: ","}) {
		f.Comments[1] = p.AcceptRunWhitespaceCommentsIfMultiline()
	} else {
		f.Comments[1] = p.AcceptRunWhitespaceCommentsNoNewlineIfMultiline()
	}

	f.Tokens = p.ToTokens()

	return nil
}

// Comprehension as defined in python@3.13.0:
// https://docs.python.org/release/3.13.0/reference/expressions.html#grammar-token-python-grammar-comprehension
//
// The first set of comments are parsed from before the Comprehension.
//
// The second set of comments are parsed from between the AssignmentExpression
// and he ComprehensionFor.
//
// The final set of comments are parsed from after the Comprehension.
//
// NB: Comments are only parsed when in a multiline structure.
type Comprehension struct {
	AssignmentExpression AssignmentExpression
	ComprehensionFor     ComprehensionFor
	Comments             [3]Comments
	Tokens               Tokens
}

func (c *Comprehension) parse(p *pyParser) error {
	c.Comments[0] = p.AcceptRunWhitespaceComments()

	p.AcceptRunWhitespace()

	q := p.NewGoal()

	if err := c.AssignmentExpression.parse(q); err != nil {
		return p.Error("Comprehension", err)
	}

	p.Score(q)

	c.Comments[1] = p.AcceptRunWhitespaceComments()

	p.AcceptRunWhitespace()

	q = p.NewGoal()

	if err := c.ComprehensionFor.parse(q); err != nil {
		return p.Error("Comprehension", err)
	}

	p.Score(q)

	c.Comments[2] = p.AcceptRunWhitespaceCommentsNoNewline()
	c.Tokens = p.ToTokens()

	return nil
}

// ComprehensionFor as defined in python@3.13.0:
// https://docs.python.org/release/3.13.0/reference/expressions.html#grammar-token-python-grammar-comp_for
//
// The first set of comments are parsed after an 'async' token.
//
// The second set of comments are parsed from after the 'in' token.
//
// NB: Comments are only parsed when in a multiline structure.
type ComprehensionFor struct {
	Async                 bool
	TargetList            TargetList
	OrTest                OrTest
	ComprehensionIterator *ComprehensionIterator
	Comments              [2]Comments
	Tokens                Tokens
}

func (c *ComprehensionFor) parse(p *pyParser) error {
	c.Async = p.AcceptToken(parser.Token{Type: TokenKeyword, Data: "async"})

	c.Comments[0] = p.AcceptRunWhitespaceCommentsIfMultiline()

	p.AcceptRunWhitespace()

	if !p.AcceptToken(parser.Token{Type: TokenKeyword, Data: "for"}) {
		return p.Error("ComprehensionFor", ErrMissingFor)
	}

	p.AcceptRunWhitespaceNoComment()

	q := p.NewGoal()

	if err := c.TargetList.parse(q); err != nil {
		return p.Error("ComprehensionFor", err)
	}

	p.Score(q)

	p.AcceptRunWhitespace()

	if !p.AcceptToken(parser.Token{Type: TokenKeyword, Data: "in"}) {
		return p.Error("ComprehensionFor", ErrMissingIn)
	}

	c.Comments[1] = p.AcceptRunWhitespaceCommentsIfMultiline()

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
		p.AcceptRunWhitespaceNoComment()

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
//
// When in a multiline structure, the comments are parsed from before and after
// the ComprehensionIterator.
type ComprehensionIterator struct {
	ComprehensionFor *ComprehensionFor
	ComprehensionIf  *ComprehensionIf
	Comments         [2]Comments
	Tokens           Tokens
}

func (c *ComprehensionIterator) parse(p *pyParser) error {
	c.Comments[0] = p.AcceptRunWhitespaceComments()

	p.AcceptRunWhitespace()

	q := p.NewGoal()

	q.AcceptRunAllWhitespace()

	if q.Peek() == (parser.Token{Type: TokenKeyword, Data: "if"}) {
		q = p.NewGoal()
		c.ComprehensionIf = new(ComprehensionIf)

		if err := c.ComprehensionIf.parse(q); err != nil {
			return p.Error("ComprehensionIterator", err)
		}
	} else {
		q = p.NewGoal()
		c.ComprehensionFor = new(ComprehensionFor)

		if err := c.ComprehensionFor.parse(q); err != nil {
			return p.Error("ComprehensionIterator", err)
		}
	}

	p.Score(q)

	c.Comments[1] = p.AcceptRunWhitespaceComments()
	c.Tokens = p.ToTokens()

	return nil
}

// ComprehensionIf as defined in python@3.13.0:
// https://docs.python.org/release/3.13.0/reference/expressions.html#grammar-token-python-grammar-comp_if
//
// When in a multiline structure, the comments are parsed from after the 'if'
// keyword.
type ComprehensionIf struct {
	OrTest                OrTest
	ComprehensionIterator *ComprehensionIterator
	Comments              Comments
	Tokens                Tokens
}

func (c *ComprehensionIf) parse(p *pyParser) error {
	if !p.AcceptToken(parser.Token{Type: TokenKeyword, Data: "if"}) {
		return p.Error("ComprehensionIf", ErrMissingIf)
	}

	c.Comments = p.AcceptRunWhitespaceCommentsIfMultiline()

	p.AcceptRunWhitespace()

	q := p.NewGoal()

	if err := c.OrTest.parse(q); err != nil {
		return p.Error("ComprehensionIf", err)
	}

	p.Score(q)

	q = p.NewGoal()

	q.AcceptRunWhitespace()

	switch q.Peek() {
	case parser.Token{Type: TokenKeyword, Data: "if"}, parser.Token{Type: TokenKeyword, Data: "async"}, parser.Token{Type: TokenKeyword, Data: "for"}:
		p.AcceptRunWhitespaceNoComment()

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

func (d *DictDisplay) parse(p *pyParser) error {
Loop:
	for {
		q := p.NewGoal()

		var di DictItem

		if err := di.parse(q); err != nil {
			return p.Error("DictDisplay", err)
		}

		p.Score(q)

		d.DictItems = append(d.DictItems, di)
		q = p.NewGoal()

		q.AcceptRunWhitespace()

		switch q.Peek() {
		case parser.Token{Type: TokenKeyword, Data: "async"}, parser.Token{Type: TokenKeyword, Data: "for"}:
			if len(d.DictItems) > 1 || d.DictItems[0].OrExpression != nil {
				return q.Error("DictDisplay", ErrInvalidKeyword)
			}

			p.AcceptRunWhitespaceNoNewline()

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
//
// The first set of comments are parsed from before the DictItem.
//
// In a key/value DictItem, the second and third comments are parsed from before
// and after the ':' token; otherwise, the second comments are parsed from after
// the '**' token.
//
// The final set of comments are parsed from after the DictItem.
type DictItem struct {
	Key          *Expression
	Value        *Expression
	OrExpression *OrExpression
	Comments     [4]Comments
	Tokens       Tokens
}

func (d *DictItem) parse(p *pyParser) error {
	d.Comments[0] = p.AcceptRunWhitespaceCommentsIfMultiline()

	p.AcceptRunWhitespace()

	if p.AcceptToken(parser.Token{Type: TokenOperator, Data: "**"}) {
		d.Comments[1] = p.AcceptRunWhitespaceCommentsIfMultiline()

		p.AcceptRunWhitespace()

		q := p.NewGoal()
		d.OrExpression = new(OrExpression)

		if err := d.OrExpression.parse(q); err != nil {
			return p.Error("DictItem", err)
		}

		p.Score(q)
	} else {
		q := p.NewGoal()
		d.Key = new(Expression)

		if err := d.Key.parse(q); err != nil {
			return p.Error("DictItem", err)
		}

		p.Score(q)

		d.Comments[1] = p.AcceptRunWhitespaceCommentsIfMultiline()

		p.AcceptRunWhitespace()

		if !p.AcceptToken(parser.Token{Type: TokenDelimiter, Data: ":"}) {
			return p.Error("DictItem", ErrMissingColon)
		}

		d.Comments[2] = p.AcceptRunWhitespaceCommentsIfMultiline()

		p.AcceptRunWhitespace()

		q = p.NewGoal()
		d.Value = new(Expression)

		if err := d.Value.parse(q); err != nil {
			return p.Error("DictItem", err)
		}

		p.Score(q)
	}

	q := p.NewGoal()

	q.AcceptRunWhitespace()

	if q.Peek() == (parser.Token{Type: TokenDelimiter, Data: ","}) {
		d.Comments[3] = p.AcceptRunWhitespaceCommentsIfMultiline()
	} else {
		d.Comments[3] = p.AcceptRunWhitespaceCommentsNoNewlineIfMultiline()
	}

	d.Tokens = p.ToTokens()

	return nil
}

// GeneratorExpression as defined in python@3.13.0:
// https://docs.python.org/release/3.13.0/reference/expressions.html#grammar-token-python-grammar-generator_expression
//
// The first set of comments are parsed from before the GeneratorExpression.
//
// The second set of comments are parsed after the expression.
//
// The third set of comments are parsed from after the GeneratorExpression.
type GeneratorExpression struct {
	Expression       Expression
	ComprehensionFor ComprehensionFor
	Comments         [3]Comments
	Tokens           Tokens
}

func (g *GeneratorExpression) parse(p *pyParser) error {
	g.Comments[0] = p.AcceptRunWhitespaceComments()
	p.AcceptRunWhitespace()

	q := p.NewGoal()

	if err := g.Expression.parse(q); err != nil {
		return p.Error("GeneratorExpression", err)
	}

	p.Score(q)

	g.Comments[1] = p.AcceptRunWhitespaceComments()

	p.AcceptRunAllWhitespace()

	q = p.NewGoal()

	if err := g.ComprehensionFor.parse(q); err != nil {
		return p.Error("GeneratorExpression", err)
	}

	p.Score(q)

	g.Comments[2] = p.AcceptRunWhitespaceComments()
	g.Tokens = p.ToTokens()

	return nil
}

// ArgumentListOrComprehension as defined in python@3.13.0:
// https://docs.python.org/release/3.13.0/reference/expressions.html#grammar-token-python-grammar-call
//
// The first set of comments are parsed from directly after the opening paren.
//
// The second set of comments are parsed from directly before the closing paren.
type ArgumentListOrComprehension struct {
	ArgumentList  *ArgumentList
	Comprehension *Comprehension
	Comments      [2]Comments
	Tokens        Tokens
}

func (a *ArgumentListOrComprehension) parse(p *pyParser) error {
	p.AcceptToken(parser.Token{Type: TokenDelimiter, Data: "("})

	a.Comments[0] = p.AcceptRunWhitespaceCommentsNoNewline()

	p.AcceptRunWhitespaceNoComment()

	q := p.NewGoal()

	q.AcceptRunWhitespace()

	if q.LookaheadLine(parser.Token{Type: TokenKeyword, Data: "for"}) == 0 {
		q = p.NewGoal()
		a.Comprehension = new(Comprehension)

		if err := a.Comprehension.parse(q); err != nil {
			return p.Error("ArgumentListOrComprehension", err)
		}
	} else {
		q = p.NewGoal()
		a.ArgumentList = new(ArgumentList)

		if err := a.ArgumentList.parse(q); err != nil {
			return p.Error("ArgumentListOrComprehension", err)
		}
	}

	p.Score(q)

	a.Comments[1] = p.AcceptRunWhitespaceComments()

	p.AcceptRunWhitespace()

	if !p.AcceptToken(parser.Token{Type: TokenDelimiter, Data: ")"}) {
		return p.Error("ArgumentListOrComprehension", ErrMissingClosingParen)
	}

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
//
// The first set of comments are parsed from directly after the opening brace.
//
// The second set of comments are parsed from directly before the closing brace.
type SliceList struct {
	SliceItems []SliceItem
	Comments   [2]Comments
	Tokens     Tokens
}

func (s *SliceList) parse(p *pyParser) error {
	p.AcceptToken(parser.Token{Type: TokenDelimiter, Data: "["})

	s.Comments[0] = p.AcceptRunWhitespaceCommentsNoNewline()

	q := p.NewGoal()

	q.AcceptRunWhitespace()

	for !q.AcceptToken(parser.Token{Type: TokenDelimiter, Data: "]"}) {
		p.AcceptRunWhitespaceNoComment()

		q = p.NewGoal()

		var si SliceItem

		if err := si.parse(q); err != nil {
			return p.Error("SliceList", err)
		}

		p.Score(q)

		s.SliceItems = append(s.SliceItems, si)
		q = p.NewGoal()

		q.AcceptRunWhitespace()

		if q.Peek() != (parser.Token{Type: TokenDelimiter, Data: "]"}) {
			if !q.AcceptToken(parser.Token{Type: TokenDelimiter, Data: ","}) {
				return q.Error("SliceList", ErrMissingComma)
			}

			p.Score(q)

			q = p.NewGoal()

			q.AcceptRunWhitespace()
		}
	}

	s.Comments[1] = p.AcceptRunWhitespaceComments()

	p.AcceptRunAllWhitespace()
	p.Next()

	s.Tokens = p.ToTokens()

	return nil
}

// SliceItem as defined in python@3.13.0:
// https://docs.python.org/release/3.13.0/reference/expressions.html#grammar-token-python-grammar-slice_item
//
// The first set of comments are parsed from directly after the opening bracket.
//
// The second and third set of comments are parsed from either side of the
// first ':' token.
//
// The fourth and fifth set of comments are parsed from either side of the
// second ':' token, if it exists.
//
// The final set of comments are parsed from directly before the closing
// bracket.
type SliceItem struct {
	Expression *Expression
	LowerBound *Expression
	UpperBound *Expression
	Stride     *Expression
	Comments   [6]Comments
	Tokens     Tokens
}

func (s *SliceItem) parse(p *pyParser) error {
	s.Comments[0] = p.AcceptRunWhitespaceComments()

	p.AcceptRunWhitespace()

	q := p.NewGoal()

	s.Expression = new(Expression)

	if err := s.Expression.parse(q); err != nil {
		return p.Error("SliceItem", err)
	}

	p.Score(q)

	q = p.NewGoal()

	q.AcceptRunWhitespace()

	if q.AcceptToken(parser.Token{Type: TokenDelimiter, Data: ":"}) {
		s.Comments[1] = p.AcceptRunWhitespaceComments()

		p.AcceptRunWhitespace()
		p.Next()

		s.Comments[2] = p.AcceptRunWhitespaceComments()

		p.AcceptRunAllWhitespace()

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
			s.Comments[3] = p.AcceptRunWhitespaceComments()

			p.AcceptRunWhitespace()
			p.Next()

			s.Comments[4] = p.AcceptRunWhitespaceComments()

			p.AcceptRunWhitespace()

			q = p.NewGoal()
			s.Stride = new(Expression)

			if err := s.Stride.parse(q); err != nil {
				return p.Error("SliceItem", err)
			}

			p.Score(q)
		}
	}

	q = p.NewGoal()

	q.AcceptRunAllWhitespace()

	if q.AcceptToken(parser.Token{Type: TokenDelimiter, Data: ","}) {
		s.Comments[5] = p.AcceptRunWhitespaceComments()
	} else {
		s.Comments[5] = p.AcceptRunWhitespaceCommentsNoNewline()
	}

	s.Tokens = p.ToTokens()

	return nil
}

// OrExpression as defined in python@3.13.0:
// https://docs.python.org/release/3.13.0/reference/expressions.html#grammar-token-python-grammar-or_expr
//
// When in a multiline structure, comments are parsed on either side of the
// '|' token.
type OrExpression struct {
	XorExpression XorExpression
	OrExpression  *OrExpression
	Comments      [2]Comments
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
		o.Comments[0] = p.AcceptRunWhitespaceCommentsIfMultiline()

		p.AcceptRunWhitespace()
		p.Next()

		o.Comments[1] = p.AcceptRunWhitespaceCommentsIfMultiline()

		p.AcceptRunWhitespace()

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

func skipOrExpression(p *pyParser) {
	skipXorExpression(p)
	p.AcceptRunWhitespace()

	if p.AcceptToken(parser.Token{Type: TokenOperator, Data: "|"}) {
		p.AcceptRunWhitespace()
		skipOrExpression(p)
	}
}

// XorExpression as defined in python@3.13.0:
// https://docs.python.org/release/3.13.0/reference/expressions.html#grammar-token-python-grammar-xor_expr
//
// When in a multiline structure, comments are parsed on either side of the
// '^' token.
type XorExpression struct {
	AndExpression AndExpression
	XorExpression *XorExpression
	Comments      [2]Comments
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
		x.Comments[0] = p.AcceptRunWhitespaceCommentsIfMultiline()

		p.AcceptRunWhitespace()
		p.Next()

		x.Comments[1] = p.AcceptRunWhitespaceCommentsIfMultiline()

		p.AcceptRunWhitespace()

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

func skipXorExpression(p *pyParser) {
	skipAndExpression(p)
	p.AcceptRunWhitespace()

	if p.AcceptToken(parser.Token{Type: TokenOperator, Data: "^"}) {
		p.AcceptRunWhitespace()
		skipXorExpression(p)
	}
}

// AndExpression as defined in python@3.13.0:
// https://docs.python.org/release/3.13.0/reference/expressions.html#grammar-token-python-grammar-and_expr
//
// When in a multiline structure, comments are parsed on either side of the
// '&' token.
type AndExpression struct {
	ShiftExpression ShiftExpression
	AndExpression   *AndExpression
	Comments        [2]Comments
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
		a.Comments[0] = p.AcceptRunWhitespaceCommentsIfMultiline()

		p.AcceptRunWhitespace()
		p.Next()

		a.Comments[1] = p.AcceptRunWhitespaceCommentsIfMultiline()

		p.AcceptRunWhitespace()

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

func skipAndExpression(p *pyParser) {
	skipShiftExpression(p)
	p.AcceptRunWhitespace()

	if p.AcceptToken(parser.Token{Type: TokenOperator, Data: "&"}) {
		p.AcceptRunWhitespace()
		skipAndExpression(p)
	}
}

// ShiftExpression as defined in python@3.13.0:
// https://docs.python.org/release/3.13.0/reference/expressions.html#grammar-token-python-grammar-shift_expr
//
// When in a multiline structure, comments are parsed on either side of the
// '<<' or '>>' tokens.
type ShiftExpression struct {
	AddExpression   AddExpression
	Shift           *Token
	ShiftExpression *ShiftExpression
	Comments        [2]Comments
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
		s.Comments[0] = p.AcceptRunWhitespaceCommentsIfMultiline()

		p.AcceptRunWhitespace()
		p.Next()
		s.Shift = q.GetLastToken()

		s.Comments[1] = p.AcceptRunWhitespaceCommentsIfMultiline()

		p.AcceptRunWhitespace()

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

func skipShiftExpression(p *pyParser) {
	skipAddExpression(p)
	p.AcceptRunWhitespace()

	if p.AcceptToken(parser.Token{Type: TokenOperator, Data: "<<"}) || p.AcceptToken(parser.Token{Type: TokenOperator, Data: ">>"}) {
		p.AcceptRunWhitespace()
		skipShiftExpression(p)
	}
}

// AddExpression as defined in python@3.13.0:
// https://docs.python.org/release/3.13.0/reference/expressions.html#grammar-token-python-grammar-a_expr
//
// When in a multiline structure, comments are parsed on either side of the
// '+' or '-' tokens.
type AddExpression struct {
	MultiplyExpression MultiplyExpression
	Add                *Token
	AddExpression      *AddExpression
	Comments           [2]Comments
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
		a.Comments[0] = p.AcceptRunWhitespaceCommentsIfMultiline()

		p.AcceptRunWhitespace()
		p.Next()
		a.Add = q.GetLastToken()

		a.Comments[1] = p.AcceptRunWhitespaceCommentsIfMultiline()

		p.AcceptRunWhitespace()

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

func skipAddExpression(p *pyParser) {
	skipMultiplyExpression(p)
	p.AcceptRunWhitespace()

	if p.AcceptToken(parser.Token{Type: TokenOperator, Data: "+"}) || p.AcceptToken(parser.Token{Type: TokenOperator, Data: "-"}) {
		p.AcceptRunWhitespace()
		skipAddExpression(p)
	}
}

// MultiplyExpression as defined in python@3.13.0:
// https://docs.python.org/release/3.13.0/reference/expressions.html#grammar-token-python-grammar-m_expr
//
// When in a multiline structure, comments are parsed on either side of the
// operator token.
type MultiplyExpression struct {
	UnaryExpression    UnaryExpression
	Multiply           *Token
	MultiplyExpression *MultiplyExpression
	Comments           [2]Comments
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
		m.Comments[0] = p.AcceptRunWhitespaceCommentsIfMultiline()

		p.AcceptRunWhitespace()
		p.Next()
		m.Multiply = q.GetLastToken()

		m.Comments[1] = p.AcceptRunWhitespaceCommentsIfMultiline()

		p.AcceptRunWhitespace()

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

func skipMultiplyExpression(p *pyParser) {
	skipUnaryExpression(p)
	p.AcceptRunWhitespace()

	if p.AcceptToken(parser.Token{Type: TokenOperator, Data: "*"}) || p.AcceptToken(parser.Token{Type: TokenOperator, Data: "@"}) ||
		p.AcceptToken(parser.Token{Type: TokenOperator, Data: "//"}) || p.AcceptToken(parser.Token{Type: TokenOperator, Data: "/"}) ||
		p.AcceptToken(parser.Token{Type: TokenOperator, Data: "%"}) {
		p.AcceptRunWhitespace()
		skipMultiplyExpression(p)
	}
}

// UnaryExpression as defined in python@3.13.0:
// https://docs.python.org/release/3.13.0/reference/expressions.html#grammar-token-python-grammar-u_expr
//
// When in a multiline structure, comments are parsed on either side of the
// operator token.
type UnaryExpression struct {
	PowerExpression *PowerExpression
	Unary           *Token
	UnaryExpression *UnaryExpression
	Comments
	Tokens Tokens
}

func (u *UnaryExpression) parse(p *pyParser) error {
	if p.AcceptToken(parser.Token{Type: TokenOperator, Data: "-"}) ||
		p.AcceptToken(parser.Token{Type: TokenOperator, Data: "+"}) ||
		p.AcceptToken(parser.Token{Type: TokenOperator, Data: "~"}) {
		u.Unary = p.GetLastToken()

		u.Comments = p.AcceptRunWhitespaceCommentsIfMultiline()
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

func skipUnaryExpression(p *pyParser) {
	if p.AcceptToken(parser.Token{Type: TokenOperator, Data: "-"}) ||
		p.AcceptToken(parser.Token{Type: TokenOperator, Data: "+"}) ||
		p.AcceptToken(parser.Token{Type: TokenOperator, Data: "~"}) {
		p.AcceptRunWhitespace()
		skipUnaryExpression(p)
	} else {
		skipPowerExpression(p)
	}
}

// PowerExpression as defined in python@3.13.0:
// https://docs.python.org/release/3.13.0/reference/expressions.html#grammar-token-python-grammar-power
//
// The first set of comments are parsed after an 'await' keyword.
//
// The second and third set of comments are parsed from before and after the
// '**' token.
//
// NB: Comments are only parsed when in a multiline structure.
type PowerExpression struct {
	AwaitExpression   bool
	PrimaryExpression PrimaryExpression
	UnaryExpression   *UnaryExpression
	Comments          [3]Comments
	Tokens            Tokens
}

func (pe *PowerExpression) parse(p *pyParser) error {
	if p.AcceptToken(parser.Token{Type: TokenKeyword, Data: "await"}) {
		pe.AwaitExpression = true

		pe.Comments[0] = p.AcceptRunWhitespaceCommentsIfMultiline()

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
		pe.Comments[1] = p.AcceptRunWhitespaceCommentsIfMultiline()

		p.AcceptRunWhitespace()
		p.Next()

		pe.Comments[2] = p.AcceptRunWhitespaceCommentsIfMultiline()

		p.AcceptRunWhitespace()

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

func skipPowerExpression(p *pyParser) {
	p.AcceptToken(parser.Token{Type: TokenKeyword, Data: "await"})
	p.AcceptRunWhitespace()
	skipPrimaryExpression(p)
	p.AcceptRunWhitespace()

	if p.AcceptToken(parser.Token{Type: TokenOperator, Data: "**"}) {
		p.AcceptRunWhitespace()
		skipUnaryExpression(p)
	}
}

// StatementType specifies the type of a SimpleStatment.
type StatementType uint8

const (
	StatementAssert StatementType = iota
	StatementAssignment
	StatementAugmentedAssignment
	StatementAnnotatedAssignment
	StatementPass
	StatementDel
	StatementReturn
	StatementYield
	StatementRaise
	StatementBreak
	StatementContinue
	StatementImport
	StatementGlobal
	StatementNonLocal
	StatementTyp
)

// StarredExpression as defined in python@3.12.6:
// https://docs.python.org/release/3.12.6/reference/expressions.html#grammar-token-python-grammar-starred_expression
//
// When in a multiline structure, the comments are parsed before and after the StarredExpression.
type StarredExpression struct {
	Expression  *Expression
	StarredList *StarredList
	Comments    [2]Comments
	Tokens      Tokens
}

func (s *StarredExpression) parse(p *pyParser) error {
	s.Comments[0] = p.AcceptRunWhitespaceCommentsIfMultiline()

	p.AcceptRunAllWhitespace()

	q := p.NewGoal()

	if q.Peek() == (parser.Token{Type: TokenOperator, Data: "*"}) || q.LookaheadLine(parser.Token{Type: TokenDelimiter, Data: ","}) == 0 {
		s.StarredList = new(StarredList)

		if err := s.StarredList.parse(q); err != nil {
			return p.Error("StarredExpression", err)
		}
	} else {
		s.Expression = new(Expression)

		if err := s.Expression.parse(q); err != nil {
			return p.Error("StarredExpression", err)
		}
	}

	p.Score(q)

	s.Comments[1] = p.AcceptRunWhitespaceCommentsNoNewlineIfMultiline()
	s.Tokens = p.ToTokens()

	return nil
}

// AssignmentExpression as defined in python@3.13.0:
// https://docs.python.org/release/3.13.0/reference/expressions.html#grammar-token-python-grammar-assignment_expression
//
// When in a multiline structure, comments are parsed on either side of the
// ':=' operator.
type AssignmentExpression struct {
	Identifier *Token
	Expression Expression
	Comments   [2]Comments
	Tokens
}

func (a *AssignmentExpression) parse(p *pyParser) error {
	q := p.NewGoal()

	if q.Accept(TokenIdentifier) {
		q.AcceptRunWhitespace()

		if q.AcceptToken(parser.Token{Type: TokenOperator, Data: ":="}) {
			p.Next()
			a.Identifier = p.GetLastToken()

			a.Comments[0] = p.AcceptRunWhitespaceCommentsIfMultiline()

			p.AcceptRunWhitespace()
			p.Next()

			a.Comments[1] = p.AcceptRunWhitespaceCommentsIfMultiline()

			p.AcceptRunWhitespace()
		}

		q = p.NewGoal()
	}

	if err := a.Expression.parse(q); err != nil {
		return p.Error("AssignmentExpression", err)
	}

	p.Score(q)

	a.Tokens = p.ToTokens()

	return nil
}

func skipAssignmentExpression(p *pyParser) {
	p.AcceptRunWhitespace()

	q := p.NewGoal()

	if q.Accept(TokenIdentifier) {
		q.AcceptRunWhitespace()

		if q.AcceptToken(parser.Token{Type: TokenOperator, Data: ":="}) {
			q.AcceptRunWhitespace()
			p.Score(q)
		}
	}

	skipExpression(p)
}

// Expression as defined in python@3.13.0:
// https://docs.python.org/release/3.13.0/reference/expressions.html#grammar-token-python-grammar-expression
type Expression struct {
	ConditionalExpression *ConditionalExpression
	LambdaExpression      *LambdaExpression
	Tokens                Tokens
}

func (e *Expression) parse(p *pyParser) error {
	if p.Peek() == (parser.Token{Type: TokenKeyword, Data: "lambda"}) {
		e.LambdaExpression = new(LambdaExpression)
		q := p.NewGoal()

		if err := e.LambdaExpression.parse(q); err != nil {
			return p.Error("Expression", err)
		}

		p.Score(q)
	} else {
		e.ConditionalExpression = new(ConditionalExpression)
		q := p.NewGoal()

		if err := e.ConditionalExpression.parse(q); err != nil {
			return p.Error("Expression", err)
		}

		p.Score(q)
	}

	e.Tokens = p.ToTokens()

	return nil
}

func skipExpression(p *pyParser) {
	p.AcceptRunWhitespace()

	if p.AcceptToken(parser.Token{Type: TokenKeyword, Data: "lambda"}) {
		skipDepth(p, parser.Token{Type: TokenKeyword, Data: "lambda"}, parser.Token{Type: TokenDelimiter, Data: ":"})
		p.AcceptRunWhitespace()
		skipExpression(p)
	} else {
		skipConditionalExpression(p)
	}
}

// ConditionalExpression as defined in python@3.13.0:
// https://docs.python.org/release/3.13.0/reference/expressions.html#grammar-token-python-grammar-conditional_expression
//
// The first and second sets of comments are parsed from before and after an
// 'if' keyword.
//
// The third and fourth sets of comments are parsed from before and after an
// 'else' keyword.
//
// NB: Comments are only parsed when in a multiline structure.
type ConditionalExpression struct {
	OrTest   OrTest
	If       *OrTest
	Else     *Expression
	Comments [4]Comments
	Tokens   Tokens
}

func (c *ConditionalExpression) parse(p *pyParser) error {
	q := p.NewGoal()

	if err := c.OrTest.parse(q); err != nil {
		return p.Error("ConditionalExpression", err)
	}

	p.Score(q)

	q = p.NewGoal()

	q.AcceptRunWhitespace()

	if q.AcceptToken(parser.Token{Type: TokenKeyword, Data: "if"}) {
		c.Comments[0] = p.AcceptRunWhitespaceCommentsIfMultiline()

		p.AcceptRunWhitespace()
		p.Next()

		c.Comments[1] = p.AcceptRunWhitespaceCommentsIfMultiline()

		p.AcceptRunWhitespace()

		q = p.NewGoal()
		c.If = new(OrTest)

		if err := c.If.parse(q); err != nil {
			return p.Error("ConditionalExpression", err)
		}

		p.Score(q)

		c.Comments[2] = p.AcceptRunWhitespaceCommentsIfMultiline()

		p.AcceptRunWhitespace()

		if !p.AcceptToken(parser.Token{Type: TokenKeyword, Data: "else"}) {
			return p.Error("ConditionalExpression", ErrMissingElse)
		}

		c.Comments[3] = p.AcceptRunWhitespaceCommentsIfMultiline()

		p.AcceptRunWhitespace()

		q = p.NewGoal()
		c.Else = new(Expression)

		if err := c.Else.parse(q); err != nil {
			return p.Error("ConditionalExpression", err)
		}

		p.Score(q)
	}

	c.Tokens = p.ToTokens()

	return nil
}

func skipConditionalExpression(p *pyParser) {
	p.AcceptRunWhitespace()
	skipOrTest(p)
	p.AcceptRunWhitespace()

	if !p.AcceptToken(parser.Token{Type: TokenKeyword, Data: "if"}) {
		return
	}

	skipOrTest(p)
	p.AcceptRunWhitespace()
	p.AcceptToken(parser.Token{Type: TokenKeyword, Data: "else"})
	p.AcceptRunWhitespace()
	skipExpression(p)
}

// LambdaExpression as defined in python@3.13.0:
// https://docs.python.org/release/3.13.0/reference/expressions.html#grammar-token-python-grammar-lambda_expr
//
// The first set of comments are parsed after the 'lambda' keyword.
//
// If there are params, the second set of comments are parsed before the ':'
// token.
//
// The third set of comments are parsed after the ':' token.
//
// NB: Comments are only parsed when in a multiline structure.
type LambdaExpression struct {
	ParameterList *ParameterList
	Expression    Expression
	Comments      [3]Comments
	Tokens        Tokens
}

func (l *LambdaExpression) parse(p *pyParser) error {
	p.Next()

	l.Comments[0] = p.AcceptRunWhitespaceCommentsIfMultiline()

	p.AcceptRunWhitespace()

	if !p.AcceptToken(parser.Token{Type: TokenDelimiter, Data: ":"}) {
		p.AcceptRunWhitespaceNoComment()

		q := p.NewGoal()
		l.ParameterList = new(ParameterList)

		if err := l.ParameterList.parse(q, false); err != nil {
			return p.Error("LambdaExpression", err)
		}

		p.Score(q)

		l.Comments[1] = p.AcceptRunWhitespaceCommentsIfMultiline()

		p.AcceptRunWhitespace()

		if !p.AcceptToken(parser.Token{Type: TokenDelimiter, Data: ":"}) {
			return p.Error("LambdaExpression", ErrMissingColon)
		}
	}

	l.Comments[2] = p.AcceptRunWhitespaceCommentsIfMultiline()

	p.AcceptRunWhitespace()

	q := p.NewGoal()

	if err := l.Expression.parse(q); err != nil {
		return p.Error("LambdaExpression", err)
	}

	p.Score(q)

	l.Tokens = p.ToTokens()

	return nil
}

// OrTest as defined in python@3.13.0:
// https://docs.python.org/release/3.13.0/reference/expressions.html#grammar-token-python-grammar-or_test
//
// When in a multiline structure, comments are parsed on either side of the
// 'or' operator.
type OrTest struct {
	AndTest  AndTest
	OrTest   *OrTest
	Comments [2]Comments
	Tokens   Tokens
}

func (o *OrTest) parse(p *pyParser) error {
	q := p.NewGoal()

	if err := o.AndTest.parse(q); err != nil {
		return p.Error("OrTest", err)
	}

	p.Score(q)

	q = p.NewGoal()

	q.AcceptRunWhitespace()

	if q.AcceptToken(parser.Token{Type: TokenKeyword, Data: "or"}) {
		o.Comments[0] = p.AcceptRunWhitespaceCommentsIfMultiline()

		p.AcceptRunWhitespace()
		p.Next()

		o.Comments[1] = p.AcceptRunWhitespaceCommentsIfMultiline()

		p.AcceptRunWhitespace()

		q = p.NewGoal()
		o.OrTest = new(OrTest)

		if err := o.OrTest.parse(q); err != nil {
			return p.Error("OrTest", err)
		}

		p.Score(q)
	}

	o.Tokens = p.ToTokens()

	return nil
}

func skipOrTest(p *pyParser) {
	skipAndTest(p)
	p.AcceptRunWhitespace()

	if p.AcceptToken(parser.Token{Type: TokenKeyword, Data: "or"}) {
		p.AcceptRunWhitespace()
		skipOrTest(p)
	}
}

// AndTest as defined in python@3.13.0:
// https://docs.python.org/release/3.13.0/reference/expressions.html#grammar-token-python-grammar-and_test
//
// When in a multiline structure, comments are parsed on either side of the
// 'and' operator.
type AndTest struct {
	NotTest  NotTest
	AndTest  *AndTest
	Comments [2]Comments
	Tokens   Tokens
}

func (a *AndTest) parse(p *pyParser) error {
	q := p.NewGoal()

	if err := a.NotTest.parse(q); err != nil {
		return p.Error("AndTest", err)
	}

	p.Score(q)

	q = p.NewGoal()

	q.AcceptRunWhitespace()

	if q.AcceptToken(parser.Token{Type: TokenKeyword, Data: "and"}) {
		a.Comments[0] = p.AcceptRunWhitespaceCommentsIfMultiline()

		p.AcceptRunWhitespace()
		p.Next()

		a.Comments[1] = p.AcceptRunWhitespaceCommentsIfMultiline()

		p.AcceptRunWhitespace()

		q = p.NewGoal()
		a.AndTest = new(AndTest)

		if err := a.AndTest.parse(q); err != nil {
			return p.Error("AndTest", err)
		}

		p.Score(q)
	}

	a.Tokens = p.ToTokens()

	return nil
}

func skipAndTest(p *pyParser) {
	skipNotTest(p)
	p.AcceptRunWhitespace()

	if p.AcceptToken(parser.Token{Type: TokenKeyword, Data: "and"}) {
		p.AcceptRunWhitespace()
		skipAndTest(p)
	}
}

// NotTest as defined in python@3.13.0:
// https://docs.python.org/release/3.13.0/reference/expressions.html#grammar-token-python-grammar-not_test
//
// When in a multiline structure, comments are parsed after every 'not' operator
type NotTest struct {
	Nots       []Comments
	Comparison Comparison
	Tokens     Tokens
}

func (n *NotTest) parse(p *pyParser) error {
	for p.AcceptToken(parser.Token{Type: TokenKeyword, Data: "not"}) {
		n.Nots = append(n.Nots, p.AcceptRunWhitespaceCommentsIfMultiline())
		p.AcceptRunWhitespace()
	}

	q := p.NewGoal()

	if err := n.Comparison.parse(q); err != nil {
		return p.Error("NotTest", err)
	}

	p.Score(q)

	n.Tokens = p.ToTokens()

	return nil
}

func skipNotTest(p *pyParser) {
	for p.AcceptToken(parser.Token{Type: TokenKeyword, Data: "not"}) {
		p.AcceptRunWhitespace()
	}

	skipComparison(p)
}

// Comparison as defined in python@3.13.0:
// https://docs.python.org/release/3.13.0/reference/expressions.html#grammar-token-python-grammar-comparison
type Comparison struct {
	OrExpression OrExpression
	Comparisons  []ComparisonExpression
	Tokens       Tokens
}

func (c *Comparison) parse(p *pyParser) error {
	q := p.NewGoal()

	if err := c.OrExpression.parse(q); err != nil {
		return p.Error("Comparison", err)
	}

	p.Score(q)

Loop:
	for {
		q = p.NewGoal()

		var ce ComparisonExpression

		ce.Comments[0] = q.AcceptRunWhitespaceCommentsIfMultiline()

		q.AcceptRunWhitespace()

		switch q.Peek() {
		case parser.Token{Type: TokenOperator, Data: "<"},
			parser.Token{Type: TokenOperator, Data: ">"},
			parser.Token{Type: TokenOperator, Data: "=="},
			parser.Token{Type: TokenOperator, Data: ">="},
			parser.Token{Type: TokenOperator, Data: "<="},
			parser.Token{Type: TokenOperator, Data: "!="},
			parser.Token{Type: TokenKeyword, Data: "in"}:
			p.Score(q)

			q = p.NewGoal()

			q.Next()

			ce.ComparisonOperator = q.ToTokens()

			p.Score(q)
		case parser.Token{Type: TokenKeyword, Data: "is"}:
			p.Score(q)

			q = p.NewGoal()

			q.Next()

			r := q.NewGoal()

			r.AcceptRunWhitespace()

			if r.AcceptToken(parser.Token{Type: TokenKeyword, Data: "not"}) {
				ce.Comments[1] = q.AcceptRunWhitespaceCommentsIfMultiline()

				q.AcceptRunWhitespace()
				q.Next()
			}

			ce.ComparisonOperator = q.ToTokens()

			p.Score(q)
		case parser.Token{Type: TokenKeyword, Data: "not"}:
			p.Score(q)

			q = p.NewGoal()

			q.Next()

			ce.Comments[1] = q.AcceptRunWhitespaceCommentsIfMultiline()

			q.AcceptRunWhitespace()

			if !q.AcceptToken(parser.Token{Type: TokenKeyword, Data: "in"}) {
				return q.Error("Comparison", ErrMissingIn)
			}

			ce.ComparisonOperator = q.ToTokens()

			p.Score(q)
		default:
			break Loop
		}

		ce.Comments[2] = p.AcceptRunWhitespaceCommentsIfMultiline()

		p.AcceptRunWhitespace()

		q = p.NewGoal()

		if err := ce.OrExpression.parse(q); err != nil {
			return p.Error("Comparison", err)
		}

		p.Score(q)

		c.Comparisons = append(c.Comparisons, ce)
	}

	c.Tokens = p.ToTokens()

	return nil
}

func skipComparison(p *pyParser) {
	skipOrExpression(p)
	p.AcceptRunWhitespace()

	for {
		switch p.Peek() {
		default:
			return
		case parser.Token{Type: TokenOperator, Data: "<"},
			parser.Token{Type: TokenOperator, Data: ">"},
			parser.Token{Type: TokenOperator, Data: "=="},
			parser.Token{Type: TokenOperator, Data: ">="},
			parser.Token{Type: TokenOperator, Data: "<="},
			parser.Token{Type: TokenOperator, Data: "!="},
			parser.Token{Type: TokenKeyword, Data: "in"}:
			p.Next()
		case parser.Token{Type: TokenKeyword, Data: "is"}:
			p.Next()
			p.AcceptRunWhitespace()
			p.AcceptToken(parser.Token{Type: TokenKeyword, Data: "not"})
		case parser.Token{Type: TokenKeyword, Data: "not"}:
			p.Next()
			p.AcceptRunWhitespace()
			p.AcceptToken(parser.Token{Type: TokenKeyword, Data: "in"})
		}

		p.AcceptRunWhitespace()
		skipOrExpression(p)
		p.AcceptRunWhitespace()
	}
}

// ComparisonExpression combines combines the operators with an OrExpression.
type ComparisonExpression struct {
	ComparisonOperator []Token
	OrExpression       OrExpression
	Comments           [3]Comments
}

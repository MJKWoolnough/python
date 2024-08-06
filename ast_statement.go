package python

import (
	"slices"

	"vimagination.zapto.org/parser"
)

type Statement struct {
	StatementList     StatementList
	CompoundStatement *CompoundStatement
	Tokens            Tokens
}

func (s *Statement) parse(p *pyParser) error {
	var (
		isCompound     bool
		isSoftCompound bool
	)

	q := p.NewGoal()

	switch tk := p.Peek(); tk.Type {
	case TokenDelimiter:
		isCompound = tk.Data == "@"
	case TokenKeyword:
		isCompound = slices.Contains(compounds[:], tk.Data)
	case TokenIdentifier:
		isCompound = tk.Data == "match"
		isSoftCompound = isCompound
	}

	if isCompound {
		var c CompoundStatement

		if err := c.parser(q); err != nil {
			if !isSoftCompound {
				return p.Error("Statement", err)
			}
		} else {
			p.Score(q)

			s.CompoundStatement = &c
			s.Tokens = p.ToTokens()

			return nil
		}
	}

	if err := s.StatementList.parse(q); err != nil {
		return p.Error("Statement", err)
	}

	return nil
}

type StatementList struct {
	Statements []SimpleStatement
	Tokens
}

func (s *StatementList) parse(p *pyParser) error {
	for {
		var ss SimpleStatement

		q := p.NewGoal()

		if err := ss.parse(q); err != nil {
			return p.Error("StatementList", err)
		}

		p.Score(q)

		s.Statements = append(s.Statements, ss)

		q = p.NewGoal()

		q.AcceptRun(TokenWhitespace)

		if !q.AcceptToken(parser.Token{Type: TokenDelimiter, Data: ";"}) {
			break
		}

		p.Score(q)
	}

	s.Tokens = p.ToTokens()

	return nil
}

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

type SimpleStatement struct {
	Type                StatementType
	AssertStatement     *AssertStatement
	AssignmentStatement *AssignmentStatement
	DelStatement        *DelStatement
	ReturnStatement     *ReturnStatement
	YieldStatement      *YieldStatement
	RaiseStatement      *RaiseStatement
	ImportStatement     *ImportStatement
	GlobalStatement     *GlobalStatement
	NonLocalStatement   *NonLocalStatement
	TypeStatement       *TypeStatement
	Tokens              Tokens
}

func (s *SimpleStatement) parse(p *pyParser) error {
	switch p.Peek() {
	case parser.Token{Type: TokenKeyword, Data: "assert"}:
		s.AssertStatement = new(AssertStatement)
		s.Type = StatementAssert

		q := p.NewGoal()

		if err := s.AssertStatement.parse(q); err != nil {
			return p.Error("SimpleStatement", err)
		}

		p.Score(q)
	case parser.Token{Type: TokenKeyword, Data: "pass"}:
		p.Skip()

		s.Type = StatementPass
	case parser.Token{Type: TokenKeyword, Data: "del"}:
		s.DelStatement = new(DelStatement)
		s.Type = StatementDel

		q := p.NewGoal()

		if err := s.DelStatement.parse(q); err != nil {
			return p.Error("SimpleStatement", err)
		}

		p.Score(q)
	case parser.Token{Type: TokenKeyword, Data: "return"}:
		s.ReturnStatement = new(ReturnStatement)
		s.Type = StatementReturn

		q := p.NewGoal()

		if err := s.ReturnStatement.parse(q); err != nil {
			return p.Error("SimpleStatement", err)
		}

		p.Score(q)
	case parser.Token{Type: TokenKeyword, Data: "yield"}:
		s.YieldStatement = new(YieldStatement)
		s.Type = StatementYield

		q := p.NewGoal()

		if err := s.YieldStatement.parse(q); err != nil {
			return p.Error("SimpleStatement", err)
		}

		p.Score(q)
	case parser.Token{Type: TokenKeyword, Data: "raise"}:
		s.RaiseStatement = new(RaiseStatement)
		s.Type = StatementRaise

		q := p.NewGoal()

		if err := s.RaiseStatement.parse(q); err != nil {
			return p.Error("SimpleStatement", err)
		}

		p.Score(q)
	case parser.Token{Type: TokenKeyword, Data: "break"}:
		p.Skip()

		s.Type = StatementBreak
	case parser.Token{Type: TokenKeyword, Data: "continue"}:
		p.Skip()

		s.Type = StatementContinue
	case parser.Token{Type: TokenKeyword, Data: "import"}, parser.Token{Type: TokenKeyword, Data: "from"}:
		s.ImportStatement = new(ImportStatement)
		s.Type = StatementImport

		q := p.NewGoal()

		if err := s.ImportStatement.parse(q); err != nil {
			return p.Error("SimpleStatement", err)
		}

		p.Score(q)
	case parser.Token{Type: TokenKeyword, Data: "global"}:
		s.GlobalStatement = new(GlobalStatement)
		s.Type = StatementGlobal

		q := p.NewGoal()

		if err := s.GlobalStatement.parse(q); err != nil {
			return p.Error("SimpleStatement", err)
		}

		p.Score(q)
	case parser.Token{Type: TokenKeyword, Data: "nonlocal"}:
		s.NonLocalStatement = new(NonLocalStatement)
		s.Type = StatementNonLocal

		q := p.NewGoal()

		if err := s.NonLocalStatement.parse(q); err != nil {
			return p.Error("SimpleStatement", err)
		}

		p.Score(q)
	case parser.Token{Type: TokenKeyword, Data: "type"}:
		s.TypeStatement = new(TypeStatement)
		s.Type = StatementTyp

		q := p.NewGoal()

		if err := s.TypeStatement.parse(q); err != nil {
			return p.Error("SimpleStatement", err)
		}

		p.Score(q)
	default:
		s.AssignmentStatement = new(AssignmentStatement)
		s.Type = StatementAssignment

		q := p.NewGoal()

		if err := s.AssignmentStatement.parse(q); err != nil {
			return p.Error("SimpleStatement", err)
		}

		p.Score(q)
	}

	s.Tokens = p.ToTokens()

	return nil
}

type AssertStatement struct {
	Expressions []Expression
	Tokens      Tokens
}

func (a *AssertStatement) parse(p *pyParser) error {
	p.Skip()

	for {
		p.AcceptRun(TokenWhitespace)

		var e Expression

		q := p.NewGoal()

		if err := e.parse(q, whitespaceToken); err != nil {
			return p.Error("AssertStatement", err)
		}

		a.Expressions = append(a.Expressions, e)

		p.Score(q)

		q = p.NewGoal()

		q.AcceptRun(TokenWhitespace)

		if !q.AcceptToken(parser.Token{Type: TokenDelimiter, Data: ","}) {
			break
		}

		p.Score(q)
	}

	a.Tokens = p.ToTokens()
	return nil
}

type AssignmentStatement struct{}

func (a *AssignmentStatement) parse(_ *pyParser) error {
	return nil
}

type DelStatement struct {
	TargetList TargetList
	Tokens     Tokens
}

func (d *DelStatement) parse(p *pyParser) error {
	p.Skip()

	p.AcceptRun(TokenWhitespace)

	q := p.NewGoal()

	if err := d.TargetList.parse(q, whitespaceToken); err != nil {
		return p.Error("DelStatement", err)
	}

	p.Score(q)

	d.Tokens = p.ToTokens()

	return nil
}

type ReturnStatement struct {
	Expression Expression
	From       *Expression
	Tokens     Tokens
}

func (r *ReturnStatement) parse(p *pyParser) error {
	p.Skip()

	p.AcceptRun(TokenWhitespace)

	q := p.NewGoal()

	if err := r.Expression.parse(q, whitespaceToken); err != nil {
		return p.Error("ReturnStatement", err)
	}

	p.Score(q)

	q = p.NewGoal()

	q.AcceptRun(TokenWhitespace)

	if q.AcceptToken(parser.Token{Type: TokenKeyword, Data: "from"}) {
		q.AcceptRun(TokenWhitespace)
		p.Score(q)

		q = p.NewGoal()
		r.From = new(Expression)

		if err := r.From.parse(q, whitespaceToken); err != nil {
			return p.Error("ReturnStatement", err)
		}

		p.Score(q)
	}

	r.Tokens = p.ToTokens()

	return nil
}

type YieldStatement struct {
	Expression *Expression
	From       *ExpressionList
	Tokens     Tokens
}

func (y *YieldStatement) parse(p *pyParser) error {
	p.Skip()

	p.AcceptRun(TokenWhitespace)

	if p.AcceptToken(parser.Token{Type: TokenKeyword, Data: "from"}) {
		y.From = new(ExpressionList)

		p.AcceptRun(TokenWhitespace)

		q := p.NewGoal()

		if err := y.From.parse(q); err != nil {
			return p.Error("YieldStatement", err)
		}

		p.Score(p)
	} else {
		y.Expression = new(Expression)

		p.AcceptRun(TokenWhitespace)

		q := p.NewGoal()

		if err := y.Expression.parse(q, whitespaceToken); err != nil {
			return p.Error("YieldStatement", err)
		}

		p.Score(p)
	}

	y.Tokens = p.ToTokens()

	return nil
}

type RaiseStatement struct {
	Expression *Expression
	From       *ExpressionList
	Tokens     Tokens
}

func (r *RaiseStatement) parse(p *pyParser) error {
	p.Skip()

	p.AcceptRun(TokenWhitespace)

	q := p.NewGoal()

	switch q.AcceptRun(TokenWhitespace) {
	case TokenLineTerminator, TokenComment:
	default:
		p.Score(q)

		q = p.NewGoal()
		r.Expression = new(Expression)

		if err := r.Expression.parse(q, whitespaceToken); err != nil {
			return p.Error("RaiseStatement", err)
		}

		p.Score(q)

		q = p.NewGoal()

		q.AcceptRun(TokenWhitespace)

		if q.AcceptToken(parser.Token{Type: TokenKeyword, Data: "from"}) {
			q.AcceptRun(TokenWhitespace)
			p.Score(q)

			q = p.NewGoal()
			r.From = new(ExpressionList)

			if err := r.From.parse(q); err != nil {
				return p.Error("RaiseStatement", err)
			}

			p.Score(q)
		}
	}

	r.Tokens = p.ToTokens()

	return nil
}

type ImportStatement struct{}

func (i *ImportStatement) parse(_ *pyParser) error {
	return nil
}

type GlobalStatement struct{}

func (g *GlobalStatement) parse(_ *pyParser) error {
	return nil
}

type NonLocalStatement struct{}

func (n *NonLocalStatement) parse(_ *pyParser) error {
	return nil
}

type TypeStatement struct{}

func (t *TypeStatement) parse(_ *pyParser) error {
	return nil
}

type AssignmentExpression struct{}

func (s *AssignmentExpression) parse(_ *pyParser) error {
	return nil
}

type Expression struct{}

func (s *Expression) parse(_ *pyParser, _ws []parser.TokenType) error {
	return nil
}

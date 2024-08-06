package python

import (
	"errors"
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
	case parser.Token{Type: TokenIdentifier, Data: "type"}:
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

type ImportStatement struct {
	RelativeModule *RelativeModule
	Modules        []ModuleAs
	Tokens         Tokens
}

func (i *ImportStatement) parse(p *pyParser) error {
	if p.AcceptToken(parser.Token{Type: TokenKeyword, Data: "from"}) {
		p.AcceptRun(TokenWhitespace)

		q := p.NewGoal()
		i.RelativeModule = new(RelativeModule)

		if err := i.RelativeModule.parse(q); err != nil {
			return p.Error("ImportStatement", err)
		}

		p.Score(q)
		p.AcceptRun(TokenWhitespace)

		if !p.AcceptToken(parser.Token{Type: TokenKeyword, Data: "import"}) {
			return p.Error("ImportStatement", ErrMissingImport)
		}
	} else {
		p.Skip()
	}

	p.AcceptRun(TokenWhitespace)

	if i.RelativeModule == nil || !p.AcceptToken(parser.Token{Type: TokenOperator, Data: "*"}) {
		parens := p.AcceptToken(parser.Token{Type: TokenDelimiter, Data: "("})

		ws := whitespaceToken
		if parens {
			ws = whitespaceCommentTokens
			p.AcceptRun(ws...)
		}

		for {

			q := p.NewGoal()

			var module ModuleAs

			if err := module.parse(q, ws); err != nil {
				return p.Error("ImportStatement", err)
			}

			p.Score(q)

			q = p.NewGoal()

			q.AcceptRun(ws...)

			if !q.AcceptToken(parser.Token{Type: TokenDelimiter, Data: ","}) && !parens {
				break
			}

			q.AcceptRun(ws...)

			if parens && q.AcceptToken(parser.Token{Type: TokenDelimiter, Data: ")"}) {
				p.Score(q)

				break
			}
		}
	}

	i.Tokens = p.ToTokens()

	return nil
}

type RelativeModule struct {
	Dots   int
	Module *Module
	Tokens Tokens
}

func (r *RelativeModule) parse(p *pyParser) error {
	dots := 0

	for p.AcceptToken(parser.Token{Type: TokenDelimiter, Data: "."}) {
		dots++
	}

	r.Dots = dots

	q := p.NewGoal()

	switch q.AcceptRun(TokenWhitespace) {
	case TokenLineTerminator, TokenComment:
		if dots == 0 {
			return q.Error("RelativeModule", ErrMissingModule)
		}
	default:
		p.Score(q)

		q = p.NewGoal()
		r.Module = new(Module)

		if err := r.Module.parse(q, whitespaceToken); err != nil {
			return p.Error("RelativeModule", err)
		}

		p.Score(q)
	}

	r.Tokens = p.ToTokens()

	return nil
}

type ModuleAs struct {
	Module Module
	As     *Token
	Tokens Tokens
}

func (m *ModuleAs) parse(p *pyParser, ws []parser.TokenType) error {
	q := p.NewGoal()

	if err := m.Module.parse(q, ws); err != nil {
		return p.Error("ModuleAs", err)
	}

	p.Score(q)

	q = p.NewGoal()

	q.AcceptRun(ws...)

	if q.AcceptToken(parser.Token{Type: TokenKeyword, Data: "as"}) {
		q.AcceptRun(ws...)
		p.Score(q)

		if !p.Accept(TokenIdentifier) {
			return p.Error("ModuleAs", ErrMissingIdentifier)
		}

		m.As = p.GetLastToken()
	}

	m.Tokens = p.ToTokens()

	return nil
}

type Module struct {
	Identifiers []*Token
	Tokens      Tokens
}

func (m *Module) parse(p *pyParser, ws []parser.TokenType) error {
	for {
		if !p.Accept(TokenIdentifier) {
			return p.Error("Module", ErrMissingIdentifier)
		}

		m.Identifiers = append(m.Identifiers, p.GetLastToken())

		q := p.NewGoal()

		q.AcceptRun(ws...)

		if !q.AcceptToken(parser.Token{Type: TokenDelimiter, Data: "."}) {
			break
		}

		q.AcceptRun(ws...)
		p.Score(q)
	}

	m.Tokens = p.ToTokens()

	return nil
}

type GlobalStatement struct {
	Identifiers []*Token
	Tokens      Tokens
}

func (g *GlobalStatement) parse(p *pyParser) error {
	p.Skip()
	p.AcceptRun(TokenWhitespace)

	for {
		if !p.Accept(TokenIdentifier) {
			return p.Error("GlobalStatement", ErrMissingIdentifier)
		}

		g.Identifiers = append(g.Identifiers, p.GetLastToken())

		q := p.NewGoal()

		q.AcceptRun(TokenWhitespace)

		if !q.AcceptToken(parser.Token{Type: TokenDelimiter, Data: ","}) {
			break
		}

		q.AcceptRun(TokenWhitespace)
		p.Score(q)
	}

	g.Tokens = p.ToTokens()

	return nil
}

type NonLocalStatement struct {
	Identifiers []*Token
	Tokens      Tokens
}

func (n *NonLocalStatement) parse(p *pyParser) error {
	p.Skip()
	p.AcceptRun(TokenWhitespace)

	for {
		if !p.Accept(TokenIdentifier) {
			return p.Error("NonLocalStatement", ErrMissingIdentifier)
		}

		n.Identifiers = append(n.Identifiers, p.GetLastToken())

		q := p.NewGoal()

		q.AcceptRun(TokenWhitespace)

		if !q.AcceptToken(parser.Token{Type: TokenDelimiter, Data: ","}) {
			break
		}

		q.AcceptRun(TokenWhitespace)
		p.Score(q)
	}

	n.Tokens = p.ToTokens()

	return nil
}

type TypeStatement struct {
	Identifier *Token
	TypeParams *TypeParams
	Expression Expression
	Tokens     Tokens
}

func (t *TypeStatement) parse(p *pyParser) error {
	p.Skip()
	p.AcceptRun(TokenWhitespace)

	if !p.Accept(TokenIdentifier) {
		return p.Error("TypeStatement", ErrMissingIdentifier)
	}

	p.AcceptRun(TokenWhitespace)

	if p.Peek() == (parser.Token{Type: TokenDelimiter, Data: "["}) {
		q := p.NewGoal()
		t.TypeParams = new(TypeParams)

		if err := t.TypeParams.parse(q); err != nil {
			return p.Error("TypeStatement", err)
		}

		p.Score(q)
		p.AcceptRun(TokenWhitespace)
	}

	if !p.AcceptToken(parser.Token{Type: TokenOperator, Data: "="}) {
		return p.Error("TypeStatement", ErrMissingEquals)
	}

	p.AcceptRun(TokenWhitespace)

	q := p.NewGoal()

	if err := t.Expression.parse(q, whitespaceToken); err != nil {
		return p.Error("TypeStatement", err)
	}

	p.Score(q)

	t.Tokens = p.ToTokens()

	return nil
}

type TypeParams struct {
	TypeParams []TypeParam
	Tokens     Tokens
}

func (t *TypeParams) parse(p *pyParser) error {
	p.Skip()
	p.AcceptRun(TokenWhitespace, TokenComment)

	for {
		q := p.NewGoal()

		var tp TypeParam

		if err := tp.parse(q); err != nil {
			return p.Error("TypeParams", err)
		}

		t.TypeParams = append(t.TypeParams, tp)

		p.Score(q)

		p.AcceptRun(TokenWhitespace, TokenComment)

		if p.AcceptToken(parser.Token{Type: TokenDelimiter, Data: "]"}) {
			break
		}
	}

	t.Tokens = p.ToTokens()
	return nil
}

type AssignmentExpression struct{}

func (s *AssignmentExpression) parse(_ *pyParser) error {
	return nil
}

type Expression struct{}

func (s *Expression) parse(_ *pyParser, _ []parser.TokenType) error {
	return nil
}

var (
	ErrMissingImport = errors.New("missing import keyword")
	ErrMissingModule = errors.New("missing module")
	ErrMissingEquals = errors.New("missing equals")
)

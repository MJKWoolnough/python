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

		q.AcceptRunWhitespace()

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
	Type                         StatementType
	AssertStatement              *AssertStatement
	ExpressionStatement          *StarredExpression
	AssignmentStatement          *AssignmentStatement
	AugmentedAssignmentStatement *AugmentedAssignmentStatement
	AnnotatedAssignmentStatement *AnnotatedAssignmentStatement
	DelStatement                 *DelStatement
	ReturnStatement              *ReturnStatement
	YieldStatement               *YieldExpression
	RaiseStatement               *RaiseStatement
	ImportStatement              *ImportStatement
	GlobalStatement              *GlobalStatement
	NonLocalStatement            *NonLocalStatement
	TypeStatement                *TypeStatement
	Tokens                       Tokens
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
		s.YieldStatement = new(YieldExpression)
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
		q := p.NewGoal()

		if n := p.LookaheadLine(
			parser.Token{Type: TokenKeyword, Data: "lambda"},
			parser.Token{Type: TokenDelimiter, Data: ":"},
			parser.Token{Type: TokenDelimiter, Data: "+="},
			parser.Token{Type: TokenDelimiter, Data: "-="},
			parser.Token{Type: TokenDelimiter, Data: "*="},
			parser.Token{Type: TokenDelimiter, Data: "@="},
			parser.Token{Type: TokenDelimiter, Data: "/="},
			parser.Token{Type: TokenDelimiter, Data: "//="},
			parser.Token{Type: TokenDelimiter, Data: "%="},
			parser.Token{Type: TokenDelimiter, Data: "**="},
			parser.Token{Type: TokenDelimiter, Data: ">>="},
			parser.Token{Type: TokenDelimiter, Data: "<<="},
			parser.Token{Type: TokenDelimiter, Data: "&="},
			parser.Token{Type: TokenDelimiter, Data: "^="},
			parser.Token{Type: TokenDelimiter, Data: "|="},
		); n <= 0 {
			s.Type = StatementAssignment
			s.AssignmentStatement = new(AssignmentStatement)

			if err := s.AssignmentStatement.parse(q); err != nil {
				return p.Error("SimpleStatement", err)
			}
		} else if n == 1 {
			s.Type = StatementAnnotatedAssignment
			s.AnnotatedAssignmentStatement = new(AnnotatedAssignmentStatement)

			if err := s.AugmentedAssignmentStatement.parse(q); err != nil {
				return p.Error("SimpleStatement", err)
			}
		} else {
			s.Type = StatementAugmentedAssignment
			s.AugmentedAssignmentStatement = new(AugmentedAssignmentStatement)

			if err := s.AugmentedAssignmentStatement.parse(q); err != nil {
				return p.Error("SimpleStatement", err)
			}
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
		p.AcceptRunWhitespace()

		var e Expression

		q := p.NewGoal()

		if err := e.parse(q); err != nil {
			return p.Error("AssertStatement", err)
		}

		a.Expressions = append(a.Expressions, e)

		p.Score(q)

		q = p.NewGoal()

		q.AcceptRunWhitespace()

		if !q.AcceptToken(parser.Token{Type: TokenDelimiter, Data: ","}) {
			break
		}

		p.Score(q)
	}

	a.Tokens = p.ToTokens()
	return nil
}

type AssignmentStatement struct {
	TargetLists       []TargetList
	StarredExpression *StarredExpression
	YieldExpression   *YieldExpression
	Tokens            Tokens
}

func (a *AssignmentStatement) parse(p *pyParser) error {
	q := p.NewGoal()

	for {
		if q.LookaheadLine(parser.Token{Type: TokenDelimiter, Data: "="}, parser.Token{Type: TokenKeyword, Data: "lambda"}) != 0 {
			break
		}

		var tl TargetList

		if err := tl.parse(q); err != nil {
			return p.Error("AssignmentStatement", err)
		}

		p.Score(q)
		p.AcceptRunWhitespace()
		p.AcceptToken(parser.Token{Type: TokenDelimiter, Data: "="})
		p.AcceptRunWhitespace()

		a.TargetLists = append(a.TargetLists, tl)
		q = p.NewGoal()
	}

	if p.Peek() == (parser.Token{Type: TokenKeyword, Data: "yield"}) {
		a.YieldExpression = new(YieldExpression)

		if err := a.YieldExpression.parse(q); err != nil {
			return p.Error("AssignmentStatement", err)
		}
	} else {
		a.StarredExpression = new(StarredExpression)

		if err := a.StarredExpression.parse(q); err != nil {
			return p.Error("AssignmentStatement", err)
		}
	}

	p.Score(q)

	a.Tokens = p.ToTokens()

	return nil
}

type AugmentedAssignmentStatement struct {
	AugTarget       AugTarget
	AugOp           *Token
	ExpressionList  *ExpressionList
	YieldExpression *YieldExpression
	Tokens          Tokens
}

func (a *AugmentedAssignmentStatement) parse(p *pyParser) error {
	q := p.NewGoal()

	if err := a.AugTarget.parse(q); err != nil {
		return p.Error("AugmentedAssignmentStatement", err)
	}

	p.Score(q)
	p.AcceptRunWhitespace()

	if p.Accept(TokenDelimiter) {
		t := p.GetLastToken()

		switch t.Data {
		case "+=", "-=", "*=", "@=", "/=", "//=", "%=", "**=", ">>=", "<<=", "&=", "^=", "|=":
			a.AugOp = t
		}
	}

	if a.AugOp == nil {
		return p.Error("AugmentedAssignmentStatement", ErrMissingOp)
	}

	p.AcceptRunWhitespace()

	q = p.NewGoal()

	if q.Peek() == (parser.Token{Type: TokenKeyword, Data: "yield"}) {
		a.YieldExpression = new(YieldExpression)

		if err := a.YieldExpression.parse(q); err != nil {
			return p.Error("AugmentedAssignmentStatement", err)
		}
	} else {
		a.ExpressionList = new(ExpressionList)

		if err := a.ExpressionList.parse(q); err != nil {
			return p.Error("AugmentedAssignmentStatement", err)
		}
	}

	p.Score(q)

	a.Tokens = p.ToTokens()

	return nil
}

type AugTarget struct {
	PrimaryExpression PrimaryExpression
	AttributeRef      *Token
	Slicing           *SliceList
	Tokens            Tokens
}

func (a *AugTarget) parse(p *pyParser) error {
	q := p.NewGoal()

	if err := a.PrimaryExpression.parse(q); err != nil {
		return p.Error("AugTarget", err)
	}

	p.Score(q)

	q = p.NewGoal()

	q.AcceptRunWhitespace()

	if q.AcceptToken(parser.Token{Type: TokenDelimiter, Data: "."}) {
		q.AcceptRunWhitespace()
		p.Score(q)

		if !p.Accept(TokenIdentifier) {
			return p.Error("AugTarget", ErrMissingIdentifier)
		}

		a.AttributeRef = p.GetLastToken()
	} else if q.AcceptToken(parser.Token{Type: TokenDelimiter, Data: "["}) {
		q.AcceptRunWhitespace()
		p.Score(q)

		q = p.NewGoal()
		a.Slicing = new(SliceList)

		if err := a.Slicing.parse(q); err != nil {
			return p.Error("AugTarget", err)
		}

		p.Score(q)
	} else if !a.PrimaryExpression.IsIdentifier() {
		return p.Error("AugTarget", ErrMissingIdentifier)
	}

	a.Tokens = p.ToTokens()

	return nil
}

type AnnotatedAssignmentStatement struct {
	AugTarget         AugTarget
	Expression        Expression
	StarredExpression *StarredExpression
	YieldExpression   *YieldExpression
	Tokens            Tokens
}

func (a *AnnotatedAssignmentStatement) parse(p *pyParser) error {
	q := p.NewGoal()

	if err := a.AugTarget.parse(q); err != nil {
		return p.Error("AnnotatedAssignmentStatement", err)
	}

	p.Score(q)
	p.AcceptRunWhitespace()

	if !p.AcceptToken(parser.Token{Type: TokenDelimiter, Data: ":"}) {
		return p.Error("AnnotatedAssignmentStatement", ErrMissingColon)
	}

	p.AcceptRunWhitespace()

	q = p.NewGoal()

	if err := a.Expression.parse(q); err != nil {
		return p.Error("AnnotatedAssignmentStatement", err)
	}

	p.Score(q)

	q = p.NewGoal()

	q.AcceptRunWhitespace()

	if q.AcceptToken(parser.Token{Type: TokenDelimiter, Data: "="}) {
		p.Score(q)
		p.AcceptRunWhitespace()

		q = p.NewGoal()

		if q.Peek() == (parser.Token{Type: TokenKeyword, Data: "yield"}) {
			a.YieldExpression = new(YieldExpression)

			if err := a.YieldExpression.parse(q); err != nil {
				return p.Error("AnnotatedAssignmentStatement", err)
			}

		} else {
			a.StarredExpression = new(StarredExpression)

			if err := a.StarredExpression.parse(q); err != nil {
				return p.Error("AnnotatedAssignmentStatement", err)
			}
		}

		p.Score(q)
	}

	a.Tokens = p.ToTokens()

	return nil
}

type StarredExpression struct {
	Expression   *Expression
	StarredItems []StarredItem
	Tokens       Tokens
}

func (s *StarredExpression) parse(p *pyParser) error {
	first := true

	for {
		var si StarredItem

		q := p.NewGoal()

		if err := si.parse(q); err != nil {
			return p.Error("StarredExpression", err)
		}

		p.Score(q)

		q = p.NewGoal()

		q.AcceptRunWhitespace()

		if first && q.Peek() != (parser.Token{Type: TokenDelimiter, Data: ","}) && si.AssignmentExpression != nil && si.AssignmentExpression.Identifier == nil {
			s.Expression = &si.AssignmentExpression.Expression

			break
		}

		first = false

		if !q.AcceptToken(parser.Token{Type: TokenDelimiter, Data: ","}) {
			break
		}

		p.Score(q)

		q = p.NewGoal()

		q.AcceptRunWhitespace()

		if t := q.Peek(); t == (parser.Token{Type: TokenOperator, Data: "="}) || t.Type == TokenLineTerminator || t.Type == TokenComment {
			break
		}

		p.Score(q)
	}

	s.Tokens = p.ToTokens()

	return nil
}

type DelStatement struct {
	TargetList TargetList
	Tokens     Tokens
}

func (d *DelStatement) parse(p *pyParser) error {
	p.Skip()

	p.AcceptRunWhitespace()

	q := p.NewGoal()

	if err := d.TargetList.parse(q); err != nil {
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

	p.AcceptRunWhitespace()

	q := p.NewGoal()

	if err := r.Expression.parse(q); err != nil {
		return p.Error("ReturnStatement", err)
	}

	p.Score(q)

	q = p.NewGoal()

	q.AcceptRunWhitespace()

	if q.AcceptToken(parser.Token{Type: TokenKeyword, Data: "from"}) {
		q.AcceptRunWhitespace()
		p.Score(q)

		q = p.NewGoal()
		r.From = new(Expression)

		if err := r.From.parse(q); err != nil {
			return p.Error("ReturnStatement", err)
		}

		p.Score(q)
	}

	r.Tokens = p.ToTokens()

	return nil
}

type YieldExpression struct {
	Expression *Expression
	From       *ExpressionList
	Tokens     Tokens
}

func (y *YieldExpression) parse(p *pyParser) error {
	p.Skip()

	p.AcceptRunWhitespace()

	if p.AcceptToken(parser.Token{Type: TokenKeyword, Data: "from"}) {
		y.From = new(ExpressionList)

		p.AcceptRunWhitespace()

		q := p.NewGoal()

		if err := y.From.parse(q); err != nil {
			return p.Error("YieldExpression", err)
		}

		p.Score(p)
	} else {
		y.Expression = new(Expression)

		p.AcceptRunWhitespace()

		q := p.NewGoal()

		if err := y.Expression.parse(q); err != nil {
			return p.Error("YieldExpression", err)
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

	p.AcceptRunWhitespace()

	q := p.NewGoal()

	switch q.AcceptRunWhitespace() {
	case TokenLineTerminator, TokenComment:
	default:
		p.Score(q)

		q = p.NewGoal()
		r.Expression = new(Expression)

		if err := r.Expression.parse(q); err != nil {
			return p.Error("RaiseStatement", err)
		}

		p.Score(q)

		q = p.NewGoal()

		q.AcceptRunWhitespace()

		if q.AcceptToken(parser.Token{Type: TokenKeyword, Data: "from"}) {
			q.AcceptRunWhitespace()
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
		p.AcceptRunWhitespace()

		q := p.NewGoal()
		i.RelativeModule = new(RelativeModule)

		if err := i.RelativeModule.parse(q); err != nil {
			return p.Error("ImportStatement", err)
		}

		p.Score(q)
		p.AcceptRunWhitespace()

		if !p.AcceptToken(parser.Token{Type: TokenKeyword, Data: "import"}) {
			return p.Error("ImportStatement", ErrMissingImport)
		}
	} else {
		p.Skip()
	}

	p.AcceptRunWhitespace()

	if i.RelativeModule == nil || !p.AcceptToken(parser.Token{Type: TokenOperator, Data: "*"}) {
		parens := p.AcceptToken(parser.Token{Type: TokenDelimiter, Data: "("})

		if parens {
			p.OpenBrackets()

			p.AcceptRunWhitespace()
		}

		for {

			q := p.NewGoal()

			var module ModuleAs

			if err := module.parse(q); err != nil {
				return p.Error("ImportStatement", err)
			}

			p.Score(q)

			q = p.NewGoal()

			q.AcceptRunWhitespace()

			if !q.AcceptToken(parser.Token{Type: TokenDelimiter, Data: ","}) && !parens {
				break
			}

			q.AcceptRunWhitespace()

			if parens && q.AcceptToken(parser.Token{Type: TokenDelimiter, Data: ")"}) {
				p.Score(q)
				p.CloseBrackets()

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

	switch q.AcceptRunWhitespace() {
	case TokenLineTerminator, TokenComment:
		if dots == 0 {
			return q.Error("RelativeModule", ErrMissingModule)
		}
	default:
		p.Score(q)

		q = p.NewGoal()
		r.Module = new(Module)

		if err := r.Module.parse(q); err != nil {
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

func (m *ModuleAs) parse(p *pyParser) error {
	q := p.NewGoal()

	if err := m.Module.parse(q); err != nil {
		return p.Error("ModuleAs", err)
	}

	p.Score(q)

	q = p.NewGoal()

	q.AcceptRunWhitespace()

	if q.AcceptToken(parser.Token{Type: TokenKeyword, Data: "as"}) {
		q.AcceptRunWhitespace()
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

func (m *Module) parse(p *pyParser) error {
	for {
		if !p.Accept(TokenIdentifier) {
			return p.Error("Module", ErrMissingIdentifier)
		}

		m.Identifiers = append(m.Identifiers, p.GetLastToken())

		q := p.NewGoal()

		q.AcceptRunWhitespace()

		if !q.AcceptToken(parser.Token{Type: TokenDelimiter, Data: "."}) {
			break
		}

		q.AcceptRunWhitespace()
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
	p.AcceptRunWhitespace()

	for {
		if !p.Accept(TokenIdentifier) {
			return p.Error("GlobalStatement", ErrMissingIdentifier)
		}

		g.Identifiers = append(g.Identifiers, p.GetLastToken())

		q := p.NewGoal()

		q.AcceptRunWhitespace()

		if !q.AcceptToken(parser.Token{Type: TokenDelimiter, Data: ","}) {
			break
		}

		q.AcceptRunWhitespace()
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
	p.AcceptRunWhitespace()

	for {
		if !p.Accept(TokenIdentifier) {
			return p.Error("NonLocalStatement", ErrMissingIdentifier)
		}

		n.Identifiers = append(n.Identifiers, p.GetLastToken())

		q := p.NewGoal()

		q.AcceptRunWhitespace()

		if !q.AcceptToken(parser.Token{Type: TokenDelimiter, Data: ","}) {
			break
		}

		q.AcceptRunWhitespace()
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
	p.AcceptRunWhitespace()

	if !p.Accept(TokenIdentifier) {
		return p.Error("TypeStatement", ErrMissingIdentifier)
	}

	p.AcceptRunWhitespace()

	if p.Peek() == (parser.Token{Type: TokenDelimiter, Data: "["}) {
		q := p.NewGoal()
		t.TypeParams = new(TypeParams)

		if err := t.TypeParams.parse(q); err != nil {
			return p.Error("TypeStatement", err)
		}

		p.Score(q)
		p.AcceptRunWhitespace()
	}

	if !p.AcceptToken(parser.Token{Type: TokenDelimiter, Data: "="}) {
		return p.Error("TypeStatement", ErrMissingEquals)
	}

	p.AcceptRunWhitespace()

	q := p.NewGoal()

	if err := t.Expression.parse(q); err != nil {
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
	p.OpenBrackets()

	for {
		q := p.NewGoal()

		var tp TypeParam

		if err := tp.parse(q); err != nil {
			return p.Error("TypeParams", err)
		}

		t.TypeParams = append(t.TypeParams, tp)

		p.Score(q)

		p.AcceptRunWhitespace()

		if p.AcceptToken(parser.Token{Type: TokenDelimiter, Data: "]"}) {
			p.CloseBrackets()

			break
		}
	}

	t.Tokens = p.ToTokens()
	return nil
}

type AssignmentExpression struct {
	Identifier *Token
	Expression Expression
	Tokens
}

func (a *AssignmentExpression) parse(p *pyParser) error {
	q := p.NewGoal()

	if q.Accept(TokenIdentifier) {
		identifier := q.GetLastToken()

		q.AcceptRunWhitespace()

		if q.AcceptToken(parser.Token{Type: TokenOperator, Data: ":="}) {
			q.AcceptRunWhitespace()
			p.Score(q)

			a.Identifier = identifier
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

type ConditionalExpression struct {
	OrTest OrTest
	If     *OrTest
	Else   *Expression
	Tokens Tokens
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
		q.AcceptRunWhitespace()
		p.Score(q)

		q = p.NewGoal()
		c.If = new(OrTest)

		if err := c.If.parse(q); err != nil {
			return p.Error("ConditionalExpression", err)
		}

		p.Score(q)
		p.AcceptRunWhitespace()

		if !p.AcceptToken(parser.Token{Type: TokenKeyword, Data: "else"}) {
			return p.Error("ConditionalExpression", ErrMissingElse)
		}

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

type LambdaExpression struct {
	ParameterList *ParameterList
	Expression    Expression
	Tokens        Tokens
}

func (l *LambdaExpression) parse(p *pyParser) error {
	p.Skip()

	p.AcceptRunWhitespace()

	if !p.AcceptToken(parser.Token{Type: TokenDelimiter, Data: ":"}) {
		p.AcceptRunWhitespace()

		q := p.NewGoal()
		l.ParameterList = new(ParameterList)

		if err := l.ParameterList.parse(q); err != nil {
			return p.Error("LambdaExpression", err)
		}

		p.Score(q)
		p.AcceptRunWhitespace()

		if !p.AcceptToken(parser.Token{Type: TokenDelimiter, Data: ":"}) {
			return p.Error("LambdaExpression", ErrMissingColon)
		}
	}

	p.AcceptRunWhitespace()

	q := p.NewGoal()

	if err := l.Expression.parse(q); err != nil {
		return p.Error("LambdaExpression", err)
	}

	p.Score(q)

	l.Tokens = p.ToTokens()

	return nil
}

type OrTest struct {
	AndTest AndTest
	OrTest  *OrTest
	Tokens  Tokens
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
		q.AcceptRunWhitespace()
		p.Score(q)

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

type AndTest struct {
	NotTest NotTest
	AndTest *AndTest
	Tokens  Tokens
}

func (a *AndTest) parse(p *pyParser) error {
	q := p.NewGoal()

	if err := a.NotTest.parse(q); err != nil {
		return p.Error("AndTest", err)
	}

	p.Score(q)

	q = p.NewGoal()

	q.AcceptRunWhitespace()

	if q.AcceptToken(parser.Token{Type: TokenKeyword, Data: "or"}) {
		q.AcceptRunWhitespace()
		p.Score(q)

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

type NotTest struct {
	Nots       uint
	Comparison Comparison
	Tokens     Tokens
}

func (n *NotTest) parse(p *pyParser) error {
	for {
		if !p.AcceptToken(parser.Token{Type: TokenKeyword, Data: "not"}) {
			break
		}

		p.AcceptRunWhitespace()

		n.Nots++
	}

	q := p.NewGoal()

	if err := n.Comparison.parse(q); err != nil {
		return p.Error("NotTest", err)
	}

	p.Score(q)

	n.Tokens = p.ToTokens()

	return nil
}

type Comparison struct {
	OrExpression
	Comparisons []ComparisonExpression
	Tokens      Tokens
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

		q.AcceptRunWhitespace()

		var ce ComparisonExpression

		switch q.Peek() {
		case parser.Token{Type: TokenOperator, Data: "<"},
			parser.Token{Type: TokenOperator, Data: ">"},
			parser.Token{Type: TokenOperator, Data: "=="},
			parser.Token{Type: TokenOperator, Data: ">="},
			parser.Token{Type: TokenOperator, Data: "<="},
			parser.Token{Type: TokenOperator, Data: "!="}:
			p.Score(q)

			q = p.NewGoal()

			q.Skip()

			ce.ComparisonOperator = q.ToTokens()

			p.Score(q)
		case parser.Token{Type: TokenKeyword, Data: "is"}:
			p.Score(q)

			q = p.NewGoal()

			q.Skip()

			r := q.NewGoal()

			r.AcceptRunWhitespace()

			if r.AcceptToken(parser.Token{Type: TokenKeyword, Data: "not"}) {
				q.Score(r)
			}

			ce.ComparisonOperator = q.ToTokens()

			p.Score(q)
		case parser.Token{Type: TokenKeyword, Data: "not"}:
			p.Score(q)

			q = p.NewGoal()

			q.Skip()
			q.AcceptRunWhitespace()

			if !q.AcceptToken(parser.Token{Type: TokenKeyword, Data: "in"}) {
				return q.Error("Comparison", ErrMissingIn)
			}

			ce.ComparisonOperator = q.ToTokens()

			p.Score(q)
		case parser.Token{Type: TokenKeyword, Data: "in"}:
			p.Score(q)

			q = p.NewGoal()

			q.Skip()

			ce.ComparisonOperator = q.ToTokens()

			p.Score(q)
		default:
			break Loop
		}

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

type ComparisonExpression struct {
	ComparisonOperator []Token
	OrExpression       OrExpression
}

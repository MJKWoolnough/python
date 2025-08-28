package python

import (
	"slices"

	"vimagination.zapto.org/parser"
)

// Statement as defined in python@3.13.0:
// https://docs.python.org/release/3.13.0/reference/compound_stmts.html#grammar-token-python-grammar-statement
type Statement struct {
	StatementList     *StatementList
	CompoundStatement *CompoundStatement
	Comments          Comments
	Tokens            Tokens
}

func (s *Statement) parse(p *pyParser) error {
	var isCompound, isSoftCompound bool

	s.Comments = p.AcceptRunWhitespaceComments()

	p.AcceptRunAllWhitespace()

	q := p.NewGoal()

	switch tk := p.Peek(); tk.Type {
	case TokenOperator:
		isCompound = tk.Data == "@"
	case TokenKeyword:
		isCompound = slices.Contains(compounds[:], tk.Data)
	case TokenIdentifier:
		isCompound = tk.Data == "match"
		isSoftCompound = isCompound
	}

	if isCompound {
		c := new(CompoundStatement)

		if err := c.parse(q); err != nil {
			if !isSoftCompound {
				return p.Error("Statement", err)
			}
		} else {
			p.Score(q)

			s.CompoundStatement = c
		}
	}

	if s.CompoundStatement == nil {
		q = p.NewGoal()
		s.StatementList = new(StatementList)

		if err := s.StatementList.parse(q); err != nil {
			return p.Error("Statement", err)
		}

		p.Score(q)
	}

	s.Tokens = p.ToTokens()

	return nil
}

// StatementList as defined in python@3.13.0:
// https://docs.python.org/release/3.13.0/reference/compound_stmts.html#grammar-token-python-grammar-stmt_list
type StatementList struct {
	Statements []SimpleStatement
	Comments   Comments
	Tokens
}

func (s *StatementList) parse(p *pyParser) error {
	for {
		q := p.NewGoal()

		var ss SimpleStatement

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

		q = p.NewGoal()

		if tk := q.AcceptRunWhitespace(); tk == TokenComment || tk == TokenLineTerminator || tk == TokenDedent || tk == parser.TokenDone {
			break
		}

		p.Score(q)
	}

	s.Comments = p.AcceptRunWhitespaceCommentsNoNewline()
	s.Tokens = p.ToTokens()

	return nil
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

// SimpleStatement as defined in python@3.13.0:
// https://docs.python.org/release/3.13.0/reference/simple_stmts.html#grammar-token-python-grammar-simple_stmt
type SimpleStatement struct {
	Type                         StatementType
	AssertStatement              *AssertStatement
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
		p.Next()

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
		p.Next()

		s.Type = StatementBreak
	case parser.Token{Type: TokenKeyword, Data: "continue"}:
		p.Next()

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

		if n := q.LookaheadLine(
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

			if err := s.AnnotatedAssignmentStatement.parse(q); err != nil {
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

// AssertStatement as defined in python@3.13.0:
// https://docs.python.org/release/3.13.0/reference/simple_stmts.html#grammar-token-python-grammar-assert_stmt
type AssertStatement struct {
	Expressions []Expression
	Tokens      Tokens
}

func (a *AssertStatement) parse(p *pyParser) error {
	p.Next()

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

// AssignmentStatement as defined in python:3.13.0:
// https://docs.python.org/release/3.13.0/reference/simple_stmts.html#grammar-token-python-grammar-assignment_stmt
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

		if p.AcceptToken(parser.Token{Type: TokenDelimiter, Data: ","}) {
			p.AcceptRunWhitespace()
		}

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

// AugmentedAssignmentStatement as defined in python@3.13.0:
// https://docs.python.org/release/3.13.0/reference/simple_stmts.html#grammar-token-python-grammar-augmented_assignment_stmt
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

	q = p.NewGoal()

	if q.Accept(TokenDelimiter) {
		t := q.GetLastToken()

		switch t.Data {
		case "+=", "-=", "*=", "@=", "/=", "//=", "%=", "**=", ">>=", "<<=", "&=", "^=", "|=":
			a.AugOp = t
		}
	}

	if a.AugOp == nil {
		return p.Error("AugmentedAssignmentStatement", ErrMissingOp)
	}

	p.Score(q)
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

// AugTarget as defined in python@3.13.0:
// https://docs.python.org/release/3.13.0/reference/simple_stmts.html#grammar-token-python-grammar-augtarget
type AugTarget struct {
	PrimaryExpression PrimaryExpression
	Tokens            Tokens
}

func (a *AugTarget) parse(p *pyParser) error {
	q := p.NewGoal()

	if err := a.PrimaryExpression.parse(q); err != nil {
		return p.Error("AugTarget", err)
	} else if a.PrimaryExpression.Call != nil || !a.PrimaryExpression.IsIdentifier() {
		return p.Error("AugTarget", ErrMissingIdentifier)
	}

	p.Score(q)

	a.Tokens = p.ToTokens()

	return nil
}

// AnnotatedAssignmentStatement as defined in python@3.13.0:
// https://docs.python.org/release/3.13.0/reference/simple_stmts.html#grammar-token-python-grammar-annotated_assignment_stmt
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

// StarredExpression as defined in python@3.12.6:
// https://docs.python.org/release/3.12.6/reference/expressions.html#grammar-token-python-grammar-starred_expression
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

	q = p.NewGoal()

	q.AcceptRunWhitespace()

	if q.AcceptToken(parser.Token{Type: TokenDelimiter, Data: ","}) {
		s.Comments[1] = p.AcceptRunWhitespaceCommentsIfMultiline()
	} else {
		s.Comments[1] = p.AcceptRunWhitespaceCommentsNoNewlineIfMultiline()
	}

	s.Tokens = p.ToTokens()

	return nil
}

// DelStatement as defined in python@3.13.0:
// https://docs.python.org/release/3.13.0/reference/simple_stmts.html#grammar-token-python-grammar-del_stmt
type DelStatement struct {
	TargetList TargetList
	Tokens     Tokens
}

func (d *DelStatement) parse(p *pyParser) error {
	p.Next()

	p.AcceptRunWhitespace()

	q := p.NewGoal()

	if err := d.TargetList.parse(q); err != nil {
		return p.Error("DelStatement", err)
	}

	p.Score(q)

	d.Tokens = p.ToTokens()

	return nil
}

// ReturnStatement as defined in python@3.13.0:
// https://docs.python.org/release/3.13.0/reference/simple_stmts.html#grammar-token-python-grammar-return_stmt
type ReturnStatement struct {
	Expression *Expression
	Tokens     Tokens
}

func (r *ReturnStatement) parse(p *pyParser) error {
	p.Next()

	q := p.NewGoal()

	q.AcceptRunWhitespace()

	if tk := q.Peek(); tk != (parser.Token{Type: TokenDelimiter, Data: "}"}) && tk != (parser.Token{Type: TokenDelimiter, Data: "]"}) && tk != (parser.Token{Type: TokenDelimiter, Data: ")"}) && tk.Type != TokenLineTerminator && tk.Type != parser.TokenDone && tk.Type != TokenDedent {
		p.Score(q)

		q = p.NewGoal()
		r.Expression = new(Expression)

		if err := r.Expression.parse(q); err != nil {
			return q.Error("ReturnStatement", err)
		}

		p.Score(q)
	}

	r.Tokens = p.ToTokens()

	return nil
}

// YieldExpression as defined in python@3.13.0:
// https://docs.python.org/release/3.13.0/reference/simple_stmts.html#grammar-token-python-grammar-yield_stmt
type YieldExpression struct {
	ExpressionList *ExpressionList
	From           *Expression
	Comments       [4]Comments
	Tokens         Tokens
}

func (y *YieldExpression) parse(p *pyParser) error {
	y.Comments[0] = p.AcceptRunWhitespaceCommentsIfMultiline()

	p.AcceptRunWhitespace()
	p.Next()

	y.Comments[1] = p.AcceptRunWhitespaceCommentsIfMultiline()

	p.AcceptRunWhitespace()

	if p.AcceptToken(parser.Token{Type: TokenKeyword, Data: "from"}) {
		y.Comments[2] = p.AcceptRunWhitespaceCommentsIfMultiline()

		p.AcceptRunWhitespace()

		q := p.NewGoal()
		y.From = new(Expression)

		if err := y.From.parse(q); err != nil {
			return p.Error("YieldExpression", err)
		}

		p.Score(q)
	} else {
		q := p.NewGoal()
		y.ExpressionList = new(ExpressionList)

		if err := y.ExpressionList.parse(q); err != nil {
			return p.Error("YieldExpression", err)
		}

		p.Score(q)

		q = p.NewGoal()

		q.AcceptRunWhitespace()

		if q.AcceptToken(parser.Token{Type: TokenDelimiter, Data: ","}) {
			y.Comments[2] = p.AcceptRunWhitespaceCommentsIfMultiline()

			p.AcceptRunWhitespace()
			p.Next()
		}
	}

	y.Comments[3] = p.AcceptRunWhitespaceCommentsNoNewlineIfMultiline()
	y.Tokens = p.ToTokens()

	return nil
}

// RaiseStatement as defined in python@3.13.0:
// https://docs.python.org/release/3.13.0/reference/simple_stmts.html#grammar-token-python-grammar-raise_stmt
type RaiseStatement struct {
	Expression *Expression
	From       *Expression
	Tokens     Tokens
}

func (r *RaiseStatement) parse(p *pyParser) error {
	p.Next()
	p.AcceptRunWhitespace()

	q := p.NewGoal()

	switch q.AcceptRunWhitespace() {
	case TokenLineTerminator, TokenComment, TokenDedent, parser.TokenDone:
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
			r.From = new(Expression)

			if err := r.From.parse(q); err != nil {
				return p.Error("RaiseStatement", err)
			}

			p.Score(q)
		}
	}

	r.Tokens = p.ToTokens()

	return nil
}

// ImportStatement as defined in python@3.13.0:
// https://docs.python.org/release/3.13.0/reference/simple_stmts.html#grammar-token-python-grammar-import_stmt
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
		p.Next()
	}

	p.AcceptRunWhitespace()

	if i.RelativeModule == nil || !p.AcceptToken(parser.Token{Type: TokenOperator, Data: "*"}) {
		parens := i.RelativeModule != nil && p.AcceptToken(parser.Token{Type: TokenDelimiter, Data: "("})

		if parens {
			p.AcceptRunWhitespace()
		}

		for {
			q := p.NewGoal()

			var module ModuleAs

			if err := module.parse(q); err != nil {
				return p.Error("ImportStatement", err)
			}

			p.Score(q)

			i.Modules = append(i.Modules, module)

			p.AcceptRunWhitespace()

			if parens {
				if p.AcceptToken(parser.Token{Type: TokenDelimiter, Data: ")"}) {
					break
				}

				if !p.AcceptToken(parser.Token{Type: TokenDelimiter, Data: ","}) {
					return p.Error("ImportStatement", ErrMissingComma)
				}
			} else if !p.AcceptToken(parser.Token{Type: TokenDelimiter, Data: ","}) {
				break
			}

			p.AcceptRunWhitespace()
		}
	}

	i.Tokens = p.ToTokens()

	return nil
}

// RelativeModule as defined in python@3.13.0:
// https://docs.python.org/release/3.13.0/reference/simple_stmts.html#grammar-token-python-grammar-relative_module
type RelativeModule struct {
	Dots   int
	Module *Module
	Tokens Tokens
}

func (r *RelativeModule) parse(p *pyParser) error {
	q := p.NewGoal()

	for q.AcceptToken(parser.Token{Type: TokenDelimiter, Data: "."}) {
		r.Dots++

		p.Score(q)

		q = p.NewGoal()

		q.AcceptRunWhitespace()
	}
	switch q.Peek().Type {
	case TokenIdentifier:
		p.Score(q)

		q = p.NewGoal()
		r.Module = new(Module)

		if err := r.Module.parse(q); err != nil {
			return p.Error("RelativeModule", err)
		}

		p.Score(q)
	default:
		if r.Dots == 0 {
			return q.Error("RelativeModule", ErrMissingModule)
		}
	}

	r.Tokens = p.ToTokens()

	return nil
}

// ModuleAs as defined in python@3.13.0:
// https://docs.python.org/release/3.13.0/reference/simple_stmts.html#grammar-token-python-grammar-import_stmt
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

// Module as defined in python@3.13.0:
// https://docs.python.org/release/3.13.0/reference/simple_stmts.html#grammar-token-python-grammar-module
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

// GlobalStatement as defined in python@3.13.0:
// https://docs.python.org/release/3.13.0/reference/simple_stmts.html#grammar-token-python-grammar-global_stmt
type GlobalStatement struct {
	Identifiers []*Token
	Tokens      Tokens
}

func (g *GlobalStatement) parse(p *pyParser) error {
	p.Next()
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

// NonLocalStatement as defined in python@3.13.0:
// https://docs.python.org/release/3.13.0/reference/simple_stmts.html#grammar-token-python-grammar-nonlocal_stmt
type NonLocalStatement struct {
	Identifiers []*Token
	Tokens      Tokens
}

func (n *NonLocalStatement) parse(p *pyParser) error {
	p.Next()
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

// TypeStatement as defined in python@3.13.0:
// https://docs.python.org/release/3.13.0/reference/simple_stmts.html#grammar-token-python-grammar-type_stmt
type TypeStatement struct {
	Identifier *Token
	TypeParams *TypeParams
	Expression Expression
	Tokens     Tokens
}

func (t *TypeStatement) parse(p *pyParser) error {
	p.Next()
	p.AcceptRunWhitespace()

	if !p.Accept(TokenIdentifier) {
		return p.Error("TypeStatement", ErrMissingIdentifier)
	}

	t.Identifier = p.GetLastToken()

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

// TypeParams as defined in python@3.13.0:
// https://docs.python.org/release/3.13.0/reference/compound_stmts.html#grammar-token-python-grammar-type_params
type TypeParams struct {
	TypeParams []TypeParam
	Comments   [2]Comments
	Tokens     Tokens
}

func (t *TypeParams) parse(p *pyParser) error {
	p.AcceptToken(parser.Token{Type: TokenDelimiter, Data: "["})

	t.Comments[0] = p.AcceptRunWhitespaceCommentsNoNewline()
	q := p.NewGoal()

	q.AcceptRunAllWhitespace()

	if !q.AcceptToken(parser.Token{Type: TokenDelimiter, Data: "]"}) {
		for {
			p.AcceptRunWhitespaceNoComment()

			q := p.NewGoal()

			var tp TypeParam

			if err := tp.parse(q); err != nil {
				return p.Error("TypeParams", err)
			}

			t.TypeParams = append(t.TypeParams, tp)

			p.Score(q)

			q = p.NewGoal()

			q.AcceptRunAllWhitespace()

			if q.AcceptToken(parser.Token{Type: TokenDelimiter, Data: "]"}) {
				break
			} else if !q.AcceptToken(parser.Token{Type: TokenDelimiter, Data: ","}) {
				return q.Error("TypeParams", ErrMissingComma)
			}

			p.AcceptRunWhitespace()
			p.AcceptToken(parser.Token{Type: TokenDelimiter, Data: ","})
		}
	}

	t.Comments[1] = p.AcceptRunWhitespaceComments()

	p.AcceptRunAllWhitespaceNoComment()

	p.AcceptToken(parser.Token{Type: TokenDelimiter, Data: "]"})

	t.Tokens = p.ToTokens()

	return nil
}

// AssignmentExpression as defined in python@3.13.0:
// https://docs.python.org/release/3.13.0/reference/expressions.html#grammar-token-python-grammar-assignment_expression
type AssignmentExpression struct {
	Identifier *Token
	Expression Expression
	Comments   [3]Comments
	Tokens
}

func (a *AssignmentExpression) parse(p *pyParser) error {
	a.Comments[0] = p.AcceptRunWhitespaceCommentsIfMultiline()

	p.AcceptRunWhitespace()

	q := p.NewGoal()

	if q.Accept(TokenIdentifier) {
		q.AcceptRunWhitespace()

		if q.AcceptToken(parser.Token{Type: TokenOperator, Data: ":="}) {
			p.Next()
			a.Identifier = p.GetLastToken()

			a.Comments[1] = p.AcceptRunWhitespaceCommentsIfMultiline()

			p.AcceptRunWhitespace()
			p.Next()

			a.Comments[2] = p.AcceptRunWhitespaceCommentsIfMultiline()

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

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

type SimpleStatement struct{}

func (s *SimpleStatement) parse(_ *pyParser) error {
	return nil
}

type AssignmentExpression struct{}

func (s *AssignmentExpression) parse(_ *pyParser) error {
	return nil
}

// Package python implements a python tokeniser and parser.
package python // import "vimagination.zapto.org/python"

import "vimagination.zapto.org/parser"

// File represents a parsed Python file.
type File struct {
	Statements []Statement
	Comments   Comments
	Tokens     Tokens
}

// Parse parses Python input into AST.
func Parse(t Tokeniser) (*File, error) {
	p, err := newPyParser(t)
	if err != nil {
		return nil, err
	}

	f := new(File)
	if err = f.parse(p); err != nil {
		return nil, err
	}

	return f, nil
}

func (f *File) parse(p *pyParser) error {
	q := p.NewGoal()

	for q.AcceptRunAllWhitespace() != parser.TokenDone {
		p.AcceptRunAllWhitespaceNoComment()

		var s Statement

		q = p.NewGoal()

		if err := s.parse(q); err != nil {
			return p.Error("File", err)
		}

		f.Statements = append(f.Statements, s)

		p.Score(q)

		q = p.NewGoal()
	}

	f.Comments = p.AcceptRunWhitespaceComments()
	f.Tokens = p.ToTokens()

	return nil
}

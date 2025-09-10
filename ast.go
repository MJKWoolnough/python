// Package python implements a python tokeniser and parser.
package python // import "vimagination.zapto.org/python"

import "vimagination.zapto.org/parser"

// File represents a parsed Python file.
//
// The first set of comments are parsed and printed at the top of the file, the
// second set are parsed and printed at the bottom of the file.
type File struct {
	Statements []Statement
	Comments   [2]Comments
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
	if p.Peek().Type == TokenComment {
		f.Comments[0] = p.AcceptRunWhitespaceCommentsNoNewline()
	}

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

	f.Comments[1] = p.AcceptRunWhitespaceComments()
	f.Tokens = p.ToTokens()

	return nil
}

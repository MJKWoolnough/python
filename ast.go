package python

import "vimagination.zapto.org/parser"

type File struct {
	Statements []Statement
	Tokens     Tokens
}

func Parse(t Tokeniser) (*File, error) {
	p, err := newPyParser(t)
	if err != nil {
		return nil, err
	}

	f := new(File)
	if err = f.parse(&p); err != nil {
		return nil, err
	}

	return f, nil
}

func (f *File) parse(p *pyParser) error {
	for p.AcceptRun(TokenLineTerminator) != parser.TokenDone {
		var s Statement

		q := p.NewGoal()
		if err := s.parse(&q); err != nil {
			return p.Error("File", err)
		}

		p.Score(q)
	}

	return nil
}

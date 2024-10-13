package python

import (
	"fmt"

	"vimagination.zapto.org/parser"
)

type Token struct {
	parser.Token
	Pos, Line, LinePos uint64
}

type Tokens []Token

type pyParser struct {
	inBrackets uint
	Tokens
}

type Tokeniser interface {
	GetToken() (parser.Token, error)
	GetError() error
	TokeniserState(parser.TokenFunc)
}

func newPyParser(t Tokeniser) (*pyParser, error) {
	p := &pyTokeniser{
		indents: []string{""},
	}

	t.TokeniserState(p.main)

	var (
		tokens             Tokens
		pos, line, linePos uint64
	)

	for {
		tk, _ := t.GetToken()
		tokens = append(tokens, Token{Token: tk, Pos: pos, Line: line, LinePos: linePos})

		switch tk.Type {
		case parser.TokenDone:
			return &pyParser{Tokens: tokens[0:0:len(tokens)]}, nil
		case parser.TokenError:
			return nil, Error{Err: t.GetError(), Parsing: "Tokens", Token: tokens[len(tokens)-1]}
		case TokenLineTerminator:
			line += uint64(len(tk.Data))
			linePos = 0
		default:
			linePos += uint64(len(tk.Data))
		}

		pos += uint64(len(tk.Data))
	}
}

func (p pyParser) NewGoal() *pyParser {
	return &pyParser{
		inBrackets: p.inBrackets,
		Tokens:     p.Tokens[len(p.Tokens):],
	}
}

func (p *pyParser) Score(k *pyParser) {
	p.Tokens = p.Tokens[:len(p.Tokens)+len(k.Tokens)]
}

func (p *pyParser) next() Token {
	l := len(p.Tokens)
	if l == cap(p.Tokens) {
		return p.Tokens[l-1]
	}

	p.Tokens = p.Tokens[:l+1]
	tk := p.Tokens[l]

	return tk
}

func (p *pyParser) backup() {
	p.Tokens = p.Tokens[:len(p.Tokens)-1]
}

func (j *pyParser) Peek() parser.Token {
	tk := j.next().Token

	j.backup()

	return tk
}

func (p *pyParser) Accept(ts ...parser.TokenType) bool {
	tt := p.next().Type

	for _, pt := range ts {
		if pt == tt {
			return true
		}
	}

	p.backup()

	return false
}

func (p *pyParser) AcceptRun(ts ...parser.TokenType) parser.TokenType {
Loop:
	for {
		tt := p.next().Type

		for _, pt := range ts {
			if pt == tt {
				continue Loop
			}
		}

		p.backup()

		return tt
	}
}

func (p *pyParser) Skip() {
	p.next()
}

func (p *pyParser) ExceptRun(ts ...parser.TokenType) parser.TokenType {
	for {
		tt := p.next().Type

		for _, pt := range ts {
			if pt == tt || tt < 0 {
				p.backup()

				return tt
			}
		}
	}
}

func (p *pyParser) AcceptToken(tk parser.Token) bool {
	if p.next().Token == tk {
		return true
	}

	p.backup()

	return false
}

func (p *pyParser) ToTokens() Tokens {
	return p.Tokens[:len(p.Tokens):len(p.Tokens)]
}

func (p *pyParser) GetLastToken() *Token {
	return &p.Tokens[len(p.Tokens)-1]
}

func (p *pyParser) OpenBrackets() {
	p.inBrackets++
}

func (p *pyParser) CloseBrackets() {
	p.inBrackets--
}

func (p *pyParser) AcceptRunWhitespace() parser.TokenType {
	if p.inBrackets > 0 {
		return p.AcceptRun(TokenWhitespace, TokenComment)
	}

	return p.AcceptRun(TokenWhitespace)
}

func (p *pyParser) LookaheadLine(tks ...parser.Token) int {
	brackets := 0

Loop:
	for _, tk := range p.Tokens[:cap(p.Tokens)] {
		if brackets > 0 {
			switch tk.Data {
			case "]", ")", "}":
				brackets--
			}

			continue
		}

		if tk.Type == TokenDelimiter {
			switch tk.Data {
			case "[", "(", "{":
				brackets++

				continue
			case "]", ")", "}":
				break Loop
			}
		}

		if tk.Type == TokenLineTerminator {
			break
		}

		for n, t := range tks {
			if t == tk.Token {
				return n
			}
		}
	}

	return -1
}

type Error struct {
	Err     error
	Parsing string
	Token   Token
}

func (e Error) Error() string {
	return fmt.Sprintf("%s: error at position %d (%d:%d):\n%s", e.Parsing, e.Token.Pos+1, e.Token.Line+1, e.Token.LinePos+1, e.Err)
}

func (e Error) Unwrap() error {
	return e.Err
}

func (p *pyParser) Error(parsingFunc string, err error) error {
	tk := p.next()

	p.backup()

	return Error{
		Err:     err,
		Parsing: parsingFunc,
		Token:   tk,
	}
}

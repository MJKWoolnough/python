package python

import (
	"fmt"
	"strings"

	"vimagination.zapto.org/parser"
)

// Token represents a parser.Token combined with positioning information.
type Token struct {
	parser.Token
	Pos, Line, LinePos uint64
}

// Tokens represents a list of tokens that have been parsed.
type Tokens []Token

type Comments []Token

type pyParser struct {
	inBrackets uint
	Tokens
}

// Tokeniser represents the methods required by the python tokeniser.
type Tokeniser interface {
	Iter(func(parser.Token) bool)
	TokeniserState(parser.TokenFunc)
	GetError() error
}

func newPyParser(t Tokeniser) (*pyParser, error) {
	p := &pyTokeniser{
		indents: []string{""},
	}

	t.TokeniserState(p.main)

	var (
		tokens             Tokens
		err                error
		pos, line, linePos uint64
	)

	for tk := range t.Iter {
		tokens = append(tokens, Token{Token: tk, Pos: pos, Line: line, LinePos: linePos})

		switch tk.Type {
		case parser.TokenDone:
		case parser.TokenError:
			err = Error{Err: t.GetError(), Parsing: "Tokens", Token: tokens[len(tokens)-1]}
		case TokenLineTerminator:
			line += uint64(len(tk.Data))
			linePos = 0
		case TokenWhitespace:
			for _, c := range tk.Data {
				if c == '\n' {
					line++
					linePos = 0
				} else {
					linePos++
				}
			}
		default:
			linePos += uint64(len(tk.Data))
		}

		pos += uint64(len(tk.Data))
	}

	return &pyParser{Tokens: tokens[0:0:len(tokens)]}, err
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
	p.Tokens = p.Tokens[:l+1]
	tk := p.Tokens[l]

	return tk
}

func (p *pyParser) backup() {
	p.Tokens = p.Tokens[:len(p.Tokens)-1]
}

func (p *pyParser) Peek() parser.Token {
	tk := p.next().Token

	p.backup()

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

func (p *pyParser) Next() Token {
	return p.next()
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
		return p.AcceptRunAllWhitespace()
	}

	return p.AcceptRun(TokenWhitespace, TokenComment)
}

func (p *pyParser) AcceptRunWhitespaceNoComment() parser.TokenType {
	if p.inBrackets > 0 {
		return p.AcceptRunAllWhitespaceNoComment()
	}

	return p.AcceptRun(TokenWhitespace)
}

func (p *pyParser) AcceptRunWhitespaceComments() Comments {
	var c Comments

	s := p.NewGoal()

	for s.AcceptRunAllWhitespaceNoComment() == TokenComment {
		c = append(c, s.Next())

		p.Score(s)

		s = p.NewGoal()
	}

	return c
}

func (p *pyParser) AcceptRunWhitespaceCommentsNoNewline() Comments {
	var c Comments

	s := p.NewGoal()

	for s.AcceptRunWhitespaceNoNewline() == TokenComment {
		p.Score(s)

		c = append(c, p.Next())
		s = p.NewGoal()

		s.AcceptNewline()
	}

	return c
}

func (p *pyParser) AcceptRunWhitespaceNoNewline() parser.TokenType {
	for {
		if tk := p.Peek(); tk.Type != TokenWhitespace || strings.Contains(tk.Data, lineTerminator) {
			return tk.Type
		}

		p.Next()
	}
}

func (p *pyParser) AcceptNewline() bool {
	if tk := p.Peek(); tk.Type == TokenWhitespace && strings.Count(tk.Data, "\n") == 1 {
		p.Next()

		return true
	}

	return p.Accept(TokenLineTerminator)
}

func (p *pyParser) AcceptRunAllWhitespace() parser.TokenType {
	return p.AcceptRun(TokenWhitespace, TokenComment, TokenLineTerminator)
}

func (p *pyParser) AcceptRunAllWhitespaceNoComment() parser.TokenType {
	return p.AcceptRun(TokenWhitespace, TokenLineTerminator)
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

// Error represents a Python parsing error.
type Error struct {
	Err     error
	Parsing string
	Token   Token
}

// Error implements the error interface.
func (e Error) Error() string {
	return fmt.Sprintf("%s: error at position %d (%d:%d):\n%s", e.Parsing, e.Token.Pos+1, e.Token.Line+1, e.Token.LinePos+1, e.Err)
}

// Unwrap returns the underlying error.
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

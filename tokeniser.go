package python

import (
	"io"
	"strings"
	"unicode"

	"vimagination.zapto.org/parser"
)

const (
	whitespace       = " \t"
	lineTerminator   = "\n"
	comment          = "#"
	singleEscapeChar = "'\"\\bfnrtv"
	binaryDigit      = "01"
	octalDigit       = "01234567"
	decimalDigit     = "0123456789"
	hexDigit         = "0123456789abcdefABCDEF"
	stringPrefix     = "rRuUfFbB"
	stringStart      = "\"'"
)

var (
	keywords = [...]string{"await", "else", "import", "pass", "None", "break", "except", "in", "raise", "class", "finally", "is", "return", "and", "continue", "for", "lambda", "try", "as", "def", "from", "nonlocal", "while", "assert", "del", "global", "not", "with", "async", "elif", "if", "or", "yield"}

	idContinue = []*unicode.RangeTable{
		unicode.L,
		unicode.Nl,
		unicode.Other_ID_Start,
		unicode.Mn,
		unicode.Mc,
		unicode.Nd,
		unicode.Pc,
		unicode.Other_ID_Continue,
	}
	idStart = []*unicode.RangeTable{
		unicode.Lu,
		unicode.Ll,
		unicode.Lt,
		unicode.Lm,
		unicode.Lo,
		unicode.Nl,
		unicode.Other_ID_Start,
	}
)

const (
	TokenWhitespace parser.TokenType = iota
	TokenLineTerminator
	TokenComment
	TokenIdentifier
	TokenKeyword
	TokenOperator
	TokenDelimiter
	TokenBooleanLiteral
	TokenNumericLiteral
	TokenStringLiteral
	TokenNullLiteral
)

func isIDStart(c rune) bool {
	return c == '_' || unicode.In(c, idStart...)
}

func isIDContinue(c rune) bool {
	return c == '_' || unicode.In(c, idContinue...)
}

type pyTokeniser struct {
	tokenDepth []byte
}

func (p *pyTokeniser) main(t *parser.Tokeniser) (parser.Token, parser.TokenFunc) {
	if t.Peek() == 0 {
		if len(p.tokenDepth) > 0 {
			t.Err = io.ErrUnexpectedEOF

			return t.Error()
		}

		return t.Done()
	}

	if t.Accept(whitespace) {
		t.AcceptRun(whitespace)

		return parser.Token{
			Type: TokenWhitespace,
			Data: t.Get(),
		}, p.main
	}

	if t.Accept(lineTerminator) {
		t.AcceptRun(lineTerminator)

		return parser.Token{
			Type: TokenLineTerminator,
			Data: t.Get(),
		}, p.main
	}

	if t.Accept(comment) {
		t.ExceptRun(lineTerminator)

		return parser.Token{
			Type: TokenComment,
			Data: t.Get(),
		}, p.main
	}

	pk := t.Peek()

	if strings.ContainsRune(stringPrefix, pk) {
		return p.stringOrIdentifier(t)
	}

	if strings.ContainsRune(stringStart, pk) {
		return p.string(t)
	}

	if isIDStart(pk) {
		return p.identifier(t)
	}

	if t.Accept("0") {
		return p.baseNumber(t)
	}

	if t.Accept(decimalDigit) {
		return p.number(t)
	}

	if t.Accept(".") {
		return p.floatOrDelimiter(t)
	}

	return p.operatorOrDelimiter(t)
}

func (p *pyTokeniser) stringOrIdentifier(t *parser.Tokeniser) (parser.Token, parser.TokenFunc) {
	return parser.Token{}, nil
}

func (p *pyTokeniser) string(t *parser.Tokeniser) (parser.Token, parser.TokenFunc) {
	return parser.Token{}, nil
}

func (p *pyTokeniser) identifier(t *parser.Tokeniser) (parser.Token, parser.TokenFunc) {
	return parser.Token{}, nil
}

func (p *pyTokeniser) baseNumber(t *parser.Tokeniser) (parser.Token, parser.TokenFunc) {
	return parser.Token{}, nil
}

func (p *pyTokeniser) number(t *parser.Tokeniser) (parser.Token, parser.TokenFunc) {
	return parser.Token{}, nil
}

func (p *pyTokeniser) floatOrDelimiter(t *parser.Tokeniser) (parser.Token, parser.TokenFunc) {
	return parser.Token{}, nil
}

func (p *pyTokeniser) operatorOrDelimiter(t *parser.Tokeniser) (parser.Token, parser.TokenFunc) {
	return parser.Token{}, nil
}

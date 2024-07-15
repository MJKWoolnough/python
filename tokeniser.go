package python

import (
	"errors"
	"io"
	"slices"
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
	keywords = [...]string{"await", "else", "import", "pass", "break", "except", "in", "raise", "class", "finally", "is", "return", "and", "continue", "for", "lambda", "try", "as", "def", "from", "nonlocal", "while", "assert", "del", "global", "not", "with", "async", "elif", "if", "or", "yield"}
	booleans = [...]string{"True", "False"}
	null     = "None"

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

func SetTokeniser(t *parser.Tokeniser) *parser.Tokeniser {
	t.TokeniserState(new(pyTokeniser).main)

	return t
}

func (p *pyTokeniser) main(t *parser.Tokeniser) (parser.Token, parser.TokenFunc) {
	if t.Peek() == -1 {
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
		return p.string(t, false)
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
	var raw bool

	switch t.Peek() {
	case 'r', 'R':
		t.Except("")
		t.Accept("fFbB")

		raw = true
	case 'b', 'B', 'f', 'F':
		t.Except("")

		raw = t.Accept("rR")
	case 'u', 'U':
		t.Except("")
	}

	if strings.ContainsRune(stringStart, t.Peek()) {
		return p.string(t, raw)
	}

	return p.identifier(t)
}

func (p *pyTokeniser) string(t *parser.Tokeniser, raw bool) (parser.Token, parser.TokenFunc) {
	var m string

	if t.Peek() == '"' {
		m = "\""
	} else {
		m = "'"
	}

	triple := false

	t.Except("")

	if t.Accept(m) {
		if !t.Accept(m) {
			return parser.Token{
				Type: TokenStringLiteral,
				Data: t.Get(),
			}, p.main
		}

		triple = true
	}

	except := "\n" + m

	if !raw {
		except += "\\"
	}

Loop:
	for {
		switch t.ExceptRun(except) {
		default:
			t.Err = io.ErrUnexpectedEOF

			return t.Error()
		case '\\':
			t.Except("")
			t.Except("")
		case '\n':
			if !triple {
				t.Err = ErrInvalidCharacter

				return t.Error()
			}

			t.Except("")
		case '\'', '"':
			t.Except("")

			if !triple || t.Accept(m) && t.Accept(m) {
				break Loop
			}
		}
	}

	return parser.Token{
		Type: TokenStringLiteral,
		Data: t.Get(),
	}, p.main
}

func (p *pyTokeniser) identifier(t *parser.Tokeniser) (parser.Token, parser.TokenFunc) {
	for isIDContinue(t.Peek()) {
		t.Except("")
	}

	ident := t.Get()
	typ := TokenIdentifier

	if slices.Contains(booleans[:], ident) {
		typ = TokenBooleanLiteral
	}

	if slices.Contains(keywords[:], ident) {
		typ = TokenKeyword
	}

	if ident == "None" {
		typ = TokenNullLiteral
	}

	return parser.Token{
		Type: typ,
		Data: ident,
	}, p.main
}

func numberWithGrouping(t *parser.Tokeniser, digits string) bool {
	for t.Accept("_") {
		if !t.Accept(digits) {
			return false
		}

		t.AcceptRun(digits)
	}

	return true
}

func (p *pyTokeniser) baseNumber(t *parser.Tokeniser) (parser.Token, parser.TokenFunc) {
	digits := "0"

	if t.Accept("xX") {
		digits = hexDigit
	} else if t.Accept("oO") {
		digits = octalDigit
	} else if t.Accept("bB") {
		digits = binaryDigit
	}

	if !t.Accept(digits) && digits != "0" {
		t.Err = ErrInvalidNumber

		return t.Error()
	}
	t.AcceptRun(digits)

	if !numberWithGrouping(t, digits) {
		t.Err = ErrInvalidNumber

		return t.Error()
	}

	if digits == "0" {
		return p.floatOrImaginary(t)
	}

	return parser.Token{
		Type: TokenNumericLiteral,
		Data: t.Get(),
	}, p.main
}

func (p *pyTokeniser) floatOrImaginary(t *parser.Tokeniser) (parser.Token, parser.TokenFunc) {
	if t.Accept(".") {
		return p.float(t)
	}

	if t.Accept("eE") {
		return p.exponential(t)
	}

	return p.imaginary(t)
}

func (p *pyTokeniser) float(t *parser.Tokeniser) (parser.Token, parser.TokenFunc) {
	if !t.Accept(decimalDigit) || !numberWithGrouping(t, decimalDigit) {
		t.Err = ErrInvalidNumber

		return t.Error()
	}

	if t.Accept("eE") {
		return p.exponential(t)
	}

	return p.imaginary(t)
}

func (p *pyTokeniser) exponential(t *parser.Tokeniser) (parser.Token, parser.TokenFunc) {
	t.Accept("+-")

	if !t.Accept(decimalDigit) || !numberWithGrouping(t, decimalDigit) {
		t.Err = ErrInvalidNumber

		return t.Error()
	}

	return parser.Token{
		Type: TokenNumericLiteral,
		Data: t.Get(),
	}, p.main
}

func (p *pyTokeniser) imaginary(t *parser.Tokeniser) (parser.Token, parser.TokenFunc) {
	t.Accept("jJ")

	return parser.Token{
		Type: TokenNumericLiteral,
		Data: t.Get(),
	}, p.main
}

func (p *pyTokeniser) number(t *parser.Tokeniser) (parser.Token, parser.TokenFunc) {
	t.AcceptRun(decimalDigit)

	if !t.Accept(decimalDigit) {
		t.Err = ErrInvalidNumber

		return t.Error()
	}

	return p.floatOrImaginary(t)
}

func (p *pyTokeniser) floatOrDelimiter(t *parser.Tokeniser) (parser.Token, parser.TokenFunc) {
	if t.Accept(decimalDigit) {
		t.AcceptRun(decimalDigit)

		return p.imaginary(t)
	}

	return parser.Token{
		Type: TokenDelimiter,
		Data: t.Get(),
	}, p.main
}

func (p *pyTokeniser) operatorOrDelimiter(t *parser.Tokeniser) (parser.Token, parser.TokenFunc) {
	typ := TokenOperator
	bracket := 0

	const brackets = "}])"

	switch c := t.Peek(); c {
	default:
		t.Err = ErrInvalidCharacter

		return t.Error()
	case '+', '%', '@', '|', '^', ':', '=':
		t.Except("")
		t.Accept("=")

		if c != '=' && c != ':' {
			typ = TokenDelimiter
		}
	case '-':
		t.Except("")

		if t.Accept("=>") {
			typ = TokenDelimiter
		}
	case '*', '/', '<', '>':
		t.Except("")

		d := t.Accept(string(c))

		if t.Accept("=") && (!d || c == '*' || c == '/') {
			typ = TokenDelimiter
		}
	case '!':
		if !t.Accept("=") {
			t.Err = ErrInvalidCharacter

			return t.Error()
		}
	case ')', '}', ']':
		if len(p.tokenDepth) == 0 || p.tokenDepth[len(p.tokenDepth)-1] != byte(c) {
			t.Err = ErrInvalidCharacter

			return t.Error()
		}

		t.Except("")
		p.tokenDepth = p.tokenDepth[:len(p.tokenDepth)-1]

		typ = TokenDelimiter
	case '(':
		bracket++

		fallthrough
	case '[':
		bracket++

		fallthrough
	case '{':
		p.tokenDepth = append(p.tokenDepth, brackets[bracket])

		fallthrough
	case ',', '.', ';':
		typ = TokenDelimiter

		fallthrough
	case '~':
		t.Except("")
	}

	return parser.Token{
		Type: typ,
		Data: t.Get(),
	}, p.main
}

var (
	ErrInvalidCharacter = errors.New("invalid character")
	ErrInvalidNumber    = errors.New("invalid number")
)

package python

import (
	"io"
	"slices"
	"strings"
	"unicode"

	"vimagination.zapto.org/parser"
)

const (
	whitespaceWithLineTerminator = " \t\n"
	whitespace                   = " \t"
	lineTerminator               = "\n"
	notIndent                    = "\\\n#"
	comment                      = "#"
	singleEscapeChar             = "'\"\\bfnrtv"
	binaryDigit                  = "01"
	octalDigit                   = "01234567"
	decimalDigit                 = "0123456789"
	hexDigit                     = "0123456789abcdefABCDEF"
	stringPrefix                 = "rRuUfFbB"
	stringStart                  = "\"'"

	singleQuotedExcept = "\\\n'"
	doubleQuotedExcept = "\\\n\""
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
	TokenIndent
	TokenDedent
)

func isIDStart(c rune) bool {
	return c == '_' || unicode.In(c, idStart...)
}

func isIDContinue(c rune) bool {
	return c == '_' || unicode.In(c, idContinue...)
}

type pyTokeniser struct {
	tokenDepth []byte
	indents    []string
	dedents    int
}

func SetTokeniser(t *parser.Tokeniser) *parser.Tokeniser {
	p := &pyTokeniser{
		indents: []string{""},
	}

	t.TokeniserState(p.main)

	return t
}

func (p *pyTokeniser) main(t *parser.Tokeniser) (parser.Token, parser.TokenFunc) {
	if t.Peek() == -1 {
		if len(p.tokenDepth) > 0 {
			t.Err = io.ErrUnexpectedEOF

			return t.Error()
		}

		if p.dedents = len(p.indents) - 1; p.dedents > 0 {
			p.indents = p.indents[:1]

			return p.dedent(t)
		}

		return t.Done()
	}

	ws := whitespace

	if len(p.tokenDepth) > 0 {
		ws = whitespaceWithLineTerminator
	}

	if t.Accept(ws) {
		t.AcceptRun(ws)

		return t.Return(TokenWhitespace, p.main)
	}

	if t.Accept(lineTerminator) {
		t.AcceptRun(lineTerminator)

		return t.Return(TokenLineTerminator, p.indent)
	}

	if t.Accept("\\") {
		if !t.Accept(lineTerminator) {
			t.Err = ErrInvalidCharacter

			return t.Error()
		}

		return t.Return(TokenWhitespace, p.main)
	}

	if t.Accept(comment) {
		t.ExceptRun(lineTerminator)

		return t.Return(TokenComment, p.main)
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

func (p *pyTokeniser) indent(t *parser.Tokeniser) (parser.Token, parser.TokenFunc) {
	t.AcceptRun(whitespace)

	indent := t.Get()
	current := p.indents[0]

	if current == indent || strings.ContainsRune(notIndent, t.Peek()) {
		if indent == "" {
			return p.main(t)
		}

		return parser.Token{
			Type: TokenWhitespace,
			Data: indent,
		}, p.main
	}

	if len(indent) > len(current) {
		if strings.HasPrefix(indent, current) {
			p.indents = slices.Insert(p.indents, 0, indent)

			return parser.Token{
				Type: TokenIndent,
				Data: indent,
			}, p.main
		}

		t.Err = ErrInvalidIndent

		return t.Error()
	}

	for n, i := range p.indents[1:] {
		if indent == i {
			p.indents = slices.Delete(p.indents, 0, n+1)
			p.dedents = n

			return parser.Token{
				Type: TokenDedent,
				Data: indent,
			}, p.dedent
		}
	}

	t.Err = ErrInvalidIndent

	return t.Error()
}

func (p *pyTokeniser) dedent(t *parser.Tokeniser) (parser.Token, parser.TokenFunc) {
	if p.dedents == 0 {
		return p.main(t)
	}

	p.dedents--

	return t.Return(TokenDedent, p.dedent)
}

func parseStringRaw(t *parser.Tokeniser) bool {
	switch t.Peek() {
	case 'r', 'R':
		t.Next()
		t.Accept("fFbB")

		return true
	case 'b', 'B', 'f', 'F':
		t.Next()

		return t.Accept("rR")
	case 'u', 'U':
		t.Next()
	}

	return false
}

func (p *pyTokeniser) stringOrIdentifier(t *parser.Tokeniser) (parser.Token, parser.TokenFunc) {
	raw := parseStringRaw(t)

	if strings.ContainsRune(stringStart, t.Peek()) {
		return p.string(t, raw)
	}

	return p.identifier(t)
}

type stringOpening byte

const (
	stringError stringOpening = iota
	stringSingle
	stringDouble
	stringTripleSingle
	stringTripleDouble
	stringEmpty
)

func (s stringOpening) Quote() string {
	if s == stringSingle || s == stringTripleSingle {
		return "'"
	} else if s == stringDouble || s == stringTripleDouble {
		return "\""
	}

	return ""
}

func (s stringOpening) Except(raw bool) string {
	var except string

	if s == stringSingle || s == stringTripleSingle {
		except = singleQuotedExcept
	} else {
		except = doubleQuotedExcept
	}

	if raw {
		except = except[1:]
	}

	return except
}

func (s stringOpening) IsTriple() bool {
	return s == stringTripleSingle || s == stringTripleDouble
}

func parseStringOpening(t *parser.Tokeniser) stringOpening {
	var (
		m       string
		opening stringOpening
	)

	if t.Accept("\"") {
		m = "\""
		opening = stringDouble
	} else if t.Accept("'") {
		m = "'"
		opening = stringSingle
	} else {
		return opening
	}

	if t.Accept(m) {
		if !t.Accept(m) {
			return stringEmpty
		}

		opening += 2
	}

	return opening
}

func (p *pyTokeniser) string(t *parser.Tokeniser, raw bool) (parser.Token, parser.TokenFunc) {
	so := parseStringOpening(t)

	if so == stringEmpty {
		return parser.Token{
			Type: TokenStringLiteral,
			Data: t.Get(),
		}, p.main
	}

	m := so.Quote()
	triple := so.IsTriple()
	except := so.Except(raw)

Loop:
	for {
		switch t.ExceptRun(except) {
		default:
			t.Err = io.ErrUnexpectedEOF

			return t.Error()
		case '\\':
			t.Next()
			t.Next()
		case '\n':
			if !triple {
				t.Err = ErrInvalidCharacter

				return t.Error()
			}

			t.Next()
		case '\'', '"':
			t.Next()

			if !triple || t.Accept(m) && t.Accept(m) {
				break Loop
			}
		}
	}

	return t.Return(TokenStringLiteral, p.main)
}

func (p *pyTokeniser) identifier(t *parser.Tokeniser) (parser.Token, parser.TokenFunc) {
	for isIDContinue(t.Peek()) {
		t.Next()
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

	t.AcceptRun(digits)

	if !numberWithGrouping(t, digits) {
		t.Err = ErrInvalidNumber

		return t.Error()
	}

	if digits == "0" {
		return p.floatOrImaginary(t)
	}

	if t.Len() == 2 {
		t.Err = ErrInvalidNumber

		return t.Error()
	}

	return t.Return(TokenNumericLiteral, p.main)
}

func (p *pyTokeniser) floatOrImaginary(t *parser.Tokeniser) (parser.Token, parser.TokenFunc) {
	if t.Accept(".") {
		if strings.ContainsRune(decimalDigit, t.Peek()) {
			return p.float(t)
		}
	} else if t.Accept("eE") {
		return p.exponential(t)
	}

	return p.imaginary(t)
}

func (p *pyTokeniser) float(t *parser.Tokeniser) (parser.Token, parser.TokenFunc) {
	if !t.Accept(decimalDigit) {
		t.Err = ErrInvalidNumber

		return t.Error()
	}

	t.AcceptRun(decimalDigit)

	if !numberWithGrouping(t, decimalDigit) {
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

	if !t.Accept(decimalDigit) {
		t.Err = ErrInvalidNumber

		return t.Error()
	}

	t.AcceptRun(decimalDigit)

	if !numberWithGrouping(t, decimalDigit) {
		t.Err = ErrInvalidNumber

		return t.Error()
	}

	return p.imaginary(t)
}

func (p *pyTokeniser) imaginary(t *parser.Tokeniser) (parser.Token, parser.TokenFunc) {
	t.Accept("jJ")

	return t.Return(TokenNumericLiteral, p.main)
}

func (p *pyTokeniser) number(t *parser.Tokeniser) (parser.Token, parser.TokenFunc) {
	t.AcceptRun(decimalDigit)

	if !numberWithGrouping(t, decimalDigit) {
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

	return t.Return(TokenDelimiter, p.main)
}

func (p *pyTokeniser) operatorOrDelimiter(t *parser.Tokeniser) (parser.Token, parser.TokenFunc) {
	typ := TokenOperator
	bracket := 0

	const brackets = "}])"

	switch c := t.Peek(); c {
	default:
		t.Err = ErrInvalidCharacter

		return t.Error()
	case '+', '%', '@', '&', '|', '^', ':', '=':
		t.Next()

		if t.Accept("=") != (c == '=' || c == ':') {
			typ = TokenDelimiter
		}
	case '-':
		t.Next()

		if t.Accept("=>") {
			typ = TokenDelimiter
		}
	case '*', '/', '<', '>':
		t.Next()

		d := t.Accept(string(c))

		if t.Accept("=") && (d || c == '*' || c == '/') {
			typ = TokenDelimiter
		}
	case '!':
		t.Next()

		if !t.Accept("=") {
			t.Err = ErrInvalidCharacter

			return t.Error()
		}
	case ')', '}', ']':
		if len(p.tokenDepth) == 0 || p.tokenDepth[len(p.tokenDepth)-1] != byte(c) {
			t.Err = ErrInvalidCharacter

			return t.Error()
		}

		t.Next()
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
		t.Next()
	}

	return t.Return(typ, p.main)
}

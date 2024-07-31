package python

import (
	"strconv"
	"strings"

	"vimagination.zapto.org/parser"
)

func Unquote(str string) (string, error) {
	t := parser.NewStringTokeniser(str)

	raw := parseStringRaw(&t)
	so := parseStringOpening(&t)

	if so == stringEmpty {
		return "", nil
	}

	m := so.Quote()
	triple := so.IsTriple()
	except := "\n" + so.Quote()

	if !raw {
		except += "\\"
	}

	t.Get()

	var ret strings.Builder

	ret.Grow(len(str))

	for {
		switch t.ExceptRun(except) {
		default:
			return "", strconv.ErrSyntax
		case '\\':
			ret.WriteString(t.Get())

			r := unescapeEscaped(&t)

			if r < 0 {
				return "", strconv.ErrSyntax
			}

			ret.WriteRune(r)

			t.Get()
		case '\n':
			if !triple {
				return "", strconv.ErrSyntax
			}

			t.Next()
		case '\'', '"':
			ret.WriteString(t.Get())
			t.Next()

			if !triple || t.Accept(m) && t.Accept(m) {
				return ret.String(), nil
			}
		}
	}
}

func unescapeEscaped(t *parser.Tokeniser) rune {
	t.Next()
	t.Get()

	c := t.Peek()

	if t.Accept(octalDigit) {
		return readEscapedDigits(t, octalDigit, 8, 2)
	}

	t.Next()

	switch c {
	case '\\', '\'', '"':
		return c
	case 'a':
		return 7
	case 'b':
		return 8
	case 'f':
		return 12
	case 'n':
		return 10
	case 'r':
		return 13
	case 't':
		return 8
	case 'v':
		return 11
	case 'x':
		t.Next()

		return readEscapedDigits(t, hexDigit, 16, 2)
	case 'N':
		return -1 // currently unsupported
	case 'u':
		t.Next()

		return readEscapedDigits(t, hexDigit, 16, 4)
	case 'U':
		t.Next()

		return readEscapedDigits(t, hexDigit, 16, 8)
	}

	return -1
}

func readEscapedDigits(t *parser.Tokeniser, digits string, base, num int) rune {
	for ; num > 0; num-- {
		if !t.Accept(digits) {
			return -1
		}
	}

	n, _ := strconv.ParseInt(t.Get(), base, 32)

	return rune(n)
}

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

func WrapConditional(p ConditionalWrappable) *ConditionalExpression {
	if c, ok := p.(*ConditionalExpression); ok {
		return c
	}

	if c, ok := p.(ConditionalExpression); ok {
		return &c
	}

	c := &ConditionalExpression{}

	switch p := p.(type) {
	case *OrTest:
		c.OrTest = *p

		goto OrTest
	case OrTest:
		c.OrTest = p

		goto OrTest
	case *AndTest:
		c.OrTest.AndTest = *p

		goto AndTest
	case AndTest:
		c.OrTest.AndTest = p

		goto AndTest
	case *NotTest:
		c.OrTest.AndTest.NotTest = *p

		goto NotTest
	case NotTest:
		c.OrTest.AndTest.NotTest = p

		goto NotTest
	case *Comparison:
		c.OrTest.AndTest.NotTest.Comparison = *p

		goto Comparison
	case Comparison:
		c.OrTest.AndTest.NotTest.Comparison = p

		goto Comparison
	case *OrExpression:
		c.OrTest.AndTest.NotTest.Comparison.OrExpression = *p

		goto OrExpression
	case OrExpression:
		c.OrTest.AndTest.NotTest.Comparison.OrExpression = p

		goto OrExpression
	case *XorExpression:
		c.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpressions = *p

		goto XorExpression
	case XorExpression:
		c.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpressions = p

		goto XorExpression
	case *AndExpression:
		c.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpressions.AndExpression = *p

		goto AndExpression
	case AndExpression:
		c.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpressions.AndExpression = p

		goto AndExpression
	case *ShiftExpression:
		c.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpressions.AndExpression.ShiftExpression = *p

		goto ShiftExpression
	case ShiftExpression:
		c.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpressions.AndExpression.ShiftExpression = p

		goto ShiftExpression
	case *AddExpression:
		c.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpressions.AndExpression.ShiftExpression.AddExpression = *p

		goto AddExpression
	case AddExpression:
		c.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpressions.AndExpression.ShiftExpression.AddExpression = p

		goto AddExpression
	case *MultiplyExpression:
		c.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpressions.AndExpression.ShiftExpression.AddExpression.MultiplyExpression = *p

		goto MultiplyExpression
	case MultiplyExpression:
		c.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpressions.AndExpression.ShiftExpression.AddExpression.MultiplyExpression = p

		goto MultiplyExpression
	case *UnaryExpression:
		c.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpressions.AndExpression.ShiftExpression.AddExpression.MultiplyExpression.UnaryExpression = *p

		goto UnaryExpression
	case UnaryExpression:
		c.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpressions.AndExpression.ShiftExpression.AddExpression.MultiplyExpression.UnaryExpression = p

		goto UnaryExpression
	case *PowerExpression:
		c.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpressions.AndExpression.ShiftExpression.AddExpression.MultiplyExpression.UnaryExpression.PowerExpression = p

		goto PowerExpression
	case PowerExpression:
		c.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpressions.AndExpression.ShiftExpression.AddExpression.MultiplyExpression.UnaryExpression.PowerExpression = &p

		goto PowerExpression
	case *PrimaryExpression:
		c.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpressions.AndExpression.ShiftExpression.AddExpression.MultiplyExpression.UnaryExpression.PowerExpression.PrimaryExpression = *p
	case PrimaryExpression:
		c.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpressions.AndExpression.ShiftExpression.AddExpression.MultiplyExpression.UnaryExpression.PowerExpression.PrimaryExpression = p
	}

	c.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpressions.AndExpression.ShiftExpression.AddExpression.MultiplyExpression.UnaryExpression.PowerExpression.Tokens = c.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpressions.AndExpression.ShiftExpression.AddExpression.MultiplyExpression.UnaryExpression.PowerExpression.PrimaryExpression.Tokens
PowerExpression:
	c.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpressions.AndExpression.ShiftExpression.AddExpression.MultiplyExpression.UnaryExpression.Tokens = c.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpressions.AndExpression.ShiftExpression.AddExpression.MultiplyExpression.UnaryExpression.PowerExpression.Tokens
UnaryExpression:
	c.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpressions.AndExpression.ShiftExpression.AddExpression.MultiplyExpression.Tokens = c.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpressions.AndExpression.ShiftExpression.AddExpression.MultiplyExpression.UnaryExpression.Tokens
MultiplyExpression:
	c.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpressions.AndExpression.ShiftExpression.AddExpression.Tokens = c.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpressions.AndExpression.ShiftExpression.AddExpression.MultiplyExpression.Tokens
AddExpression:
	c.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpressions.AndExpression.ShiftExpression.Tokens = c.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpressions.AndExpression.ShiftExpression.AddExpression.Tokens
ShiftExpression:
	c.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpressions.AndExpression.Tokens = c.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpressions.AndExpression.ShiftExpression.Tokens
AndExpression:
	c.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpressions.Tokens = c.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpressions.AndExpression.Tokens
XorExpression:
	c.OrTest.AndTest.NotTest.Comparison.OrExpression.Tokens = c.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpressions.Tokens
OrExpression:
	c.OrTest.AndTest.NotTest.Comparison.Tokens = c.OrTest.AndTest.NotTest.Comparison.OrExpression.Tokens
Comparison:
	c.OrTest.AndTest.NotTest.Tokens = c.OrTest.AndTest.NotTest.Comparison.Tokens
NotTest:
	c.OrTest.AndTest.Tokens = c.OrTest.AndTest.NotTest.Tokens
AndTest:
	c.OrTest.Tokens = c.OrTest.AndTest.Tokens
OrTest:
	c.Tokens = c.OrTest.Tokens

	return c
}

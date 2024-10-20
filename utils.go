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
		return '\a'
	case 'b':
		return '\b'
	case 'f':
		return '\f'
	case 'n':
		return '\n'
	case 'r':
		return '\r'
	case 't':
		return '\t'
	case 'v':
		return '\v'
	case 'x':
		t.Get()

		return readEscapedDigits(t, hexDigit, 16, 2)
	case 'N':
		return -1 // currently unsupported
	case 'u':
		t.Get()

		return readEscapedDigits(t, hexDigit, 16, 4)
	case 'U':
		t.Get()

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
		c.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpression = *p

		goto XorExpression
	case XorExpression:
		c.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpression = p

		goto XorExpression
	case *AndExpression:
		c.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpression.AndExpression = *p

		goto AndExpression
	case AndExpression:
		c.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpression.AndExpression = p

		goto AndExpression
	case *ShiftExpression:
		c.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpression.AndExpression.ShiftExpression = *p

		goto ShiftExpression
	case ShiftExpression:
		c.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpression.AndExpression.ShiftExpression = p

		goto ShiftExpression
	case *AddExpression:
		c.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpression.AndExpression.ShiftExpression.AddExpression = *p

		goto AddExpression
	case AddExpression:
		c.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpression.AndExpression.ShiftExpression.AddExpression = p

		goto AddExpression
	case *MultiplyExpression:
		c.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpression.AndExpression.ShiftExpression.AddExpression.MultiplyExpression = *p

		goto MultiplyExpression
	case MultiplyExpression:
		c.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpression.AndExpression.ShiftExpression.AddExpression.MultiplyExpression = p

		goto MultiplyExpression
	case *UnaryExpression:
		c.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpression.AndExpression.ShiftExpression.AddExpression.MultiplyExpression.UnaryExpression = *p

		goto UnaryExpression
	case UnaryExpression:
		c.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpression.AndExpression.ShiftExpression.AddExpression.MultiplyExpression.UnaryExpression = p

		goto UnaryExpression
	case *PowerExpression:
		c.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpression.AndExpression.ShiftExpression.AddExpression.MultiplyExpression.UnaryExpression.PowerExpression = p

		goto PowerExpression
	case PowerExpression:
		c.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpression.AndExpression.ShiftExpression.AddExpression.MultiplyExpression.UnaryExpression.PowerExpression = &p

		goto PowerExpression
	case *PrimaryExpression:
		c.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpression.AndExpression.ShiftExpression.AddExpression.MultiplyExpression.UnaryExpression.PowerExpression = &PowerExpression{
			PrimaryExpression: *p,
		}

		goto PrimaryExpression
	case PrimaryExpression:
		c.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpression.AndExpression.ShiftExpression.AddExpression.MultiplyExpression.UnaryExpression.PowerExpression = &PowerExpression{
			PrimaryExpression: p,
		}

		goto PrimaryExpression
	case *Atom:
		c.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpression.AndExpression.ShiftExpression.AddExpression.MultiplyExpression.UnaryExpression.PowerExpression = &PowerExpression{
			PrimaryExpression: PrimaryExpression{
				Atom: p,
			},
		}
	case Atom:
		c.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpression.AndExpression.ShiftExpression.AddExpression.MultiplyExpression.UnaryExpression.PowerExpression = &PowerExpression{
			PrimaryExpression: PrimaryExpression{
				Atom: &p,
			},
		}
	}

	c.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpression.AndExpression.ShiftExpression.AddExpression.MultiplyExpression.UnaryExpression.PowerExpression.PrimaryExpression.Tokens = c.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpression.AndExpression.ShiftExpression.AddExpression.MultiplyExpression.UnaryExpression.PowerExpression.PrimaryExpression.Atom.Tokens
PrimaryExpression:
	c.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpression.AndExpression.ShiftExpression.AddExpression.MultiplyExpression.UnaryExpression.PowerExpression.Tokens = c.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpression.AndExpression.ShiftExpression.AddExpression.MultiplyExpression.UnaryExpression.PowerExpression.PrimaryExpression.Tokens
PowerExpression:
	c.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpression.AndExpression.ShiftExpression.AddExpression.MultiplyExpression.UnaryExpression.Tokens = c.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpression.AndExpression.ShiftExpression.AddExpression.MultiplyExpression.UnaryExpression.PowerExpression.Tokens
UnaryExpression:
	c.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpression.AndExpression.ShiftExpression.AddExpression.MultiplyExpression.Tokens = c.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpression.AndExpression.ShiftExpression.AddExpression.MultiplyExpression.UnaryExpression.Tokens
MultiplyExpression:
	c.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpression.AndExpression.ShiftExpression.AddExpression.Tokens = c.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpression.AndExpression.ShiftExpression.AddExpression.MultiplyExpression.Tokens
AddExpression:
	c.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpression.AndExpression.ShiftExpression.Tokens = c.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpression.AndExpression.ShiftExpression.AddExpression.Tokens
ShiftExpression:
	c.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpression.AndExpression.Tokens = c.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpression.AndExpression.ShiftExpression.Tokens
AndExpression:
	c.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpression.Tokens = c.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpression.AndExpression.Tokens
XorExpression:
	c.OrTest.AndTest.NotTest.Comparison.OrExpression.Tokens = c.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpression.Tokens
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

func UnwrapConditional(c *ConditionalExpression) ConditionalWrappable {
	if c == nil {
		return nil
	} else if c.If != nil {
		return c
	} else if c.OrTest.OrTest != nil {
		return &c.OrTest
	} else if c.OrTest.AndTest.AndTest != nil {
		return &c.OrTest.AndTest
	} else if c.OrTest.AndTest.NotTest.Nots != 0 {
		return &c.OrTest.AndTest.NotTest
	} else if c.OrTest.AndTest.NotTest.Comparison.Comparisons != nil {
		return &c.OrTest.AndTest.NotTest.Comparison
	} else if c.OrTest.AndTest.NotTest.Comparison.OrExpression.OrExpression != nil {
		return &c.OrTest.AndTest.NotTest.Comparison.OrExpression
	} else if c.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpression.XorExpression != nil {
		return &c.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpression
	} else if c.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpression.AndExpression.AndExpression != nil {
		return &c.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpression.AndExpression
	} else if c.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpression.AndExpression.ShiftExpression.ShiftExpression != nil {
		return &c.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpression.AndExpression.ShiftExpression
	} else if c.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpression.AndExpression.ShiftExpression.AddExpression.AddExpression != nil {
		return &c.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpression.AndExpression.ShiftExpression.AddExpression
	} else if c.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpression.AndExpression.ShiftExpression.AddExpression.MultiplyExpression.MultiplyExpression != nil {
		return &c.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpression.AndExpression.ShiftExpression.AddExpression.MultiplyExpression
	} else if c.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpression.AndExpression.ShiftExpression.AddExpression.MultiplyExpression.UnaryExpression.UnaryExpression != nil {
		return &c.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpression.AndExpression.ShiftExpression.AddExpression.MultiplyExpression.UnaryExpression
	} else if c.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpression.AndExpression.ShiftExpression.AddExpression.MultiplyExpression.UnaryExpression.PowerExpression == nil {
		return nil
	} else if c.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpression.AndExpression.ShiftExpression.AddExpression.MultiplyExpression.UnaryExpression.PowerExpression.AwaitExpression || c.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpression.AndExpression.ShiftExpression.AddExpression.MultiplyExpression.UnaryExpression.PowerExpression.UnaryExpression != nil {
		return c.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpression.AndExpression.ShiftExpression.AddExpression.MultiplyExpression.UnaryExpression.PowerExpression
	} else if c.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpression.AndExpression.ShiftExpression.AddExpression.MultiplyExpression.UnaryExpression.PowerExpression.PrimaryExpression.Atom == nil {
		return &c.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpression.AndExpression.ShiftExpression.AddExpression.MultiplyExpression.UnaryExpression.PowerExpression.PrimaryExpression
	} else {
		return c.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpression.AndExpression.ShiftExpression.AddExpression.MultiplyExpression.UnaryExpression.PowerExpression.PrimaryExpression.Atom
	}
}

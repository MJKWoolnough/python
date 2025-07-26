package python

import (
	"fmt"
	"io"
	"strings"
	"unsafe"

	"vimagination.zapto.org/parser"
)

var (
	indent = []byte{'\t'}
	space  = []byte{' '}
)

type writer interface {
	io.Writer
	io.StringWriter
	Underlying() writer
	Indent() writer
	Printf(string, ...any)
}

type indentPrinter struct {
	writer
	hadNewline bool
}

func (i *indentPrinter) Write(p []byte) (int, error) {
	var (
		total int
		last  int
	)

	for n, c := range p {
		if c == '\n' {
			if last != n {
				if err := i.printIndent(); err != nil {
					return total, err
				}
			}

			m, err := i.writer.Write(p[last : n+1])
			total += m

			if err != nil {
				return total, err
			}

			i.hadNewline = true
			last = n + 1
		}
	}

	if last != len(p) {
		if err := i.printIndent(); err != nil {
			return total, err
		}

		m, err := i.writer.Write(p[last:])
		total += m

		if err != nil {
			return total, err
		}
	}

	return total, nil
}

func (i *indentPrinter) printIndent() error {
	if i.hadNewline {
		if _, err := i.writer.Write(indent); err != nil {
			return err
		}

		i.hadNewline = false
	}

	return nil
}

func (i *indentPrinter) Printf(format string, args ...any) {
	fmt.Fprintf(i, format, args...)
}

func (i *indentPrinter) WriteString(s string) (int, error) {
	return i.Write(unsafe.Slice(unsafe.StringData(s), len(s)))
}

func (i *indentPrinter) Indent() writer {
	return &indentPrinter{writer: i}
}

type countPrinter struct {
	io.Writer
	pos int
}

func (c *countPrinter) Write(p []byte) (int, error) {
	for _, b := range p {
		if b == '\n' {
			c.pos = 0
		} else if b != '\t' || c.pos > 0 {
			c.pos++
		}
	}

	return c.Writer.Write(p)
}

func (c *countPrinter) WriteString(s string) (int, error) {
	return c.Write(unsafe.Slice(unsafe.StringData(s), len(s)))
}

func (c *countPrinter) Underlying() writer {
	return c
}

func (c *countPrinter) Indent() writer {
	return &indentPrinter{writer: c}
}

func (c *countPrinter) Printf(format string, args ...any) {
	fmt.Fprintf(c, format, args...)
}

func (t Token) printType(w writer, v bool) {
	var typ string

	switch t.Type {
	case TokenWhitespace:
		typ = "Whitespace"
	case TokenLineTerminator:
		typ = "LineTerminator"
	case TokenComment:
		typ = "Comment"
	case TokenIdentifier:
		typ = "Identifier"
	case TokenKeyword:
		typ = "Keyword"
	case TokenOperator:
		typ = "Operator"
	case TokenDelimiter:
		typ = "Delimiter"
	case TokenBooleanLiteral:
		typ = "BooleanLiteral"
	case TokenNumericLiteral:
		typ = "NumericLiteral"
	case TokenStringLiteral:
		typ = "StringLiteral"
	case TokenNullLiteral:
		typ = "NullLiteral"
	case TokenIndent:
		typ = "Indent"
	case TokenDedent:
		typ = "Dedent"
	case parser.TokenDone:
		typ = "Done"
	case parser.TokenError:
		typ = "Error"
	default:
		typ = "Unknown"
	}

	fmt.Fprintf(w, "Type: %s - Data: %q", typ, t.Data)

	if v {
		fmt.Fprintf(w, " - Position: %d (%d: %d)", t.Pos, t.Line, t.LinePos)
	}
}

func (t Tokens) printType(w writer, v bool) {
	if t == nil {
		io.WriteString(w, "nil")

		return
	}

	if len(t) == 0 {
		io.WriteString(w, "[]")

		return
	}

	io.WriteString(w, "[")

	ipp := indentPrinter{writer: w}

	for n, t := range t {
		ipp.Printf("\n%d: ", n)
		t.printType(w, v)
	}

	io.WriteString(w, "\n]")
}

func (c Comments) printType(w writer, v bool) {
	Tokens(c).printType(w, v)
}

func (c Comments) printSource(w writer, v bool) {
	if len(c) > 0 {
		printComment(w, c[0].Data)

		line := c[0].Line

		for _, c := range c[1:] {
			io.WriteString(w, "\n")

			line++

			if line < c.Line {
				io.WriteString(w, "\n")

				line++
			}

			printComment(w, c.Data)
		}

		if v {
			io.WriteString(w, "\n")
		}
	}
}

func printComment(w writer, c string) {
	if !strings.HasPrefix(c, "#") {
		io.WriteString(w, "#")
	}

	io.WriteString(w, c)
}

func (s StatementType) printType(w writer, _ bool) {
	io.WriteString(w, s.String())
}

// String implements the fmt.Stringer interface.
func (s StatementType) String() string {
	switch s {
	case StatementAssert:
		return "StatementAssert"
	case StatementAssignment:
		return "StatementAssignment"
	case StatementAugmentedAssignment:
		return "StatementAugmentedAssignment"
	case StatementAnnotatedAssignment:
		return "StatementAnnotatedAssignment"
	case StatementPass:
		return "StatementPass"
	case StatementDel:
		return "StatementDel"
	case StatementReturn:
		return "StatementReturn"
	case StatementYield:
		return "StatementYield"
	case StatementRaise:
		return "StatementRaise"
	case StatementBreak:
		return "StatementBreak"
	case StatementContinue:
		return "StatementContinue"
	case StatementImport:
		return "StatementImport"
	case StatementGlobal:
		return "StatementGlobal"
	case StatementNonLocal:
		return "StatementNonLocal"
	case StatementTyp:
		return "StatementTyp"
	default:
		return "Unknown"
	}
}

func (t TypeParamType) printType(w writer, _ bool) {
	io.WriteString(w, t.String())
}

// String implements the fmt.Stringer interface.
func (t TypeParamType) String() string {
	switch t {
	case TypeParamIdentifer:
		return "TypeParamIdentifer"
	case TypeParamVar:
		return "TypeParamVar"
	case TypeParamVarTuple:
		return "TypeParamVarTuple"
	default:
		return "Unknown"
	}
}

type formatter interface {
	printType(writer, bool)
	printSource(writer, bool)
}

func format(f formatter, s fmt.State, v rune) {
	switch v {
	case 'v':
		f.printType(&countPrinter{Writer: s}, s.Flag('+'))
	case 's':
		f.printSource(&countPrinter{Writer: s}, s.Flag('+'))
	}
}

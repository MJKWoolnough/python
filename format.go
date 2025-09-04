package python

import (
	"fmt"
	"io"
	"strings"
	"unsafe"

	"vimagination.zapto.org/parser"
)

var indent = []byte{'\t'}

type writer interface {
	io.Writer
	WriteString(string)
	Underlying() writer
	LastChar() byte
	Indent() writer
	IndentMultiline() writer
	InMultiline() bool
	Printf(string, ...any)
}

type indentPrinter struct {
	writer
	hadNewline  bool
	inMultiline bool
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

func (i *indentPrinter) WriteString(s string) {
	i.Write(unsafe.Slice(unsafe.StringData(s), len(s)))
}

func (i *indentPrinter) Indent() writer {
	return &indentPrinter{writer: i, inMultiline: i.inMultiline}
}

func (i *indentPrinter) IndentMultiline() writer {
	return &indentPrinter{writer: i, inMultiline: true}
}

func (i *indentPrinter) InMultiline() bool {
	return i.inMultiline
}

type underlyingWriter struct {
	io.Writer
	lastChar byte
}

func (u *underlyingWriter) Write(p []byte) (int, error) {
	if len(p) > 0 {
		u.lastChar = p[len(p)-1]
	}

	return u.Writer.Write(p)
}

func (u *underlyingWriter) WriteString(s string) {
	u.Write(unsafe.Slice(unsafe.StringData(s), len(s)))
}

func (u *underlyingWriter) Underlying() writer {
	return u
}

func (u *underlyingWriter) LastChar() byte {
	return u.lastChar
}

func (u *underlyingWriter) Indent() writer {
	return &indentPrinter{writer: u}
}

func (u *underlyingWriter) IndentMultiline() writer {
	return &indentPrinter{writer: u, inMultiline: true}
}

func (u *underlyingWriter) InMultiline() bool {
	return false
}

func (u *underlyingWriter) Printf(format string, args ...any) {
	fmt.Fprintf(u, format, args...)
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

	w.Printf("Type: %s - Data: %q", typ, t.Data)

	if v {
		w.Printf(" - Position: %d (%d: %d)", t.Pos, t.Line, t.LinePos)
	}
}

func (t Tokens) printType(w writer, v bool) {
	if t == nil {
		w.WriteString("nil")

		return
	}

	if len(t) == 0 {
		w.WriteString("[]")

		return
	}

	w.WriteString("[")

	ipp := indentPrinter{writer: w}

	for n, t := range t {
		ipp.Printf("\n%d: ", n)
		t.printType(w, v)
	}

	w.WriteString("\n]")
}

func (c Comments) printType(w writer, v bool) {
	Tokens(c).printType(w, v)
}

func (c Comments) printSource(w writer, v bool) {
	if len(c) > 0 {
		switch w.LastChar() {
		case 0, ' ', '\n', '\t':
		default:
			w.WriteString(" ")
		}

		printComment(w, c[0].Data)

		line := c[0].Line

		for _, c := range c[1:] {
			w.WriteString("\n")

			line++

			if line < c.Line {
				w.WriteString("\n")

				line++
			}

			printComment(w, c.Data)
		}

		if v {
			w.WriteString("\n")
		}
	}
}

func printComment(w writer, c string) {
	if !strings.HasPrefix(c, "#") {
		w.WriteString("#")
	}

	w.WriteString(c)
}

func (s StatementType) printType(w writer, _ bool) {
	w.WriteString(s.String())
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
	w.WriteString(t.String())
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
		f.printType(&underlyingWriter{Writer: s}, s.Flag('+'))
	case 's':
		f.printSource(&underlyingWriter{Writer: s}, s.Flag('+'))
	}
}

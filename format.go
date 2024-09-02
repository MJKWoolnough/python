package python

import (
	"fmt"
	"io"
)

var indent = []byte{'	'}

type indentPrinter struct {
	io.Writer
}

func (i *indentPrinter) Write(p []byte) (int, error) {
	var (
		total int
		last  int
	)

	for n, c := range p {
		if c == '\n' {
			m, err := i.Writer.Write(p[last : n+1])
			total += m

			if err != nil {
				return total, err
			}

			_, err = i.Writer.Write(indent)
			if err != nil {
				return total, err
			}

			last = n + 1
		}
	}

	if last != len(p) {
		m, err := i.Writer.Write(p[last:])
		total += m

		if err != nil {
			return total, err
		}
	}

	return total, nil
}

func (i *indentPrinter) Print(args ...interface{}) {
	fmt.Fprint(i, args...)
}

func (i *indentPrinter) Printf(format string, args ...interface{}) {
	fmt.Fprintf(i, format, args...)
}

func (i *indentPrinter) WriteString(s string) (int, error) {
	return i.Write([]byte(s))
}

func (t Token) printType(w io.Writer, v bool) {
	var typ string

	switch t.Type {
	default:
		typ = "Unknown"
	}

	fmt.Fprintf(w, "Type: %s - Data: %q", typ, t.Data)

	if v {
		fmt.Fprintf(w, " - Position: %d (%d: %d)", t.Pos, t.Line, t.LinePos)
	}
}

func (t Tokens) printType(w io.Writer, v bool) {
	if len(t) == 0 {
		io.WriteString(w, "[]")

		return
	}

	io.WriteString(w, "[")

	ipp := indentPrinter{w}

	for n, t := range t {
		ipp.Printf("\n%d: ", n)
		t.printType(w, v)
	}

	io.WriteString(w, "]")
}

func (s StatementType) printType(w io.Writer, _ bool) {
	io.WriteString(w, s.String())
}

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

func (t TypeParamType) printType(w io.Writer, _ bool) {
	io.WriteString(w, t.String())
}

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
	printType(io.Writer, bool)
	// printSource(io.Writer, bool)
}

func format(f formatter, s fmt.State, v rune) {
	switch v {
	case 'v':
		f.printType(s, s.Flag('+'))
		// case 's':
		// f.printSource(s, s.Flag('+'))
	}
}

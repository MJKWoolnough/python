package walk

import (
	"errors"
	"reflect"
	"testing"

	"vimagination.zapto.org/parser"
	"vimagination.zapto.org/python"
)

var (
	sentinel = errors.New("")
	nilErr   = errors.New("nil received")
	nilRet   = func(_ *python.File) python.Type { return nil }
)

type walker struct {
	end   python.Type
	level []string
}

func (w *walker) Handle(t python.Type) error {
	if reflect.ValueOf(t).IsNil() {
		return nilErr
	}

	if t == w.end {
		w.level = append(w.level, reflect.TypeOf(t).Elem().Name())

		return sentinel
	}

	err := Walk(t, w)
	if err != nil {
		w.level = append(w.level, reflect.TypeOf(t).Elem().Name())
	}

	return err
}

func TestWalk(t *testing.T) {
	for n, test := range [...]struct {
		Input string
		End   func(f *python.File) python.Type
		Level []string
	}{
		{ // 1
			"",
			nilRet,
			nil,
		},
		{ // 2
			"a\nb",
			func(f *python.File) python.Type {
				return &f.Statements[0]
			},
			[]string{"File", "Statement"},
		},
		{ // 3
			"a\nb",
			func(f *python.File) python.Type {
				return &f.Statements[1]
			},
			[]string{"File", "Statement"},
		},
		{ // 4
			"if a:\n\tb",
			func(f *python.File) python.Type {
				return f.Statements[0].CompoundStatement
			},
			[]string{"File", "Statement", "CompoundStatement"},
		},
		{ // 5
			"a",
			func(f *python.File) python.Type {
				return f.Statements[0].StatementList
			},
			[]string{"File", "Statement", "StatementList"},
		},
		{ // 6
			"if a:\n\tb",
			func(f *python.File) python.Type {
				return f.Statements[0].CompoundStatement.If
			},
			[]string{"File", "Statement", "CompoundStatement", "IfStatement"},
		},
		{ // 7
			"while a:\n\tb",
			func(f *python.File) python.Type {
				return f.Statements[0].CompoundStatement.While
			},
			[]string{"File", "Statement", "CompoundStatement", "WhileStatement"},
		},
		{ // 8
			"for a in b:\n\tc",
			func(f *python.File) python.Type {
				return f.Statements[0].CompoundStatement.For
			},
			[]string{"File", "Statement", "CompoundStatement", "ForStatement"},
		},
		{ // 9
			"try:\n\ta\nexcept b as c:\n\td",
			func(f *python.File) python.Type {
				return f.Statements[0].CompoundStatement.Try
			},
			[]string{"File", "Statement", "CompoundStatement", "TryStatement"},
		},
		{ // 10
			"with a:\n\tb",
			func(f *python.File) python.Type {
				return f.Statements[0].CompoundStatement.With
			},
			[]string{"File", "Statement", "CompoundStatement", "WithStatement"},
		},
		{ // 11
			"def a():\n\tb",
			func(f *python.File) python.Type {
				return f.Statements[0].CompoundStatement.Func
			},
			[]string{"File", "Statement", "CompoundStatement", "FuncDefinition"},
		},
		{ // 12
			"class a():\n\tb",
			func(f *python.File) python.Type {
				return f.Statements[0].CompoundStatement.Class
			},
			[]string{"File", "Statement", "CompoundStatement", "ClassDefinition"},
		},
	} {
		tk := parser.NewStringTokeniser(test.Input)

		m, err := python.Parse(&tk)
		if err != nil {
			t.Errorf("test %d: unexpected error parsing file: %s", n+1, err)
		} else {
			w := walker{end: test.End(m)}

			if err := w.Handle(m); err == nil && test.Level != nil {
				t.Errorf("test %d: expected to recieve sentinel error, but didn't", n+1)
			} else if err != nil && test.Level == nil {
				t.Errorf("test %d: expected no error, but recieved %v", n+1, err)
			} else if len(w.level) != len(test.Level) {
				t.Errorf("test %d: expected to have %d levels, got %d", n+1, len(test.Level), len(w.level))
			} else {
				for m, l := range w.level {
					if e := test.Level[len(test.Level)-m-1]; e != l {
						t.Errorf("test %d.%d: expected to read level %s, got %s", n+1, m+1, e, l)
					}
				}
			}
		}
	}
}

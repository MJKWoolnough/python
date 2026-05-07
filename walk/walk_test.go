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
		{ // 13
			"if a:\n\tb\nelif c: d\nelif e: f\nelse: g",
			nilRet,
			nil,
		},
		{ // 14
			"if a:\n\tb\nelif c: d\nelif e: f\nelse: g",
			func(f *python.File) python.Type {
				return &f.Statements[0].CompoundStatement.If.AssignmentExpression
			},
			[]string{"File", "Statement", "CompoundStatement", "IfStatement", "AssignmentExpression"},
		},
		{ // 15
			"if a:\n\tb\nelif c: d\nelif e: f\nelse: g",
			func(f *python.File) python.Type {
				return &f.Statements[0].CompoundStatement.If.Suite
			},
			[]string{"File", "Statement", "CompoundStatement", "IfStatement", "Suite"},
		},
		{ // 16
			"if a:\n\tb\nelif c: d\nelif e: f\nelse: g",
			func(f *python.File) python.Type {
				return &f.Statements[0].CompoundStatement.If.Elif[0]
			},
			[]string{"File", "Statement", "CompoundStatement", "IfStatement", "AssignmentExpressionAndSuite"},
		},
		{ // 17
			"if a:\n\tb\nelif c: d\nelif e: f\nelse: g",
			func(f *python.File) python.Type {
				return &f.Statements[0].CompoundStatement.If.Elif[1]
			},
			[]string{"File", "Statement", "CompoundStatement", "IfStatement", "AssignmentExpressionAndSuite"},
		},
		{ // 18
			"if a:\n\tb\nelif c: d\nelif e: f\nelse: g",
			func(f *python.File) python.Type {
				return f.Statements[0].CompoundStatement.If.Else
			},
			[]string{"File", "Statement", "CompoundStatement", "IfStatement", "Suite"},
		},
		{ // 19
			"if a: b",
			func(f *python.File) python.Type {
				return f.Statements[0].CompoundStatement.If.Suite.StatementList
			},
			[]string{"File", "Statement", "CompoundStatement", "IfStatement", "Suite", "StatementList"},
		},
		{ // 20
			"if a:\n\tb\n\tc",
			func(f *python.File) python.Type {
				return &f.Statements[0].CompoundStatement.If.Suite.Statements[0]
			},
			[]string{"File", "Statement", "CompoundStatement", "IfStatement", "Suite", "Statement"},
		},
		{ // 21
			"if a:\n\tb\n\tc",
			func(f *python.File) python.Type {
				return &f.Statements[0].CompoundStatement.If.Suite.Statements[1]
			},
			[]string{"File", "Statement", "CompoundStatement", "IfStatement", "Suite", "Statement"},
		},
		{ // 22
			"if a:\n\tb\nelif c: d",
			func(f *python.File) python.Type {
				return &f.Statements[0].CompoundStatement.If.Elif[0].AssignmentExpression
			},
			[]string{"File", "Statement", "CompoundStatement", "IfStatement", "AssignmentExpressionAndSuite", "AssignmentExpression"},
		},
		{ // 23
			"if a:\n\tb\nelif c: d",
			func(f *python.File) python.Type {
				return &f.Statements[0].CompoundStatement.If.Elif[0].Suite
			},
			[]string{"File", "Statement", "CompoundStatement", "IfStatement", "AssignmentExpressionAndSuite", "Suite"},
		},
		{ // 24
			"while a:\n\tb",
			nilRet,
			nil,
		},
		{ // 25
			"while a:\n\tb\nelse: c",
			func(f *python.File) python.Type {
				return &f.Statements[0].CompoundStatement.While.AssignmentExpression
			},
			[]string{"File", "Statement", "CompoundStatement", "WhileStatement", "AssignmentExpression"},
		},
		{ // 26
			"while a:\n\tb\nelse: c",
			func(f *python.File) python.Type {
				return &f.Statements[0].CompoundStatement.While.Suite
			},
			[]string{"File", "Statement", "CompoundStatement", "WhileStatement", "Suite"},
		},
		{ // 27
			"while a:\n\tb\nelse: c",
			func(f *python.File) python.Type {
				return f.Statements[0].CompoundStatement.While.Else
			},
			[]string{"File", "Statement", "CompoundStatement", "WhileStatement", "Suite"},
		},
		{ // 28
			"for a in b:\n\tc",
			nilRet,
			nil,
		},
		{ // 29
			"for a in b:\n\tc\nelse: d",
			func(f *python.File) python.Type {
				return &f.Statements[0].CompoundStatement.For.TargetList
			},
			[]string{"File", "Statement", "CompoundStatement", "ForStatement", "TargetList"},
		},
		{ // 30
			"for a in b:\n\tc\nelse: d",
			func(f *python.File) python.Type {
				return &f.Statements[0].CompoundStatement.For.StarredList
			},
			[]string{"File", "Statement", "CompoundStatement", "ForStatement", "StarredList"},
		},
		{ // 31
			"for a in b:\n\tc\nelse: d",
			func(f *python.File) python.Type {
				return &f.Statements[0].CompoundStatement.For.Suite
			},
			[]string{"File", "Statement", "CompoundStatement", "ForStatement", "Suite"},
		},
		{ // 32
			"for a in b:\n\tc\nelse: d",
			func(f *python.File) python.Type {
				return f.Statements[0].CompoundStatement.For.Else
			},
			[]string{"File", "Statement", "CompoundStatement", "ForStatement", "Suite"},
		},
		{ // 33
			"for a, b in c: d",
			func(f *python.File) python.Type {
				return &f.Statements[0].CompoundStatement.For.TargetList.Targets[0]
			},
			[]string{"File", "Statement", "CompoundStatement", "ForStatement", "TargetList", "Target"},
		},
		{ // 34
			"for a, b in c: d",
			func(f *python.File) python.Type {
				return &f.Statements[0].CompoundStatement.For.TargetList.Targets[1]
			},
			[]string{"File", "Statement", "CompoundStatement", "ForStatement", "TargetList", "Target"},
		},
		{ // 35
			"for a in b, c: d",
			func(f *python.File) python.Type {
				return &f.Statements[0].CompoundStatement.For.StarredList.StarredItems[0]
			},
			[]string{"File", "Statement", "CompoundStatement", "ForStatement", "StarredList", "StarredItem"},
		},
		{ // 36
			"for a in b, c: d",
			func(f *python.File) python.Type {
				return &f.Statements[0].CompoundStatement.For.StarredList.StarredItems[1]
			},
			[]string{"File", "Statement", "CompoundStatement", "ForStatement", "StarredList", "StarredItem"},
		},
		{ // 37
			"try:\n\ta\nexcept b as c:\n\td\nexcept e as f: g\nelse: h",
			nilRet,
			nil,
		},
		{ // 38
			"try:\n\ta\nexcept b as c:\n\td\nexcept e as f: g\nelse: h\nfinally: i",
			func(f *python.File) python.Type {
				return &f.Statements[0].CompoundStatement.Try.Try
			},
			[]string{"File", "Statement", "CompoundStatement", "TryStatement", "Suite"},
		},
		{ // 39
			"try:\n\ta\nexcept b as c:\n\td\nexcept e as f: g\nelse: h\nfinally: i",
			func(f *python.File) python.Type {
				return &f.Statements[0].CompoundStatement.Try.Except[0]
			},
			[]string{"File", "Statement", "CompoundStatement", "TryStatement", "Except"},
		},
		{ // 40
			"try:\n\ta\nexcept b as c:\n\td\nexcept e as f: g\nelse: h\nfinally: i",
			func(f *python.File) python.Type {
				return &f.Statements[0].CompoundStatement.Try.Except[1]
			},
			[]string{"File", "Statement", "CompoundStatement", "TryStatement", "Except"},
		},
		{ // 41
			"try:\n\ta\nexcept b as c:\n\td\nexcept e as f: g\nelse: h\nfinally: i",
			func(f *python.File) python.Type {
				return f.Statements[0].CompoundStatement.Try.Else
			},
			[]string{"File", "Statement", "CompoundStatement", "TryStatement", "Suite"},
		},
		{ // 42
			"try:\n\ta\nexcept b as c:\n\td\nexcept e as f: g\nelse: h\nfinally: i",
			func(f *python.File) python.Type {
				return f.Statements[0].CompoundStatement.Try.Finally
			},
			[]string{"File", "Statement", "CompoundStatement", "TryStatement", "Suite"},
		},
		{ // 43
			"try:\n\ta\nexcept b as c: d",
			func(f *python.File) python.Type {
				return &f.Statements[0].CompoundStatement.Try.Except[0].Expression
			},
			[]string{"File", "Statement", "CompoundStatement", "TryStatement", "Except", "Expression"},
		},
		{ // 44
			"try:\n\ta\nexcept b as c: d",
			func(f *python.File) python.Type {
				return &f.Statements[0].CompoundStatement.Try.Except[0].Suite
			},
			[]string{"File", "Statement", "CompoundStatement", "TryStatement", "Except", "Suite"},
		},
		{ // 45
			"with a:\n\tb",
			func(f *python.File) python.Type {
				return &f.Statements[0].CompoundStatement.With.Contents
			},
			[]string{"File", "Statement", "CompoundStatement", "WithStatement", "WithStatementContents"},
		},
		{ // 46
			"with a:\n\tb",
			func(f *python.File) python.Type {
				return &f.Statements[0].CompoundStatement.With.Suite
			},
			[]string{"File", "Statement", "CompoundStatement", "WithStatement", "Suite"},
		},
		{ // 47
			"with a, b:\n\tc",
			func(f *python.File) python.Type {
				return &f.Statements[0].CompoundStatement.With.Contents.Items[0]
			},
			[]string{"File", "Statement", "CompoundStatement", "WithStatement", "WithStatementContents", "WithItem"},
		},
		{ // 48
			"with a, b:\n\tc",
			func(f *python.File) python.Type {
				return &f.Statements[0].CompoundStatement.With.Contents.Items[1]
			},
			[]string{"File", "Statement", "CompoundStatement", "WithStatement", "WithStatementContents", "WithItem"},
		},
		{ // 49
			"with a as b:\n\tc",
			func(f *python.File) python.Type {
				return &f.Statements[0].CompoundStatement.With.Contents.Items[0].Expression
			},
			[]string{"File", "Statement", "CompoundStatement", "WithStatement", "WithStatementContents", "WithItem", "Expression"},
		},
		{ // 50
			"with a as b:\n\tc",
			func(f *python.File) python.Type {
				return f.Statements[0].CompoundStatement.With.Contents.Items[0].Target
			},
			[]string{"File", "Statement", "CompoundStatement", "WithStatement", "WithStatementContents", "WithItem", "Target"},
		},
		{ // 51
			"@a\ndef b[c]() -> d: e",
			func(f *python.File) python.Type {
				return f.Statements[0].CompoundStatement.Func.Decorators
			},
			[]string{"File", "Statement", "CompoundStatement", "FuncDefinition", "Decorators"},
		},
		{ // 52
			"@a\ndef b[c]() -> d: e",
			func(f *python.File) python.Type {
				return f.Statements[0].CompoundStatement.Func.TypeParams
			},
			[]string{"File", "Statement", "CompoundStatement", "FuncDefinition", "TypeParams"},
		},
		{ // 53
			"@a\ndef b[c]() -> d: e",
			func(f *python.File) python.Type {
				return &f.Statements[0].CompoundStatement.Func.ParameterList
			},
			[]string{"File", "Statement", "CompoundStatement", "FuncDefinition", "ParameterList"},
		},
		{ // 54
			"@a\ndef b[c]() -> d: e",
			func(f *python.File) python.Type {
				return f.Statements[0].CompoundStatement.Func.Expression
			},
			[]string{"File", "Statement", "CompoundStatement", "FuncDefinition", "Expression"},
		},
		{ // 55
			"@a\ndef b[c]() -> d: e",
			func(f *python.File) python.Type {
				return &f.Statements[0].CompoundStatement.Func.Suite
			},
			[]string{"File", "Statement", "CompoundStatement", "FuncDefinition", "Suite"},
		},
		{ // 56
			"@a\n@b\ndef c(): d",
			func(f *python.File) python.Type {
				return &f.Statements[0].CompoundStatement.Func.Decorators.Decorators[0]
			},
			[]string{"File", "Statement", "CompoundStatement", "FuncDefinition", "Decorators", "Decorator"},
		},
		{ // 57
			"@a\n@b\ndef c(): d",
			func(f *python.File) python.Type {
				return &f.Statements[0].CompoundStatement.Func.Decorators.Decorators[1]
			},
			[]string{"File", "Statement", "CompoundStatement", "FuncDefinition", "Decorators", "Decorator"},
		},
		{ // 58
			"@a\ndef b(): c",
			func(f *python.File) python.Type {
				return &f.Statements[0].CompoundStatement.Func.Decorators.Decorators[0].Decorator
			},
			[]string{"File", "Statement", "CompoundStatement", "FuncDefinition", "Decorators", "Decorator", "AssignmentExpression"},
		},
		{ // 59
			"def a[b, c]() -> d: e",
			func(f *python.File) python.Type {
				return &f.Statements[0].CompoundStatement.Func.TypeParams.TypeParams[0]
			},
			[]string{"File", "Statement", "CompoundStatement", "FuncDefinition", "TypeParams", "TypeParam"},
		},
		{ // 60
			"def a[b, c]() -> d: e",
			func(f *python.File) python.Type {
				return &f.Statements[0].CompoundStatement.Func.TypeParams.TypeParams[1]
			},
			[]string{"File", "Statement", "CompoundStatement", "FuncDefinition", "TypeParams", "TypeParam"},
		},
		{ // 61
			"def a[b: c]() -> d: e",
			func(f *python.File) python.Type {
				return f.Statements[0].CompoundStatement.Func.TypeParams.TypeParams[0].Expression
			},
			[]string{"File", "Statement", "CompoundStatement", "FuncDefinition", "TypeParams", "TypeParam", "Expression"},
		},
		{ // 62
			"def a(b, c, /, d, e, *f, g, h, **i): j",
			func(f *python.File) python.Type {
				return &f.Statements[0].CompoundStatement.Func.ParameterList.DefParameters[0]
			},
			[]string{"File", "Statement", "CompoundStatement", "FuncDefinition", "ParameterList", "DefParameter"},
		},
		{ // 63
			"def a(b, c, /, d, e, *f, g, h, **i): j",
			func(f *python.File) python.Type {
				return &f.Statements[0].CompoundStatement.Func.ParameterList.DefParameters[1]
			},
			[]string{"File", "Statement", "CompoundStatement", "FuncDefinition", "ParameterList", "DefParameter"},
		},
		{ // 64
			"def a(b, c, /, d, e, *f, g, h, **i): j",
			func(f *python.File) python.Type {
				return &f.Statements[0].CompoundStatement.Func.ParameterList.NoPosOnly[0]
			},
			[]string{"File", "Statement", "CompoundStatement", "FuncDefinition", "ParameterList", "DefParameter"},
		},
		{ // 65
			"def a(b, c, /, d, e, *f, g, h, **i): j",
			func(f *python.File) python.Type {
				return &f.Statements[0].CompoundStatement.Func.ParameterList.NoPosOnly[1]
			},
			[]string{"File", "Statement", "CompoundStatement", "FuncDefinition", "ParameterList", "DefParameter"},
		},
		{ // 66
			"def a(b, c, /, d, e, *f, g, h, **i): j",
			func(f *python.File) python.Type {
				return f.Statements[0].CompoundStatement.Func.ParameterList.StarArg
			},
			[]string{"File", "Statement", "CompoundStatement", "FuncDefinition", "ParameterList", "Parameter"},
		},
		{ // 67
			"def a(b, c, /, d, e, *f, g, h, **i): j",
			func(f *python.File) python.Type {
				return &f.Statements[0].CompoundStatement.Func.ParameterList.StarArgs[0]
			},
			[]string{"File", "Statement", "CompoundStatement", "FuncDefinition", "ParameterList", "DefParameter"},
		},
		{ // 68
			"def a(b, c, /, d, e, *f, g, h, **i): j",
			func(f *python.File) python.Type {
				return &f.Statements[0].CompoundStatement.Func.ParameterList.StarArgs[1]
			},
			[]string{"File", "Statement", "CompoundStatement", "FuncDefinition", "ParameterList", "DefParameter"},
		},
		{ // 69
			"def a(b, c, /, d, e, *f, g, h, **i): j",
			func(f *python.File) python.Type {
				return f.Statements[0].CompoundStatement.Func.ParameterList.StarStarArg
			},
			[]string{"File", "Statement", "CompoundStatement", "FuncDefinition", "ParameterList", "Parameter"},
		},
		{ // 70
			"def a(b = c): d",
			func(f *python.File) python.Type {
				return &f.Statements[0].CompoundStatement.Func.ParameterList.NoPosOnly[0].Parameter
			},
			[]string{"File", "Statement", "CompoundStatement", "FuncDefinition", "ParameterList", "DefParameter", "Parameter"},
		},
		{ // 71
			"def a(b = c): d",
			func(f *python.File) python.Type {
				return f.Statements[0].CompoundStatement.Func.ParameterList.NoPosOnly[0].Value
			},
			[]string{"File", "Statement", "CompoundStatement", "FuncDefinition", "ParameterList", "DefParameter", "Expression"},
		},
		{ // 72
			"def a(b: c): d",
			func(f *python.File) python.Type {
				return f.Statements[0].CompoundStatement.Func.ParameterList.NoPosOnly[0].Parameter.Type
			},
			[]string{"File", "Statement", "CompoundStatement", "FuncDefinition", "ParameterList", "DefParameter", "Parameter", "Expression"},
		},
		{ // 73
			"@a\nclass b[c](d): e",
			func(f *python.File) python.Type {
				return f.Statements[0].CompoundStatement.Class.Decorators
			},
			[]string{"File", "Statement", "CompoundStatement", "ClassDefinition", "Decorators"},
		},
		{ // 74
			"@a\nclass b[c](d): e",
			func(f *python.File) python.Type {
				return f.Statements[0].CompoundStatement.Class.TypeParams
			},
			[]string{"File", "Statement", "CompoundStatement", "ClassDefinition", "TypeParams"},
		},
		{ // 75
			"@a\nclass b[c](d): e",
			func(f *python.File) python.Type {
				return f.Statements[0].CompoundStatement.Class.Inheritance
			},
			[]string{"File", "Statement", "CompoundStatement", "ClassDefinition", "ArgumentList"},
		},
		{ // 76
			"@a\nclass b[c](d): e",
			func(f *python.File) python.Type {
				return &f.Statements[0].CompoundStatement.Class.Suite
			},
			[]string{"File", "Statement", "CompoundStatement", "ClassDefinition", "Suite"},
		},
		{ // 77
			"a; b",
			func(f *python.File) python.Type {
				return &f.Statements[0].StatementList.Statements[0]
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement"},
		},
		{ // 78
			"a; b",
			func(f *python.File) python.Type {
				return &f.Statements[0].StatementList.Statements[1]
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement"},
		},
		{ // 79
			"a",
			func(f *python.File) python.Type {
				return &f.Statements[0].StatementList.Statements[0]
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement"},
		},
		{ // 80
			"assert a",
			func(f *python.File) python.Type {
				return f.Statements[0].StatementList.Statements[0].AssertStatement
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "AssertStatement"},
		},
		{ // 81
			"a = b",
			func(f *python.File) python.Type {
				return f.Statements[0].StatementList.Statements[0].AssignmentStatement
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "AssignmentStatement"},
		},
		{ // 82
			"a += b",
			func(f *python.File) python.Type {
				return f.Statements[0].StatementList.Statements[0].AugmentedAssignmentStatement
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "AugmentedAssignmentStatement"},
		},
		{ // 83
			"a: b = c",
			func(f *python.File) python.Type {
				return f.Statements[0].StatementList.Statements[0].AnnotatedAssignmentStatement
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "AnnotatedAssignmentStatement"},
		},
		{ // 84
			"del a",
			func(f *python.File) python.Type {
				return f.Statements[0].StatementList.Statements[0].DelStatement
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "DelStatement"},
		},
		{ // 85
			"def a(): return b",
			func(f *python.File) python.Type {
				return f.Statements[0].CompoundStatement.Func.Suite.StatementList.Statements[0].ReturnStatement
			},
			[]string{"File", "Statement", "CompoundStatement", "FuncDefinition", "Suite", "StatementList", "SimpleStatement", "ReturnStatement"},
		},
		{ // 86
			"def a(): yield b",
			func(f *python.File) python.Type {
				return f.Statements[0].CompoundStatement.Func.Suite.StatementList.Statements[0].YieldStatement
			},
			[]string{"File", "Statement", "CompoundStatement", "FuncDefinition", "Suite", "StatementList", "SimpleStatement", "YieldExpression"},
		},
		{ // 87
			"raise a",
			func(f *python.File) python.Type {
				return f.Statements[0].StatementList.Statements[0].RaiseStatement
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "RaiseStatement"},
		},
		{ // 88
			"import a",
			func(f *python.File) python.Type {
				return f.Statements[0].StatementList.Statements[0].ImportStatement
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "ImportStatement"},
		},
		{ // 89
			"global a",
			func(f *python.File) python.Type {
				return f.Statements[0].StatementList.Statements[0].GlobalStatement
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "GlobalStatement"},
		},
		{ // 90
			"nonlocal a",
			func(f *python.File) python.Type {
				return f.Statements[0].StatementList.Statements[0].NonLocalStatement
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "NonLocalStatement"},
		},
		{ // 91
			"type a = b",
			func(f *python.File) python.Type {
				return f.Statements[0].StatementList.Statements[0].TypeStatement
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "TypeStatement"},
		},
		{ // 92
			"assert a, b",
			nilRet,
			nil,
		},
		{ // 93
			"assert a, b",
			func(f *python.File) python.Type {
				return &f.Statements[0].StatementList.Statements[0].AssertStatement.Expressions[0]
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "AssertStatement", "Expression"},
		},
		{ // 94
			"assert a, b",
			func(f *python.File) python.Type {
				return &f.Statements[0].StatementList.Statements[0].AssertStatement.Expressions[1]
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "AssertStatement", "Expression"},
		},
		{ // 95
			"a = b = c",
			func(f *python.File) python.Type {
				return &f.Statements[0].StatementList.Statements[0].AssignmentStatement.TargetLists[0]
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "AssignmentStatement", "TargetList"},
		},
		{ // 96
			"a = b = c",
			func(f *python.File) python.Type {
				return &f.Statements[0].StatementList.Statements[0].AssignmentStatement.TargetLists[1]
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "AssignmentStatement", "TargetList"},
		},
		{ // 97
			"a = b",
			func(f *python.File) python.Type {
				return f.Statements[0].StatementList.Statements[0].AssignmentStatement.StarredExpression
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "AssignmentStatement", "StarredExpression"},
		},
		{ // 98
			"a = yield b",
			func(f *python.File) python.Type {
				return f.Statements[0].StatementList.Statements[0].AssignmentStatement.YieldExpression
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "AssignmentStatement", "YieldExpression"},
		},
		{ // 99
			"a",
			func(f *python.File) python.Type {
				return f.Statements[0].StatementList.Statements[0].AssignmentStatement.StarredExpression.Expression
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "AssignmentStatement", "StarredExpression", "Expression"},
		},
		{ // 100
			"*a",
			func(f *python.File) python.Type {
				return f.Statements[0].StatementList.Statements[0].AssignmentStatement.StarredExpression.StarredList
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "AssignmentStatement", "StarredExpression", "StarredList"},
		},
		{ // 101
			"a",
			func(f *python.File) python.Type {
				return f.Statements[0].StatementList.Statements[0].AssignmentStatement.StarredExpression.Expression.ConditionalExpression
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "AssignmentStatement", "StarredExpression", "Expression", "ConditionalExpression"},
		},
		{ // 102
			"lambda b: c",
			func(f *python.File) python.Type {
				return f.Statements[0].StatementList.Statements[0].AssignmentStatement.StarredExpression.Expression.LambdaExpression
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "AssignmentStatement", "StarredExpression", "Expression", "LambdaExpression"},
		},
		{ // 103
			"a if b else c",
			func(f *python.File) python.Type {
				return &f.Statements[0].StatementList.Statements[0].AssignmentStatement.StarredExpression.Expression.ConditionalExpression.OrTest
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "AssignmentStatement", "StarredExpression", "Expression", "ConditionalExpression", "OrTest"},
		},
		{ // 104
			"a if b else c",
			func(f *python.File) python.Type {
				return f.Statements[0].StatementList.Statements[0].AssignmentStatement.StarredExpression.Expression.ConditionalExpression.If
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "AssignmentStatement", "StarredExpression", "Expression", "ConditionalExpression", "OrTest"},
		},
		{ // 105
			"a if b else c",
			func(f *python.File) python.Type {
				return f.Statements[0].StatementList.Statements[0].AssignmentStatement.StarredExpression.Expression.ConditionalExpression.Else
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "AssignmentStatement", "StarredExpression", "Expression", "ConditionalExpression", "Expression"},
		},
		{ // 106
			"a or b",
			func(f *python.File) python.Type {
				return &f.Statements[0].StatementList.Statements[0].AssignmentStatement.StarredExpression.Expression.ConditionalExpression.OrTest.AndTest
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "AssignmentStatement", "StarredExpression", "Expression", "ConditionalExpression", "OrTest", "AndTest"},
		},
		{ // 107
			"a or b",
			func(f *python.File) python.Type {
				return f.Statements[0].StatementList.Statements[0].AssignmentStatement.StarredExpression.Expression.ConditionalExpression.OrTest.OrTest
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "AssignmentStatement", "StarredExpression", "Expression", "ConditionalExpression", "OrTest", "OrTest"},
		},
		{ // 108
			"a and b",
			func(f *python.File) python.Type {
				return &f.Statements[0].StatementList.Statements[0].AssignmentStatement.StarredExpression.Expression.ConditionalExpression.OrTest.AndTest.NotTest
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "AssignmentStatement", "StarredExpression", "Expression", "ConditionalExpression", "OrTest", "AndTest", "NotTest"},
		},
		{ // 109
			"a and b",
			func(f *python.File) python.Type {
				return f.Statements[0].StatementList.Statements[0].AssignmentStatement.StarredExpression.Expression.ConditionalExpression.OrTest.AndTest.AndTest
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "AssignmentStatement", "StarredExpression", "Expression", "ConditionalExpression", "OrTest", "AndTest", "AndTest"},
		},
		{ // 110
			"not a",
			func(f *python.File) python.Type {
				return &f.Statements[0].StatementList.Statements[0].AssignmentStatement.StarredExpression.Expression.ConditionalExpression.OrTest.AndTest.NotTest.Comparison
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "AssignmentStatement", "StarredExpression", "Expression", "ConditionalExpression", "OrTest", "AndTest", "NotTest", "Comparison"},
		},
		{ // 111
			"a in b == c",
			func(f *python.File) python.Type {
				return &f.Statements[0].StatementList.Statements[0].AssignmentStatement.StarredExpression.Expression.ConditionalExpression.OrTest.AndTest.NotTest.Comparison.OrExpression
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "AssignmentStatement", "StarredExpression", "Expression", "ConditionalExpression", "OrTest", "AndTest", "NotTest", "Comparison", "OrExpression"},
		},
		{ // 112
			"a in b == c",
			func(f *python.File) python.Type {
				return &f.Statements[0].StatementList.Statements[0].AssignmentStatement.StarredExpression.Expression.ConditionalExpression.OrTest.AndTest.NotTest.Comparison.Comparisons[0]
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "AssignmentStatement", "StarredExpression", "Expression", "ConditionalExpression", "OrTest", "AndTest", "NotTest", "Comparison", "ComparisonExpression"},
		},
		{ // 113
			"a in b == c",
			func(f *python.File) python.Type {
				return &f.Statements[0].StatementList.Statements[0].AssignmentStatement.StarredExpression.Expression.ConditionalExpression.OrTest.AndTest.NotTest.Comparison.Comparisons[1]
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "AssignmentStatement", "StarredExpression", "Expression", "ConditionalExpression", "OrTest", "AndTest", "NotTest", "Comparison", "ComparisonExpression"},
		},
		{ // 114
			"a in b",
			func(f *python.File) python.Type {
				return &f.Statements[0].StatementList.Statements[0].AssignmentStatement.StarredExpression.Expression.ConditionalExpression.OrTest.AndTest.NotTest.Comparison.Comparisons[0].OrExpression
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "AssignmentStatement", "StarredExpression", "Expression", "ConditionalExpression", "OrTest", "AndTest", "NotTest", "Comparison", "ComparisonExpression", "OrExpression"},
		},
		{ // 115
			"a | b",
			func(f *python.File) python.Type {
				return &f.Statements[0].StatementList.Statements[0].AssignmentStatement.StarredExpression.Expression.ConditionalExpression.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpression
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "AssignmentStatement", "StarredExpression", "Expression", "ConditionalExpression", "OrTest", "AndTest", "NotTest", "Comparison", "OrExpression", "XorExpression"},
		},
		{ // 116
			"a | b",
			func(f *python.File) python.Type {
				return f.Statements[0].StatementList.Statements[0].AssignmentStatement.StarredExpression.Expression.ConditionalExpression.OrTest.AndTest.NotTest.Comparison.OrExpression.OrExpression
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "AssignmentStatement", "StarredExpression", "Expression", "ConditionalExpression", "OrTest", "AndTest", "NotTest", "Comparison", "OrExpression", "OrExpression"},
		},
		{ // 117
			"a ^ b",
			func(f *python.File) python.Type {
				return &f.Statements[0].StatementList.Statements[0].AssignmentStatement.StarredExpression.Expression.ConditionalExpression.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpression.AndExpression
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "AssignmentStatement", "StarredExpression", "Expression", "ConditionalExpression", "OrTest", "AndTest", "NotTest", "Comparison", "OrExpression", "XorExpression", "AndExpression"},
		},
		{ // 118
			"a ^ b",
			func(f *python.File) python.Type {
				return f.Statements[0].StatementList.Statements[0].AssignmentStatement.StarredExpression.Expression.ConditionalExpression.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpression.XorExpression
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "AssignmentStatement", "StarredExpression", "Expression", "ConditionalExpression", "OrTest", "AndTest", "NotTest", "Comparison", "OrExpression", "XorExpression", "XorExpression"},
		},
		{ // 119
			"a & b",
			func(f *python.File) python.Type {
				return &f.Statements[0].StatementList.Statements[0].AssignmentStatement.StarredExpression.Expression.ConditionalExpression.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpression.AndExpression.ShiftExpression
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "AssignmentStatement", "StarredExpression", "Expression", "ConditionalExpression", "OrTest", "AndTest", "NotTest", "Comparison", "OrExpression", "XorExpression", "AndExpression", "ShiftExpression"},
		},
		{ // 120
			"a & b",
			func(f *python.File) python.Type {
				return f.Statements[0].StatementList.Statements[0].AssignmentStatement.StarredExpression.Expression.ConditionalExpression.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpression.AndExpression.AndExpression
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "AssignmentStatement", "StarredExpression", "Expression", "ConditionalExpression", "OrTest", "AndTest", "NotTest", "Comparison", "OrExpression", "XorExpression", "AndExpression", "AndExpression"},
		},
		{ // 121
			"a << b",
			func(f *python.File) python.Type {
				return &f.Statements[0].StatementList.Statements[0].AssignmentStatement.StarredExpression.Expression.ConditionalExpression.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpression.AndExpression.ShiftExpression.AddExpression
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "AssignmentStatement", "StarredExpression", "Expression", "ConditionalExpression", "OrTest", "AndTest", "NotTest", "Comparison", "OrExpression", "XorExpression", "AndExpression", "ShiftExpression", "AddExpression"},
		},
		{ // 122
			"a << b",
			func(f *python.File) python.Type {
				return f.Statements[0].StatementList.Statements[0].AssignmentStatement.StarredExpression.Expression.ConditionalExpression.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpression.AndExpression.ShiftExpression.ShiftExpression
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "AssignmentStatement", "StarredExpression", "Expression", "ConditionalExpression", "OrTest", "AndTest", "NotTest", "Comparison", "OrExpression", "XorExpression", "AndExpression", "ShiftExpression", "ShiftExpression"},
		},
		{ // 123
			"a + b",
			func(f *python.File) python.Type {
				return &f.Statements[0].StatementList.Statements[0].AssignmentStatement.StarredExpression.Expression.ConditionalExpression.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpression.AndExpression.ShiftExpression.AddExpression.MultiplyExpression
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "AssignmentStatement", "StarredExpression", "Expression", "ConditionalExpression", "OrTest", "AndTest", "NotTest", "Comparison", "OrExpression", "XorExpression", "AndExpression", "ShiftExpression", "AddExpression", "MultiplyExpression"},
		},
		{ // 124
			"a + b",
			func(f *python.File) python.Type {
				return f.Statements[0].StatementList.Statements[0].AssignmentStatement.StarredExpression.Expression.ConditionalExpression.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpression.AndExpression.ShiftExpression.AddExpression.AddExpression
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "AssignmentStatement", "StarredExpression", "Expression", "ConditionalExpression", "OrTest", "AndTest", "NotTest", "Comparison", "OrExpression", "XorExpression", "AndExpression", "ShiftExpression", "AddExpression", "AddExpression"},
		},
		{ // 125
			"a * b",
			func(f *python.File) python.Type {
				return &f.Statements[0].StatementList.Statements[0].AssignmentStatement.StarredExpression.Expression.ConditionalExpression.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpression.AndExpression.ShiftExpression.AddExpression.MultiplyExpression.UnaryExpression
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "AssignmentStatement", "StarredExpression", "Expression", "ConditionalExpression", "OrTest", "AndTest", "NotTest", "Comparison", "OrExpression", "XorExpression", "AndExpression", "ShiftExpression", "AddExpression", "MultiplyExpression", "UnaryExpression"},
		},
		{ // 126
			"a * b",
			func(f *python.File) python.Type {
				return f.Statements[0].StatementList.Statements[0].AssignmentStatement.StarredExpression.Expression.ConditionalExpression.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpression.AndExpression.ShiftExpression.AddExpression.MultiplyExpression.MultiplyExpression
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "AssignmentStatement", "StarredExpression", "Expression", "ConditionalExpression", "OrTest", "AndTest", "NotTest", "Comparison", "OrExpression", "XorExpression", "AndExpression", "ShiftExpression", "AddExpression", "MultiplyExpression", "MultiplyExpression"},
		},
		{ // 127
			"a",
			func(f *python.File) python.Type {
				return f.Statements[0].StatementList.Statements[0].AssignmentStatement.StarredExpression.Expression.ConditionalExpression.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpression.AndExpression.ShiftExpression.AddExpression.MultiplyExpression.UnaryExpression.PowerExpression
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "AssignmentStatement", "StarredExpression", "Expression", "ConditionalExpression", "OrTest", "AndTest", "NotTest", "Comparison", "OrExpression", "XorExpression", "AndExpression", "ShiftExpression", "AddExpression", "MultiplyExpression", "UnaryExpression", "PowerExpression"},
		},
		{ // 128
			"+a",
			func(f *python.File) python.Type {
				return f.Statements[0].StatementList.Statements[0].AssignmentStatement.StarredExpression.Expression.ConditionalExpression.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpression.AndExpression.ShiftExpression.AddExpression.MultiplyExpression.UnaryExpression.UnaryExpression
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "AssignmentStatement", "StarredExpression", "Expression", "ConditionalExpression", "OrTest", "AndTest", "NotTest", "Comparison", "OrExpression", "XorExpression", "AndExpression", "ShiftExpression", "AddExpression", "MultiplyExpression", "UnaryExpression", "UnaryExpression"},
		},
		{ // 129
			"a ** b",
			func(f *python.File) python.Type {
				return &f.Statements[0].StatementList.Statements[0].AssignmentStatement.StarredExpression.Expression.ConditionalExpression.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpression.AndExpression.ShiftExpression.AddExpression.MultiplyExpression.UnaryExpression.PowerExpression.PrimaryExpression
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "AssignmentStatement", "StarredExpression", "Expression", "ConditionalExpression", "OrTest", "AndTest", "NotTest", "Comparison", "OrExpression", "XorExpression", "AndExpression", "ShiftExpression", "AddExpression", "MultiplyExpression", "UnaryExpression", "PowerExpression", "PrimaryExpression"},
		},
		{ // 130
			"a ** b",
			func(f *python.File) python.Type {
				return f.Statements[0].StatementList.Statements[0].AssignmentStatement.StarredExpression.Expression.ConditionalExpression.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpression.AndExpression.ShiftExpression.AddExpression.MultiplyExpression.UnaryExpression.PowerExpression.UnaryExpression
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "AssignmentStatement", "StarredExpression", "Expression", "ConditionalExpression", "OrTest", "AndTest", "NotTest", "Comparison", "OrExpression", "XorExpression", "AndExpression", "ShiftExpression", "AddExpression", "MultiplyExpression", "UnaryExpression", "PowerExpression", "UnaryExpression"},
		},
		{ // 131
			"a",
			func(f *python.File) python.Type {
				return f.Statements[0].StatementList.Statements[0].AssignmentStatement.StarredExpression.Expression.ConditionalExpression.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpression.AndExpression.ShiftExpression.AddExpression.MultiplyExpression.UnaryExpression.PowerExpression.PrimaryExpression.Atom
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "AssignmentStatement", "StarredExpression", "Expression", "ConditionalExpression", "OrTest", "AndTest", "NotTest", "Comparison", "OrExpression", "XorExpression", "AndExpression", "ShiftExpression", "AddExpression", "MultiplyExpression", "UnaryExpression", "PowerExpression", "PrimaryExpression", "Atom"},
		},
		{ // 132
			"a()",
			func(f *python.File) python.Type {
				return f.Statements[0].StatementList.Statements[0].AssignmentStatement.StarredExpression.Expression.ConditionalExpression.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpression.AndExpression.ShiftExpression.AddExpression.MultiplyExpression.UnaryExpression.PowerExpression.PrimaryExpression.PrimaryExpression
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "AssignmentStatement", "StarredExpression", "Expression", "ConditionalExpression", "OrTest", "AndTest", "NotTest", "Comparison", "OrExpression", "XorExpression", "AndExpression", "ShiftExpression", "AddExpression", "MultiplyExpression", "UnaryExpression", "PowerExpression", "PrimaryExpression", "PrimaryExpression"},
		},
		{ // 133
			"a()",
			func(f *python.File) python.Type {
				return f.Statements[0].StatementList.Statements[0].AssignmentStatement.StarredExpression.Expression.ConditionalExpression.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpression.AndExpression.ShiftExpression.AddExpression.MultiplyExpression.UnaryExpression.PowerExpression.PrimaryExpression.Call
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "AssignmentStatement", "StarredExpression", "Expression", "ConditionalExpression", "OrTest", "AndTest", "NotTest", "Comparison", "OrExpression", "XorExpression", "AndExpression", "ShiftExpression", "AddExpression", "MultiplyExpression", "UnaryExpression", "PowerExpression", "PrimaryExpression", "ArgumentListOrComprehension"},
		},
		{ // 134
			"a[]",
			func(f *python.File) python.Type {
				return f.Statements[0].StatementList.Statements[0].AssignmentStatement.StarredExpression.Expression.ConditionalExpression.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpression.AndExpression.ShiftExpression.AddExpression.MultiplyExpression.UnaryExpression.PowerExpression.PrimaryExpression.PrimaryExpression
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "AssignmentStatement", "StarredExpression", "Expression", "ConditionalExpression", "OrTest", "AndTest", "NotTest", "Comparison", "OrExpression", "XorExpression", "AndExpression", "ShiftExpression", "AddExpression", "MultiplyExpression", "UnaryExpression", "PowerExpression", "PrimaryExpression", "PrimaryExpression"},
		},
		{ // 135
			"a[]",
			func(f *python.File) python.Type {
				return f.Statements[0].StatementList.Statements[0].AssignmentStatement.StarredExpression.Expression.ConditionalExpression.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpression.AndExpression.ShiftExpression.AddExpression.MultiplyExpression.UnaryExpression.PowerExpression.PrimaryExpression.Slicing
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "AssignmentStatement", "StarredExpression", "Expression", "ConditionalExpression", "OrTest", "AndTest", "NotTest", "Comparison", "OrExpression", "XorExpression", "AndExpression", "ShiftExpression", "AddExpression", "MultiplyExpression", "UnaryExpression", "PowerExpression", "PrimaryExpression", "SliceList"},
		},
		{ // 136
			"a(b)",
			func(f *python.File) python.Type {
				return f.Statements[0].StatementList.Statements[0].AssignmentStatement.StarredExpression.Expression.ConditionalExpression.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpression.AndExpression.ShiftExpression.AddExpression.MultiplyExpression.UnaryExpression.PowerExpression.PrimaryExpression.Call.ArgumentList
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "AssignmentStatement", "StarredExpression", "Expression", "ConditionalExpression", "OrTest", "AndTest", "NotTest", "Comparison", "OrExpression", "XorExpression", "AndExpression", "ShiftExpression", "AddExpression", "MultiplyExpression", "UnaryExpression", "PowerExpression", "PrimaryExpression", "ArgumentListOrComprehension", "ArgumentList"},
		},
		{ // 137
			"a(b for c in d)",
			func(f *python.File) python.Type {
				return f.Statements[0].StatementList.Statements[0].AssignmentStatement.StarredExpression.Expression.ConditionalExpression.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpression.AndExpression.ShiftExpression.AddExpression.MultiplyExpression.UnaryExpression.PowerExpression.PrimaryExpression.Call.Comprehension
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "AssignmentStatement", "StarredExpression", "Expression", "ConditionalExpression", "OrTest", "AndTest", "NotTest", "Comparison", "OrExpression", "XorExpression", "AndExpression", "ShiftExpression", "AddExpression", "MultiplyExpression", "UnaryExpression", "PowerExpression", "PrimaryExpression", "ArgumentListOrComprehension", "Comprehension"},
		},
		{ // 138
			"a(b, c, d=e, f=g, **h, **i)",
			func(f *python.File) python.Type {
				return &f.Statements[0].StatementList.Statements[0].AssignmentStatement.StarredExpression.Expression.ConditionalExpression.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpression.AndExpression.ShiftExpression.AddExpression.MultiplyExpression.UnaryExpression.PowerExpression.PrimaryExpression.Call.ArgumentList.PositionalArguments[0]
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "AssignmentStatement", "StarredExpression", "Expression", "ConditionalExpression", "OrTest", "AndTest", "NotTest", "Comparison", "OrExpression", "XorExpression", "AndExpression", "ShiftExpression", "AddExpression", "MultiplyExpression", "UnaryExpression", "PowerExpression", "PrimaryExpression", "ArgumentListOrComprehension", "ArgumentList", "PositionalArgument"},
		},
		{ // 139
			"a(b, c, d=e, f=g, **h, **i)",
			func(f *python.File) python.Type {
				return &f.Statements[0].StatementList.Statements[0].AssignmentStatement.StarredExpression.Expression.ConditionalExpression.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpression.AndExpression.ShiftExpression.AddExpression.MultiplyExpression.UnaryExpression.PowerExpression.PrimaryExpression.Call.ArgumentList.PositionalArguments[1]
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "AssignmentStatement", "StarredExpression", "Expression", "ConditionalExpression", "OrTest", "AndTest", "NotTest", "Comparison", "OrExpression", "XorExpression", "AndExpression", "ShiftExpression", "AddExpression", "MultiplyExpression", "UnaryExpression", "PowerExpression", "PrimaryExpression", "ArgumentListOrComprehension", "ArgumentList", "PositionalArgument"},
		},
		{ // 140
			"a(b, c, d=e, f=g, **h, **i)",
			func(f *python.File) python.Type {
				return &f.Statements[0].StatementList.Statements[0].AssignmentStatement.StarredExpression.Expression.ConditionalExpression.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpression.AndExpression.ShiftExpression.AddExpression.MultiplyExpression.UnaryExpression.PowerExpression.PrimaryExpression.Call.ArgumentList.StarredAndKeywordArguments[0]
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "AssignmentStatement", "StarredExpression", "Expression", "ConditionalExpression", "OrTest", "AndTest", "NotTest", "Comparison", "OrExpression", "XorExpression", "AndExpression", "ShiftExpression", "AddExpression", "MultiplyExpression", "UnaryExpression", "PowerExpression", "PrimaryExpression", "ArgumentListOrComprehension", "ArgumentList", "StarredOrKeyword"},
		},
		{ // 141
			"a(b, c, d=e, f=g, **h, **i)",
			func(f *python.File) python.Type {
				return &f.Statements[0].StatementList.Statements[0].AssignmentStatement.StarredExpression.Expression.ConditionalExpression.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpression.AndExpression.ShiftExpression.AddExpression.MultiplyExpression.UnaryExpression.PowerExpression.PrimaryExpression.Call.ArgumentList.StarredAndKeywordArguments[1]
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "AssignmentStatement", "StarredExpression", "Expression", "ConditionalExpression", "OrTest", "AndTest", "NotTest", "Comparison", "OrExpression", "XorExpression", "AndExpression", "ShiftExpression", "AddExpression", "MultiplyExpression", "UnaryExpression", "PowerExpression", "PrimaryExpression", "ArgumentListOrComprehension", "ArgumentList", "StarredOrKeyword"},
		},
		{ // 142
			"a(b, c, d=e, f=g, **h, **i)",
			func(f *python.File) python.Type {
				return &f.Statements[0].StatementList.Statements[0].AssignmentStatement.StarredExpression.Expression.ConditionalExpression.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpression.AndExpression.ShiftExpression.AddExpression.MultiplyExpression.UnaryExpression.PowerExpression.PrimaryExpression.Call.ArgumentList.KeywordArguments[0]
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "AssignmentStatement", "StarredExpression", "Expression", "ConditionalExpression", "OrTest", "AndTest", "NotTest", "Comparison", "OrExpression", "XorExpression", "AndExpression", "ShiftExpression", "AddExpression", "MultiplyExpression", "UnaryExpression", "PowerExpression", "PrimaryExpression", "ArgumentListOrComprehension", "ArgumentList", "KeywordArgument"},
		},
		{ // 143
			"a(b, c, d=e, f=g, **h, **i)",
			func(f *python.File) python.Type {
				return &f.Statements[0].StatementList.Statements[0].AssignmentStatement.StarredExpression.Expression.ConditionalExpression.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpression.AndExpression.ShiftExpression.AddExpression.MultiplyExpression.UnaryExpression.PowerExpression.PrimaryExpression.Call.ArgumentList.KeywordArguments[1]
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "AssignmentStatement", "StarredExpression", "Expression", "ConditionalExpression", "OrTest", "AndTest", "NotTest", "Comparison", "OrExpression", "XorExpression", "AndExpression", "ShiftExpression", "AddExpression", "MultiplyExpression", "UnaryExpression", "PowerExpression", "PrimaryExpression", "ArgumentListOrComprehension", "ArgumentList", "KeywordArgument"},
		},
		{ // 144
			"a(b)",
			func(f *python.File) python.Type {
				return f.Statements[0].StatementList.Statements[0].AssignmentStatement.StarredExpression.Expression.ConditionalExpression.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpression.AndExpression.ShiftExpression.AddExpression.MultiplyExpression.UnaryExpression.PowerExpression.PrimaryExpression.Call.ArgumentList.PositionalArguments[0].AssignmentExpression
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "AssignmentStatement", "StarredExpression", "Expression", "ConditionalExpression", "OrTest", "AndTest", "NotTest", "Comparison", "OrExpression", "XorExpression", "AndExpression", "ShiftExpression", "AddExpression", "MultiplyExpression", "UnaryExpression", "PowerExpression", "PrimaryExpression", "ArgumentListOrComprehension", "ArgumentList", "PositionalArgument", "AssignmentExpression"},
		},
		{ // 145
			"a(*b)",
			func(f *python.File) python.Type {
				return f.Statements[0].StatementList.Statements[0].AssignmentStatement.StarredExpression.Expression.ConditionalExpression.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpression.AndExpression.ShiftExpression.AddExpression.MultiplyExpression.UnaryExpression.PowerExpression.PrimaryExpression.Call.ArgumentList.PositionalArguments[0].Expression
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "AssignmentStatement", "StarredExpression", "Expression", "ConditionalExpression", "OrTest", "AndTest", "NotTest", "Comparison", "OrExpression", "XorExpression", "AndExpression", "ShiftExpression", "AddExpression", "MultiplyExpression", "UnaryExpression", "PowerExpression", "PrimaryExpression", "ArgumentListOrComprehension", "ArgumentList", "PositionalArgument", "Expression"},
		},
		{ // 146
			"a(b=c)",
			func(f *python.File) python.Type {
				return f.Statements[0].StatementList.Statements[0].AssignmentStatement.StarredExpression.Expression.ConditionalExpression.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpression.AndExpression.ShiftExpression.AddExpression.MultiplyExpression.UnaryExpression.PowerExpression.PrimaryExpression.Call.ArgumentList.StarredAndKeywordArguments[0].KeywordItem
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "AssignmentStatement", "StarredExpression", "Expression", "ConditionalExpression", "OrTest", "AndTest", "NotTest", "Comparison", "OrExpression", "XorExpression", "AndExpression", "ShiftExpression", "AddExpression", "MultiplyExpression", "UnaryExpression", "PowerExpression", "PrimaryExpression", "ArgumentListOrComprehension", "ArgumentList", "StarredOrKeyword", "KeywordItem"},
		},
		{ // 147
			"a(b=c, *d)",
			func(f *python.File) python.Type {
				return f.Statements[0].StatementList.Statements[0].AssignmentStatement.StarredExpression.Expression.ConditionalExpression.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpression.AndExpression.ShiftExpression.AddExpression.MultiplyExpression.UnaryExpression.PowerExpression.PrimaryExpression.Call.ArgumentList.StarredAndKeywordArguments[1].Expression
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "AssignmentStatement", "StarredExpression", "Expression", "ConditionalExpression", "OrTest", "AndTest", "NotTest", "Comparison", "OrExpression", "XorExpression", "AndExpression", "ShiftExpression", "AddExpression", "MultiplyExpression", "UnaryExpression", "PowerExpression", "PrimaryExpression", "ArgumentListOrComprehension", "ArgumentList", "StarredOrKeyword", "Expression"},
		},
		{ // 148
			"a(b=c)",
			func(f *python.File) python.Type {
				return &f.Statements[0].StatementList.Statements[0].AssignmentStatement.StarredExpression.Expression.ConditionalExpression.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpression.AndExpression.ShiftExpression.AddExpression.MultiplyExpression.UnaryExpression.PowerExpression.PrimaryExpression.Call.ArgumentList.StarredAndKeywordArguments[0].KeywordItem.Expression
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "AssignmentStatement", "StarredExpression", "Expression", "ConditionalExpression", "OrTest", "AndTest", "NotTest", "Comparison", "OrExpression", "XorExpression", "AndExpression", "ShiftExpression", "AddExpression", "MultiplyExpression", "UnaryExpression", "PowerExpression", "PrimaryExpression", "ArgumentListOrComprehension", "ArgumentList", "StarredOrKeyword", "KeywordItem", "Expression"},
		},
		{ // 149
			"a(**b, c=d)",
			func(f *python.File) python.Type {
				return f.Statements[0].StatementList.Statements[0].AssignmentStatement.StarredExpression.Expression.ConditionalExpression.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpression.AndExpression.ShiftExpression.AddExpression.MultiplyExpression.UnaryExpression.PowerExpression.PrimaryExpression.Call.ArgumentList.KeywordArguments[1].KeywordItem
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "AssignmentStatement", "StarredExpression", "Expression", "ConditionalExpression", "OrTest", "AndTest", "NotTest", "Comparison", "OrExpression", "XorExpression", "AndExpression", "ShiftExpression", "AddExpression", "MultiplyExpression", "UnaryExpression", "PowerExpression", "PrimaryExpression", "ArgumentListOrComprehension", "ArgumentList", "KeywordArgument", "KeywordItem"},
		},
		{ // 150
			"a(**b, c=d)",
			func(f *python.File) python.Type {
				return f.Statements[0].StatementList.Statements[0].AssignmentStatement.StarredExpression.Expression.ConditionalExpression.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpression.AndExpression.ShiftExpression.AddExpression.MultiplyExpression.UnaryExpression.PowerExpression.PrimaryExpression.Call.ArgumentList.KeywordArguments[0].Expression
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "AssignmentStatement", "StarredExpression", "Expression", "ConditionalExpression", "OrTest", "AndTest", "NotTest", "Comparison", "OrExpression", "XorExpression", "AndExpression", "ShiftExpression", "AddExpression", "MultiplyExpression", "UnaryExpression", "PowerExpression", "PrimaryExpression", "ArgumentListOrComprehension", "ArgumentList", "KeywordArgument", "Expression"},
		},
		{ // 151
			"a[b,c]",
			nilRet,
			nil,
		},
		{ // 152
			"a[b,c]",
			func(f *python.File) python.Type {
				return &f.Statements[0].StatementList.Statements[0].AssignmentStatement.StarredExpression.Expression.ConditionalExpression.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpression.AndExpression.ShiftExpression.AddExpression.MultiplyExpression.UnaryExpression.PowerExpression.PrimaryExpression.Slicing.SliceItems[0]
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "AssignmentStatement", "StarredExpression", "Expression", "ConditionalExpression", "OrTest", "AndTest", "NotTest", "Comparison", "OrExpression", "XorExpression", "AndExpression", "ShiftExpression", "AddExpression", "MultiplyExpression", "UnaryExpression", "PowerExpression", "PrimaryExpression", "SliceList", "SliceItem"},
		},
		{ // 153
			"a[b,c]",
			func(f *python.File) python.Type {
				return &f.Statements[0].StatementList.Statements[0].AssignmentStatement.StarredExpression.Expression.ConditionalExpression.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpression.AndExpression.ShiftExpression.AddExpression.MultiplyExpression.UnaryExpression.PowerExpression.PrimaryExpression.Slicing.SliceItems[1]
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "AssignmentStatement", "StarredExpression", "Expression", "ConditionalExpression", "OrTest", "AndTest", "NotTest", "Comparison", "OrExpression", "XorExpression", "AndExpression", "ShiftExpression", "AddExpression", "MultiplyExpression", "UnaryExpression", "PowerExpression", "PrimaryExpression", "SliceList", "SliceItem"},
		},
		{ // 154
			"a[b]",
			func(f *python.File) python.Type {
				return f.Statements[0].StatementList.Statements[0].AssignmentStatement.StarredExpression.Expression.ConditionalExpression.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpression.AndExpression.ShiftExpression.AddExpression.MultiplyExpression.UnaryExpression.PowerExpression.PrimaryExpression.Slicing.SliceItems[0].Expression
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "AssignmentStatement", "StarredExpression", "Expression", "ConditionalExpression", "OrTest", "AndTest", "NotTest", "Comparison", "OrExpression", "XorExpression", "AndExpression", "ShiftExpression", "AddExpression", "MultiplyExpression", "UnaryExpression", "PowerExpression", "PrimaryExpression", "SliceList", "SliceItem", "Expression"},
		},
		{ // 155
			"a[b:c:d]",
			func(f *python.File) python.Type {
				return f.Statements[0].StatementList.Statements[0].AssignmentStatement.StarredExpression.Expression.ConditionalExpression.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpression.AndExpression.ShiftExpression.AddExpression.MultiplyExpression.UnaryExpression.PowerExpression.PrimaryExpression.Slicing.SliceItems[0].LowerBound
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "AssignmentStatement", "StarredExpression", "Expression", "ConditionalExpression", "OrTest", "AndTest", "NotTest", "Comparison", "OrExpression", "XorExpression", "AndExpression", "ShiftExpression", "AddExpression", "MultiplyExpression", "UnaryExpression", "PowerExpression", "PrimaryExpression", "SliceList", "SliceItem", "Expression"},
		},
		{ // 156
			"a[b:c:d]",
			func(f *python.File) python.Type {
				return f.Statements[0].StatementList.Statements[0].AssignmentStatement.StarredExpression.Expression.ConditionalExpression.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpression.AndExpression.ShiftExpression.AddExpression.MultiplyExpression.UnaryExpression.PowerExpression.PrimaryExpression.Slicing.SliceItems[0].UpperBound
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "AssignmentStatement", "StarredExpression", "Expression", "ConditionalExpression", "OrTest", "AndTest", "NotTest", "Comparison", "OrExpression", "XorExpression", "AndExpression", "ShiftExpression", "AddExpression", "MultiplyExpression", "UnaryExpression", "PowerExpression", "PrimaryExpression", "SliceList", "SliceItem", "Expression"},
		},
		{ // 157
			"a[b:c:d]",
			func(f *python.File) python.Type {
				return f.Statements[0].StatementList.Statements[0].AssignmentStatement.StarredExpression.Expression.ConditionalExpression.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpression.AndExpression.ShiftExpression.AddExpression.MultiplyExpression.UnaryExpression.PowerExpression.PrimaryExpression.Slicing.SliceItems[0].Stride
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "AssignmentStatement", "StarredExpression", "Expression", "ConditionalExpression", "OrTest", "AndTest", "NotTest", "Comparison", "OrExpression", "XorExpression", "AndExpression", "ShiftExpression", "AddExpression", "MultiplyExpression", "UnaryExpression", "PowerExpression", "PrimaryExpression", "SliceList", "SliceItem", "Expression"},
		},
		{ // 158
			"a[b:c]",
			func(f *python.File) python.Type {
				return f.Statements[0].StatementList.Statements[0].AssignmentStatement.StarredExpression.Expression.ConditionalExpression.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpression.AndExpression.ShiftExpression.AddExpression.MultiplyExpression.UnaryExpression.PowerExpression.PrimaryExpression.Slicing.SliceItems[0].Stride
			},
			nil,
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

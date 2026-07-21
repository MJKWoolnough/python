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
				return &f.Statements[0].CompoundStatement.Try.Except[0].Expression[0]
			},
			[]string{"File", "Statement", "CompoundStatement", "TryStatement", "Except", "Expression"},
		},
		{ // 44
			"try:\n\ta\nexcept b, c as d: e",
			func(f *python.File) python.Type {
				return &f.Statements[0].CompoundStatement.Try.Except[0].Expression[1]
			},
			[]string{"File", "Statement", "CompoundStatement", "TryStatement", "Except", "Expression"},
		},
		{ // 45
			"try:\n\ta\nexcept b as c: d",
			func(f *python.File) python.Type {
				return &f.Statements[0].CompoundStatement.Try.Except[0].Suite
			},
			[]string{"File", "Statement", "CompoundStatement", "TryStatement", "Except", "Suite"},
		},
		{ // 46
			"with a:\n\tb",
			func(f *python.File) python.Type {
				return &f.Statements[0].CompoundStatement.With.Contents
			},
			[]string{"File", "Statement", "CompoundStatement", "WithStatement", "WithStatementContents"},
		},
		{ // 47
			"with a:\n\tb",
			func(f *python.File) python.Type {
				return &f.Statements[0].CompoundStatement.With.Suite
			},
			[]string{"File", "Statement", "CompoundStatement", "WithStatement", "Suite"},
		},
		{ // 48
			"with a, b:\n\tc",
			func(f *python.File) python.Type {
				return &f.Statements[0].CompoundStatement.With.Contents.Items[0]
			},
			[]string{"File", "Statement", "CompoundStatement", "WithStatement", "WithStatementContents", "WithItem"},
		},
		{ // 49
			"with a, b:\n\tc",
			func(f *python.File) python.Type {
				return &f.Statements[0].CompoundStatement.With.Contents.Items[1]
			},
			[]string{"File", "Statement", "CompoundStatement", "WithStatement", "WithStatementContents", "WithItem"},
		},
		{ // 50
			"with a as b:\n\tc",
			func(f *python.File) python.Type {
				return &f.Statements[0].CompoundStatement.With.Contents.Items[0].Expression
			},
			[]string{"File", "Statement", "CompoundStatement", "WithStatement", "WithStatementContents", "WithItem", "Expression"},
		},
		{ // 51
			"with a as b:\n\tc",
			func(f *python.File) python.Type {
				return f.Statements[0].CompoundStatement.With.Contents.Items[0].Target
			},
			[]string{"File", "Statement", "CompoundStatement", "WithStatement", "WithStatementContents", "WithItem", "Target"},
		},
		{ // 52
			"@a\ndef b[c]() -> d: e",
			func(f *python.File) python.Type {
				return f.Statements[0].CompoundStatement.Func.Decorators
			},
			[]string{"File", "Statement", "CompoundStatement", "FuncDefinition", "Decorators"},
		},
		{ // 53
			"@a\ndef b[c]() -> d: e",
			func(f *python.File) python.Type {
				return f.Statements[0].CompoundStatement.Func.TypeParams
			},
			[]string{"File", "Statement", "CompoundStatement", "FuncDefinition", "TypeParams"},
		},
		{ // 54
			"@a\ndef b[c]() -> d: e",
			func(f *python.File) python.Type {
				return &f.Statements[0].CompoundStatement.Func.ParameterList
			},
			[]string{"File", "Statement", "CompoundStatement", "FuncDefinition", "ParameterList"},
		},
		{ // 55
			"@a\ndef b[c]() -> d: e",
			func(f *python.File) python.Type {
				return f.Statements[0].CompoundStatement.Func.Expression
			},
			[]string{"File", "Statement", "CompoundStatement", "FuncDefinition", "Expression"},
		},
		{ // 56
			"@a\ndef b[c]() -> d: e",
			func(f *python.File) python.Type {
				return &f.Statements[0].CompoundStatement.Func.Suite
			},
			[]string{"File", "Statement", "CompoundStatement", "FuncDefinition", "Suite"},
		},
		{ // 57
			"@a\n@b\ndef c(): d",
			func(f *python.File) python.Type {
				return &f.Statements[0].CompoundStatement.Func.Decorators.Decorators[0]
			},
			[]string{"File", "Statement", "CompoundStatement", "FuncDefinition", "Decorators", "Decorator"},
		},
		{ // 58
			"@a\n@b\ndef c(): d",
			func(f *python.File) python.Type {
				return &f.Statements[0].CompoundStatement.Func.Decorators.Decorators[1]
			},
			[]string{"File", "Statement", "CompoundStatement", "FuncDefinition", "Decorators", "Decorator"},
		},
		{ // 59
			"@a\ndef b(): c",
			func(f *python.File) python.Type {
				return &f.Statements[0].CompoundStatement.Func.Decorators.Decorators[0].Decorator
			},
			[]string{"File", "Statement", "CompoundStatement", "FuncDefinition", "Decorators", "Decorator", "AssignmentExpression"},
		},
		{ // 60
			"def a[b, c]() -> d: e",
			func(f *python.File) python.Type {
				return &f.Statements[0].CompoundStatement.Func.TypeParams.TypeParams[0]
			},
			[]string{"File", "Statement", "CompoundStatement", "FuncDefinition", "TypeParams", "TypeParam"},
		},
		{ // 61
			"def a[b, c]() -> d: e",
			func(f *python.File) python.Type {
				return &f.Statements[0].CompoundStatement.Func.TypeParams.TypeParams[1]
			},
			[]string{"File", "Statement", "CompoundStatement", "FuncDefinition", "TypeParams", "TypeParam"},
		},
		{ // 62
			"def a[b: c]() -> d: e",
			func(f *python.File) python.Type {
				return f.Statements[0].CompoundStatement.Func.TypeParams.TypeParams[0].Expression
			},
			[]string{"File", "Statement", "CompoundStatement", "FuncDefinition", "TypeParams", "TypeParam", "Expression"},
		},
		{ // 63
			"def a(b, c, /, d, e, *f, g, h, **i): j",
			func(f *python.File) python.Type {
				return &f.Statements[0].CompoundStatement.Func.ParameterList.DefParameters[0]
			},
			[]string{"File", "Statement", "CompoundStatement", "FuncDefinition", "ParameterList", "DefParameter"},
		},
		{ // 64
			"def a(b, c, /, d, e, *f, g, h, **i): j",
			func(f *python.File) python.Type {
				return &f.Statements[0].CompoundStatement.Func.ParameterList.DefParameters[1]
			},
			[]string{"File", "Statement", "CompoundStatement", "FuncDefinition", "ParameterList", "DefParameter"},
		},
		{ // 65
			"def a(b, c, /, d, e, *f, g, h, **i): j",
			func(f *python.File) python.Type {
				return &f.Statements[0].CompoundStatement.Func.ParameterList.NoPosOnly[0]
			},
			[]string{"File", "Statement", "CompoundStatement", "FuncDefinition", "ParameterList", "DefParameter"},
		},
		{ // 66
			"def a(b, c, /, d, e, *f, g, h, **i): j",
			func(f *python.File) python.Type {
				return &f.Statements[0].CompoundStatement.Func.ParameterList.NoPosOnly[1]
			},
			[]string{"File", "Statement", "CompoundStatement", "FuncDefinition", "ParameterList", "DefParameter"},
		},
		{ // 67
			"def a(b, c, /, d, e, *f, g, h, **i): j",
			func(f *python.File) python.Type {
				return f.Statements[0].CompoundStatement.Func.ParameterList.StarArg
			},
			[]string{"File", "Statement", "CompoundStatement", "FuncDefinition", "ParameterList", "Parameter"},
		},
		{ // 68
			"def a(b, c, /, d, e, *f, g, h, **i): j",
			func(f *python.File) python.Type {
				return &f.Statements[0].CompoundStatement.Func.ParameterList.StarArgs[0]
			},
			[]string{"File", "Statement", "CompoundStatement", "FuncDefinition", "ParameterList", "DefParameter"},
		},
		{ // 69
			"def a(b, c, /, d, e, *f, g, h, **i): j",
			func(f *python.File) python.Type {
				return &f.Statements[0].CompoundStatement.Func.ParameterList.StarArgs[1]
			},
			[]string{"File", "Statement", "CompoundStatement", "FuncDefinition", "ParameterList", "DefParameter"},
		},
		{ // 70
			"def a(b, c, /, d, e, *f, g, h, **i): j",
			func(f *python.File) python.Type {
				return f.Statements[0].CompoundStatement.Func.ParameterList.StarStarArg
			},
			[]string{"File", "Statement", "CompoundStatement", "FuncDefinition", "ParameterList", "Parameter"},
		},
		{ // 71
			"def a(b = c): d",
			func(f *python.File) python.Type {
				return &f.Statements[0].CompoundStatement.Func.ParameterList.NoPosOnly[0].Parameter
			},
			[]string{"File", "Statement", "CompoundStatement", "FuncDefinition", "ParameterList", "DefParameter", "Parameter"},
		},
		{ // 72
			"def a(b = c): d",
			func(f *python.File) python.Type {
				return f.Statements[0].CompoundStatement.Func.ParameterList.NoPosOnly[0].Value
			},
			[]string{"File", "Statement", "CompoundStatement", "FuncDefinition", "ParameterList", "DefParameter", "Expression"},
		},
		{ // 73
			"def a(b: c): d",
			func(f *python.File) python.Type {
				return f.Statements[0].CompoundStatement.Func.ParameterList.NoPosOnly[0].Parameter.Type
			},
			[]string{"File", "Statement", "CompoundStatement", "FuncDefinition", "ParameterList", "DefParameter", "Parameter", "Expression"},
		},
		{ // 74
			"@a\nclass b[c](d): e",
			func(f *python.File) python.Type {
				return f.Statements[0].CompoundStatement.Class.Decorators
			},
			[]string{"File", "Statement", "CompoundStatement", "ClassDefinition", "Decorators"},
		},
		{ // 75
			"@a\nclass b[c](d): e",
			func(f *python.File) python.Type {
				return f.Statements[0].CompoundStatement.Class.TypeParams
			},
			[]string{"File", "Statement", "CompoundStatement", "ClassDefinition", "TypeParams"},
		},
		{ // 76
			"@a\nclass b[c](d): e",
			func(f *python.File) python.Type {
				return f.Statements[0].CompoundStatement.Class.Inheritance
			},
			[]string{"File", "Statement", "CompoundStatement", "ClassDefinition", "ArgumentList"},
		},
		{ // 77
			"@a\nclass b[c](d): e",
			func(f *python.File) python.Type {
				return &f.Statements[0].CompoundStatement.Class.Suite
			},
			[]string{"File", "Statement", "CompoundStatement", "ClassDefinition", "Suite"},
		},
		{ // 78
			"a; b",
			func(f *python.File) python.Type {
				return &f.Statements[0].StatementList.Statements[0]
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement"},
		},
		{ // 79
			"a; b",
			func(f *python.File) python.Type {
				return &f.Statements[0].StatementList.Statements[1]
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement"},
		},
		{ // 80
			"a",
			func(f *python.File) python.Type {
				return &f.Statements[0].StatementList.Statements[0]
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement"},
		},
		{ // 81
			"assert a",
			func(f *python.File) python.Type {
				return f.Statements[0].StatementList.Statements[0].AssertStatement
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "AssertStatement"},
		},
		{ // 82
			"a = b",
			func(f *python.File) python.Type {
				return f.Statements[0].StatementList.Statements[0].AssignmentStatement
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "AssignmentStatement"},
		},
		{ // 83
			"a += b",
			func(f *python.File) python.Type {
				return f.Statements[0].StatementList.Statements[0].AugmentedAssignmentStatement
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "AugmentedAssignmentStatement"},
		},
		{ // 84
			"a: b = c",
			func(f *python.File) python.Type {
				return f.Statements[0].StatementList.Statements[0].AnnotatedAssignmentStatement
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "AnnotatedAssignmentStatement"},
		},
		{ // 85
			"del a",
			func(f *python.File) python.Type {
				return f.Statements[0].StatementList.Statements[0].DelStatement
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "DelStatement"},
		},
		{ // 86
			"def a(): return b",
			func(f *python.File) python.Type {
				return f.Statements[0].CompoundStatement.Func.Suite.StatementList.Statements[0].ReturnStatement
			},
			[]string{"File", "Statement", "CompoundStatement", "FuncDefinition", "Suite", "StatementList", "SimpleStatement", "ReturnStatement"},
		},
		{ // 87
			"def a(): yield b",
			func(f *python.File) python.Type {
				return f.Statements[0].CompoundStatement.Func.Suite.StatementList.Statements[0].YieldStatement
			},
			[]string{"File", "Statement", "CompoundStatement", "FuncDefinition", "Suite", "StatementList", "SimpleStatement", "YieldExpression"},
		},
		{ // 88
			"raise a",
			func(f *python.File) python.Type {
				return f.Statements[0].StatementList.Statements[0].RaiseStatement
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "RaiseStatement"},
		},
		{ // 89
			"import a",
			func(f *python.File) python.Type {
				return f.Statements[0].StatementList.Statements[0].ImportStatement
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "ImportStatement"},
		},
		{ // 90
			"global a",
			func(f *python.File) python.Type {
				return f.Statements[0].StatementList.Statements[0].GlobalStatement
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "GlobalStatement"},
		},
		{ // 91
			"nonlocal a",
			func(f *python.File) python.Type {
				return f.Statements[0].StatementList.Statements[0].NonLocalStatement
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "NonLocalStatement"},
		},
		{ // 92
			"type a = b",
			func(f *python.File) python.Type {
				return f.Statements[0].StatementList.Statements[0].TypeStatement
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "TypeStatement"},
		},
		{ // 93
			"assert a, b",
			nilRet,
			nil,
		},
		{ // 94
			"assert a, b",
			func(f *python.File) python.Type {
				return &f.Statements[0].StatementList.Statements[0].AssertStatement.Expressions[0]
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "AssertStatement", "Expression"},
		},
		{ // 95
			"assert a, b",
			func(f *python.File) python.Type {
				return &f.Statements[0].StatementList.Statements[0].AssertStatement.Expressions[1]
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "AssertStatement", "Expression"},
		},
		{ // 96
			"a = b = c",
			func(f *python.File) python.Type {
				return &f.Statements[0].StatementList.Statements[0].AssignmentStatement.TargetLists[0]
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "AssignmentStatement", "TargetList"},
		},
		{ // 97
			"a = b = c",
			func(f *python.File) python.Type {
				return &f.Statements[0].StatementList.Statements[0].AssignmentStatement.TargetLists[1]
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "AssignmentStatement", "TargetList"},
		},
		{ // 98
			"a = b",
			func(f *python.File) python.Type {
				return f.Statements[0].StatementList.Statements[0].AssignmentStatement.StarredExpression
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "AssignmentStatement", "StarredExpression"},
		},
		{ // 99
			"def a(): b = yield c",
			func(f *python.File) python.Type {
				return f.Statements[0].CompoundStatement.Func.Suite.StatementList.Statements[0].AssignmentStatement.YieldExpression
			},
			[]string{"File", "Statement", "CompoundStatement", "FuncDefinition", "Suite", "StatementList", "SimpleStatement", "AssignmentStatement", "YieldExpression"},
		},
		{ // 100
			"a",
			func(f *python.File) python.Type {
				return f.Statements[0].StatementList.Statements[0].AssignmentStatement.StarredExpression.Expression
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "AssignmentStatement", "StarredExpression", "Expression"},
		},
		{ // 101
			"*a",
			func(f *python.File) python.Type {
				return f.Statements[0].StatementList.Statements[0].AssignmentStatement.StarredExpression.StarredList
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "AssignmentStatement", "StarredExpression", "StarredList"},
		},
		{ // 102
			"a",
			func(f *python.File) python.Type {
				return f.Statements[0].StatementList.Statements[0].AssignmentStatement.StarredExpression.Expression.ConditionalExpression
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "AssignmentStatement", "StarredExpression", "Expression", "ConditionalExpression"},
		},
		{ // 103
			"lambda b: c",
			func(f *python.File) python.Type {
				return f.Statements[0].StatementList.Statements[0].AssignmentStatement.StarredExpression.Expression.LambdaExpression
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "AssignmentStatement", "StarredExpression", "Expression", "LambdaExpression"},
		},
		{ // 104
			"a if b else c",
			func(f *python.File) python.Type {
				return &f.Statements[0].StatementList.Statements[0].AssignmentStatement.StarredExpression.Expression.ConditionalExpression.OrTest
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "AssignmentStatement", "StarredExpression", "Expression", "ConditionalExpression", "OrTest"},
		},
		{ // 105
			"a if b else c",
			func(f *python.File) python.Type {
				return f.Statements[0].StatementList.Statements[0].AssignmentStatement.StarredExpression.Expression.ConditionalExpression.If
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "AssignmentStatement", "StarredExpression", "Expression", "ConditionalExpression", "OrTest"},
		},
		{ // 106
			"a if b else c",
			func(f *python.File) python.Type {
				return f.Statements[0].StatementList.Statements[0].AssignmentStatement.StarredExpression.Expression.ConditionalExpression.Else
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "AssignmentStatement", "StarredExpression", "Expression", "ConditionalExpression", "Expression"},
		},
		{ // 107
			"a or b",
			func(f *python.File) python.Type {
				return &f.Statements[0].StatementList.Statements[0].AssignmentStatement.StarredExpression.Expression.ConditionalExpression.OrTest.AndTest
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "AssignmentStatement", "StarredExpression", "Expression", "ConditionalExpression", "OrTest", "AndTest"},
		},
		{ // 108
			"a or b",
			func(f *python.File) python.Type {
				return f.Statements[0].StatementList.Statements[0].AssignmentStatement.StarredExpression.Expression.ConditionalExpression.OrTest.OrTest
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "AssignmentStatement", "StarredExpression", "Expression", "ConditionalExpression", "OrTest", "OrTest"},
		},
		{ // 109
			"a and b",
			func(f *python.File) python.Type {
				return &f.Statements[0].StatementList.Statements[0].AssignmentStatement.StarredExpression.Expression.ConditionalExpression.OrTest.AndTest.NotTest
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "AssignmentStatement", "StarredExpression", "Expression", "ConditionalExpression", "OrTest", "AndTest", "NotTest"},
		},
		{ // 110
			"a and b",
			func(f *python.File) python.Type {
				return f.Statements[0].StatementList.Statements[0].AssignmentStatement.StarredExpression.Expression.ConditionalExpression.OrTest.AndTest.AndTest
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "AssignmentStatement", "StarredExpression", "Expression", "ConditionalExpression", "OrTest", "AndTest", "AndTest"},
		},
		{ // 111
			"not a",
			func(f *python.File) python.Type {
				return &f.Statements[0].StatementList.Statements[0].AssignmentStatement.StarredExpression.Expression.ConditionalExpression.OrTest.AndTest.NotTest.Comparison
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "AssignmentStatement", "StarredExpression", "Expression", "ConditionalExpression", "OrTest", "AndTest", "NotTest", "Comparison"},
		},
		{ // 112
			"a in b == c",
			func(f *python.File) python.Type {
				return &f.Statements[0].StatementList.Statements[0].AssignmentStatement.StarredExpression.Expression.ConditionalExpression.OrTest.AndTest.NotTest.Comparison.OrExpression
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "AssignmentStatement", "StarredExpression", "Expression", "ConditionalExpression", "OrTest", "AndTest", "NotTest", "Comparison", "OrExpression"},
		},
		{ // 113
			"a in b == c",
			func(f *python.File) python.Type {
				return &f.Statements[0].StatementList.Statements[0].AssignmentStatement.StarredExpression.Expression.ConditionalExpression.OrTest.AndTest.NotTest.Comparison.Comparisons[0]
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "AssignmentStatement", "StarredExpression", "Expression", "ConditionalExpression", "OrTest", "AndTest", "NotTest", "Comparison", "ComparisonExpression"},
		},
		{ // 114
			"a in b == c",
			func(f *python.File) python.Type {
				return &f.Statements[0].StatementList.Statements[0].AssignmentStatement.StarredExpression.Expression.ConditionalExpression.OrTest.AndTest.NotTest.Comparison.Comparisons[1]
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "AssignmentStatement", "StarredExpression", "Expression", "ConditionalExpression", "OrTest", "AndTest", "NotTest", "Comparison", "ComparisonExpression"},
		},
		{ // 115
			"a in b",
			func(f *python.File) python.Type {
				return &f.Statements[0].StatementList.Statements[0].AssignmentStatement.StarredExpression.Expression.ConditionalExpression.OrTest.AndTest.NotTest.Comparison.Comparisons[0].OrExpression
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "AssignmentStatement", "StarredExpression", "Expression", "ConditionalExpression", "OrTest", "AndTest", "NotTest", "Comparison", "ComparisonExpression", "OrExpression"},
		},
		{ // 116
			"a | b",
			func(f *python.File) python.Type {
				return &f.Statements[0].StatementList.Statements[0].AssignmentStatement.StarredExpression.Expression.ConditionalExpression.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpression
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "AssignmentStatement", "StarredExpression", "Expression", "ConditionalExpression", "OrTest", "AndTest", "NotTest", "Comparison", "OrExpression", "XorExpression"},
		},
		{ // 117
			"a | b",
			func(f *python.File) python.Type {
				return f.Statements[0].StatementList.Statements[0].AssignmentStatement.StarredExpression.Expression.ConditionalExpression.OrTest.AndTest.NotTest.Comparison.OrExpression.OrExpression
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "AssignmentStatement", "StarredExpression", "Expression", "ConditionalExpression", "OrTest", "AndTest", "NotTest", "Comparison", "OrExpression", "OrExpression"},
		},
		{ // 118
			"a ^ b",
			func(f *python.File) python.Type {
				return &f.Statements[0].StatementList.Statements[0].AssignmentStatement.StarredExpression.Expression.ConditionalExpression.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpression.AndExpression
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "AssignmentStatement", "StarredExpression", "Expression", "ConditionalExpression", "OrTest", "AndTest", "NotTest", "Comparison", "OrExpression", "XorExpression", "AndExpression"},
		},
		{ // 119
			"a ^ b",
			func(f *python.File) python.Type {
				return f.Statements[0].StatementList.Statements[0].AssignmentStatement.StarredExpression.Expression.ConditionalExpression.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpression.XorExpression
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "AssignmentStatement", "StarredExpression", "Expression", "ConditionalExpression", "OrTest", "AndTest", "NotTest", "Comparison", "OrExpression", "XorExpression", "XorExpression"},
		},
		{ // 120
			"a & b",
			func(f *python.File) python.Type {
				return &f.Statements[0].StatementList.Statements[0].AssignmentStatement.StarredExpression.Expression.ConditionalExpression.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpression.AndExpression.ShiftExpression
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "AssignmentStatement", "StarredExpression", "Expression", "ConditionalExpression", "OrTest", "AndTest", "NotTest", "Comparison", "OrExpression", "XorExpression", "AndExpression", "ShiftExpression"},
		},
		{ // 121
			"a & b",
			func(f *python.File) python.Type {
				return f.Statements[0].StatementList.Statements[0].AssignmentStatement.StarredExpression.Expression.ConditionalExpression.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpression.AndExpression.AndExpression
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "AssignmentStatement", "StarredExpression", "Expression", "ConditionalExpression", "OrTest", "AndTest", "NotTest", "Comparison", "OrExpression", "XorExpression", "AndExpression", "AndExpression"},
		},
		{ // 122
			"a << b",
			func(f *python.File) python.Type {
				return &f.Statements[0].StatementList.Statements[0].AssignmentStatement.StarredExpression.Expression.ConditionalExpression.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpression.AndExpression.ShiftExpression.AddExpression
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "AssignmentStatement", "StarredExpression", "Expression", "ConditionalExpression", "OrTest", "AndTest", "NotTest", "Comparison", "OrExpression", "XorExpression", "AndExpression", "ShiftExpression", "AddExpression"},
		},
		{ // 123
			"a << b",
			func(f *python.File) python.Type {
				return f.Statements[0].StatementList.Statements[0].AssignmentStatement.StarredExpression.Expression.ConditionalExpression.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpression.AndExpression.ShiftExpression.ShiftExpression
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "AssignmentStatement", "StarredExpression", "Expression", "ConditionalExpression", "OrTest", "AndTest", "NotTest", "Comparison", "OrExpression", "XorExpression", "AndExpression", "ShiftExpression", "ShiftExpression"},
		},
		{ // 124
			"a + b",
			func(f *python.File) python.Type {
				return &f.Statements[0].StatementList.Statements[0].AssignmentStatement.StarredExpression.Expression.ConditionalExpression.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpression.AndExpression.ShiftExpression.AddExpression.MultiplyExpression
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "AssignmentStatement", "StarredExpression", "Expression", "ConditionalExpression", "OrTest", "AndTest", "NotTest", "Comparison", "OrExpression", "XorExpression", "AndExpression", "ShiftExpression", "AddExpression", "MultiplyExpression"},
		},
		{ // 125
			"a + b",
			func(f *python.File) python.Type {
				return f.Statements[0].StatementList.Statements[0].AssignmentStatement.StarredExpression.Expression.ConditionalExpression.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpression.AndExpression.ShiftExpression.AddExpression.AddExpression
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "AssignmentStatement", "StarredExpression", "Expression", "ConditionalExpression", "OrTest", "AndTest", "NotTest", "Comparison", "OrExpression", "XorExpression", "AndExpression", "ShiftExpression", "AddExpression", "AddExpression"},
		},
		{ // 126
			"a * b",
			func(f *python.File) python.Type {
				return &f.Statements[0].StatementList.Statements[0].AssignmentStatement.StarredExpression.Expression.ConditionalExpression.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpression.AndExpression.ShiftExpression.AddExpression.MultiplyExpression.UnaryExpression
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "AssignmentStatement", "StarredExpression", "Expression", "ConditionalExpression", "OrTest", "AndTest", "NotTest", "Comparison", "OrExpression", "XorExpression", "AndExpression", "ShiftExpression", "AddExpression", "MultiplyExpression", "UnaryExpression"},
		},
		{ // 127
			"a * b",
			func(f *python.File) python.Type {
				return f.Statements[0].StatementList.Statements[0].AssignmentStatement.StarredExpression.Expression.ConditionalExpression.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpression.AndExpression.ShiftExpression.AddExpression.MultiplyExpression.MultiplyExpression
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "AssignmentStatement", "StarredExpression", "Expression", "ConditionalExpression", "OrTest", "AndTest", "NotTest", "Comparison", "OrExpression", "XorExpression", "AndExpression", "ShiftExpression", "AddExpression", "MultiplyExpression", "MultiplyExpression"},
		},
		{ // 128
			"a",
			func(f *python.File) python.Type {
				return f.Statements[0].StatementList.Statements[0].AssignmentStatement.StarredExpression.Expression.ConditionalExpression.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpression.AndExpression.ShiftExpression.AddExpression.MultiplyExpression.UnaryExpression.PowerExpression
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "AssignmentStatement", "StarredExpression", "Expression", "ConditionalExpression", "OrTest", "AndTest", "NotTest", "Comparison", "OrExpression", "XorExpression", "AndExpression", "ShiftExpression", "AddExpression", "MultiplyExpression", "UnaryExpression", "PowerExpression"},
		},
		{ // 129
			"+a",
			func(f *python.File) python.Type {
				return f.Statements[0].StatementList.Statements[0].AssignmentStatement.StarredExpression.Expression.ConditionalExpression.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpression.AndExpression.ShiftExpression.AddExpression.MultiplyExpression.UnaryExpression.UnaryExpression
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "AssignmentStatement", "StarredExpression", "Expression", "ConditionalExpression", "OrTest", "AndTest", "NotTest", "Comparison", "OrExpression", "XorExpression", "AndExpression", "ShiftExpression", "AddExpression", "MultiplyExpression", "UnaryExpression", "UnaryExpression"},
		},
		{ // 130
			"a ** b",
			func(f *python.File) python.Type {
				return &f.Statements[0].StatementList.Statements[0].AssignmentStatement.StarredExpression.Expression.ConditionalExpression.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpression.AndExpression.ShiftExpression.AddExpression.MultiplyExpression.UnaryExpression.PowerExpression.PrimaryExpression
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "AssignmentStatement", "StarredExpression", "Expression", "ConditionalExpression", "OrTest", "AndTest", "NotTest", "Comparison", "OrExpression", "XorExpression", "AndExpression", "ShiftExpression", "AddExpression", "MultiplyExpression", "UnaryExpression", "PowerExpression", "PrimaryExpression"},
		},
		{ // 131
			"a ** b",
			func(f *python.File) python.Type {
				return f.Statements[0].StatementList.Statements[0].AssignmentStatement.StarredExpression.Expression.ConditionalExpression.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpression.AndExpression.ShiftExpression.AddExpression.MultiplyExpression.UnaryExpression.PowerExpression.UnaryExpression
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "AssignmentStatement", "StarredExpression", "Expression", "ConditionalExpression", "OrTest", "AndTest", "NotTest", "Comparison", "OrExpression", "XorExpression", "AndExpression", "ShiftExpression", "AddExpression", "MultiplyExpression", "UnaryExpression", "PowerExpression", "UnaryExpression"},
		},
		{ // 132
			"a",
			func(f *python.File) python.Type {
				return f.Statements[0].StatementList.Statements[0].AssignmentStatement.StarredExpression.Expression.ConditionalExpression.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpression.AndExpression.ShiftExpression.AddExpression.MultiplyExpression.UnaryExpression.PowerExpression.PrimaryExpression.Atom
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "AssignmentStatement", "StarredExpression", "Expression", "ConditionalExpression", "OrTest", "AndTest", "NotTest", "Comparison", "OrExpression", "XorExpression", "AndExpression", "ShiftExpression", "AddExpression", "MultiplyExpression", "UnaryExpression", "PowerExpression", "PrimaryExpression", "Atom"},
		},
		{ // 133
			"a()",
			func(f *python.File) python.Type {
				return f.Statements[0].StatementList.Statements[0].AssignmentStatement.StarredExpression.Expression.ConditionalExpression.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpression.AndExpression.ShiftExpression.AddExpression.MultiplyExpression.UnaryExpression.PowerExpression.PrimaryExpression.PrimaryExpression
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "AssignmentStatement", "StarredExpression", "Expression", "ConditionalExpression", "OrTest", "AndTest", "NotTest", "Comparison", "OrExpression", "XorExpression", "AndExpression", "ShiftExpression", "AddExpression", "MultiplyExpression", "UnaryExpression", "PowerExpression", "PrimaryExpression", "PrimaryExpression"},
		},
		{ // 134
			"a()",
			func(f *python.File) python.Type {
				return f.Statements[0].StatementList.Statements[0].AssignmentStatement.StarredExpression.Expression.ConditionalExpression.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpression.AndExpression.ShiftExpression.AddExpression.MultiplyExpression.UnaryExpression.PowerExpression.PrimaryExpression.Call
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "AssignmentStatement", "StarredExpression", "Expression", "ConditionalExpression", "OrTest", "AndTest", "NotTest", "Comparison", "OrExpression", "XorExpression", "AndExpression", "ShiftExpression", "AddExpression", "MultiplyExpression", "UnaryExpression", "PowerExpression", "PrimaryExpression", "ArgumentListOrComprehension"},
		},
		{ // 135
			"a[]",
			func(f *python.File) python.Type {
				return f.Statements[0].StatementList.Statements[0].AssignmentStatement.StarredExpression.Expression.ConditionalExpression.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpression.AndExpression.ShiftExpression.AddExpression.MultiplyExpression.UnaryExpression.PowerExpression.PrimaryExpression.PrimaryExpression
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "AssignmentStatement", "StarredExpression", "Expression", "ConditionalExpression", "OrTest", "AndTest", "NotTest", "Comparison", "OrExpression", "XorExpression", "AndExpression", "ShiftExpression", "AddExpression", "MultiplyExpression", "UnaryExpression", "PowerExpression", "PrimaryExpression", "PrimaryExpression"},
		},
		{ // 136
			"a[]",
			func(f *python.File) python.Type {
				return f.Statements[0].StatementList.Statements[0].AssignmentStatement.StarredExpression.Expression.ConditionalExpression.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpression.AndExpression.ShiftExpression.AddExpression.MultiplyExpression.UnaryExpression.PowerExpression.PrimaryExpression.Slicing
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "AssignmentStatement", "StarredExpression", "Expression", "ConditionalExpression", "OrTest", "AndTest", "NotTest", "Comparison", "OrExpression", "XorExpression", "AndExpression", "ShiftExpression", "AddExpression", "MultiplyExpression", "UnaryExpression", "PowerExpression", "PrimaryExpression", "SliceList"},
		},
		{ // 137
			"a(b)",
			func(f *python.File) python.Type {
				return f.Statements[0].StatementList.Statements[0].AssignmentStatement.StarredExpression.Expression.ConditionalExpression.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpression.AndExpression.ShiftExpression.AddExpression.MultiplyExpression.UnaryExpression.PowerExpression.PrimaryExpression.Call.ArgumentList
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "AssignmentStatement", "StarredExpression", "Expression", "ConditionalExpression", "OrTest", "AndTest", "NotTest", "Comparison", "OrExpression", "XorExpression", "AndExpression", "ShiftExpression", "AddExpression", "MultiplyExpression", "UnaryExpression", "PowerExpression", "PrimaryExpression", "ArgumentListOrComprehension", "ArgumentList"},
		},
		{ // 138
			"a(b for c in d)",
			func(f *python.File) python.Type {
				return f.Statements[0].StatementList.Statements[0].AssignmentStatement.StarredExpression.Expression.ConditionalExpression.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpression.AndExpression.ShiftExpression.AddExpression.MultiplyExpression.UnaryExpression.PowerExpression.PrimaryExpression.Call.Comprehension
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "AssignmentStatement", "StarredExpression", "Expression", "ConditionalExpression", "OrTest", "AndTest", "NotTest", "Comparison", "OrExpression", "XorExpression", "AndExpression", "ShiftExpression", "AddExpression", "MultiplyExpression", "UnaryExpression", "PowerExpression", "PrimaryExpression", "ArgumentListOrComprehension", "Comprehension"},
		},
		{ // 139
			"a(b, c, d=e, f=g, **h, **i)",
			func(f *python.File) python.Type {
				return &f.Statements[0].StatementList.Statements[0].AssignmentStatement.StarredExpression.Expression.ConditionalExpression.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpression.AndExpression.ShiftExpression.AddExpression.MultiplyExpression.UnaryExpression.PowerExpression.PrimaryExpression.Call.ArgumentList.PositionalArguments[0]
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "AssignmentStatement", "StarredExpression", "Expression", "ConditionalExpression", "OrTest", "AndTest", "NotTest", "Comparison", "OrExpression", "XorExpression", "AndExpression", "ShiftExpression", "AddExpression", "MultiplyExpression", "UnaryExpression", "PowerExpression", "PrimaryExpression", "ArgumentListOrComprehension", "ArgumentList", "PositionalArgument"},
		},
		{ // 140
			"a(b, c, d=e, f=g, **h, **i)",
			func(f *python.File) python.Type {
				return &f.Statements[0].StatementList.Statements[0].AssignmentStatement.StarredExpression.Expression.ConditionalExpression.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpression.AndExpression.ShiftExpression.AddExpression.MultiplyExpression.UnaryExpression.PowerExpression.PrimaryExpression.Call.ArgumentList.PositionalArguments[1]
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "AssignmentStatement", "StarredExpression", "Expression", "ConditionalExpression", "OrTest", "AndTest", "NotTest", "Comparison", "OrExpression", "XorExpression", "AndExpression", "ShiftExpression", "AddExpression", "MultiplyExpression", "UnaryExpression", "PowerExpression", "PrimaryExpression", "ArgumentListOrComprehension", "ArgumentList", "PositionalArgument"},
		},
		{ // 141
			"a(b, c, d=e, f=g, **h, **i)",
			func(f *python.File) python.Type {
				return &f.Statements[0].StatementList.Statements[0].AssignmentStatement.StarredExpression.Expression.ConditionalExpression.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpression.AndExpression.ShiftExpression.AddExpression.MultiplyExpression.UnaryExpression.PowerExpression.PrimaryExpression.Call.ArgumentList.StarredAndKeywordArguments[0]
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "AssignmentStatement", "StarredExpression", "Expression", "ConditionalExpression", "OrTest", "AndTest", "NotTest", "Comparison", "OrExpression", "XorExpression", "AndExpression", "ShiftExpression", "AddExpression", "MultiplyExpression", "UnaryExpression", "PowerExpression", "PrimaryExpression", "ArgumentListOrComprehension", "ArgumentList", "StarredOrKeyword"},
		},
		{ // 142
			"a(b, c, d=e, f=g, **h, **i)",
			func(f *python.File) python.Type {
				return &f.Statements[0].StatementList.Statements[0].AssignmentStatement.StarredExpression.Expression.ConditionalExpression.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpression.AndExpression.ShiftExpression.AddExpression.MultiplyExpression.UnaryExpression.PowerExpression.PrimaryExpression.Call.ArgumentList.StarredAndKeywordArguments[1]
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "AssignmentStatement", "StarredExpression", "Expression", "ConditionalExpression", "OrTest", "AndTest", "NotTest", "Comparison", "OrExpression", "XorExpression", "AndExpression", "ShiftExpression", "AddExpression", "MultiplyExpression", "UnaryExpression", "PowerExpression", "PrimaryExpression", "ArgumentListOrComprehension", "ArgumentList", "StarredOrKeyword"},
		},
		{ // 143
			"a(b, c, d=e, f=g, **h, **i)",
			func(f *python.File) python.Type {
				return &f.Statements[0].StatementList.Statements[0].AssignmentStatement.StarredExpression.Expression.ConditionalExpression.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpression.AndExpression.ShiftExpression.AddExpression.MultiplyExpression.UnaryExpression.PowerExpression.PrimaryExpression.Call.ArgumentList.KeywordArguments[0]
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "AssignmentStatement", "StarredExpression", "Expression", "ConditionalExpression", "OrTest", "AndTest", "NotTest", "Comparison", "OrExpression", "XorExpression", "AndExpression", "ShiftExpression", "AddExpression", "MultiplyExpression", "UnaryExpression", "PowerExpression", "PrimaryExpression", "ArgumentListOrComprehension", "ArgumentList", "KeywordArgument"},
		},
		{ // 144
			"a(b, c, d=e, f=g, **h, **i)",
			func(f *python.File) python.Type {
				return &f.Statements[0].StatementList.Statements[0].AssignmentStatement.StarredExpression.Expression.ConditionalExpression.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpression.AndExpression.ShiftExpression.AddExpression.MultiplyExpression.UnaryExpression.PowerExpression.PrimaryExpression.Call.ArgumentList.KeywordArguments[1]
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "AssignmentStatement", "StarredExpression", "Expression", "ConditionalExpression", "OrTest", "AndTest", "NotTest", "Comparison", "OrExpression", "XorExpression", "AndExpression", "ShiftExpression", "AddExpression", "MultiplyExpression", "UnaryExpression", "PowerExpression", "PrimaryExpression", "ArgumentListOrComprehension", "ArgumentList", "KeywordArgument"},
		},
		{ // 145
			"a(b)",
			func(f *python.File) python.Type {
				return f.Statements[0].StatementList.Statements[0].AssignmentStatement.StarredExpression.Expression.ConditionalExpression.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpression.AndExpression.ShiftExpression.AddExpression.MultiplyExpression.UnaryExpression.PowerExpression.PrimaryExpression.Call.ArgumentList.PositionalArguments[0].AssignmentExpression
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "AssignmentStatement", "StarredExpression", "Expression", "ConditionalExpression", "OrTest", "AndTest", "NotTest", "Comparison", "OrExpression", "XorExpression", "AndExpression", "ShiftExpression", "AddExpression", "MultiplyExpression", "UnaryExpression", "PowerExpression", "PrimaryExpression", "ArgumentListOrComprehension", "ArgumentList", "PositionalArgument", "AssignmentExpression"},
		},
		{ // 146
			"a(*b)",
			func(f *python.File) python.Type {
				return f.Statements[0].StatementList.Statements[0].AssignmentStatement.StarredExpression.Expression.ConditionalExpression.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpression.AndExpression.ShiftExpression.AddExpression.MultiplyExpression.UnaryExpression.PowerExpression.PrimaryExpression.Call.ArgumentList.PositionalArguments[0].Expression
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "AssignmentStatement", "StarredExpression", "Expression", "ConditionalExpression", "OrTest", "AndTest", "NotTest", "Comparison", "OrExpression", "XorExpression", "AndExpression", "ShiftExpression", "AddExpression", "MultiplyExpression", "UnaryExpression", "PowerExpression", "PrimaryExpression", "ArgumentListOrComprehension", "ArgumentList", "PositionalArgument", "Expression"},
		},
		{ // 147
			"a(b=c)",
			func(f *python.File) python.Type {
				return f.Statements[0].StatementList.Statements[0].AssignmentStatement.StarredExpression.Expression.ConditionalExpression.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpression.AndExpression.ShiftExpression.AddExpression.MultiplyExpression.UnaryExpression.PowerExpression.PrimaryExpression.Call.ArgumentList.StarredAndKeywordArguments[0].KeywordItem
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "AssignmentStatement", "StarredExpression", "Expression", "ConditionalExpression", "OrTest", "AndTest", "NotTest", "Comparison", "OrExpression", "XorExpression", "AndExpression", "ShiftExpression", "AddExpression", "MultiplyExpression", "UnaryExpression", "PowerExpression", "PrimaryExpression", "ArgumentListOrComprehension", "ArgumentList", "StarredOrKeyword", "KeywordItem"},
		},
		{ // 148
			"a(b=c, *d)",
			func(f *python.File) python.Type {
				return f.Statements[0].StatementList.Statements[0].AssignmentStatement.StarredExpression.Expression.ConditionalExpression.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpression.AndExpression.ShiftExpression.AddExpression.MultiplyExpression.UnaryExpression.PowerExpression.PrimaryExpression.Call.ArgumentList.StarredAndKeywordArguments[1].Expression
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "AssignmentStatement", "StarredExpression", "Expression", "ConditionalExpression", "OrTest", "AndTest", "NotTest", "Comparison", "OrExpression", "XorExpression", "AndExpression", "ShiftExpression", "AddExpression", "MultiplyExpression", "UnaryExpression", "PowerExpression", "PrimaryExpression", "ArgumentListOrComprehension", "ArgumentList", "StarredOrKeyword", "Expression"},
		},
		{ // 149
			"a(b=c)",
			func(f *python.File) python.Type {
				return &f.Statements[0].StatementList.Statements[0].AssignmentStatement.StarredExpression.Expression.ConditionalExpression.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpression.AndExpression.ShiftExpression.AddExpression.MultiplyExpression.UnaryExpression.PowerExpression.PrimaryExpression.Call.ArgumentList.StarredAndKeywordArguments[0].KeywordItem.Expression
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "AssignmentStatement", "StarredExpression", "Expression", "ConditionalExpression", "OrTest", "AndTest", "NotTest", "Comparison", "OrExpression", "XorExpression", "AndExpression", "ShiftExpression", "AddExpression", "MultiplyExpression", "UnaryExpression", "PowerExpression", "PrimaryExpression", "ArgumentListOrComprehension", "ArgumentList", "StarredOrKeyword", "KeywordItem", "Expression"},
		},
		{ // 150
			"a(**b, c=d)",
			func(f *python.File) python.Type {
				return f.Statements[0].StatementList.Statements[0].AssignmentStatement.StarredExpression.Expression.ConditionalExpression.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpression.AndExpression.ShiftExpression.AddExpression.MultiplyExpression.UnaryExpression.PowerExpression.PrimaryExpression.Call.ArgumentList.KeywordArguments[1].KeywordItem
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "AssignmentStatement", "StarredExpression", "Expression", "ConditionalExpression", "OrTest", "AndTest", "NotTest", "Comparison", "OrExpression", "XorExpression", "AndExpression", "ShiftExpression", "AddExpression", "MultiplyExpression", "UnaryExpression", "PowerExpression", "PrimaryExpression", "ArgumentListOrComprehension", "ArgumentList", "KeywordArgument", "KeywordItem"},
		},
		{ // 151
			"a(**b, c=d)",
			func(f *python.File) python.Type {
				return f.Statements[0].StatementList.Statements[0].AssignmentStatement.StarredExpression.Expression.ConditionalExpression.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpression.AndExpression.ShiftExpression.AddExpression.MultiplyExpression.UnaryExpression.PowerExpression.PrimaryExpression.Call.ArgumentList.KeywordArguments[0].Expression
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "AssignmentStatement", "StarredExpression", "Expression", "ConditionalExpression", "OrTest", "AndTest", "NotTest", "Comparison", "OrExpression", "XorExpression", "AndExpression", "ShiftExpression", "AddExpression", "MultiplyExpression", "UnaryExpression", "PowerExpression", "PrimaryExpression", "ArgumentListOrComprehension", "ArgumentList", "KeywordArgument", "Expression"},
		},
		{ // 152
			"a[b,c]",
			nilRet,
			nil,
		},
		{ // 153
			"a[b,c]",
			func(f *python.File) python.Type {
				return &f.Statements[0].StatementList.Statements[0].AssignmentStatement.StarredExpression.Expression.ConditionalExpression.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpression.AndExpression.ShiftExpression.AddExpression.MultiplyExpression.UnaryExpression.PowerExpression.PrimaryExpression.Slicing.SliceItems[0]
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "AssignmentStatement", "StarredExpression", "Expression", "ConditionalExpression", "OrTest", "AndTest", "NotTest", "Comparison", "OrExpression", "XorExpression", "AndExpression", "ShiftExpression", "AddExpression", "MultiplyExpression", "UnaryExpression", "PowerExpression", "PrimaryExpression", "SliceList", "SliceItem"},
		},
		{ // 154
			"a[b,c]",
			func(f *python.File) python.Type {
				return &f.Statements[0].StatementList.Statements[0].AssignmentStatement.StarredExpression.Expression.ConditionalExpression.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpression.AndExpression.ShiftExpression.AddExpression.MultiplyExpression.UnaryExpression.PowerExpression.PrimaryExpression.Slicing.SliceItems[1]
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "AssignmentStatement", "StarredExpression", "Expression", "ConditionalExpression", "OrTest", "AndTest", "NotTest", "Comparison", "OrExpression", "XorExpression", "AndExpression", "ShiftExpression", "AddExpression", "MultiplyExpression", "UnaryExpression", "PowerExpression", "PrimaryExpression", "SliceList", "SliceItem"},
		},
		{ // 155
			"a[b]",
			func(f *python.File) python.Type {
				return f.Statements[0].StatementList.Statements[0].AssignmentStatement.StarredExpression.Expression.ConditionalExpression.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpression.AndExpression.ShiftExpression.AddExpression.MultiplyExpression.UnaryExpression.PowerExpression.PrimaryExpression.Slicing.SliceItems[0].Expression
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "AssignmentStatement", "StarredExpression", "Expression", "ConditionalExpression", "OrTest", "AndTest", "NotTest", "Comparison", "OrExpression", "XorExpression", "AndExpression", "ShiftExpression", "AddExpression", "MultiplyExpression", "UnaryExpression", "PowerExpression", "PrimaryExpression", "SliceList", "SliceItem", "Expression"},
		},
		{ // 156
			"a[b:c:d]",
			func(f *python.File) python.Type {
				return f.Statements[0].StatementList.Statements[0].AssignmentStatement.StarredExpression.Expression.ConditionalExpression.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpression.AndExpression.ShiftExpression.AddExpression.MultiplyExpression.UnaryExpression.PowerExpression.PrimaryExpression.Slicing.SliceItems[0].LowerBound
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "AssignmentStatement", "StarredExpression", "Expression", "ConditionalExpression", "OrTest", "AndTest", "NotTest", "Comparison", "OrExpression", "XorExpression", "AndExpression", "ShiftExpression", "AddExpression", "MultiplyExpression", "UnaryExpression", "PowerExpression", "PrimaryExpression", "SliceList", "SliceItem", "Expression"},
		},
		{ // 157
			"a[b:c:d]",
			func(f *python.File) python.Type {
				return f.Statements[0].StatementList.Statements[0].AssignmentStatement.StarredExpression.Expression.ConditionalExpression.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpression.AndExpression.ShiftExpression.AddExpression.MultiplyExpression.UnaryExpression.PowerExpression.PrimaryExpression.Slicing.SliceItems[0].UpperBound
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "AssignmentStatement", "StarredExpression", "Expression", "ConditionalExpression", "OrTest", "AndTest", "NotTest", "Comparison", "OrExpression", "XorExpression", "AndExpression", "ShiftExpression", "AddExpression", "MultiplyExpression", "UnaryExpression", "PowerExpression", "PrimaryExpression", "SliceList", "SliceItem", "Expression"},
		},
		{ // 158
			"a[b:c:d]",
			func(f *python.File) python.Type {
				return f.Statements[0].StatementList.Statements[0].AssignmentStatement.StarredExpression.Expression.ConditionalExpression.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpression.AndExpression.ShiftExpression.AddExpression.MultiplyExpression.UnaryExpression.PowerExpression.PrimaryExpression.Slicing.SliceItems[0].Stride
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "AssignmentStatement", "StarredExpression", "Expression", "ConditionalExpression", "OrTest", "AndTest", "NotTest", "Comparison", "OrExpression", "XorExpression", "AndExpression", "ShiftExpression", "AddExpression", "MultiplyExpression", "UnaryExpression", "PowerExpression", "PrimaryExpression", "SliceList", "SliceItem", "Expression"},
		},
		{ // 159
			"a[b:c]",
			func(f *python.File) python.Type {
				return f.Statements[0].StatementList.Statements[0].AssignmentStatement.StarredExpression.Expression.ConditionalExpression.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpression.AndExpression.ShiftExpression.AddExpression.MultiplyExpression.UnaryExpression.PowerExpression.PrimaryExpression.Slicing.SliceItems[0].Stride
			},
			nil,
		},
		{ // 160
			"[]",
			func(f *python.File) python.Type {
				return f.Statements[0].StatementList.Statements[0].AssignmentStatement.StarredExpression.Expression.ConditionalExpression.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpression.AndExpression.ShiftExpression.AddExpression.MultiplyExpression.UnaryExpression.PowerExpression.PrimaryExpression.Atom.Enclosure
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "AssignmentStatement", "StarredExpression", "Expression", "ConditionalExpression", "OrTest", "AndTest", "NotTest", "Comparison", "OrExpression", "XorExpression", "AndExpression", "ShiftExpression", "AddExpression", "MultiplyExpression", "UnaryExpression", "PowerExpression", "PrimaryExpression", "Atom", "Enclosure"},
		},
		{ // 161
			"()",
			func(f *python.File) python.Type {
				return f.Statements[0].StatementList.Statements[0].AssignmentStatement.StarredExpression.Expression.ConditionalExpression.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpression.AndExpression.ShiftExpression.AddExpression.MultiplyExpression.UnaryExpression.PowerExpression.PrimaryExpression.Atom.Enclosure.ParenthForm
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "AssignmentStatement", "StarredExpression", "Expression", "ConditionalExpression", "OrTest", "AndTest", "NotTest", "Comparison", "OrExpression", "XorExpression", "AndExpression", "ShiftExpression", "AddExpression", "MultiplyExpression", "UnaryExpression", "PowerExpression", "PrimaryExpression", "Atom", "Enclosure", "StarredExpression"},
		},
		{ // 162
			"[]",
			func(f *python.File) python.Type {
				return f.Statements[0].StatementList.Statements[0].AssignmentStatement.StarredExpression.Expression.ConditionalExpression.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpression.AndExpression.ShiftExpression.AddExpression.MultiplyExpression.UnaryExpression.PowerExpression.PrimaryExpression.Atom.Enclosure.ListDisplay
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "AssignmentStatement", "StarredExpression", "Expression", "ConditionalExpression", "OrTest", "AndTest", "NotTest", "Comparison", "OrExpression", "XorExpression", "AndExpression", "ShiftExpression", "AddExpression", "MultiplyExpression", "UnaryExpression", "PowerExpression", "PrimaryExpression", "Atom", "Enclosure", "FlexibleExpressionListOrComprehension"},
		},
		{ // 163
			"{}",
			func(f *python.File) python.Type {
				return f.Statements[0].StatementList.Statements[0].AssignmentStatement.StarredExpression.Expression.ConditionalExpression.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpression.AndExpression.ShiftExpression.AddExpression.MultiplyExpression.UnaryExpression.PowerExpression.PrimaryExpression.Atom.Enclosure.DictDisplay
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "AssignmentStatement", "StarredExpression", "Expression", "ConditionalExpression", "OrTest", "AndTest", "NotTest", "Comparison", "OrExpression", "XorExpression", "AndExpression", "ShiftExpression", "AddExpression", "MultiplyExpression", "UnaryExpression", "PowerExpression", "PrimaryExpression", "Atom", "Enclosure", "DictDisplay"},
		},
		{ // 164
			"{a}",
			func(f *python.File) python.Type {
				return f.Statements[0].StatementList.Statements[0].AssignmentStatement.StarredExpression.Expression.ConditionalExpression.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpression.AndExpression.ShiftExpression.AddExpression.MultiplyExpression.UnaryExpression.PowerExpression.PrimaryExpression.Atom.Enclosure.SetDisplay
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "AssignmentStatement", "StarredExpression", "Expression", "ConditionalExpression", "OrTest", "AndTest", "NotTest", "Comparison", "OrExpression", "XorExpression", "AndExpression", "ShiftExpression", "AddExpression", "MultiplyExpression", "UnaryExpression", "PowerExpression", "PrimaryExpression", "Atom", "Enclosure", "FlexibleExpressionListOrComprehension"},
		},
		{ // 165
			"(a for b in c)",
			func(f *python.File) python.Type {
				return f.Statements[0].StatementList.Statements[0].AssignmentStatement.StarredExpression.Expression.ConditionalExpression.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpression.AndExpression.ShiftExpression.AddExpression.MultiplyExpression.UnaryExpression.PowerExpression.PrimaryExpression.Atom.Enclosure.GeneratorExpression
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "AssignmentStatement", "StarredExpression", "Expression", "ConditionalExpression", "OrTest", "AndTest", "NotTest", "Comparison", "OrExpression", "XorExpression", "AndExpression", "ShiftExpression", "AddExpression", "MultiplyExpression", "UnaryExpression", "PowerExpression", "PrimaryExpression", "Atom", "Enclosure", "GeneratorExpression"},
		},
		{ // 166
			"(yield a)",
			func(f *python.File) python.Type {
				return f.Statements[0].StatementList.Statements[0].AssignmentStatement.StarredExpression.Expression.ConditionalExpression.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpression.AndExpression.ShiftExpression.AddExpression.MultiplyExpression.UnaryExpression.PowerExpression.PrimaryExpression.Atom.Enclosure.YieldAtom
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "AssignmentStatement", "StarredExpression", "Expression", "ConditionalExpression", "OrTest", "AndTest", "NotTest", "Comparison", "OrExpression", "XorExpression", "AndExpression", "ShiftExpression", "AddExpression", "MultiplyExpression", "UnaryExpression", "PowerExpression", "PrimaryExpression", "Atom", "Enclosure", "YieldExpression"},
		},
		{ // 167
			"[]",
			nilRet,
			nil,
		},
		{ // 168
			"[a]",
			func(f *python.File) python.Type {
				return f.Statements[0].StatementList.Statements[0].AssignmentStatement.StarredExpression.Expression.ConditionalExpression.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpression.AndExpression.ShiftExpression.AddExpression.MultiplyExpression.UnaryExpression.PowerExpression.PrimaryExpression.Atom.Enclosure.ListDisplay.FlexibleExpressionList
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "AssignmentStatement", "StarredExpression", "Expression", "ConditionalExpression", "OrTest", "AndTest", "NotTest", "Comparison", "OrExpression", "XorExpression", "AndExpression", "ShiftExpression", "AddExpression", "MultiplyExpression", "UnaryExpression", "PowerExpression", "PrimaryExpression", "Atom", "Enclosure", "FlexibleExpressionListOrComprehension", "FlexibleExpressionList"},
		},
		{ // 169
			"[a for b in c]",
			func(f *python.File) python.Type {
				return f.Statements[0].StatementList.Statements[0].AssignmentStatement.StarredExpression.Expression.ConditionalExpression.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpression.AndExpression.ShiftExpression.AddExpression.MultiplyExpression.UnaryExpression.PowerExpression.PrimaryExpression.Atom.Enclosure.ListDisplay.Comprehension
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "AssignmentStatement", "StarredExpression", "Expression", "ConditionalExpression", "OrTest", "AndTest", "NotTest", "Comparison", "OrExpression", "XorExpression", "AndExpression", "ShiftExpression", "AddExpression", "MultiplyExpression", "UnaryExpression", "PowerExpression", "PrimaryExpression", "Atom", "Enclosure", "FlexibleExpressionListOrComprehension", "Comprehension"},
		},
		{ // 170
			"[a, b]",
			nilRet,
			nil,
		},
		{ // 171
			"[a, b]",
			func(f *python.File) python.Type {
				return &f.Statements[0].StatementList.Statements[0].AssignmentStatement.StarredExpression.Expression.ConditionalExpression.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpression.AndExpression.ShiftExpression.AddExpression.MultiplyExpression.UnaryExpression.PowerExpression.PrimaryExpression.Atom.Enclosure.ListDisplay.FlexibleExpressionList.FlexibleExpressions[0]
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "AssignmentStatement", "StarredExpression", "Expression", "ConditionalExpression", "OrTest", "AndTest", "NotTest", "Comparison", "OrExpression", "XorExpression", "AndExpression", "ShiftExpression", "AddExpression", "MultiplyExpression", "UnaryExpression", "PowerExpression", "PrimaryExpression", "Atom", "Enclosure", "FlexibleExpressionListOrComprehension", "FlexibleExpressionList", "FlexibleExpression"},
		},
		{ // 172
			"[a, b]",
			func(f *python.File) python.Type {
				return &f.Statements[0].StatementList.Statements[0].AssignmentStatement.StarredExpression.Expression.ConditionalExpression.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpression.AndExpression.ShiftExpression.AddExpression.MultiplyExpression.UnaryExpression.PowerExpression.PrimaryExpression.Atom.Enclosure.ListDisplay.FlexibleExpressionList.FlexibleExpressions[1]
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "AssignmentStatement", "StarredExpression", "Expression", "ConditionalExpression", "OrTest", "AndTest", "NotTest", "Comparison", "OrExpression", "XorExpression", "AndExpression", "ShiftExpression", "AddExpression", "MultiplyExpression", "UnaryExpression", "PowerExpression", "PrimaryExpression", "Atom", "Enclosure", "FlexibleExpressionListOrComprehension", "FlexibleExpressionList", "FlexibleExpression"},
		},
		{ // 173
			"[a]",
			func(f *python.File) python.Type {
				return f.Statements[0].StatementList.Statements[0].AssignmentStatement.StarredExpression.Expression.ConditionalExpression.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpression.AndExpression.ShiftExpression.AddExpression.MultiplyExpression.UnaryExpression.PowerExpression.PrimaryExpression.Atom.Enclosure.ListDisplay.FlexibleExpressionList.FlexibleExpressions[0].AssignmentExpression
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "AssignmentStatement", "StarredExpression", "Expression", "ConditionalExpression", "OrTest", "AndTest", "NotTest", "Comparison", "OrExpression", "XorExpression", "AndExpression", "ShiftExpression", "AddExpression", "MultiplyExpression", "UnaryExpression", "PowerExpression", "PrimaryExpression", "Atom", "Enclosure", "FlexibleExpressionListOrComprehension", "FlexibleExpressionList", "FlexibleExpression", "AssignmentExpression"},
		},
		{ // 174
			"[*a]",
			func(f *python.File) python.Type {
				return f.Statements[0].StatementList.Statements[0].AssignmentStatement.StarredExpression.Expression.ConditionalExpression.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpression.AndExpression.ShiftExpression.AddExpression.MultiplyExpression.UnaryExpression.PowerExpression.PrimaryExpression.Atom.Enclosure.ListDisplay.FlexibleExpressionList.FlexibleExpressions[0].StarredExpression
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "AssignmentStatement", "StarredExpression", "Expression", "ConditionalExpression", "OrTest", "AndTest", "NotTest", "Comparison", "OrExpression", "XorExpression", "AndExpression", "ShiftExpression", "AddExpression", "MultiplyExpression", "UnaryExpression", "PowerExpression", "PrimaryExpression", "Atom", "Enclosure", "FlexibleExpressionListOrComprehension", "FlexibleExpressionList", "FlexibleExpression", "OrExpression"},
		},
		{ // 175
			"[a for b in c]",
			func(f *python.File) python.Type {
				return &f.Statements[0].StatementList.Statements[0].AssignmentStatement.StarredExpression.Expression.ConditionalExpression.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpression.AndExpression.ShiftExpression.AddExpression.MultiplyExpression.UnaryExpression.PowerExpression.PrimaryExpression.Atom.Enclosure.ListDisplay.Comprehension.AssignmentExpression
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "AssignmentStatement", "StarredExpression", "Expression", "ConditionalExpression", "OrTest", "AndTest", "NotTest", "Comparison", "OrExpression", "XorExpression", "AndExpression", "ShiftExpression", "AddExpression", "MultiplyExpression", "UnaryExpression", "PowerExpression", "PrimaryExpression", "Atom", "Enclosure", "FlexibleExpressionListOrComprehension", "Comprehension", "AssignmentExpression"},
		},
		{ // 176
			"[a for b in c]",
			func(f *python.File) python.Type {
				return &f.Statements[0].StatementList.Statements[0].AssignmentStatement.StarredExpression.Expression.ConditionalExpression.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpression.AndExpression.ShiftExpression.AddExpression.MultiplyExpression.UnaryExpression.PowerExpression.PrimaryExpression.Atom.Enclosure.ListDisplay.Comprehension.ComprehensionFor
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "AssignmentStatement", "StarredExpression", "Expression", "ConditionalExpression", "OrTest", "AndTest", "NotTest", "Comparison", "OrExpression", "XorExpression", "AndExpression", "ShiftExpression", "AddExpression", "MultiplyExpression", "UnaryExpression", "PowerExpression", "PrimaryExpression", "Atom", "Enclosure", "FlexibleExpressionListOrComprehension", "Comprehension", "ComprehensionFor"},
		},
		{ // 177
			"{**a}",
			nilRet,
			nil,
		},
		{ // 178
			"{**a, b: c}",
			func(f *python.File) python.Type {
				return &f.Statements[0].StatementList.Statements[0].AssignmentStatement.StarredExpression.Expression.ConditionalExpression.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpression.AndExpression.ShiftExpression.AddExpression.MultiplyExpression.UnaryExpression.PowerExpression.PrimaryExpression.Atom.Enclosure.DictDisplay.DictItems[0]
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "AssignmentStatement", "StarredExpression", "Expression", "ConditionalExpression", "OrTest", "AndTest", "NotTest", "Comparison", "OrExpression", "XorExpression", "AndExpression", "ShiftExpression", "AddExpression", "MultiplyExpression", "UnaryExpression", "PowerExpression", "PrimaryExpression", "Atom", "Enclosure", "DictDisplay", "DictItem"},
		},
		{ // 179
			"{**a, b: c}",
			func(f *python.File) python.Type {
				return &f.Statements[0].StatementList.Statements[0].AssignmentStatement.StarredExpression.Expression.ConditionalExpression.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpression.AndExpression.ShiftExpression.AddExpression.MultiplyExpression.UnaryExpression.PowerExpression.PrimaryExpression.Atom.Enclosure.DictDisplay.DictItems[1]
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "AssignmentStatement", "StarredExpression", "Expression", "ConditionalExpression", "OrTest", "AndTest", "NotTest", "Comparison", "OrExpression", "XorExpression", "AndExpression", "ShiftExpression", "AddExpression", "MultiplyExpression", "UnaryExpression", "PowerExpression", "PrimaryExpression", "Atom", "Enclosure", "DictDisplay", "DictItem"},
		},
		{ // 180
			"{a: b for c in d}",
			func(f *python.File) python.Type {
				return f.Statements[0].StatementList.Statements[0].AssignmentStatement.StarredExpression.Expression.ConditionalExpression.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpression.AndExpression.ShiftExpression.AddExpression.MultiplyExpression.UnaryExpression.PowerExpression.PrimaryExpression.Atom.Enclosure.DictDisplay.DictComprehension
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "AssignmentStatement", "StarredExpression", "Expression", "ConditionalExpression", "OrTest", "AndTest", "NotTest", "Comparison", "OrExpression", "XorExpression", "AndExpression", "ShiftExpression", "AddExpression", "MultiplyExpression", "UnaryExpression", "PowerExpression", "PrimaryExpression", "Atom", "Enclosure", "DictDisplay", "ComprehensionFor"},
		},
		{ // 181
			"{**a}",
			func(f *python.File) python.Type {
				return f.Statements[0].StatementList.Statements[0].AssignmentStatement.StarredExpression.Expression.ConditionalExpression.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpression.AndExpression.ShiftExpression.AddExpression.MultiplyExpression.UnaryExpression.PowerExpression.PrimaryExpression.Atom.Enclosure.DictDisplay.DictItems[0].OrExpression
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "AssignmentStatement", "StarredExpression", "Expression", "ConditionalExpression", "OrTest", "AndTest", "NotTest", "Comparison", "OrExpression", "XorExpression", "AndExpression", "ShiftExpression", "AddExpression", "MultiplyExpression", "UnaryExpression", "PowerExpression", "PrimaryExpression", "Atom", "Enclosure", "DictDisplay", "DictItem", "OrExpression"},
		},
		{ // 182
			"{a: b}",
			func(f *python.File) python.Type {
				return f.Statements[0].StatementList.Statements[0].AssignmentStatement.StarredExpression.Expression.ConditionalExpression.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpression.AndExpression.ShiftExpression.AddExpression.MultiplyExpression.UnaryExpression.PowerExpression.PrimaryExpression.Atom.Enclosure.DictDisplay.DictItems[0].Key
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "AssignmentStatement", "StarredExpression", "Expression", "ConditionalExpression", "OrTest", "AndTest", "NotTest", "Comparison", "OrExpression", "XorExpression", "AndExpression", "ShiftExpression", "AddExpression", "MultiplyExpression", "UnaryExpression", "PowerExpression", "PrimaryExpression", "Atom", "Enclosure", "DictDisplay", "DictItem", "Expression"},
		},
		{ // 183
			"{a: b}",
			func(f *python.File) python.Type {
				return f.Statements[0].StatementList.Statements[0].AssignmentStatement.StarredExpression.Expression.ConditionalExpression.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpression.AndExpression.ShiftExpression.AddExpression.MultiplyExpression.UnaryExpression.PowerExpression.PrimaryExpression.Atom.Enclosure.DictDisplay.DictItems[0].Value
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "AssignmentStatement", "StarredExpression", "Expression", "ConditionalExpression", "OrTest", "AndTest", "NotTest", "Comparison", "OrExpression", "XorExpression", "AndExpression", "ShiftExpression", "AddExpression", "MultiplyExpression", "UnaryExpression", "PowerExpression", "PrimaryExpression", "Atom", "Enclosure", "DictDisplay", "DictItem", "Expression"},
		},
		{ // 184
			"(a for b in c)",
			func(f *python.File) python.Type {
				return &f.Statements[0].StatementList.Statements[0].AssignmentStatement.StarredExpression.Expression.ConditionalExpression.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpression.AndExpression.ShiftExpression.AddExpression.MultiplyExpression.UnaryExpression.PowerExpression.PrimaryExpression.Atom.Enclosure.GeneratorExpression.Expression
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "AssignmentStatement", "StarredExpression", "Expression", "ConditionalExpression", "OrTest", "AndTest", "NotTest", "Comparison", "OrExpression", "XorExpression", "AndExpression", "ShiftExpression", "AddExpression", "MultiplyExpression", "UnaryExpression", "PowerExpression", "PrimaryExpression", "Atom", "Enclosure", "GeneratorExpression", "Expression"},
		},
		{ // 185
			"(a for b in c)",
			func(f *python.File) python.Type {
				return &f.Statements[0].StatementList.Statements[0].AssignmentStatement.StarredExpression.Expression.ConditionalExpression.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpression.AndExpression.ShiftExpression.AddExpression.MultiplyExpression.UnaryExpression.PowerExpression.PrimaryExpression.Atom.Enclosure.GeneratorExpression.ComprehensionFor
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "AssignmentStatement", "StarredExpression", "Expression", "ConditionalExpression", "OrTest", "AndTest", "NotTest", "Comparison", "OrExpression", "XorExpression", "AndExpression", "ShiftExpression", "AddExpression", "MultiplyExpression", "UnaryExpression", "PowerExpression", "PrimaryExpression", "Atom", "Enclosure", "GeneratorExpression", "ComprehensionFor"},
		},
		{ // 186
			"lambda b: c",
			func(f *python.File) python.Type {
				return f.Statements[0].StatementList.Statements[0].AssignmentStatement.StarredExpression.Expression.LambdaExpression.ParameterList
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "AssignmentStatement", "StarredExpression", "Expression", "LambdaExpression", "ParameterList"},
		},
		{ // 187
			"lambda b: c",
			func(f *python.File) python.Type {
				return &f.Statements[0].StatementList.Statements[0].AssignmentStatement.StarredExpression.Expression.LambdaExpression.Expression
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "AssignmentStatement", "StarredExpression", "Expression", "LambdaExpression", "Expression"},
		},
		{ // 188
			"def a(): b = yield from c",
			func(f *python.File) python.Type {
				return f.Statements[0].CompoundStatement.Func.Suite.StatementList.Statements[0].AssignmentStatement.YieldExpression.From
			},
			[]string{"File", "Statement", "CompoundStatement", "FuncDefinition", "Suite", "StatementList", "SimpleStatement", "AssignmentStatement", "YieldExpression", "Expression"},
		},
		{ // 189
			"def a(): b = yield c",
			func(f *python.File) python.Type {
				return f.Statements[0].CompoundStatement.Func.Suite.StatementList.Statements[0].AssignmentStatement.YieldExpression.ExpressionList
			},
			[]string{"File", "Statement", "CompoundStatement", "FuncDefinition", "Suite", "StatementList", "SimpleStatement", "AssignmentStatement", "YieldExpression", "ExpressionList"},
		},
		{ // 190
			"def a(): b = yield c, d",
			nilRet,
			nil,
		},
		{ // 191
			"def a(): b = yield c, d",
			func(f *python.File) python.Type {
				return &f.Statements[0].CompoundStatement.Func.Suite.StatementList.Statements[0].AssignmentStatement.YieldExpression.ExpressionList.Expressions[1]
			},
			[]string{"File", "Statement", "CompoundStatement", "FuncDefinition", "Suite", "StatementList", "SimpleStatement", "AssignmentStatement", "YieldExpression", "ExpressionList", "Expression"},
		},
		{ // 192
			"def a(): b = yield c, d",
			func(f *python.File) python.Type {
				return &f.Statements[0].CompoundStatement.Func.Suite.StatementList.Statements[0].AssignmentStatement.YieldExpression.ExpressionList.Expressions[1]
			},
			[]string{"File", "Statement", "CompoundStatement", "FuncDefinition", "Suite", "StatementList", "SimpleStatement", "AssignmentStatement", "YieldExpression", "ExpressionList", "Expression"},
		},
		{ // 193
			"del a",
			func(f *python.File) python.Type {
				return &f.Statements[0].StatementList.Statements[0].DelStatement.TargetList
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "DelStatement", "TargetList"},
		},
		{ // 194
			"def a(): return b",
			func(f *python.File) python.Type {
				return f.Statements[0].CompoundStatement.Func.Suite.StatementList.Statements[0].ReturnStatement.Expression
			},
			[]string{"File", "Statement", "CompoundStatement", "FuncDefinition", "Suite", "StatementList", "SimpleStatement", "ReturnStatement", "Expression"},
		},
		{ // 195
			"def a(): return",
			func(f *python.File) python.Type {
				return f.Statements[0].CompoundStatement.Func.Suite.StatementList.Statements[0].ReturnStatement.Expression
			},
			nil,
		},
		{ // 196
			"raise",
			func(f *python.File) python.Type {
				return f.Statements[0].StatementList.Statements[0].RaiseStatement.Expression
			},
			nil,
		},
		{ // 197
			"raise a",
			func(f *python.File) python.Type {
				return f.Statements[0].StatementList.Statements[0].RaiseStatement.Expression
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "RaiseStatement", "Expression"},
		},
		{ // 198
			"raise a from b",
			func(f *python.File) python.Type {
				return f.Statements[0].StatementList.Statements[0].RaiseStatement.From
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "RaiseStatement", "Expression"},
		},
		{ // 199
			"import a, b",
			nilRet,
			nil,
		},
		{ // 200
			"import a.b",
			func(f *python.File) python.Type {
				return f.Statements[0].StatementList.Statements[0].ImportStatement.Modules[0]
			},
			nil,
		},
		{ // 201
			"import a, b",
			func(f *python.File) python.Type {
				return &f.Statements[0].StatementList.Statements[0].ImportStatement.Modules[0]
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "ImportStatement", "ModuleAs"},
		},
		{ // 202
			"import a, b",
			func(f *python.File) python.Type {
				return &f.Statements[0].StatementList.Statements[0].ImportStatement.Modules[1]
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "ImportStatement", "ModuleAs"},
		},
		{ // 203

			"from a import b, c",
			func(f *python.File) python.Type {
				return f.Statements[0].StatementList.Statements[0].ImportStatement.RelativeModule
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "ImportStatement", "RelativeModule"},
		},
		{ // 204

			"from a import b, c",
			func(f *python.File) python.Type {
				return &f.Statements[0].StatementList.Statements[0].ImportStatement.Modules[0]
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "ImportStatement", "ModuleAs"},
		},
		{ // 205
			"from a import b, c",
			func(f *python.File) python.Type {
				return &f.Statements[0].StatementList.Statements[0].ImportStatement.Modules[1]
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "ImportStatement", "ModuleAs"},
		},
		{ // 206
			"import a",
			func(f *python.File) python.Type {
				return &f.Statements[0].StatementList.Statements[0].ImportStatement.Modules[0].Module
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "ImportStatement", "ModuleAs", "Module"},
		},
		{ // 207
			"import a",
			func(f *python.File) python.Type {
				return &f.Statements[0].StatementList.Statements[0].ImportStatement.Modules[0].Module.Identifiers[0]
			},
			nil,
		},
		{ // 208

			"from . import a",
			func(f *python.File) python.Type {
				return f.Statements[0].StatementList.Statements[0].ImportStatement.RelativeModule.Module
			},
			nil,
		},
		{ // 209

			"from a import b",
			func(f *python.File) python.Type {
				return f.Statements[0].StatementList.Statements[0].ImportStatement.RelativeModule.Module
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "ImportStatement", "RelativeModule", "Module"},
		},
		{ // 210
			"global a",
			nilRet,
			nil,
		},
		{ // 211
			"nonlocal a",
			nilRet,
			nil,
		},
		{ // 212
			"type a[b, c] = d",
			func(f *python.File) python.Type {
				return f.Statements[0].StatementList.Statements[0].TypeStatement.TypeParams
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "TypeStatement", "TypeParams"},
		},
		{ // 213
			"type a[b, c] = d",
			func(f *python.File) python.Type {
				return &f.Statements[0].StatementList.Statements[0].TypeStatement.Expression
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "TypeStatement", "Expression"},
		},
		{ // 214
			"a, *b",
			func(f *python.File) python.Type {
				return f.Statements[0].StatementList.Statements[0].AssignmentStatement.StarredExpression.StarredList.StarredItems[0].AssignmentExpression
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "AssignmentStatement", "StarredExpression", "StarredList", "StarredItem", "AssignmentExpression"},
		},
		{ // 215
			"a, *b",
			func(f *python.File) python.Type {
				return f.Statements[0].StatementList.Statements[0].AssignmentStatement.StarredExpression.StarredList.StarredItems[1].OrExpr
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "AssignmentStatement", "StarredExpression", "StarredList", "StarredItem", "OrExpression"},
		},
		{ // 216
			"a = b",
			func(f *python.File) python.Type {
				return f.Statements[0].StatementList.Statements[0].AssignmentStatement.TargetLists[0].Targets[0].PrimaryExpression
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "AssignmentStatement", "TargetList", "Target", "PrimaryExpression"},
		},
		{ // 217
			"[a] = b",
			func(f *python.File) python.Type {
				return f.Statements[0].StatementList.Statements[0].AssignmentStatement.TargetLists[0].Targets[0].Array
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "AssignmentStatement", "TargetList", "Target", "TargetList"},
		},
		{ // 218
			"(a) = b",
			func(f *python.File) python.Type {
				return f.Statements[0].StatementList.Statements[0].AssignmentStatement.TargetLists[0].Targets[0].Tuple
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "AssignmentStatement", "TargetList", "Target", "TargetList"},
		},
		{ // 219
			"*a = b",
			func(f *python.File) python.Type {
				return f.Statements[0].StatementList.Statements[0].AssignmentStatement.TargetLists[0].Targets[0].Star
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "AssignmentStatement", "TargetList", "Target", "Target"},
		},
		{ // 220
			"a += b",
			func(f *python.File) python.Type {
				return &f.Statements[0].StatementList.Statements[0].AugmentedAssignmentStatement.AugTarget
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "AugmentedAssignmentStatement", "AugTarget"},
		},
		{ // 221
			"a += b",
			func(f *python.File) python.Type {
				return f.Statements[0].StatementList.Statements[0].AugmentedAssignmentStatement.ExpressionList
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "AugmentedAssignmentStatement", "ExpressionList"},
		},
		{ // 222
			"a += yield b",
			func(f *python.File) python.Type {
				return f.Statements[0].StatementList.Statements[0].AugmentedAssignmentStatement.YieldExpression
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "AugmentedAssignmentStatement", "YieldExpression"},
		},
		{ // 223
			"a: b = c",
			func(f *python.File) python.Type {
				return &f.Statements[0].StatementList.Statements[0].AnnotatedAssignmentStatement.AugTarget
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "AnnotatedAssignmentStatement", "AugTarget"},
		},
		{ // 224
			"a: b = c",
			func(f *python.File) python.Type {
				return &f.Statements[0].StatementList.Statements[0].AnnotatedAssignmentStatement.Expression
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "AnnotatedAssignmentStatement", "Expression"},
		},
		{ // 225
			"a: b = c",
			func(f *python.File) python.Type {
				return f.Statements[0].StatementList.Statements[0].AnnotatedAssignmentStatement.StarredExpression
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "AnnotatedAssignmentStatement", "StarredExpression"},
		},
		{ // 226
			"a: b = yield c",
			func(f *python.File) python.Type {
				return f.Statements[0].StatementList.Statements[0].AnnotatedAssignmentStatement.YieldExpression
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "AnnotatedAssignmentStatement", "YieldExpression"},
		},
		{ // 227
			"(a for b in c if d)",
			func(f *python.File) python.Type {
				return &f.Statements[0].StatementList.Statements[0].AssignmentStatement.StarredExpression.Expression.ConditionalExpression.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpression.AndExpression.ShiftExpression.AddExpression.MultiplyExpression.UnaryExpression.PowerExpression.PrimaryExpression.Atom.Enclosure.GeneratorExpression.ComprehensionFor.TargetList
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "AssignmentStatement", "StarredExpression", "Expression", "ConditionalExpression", "OrTest", "AndTest", "NotTest", "Comparison", "OrExpression", "XorExpression", "AndExpression", "ShiftExpression", "AddExpression", "MultiplyExpression", "UnaryExpression", "PowerExpression", "PrimaryExpression", "Atom", "Enclosure", "GeneratorExpression", "ComprehensionFor", "TargetList"},
		},
		{ // 228
			"(a for b in c if d)",
			func(f *python.File) python.Type {
				return &f.Statements[0].StatementList.Statements[0].AssignmentStatement.StarredExpression.Expression.ConditionalExpression.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpression.AndExpression.ShiftExpression.AddExpression.MultiplyExpression.UnaryExpression.PowerExpression.PrimaryExpression.Atom.Enclosure.GeneratorExpression.ComprehensionFor.OrTest
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "AssignmentStatement", "StarredExpression", "Expression", "ConditionalExpression", "OrTest", "AndTest", "NotTest", "Comparison", "OrExpression", "XorExpression", "AndExpression", "ShiftExpression", "AddExpression", "MultiplyExpression", "UnaryExpression", "PowerExpression", "PrimaryExpression", "Atom", "Enclosure", "GeneratorExpression", "ComprehensionFor", "OrTest"},
		},
		{ // 229
			"(a for b in c if d)",
			func(f *python.File) python.Type {
				return f.Statements[0].StatementList.Statements[0].AssignmentStatement.StarredExpression.Expression.ConditionalExpression.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpression.AndExpression.ShiftExpression.AddExpression.MultiplyExpression.UnaryExpression.PowerExpression.PrimaryExpression.Atom.Enclosure.GeneratorExpression.ComprehensionFor.ComprehensionIterator
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "AssignmentStatement", "StarredExpression", "Expression", "ConditionalExpression", "OrTest", "AndTest", "NotTest", "Comparison", "OrExpression", "XorExpression", "AndExpression", "ShiftExpression", "AddExpression", "MultiplyExpression", "UnaryExpression", "PowerExpression", "PrimaryExpression", "Atom", "Enclosure", "GeneratorExpression", "ComprehensionFor", "ComprehensionIterator"},
		},
		{ // 230
			"(a for b in c)",
			func(f *python.File) python.Type {
				return f.Statements[0].StatementList.Statements[0].AssignmentStatement.StarredExpression.Expression.ConditionalExpression.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpression.AndExpression.ShiftExpression.AddExpression.MultiplyExpression.UnaryExpression.PowerExpression.PrimaryExpression.Atom.Enclosure.GeneratorExpression.ComprehensionFor.ComprehensionIterator
			},
			nil,
		},
		{ // 231
			"(a for b in c if d)",
			func(f *python.File) python.Type {
				return f.Statements[0].StatementList.Statements[0].AssignmentStatement.StarredExpression.Expression.ConditionalExpression.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpression.AndExpression.ShiftExpression.AddExpression.MultiplyExpression.UnaryExpression.PowerExpression.PrimaryExpression.Atom.Enclosure.GeneratorExpression.ComprehensionFor.ComprehensionIterator.ComprehensionIf
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "AssignmentStatement", "StarredExpression", "Expression", "ConditionalExpression", "OrTest", "AndTest", "NotTest", "Comparison", "OrExpression", "XorExpression", "AndExpression", "ShiftExpression", "AddExpression", "MultiplyExpression", "UnaryExpression", "PowerExpression", "PrimaryExpression", "Atom", "Enclosure", "GeneratorExpression", "ComprehensionFor", "ComprehensionIterator", "ComprehensionIf"},
		},
		{ // 232
			"(a for b in c for d in e)",
			func(f *python.File) python.Type {
				return f.Statements[0].StatementList.Statements[0].AssignmentStatement.StarredExpression.Expression.ConditionalExpression.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpression.AndExpression.ShiftExpression.AddExpression.MultiplyExpression.UnaryExpression.PowerExpression.PrimaryExpression.Atom.Enclosure.GeneratorExpression.ComprehensionFor.ComprehensionIterator.ComprehensionFor
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "AssignmentStatement", "StarredExpression", "Expression", "ConditionalExpression", "OrTest", "AndTest", "NotTest", "Comparison", "OrExpression", "XorExpression", "AndExpression", "ShiftExpression", "AddExpression", "MultiplyExpression", "UnaryExpression", "PowerExpression", "PrimaryExpression", "Atom", "Enclosure", "GeneratorExpression", "ComprehensionFor", "ComprehensionIterator", "ComprehensionFor"},
		},
		{ // 233
			"(a for b in c if d for e in f)",
			func(f *python.File) python.Type {
				return &f.Statements[0].StatementList.Statements[0].AssignmentStatement.StarredExpression.Expression.ConditionalExpression.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpression.AndExpression.ShiftExpression.AddExpression.MultiplyExpression.UnaryExpression.PowerExpression.PrimaryExpression.Atom.Enclosure.GeneratorExpression.ComprehensionFor.ComprehensionIterator.ComprehensionIf.OrTest
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "AssignmentStatement", "StarredExpression", "Expression", "ConditionalExpression", "OrTest", "AndTest", "NotTest", "Comparison", "OrExpression", "XorExpression", "AndExpression", "ShiftExpression", "AddExpression", "MultiplyExpression", "UnaryExpression", "PowerExpression", "PrimaryExpression", "Atom", "Enclosure", "GeneratorExpression", "ComprehensionFor", "ComprehensionIterator", "ComprehensionIf", "OrTest"},
		},
		{ // 234
			"(a for b in c if d for e in f)",
			func(f *python.File) python.Type {
				return f.Statements[0].StatementList.Statements[0].AssignmentStatement.StarredExpression.Expression.ConditionalExpression.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpression.AndExpression.ShiftExpression.AddExpression.MultiplyExpression.UnaryExpression.PowerExpression.PrimaryExpression.Atom.Enclosure.GeneratorExpression.ComprehensionFor.ComprehensionIterator.ComprehensionIf.ComprehensionIterator
			},
			[]string{"File", "Statement", "StatementList", "SimpleStatement", "AssignmentStatement", "StarredExpression", "Expression", "ConditionalExpression", "OrTest", "AndTest", "NotTest", "Comparison", "OrExpression", "XorExpression", "AndExpression", "ShiftExpression", "AddExpression", "MultiplyExpression", "UnaryExpression", "PowerExpression", "PrimaryExpression", "Atom", "Enclosure", "GeneratorExpression", "ComprehensionFor", "ComprehensionIterator", "ComprehensionIf", "ComprehensionIterator"},
		},
		{ // 235
			"(a for b in c if d)",
			func(f *python.File) python.Type {
				return f.Statements[0].StatementList.Statements[0].AssignmentStatement.StarredExpression.Expression.ConditionalExpression.OrTest.AndTest.NotTest.Comparison.OrExpression.XorExpression.AndExpression.ShiftExpression.AddExpression.MultiplyExpression.UnaryExpression.PowerExpression.PrimaryExpression.Atom.Enclosure.GeneratorExpression.ComprehensionFor.ComprehensionIterator.ComprehensionIf.ComprehensionIterator
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
				t.Errorf("test %d: expected to receive sentinel error, but didn't", n+1)
			} else if err != nil && test.Level == nil {
				t.Errorf("test %d: expected no error, but received %v", n+1, err)
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

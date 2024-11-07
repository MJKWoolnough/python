package python

import (
	"fmt"
	"testing"

	"vimagination.zapto.org/parser"
)

func TestPrintSource(t *testing.T) {
	for n, test := range [...][3]string{
		{ // 1
			"1+ 2",
			"1+2\n",
			"1 + 2\n",
		},
		{ // 2
			"a+ 1",
			"a+1\n",
			"a + 1\n",
		},
		{ // 3
			"1 + 2 + 3\n",
			"1+2+3\n",
			"1 + 2 + 3\n",
		},
		{ // 4
			"1- 2",
			"1-2\n",
			"1 - 2\n",
		},
		{ // 5
			"a- 1",
			"a-1\n",
			"a - 1\n",
		},
		{ // 6
			"1  -  2  -  3\n",
			"1-2-3\n",
			"1 - 2 - 3\n",
		},
		{ // 7
			"1+2-3\n",
			"1+2-3\n",
			"1 + 2 - 3\n",
		},
		{ // 8
			"a-b+c\n",
			"a-b+c\n",
			"a - b + c\n",
		},
		{ // 9
			"1& 2",
			"1&2\n",
			"1 & 2\n",
		},
		{ // 10
			"a& 1",
			"a&1\n",
			"a & 1\n",
		},
		{ // 11
			"a&b&c\n",
			"a&b&c\n",
			"a & b & c\n",
		},
		{ // 12
			"1 and 2",
			"1 and 2\n",
			"1 and 2\n",
		},
		{ // 13
			"a and 1",
			"a and 1\n",
			"a and 1\n",
		},
		{ // 14
			"a and b and c\n",
			"a and b and c\n",
			"a and b and c\n",
		},
		{ // 15
			"a:b=c\n",
			"a:b=c\n",
			"a: b = c\n",
		},
		{ // 16
			"a : b = yield c",
			"a:b=yield c\n",
			"a: b = yield c\n",
		},
		{ // 17
			"a()",
			"a()\n",
			"a()\n",
		},
		{ // 18
			"a(b)",
			"a(b)\n",
			"a(b)\n",
		},
		{ // 19
			"a(b,c)",
			"a(b,c)\n",
			"a(b, c)\n",
		},
		{ // 20
			"a(b, *c)",
			"a(b,*c)\n",
			"a(b, *c)\n",
		},
		{ // 21
			"a(*b)",
			"a(*b)\n",
			"a(*b)\n",
		},
		{ // 22
			"a(b, **c)",
			"a(b,**c)\n",
			"a(b, **c)\n",
		},
		{ // 23
			"a(*b,**c)",
			"a(*b,**c)\n",
			"a(*b, **c)\n",
		},
		{ // 24
			"a(b, *c, **d)",
			"a(b,*c,**d)\n",
			"a(b, *c, **d)\n",
		},
		{ // 25
			"a(b for c in d)",
			"a(b for c in d)\n",
			"a(b for c in d)\n",
		},
		{ // 26
			"assert a",
			"assert a\n",
			"assert a\n",
		},
		{ // 27
			"assert a,b",
			"assert a,b\n",
			"assert a, b\n",
		},
		{ // 28
			"if a:b\nelif c:d",
			"if a:b\nelif c:d\n",
			"if a: b\nelif c: d\n",
		},
		{ // 29
			"@a:=b\ndef c():d",
			"@a:=b\ndef c():d\n",
			"@a := b\ndef c(): d\n",
		},
		{ // 30
			"a=b",
			"a=b\n",
			"a = b\n",
		},
		{ // 31
			"a=yield b",
			"a=yield b\n",
			"a = yield b\n",
		},
		{ // 32
			"a=b=c",
			"a=b=c\n",
			"a = b = c\n",
		},
		{ // 33
			"a",
			"a\n",
			"a\n",
		},
		{ // 34
			"1",
			"1\n",
			"1\n",
		},
		{ // 35
			"(a)",
			"(a)\n",
			"(a)\n",
		},
		{ // 36
			"try:a\nexcept b:c",
			"try:a\nexcept b:c\n",
			"try: a\nexcept b: c\n",
		},
		{ // 37
			"try:a\nexcept b as c:d",
			"try:a\nexcept b as c:d\n",
			"try: a\nexcept b as c: d\n",
		},
		{ // 38
			"assert a",
			"assert a\n",
			"assert a\n",
		},
		{ // 39
			"assert lambda:a",
			"assert lambda:a\n",
			"assert lambda: a\n",
		},
		{ // 40
			"yield a",
			"yield a\n",
			"yield a\n",
		},
		{ // 41
			"yield a,b",
			"yield a,b\n",
			"yield a, b\n",
		},
		{ // 42
			"a\nb\nc",
			"a\nb\nc\n",
			"a\nb\nc\n",
		},
		{ // 43
			"[a]",
			"[a]\n",
			"[a]\n",
		},
		{ // 44
			"[a for b in c]",
			"[a for b in c]\n",
			"[a for b in c]\n",
		},
		{ // 45
			"[a,b]",
			"[a,b]\n",
			"[a, b]\n",
		},
		{ // 46
			"[a := b]",
			"[a:=b]\n",
			"[a := b]\n",
		},
		{ // 47
			"[* a]",
			"[*a]\n",
			"[*a]\n",
		},
		{ // 48
			"for a in b:c",
			"for a in b:c\n",
			"for a in b: c\n",
		},
		{ // 49
			"async for a in b:c",
			"async for a in b:c\n",
			"async for a in b: c\n",
		},
		{ // 50
			"for a in b:c\nelse:d",
			"for a in b:c\nelse:d\n",
			"for a in b: c\nelse: d\n",
		},
		{ // 51
			"def a():b",
			"def a():b\n",
			"def a(): b\n",
		},
		{ // 52
			"def a():\n\tb",
			"def a():\n\tb\n",
			"def a():\n\tb\n",
		},
		{ // 53
			"def a [b] ():c",
			"def a[b]():c\n",
			"def a[b](): c\n",
		},
		{ // 54
			"def a() -> b:c",
			"def a()->b:c\n",
			"def a() -> b: c\n",
		},
		{ // 55
			"(a for b in c)",
			"(a for b in c)\n",
			"(a for b in c)\n",
		},
		{ // 56
			"global a",
			"global a\n",
			"global a\n",
		},
		{ // 57
			"global a , b",
			"global a,b\n",
			"global a, b\n",
		},
		{ // 58
			"if a:b",
			"if a:b\n",
			"if a: b\n",
		},
		{ // 59
			"if a :b\nelse : c",
			"if a:b\nelse:c\n",
			"if a: b\nelse: c\n",
		},
		{ // 60
			"if a :b\nelif c : d",
			"if a:b\nelif c:d\n",
			"if a: b\nelif c: d\n",
		},
		{ // 61
			"if a :b\nelif c : d\nelif e:f",
			"if a:b\nelif c:d\nelif e:f\n",
			"if a: b\nelif c: d\nelif e: f\n",
		},
		{ // 62
			"if a :b\nelif c : d\nelif e:f\nelse:g",
			"if a:b\nelif c:d\nelif e:f\nelse:g\n",
			"if a: b\nelif c: d\nelif e: f\nelse: g\n",
		},
		{ // 63
			"from a import b",
			"from a import b\n",
			"from a import b\n",
		},
		{ // 64
			"from a import b,c",
			"from a import b,c\n",
			"from a import b, c\n",
		},
		{ // 65
			"import a",
			"import a\n",
			"import a\n",
		},
		{ // 66
			"import a , b",
			"import a,b\n",
			"import a, b\n",
		},
		{ // 67
			"class a(**b):c",
			"class a(**b):c\n",
			"class a(**b): c\n",
		},
		{ // 68
			"class a ( **b, c = d) : e",
			"class a(**b,c=d):e\n",
			"class a(**b, c = d): e\n",
		},
		{ // 69
			"assert lambda:a",
			"assert lambda:a\n",
			"assert lambda: a\n",
		},
		{ // 70
			"assert lambda a : b",
			"assert lambda a:b\n",
			"assert lambda a: b\n",
		},
		{ // 71
			"import a as b",
			"import a as b\n",
			"import a as b\n",
		},
		{ // 72
			"import a.b",
			"import a.b\n",
			"import a.b\n",
		},
		{ // 73
			"import a.b.c",
			"import a.b.c\n",
			"import a.b.c\n",
		},
		{ // 74
			"a*b",
			"a*b\n",
			"a * b\n",
		},
		{ // 75
			"a / b",
			"a/b\n",
			"a / b\n",
		},
		{ // 76
			"a * b / c*d",
			"a*b/c*d\n",
			"a * b / c * d\n",
		},
		{ // 77
			"nonlocal a",
			"nonlocal a\n",
			"nonlocal a\n",
		},
		{ // 78
			"nonlocal a,b , c",
			"nonlocal a,b,c\n",
			"nonlocal a, b, c\n",
		},
		{ // 79
			"not a",
			"not a\n",
			"not a\n",
		},
		{ // 80
			"not not not not a",
			"not not not not a\n",
			"not not not not a\n",
		},
		{ // 81
			"a|b",
			"a|b\n",
			"a | b\n",
		},
		{ // 82
			"a|b | c",
			"a|b|c\n",
			"a | b | c\n",
		},
		{ // 83
			"a or b",
			"a or b\n",
			"a or b\n",
		},
		{ // 84
			"a or b or c",
			"a or b or c\n",
			"a or b or c\n",
		},
		{ // 85
			"def a():b",
			"def a():b\n",
			"def a(): b\n",
		},
		{ // 86
			"def a(b):c",
			"def a(b):c\n",
			"def a(b): c\n",
		},
		{ // 87
			"def a(b,/,c):d",
			"def a(b,/,c):d\n",
			"def a(b, /, c): d\n",
		},
		{ // 88
			"def a(b,c,/,d,e):f",
			"def a(b,c,/,d,e):f\n",
			"def a(b, c, /, d, e): f\n",
		},
		{ // 89
			"def a(b, *c):d",
			"def a(b,*c):d\n",
			"def a(b, *c): d\n",
		},
		{ // 90
			"def a(b, *c, d):e",
			"def a(b,*c,d):e\n",
			"def a(b, *c, d): e\n",
		},
		{ // 91
			"def a(b,**c):d",
			"def a(b,**c):d\n",
			"def a(b, **c): d\n",
		},
		{ // 92
			"def a(b , / , *c):d",
			"def a(b,/,*c):d\n",
			"def a(b, /, *c): d\n",
		},
		{ // 93
			"def a(b , / , *c,d):e",
			"def a(b,/,*c,d):e\n",
			"def a(b, /, *c, d): e\n",
		},
		{ // 94
			"def a(b , / , **c):d",
			"def a(b,/,**c):d\n",
			"def a(b, /, **c): d\n",
		},
		{ // 95
			"def a(b , / , c, d, *e, f, g, **h):i",
			"def a(b,/,c,d,*e,f,g,**h):i\n",
			"def a(b, /, c, d, *e, f, g, **h): i\n",
		},
		{ // 96
			"def a(b , c, d, *e, f, g, **h):i",
			"def a(b,c,d,*e,f,g,**h):i\n",
			"def a(b, c, d, *e, f, g, **h): i\n",
		},
		{ // 97
			"a( b )",
			"a(b)\n",
			"a(b)\n",
		},
		{ // 98
			"a( *b )",
			"a(*b)\n",
			"a(*b)\n",
		},
		{ // 99
			"a ** b",
			"a**b\n",
			"a ** b\n",
		},
		{ // 100
			"await a**b",
			"await a**b\n",
			"await a ** b\n",
		},
		{ // 101
			"a",
			"a\n",
			"a\n",
		},
		{ // 102
			"a.b",
			"a.b\n",
			"a.b\n",
		},
		{ // 103
			"a . b",
			"a.b\n",
			"a.b\n",
		},
		{ // 104
			"a[b]",
			"a[b]\n",
			"a[b]\n",
		},
		{ // 105
			"a[ b ]",
			"a[b]\n",
			"a[b]\n",
		},
		{ // 106
			"a(b)",
			"a(b)\n",
			"a(b)\n",
		},
		{ // 107
			"a( b )",
			"a(b)\n",
			"a(b)\n",
		},
		{ // 108
			"raise",
			"raise\n",
			"raise\n",
		},
		{ // 109
			"raise a",
			"raise a\n",
			"raise a\n",
		},
		{ // 110
			"raise a from b",
			"raise a from b\n",
			"raise a from b\n",
		},
		{ // 111
			"from . a import b",
			"from .a import b\n",
			"from .a import b\n",
		},
		{ // 112
			"from ..a import b",
			"from ..a import b\n",
			"from ..a import b\n",
		},
		{ // 113
			"from ... import a",
			"from ... import a\n",
			"from ... import a\n",
		},
		{ // 114
			"from ....a.b import c",
			"from ....a.b import c\n",
			"from ....a.b import c\n",
		},
		{ // 115
			"def a():\n\treturn",
			"def a():\n\treturn\n",
			"def a():\n\treturn\n",
		},
		{ // 116
			"def a():\n\treturn b",
			"def a():\n\treturn b\n",
			"def a():\n\treturn b\n",
		},
		{ // 117
			"a>>b",
			"a>>b\n",
			"a >> b\n",
		},
		{ // 118
			"a << b",
			"a<<b\n",
			"a << b\n",
		},
		{ // 119
			"assert a",
			"assert a\n",
			"assert a\n",
		},
		{ // 120
			"del b",
			"del b\n",
			"del b\n",
		},
		{ // 121
			"return a",
			"return a\n",
			"return a\n",
		},
		{ // 122
			"yield a",
			"yield a\n",
			"yield a\n",
		},
		{ // 123
			"raise a",
			"raise a\n",
			"raise a\n",
		},
		{ // 124
			"import a",
			"import a\n",
			"import a\n",
		},
		{ // 125
			"global a",
			"global a\n",
			"global a\n",
		},
		{ // 126
			"nonlocal a",
			"nonlocal a\n",
			"nonlocal a\n",
		},
		{ // 127
			"type a = b",
			"type a = b\n",
			"type a = b\n",
		},
		{ // 128
			"a = b",
			"a=b\n",
			"a = b\n",
		},
		{ // 129
			"a: b = c",
			"a:b=c\n",
			"a: b = c\n",
		},
		{ // 130
			"a += b",
			"a+=b\n",
			"a += b\n",
		},
		{ // 131
			"pass",
			"pass\n",
			"pass\n",
		},
		{ // 132
			"break",
			"break\n",
			"break\n",
		},
		{ // 133
			"continue",
			"continue\n",
			"continue\n",
		},
		{ // 134
			"a[b]",
			"a[b]\n",
			"a[b]\n",
		},
		{ // 135
			"a [ b : c ] ",
			"a[b:c]\n",
			"a[b : c]\n",
		},
		{ // 136
			"a[ b : c : d]",
			"a[b:c:d]\n",
			"a[b : c : d]\n",
		},
		{ // 137
			"a[ b,c ]",
			"a[b,c]\n",
			"a[b, c]\n",
		},
		{ // 138
			"a[ b,c ,d]",
			"a[b,c,d]\n",
			"a[b, c, d]\n",
		},
		{ // 139
			"a = b",
			"a=b\n",
			"a = b\n",
		},
		{ // 140
			"a = *b",
			"a=*b\n",
			"a = *b\n",
		},
		{ // 141
			"a = *b, c",
			"a=*b,c\n",
			"a = *b, c\n",
		},
		{ // 142
			"a = b ,",
			"a=b,\n",
			"a = b,\n",
		},
		{ // 143
			"a = *b,",
			"a=*b,\n",
			"a = *b,\n",
		},
		{ // 144
			"a = *b, c",
			"a=*b,c\n",
			"a = *b, c\n",
		},
		{ // 145
			"a(*b)",
			"a(*b)\n",
			"a(*b)\n",
		},
		{ // 146
			"a(*b, c)",
			"a(*b,c)\n",
			"a(*b, c)\n",
		},
		{ // 147
			"a(*b, *c)",
			"a(*b,*c)\n",
			"a(*b, *c)\n",
		},
		{ // 148
			"a(*b, c = d)",
			"a(*b,c=d)\n",
			"a(*b, c = d)\n",
		},
		{ // 149
			"a",
			"a\n",
			"a\n",
		},
		{ // 150
			"if a: b",
			"if a:b\n",
			"if a: b\n",
		},
		{ // 151
			"a;b",
			"a;b\n",
			"a; b\n",
		},
		{ // 152
			"if a: \n\tb",
			"if a:\n\tb\n",
			"if a:\n\tb\n",
		},
		{ // 153
			"if a: \n\tb\n\tc",
			"if a:\n\tb\n\tc\n",
			"if a:\n\tb\n\tc\n",
		},
		{ // 154
			"if a:\n\t(\nb\n)",
			"if a:\n\t(b)\n",
			"if a:\n\t(b)\n",
		},
		{ // 155
			"if a:\n\tif b:\n\t\tc\n\t\td",
			"if a:\n\tif b:\n\t\tc\n\t\td\n",
			"if a:\n\tif b:\n\t\tc\n\t\td\n",
		},
		{ // 156
			"a = b",
			"a=b\n",
			"a = b\n",
		},
		{ // 157
			"a.b = c",
			"a.b=c\n",
			"a.b = c\n",
		},
		{ // 158
			"(a) = b",
			"(a)=b\n",
			"(a) = b\n",
		},
		{ // 159
			"[a] = b",
			"[a]=b\n",
			"[a] = b\n",
		},
		{ // 160
			"*a = b",
			"*a=b\n",
			"*a = b\n",
		},
		{ // 161
			"a, b = c",
			"a,b=c\n",
			"a, b = c\n",
		},
	} {
		for m, input := range test {
			tk := parser.NewStringTokeniser(input)

			if f, err := Parse(&tk); err != nil {
				t.Errorf("test %d.%d: unexpected error: %s", n+1, m+1, err)
			} else if simple := fmt.Sprintf("%s", f); simple != test[1] {
				t.Errorf("test %d.%d.1: expecting output %q, got %q", n+1, m+1, test[1], simple)
			} else if verbose := fmt.Sprintf("%+s", f); verbose != test[2] {
				t.Errorf("test %d.%d.2: expecting output %q, got %q", n+1, m+1, test[2], verbose)
			}
		}
	}
}

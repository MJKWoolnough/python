package python

import (
	"fmt"
	"strings"
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
			"def a(): b : c = yield d",
			"def a():b:c=yield d\n",
			"def a(): b: c = yield d\n",
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
			"def a(): b=yield c",
			"def a():b=yield c\n",
			"def a(): b = yield c\n",
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
			"try:a\nexcept b as c:d",
			"try:a\nexcept b as c:d\n",
			"try: a\nexcept b as c: d\n",
		},
		{ // 37
			"try:a\nexcept b, c as d:e",
			"try:a\nexcept b, c as d:e\n",
			"try: a\nexcept b, c as d: e\n",
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
			"def a(): yield b",
			"def a():yield b\n",
			"def a(): yield b\n",
		},
		{ // 41
			"def a(): yield b,c",
			"def a():yield b,c\n",
			"def a(): yield b, c\n",
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
			"a+=b",
			"a+=b\n",
			"a += b\n",
		},
		{ // 49
			"def a(): b -= yield c",
			"def a():b-=yield c\n",
			"def a(): b -= yield c\n",
		},
		{ // 50
			"try:a\nexcept b:c",
			"try:a\nexcept b:c\n",
			"try: a\nexcept b: c\n",
		},
		{ // 51
			"for a in b:c",
			"for a in b:c\n",
			"for a in b: c\n",
		},
		{ // 52
			"async for a in b:c",
			"async for a in b:c\n",
			"async for a in b: c\n",
		},
		{ // 53
			"for a in b:c\nelse:d",
			"for a in b:c\nelse:d\n",
			"for a in b: c\nelse: d\n",
		},
		{ // 54
			"def a():b",
			"def a():b\n",
			"def a(): b\n",
		},
		{ // 55
			"def a():\n\tb",
			"def a():\n\tb\n",
			"def a():\n\tb\n",
		},
		{ // 56
			"def a [b] ():c",
			"def a[b]():c\n",
			"def a[b](): c\n",
		},
		{ // 57
			"def a [] ():c",
			"def a[]():c\n",
			"def a[](): c\n",
		},
		{ // 58
			"def a() -> b:c",
			"def a()->b:c\n",
			"def a() -> b: c\n",
		},
		{ // 59
			"(a for b in c)",
			"(a for b in c)\n",
			"(a for b in c)\n",
		},
		{ // 60
			"global a",
			"global a\n",
			"global a\n",
		},
		{ // 61
			"global a , b",
			"global a,b\n",
			"global a, b\n",
		},
		{ // 62
			"if a:b",
			"if a:b\n",
			"if a: b\n",
		},
		{ // 63
			"if a :b\nelse : c",
			"if a:b\nelse:c\n",
			"if a: b\nelse: c\n",
		},
		{ // 64
			"if a :b\nelif c : d",
			"if a:b\nelif c:d\n",
			"if a: b\nelif c: d\n",
		},
		{ // 65
			"if a :b\nelif c : d\nelif e:f",
			"if a:b\nelif c:d\nelif e:f\n",
			"if a: b\nelif c: d\nelif e: f\n",
		},
		{ // 66
			"if a :b\nelif c : d\nelif e:f\nelse:g",
			"if a:b\nelif c:d\nelif e:f\nelse:g\n",
			"if a: b\nelif c: d\nelif e: f\nelse: g\n",
		},
		{ // 67
			"from a import b",
			"from a import b\n",
			"from a import b\n",
		},
		{ // 68
			"lazy from a import b",
			"lazy from a import b\n",
			"lazy from a import b\n",
		},
		{ // 69
			"from a import b,c",
			"from a import b,c\n",
			"from a import b, c\n",
		},
		{ // 70
			"import a",
			"import a\n",
			"import a\n",
		},
		{ // 71
			"lazy import a",
			"lazy import a\n",
			"lazy import a\n",
		},
		{ // 72
			"lazy = 1",
			"lazy=1\n",
			"lazy = 1\n",
		},
		{ // 73
			"import a , b",
			"import a,b\n",
			"import a, b\n",
		},
		{ // 74
			"class a(**b):c",
			"class a(**b):c\n",
			"class a(**b): c\n",
		},
		{ // 75
			"class a ( **b, c = d) : e",
			"class a(**b,c=d):e\n",
			"class a(**b, c = d): e\n",
		},
		{ // 76
			"@a\nclass b:c",
			"@a\nclass b:c\n",
			"@a\nclass b: c\n",
		},
		{ // 77
			"@a\n@b\nclass c:d",
			"@a\n@b\nclass c:d\n",
			"@a\n@b\nclass c: d\n",
		},
		{ // 78
			"@a\n# A\n\n# B\n@b\nclass c:d",
			"@a\n@b\nclass c:d\n",
			"@a\n# A\n\n# B\n@b\nclass c: d\n",
		},
		{ // 79
			"a == b",
			"a==b\n",
			"a == b\n",
		},
		{ // 80
			"a == b == c",
			"a==b==c\n",
			"a == b == c\n",
		},
		{ // 81
			"a is b",
			"a is b\n",
			"a is b\n",
		},
		{ // 82
			"a is not b",
			"a is not b\n",
			"a is not b\n",
		},
		{ // 83
			"a in b",
			"a in b\n",
			"a in b\n",
		},
		{ // 84
			"a not in b",
			"a not in b\n",
			"a not in b\n",
		},
		{ // 85
			"assert lambda:a",
			"assert lambda:a\n",
			"assert lambda: a\n",
		},
		{ // 86
			"assert lambda a : b",
			"assert lambda a:b\n",
			"assert lambda a: b\n",
		},
		{ // 87
			"import a as b",
			"import a as b\n",
			"import a as b\n",
		},
		{ // 88
			"import a.b",
			"import a.b\n",
			"import a.b\n",
		},
		{ // 89
			"import a.b.c",
			"import a.b.c\n",
			"import a.b.c\n",
		},
		{ // 90
			"a*b",
			"a*b\n",
			"a * b\n",
		},
		{ // 91
			"a / b",
			"a/b\n",
			"a / b\n",
		},
		{ // 92
			"a * b / c*d",
			"a*b/c*d\n",
			"a * b / c * d\n",
		},
		{ // 93
			"nonlocal a",
			"nonlocal a\n",
			"nonlocal a\n",
		},
		{ // 94
			"nonlocal a,b , c",
			"nonlocal a,b,c\n",
			"nonlocal a, b, c\n",
		},
		{ // 95
			"not a",
			"not a\n",
			"not a\n",
		},
		{ // 96
			"not not not not a",
			"not not not not a\n",
			"not not not not a\n",
		},
		{ // 97
			"a|b",
			"a|b\n",
			"a | b\n",
		},
		{ // 98
			"a|b | c",
			"a|b|c\n",
			"a | b | c\n",
		},
		{ // 99
			"a or b",
			"a or b\n",
			"a or b\n",
		},
		{ // 100
			"a or b or c",
			"a or b or c\n",
			"a or b or c\n",
		},
		{ // 101
			"def a():b",
			"def a():b\n",
			"def a(): b\n",
		},
		{ // 102
			"def a(b):c",
			"def a(b):c\n",
			"def a(b): c\n",
		},
		{ // 103
			"def a(b,/,c):d",
			"def a(b,/,c):d\n",
			"def a(b, /, c): d\n",
		},
		{ // 104
			"def a(b,c,/,d,e):f",
			"def a(b,c,/,d,e):f\n",
			"def a(b, c, /, d, e): f\n",
		},
		{ // 105
			"def a(b, *c):d",
			"def a(b,*c):d\n",
			"def a(b, *c): d\n",
		},
		{ // 106
			"def a(b, *c, d):e",
			"def a(b,*c,d):e\n",
			"def a(b, *c, d): e\n",
		},
		{ // 107
			"def a(b,**c):d",
			"def a(b,**c):d\n",
			"def a(b, **c): d\n",
		},
		{ // 108
			"def a(b , / , *c):d",
			"def a(b,/,*c):d\n",
			"def a(b, /, *c): d\n",
		},
		{ // 109
			"def a(b , / , *c,d):e",
			"def a(b,/,*c,d):e\n",
			"def a(b, /, *c, d): e\n",
		},
		{ // 110
			"def a(b , / , **c):d",
			"def a(b,/,**c):d\n",
			"def a(b, /, **c): d\n",
		},
		{ // 111
			"def a(b , / , c, d, *e, f, g, **h):i",
			"def a(b,/,c,d,*e,f,g,**h):i\n",
			"def a(b, /, c, d, *e, f, g, **h): i\n",
		},
		{ // 112
			"def a(b , c, d, *e, f, g, **h):i",
			"def a(b,c,d,*e,f,g,**h):i\n",
			"def a(b, c, d, *e, f, g, **h): i\n",
		},
		{ // 113
			"def a(*b, **c): d",
			"def a(*b,**c):d\n",
			"def a(*b, **c): d\n",
		},
		{ // 114
			"def a(**b):c",
			"def a(**b):c\n",
			"def a(**b): c\n",
		},
		{ // 115
			"def a(b = c): d",
			"def a(b=c):d\n",
			"def a(b = c): d\n",
		},
		{ // 116
			"def a(b:c = d): e",
			"def a(b:c=d):e\n",
			"def a(b: c = d): e\n",
		},
		{ // 117
			"async def a(b = c): d",
			"async def a(b=c):d\n",
			"async def a(b = c): d\n",
		},
		{ // 118
			"a( b )",
			"a(b)\n",
			"a(b)\n",
		},
		{ // 119
			"a( *b )",
			"a(*b)\n",
			"a(*b)\n",
		},
		{ // 120
			"a(b=c, *d)",
			"a(b=c,*d)\n",
			"a(b = c, *d)\n",
		},
		{ // 121
			"a ** b",
			"a**b\n",
			"a ** b\n",
		},
		{ // 122
			"await a**b",
			"await a**b\n",
			"await a ** b\n",
		},
		{ // 123
			"a",
			"a\n",
			"a\n",
		},
		{ // 124
			"a.b",
			"a.b\n",
			"a.b\n",
		},
		{ // 125
			"a . b",
			"a.b\n",
			"a.b\n",
		},
		{ // 126
			"a[b]",
			"a[b]\n",
			"a[b]\n",
		},
		{ // 127
			"a[ b ]",
			"a[b]\n",
			"a[b]\n",
		},
		{ // 128
			"a(b)",
			"a(b)\n",
			"a(b)\n",
		},
		{ // 129
			"a( b )",
			"a(b)\n",
			"a(b)\n",
		},
		{ // 130
			"raise",
			"raise\n",
			"raise\n",
		},
		{ // 131
			"raise a",
			"raise a\n",
			"raise a\n",
		},
		{ // 132
			"raise a from b",
			"raise a from b\n",
			"raise a from b\n",
		},
		{ // 133
			"from . a import b",
			"from .a import b\n",
			"from .a import b\n",
		},
		{ // 134
			"from ..a import b",
			"from ..a import b\n",
			"from ..a import b\n",
		},
		{ // 135
			"from ... import a",
			"from ... import a\n",
			"from ... import a\n",
		},
		{ // 136
			"from ....a.b import c",
			"from ....a.b import c\n",
			"from ....a.b import c\n",
		},
		{ // 137
			"def a():\n\treturn",
			"def a():\n\treturn\n",
			"def a():\n\treturn\n",
		},
		{ // 138
			"def a():\n\treturn b",
			"def a():\n\treturn b\n",
			"def a():\n\treturn b\n",
		},
		{ // 139
			"a>>b",
			"a>>b\n",
			"a >> b\n",
		},
		{ // 140
			"a << b",
			"a<<b\n",
			"a << b\n",
		},
		{ // 141
			"assert a",
			"assert a\n",
			"assert a\n",
		},
		{ // 142
			"del b",
			"del b\n",
			"del b\n",
		},
		{ // 143
			"def a():\n\treturn a",
			"def a():\n\treturn a\n",
			"def a():\n\treturn a\n",
		},
		{ // 144
			"def a(): yield b",
			"def a():yield b\n",
			"def a(): yield b\n",
		},
		{ // 145
			"raise a",
			"raise a\n",
			"raise a\n",
		},
		{ // 146
			"import a",
			"import a\n",
			"import a\n",
		},
		{ // 147
			"global a",
			"global a\n",
			"global a\n",
		},
		{ // 148
			"nonlocal a",
			"nonlocal a\n",
			"nonlocal a\n",
		},
		{ // 149
			"type a = b",
			"type a=b\n",
			"type a = b\n",
		},
		{ // 150
			"a = b",
			"a=b\n",
			"a = b\n",
		},
		{ // 151
			"a: b = c",
			"a:b=c\n",
			"a: b = c\n",
		},
		{ // 152
			"a += b",
			"a+=b\n",
			"a += b\n",
		},
		{ // 153
			"pass",
			"pass\n",
			"pass\n",
		},
		{ // 154
			"while a: break",
			"while a:break\n",
			"while a: break\n",
		},
		{ // 155
			"while a: continue",
			"while a:continue\n",
			"while a: continue\n",
		},
		{ // 156
			"a[b]",
			"a[b]\n",
			"a[b]\n",
		},
		{ // 157
			"a [ b : c ] ",
			"a[b:c]\n",
			"a[b : c]\n",
		},
		{ // 158
			"a[ b : c : d]",
			"a[b:c:d]\n",
			"a[b : c : d]\n",
		},
		{ // 159
			"a[ b,c ]",
			"a[b,c]\n",
			"a[b, c]\n",
		},
		{ // 160
			"a[ b,c ,d]",
			"a[b,c,d]\n",
			"a[b, c, d]\n",
		},
		{ // 161
			"a = b",
			"a=b\n",
			"a = b\n",
		},
		{ // 162
			"a = *b",
			"a=*b\n",
			"a = *b\n",
		},
		{ // 163
			"a = *b, c",
			"a=*b,c\n",
			"a = *b, c\n",
		},
		{ // 164
			"a = b ,",
			"a=b,\n",
			"a = b,\n",
		},
		{ // 165
			"a = *b,",
			"a=*b,\n",
			"a = *b,\n",
		},
		{ // 166
			"a = *b, c",
			"a=*b,c\n",
			"a = *b, c\n",
		},
		{ // 167
			"a(*b)",
			"a(*b)\n",
			"a(*b)\n",
		},
		{ // 168
			"a(*b, c)",
			"a(*b,c)\n",
			"a(*b, c)\n",
		},
		{ // 169
			"a(*b, *c)",
			"a(*b,*c)\n",
			"a(*b, *c)\n",
		},
		{ // 170
			"a(*b, c = d)",
			"a(*b,c=d)\n",
			"a(*b, c = d)\n",
		},
		{ // 171
			"a",
			"a\n",
			"a\n",
		},
		{ // 172
			"if a: b",
			"if a:b\n",
			"if a: b\n",
		},
		{ // 173
			"a;b",
			"a;b\n",
			"a; b\n",
		},
		{ // 174
			"if a: \n\tb",
			"if a:\n\tb\n",
			"if a:\n\tb\n",
		},
		{ // 175
			"if a: \n\tb\n\tc",
			"if a:\n\tb\n\tc\n",
			"if a:\n\tb\n\tc\n",
		},
		{ // 176
			"if a:\n\t(\nb\n)",
			"if a:\n\t(b)\n",
			"if a:\n\t(b)\n",
		},
		{ // 177
			"if a:\n\tif b:\n\t\tc\n\t\td",
			"if a:\n\tif b:\n\t\tc\n\t\td\n",
			"if a:\n\tif b:\n\t\tc\n\t\td\n",
		},
		{ // 178
			"a = b",
			"a=b\n",
			"a = b\n",
		},
		{ // 179
			"a.b = c",
			"a.b=c\n",
			"a.b = c\n",
		},
		{ // 180
			"(a) = b",
			"(a)=b\n",
			"(a) = b\n",
		},
		{ // 181
			"{a: b}",
			"{a:b}\n",
			"{a: b}\n",
		},
		{ // 182
			"{a: b,c: d}",
			"{a:b,c:d}\n",
			"{a: b, c: d}\n",
		},
		{ // 183
			"{a:b for c in d}",
			"{a:b for c in d}\n",
			"{a: b for c in d}\n",
		},
		{ // 184
			"{**a}",
			"{**a}\n",
			"{**a}\n",
		},
		{ // 185
			"[a] = b",
			"[a]=b\n",
			"[a] = b\n",
		},
		{ // 186
			"*a = b",
			"*a=b\n",
			"*a = b\n",
		},
		{ // 187
			"a, b = c",
			"a,b=c\n",
			"a, b = c\n",
		},
		{ // 188
			"try:a\nexcept b:c",
			"try:a\nexcept b:c\n",
			"try: a\nexcept b: c\n",
		},
		{ // 189
			"try:a\nexcept b, c:d",
			"try:a\nexcept b, c:d\n",
			"try: a\nexcept b, c: d\n",
		},
		{ // 190
			"try:a\nexcept b:c\nexcept d:e",
			"try:a\nexcept b:c\nexcept d:e\n",
			"try: a\nexcept b: c\nexcept d: e\n",
		},
		{ // 191
			"try:a\nexcept *b:c",
			"try:a\nexcept *b:c\n",
			"try: a\nexcept *b: c\n",
		},
		{ // 192
			"try:a\nexcept *b, c:d",
			"try:a\nexcept *b, c:d\n",
			"try: a\nexcept *b, c: d\n",
		},
		{ // 193
			"try:a\nexcept *b:c\nexcept *d:e",
			"try:a\nexcept *b:c\nexcept *d:e\n",
			"try: a\nexcept *b: c\nexcept *d: e\n",
		},
		{ // 194
			"try:a\nexcept b:c\nelse: d",
			"try:a\nexcept b:c\nelse:d\n",
			"try: a\nexcept b: c\nelse: d\n",
		},
		{ // 195
			"try:a\nexcept b:c\nfinally: d",
			"try:a\nexcept b:c\nfinally:d\n",
			"try: a\nexcept b: c\nfinally: d\n",
		},
		{ // 196
			"try:a\nexcept b:c\nelse: d\nfinally:e",
			"try:a\nexcept b:c\nelse:d\nfinally:e\n",
			"try: a\nexcept b: c\nelse: d\nfinally: e\n",
		},
		{ // 197
			"def a[b](): c",
			"def a[b]():c\n",
			"def a[b](): c\n",
		},
		{ // 198
			"def a[b:c](): d",
			"def a[b:c]():d\n",
			"def a[b: c](): d\n",
		},
		{ // 199
			"def a[*b](): c",
			"def a[*b]():c\n",
			"def a[*b](): c\n",
		},
		{ // 200
			"def a[**b](): c",
			"def a[**b]():c\n",
			"def a[**b](): c\n",
		},
		{ // 201
			"class a[b,c, d ](): e",
			"class a[b,c,d]():e\n",
			"class a[b, c, d](): e\n",
		},
		{ // 202
			"type a = b",
			"type a=b\n",
			"type a = b\n",
		},
		{ // 203
			"type a[b] = c",
			"type a[b]=c\n",
			"type a[b] = c\n",
		},
		{ // 204
			"+a",
			"+a\n",
			"+a\n",
		},
		{ // 205
			"-a",
			"-a\n",
			"-a\n",
		},
		{ // 206
			"~a",
			"~a\n",
			"~a\n",
		},
		{ // 207
			"while a:b",
			"while a:b\n",
			"while a: b\n",
		},
		{ // 208
			"while a:b\nelse: c",
			"while a:b\nelse:c\n",
			"while a: b\nelse: c\n",
		},
		{ // 209
			"with a: b",
			"with a:b\n",
			"with a: b\n",
		},
		{ // 210
			"with a as b:c",
			"with a as b:c\n",
			"with a as b: c\n",
		},
		{ // 211
			"with a,b: c",
			"with a,b:c\n",
			"with a, b: c\n",
		},
		{ // 212
			"with a as b, c,d as e:f",
			"with a as b,c,d as e:f\n",
			"with a as b, c, d as e: f\n",
		},
		{ // 213
			"a^b",
			"a^b\n",
			"a ^ b\n",
		},
		{ // 214
			"def a(): yield b",
			"def a():yield b\n",
			"def a(): yield b\n",
		},
		{ // 215
			"def a(): yield from b",
			"def a():yield from b\n",
			"def a(): yield from b\n",
		},
		{ // 216
			"(a for b in c if d)",
			"(a for b in c if d)\n",
			"(a for b in c if d)\n",
		},
		{ // 217
			"(a async for b in c if d)",
			"(a async for b in c if d)\n",
			"(a async for b in c if d)\n",
		},
		{ // 218
			"(a for b in c if d for e in f)",
			"(a for b in c if d for e in f)\n",
			"(a for b in c if d for e in f)\n",
		},
		{ // 219
			"def a(): (yield b)",
			"def a():(yield b)\n",
			"def a(): (yield b)\n",
		},
		{ // 220
			"{a for b in c}",
			"{a for b in c}\n",
			"{a for b in c}\n",
		},
		{ // 221
			"a if b else c",
			"a if b else c\n",
			"a if b else c\n",
		},
		{ // 222
			"if a:\n\t\"\"",
			"if a:\n\t\"\"\n",
			"if a:\n\t\"\"\n",
		},
		{ // 223
			"if a:\n\t\"\"\"a\nb\"\"\"",
			"if a:\n\t\"\"\"a\nb\"\"\"\n",
			"if a:\n\t\"\"\"a\nb\"\"\"\n",
		},
		{ // 224
			"a\n# A Comment",
			"a\n",
			"a\n\n# A Comment\n",
		},
		{ // 225
			"# A comment\na",
			"a\n",
			"# A comment\na\n",
		},
		{ // 226
			"a # A comment",
			"a\n",
			"a # A comment\n",
		},
		{ // 227
			"a # A comment\n# B comment\n\n# EOF Comment",
			"a\n",
			"a # A comment\n  # B comment\n\n# EOF Comment\n",
		},
		{ // 228
			"while a: # A comment\n# B comment\n\t#abc\n\tb #def\n\n#efg",
			"while a:\n\tb\n",
			"while a: # A comment\n         # B comment\n         #abc\n\tb #def\n\n\t#efg\n",
		},
		{ // 229
			"while a:\n# A comment\n# B comment\n\t#abc\n\tb #def\n\n#efg",
			"while a:\n\tb\n",
			"while a:\n\t# A comment\n\t# B comment\n\t#abc\n\tb #def\n\n\t#efg\n",
		},
		{ // 230
			"(#abc\n)",
			"()\n",
			"( #abc\n)\n",
		},
		{ // 231
			"[#abc\n]",
			"[]\n",
			"[ #abc\n]\n",
		},
		{ // 232
			"[#abc\n]",
			"[]\n",
			"[ #abc\n]\n",
		},
		{ // 233
			"( #abc\n)",
			"()\n",
			"( #abc\n)\n",
		},
		{ // 234
			"[ #abc\n]",
			"[]\n",
			"[ #abc\n]\n",
		},
		{ // 235
			"[ #abc\na\n#def\n]",
			"[a]\n",
			"[ #abc\n\ta\n#def\n]\n",
		},
		{ // 236
			"[ #abc\n]",
			"[]\n",
			"[ #abc\n]\n",
		},
		{ // 237
			"def a(#abc\n): b",
			"def a():b\n",
			"def a( #abc\n): b\n",
		},
		{ // 238
			"def a(# A\n# B\n\n# C\n\n# D\n): b",
			"def a():b\n",
			"def a( # A\n       # B\n\n\t# C\n\n\t# D\n): b\n",
		},
		{ // 239
			"[ #abc\na #def\n] = b",
			"[a]=b\n",
			"[ #abc\n\ta #def\n] = b\n",
		},
		{ // 240
			"[ # A\n* # B\na # C\n] = b",
			"[*a]=b\n",
			"[ # A\n\t* # B\n\ta # C\n] = b\n",
		},
		{ // 241
			"a\nb\n\nc\n\nd\n\n\n\n\ne",
			"a\nb\nc\nd\ne\n",
			"a\nb\n\nc\n\nd\n\ne\n",
		},
		{ // 242
			"if a:\n\tb\n\tc\n\t\n\t\n\td",
			"if a:\n\tb\n\tc\n\td\n",
			"if a:\n\tb\n\tc\n\n\td\n",
		},
		{ // 243
			"def a[b # A\n, # B\nc # C\n](): b",
			"def a[b,c]():b\n",
			"def a[b # A\n\t, # B\n\tc # C\n](): b\n",
		},
		{ // 244
			"def a[# A\n# B\n\n# C\nb, c # D\n# E\n\n# F\n# G\n\n](): b",
			"def a[b,c]():b\n",
			"def a[ # A\n       # B\n\n\t# C\n\tb, c # D\n\t     # E\n\n# F\n# G\n](): b\n",
		},
		{ // 245
			"def a(\n# A\nb = 1 # B\n): c",
			"def a(b=1):c\n",
			"def a(\n\t# A\n\tb = 1 # B\n): c\n",
		},
		{ // 246
			"def a(\n# A\nb = 1 # B\n, /, # C\nc # D\n): d",
			"def a(b=1,/,c):d\n",
			"def a(\n\t# A\n\tb = 1 # B\n\t, /, # C\n\tc # D\n): d\n",
		},
		{ // 247
			"def a(# A\n# B\n\n# C\n\n#D\nb # E\n\n#F\n, # G\n\n# H\n\n/# I\n\n# J\n, # K\n\n# L\n*# M\n\n# N\nc\n# O\n\n# P\n, # Q\n**# R\n\n# S\nd\n# T\n): e",
			"def a(b,/,*c,**d):e\n",
			"def a( # A\n\t       # B\n\n\t# C\n\n\t#D\n\tb # E\n\n\t#F\n\t, # G\n\n\t# H\n\t/ # I\n\n\t# J\n\t, # K\n\n\t# L\n\t* # M\n\n\t# N\n\tc # O\n\n\t# P\n\t, # Q\n\t** # R\n\n\t# S\n\td\n\t# T\n): e\n",
		},
		{ // 248
			"def a( # A\n\n# B\nb = 1): c",
			"def a(b=1):c\n",
			"def a( # A\n\n\t# B\n\tb = 1): c\n",
		},
		{ // 249
			"def a( # A\n\n# B\n*b # C\n, c): d",
			"def a(*b,c):d\n",
			"def a( # A\n\n\t# B\n\t*b # C\n\t, c): d\n",
		},
		{ // 250
			"def a( # A\n\n# B\n**b # C\n\n# D\n): d",
			"def a(**b):d\n",
			"def a( # A\n\n\t# B\n\t**b # C\n\n\t# D\n): d\n",
		},
		{ // 251
			"def a(): ( # A\n\n #B\nyield # C\nb #D\n)",
			"def a():(yield b)\n",
			"def a(): ( # A\n\n\t#B\n\tyield # C\n\tb #D\n)\n",
		},
		{ // 252
			"def a(): ( # A\n\n # B\nyield # C\nb # D\n, # E\n)",
			"def a():(yield b)\n",
			"def a(): ( # A\n\n\t# B\n\tyield # C\n\tb # D\n\t, # E\n)\n",
		},
		{ // 253
			"def a(): ( # A\n\n # B\nyield # C\nfrom # D\nb # E\n\n# F\n)",
			"def a():(yield from b)\n",
			"def a(): ( # A\n\n\t# B\n\tyield # C\n\tfrom # D\n\tb # E\n\n# F\n)\n",
		},
		{ // 254
			"( # A\n\n # B\na # C\nfor b in c # D\n\n# E\n)",
			"(a for b in c)\n",
			"( # A\n\n\t# B\n\ta # C\n\tfor b in c # D\n\n\t# E\n)\n",
		},
		{ // 255
			"( # A\n\n # B\na # C\n\n# D\n)",
			"(a)\n",
			"( # A\n\n\t# B\n\ta # C\n\n# D\n)\n",
		},
		{ // 256
			"( # A\n\n # B\n*a ,# C\n\n# D\n)",
			"(*a,)\n",
			"( # A\n\n\t# B\n\t*a, # C\n\n# D\n)\n",
		},
		{ // 257
			"( # A\n\n# B\na # C\n, # D\n* # E\nb # F\n\n# G\n)",
			"(a,*b)\n",
			"( # A\n\n\t# B\n\ta # C\n\t, # D\n\t* # E\n\tb # F\n\n# G\n)\n",
		},
		{ // 258
			"( # A\n\n# B\na # C\nfor # D\nb # E\nin # F\nc # G\n\n# H\n)",
			"(a for b in c)\n",
			"( # A\n\n\t# B\n\ta # C\n\tfor # D\n\tb # E\n\tin # F\n\tc # G\n\n\t# H\n)\n",
		},
		{ // 259
			"( # A\n\n# B\na # C\nasync # D\nfor # E\nb # F\nin # G\nc # H\n\n# I\n)",
			"(a async for b in c)\n",
			"( # A\n\n\t# B\n\ta # C\n\tasync # D\n\tfor # E\n\tb # F\n\tin # G\n\tc # H\n\n\t# I\n)\n",
		},
		{ // 260
			"( # A\n\n# B\na # C\nfor # D\nb # E\nin # F\nc # G\nif # H\nd # I\n\n# J\n)",
			"(a for b in c if d)\n",
			"( # A\n\n\t# B\n\ta # C\n\tfor # D\n\tb # E\n\tin # F\n\tc # G\n\tif # H\n\td # I\n\n\t# J\n)\n",
		},
		{ // 261
			"(a # A\nor # B\nb)",
			"(a or b)\n",
			"(a # A\n\tor # B\n\tb)\n",
		},
		{ // 262
			"(a # A\nand # B\nb)",
			"(a and b)\n",
			"(a # A\n\tand # B\n\tb)\n",
		},
		{ // 263
			"(not # A\na)",
			"(not a)\n",
			"(not # A\n\ta)\n",
		},
		{ // 264
			"(not # A\nnot not # B\na)",
			"(not not not a)\n",
			"(not # A\n\tnot not # B\n\ta)\n",
		},
		{ // 265
			"(a # A\n== # B\nb)",
			"(a==b)\n",
			"(a # A\n\t== # B\n\tb)\n",
		},
		{ // 266
			"(a # A\nin # B\nb)",
			"(a in b)\n",
			"(a # A\n\tin # B\n\tb)\n",
		},
		{ // 267
			"(a # A\nnot # B\nin # C\nb)",
			"(a not in b)\n",
			"(a # A\n\tnot # B\n\tin # C\n\tb)\n",
		},
		{ // 268
			"(a # A\nis # B\nb)",
			"(a is b)\n",
			"(a # A\n\tis # B\n\tb)\n",
		},
		{ // 269
			"(a # A\nis # B\nnot # C\nb)",
			"(a is not b)\n",
			"(a # A\n\tis # B\n\tnot # C\n\tb)\n",
		},
		{ // 270
			"(a # A\n| # B\nb)",
			"(a|b)\n",
			"(a # A\n\t| # B\n\tb)\n",
		},
		{ // 271
			"(a # A\n^ # B\nb)",
			"(a^b)\n",
			"(a # A\n\t^ # B\n\tb)\n",
		},
		{ // 272
			"(a # A\n& # B\nb)",
			"(a&b)\n",
			"(a # A\n\t& # B\n\tb)\n",
		},
		{ // 273
			"(a # A\n<< # B\nb)",
			"(a<<b)\n",
			"(a # A\n\t<< # B\n\tb)\n",
		},
		{ // 274
			"(a # A\n>> # B\nb)",
			"(a>>b)\n",
			"(a # A\n\t>> # B\n\tb)\n",
		},
		{ // 275
			"(a # A\n+ # B\nb)",
			"(a+b)\n",
			"(a # A\n\t+ # B\n\tb)\n",
		},
		{ // 276
			"(a # A\n- # B\nb)",
			"(a-b)\n",
			"(a # A\n\t- # B\n\tb)\n",
		},
		{ // 277
			"(a # A\n* # B\nb)",
			"(a*b)\n",
			"(a # A\n\t* # B\n\tb)\n",
		},
		{ // 278
			"(a # A\n// # B\nb)",
			"(a//b)\n",
			"(a # A\n\t// # B\n\tb)\n",
		},
		{ // 279
			"(- # A\na)",
			"(-a)\n",
			"(- # A\n\ta)\n",
		},
		{ // 280
			"(await # A\na)",
			"(await a)\n",
			"(await # A\n\ta)\n",
		},
		{ // 281
			"(a # A\n** # B\nb)",
			"(a**b)\n",
			"(a # A\n\t** # B\n\tb)\n",
		},
		{ // 282
			"(await # A\na # B\n** # C\nb)",
			"(await a**b)\n",
			"(await # A\n\ta # B\n\t** # C\n\tb)\n",
		},
		{ // 283
			"(a # A\n. # B\nb)",
			"(a.b)\n",
			"(a # A\n\t. # B\n\tb)\n",
		},
		{ // 284
			"(a # A\n[b])",
			"(a[b])\n",
			"(a # A\n\t[b])\n",
		},
		{ // 285
			"(a # A\n(b))",
			"(a(b))\n",
			"(a # A\n\t(b))\n",
		},
		{ // 286
			"a[ # A\nb\n# B\n]",
			"a[b]\n",
			"a[ # A\n\tb\n# B\n]\n",
		},
		{ // 287
			"a[ # A\nb, # B\n]",
			"a[b]\n",
			"a[ # A\n\tb\n# B\n]\n",
		},
		{ // 288
			"a[ # A\nb, c\n# B\n]",
			"a[b,c]\n",
			"a[ # A\n\tb, c\n# B\n]\n",
		},
		{ // 289
			"a[ # A\n\n# B\n b # C\n: # D\n c # E\n: # F\nd # G\n\n# H\n]",
			"a[b:c:d]\n",
			"a[ # A\n\n\t# B\n\tb # C\n\t: # D\n\tc # E\n\t: # F\n\td # G\n\n# H\n]\n",
		},
		{ // 290
			"a[ # A\n\n# B\n b # C\n\n# D\n]",
			"a[b]\n",
			"a[ # A\n\n\t# B\n\tb # C\n\n# D\n]\n",
		},
		{ // 291
			"(a # A\nif # B\nb # C\nelse # D\nc)",
			"(a if b else c)\n",
			"(a # A\n\tif # B\n\tb # C\n\telse # D\n\tc)\n",
		},
		{ // 292
			"(# A\n\n# B\nlambda # C\n: # D\na # E\n\n# F\n)",
			"(lambda:a)\n",
			"( # A\n\n\t# B\n\tlambda # C\n\t: # D\n\ta # E\n\n# F\n)\n",
		},
		{ // 293
			"(# A\n\n# B\nlambda # C\na # D\n\n# E\n: # F\nb # G\n\n# H\n)",
			"(lambda a:b)\n",
			"( # A\n\n\t# B\n\tlambda # C\n\ta # D\n\n\t# E\n\t: # F\n\tb # G\n\n# H\n)\n",
		},
		{ // 294
			"{a # A\n:= # B\nb}",
			"{a:=b}\n",
			"{a # A\n\t:= # B\n\tb}\n",
		},
		{ // 295
			"{# A\n\n# B\na # C\n\n# D\n}",
			"{a}\n",
			"{ # A\n\n\t# B\n\ta # C\n\n# D\n}\n",
		},
		{ // 296
			"{# A\n\n# B\n*a # C\n\n# D\n}",
			"{*a}\n",
			"{ # A\n\n\t# B\n\t*a # C\n\n# D\n}\n",
		},
		{ // 297
			"{# A\n\n# B\na # C\n: # D\nb # E\n\n # F\n}",
			"{a:b}\n",
			"{ # A\n\n\t# B\n\ta # C\n\t: # D\n\tb # E\n\n# F\n}\n",
		},
		{ // 298
			"{# A\n\n# B\n** # C\na # D\n\n # F\n}",
			"{**a}\n",
			"{ # A\n\n\t# B\n\t** # C\n\ta # D\n\n# F\n}\n",
		},
		{ // 299
			"a(# A\nb\n# B\n)",
			"a(b)\n",
			"a( # A\n\tb\n# B\n)\n",
		},
		{ // 300
			"a(# A\n\n# B\nb=c # C\n\n# D\n)",
			"a(b=c)\n",
			"a( # A\n\n\t# B\n\tb = c # C\n\n# D\n)\n",
		},
		{ // 301
			"a(# A\n* # B\nb # C\n\n# D\n,)",
			"a(*b)\n",
			"a( # A\n\t* # B\n\tb # C\n\n\t# D\n\t,)\n",
		},
		{ // 302
			"def a(# A\n** # B\nb # C\n\n# D\n,): c",
			"def a(**b):c\n",
			"def a( # A\n\t** # B\n\tb # C\n\n\t# D\n): c\n",
		},
		{ // 303
			"{ # A\n\n # B\na # C\nfor b in c # D\n\n# E\n}",
			"{a for b in c}\n",
			"{ # A\n\n\t# B\n\ta # C\n\tfor b in c # D\n\n# E\n}\n",
		},
		{ // 304
			"a(# A\n\n# B\n** # C\nb # D\n\n# E\n)",
			"a(**b)\n",
			"a( # A\n\n\t# B\n\t** # C\n\tb # D\n\n# E\n)\n",
		},
		{ // 305
			"def a[# A\n\n# B\nb # C\n\n# D\n]():c",
			"def a[b]():c\n",
			"def a[ # A\n\n\t# B\n\tb # C\n\n# D\n](): c\n",
		},
		{ // 306
			"def a[# A\n\n# B\nb # C\n: # D\nc # E\n\n# F\n]():d",
			"def a[b:c]():d\n",
			"def a[ # A\n\n\t# B\n\tb # C\n\t: # D\n\tc # E\n\n# F\n](): d\n",
		},
		{ // 307
			"def a[# A\n\n# B\nb # C\n: # D\nc # E\n\n# F\n, # G\n** # H\nd # I\n]():e",
			"def a[b:c,**d]():e\n",
			"def a[ # A\n\n\t# B\n\tb # C\n\t: # D\n\tc # E\n\n\t# F\n\t, # G\n\t** # H\n\td # I\n](): e\n",
		},
		{ // 308
			"def a[# A\n\n# B\n*b # C\n\n# D\n]():c",
			"def a[*b]():c\n",
			"def a[ # A\n\n\t# B\n\t*b # C\n\n# D\n](): c\n",
		},
		{ // 309
			"class a( # A\n\n# B\n):b",
			"class a():b\n",
			"class a( # A\n\n# B\n): b\n",
		},
		{ // 310
			"class a( # A\n\nb\n# B\n):c",
			"class a(b):c\n",
			"class a( # A\nb\n# B\n): c\n",
		},
		{ // 311
			"with (# A\na,b\n# B\n): c",
			"with a,b:c\n",
			"with ( # A\na, b\n# B\n): c\n",
		},
		{ // 312
			"with (# A\n\n# B\na # C\n\n# D\n,# E\nb # F\nas # G\nc # H\n\n# I\n): d",
			"with a,b as c:d\n",
			"with ( # A\n\n\t# B\n\ta # C\n\n\t# D\n\t, \n\t# E\n\tb # F\n\tas # G\n\tc # H\n\n# I\n): d\n",
		},
		{ // 313
			"# A\n@a() # B\n# C\n\n# D\n@b()\ndef c():d",
			"@a()\n@b()\ndef c():d\n",
			"# A\n@a() # B\n     # C\n\n# D\n@b()\ndef c(): d\n",
		},
		{ // 314
			"# A\n@a # B\n# C\n\n# D\n@b # E\n\n# F\n\n# G\ndef c(# H\n):d",
			"@a\n@b\ndef c():d\n",
			"# A\n@a # B\n   # C\n\n# D\n@b # E\n\n# F\n\n# G\ndef c( # H\n): d\n",
		},
		{ // 315
			"# A\n@a # B\n# C\n\n# D\n@b # E\n\n# F\n\n# G\nclass c():d",
			"@a\n@b\nclass c():d\n",
			"# A\n@a # B\n   # C\n\n# D\n@b # E\n\n# F\n\n# G\nclass c(): d\n",
		},
		{ // 316
			"def a(b # A\n: # B\n c): d",
			"def a(b:c):d\n",
			"def a(b # A\n\t: # B\n\tc): d\n",
		},
		{ // 317
			"def a(b # A\n: # B\n c # C\n= # D\nd): e",
			"def a(b:c=d):e\n",
			"def a(b # A\n\t: # B\n\tc # C\n\t= # D\n\td): e\n",
		},
		{ // 318
			"a(b=c, * # A\nd)",
			"a(b=c,*d)\n",
			"a(b = c, * # A\n\td)\n",
		},
		{ // 319
			"from a import ( # A\nb\n# B\n)",
			"from a import b\n",
			"from a import ( # A\n\tb\n# B\n)\n",
		},
		{ // 320
			"from a import ( # A\nb # B\n. c # C\n\n# D\n)",
			"from a import b.c\n",
			"from a import ( # A\n\tb # B\n\t.c # C\n\n# D\n)\n",
		},
		{ // 321
			"from a import ( # A\nb # B\n. # C\nc # D\n\n# E\n)",
			"from a import b.c\n",
			"from a import ( # A\n\tb # B\n\t. # C\n\tc # D\n\n# E\n)\n",
		},
		{ // 322
			"from a import ( # A\nb # B\n. # C\nc # D\n\n# E\nas # F\nd # G\n\n# H\n)",
			"from a import b.c as d\n",
			"from a import ( # A\n\tb # B\n\t. # C\n\tc # D\n\n\t# E\n\t as # F\n\td # G\n\n# H\n)\n",
		},
		{ // 323
			"a = {**b for c in d}",
			"a={**b for c in d}\n",
			"a = {**b for c in d}\n",
		},
		{ // 324
			"a = {**b async for c in d}",
			"a={**b async for c in d}\n",
			"a = {**b async for c in d}\n",
		},
	} {
		for m, input := range test {
			tk := parser.NewStringTokeniser(input)

			if f, err := Parse(&tk); err != nil {
				t.Errorf("test %d.%d: unexpected error: %s", n+1, m+1, err)
			} else if simple := fmt.Sprintf("%s", f); simple != test[1] {
				t.Errorf("test %d.%d.1: expecting output %q, got %q", n+1, m+1, test[1], simple)
			} else if verbose := fmt.Sprintf("%+s", f); verbose != test[2] && (m != 1 || !strings.ContainsRune(test[2], '#') && !strings.Contains(test[2], "\n\n")) {
				t.Errorf("test %d.%d.2: expecting output %q, got %q", n+1, m+1, test[2], verbose)
			}
		}
	}
}

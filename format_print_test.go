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
			"try:a\nexcept b as c:d",
			"try:a\nexcept b as c:d\n",
			"try: a\nexcept b as c: d\n",
		},
		{ // 37
			"assert a",
			"assert a\n",
			"assert a\n",
		},
		{ // 38
			"assert lambda:a",
			"assert lambda:a\n",
			"assert lambda: a\n",
		},
		{ // 39
			"yield a",
			"yield a\n",
			"yield a\n",
		},
		{ // 40
			"yield a,b",
			"yield a,b\n",
			"yield a, b\n",
		},
		{ // 41
			"a\nb\nc",
			"a\nb\nc\n",
			"a\nb\nc\n",
		},
		{ // 42
			"[a]",
			"[a]\n",
			"[a]\n",
		},
		{ // 43
			"[a for b in c]",
			"[a for b in c]\n",
			"[a for b in c]\n",
		},
		{ // 44
			"[a,b]",
			"[a,b]\n",
			"[a, b]\n",
		},
		{ // 45
			"[a := b]",
			"[a:=b]\n",
			"[a := b]\n",
		},
		{ // 46
			"[* a]",
			"[*a]\n",
			"[*a]\n",
		},
		{ // 47
			"a+=b",
			"a+=b\n",
			"a += b\n",
		},
		{ // 48
			"a -= yield b",
			"a-=yield b\n",
			"a -= yield b\n",
		},
		{ // 49
			"try:a\nexcept b:c",
			"try:a\nexcept b:c\n",
			"try: a\nexcept b: c\n",
		},
		{ // 50
			"for a in b:c",
			"for a in b:c\n",
			"for a in b: c\n",
		},
		{ // 51
			"async for a in b:c",
			"async for a in b:c\n",
			"async for a in b: c\n",
		},
		{ // 52
			"for a in b:c\nelse:d",
			"for a in b:c\nelse:d\n",
			"for a in b: c\nelse: d\n",
		},
		{ // 53
			"def a():b",
			"def a():b\n",
			"def a(): b\n",
		},
		{ // 54
			"def a():\n\tb",
			"def a():\n\tb\n",
			"def a():\n\tb\n",
		},
		{ // 55
			"def a [b] ():c",
			"def a[b]():c\n",
			"def a[b](): c\n",
		},
		{ // 56
			"def a() -> b:c",
			"def a()->b:c\n",
			"def a() -> b: c\n",
		},
		{ // 57
			"(a for b in c)",
			"(a for b in c)\n",
			"(a for b in c)\n",
		},
		{ // 58
			"global a",
			"global a\n",
			"global a\n",
		},
		{ // 59
			"global a , b",
			"global a,b\n",
			"global a, b\n",
		},
		{ // 60
			"if a:b",
			"if a:b\n",
			"if a: b\n",
		},
		{ // 61
			"if a :b\nelse : c",
			"if a:b\nelse:c\n",
			"if a: b\nelse: c\n",
		},
		{ // 62
			"if a :b\nelif c : d",
			"if a:b\nelif c:d\n",
			"if a: b\nelif c: d\n",
		},
		{ // 63
			"if a :b\nelif c : d\nelif e:f",
			"if a:b\nelif c:d\nelif e:f\n",
			"if a: b\nelif c: d\nelif e: f\n",
		},
		{ // 64
			"if a :b\nelif c : d\nelif e:f\nelse:g",
			"if a:b\nelif c:d\nelif e:f\nelse:g\n",
			"if a: b\nelif c: d\nelif e: f\nelse: g\n",
		},
		{ // 65
			"from a import b",
			"from a import b\n",
			"from a import b\n",
		},
		{ // 66
			"from a import b,c",
			"from a import b,c\n",
			"from a import b, c\n",
		},
		{ // 67
			"import a",
			"import a\n",
			"import a\n",
		},
		{ // 68
			"import a , b",
			"import a,b\n",
			"import a, b\n",
		},
		{ // 69
			"class a(**b):c",
			"class a(**b):c\n",
			"class a(**b): c\n",
		},
		{ // 70
			"class a ( **b, c = d) : e",
			"class a(**b,c=d):e\n",
			"class a(**b, c = d): e\n",
		},
		{ // 71
			"@a\nclass b:c",
			"@a\nclass b:c\n",
			"@a\nclass b: c\n",
		},
		{ // 72
			"@a\n@b\nclass c:d",
			"@a\n@b\nclass c:d\n",
			"@a\n@b\nclass c: d\n",
		},
		{ // 73
			"a == b",
			"a==b\n",
			"a == b\n",
		},
		{ // 74
			"a == b == c",
			"a==b==c\n",
			"a == b == c\n",
		},
		{ // 75
			"a is b",
			"a is b\n",
			"a is b\n",
		},
		{ // 76
			"a is not b",
			"a is not b\n",
			"a is not b\n",
		},
		{ // 77
			"a in b",
			"a in b\n",
			"a in b\n",
		},
		{ // 78
			"a not in b",
			"a not in b\n",
			"a not in b\n",
		},
		{ // 79
			"assert lambda:a",
			"assert lambda:a\n",
			"assert lambda: a\n",
		},
		{ // 80
			"assert lambda a : b",
			"assert lambda a:b\n",
			"assert lambda a: b\n",
		},
		{ // 81
			"import a as b",
			"import a as b\n",
			"import a as b\n",
		},
		{ // 82
			"import a.b",
			"import a.b\n",
			"import a.b\n",
		},
		{ // 83
			"import a.b.c",
			"import a.b.c\n",
			"import a.b.c\n",
		},
		{ // 84
			"a*b",
			"a*b\n",
			"a * b\n",
		},
		{ // 85
			"a / b",
			"a/b\n",
			"a / b\n",
		},
		{ // 86
			"a * b / c*d",
			"a*b/c*d\n",
			"a * b / c * d\n",
		},
		{ // 87
			"nonlocal a",
			"nonlocal a\n",
			"nonlocal a\n",
		},
		{ // 88
			"nonlocal a,b , c",
			"nonlocal a,b,c\n",
			"nonlocal a, b, c\n",
		},
		{ // 89
			"not a",
			"not a\n",
			"not a\n",
		},
		{ // 90
			"not not not not a",
			"not not not not a\n",
			"not not not not a\n",
		},
		{ // 91
			"a|b",
			"a|b\n",
			"a | b\n",
		},
		{ // 92
			"a|b | c",
			"a|b|c\n",
			"a | b | c\n",
		},
		{ // 93
			"a or b",
			"a or b\n",
			"a or b\n",
		},
		{ // 94
			"a or b or c",
			"a or b or c\n",
			"a or b or c\n",
		},
		{ // 95
			"def a():b",
			"def a():b\n",
			"def a(): b\n",
		},
		{ // 96
			"def a(b):c",
			"def a(b):c\n",
			"def a(b): c\n",
		},
		{ // 97
			"def a(b,/,c):d",
			"def a(b,/,c):d\n",
			"def a(b, /, c): d\n",
		},
		{ // 98
			"def a(b,c,/,d,e):f",
			"def a(b,c,/,d,e):f\n",
			"def a(b, c, /, d, e): f\n",
		},
		{ // 99
			"def a(b, *c):d",
			"def a(b,*c):d\n",
			"def a(b, *c): d\n",
		},
		{ // 100
			"def a(b, *c, d):e",
			"def a(b,*c,d):e\n",
			"def a(b, *c, d): e\n",
		},
		{ // 101
			"def a(b,**c):d",
			"def a(b,**c):d\n",
			"def a(b, **c): d\n",
		},
		{ // 102
			"def a(b , / , *c):d",
			"def a(b,/,*c):d\n",
			"def a(b, /, *c): d\n",
		},
		{ // 103
			"def a(b , / , *c,d):e",
			"def a(b,/,*c,d):e\n",
			"def a(b, /, *c, d): e\n",
		},
		{ // 104
			"def a(b , / , **c):d",
			"def a(b,/,**c):d\n",
			"def a(b, /, **c): d\n",
		},
		{ // 105
			"def a(b , / , c, d, *e, f, g, **h):i",
			"def a(b,/,c,d,*e,f,g,**h):i\n",
			"def a(b, /, c, d, *e, f, g, **h): i\n",
		},
		{ // 106
			"def a(b , c, d, *e, f, g, **h):i",
			"def a(b,c,d,*e,f,g,**h):i\n",
			"def a(b, c, d, *e, f, g, **h): i\n",
		},
		{ // 106
			"def a(b = c): d",
			"def a(b=c):d\n",
			"def a(b = c): d\n",
		},
		{ // 107
			"def a(b:c = d): e",
			"def a(b:c=d):e\n",
			"def a(b: c = d): e\n",
		},
		{ // 107
			"async def a(b = c): d",
			"async def a(b=c):d\n",
			"async def a(b = c): d\n",
		},
		{ // 107
			"a( b )",
			"a(b)\n",
			"a(b)\n",
		},
		{ // 108
			"a( *b )",
			"a(*b)\n",
			"a(*b)\n",
		},
		{ // 109
			"a ** b",
			"a**b\n",
			"a ** b\n",
		},
		{ // 110
			"await a**b",
			"await a**b\n",
			"await a ** b\n",
		},
		{ // 111
			"a",
			"a\n",
			"a\n",
		},
		{ // 112
			"a.b",
			"a.b\n",
			"a.b\n",
		},
		{ // 113
			"a . b",
			"a.b\n",
			"a.b\n",
		},
		{ // 114
			"a[b]",
			"a[b]\n",
			"a[b]\n",
		},
		{ // 115
			"a[ b ]",
			"a[b]\n",
			"a[b]\n",
		},
		{ // 116
			"a(b)",
			"a(b)\n",
			"a(b)\n",
		},
		{ // 117
			"a( b )",
			"a(b)\n",
			"a(b)\n",
		},
		{ // 118
			"raise",
			"raise\n",
			"raise\n",
		},
		{ // 119
			"raise a",
			"raise a\n",
			"raise a\n",
		},
		{ // 120
			"raise a from b",
			"raise a from b\n",
			"raise a from b\n",
		},
		{ // 121
			"from . a import b",
			"from .a import b\n",
			"from .a import b\n",
		},
		{ // 122
			"from ..a import b",
			"from ..a import b\n",
			"from ..a import b\n",
		},
		{ // 123
			"from ... import a",
			"from ... import a\n",
			"from ... import a\n",
		},
		{ // 124
			"from ....a.b import c",
			"from ....a.b import c\n",
			"from ....a.b import c\n",
		},
		{ // 125
			"def a():\n\treturn",
			"def a():\n\treturn\n",
			"def a():\n\treturn\n",
		},
		{ // 126
			"def a():\n\treturn b",
			"def a():\n\treturn b\n",
			"def a():\n\treturn b\n",
		},
		{ // 127
			"a>>b",
			"a>>b\n",
			"a >> b\n",
		},
		{ // 128
			"a << b",
			"a<<b\n",
			"a << b\n",
		},
		{ // 129
			"assert a",
			"assert a\n",
			"assert a\n",
		},
		{ // 130
			"del b",
			"del b\n",
			"del b\n",
		},
		{ // 131
			"return a",
			"return a\n",
			"return a\n",
		},
		{ // 132
			"yield a",
			"yield a\n",
			"yield a\n",
		},
		{ // 133
			"raise a",
			"raise a\n",
			"raise a\n",
		},
		{ // 134
			"import a",
			"import a\n",
			"import a\n",
		},
		{ // 135
			"global a",
			"global a\n",
			"global a\n",
		},
		{ // 136
			"nonlocal a",
			"nonlocal a\n",
			"nonlocal a\n",
		},
		{ // 137
			"type a = b",
			"type a=b\n",
			"type a = b\n",
		},
		{ // 138
			"a = b",
			"a=b\n",
			"a = b\n",
		},
		{ // 139
			"a: b = c",
			"a:b=c\n",
			"a: b = c\n",
		},
		{ // 140
			"a += b",
			"a+=b\n",
			"a += b\n",
		},
		{ // 141
			"pass",
			"pass\n",
			"pass\n",
		},
		{ // 142
			"break",
			"break\n",
			"break\n",
		},
		{ // 143
			"continue",
			"continue\n",
			"continue\n",
		},
		{ // 144
			"a[b]",
			"a[b]\n",
			"a[b]\n",
		},
		{ // 145
			"a [ b : c ] ",
			"a[b:c]\n",
			"a[b : c]\n",
		},
		{ // 146
			"a[ b : c : d]",
			"a[b:c:d]\n",
			"a[b : c : d]\n",
		},
		{ // 147
			"a[ b,c ]",
			"a[b,c]\n",
			"a[b, c]\n",
		},
		{ // 148
			"a[ b,c ,d]",
			"a[b,c,d]\n",
			"a[b, c, d]\n",
		},
		{ // 149
			"a = b",
			"a=b\n",
			"a = b\n",
		},
		{ // 150
			"a = *b",
			"a=*b\n",
			"a = *b\n",
		},
		{ // 151
			"a = *b, c",
			"a=*b,c\n",
			"a = *b, c\n",
		},
		{ // 152
			"a = b ,",
			"a=b,\n",
			"a = b,\n",
		},
		{ // 153
			"a = *b,",
			"a=*b,\n",
			"a = *b,\n",
		},
		{ // 154
			"a = *b, c",
			"a=*b,c\n",
			"a = *b, c\n",
		},
		{ // 155
			"a(*b)",
			"a(*b)\n",
			"a(*b)\n",
		},
		{ // 156
			"a(*b, c)",
			"a(*b,c)\n",
			"a(*b, c)\n",
		},
		{ // 157
			"a(*b, *c)",
			"a(*b,*c)\n",
			"a(*b, *c)\n",
		},
		{ // 158
			"a(*b, c = d)",
			"a(*b,c=d)\n",
			"a(*b, c = d)\n",
		},
		{ // 159
			"a",
			"a\n",
			"a\n",
		},
		{ // 160
			"if a: b",
			"if a:b\n",
			"if a: b\n",
		},
		{ // 161
			"a;b",
			"a;b\n",
			"a; b\n",
		},
		{ // 162
			"if a: \n\tb",
			"if a:\n\tb\n",
			"if a:\n\tb\n",
		},
		{ // 163
			"if a: \n\tb\n\tc",
			"if a:\n\tb\n\tc\n",
			"if a:\n\tb\n\tc\n",
		},
		{ // 164
			"if a:\n\t(\nb\n)",
			"if a:\n\t(b)\n",
			"if a:\n\t(b)\n",
		},
		{ // 165
			"if a:\n\tif b:\n\t\tc\n\t\td",
			"if a:\n\tif b:\n\t\tc\n\t\td\n",
			"if a:\n\tif b:\n\t\tc\n\t\td\n",
		},
		{ // 166
			"a = b",
			"a=b\n",
			"a = b\n",
		},
		{ // 167
			"a.b = c",
			"a.b=c\n",
			"a.b = c\n",
		},
		{ // 168
			"(a) = b",
			"(a)=b\n",
			"(a) = b\n",
		},
		{ // 169
			"[a] = b",
			"[a]=b\n",
			"[a] = b\n",
		},
		{ // 170
			"*a = b",
			"*a=b\n",
			"*a = b\n",
		},
		{ // 171
			"a, b = c",
			"a,b=c\n",
			"a, b = c\n",
		},
		{ // 172
			"try:a\nexcept b:c",
			"try:a\nexcept b:c\n",
			"try: a\nexcept b: c\n",
		},
		{ // 173
			"try:a\nexcept b:c\nexcept d:e",
			"try:a\nexcept b:c\nexcept d:e\n",
			"try: a\nexcept b: c\nexcept d: e\n",
		},
		{ // 174
			"try:a\nexcept *b:c",
			"try:a\nexcept *b:c\n",
			"try: a\nexcept *b: c\n",
		},
		{ // 175
			"try:a\nexcept *b:c\nexcept *d:e",
			"try:a\nexcept *b:c\nexcept *d:e\n",
			"try: a\nexcept *b: c\nexcept *d: e\n",
		},
		{ // 176
			"try:a\nexcept b:c\nelse: d",
			"try:a\nexcept b:c\nelse:d\n",
			"try: a\nexcept b: c\nelse: d\n",
		},
		{ // 177
			"try:a\nexcept b:c\nfinally: d",
			"try:a\nexcept b:c\nfinally:d\n",
			"try: a\nexcept b: c\nfinally: d\n",
		},
		{ // 178
			"try:a\nexcept b:c\nelse: d\nfinally:e",
			"try:a\nexcept b:c\nelse:d\nfinally:e\n",
			"try: a\nexcept b: c\nelse: d\nfinally: e\n",
		},
		{ // 179
			"def a[b](): c",
			"def a[b]():c\n",
			"def a[b](): c\n",
		},
		{ // 180
			"def a[b:c](): d",
			"def a[b:c]():d\n",
			"def a[b: c](): d\n",
		},
		{ // 181
			"def a[*b](): c",
			"def a[*b]():c\n",
			"def a[*b](): c\n",
		},
		{ // 182
			"def a[**b](): c",
			"def a[**b]():c\n",
			"def a[**b](): c\n",
		},
		{ // 183
			"class a[b,c, d ](): e",
			"class a[b,c,d]():e\n",
			"class a[b, c, d](): e\n",
		},
		{ // 184
			"type a = b",
			"type a=b\n",
			"type a = b\n",
		},
		{ // 185
			"type a[b] = c",
			"type a[b]=c\n",
			"type a[b] = c\n",
		},
		{ // 186
			"+a",
			"+a\n",
			"+a\n",
		},
		{ // 187
			"-a",
			"-a\n",
			"-a\n",
		},
		{ // 188
			"~a",
			"~a\n",
			"~a\n",
		},
		{ // 189
			"while a:b",
			"while a:b\n",
			"while a: b\n",
		},
		{ // 190
			"while a:b\nelse: c",
			"while a:b\nelse:c\n",
			"while a: b\nelse: c\n",
		},
		{ // 191
			"with a: b",
			"with a:b\n",
			"with a: b\n",
		},
		{ // 192
			"with a as b:c",
			"with a as b:c\n",
			"with a as b: c\n",
		},
		{ // 193
			"with a,b: c",
			"with a,b:c\n",
			"with a, b: c\n",
		},
		{ // 194
			"with a as b, c,d as e:f",
			"with a as b,c,d as e:f\n",
			"with a as b, c, d as e: f\n",
		},
		{ // 195
			"a^b",
			"a^b\n",
			"a ^ b\n",
		},
		{ // 196
			"yield a",
			"yield a\n",
			"yield a\n",
		},
		{ // 197
			"yield from a",
			"yield from a\n",
			"yield from a\n",
		},
		{ // 198
			"(a for b in c if d)",
			"(a for b in c if d)\n",
			"(a for b in c if d)\n",
		},
		{ // 199
			"(a async for b in c if d)",
			"(a async for b in c if d)\n",
			"(a async for b in c if d)\n",
		},
		{ // 200
			"(a for b in c if d for e in f)",
			"(a for b in c if d for e in f)\n",
			"(a for b in c if d for e in f)\n",
		},
		{ // 201
			"a if b else c",
			"a if b else c\n",
			"a if b else c\n",
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

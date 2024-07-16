package python

import (
	"testing"

	"vimagination.zapto.org/parser"
)

func TestTokeniser(t *testing.T) {
	for n, test := range [...]struct {
		Input  string
		Output []parser.Token
	}{
		{ // 1
			"",
			[]parser.Token{
				{Type: parser.TokenDone, Data: ""},
			},
		},
		{ // 2
			" ",
			[]parser.Token{
				{Type: TokenWhitespace, Data: " "},
				{Type: parser.TokenDone, Data: ""},
			},
		},
		{ // 3
			" \t",
			[]parser.Token{
				{Type: TokenWhitespace, Data: " \t"},
				{Type: parser.TokenDone, Data: ""},
			},
		},
		{ // 4
			" \n ",
			[]parser.Token{
				{Type: TokenWhitespace, Data: " "},
				{Type: TokenLineTerminator, Data: "\n"},
				{Type: TokenWhitespace, Data: " "},
				{Type: parser.TokenDone, Data: ""},
			},
		},
		{ // 5
			"\"\"",
			[]parser.Token{
				{Type: TokenStringLiteral, Data: "\"\""},
				{Type: parser.TokenDone, Data: ""},
			},
		},
		{ // 6
			"\"\\\"\"",
			[]parser.Token{
				{Type: TokenStringLiteral, Data: "\"\\\"\""},
				{Type: parser.TokenDone, Data: ""},
			},
		},
		{ // 7
			"\"\n\"",
			[]parser.Token{
				{Type: parser.TokenError, Data: "invalid character"},
			},
		},
		{ // 8
			"\"\\n\"",
			[]parser.Token{
				{Type: TokenStringLiteral, Data: "\"\\n\""},
				{Type: parser.TokenDone, Data: ""},
			},
		},
		{ // 9
			"\"\\0\"",
			[]parser.Token{
				{Type: TokenStringLiteral, Data: "\"\\0\""},
				{Type: parser.TokenDone, Data: ""},
			},
		},
		{ // 10
			"\"\\x20\"",
			[]parser.Token{
				{Type: TokenStringLiteral, Data: "\"\\x20\""},
				{Type: parser.TokenDone, Data: ""},
			},
		},
		{ // 11
			"\"\\u2020\"",
			[]parser.Token{
				{Type: TokenStringLiteral, Data: "\"\\u2020\""},
				{Type: parser.TokenDone, Data: ""},
			},
		},
		{ // 12
			"b\"abc123\\\"\\'\"",
			[]parser.Token{
				{Type: TokenStringLiteral, Data: "b\"abc123\\\"\\'\""},
				{Type: parser.TokenDone, Data: ""},
			},
		},
		{ // 13
			"B\"abc123\\\"\\'\"",
			[]parser.Token{
				{Type: TokenStringLiteral, Data: "B\"abc123\\\"\\'\""},
				{Type: parser.TokenDone, Data: ""},
			},
		},
		{ // 14
			"f\"abc123\\\"\\'\"",
			[]parser.Token{
				{Type: TokenStringLiteral, Data: "f\"abc123\\\"\\'\""},
				{Type: parser.TokenDone, Data: ""},
			},
		},
		{ // 15
			"F\"abc123\\\"\\'\"",
			[]parser.Token{
				{Type: TokenStringLiteral, Data: "F\"abc123\\\"\\'\""},
				{Type: parser.TokenDone, Data: ""},
			},
		},
		{ // 16
			"r\"abc123\\\"\\'\"",
			[]parser.Token{
				{Type: TokenStringLiteral, Data: "r\"abc123\\\""},
				{Type: parser.TokenError, Data: "invalid character"},
			},
		},
		{ // 17
			"R\"abc123\\\"\\'\"",
			[]parser.Token{
				{Type: TokenStringLiteral, Data: "R\"abc123\\\""},
				{Type: parser.TokenError, Data: "invalid character"},
			},
		},
		{ // 18
			"fR\"abc123\\\"\\'\"",
			[]parser.Token{
				{Type: TokenStringLiteral, Data: "fR\"abc123\\\""},
				{Type: parser.TokenError, Data: "invalid character"},
			},
		},
		{ // 19
			"Fr\"abc123\\\"\\'\"",
			[]parser.Token{
				{Type: TokenStringLiteral, Data: "Fr\"abc123\\\""},
				{Type: parser.TokenError, Data: "invalid character"},
			},
		},
		{ // 20
			"RF\"abc123\\\"\\'\"",
			[]parser.Token{
				{Type: TokenStringLiteral, Data: "RF\"abc123\\\""},
				{Type: parser.TokenError, Data: "invalid character"},
			},
		},
		{ // 21
			"rf\"abc123\\\"\\'\"",
			[]parser.Token{
				{Type: TokenStringLiteral, Data: "rf\"abc123\\\""},
				{Type: parser.TokenError, Data: "invalid character"},
			},
		},
		{ // 22
			"Br\"abc123\\\"\\'\"",
			[]parser.Token{
				{Type: TokenStringLiteral, Data: "Br\"abc123\\\""},
				{Type: parser.TokenError, Data: "invalid character"},
			},
		},
		{ // 23
			"bR\"abc123\\\"\\'\"",
			[]parser.Token{
				{Type: TokenStringLiteral, Data: "bR\"abc123\\\""},
				{Type: parser.TokenError, Data: "invalid character"},
			},
		},
		{ // 24
			"rb\"abc123\\\"\\'\"",
			[]parser.Token{
				{Type: TokenStringLiteral, Data: "rb\"abc123\\\""},
				{Type: parser.TokenError, Data: "invalid character"},
			},
		},
		{ // 25
			"RB\"abc123\\\"\\'\"",
			[]parser.Token{
				{Type: TokenStringLiteral, Data: "RB\"abc123\\\""},
				{Type: parser.TokenError, Data: "invalid character"},
			},
		},
		{ // 26
			"\"\"\"abc123\"'456\"\"\"",
			[]parser.Token{
				{Type: TokenStringLiteral, Data: "\"\"\"abc123\"'456\"\"\""},
				{Type: parser.TokenDone, Data: ""},
			},
		},
		{ // 27
			"'''abc123\"'456'''",
			[]parser.Token{
				{Type: TokenStringLiteral, Data: "'''abc123\"'456'''"},
				{Type: parser.TokenDone, Data: ""},
			},
		},
		{ // 28
			"b\"\"\"abc123\"'456\"\"\"",
			[]parser.Token{
				{Type: TokenStringLiteral, Data: "b\"\"\"abc123\"'456\"\"\""},
				{Type: parser.TokenDone, Data: ""},
			},
		},
		{ // 29
			"b'''abc123\"'456'''",
			[]parser.Token{
				{Type: TokenStringLiteral, Data: "b'''abc123\"'456'''"},
				{Type: parser.TokenDone, Data: ""},
			},
		},
		{ // 30
			"f\"\"\"abc123\"'456\"\"\"",
			[]parser.Token{
				{Type: TokenStringLiteral, Data: "f\"\"\"abc123\"'456\"\"\""},
				{Type: parser.TokenDone, Data: ""},
			},
		},
		{ // 31
			"f'''abc123\"'456'''",
			[]parser.Token{
				{Type: TokenStringLiteral, Data: "f'''abc123\"'456'''"},
				{Type: parser.TokenDone, Data: ""},
			},
		},
		{ // 32
			"r\"\"\"abc123\"'456\"\"\"",
			[]parser.Token{
				{Type: TokenStringLiteral, Data: "r\"\"\"abc123\"'456\"\"\""},
				{Type: parser.TokenDone, Data: ""},
			},
		},
		{ // 33
			"r'''abc123\"'456'''",
			[]parser.Token{
				{Type: TokenStringLiteral, Data: "r'''abc123\"'456'''"},
				{Type: parser.TokenDone, Data: ""},
			},
		},
		{ // 34
			"u\"abc123\\\"\\'\"",
			[]parser.Token{
				{Type: TokenStringLiteral, Data: "u\"abc123\\\"\\'\""},
				{Type: parser.TokenDone, Data: ""},
			},
		},
		{ // 35
			"u'abc123\\\"\\''",
			[]parser.Token{
				{Type: TokenStringLiteral, Data: "u'abc123\\\"\\''"},
				{Type: parser.TokenDone, Data: ""},
			},
		},
		{ // 36
			"o\"abc123\\\"'456\"",
			[]parser.Token{
				{Type: TokenIdentifier, Data: "o"},
				{Type: TokenStringLiteral, Data: "\"abc123\\\"'456\""},
				{Type: parser.TokenDone, Data: ""},
			},
		},
		{ // 37
			"o'abc123\"\\'456'",
			[]parser.Token{
				{Type: TokenIdentifier, Data: "o"},
				{Type: TokenStringLiteral, Data: "'abc123\"\\'456'"},
				{Type: parser.TokenDone, Data: ""},
			},
		},
		{ // 38
			"o\"\"\"abc123\"'456\"\"\"",
			[]parser.Token{
				{Type: TokenIdentifier, Data: "o"},
				{Type: TokenStringLiteral, Data: "\"\"\"abc123\"'456\"\"\""},
				{Type: parser.TokenDone, Data: ""},
			},
		},
		{ // 39
			"o'''abc123\"'456'''",
			[]parser.Token{
				{Type: TokenIdentifier, Data: "o"},
				{Type: TokenStringLiteral, Data: "'''abc123\"'456'''"},
				{Type: parser.TokenDone, Data: ""},
			},
		},
		{ // 40
			"\"\"\"abc123\n456\"\"\"",
			[]parser.Token{
				{Type: TokenStringLiteral, Data: "\"\"\"abc123\n456\"\"\""},
				{Type: parser.TokenDone, Data: ""},
			},
		},
		{ // 41
			"'''abc123\n456'''",
			[]parser.Token{
				{Type: TokenStringLiteral, Data: "'''abc123\n456'''"},
				{Type: parser.TokenDone, Data: ""},
			},
		},
		{ // 42
			"\"an unclosed string",
			[]parser.Token{
				{Type: parser.TokenError, Data: "unexpected EOF"},
			},
		},
		{ // 43
			"'an unclosed string",
			[]parser.Token{
				{Type: parser.TokenError, Data: "unexpected EOF"},
			},
		},
		{ // 44
			"\"\"\"an unclosed string",
			[]parser.Token{
				{Type: parser.TokenError, Data: "unexpected EOF"},
			},
		},
		{ // 45
			"'''an unclosed string",
			[]parser.Token{
				{Type: parser.TokenError, Data: "unexpected EOF"},
			},
		},
		{ // 46
			"False",
			[]parser.Token{
				{Type: TokenBooleanLiteral, Data: "False"},
				{Type: parser.TokenDone, Data: ""},
			},
		},
		{ // 47
			"True",
			[]parser.Token{
				{Type: TokenBooleanLiteral, Data: "True"},
				{Type: parser.TokenDone, Data: ""},
			},
		},
		{ // 48
			"false",
			[]parser.Token{
				{Type: TokenIdentifier, Data: "false"},
				{Type: parser.TokenDone, Data: ""},
			},
		},
		{ // 49
			"true",
			[]parser.Token{
				{Type: TokenIdentifier, Data: "true"},
				{Type: parser.TokenDone, Data: ""},
			},
		},
		{ // 50
			"None",
			[]parser.Token{
				{Type: TokenNullLiteral, Data: "None"},
				{Type: parser.TokenDone, Data: ""},
			},
		},
		{ // 51
			"none",
			[]parser.Token{
				{Type: TokenIdentifier, Data: "none"},
				{Type: parser.TokenDone, Data: ""},
			},
		},
		{ // 52
			"# A Comment\n\"A string\"\n\"another string\"# Another Comment\n\"\"",
			[]parser.Token{
				{Type: TokenComment, Data: "# A Comment"},
				{Type: TokenLineTerminator, Data: "\n"},
				{Type: TokenStringLiteral, Data: "\"A string\""},
				{Type: TokenLineTerminator, Data: "\n"},
				{Type: TokenStringLiteral, Data: "\"another string\""},
				{Type: TokenComment, Data: "# Another Comment"},
				{Type: TokenLineTerminator, Data: "\n"},
				{Type: TokenStringLiteral, Data: "\"\""},
				{Type: parser.TokenDone, Data: ""},
			},
		},
		{ // 53
			"identifier",
			[]parser.Token{
				{Type: TokenIdentifier, Data: "identifier"},
				{Type: parser.TokenDone, Data: ""},
			},
		},
		{ // 54
			"another identifier",
			[]parser.Token{
				{Type: TokenIdentifier, Data: "another"},
				{Type: TokenWhitespace, Data: " "},
				{Type: TokenIdentifier, Data: "identifier"},
				{Type: parser.TokenDone, Data: ""},
			},
		},
		{ // 55
			"f r u fR rB farm",
			[]parser.Token{
				{Type: TokenIdentifier, Data: "f"},
				{Type: TokenWhitespace, Data: " "},
				{Type: TokenIdentifier, Data: "r"},
				{Type: TokenWhitespace, Data: " "},
				{Type: TokenIdentifier, Data: "u"},
				{Type: TokenWhitespace, Data: " "},
				{Type: TokenIdentifier, Data: "fR"},
				{Type: TokenWhitespace, Data: " "},
				{Type: TokenIdentifier, Data: "rB"},
				{Type: TokenWhitespace, Data: " "},
				{Type: TokenIdentifier, Data: "farm"},
				{Type: parser.TokenDone, Data: ""},
			},
		},
		{ // 56
			"await if for else global not yield from",
			[]parser.Token{
				{Type: TokenKeyword, Data: "await"},
				{Type: TokenWhitespace, Data: " "},
				{Type: TokenKeyword, Data: "if"},
				{Type: TokenWhitespace, Data: " "},
				{Type: TokenKeyword, Data: "for"},
				{Type: TokenWhitespace, Data: " "},
				{Type: TokenKeyword, Data: "else"},
				{Type: TokenWhitespace, Data: " "},
				{Type: TokenKeyword, Data: "global"},
				{Type: TokenWhitespace, Data: " "},
				{Type: TokenKeyword, Data: "not"},
				{Type: TokenWhitespace, Data: " "},
				{Type: TokenKeyword, Data: "yield"},
				{Type: TokenWhitespace, Data: " "},
				{Type: TokenKeyword, Data: "from"},
				{Type: parser.TokenDone, Data: ""},
			},
		},
		{ // 57
			"+ % | = == : := - * ** < >= != ~",
			[]parser.Token{
				{Type: TokenOperator, Data: "+"},
				{Type: TokenWhitespace, Data: " "},
				{Type: TokenOperator, Data: "%"},
				{Type: TokenWhitespace, Data: " "},
				{Type: TokenOperator, Data: "|"},
				{Type: TokenWhitespace, Data: " "},
				{Type: TokenOperator, Data: "="},
				{Type: TokenWhitespace, Data: " "},
				{Type: TokenOperator, Data: "=="},
				{Type: TokenWhitespace, Data: " "},
				{Type: TokenOperator, Data: ":"},
				{Type: TokenWhitespace, Data: " "},
				{Type: TokenOperator, Data: ":="},
				{Type: TokenWhitespace, Data: " "},
				{Type: TokenOperator, Data: "-"},
				{Type: TokenWhitespace, Data: " "},
				{Type: TokenOperator, Data: "*"},
				{Type: TokenWhitespace, Data: " "},
				{Type: TokenOperator, Data: "**"},
				{Type: TokenWhitespace, Data: " "},
				{Type: TokenOperator, Data: "<"},
				{Type: TokenWhitespace, Data: " "},
				{Type: TokenOperator, Data: ">="},
				{Type: TokenWhitespace, Data: " "},
				{Type: TokenOperator, Data: "!="},
				{Type: TokenWhitespace, Data: " "},
				{Type: TokenOperator, Data: "~"},
				{Type: parser.TokenDone, Data: ""},
			},
		},
		{ // 58
			"+= %= |= -= -> *= **= /= //= <<= >>= , . ;",
			[]parser.Token{
				{Type: TokenDelimiter, Data: "+="},
				{Type: TokenWhitespace, Data: " "},
				{Type: TokenDelimiter, Data: "%="},
				{Type: TokenWhitespace, Data: " "},
				{Type: TokenDelimiter, Data: "|="},
				{Type: TokenWhitespace, Data: " "},
				{Type: TokenDelimiter, Data: "-="},
				{Type: TokenWhitespace, Data: " "},
				{Type: TokenDelimiter, Data: "->"},
				{Type: TokenWhitespace, Data: " "},
				{Type: TokenDelimiter, Data: "*="},
				{Type: TokenWhitespace, Data: " "},
				{Type: TokenDelimiter, Data: "**="},
				{Type: TokenWhitespace, Data: " "},
				{Type: TokenDelimiter, Data: "/="},
				{Type: TokenWhitespace, Data: " "},
				{Type: TokenDelimiter, Data: "//="},
				{Type: TokenWhitespace, Data: " "},
				{Type: TokenDelimiter, Data: "<<="},
				{Type: TokenWhitespace, Data: " "},
				{Type: TokenDelimiter, Data: ">>="},
				{Type: TokenWhitespace, Data: " "},
				{Type: TokenDelimiter, Data: ","},
				{Type: TokenWhitespace, Data: " "},
				{Type: TokenDelimiter, Data: "."},
				{Type: TokenWhitespace, Data: " "},
				{Type: TokenDelimiter, Data: ";"},
				{Type: parser.TokenDone, Data: ""},
			},
		},
		{ // 59
			"!",
			[]parser.Token{
				{Type: parser.TokenError, Data: "invalid character"},
			},
		},
		{ // 60
			"( )",
			[]parser.Token{
				{Type: TokenDelimiter, Data: "("},
				{Type: TokenWhitespace, Data: " "},
				{Type: TokenDelimiter, Data: ")"},
				{Type: parser.TokenDone, Data: ""},
			},
		},
		{ // 61
			"( ",
			[]parser.Token{
				{Type: TokenDelimiter, Data: "("},
				{Type: TokenWhitespace, Data: " "},
				{Type: parser.TokenError, Data: "unexpected EOF"},
			},
		},
		{ // 62
			"[ ]",
			[]parser.Token{
				{Type: TokenDelimiter, Data: "["},
				{Type: TokenWhitespace, Data: " "},
				{Type: TokenDelimiter, Data: "]"},
				{Type: parser.TokenDone, Data: ""},
			},
		},
		{ // 63
			"[ )",
			[]parser.Token{
				{Type: TokenDelimiter, Data: "["},
				{Type: TokenWhitespace, Data: " "},
				{Type: parser.TokenError, Data: "invalid character"},
			},
		},
		{ // 64
			"{ ( [ ] ) }",
			[]parser.Token{
				{Type: TokenDelimiter, Data: "{"},
				{Type: TokenWhitespace, Data: " "},
				{Type: TokenDelimiter, Data: "("},
				{Type: TokenWhitespace, Data: " "},
				{Type: TokenDelimiter, Data: "["},
				{Type: TokenWhitespace, Data: " "},
				{Type: TokenDelimiter, Data: "]"},
				{Type: TokenWhitespace, Data: " "},
				{Type: TokenDelimiter, Data: ")"},
				{Type: TokenWhitespace, Data: " "},
				{Type: TokenDelimiter, Data: "}"},
				{Type: parser.TokenDone, Data: ""},
			},
		},
		{ // 65
			"{ ( [ ] ) )",
			[]parser.Token{
				{Type: TokenDelimiter, Data: "{"},
				{Type: TokenWhitespace, Data: " "},
				{Type: TokenDelimiter, Data: "("},
				{Type: TokenWhitespace, Data: " "},
				{Type: TokenDelimiter, Data: "["},
				{Type: TokenWhitespace, Data: " "},
				{Type: TokenDelimiter, Data: "]"},
				{Type: TokenWhitespace, Data: " "},
				{Type: TokenDelimiter, Data: ")"},
				{Type: TokenWhitespace, Data: " "},
				{Type: parser.TokenError, Data: "invalid character"},
			},
		},
	} {
		p := parser.NewStringTokeniser(test.Input)

		SetTokeniser(&p)

		for m, tkn := range test.Output {
			tk, _ := p.GetToken()
			if tk.Type != tkn.Type {
				if tk.Type == parser.TokenError {
					t.Errorf("test %d.%d: unexpected error: %s", n+1, m+1, tk.Data)
				} else {
					t.Errorf("test %d.%d: Incorrect type, expecting %d, got %d", n+1, m+1, tkn.Type, tk.Type)
				}

				break
			} else if tk.Data != tkn.Data {
				t.Errorf("test %d.%d: Incorrect data, expecting %q, got %q", n+1, m+1, tkn.Data, tk.Data)

				break
			}
		}
	}
}

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
				{parser.TokenDone, ""},
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

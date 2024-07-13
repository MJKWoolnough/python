package python

import "vimagination.zapto.org/parser"

const (
	whitespace       = " \t"
	lineTerminator   = "\n"
	singleEscapeChar = "'\"\\bfnrtv"
	binaryDigit      = "01"
	octalDigit       = "01234567"
	decimalDigit     = "0123456789"
	hexDigit         = "0123456789abcdefABCDEF"
)

var keywords = [...]string{"await", "else", "import", "pass", "None", "break", "except", "in", "raise", "class", "finally", "is", "return", "and", "continue", "for", "lambda", "try", "as", "def", "from", "nonlocal", "while", "assert", "del", "global", "not", "with", "async", "elif", "if", "or", "yield"}

const (
	TokenWhitespace parser.TokenType = iota
	TokenLineTerminator
	TokenSingleLineComment
	TokenWord
	TokenKeyword
	TokenOperator
	TokenDelimiter
	TokenBooleanLiteral
	TokenNumericLiteral
	TokenStringLiteral
	TokenNullLiteral
)

package python

import "errors"

// Errors
var (
	ErrInvalidCharacter      = errors.New("invalid character")
	ErrInvalidNumber         = errors.New("invalid number")
	ErrInvalidIndent         = errors.New("invalid indent")
	ErrMissingImport         = errors.New("missing import keyword")
	ErrMissingModule         = errors.New("missing module")
	ErrMissingEquals         = errors.New("missing equals")
	ErrMissingElse           = errors.New("missing else")
	ErrMissingOp             = errors.New("missing operator")
	ErrInvalidCompound       = errors.New("invalid compound statement")
	ErrMissingNewline        = errors.New("missing newline")
	ErrMissingColon          = errors.New("missing colon")
	ErrMissingIn             = errors.New("missing in")
	ErrMissingFinally        = errors.New("missing finally")
	ErrMissingIdentifier     = errors.New("missing identifier")
	ErrMismatchedGroups      = errors.New("mismatched groups in except")
	ErrMissingOpeningParen   = errors.New("missing opening paren")
	ErrMissingClosingParen   = errors.New("missing closing paren")
	ErrMissingClosingBracket = errors.New("missing closing bracket")
	ErrMissingIndent         = errors.New("missing indent")
	ErrMissingComma          = errors.New("missing comma")
	ErrMissingClosingBrace   = errors.New("missing closing brace")
	ErrMissingFor            = errors.New("missing for keyword")
	ErrMissingIf             = errors.New("missing for keyword")
)

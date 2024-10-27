package python

import "errors"

// Errors
var (
	ErrInvalidCharacter      = errors.New("invalid character")
	ErrInvalidCompound       = errors.New("invalid compound statement")
	ErrInvalidEnclosure      = errors.New("invalid enclosure")
	ErrInvalidIndent         = errors.New("invalid indent")
	ErrInvalidKeyword        = errors.New("unexpected keyword")
	ErrInvalidNumber         = errors.New("invalid number")
	ErrMismatchedGroups      = errors.New("mismatched groups in except")
	ErrMissingClosingBrace   = errors.New("missing closing brace")
	ErrMissingClosingBracket = errors.New("missing closing bracket")
	ErrMissingClosingParen   = errors.New("missing closing paren")
	ErrMissingColon          = errors.New("missing colon")
	ErrMissingComma          = errors.New("missing comma")
	ErrMissingElse           = errors.New("missing else")
	ErrMissingEquals         = errors.New("missing equals")
	ErrMissingFinally        = errors.New("missing finally")
	ErrMissingFor            = errors.New("missing for keyword")
	ErrMissingIdentifier     = errors.New("missing identifier")
	ErrMissingIf             = errors.New("missing if keyword")
	ErrMissingImport         = errors.New("missing import keyword")
	ErrMissingIn             = errors.New("missing in")
	ErrMissingIndent         = errors.New("missing indent")
	ErrMissingModule         = errors.New("missing module")
	ErrMissingNewline        = errors.New("missing newline")
	ErrMissingOp             = errors.New("missing operator")
	ErrMissingOpeningParen   = errors.New("missing opening paren")
)

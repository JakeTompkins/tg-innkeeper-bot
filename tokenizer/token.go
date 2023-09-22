package tokenizer

type TokenType string

type Token struct {
	Type  TokenType
	Value string
}

const (
	COMMAND        = "COMMAND"
	IMPORTANT_WORD = "IMPORTANT"
	IGNORED_WORD   = "IGNORED"
	NUMBER         = "NUMBER"

	EOF = ""
)

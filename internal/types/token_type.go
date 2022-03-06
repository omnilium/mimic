package types

type TokenType uint8

const (
	TOKEN_TEXT TokenType = iota
	TOKEN_VAR
	TOKEN_BLOCK
	TOKEN_COMMENT
)

var TokenTypeNames = map[TokenType]string{
	TOKEN_TEXT:    "text",
	TOKEN_VAR:     "var",
	TOKEN_BLOCK:   "block",
	TOKEN_COMMENT: "comment",
}

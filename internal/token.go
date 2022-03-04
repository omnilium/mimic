package internal

import (
	"fmt"
	"strings"
)

// Token represents a string from the template.
// 		Token.TokenType		- .TEXT, .VAR, .BLOCK or .COMMENT
//		Token.Content		- The token source string
//		Token.PositionRow	- The row of the template where the token was found.
//		Token.PositionCol	- The column of the template where the token was found.
// 		Token.LineNo		- The line number of the template where the token was found.
type Token struct {
	TokenType   int
	Content     string
	PositionRow int
	PositionCol int
	LineNo      int
}

func (t *Token) String() string {
	return fmt.Sprintf("<Token type=\"%d\" content=\"%s\">", t.TokenType, t.Content)
}

func (t *Token) SplitContents() []string {
	return strings.Split(t.Content, " ")
}

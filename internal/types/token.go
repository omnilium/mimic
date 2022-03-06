package types

import (
	"fmt"
	"github.com/omnilium/mimic/internal/utility"
	"strings"
)

// Token represents a string from the template.
// 		Token.TokenType		- .TEXT, .VAR, .BLOCK or .COMMENT
//		Token.Content		- The token source string
//		Token.PositionRow	- The row of the template where the token was found.
//		Token.PositionCol	- The column of the template where the token was found.
// 		Token.LineNo		- The line number of the template where the token was found.
type Token struct {
	TokenType TokenType
	Content   string
	LineNo    int
}

func (t *Token) Representation(debug bool) string {
	switch debug {
	case true:
		return fmt.Sprintf("<Token type=\"%s\" lineNo=\"%d\" content=\"%s\" />",
			TokenTypeNames[t.TokenType],
			t.LineNo,
			t.Content)
	default:
		maxLen := 20
		if len(t.Content) < 20 {
			maxLen = len(t.Content)
		}
		return fmt.Sprintf("<Token type=\"%s\" lineNo=\"%d\" content=\"%s...\" />",
			TokenTypeNames[t.TokenType],
			t.LineNo,
			strings.Replace(t.Content[0:maxLen], "\n", "", -1))
	}
}

func (t *Token) SplitContents() []string {
	return utility.SmartSplit(t.Content)
}

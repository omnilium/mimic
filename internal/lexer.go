package internal

import (
	"fmt"
	"regexp"
	"strings"
)

const FILTER_SEPARATOR = "|"
const FILTER_ARGUMENT_SEPARATOR = ":"
const VARIABLE_ATTRIBUTE_SEPARATOR = "."
const BLOCK_TAG_START = "{%"
const BLOCK_TAG_END = "%}"
const VARIABLE_TAG_START = "{{"
const VARIABLE_TAG_END = "}}"
const COMMENT_TAG_START = "{#"
const COMMENT_TAG_END = "#}"
const SINGLE_BRACE_START = "{"
const SINGLE_BRACE_END = "}"

const (
	TOKEN_TEXT int = iota
	TOKEN_VAR
	TOKEN_BLOCK
	TOKEN_COMMENT
)

type Lexer struct {
	template string
	verbatim bool
	r        *regexp.Regexp
}

func (l *Lexer) String() string {
	return fmt.Sprintf("<Lexer template_string=\"%s...\" verbatim=\"%t\">",
		strings.Replace(l.template[:20], "\n", "", -1),
		l.verbatim)
}

// NewLexer initializes a new Lexer with the given template string and a flag for verbatim usage.
func NewLexer(template string, verbatim bool) (*Lexer, error) {
	r, err := regexp.Compile(`({%.*?%}|{{.*?}}|{#.*?#})`)
	if err != nil {
		return nil, err
	}

	return &Lexer{template, verbatim, r}, nil
}

func (l *Lexer) createToken(tokenString string, position int, lineNo int, inTag bool) *Token {
	//if inTag {
	//	tokenStart := tokenString[0:2]
	//	if tokenStart == BLOCK_TAG_START {
	//		content := strings.TrimSpace(tokenString[2 : len(tokenString)-2])
	//
	//		if l.verbatim {
	//			if content != "verbatim" {
	//				return
	//			}
	//		}
	//	}
	//}
	return nil
}

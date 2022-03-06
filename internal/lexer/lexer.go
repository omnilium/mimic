package lexer

import (
	"fmt"
	"github.com/omnilium/mimic/internal/types"
	"github.com/omnilium/mimic/internal/utility"
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

type Lexer struct {
	template      string
	verbatim      bool
	r             *regexp.Regexp
	verbatimBlock string
}

// NewLexer initializes a new Lexer with the given template string and a flag for verbatim usage.
func NewLexer(template string) (*Lexer, error) {
	r, err := regexp.Compile(`({%.*?%}|{{.*?}}|{#.*?#})`)
	if err != nil {
		return nil, err
	}

	return &Lexer{
		template:      template,
		verbatim:      false,
		r:             r,
		verbatimBlock: "",
	}, nil
}

func (l *Lexer) Tokenize() []types.Token {
	inTag := false
	lineNo := 1
	var result []types.Token

	for _, tokenString := range utility.SmartSplitWithRegex(l.r, l.template, -1) {
		if tokenString != "" {
			result = append(result, l.createToken(tokenString, lineNo, inTag))
			lineNo += strings.Count(tokenString, "\n")
		}

		inTag = !inTag
	}

	return result
}

func (l *Lexer) createToken(tokenString string, lineNo int, inTag bool) types.Token {
	if inTag {
		tokenStart := tokenString[0:2]
		if tokenStart == BLOCK_TAG_START {
			content := strings.TrimSpace(tokenString[2 : len(tokenString)-2])

			if l.verbatim {
				if content != l.verbatimBlock {
					return types.Token{
						TokenType: types.TOKEN_TEXT,
						Content:   tokenString,
						LineNo:    lineNo,
					}
				}

				l.verbatim = false
			} else if content[:8] == "verbatim" {
				l.verbatim = true
				l.verbatimBlock = fmt.Sprintf("end%s", content)
			}

			return types.Token{
				TokenType: types.TOKEN_BLOCK,
				Content:   content,
				LineNo:    lineNo,
			}
		}

		if l.verbatim == false {
			content := strings.TrimSpace(tokenString[2 : len(tokenString)-2])

			if tokenStart == VARIABLE_TAG_START {
				return types.Token{
					TokenType: types.TOKEN_VAR,
					Content:   content,
					LineNo:    lineNo,
				}
			}

			if tokenStart == COMMENT_TAG_START {
				return types.Token{
					TokenType: types.TOKEN_COMMENT,
					Content:   content,
					LineNo:    lineNo,
				}
			}
		}
	}

	return types.Token{
		TokenType: types.TOKEN_TEXT,
		Content:   tokenString,
		LineNo:    lineNo,
	}
}

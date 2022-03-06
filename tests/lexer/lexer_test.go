package lexer

import (
	lexerImpl "github.com/omnilium/mimic/internal/lexer"
	"github.com/omnilium/mimic/internal/types"
	"testing"
)

func TestLexer(t *testing.T) {
	lexer, err := lexerImpl.NewLexer(TestTemplate)
	if err != nil {
		t.Errorf("Error creating lexer: %s", err)
	}

	tokens := lexer.Tokenize()
	if len(tokens) != 9 {
		t.Errorf("Expected 9 tokens, got %d", len(tokens))
	}

	if tokens[0].TokenType != types.TOKEN_TEXT {
		t.Errorf("Expected token type %s, got %s", types.TokenTypeNames[types.TOKEN_TEXT], types.TokenTypeNames[tokens[0].TokenType])
	}

	if tokens[1].TokenType != types.TOKEN_VAR {
		t.Errorf("Expected token type %s, got %s", types.TokenTypeNames[types.TOKEN_VAR], types.TokenTypeNames[tokens[1].TokenType])
	}

	if tokens[2].TokenType != types.TOKEN_TEXT {
		t.Errorf("Expected token type %s, got %s", types.TokenTypeNames[types.TOKEN_TEXT], types.TokenTypeNames[tokens[2].TokenType])
	}

	if tokens[3].TokenType != types.TOKEN_COMMENT {
		t.Errorf("Expected token type %s, got %s", types.TokenTypeNames[types.TOKEN_COMMENT], types.TokenTypeNames[tokens[3].TokenType])
	}

	if tokens[4].TokenType != types.TOKEN_TEXT {
		t.Errorf("Expected token type %s, got %s", types.TokenTypeNames[types.TOKEN_TEXT], types.TokenTypeNames[tokens[4].TokenType])
	}

	if tokens[5].TokenType != types.TOKEN_BLOCK {
		t.Errorf("Expected token type %s, got %s", types.TokenTypeNames[types.TOKEN_BLOCK], types.TokenTypeNames[tokens[5].TokenType])
	}

	if tokens[6].TokenType != types.TOKEN_TEXT {
		t.Errorf("Expected token type %s, got %s", types.TokenTypeNames[types.TOKEN_TEXT], types.TokenTypeNames[tokens[6].TokenType])
	}

	if tokens[7].TokenType != types.TOKEN_BLOCK {
		t.Errorf("Expected token type %s, got %s", types.TokenTypeNames[types.TOKEN_BLOCK], types.TokenTypeNames[tokens[7].TokenType])
	}

	if tokens[8].TokenType != types.TOKEN_TEXT {
		t.Errorf("Expected token type %s, got %s", types.TokenTypeNames[types.TOKEN_TEXT], types.TokenTypeNames[tokens[8].TokenType])
	}
}

const TestTemplate = "<!DOCTYPE html>\n<html lang=\"en\">\n<head>\n    <meta charset=\"UTF-8\">\n    <title>{{ page.title }}</title>\n</head>\n<body>\n\n{# This is a comment #}\n\n{% This is a block tag %}\n\n<div>\n    {% This is a nested block tag %}\n</div>\n\n</body>\n</html>"

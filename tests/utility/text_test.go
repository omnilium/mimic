package utility

import (
	"github.com/omnilium/mimic/internal/utility"
	"regexp"
	"testing"
)

const SmartSplitNoQuotedText = "<tag some_token></tag>"
const SmartSplitQuotedText = "<tag some_token=\"test text\"></tag>"

func TestSmartSplitNoQuotedText(t *testing.T) {
	result := utility.SmartSplit(SmartSplitNoQuotedText)
	if len(result) != 2 {
		t.Errorf("Expected 2 elements, got %d", len(result))
	}

	if result[0] != "<tag" {
		t.Errorf("Expected <tag, got %s", result[0])
	}

	if result[1] != "some_token></tag>" {
		t.Errorf("Expected some_token></tag>, got %s", result[1])
	}
}

func TestSmartSplitQuotedText(t *testing.T) {
	result := utility.SmartSplit(SmartSplitQuotedText)
	if len(result) != 2 {
		t.Errorf("Expected 2 elements, got %d", len(result))
	}

	if result[0] != "<tag" {
		t.Errorf("Expected <tag, got %s", result[0])
	}

	if result[1] != "some_token=\"test text\"></tag>" {
		t.Errorf("Expected some_token=\"test text\"></tag>, got %s", result[1])
	}
}

const SmartSplitWithRegexFilter = `\s+`
const SmartSplitWithRegexText = "this is a test"

func TestSmartSplitWithRegexNZero(t *testing.T) {
	r, err := regexp.Compile(SmartSplitWithRegexFilter)
	if err != nil {
		t.Errorf("Error compiling regex: %s", err)
	}

	result := utility.SmartSplitWithRegex(r, SmartSplitWithRegexText, 0)
	if result != nil {
		t.Errorf("Expected nil, got an array with len = %d", len(result))
	}
}

func TestSmartSplitWithRegexNNonZero(t *testing.T) {
	r, err := regexp.Compile(SmartSplitWithRegexFilter)
	if err != nil {
		t.Errorf("Error compiling regex: %s", err)
	}

	result := utility.SmartSplitWithRegex(r, SmartSplitWithRegexText, 2)
	if len(result) != 3 {
		t.Errorf("Expected 2 elements, got %d", len(result))
	}

	if result[0] != "this" {
		t.Errorf("Expected this, got %s", result[0])
	}

	if result[1] != " " {
		t.Errorf("Expected whitespace, got %s", result[1])
	}

	if result[2] != "is a test" {
		t.Errorf("Expected is a test, got %s", result[2])
	}
}

func TestSmartSplitWithRegexSLenZero(t *testing.T) {
	r, err := regexp.Compile(SmartSplitWithRegexFilter)
	if err != nil {
		t.Errorf("Error compiling regex: %s", err)
	}

	result := utility.SmartSplitWithRegex(r, "", -1)
	if len(result) != 1 {
		t.Errorf("Expected 1 element, got %d", len(result))
	}

	if result[0] != "" {
		t.Errorf("Expected empty string, got %s", result[0])
	}
}

func TestSmartSplitWithRegex(t *testing.T) {
	r, err := regexp.Compile(SmartSplitWithRegexFilter)
	if err != nil {
		t.Errorf("Error compiling regex: %s", err)
	}

	result := utility.SmartSplitWithRegex(r, SmartSplitWithRegexText, -1)
	if len(result) != 7 {
		t.Errorf("Expected 7 elements, got %d", len(result))
	}

	if result[0] != "this" {
		t.Errorf("Expected this, got %s", result[0])
	}

	if result[1] != " " {
		t.Errorf("Expected whitespace, got %s", result[1])
	}

	if result[2] != "is" {
		t.Errorf("Expected is, got %s", result[2])
	}

	if result[3] != " " {
		t.Errorf("Expected whitespace, got %s", result[3])
	}

	if result[4] != "a" {
		t.Errorf("Expected a, got %s", result[4])
	}

	if result[5] != " " {
		t.Errorf("Expected whitespace, got %s", result[5])
	}

	if result[6] != "test" {
		t.Errorf("Expected test, got %s", result[6])
	}
}

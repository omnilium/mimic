package utility

import "regexp"

// SMART_SPLIT Expression to match some_token and some_token="with spaces" (and similarly
// for single-quoted strings).
const SMART_SPLIT = `((?:[^\s'"]*(?:(?:"(?:[^"\\]|\\.)*"|'(?:[^'\\]|\\.)*')[^\s'"]*)+)|\S+)`

// SmartSplit splits a string by spaces, leaving quoted phrases together.
// Supports both single and double quotes, and supports escaping quotes with
// backslashes. In the output, strings will keep their initial and trailing
// quote marks and escaped quotes will remain escaped.
func SmartSplit(text string) []string {
	compiledRegex, _ := regexp.Compile(SMART_SPLIT)
	return compiledRegex.FindAllString(text, -1)
}

// SmartSplitWithRegex mimics Python's Regex Split function, by adding the
// matches to the resulting array. Go's Regex Split only returns the strings
// that don't match the filter, as per their documentation.
func SmartSplitWithRegex(re *regexp.Regexp, s string, n int) []string {
	if n == 0 {
		return nil
	}

	if len(s) == 0 {
		return []string{""}
	}

	matches := re.FindAllStringIndex(s, n)
	strings := make([]string, 0, len(matches))

	beg := 0
	end := 0
	for _, match := range matches {
		if n > 0 && len(strings) >= n-1 {
			break
		}

		end = match[0]
		if match[1] != 0 {
			strings = append(strings, s[beg:end])
			strings = append(strings, s[match[0]:match[1]])
		}
		beg = match[1]
	}

	if end != len(s) {
		strings = append(strings, s[beg:])
	}

	return strings
}

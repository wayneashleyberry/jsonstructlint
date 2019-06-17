package stringutil

import (
	"strings"
	"unicode"
)

const (
	ignoreString = "jsonstructlint"
)

// IsCamelCase implementation
func IsCamelCase(val string) bool {
	if strings.Contains(val, "_") {
		return false
	}

	if strings.ToLower(string(val[0])) != string(val[0]) {
		return false
	}

	return true
}

// IsTrimmed implementation
func IsTrimmed(in string) bool {
	trimmed := strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}
		return r
	}, in)

	return in == trimmed
}

// ContainsIgnoreString implementation
func ContainsIgnoreString(in string) bool {
	if !strings.Contains(in, "nolint:") {
		return false
	}

	parts := strings.Split(in, ":")

	for _, part := range parts[1:] {
		if strings.Contains(part, ignoreString) {
			return true
		}
	}

	return false
}

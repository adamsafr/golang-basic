package str

import (
	"strings"
	"unicode/utf8"
)

// IsEmpty ...
func IsEmpty(value string) bool {
	return utf8.RuneCountInString(strings.TrimSpace(value)) == 0
}

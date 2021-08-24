// Package strings contain function that can return length of string and revert string
package strings

import (
	"strings"
	"unicode/utf8"
)

// ReverseString reverse provided string, can be used for UTF8 characters.
func ReverseString(str string) string {
	var sb strings.Builder

	sb.Grow(len(str))

	lengthToProcess := len(str)

	for lengthToProcess > 0 {
		r, width := utf8.DecodeLastRuneInString(str[:lengthToProcess])
		sb.WriteRune(r)

		lengthToProcess -= width
	}

	return sb.String()
}

// StringLength returns length of string
// Can be used for UTF-8 string
func StringLength(str string) int {
	return utf8.RuneCountInString(str)
}

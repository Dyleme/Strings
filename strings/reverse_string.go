// Package contain function that can reverse string
package strings

import (
	"strings"
	"unicode/utf8"
)

// Reverse provided string, can be used for UTF8 characters.
// Return reversed string
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

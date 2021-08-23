// Package strings contain function that can return length of string and revert string
package strings

import "unicode/utf8"

// StringLength returns length of string
// Can be used for UTF-8 string
func StringLength(str string) int {
	return utf8.RuneCountInString(str)
}

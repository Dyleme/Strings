//Package contain function that can reverse string
package strings

import (
	"strings"
	"unicode/utf8"
)

//Reverse provided string can be usign for UTF8
//return reversed string
func ReverseString(str string) string {
	var sb strings.Builder

	sb.Grow(len(str))

	b := []byte(str)
	length := len(str)
	composeWidth := 0

	for composeWidth < length {
		r, width := utf8.DecodeLastRune(b[:length-composeWidth])
		sb.WriteRune(r)

		composeWidth += width
	}

	return sb.String()
}

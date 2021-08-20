package reversestring

import (
	"strings"
	"unicode/utf8"
)

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

package reverseString

import (
	"strings"
	"unicode/utf8"
)

func ReverseString(str string) string {
	var sb strings.Builder
	b := []byte(str)
	length := len(str)
	sb.Grow(length)
	composeWidth := 0
	for composeWidth < length {
		r, width := utf8.DecodeLastRune(b[:length-composeWidth])
		sb.WriteRune(r)
		composeWidth += width
	}
	return sb.String()
}

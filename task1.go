package task1

import (
	"strings"
)

func ReverseString(str string) string {
	var sb strings.Builder
	sb.Grow(len(str))
	for i := len(str) - 1; i >= 0; i-- {
		sb.WriteByte(str[i])
	}
	return sb.String()
}

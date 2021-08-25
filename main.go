package main

import (
	"fmt"
	"strings"

	mystrings "github.com/Dyleme/Strings/strings"
)

func main() {
	var str = "double  spaces"

	var sb strings.Builder

	sb.Grow(len(str))

	tempStr := str

	for {
		index := strings.IndexRune(tempStr, ' ')

		if index != -1 {
			sb.WriteString(mystrings.ReverseString(tempStr[:index]))
			sb.WriteRune(' ')
		} else {
			sb.WriteString(mystrings.ReverseString(tempStr))

			break
		}

		tempStr = tempStr[index+1:]
	}

	fmt.Println(str)
	fmt.Println(sb.String())
}

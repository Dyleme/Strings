package task1

import (
	"testing"
)

func TestReverseString(t *testing.T) {
	t.Run("Hello world", testReverseStringFunc("Hello world!", "!dlrow olleH"))
	t.Run("123MaiN678", testReverseStringFunc("123MaiN678", "876NiaM321"))
}

func testReverseStringFunc(str string, expexted string) func(t *testing.T) {
	return func(t *testing.T) {
		if got := ReverseString(str); got != expexted {
			t.Errorf("Reverse string = %s, want %s", got, expexted)
		}
	}
}

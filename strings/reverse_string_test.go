package strings_test

import (
	"fmt"
	"testing"

	"github.com/Dyleme/Strings/strings"
)

func TestReverseString(t *testing.T) {
	t.Run("Hello world", testReverseStringFunc("Hello world!", "!dlrow olleH"))
	t.Run("123MaiN678", testReverseStringFunc("123MaiN678", "876NiaM321"))
	t.Run("Привет", testReverseStringFunc("Привет", "тевирП"))
	t.Run("日本語", testReverseStringFunc("日本語", "語本日"))
	t.Run("Void string", testReverseStringFunc("", ""))
}

func testReverseStringFunc(str, expexted string) func(t *testing.T) {
	return func(t *testing.T) {
		t.Helper()
		t.Parallel()

		if got := strings.ReverseString(str); got != expexted {
			t.Errorf("Reverse string = %s, want %s", got, expexted)
		}
	}
}

func ExampleReverseString() {
	var str = "Nice weather today"

	fmt.Println(strings.ReverseString(str))

	// Output:yadot rehtaew eciN
}

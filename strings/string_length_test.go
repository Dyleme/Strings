package strings_test

import (
	"fmt"
	"testing"

	"github.com/Dyleme/Strings/strings"
)

var flagTest = []struct {
	in  string
	out int
}{
	{"Hello", 5},
	{"日本語", 3},
	{"Привет", 6},
	{"Hello мир!", 10},
	{"", 0},
}

func TestStringLength(t *testing.T) {
	for _, tt := range flagTest {
		t.Run(tt.in, func(t *testing.T) {
			if result := strings.StringLength(tt.in); result != tt.out {
				t.Errorf("Amount of characters got %v, want %v", result, tt.out)
			}
		})
	}
}

func ExampleStringLength() {
	var str = "Nice weather today"

	fmt.Println(strings.StringLength(str))

	// Output: 18
}

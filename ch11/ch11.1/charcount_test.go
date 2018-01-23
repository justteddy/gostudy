package chr

import (
	"bufio"
	"strings"
	"testing"
)

func TestCharcount(t *testing.T) {
	var tests = []struct {
		input string
		char  rune
		count int
	}{
		{"abcd", 'a', 1},
		{"aaabbbbcccccc", 'b', 4},
		{"⌘⌘⌘⌘⌘", '⌘', 5},
		{"hello world\n hello", 'l', 5},
		{"word word\n and one more\n", '\n', 2},
		{"你好你好some words here", '好', 2},
	}

	for _, test := range tests {
		in := bufio.NewReader(strings.NewReader(test.input))
		counts, _, _, _ := charCount(in)
		if counts[test.char] != test.count {
			t.Errorf("charCount(%s) = %v, expected %v", test.input, test.count, counts[test.char])
		}
	}

}

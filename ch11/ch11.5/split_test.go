package split

import (
	"strings"
	"testing"
)

func TestSplit(t *testing.T) {
	tests := []struct {
		input   string
		sep     string
		wantcnt int
	}{
		{"a:b:c", ":", 3},
		{"a:b:c", ".", 1},
		{"some,words,here,are", ",", 4},
		{"some,,,words", ",", 4},
	}

	for _, test := range tests {
		if got := len(strings.Split(test.input, test.sep)); got != test.wantcnt {
			t.Errorf("Split(%s, %s) return %d words, %d expected", test.input, test.sep, got, test.wantcnt)
		}
	}
}

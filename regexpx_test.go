package regexpx_test

import (
	"regexp"
	"testing"

	rx "github.com/yargevad/regexpx"
)

var testMatch = rx.RegexpSet{
	regexp.MustCompile(`^abc+$`),
	regexp.MustCompile(`^abc+d$`),
}

type MatchTest struct {
	Input string
	Match bool
	Index int
}

func TestMatch(t *testing.T) {
	for _, test := range []MatchTest{
		{"", false, -1},
		{"a", false, -1},
		{"abc", true, 0},
		{"abcc", true, 0},
		{"abcd", true, 1},
		{"abcde", false, -1},
	} {
		actual, idx := testMatch.Match(test.Input)
		if actual != test.Match {
			t.Fatalf("string [%s] expected %t actual %t", test.Input, test.Match, actual)
		} else if idx != test.Index {
			t.Fatalf("string [%s] expected index %d, actual %d", test.Input, test.Index, idx)
		}
	}
}

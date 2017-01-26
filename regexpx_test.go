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

type SplitTest struct {
	Input  string
	Index  int
	Output []string
}

var testSplit = rx.RegexpSet{
	regexp.MustCompile(`[aeiou]`),
}

func TestSplit(t *testing.T) {
	for _, test := range []SplitTest{
		{"paqerisotu", 0, []string{"p", "q", "r", "s", "t", ""}},
		{"fffffff", -1, nil},
	} {
		out, idx := testSplit.Split(test.Input, -1)
		if idx != test.Index {
			t.Fatalf("string [%s] expected index %d, actual %d", test.Input, test.Index, idx)
		} else if len(out) != len(test.Output) {
			t.Fatalf("string [%s] expected length %d, actual %d", test.Input, len(test.Output), len(out))
		} else {
			for i, elt := range out {
				if test.Output[i] != elt {
					t.Fatalf("string [%s] expected %q, actual %q", test.Input, test.Output, out)
				}
			}
		}
	}
}

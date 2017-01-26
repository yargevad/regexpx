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

type Expected struct {
	Match bool
	Index int
}

func TestMatch(t *testing.T) {
	expected := []Expected{{false, -1}, {false, -1}, {true, 0}, {true, 0}, {true, 1}, {false, -1}}
	for i, str := range []string{"", "a", "abc", "abcc", "abcd", "abcde"} {
		actual, idx := testMatch.Match(str)
		if actual != expected[i].Match {
			t.Fatalf("string [%s] expected %t actual %t", str, expected[i].Match, actual)
		} else if idx != expected[i].Index {
			t.Fatalf("string [%s] expected index %d, actual %d", str, expected[i].Index, idx)
		}
	}
}

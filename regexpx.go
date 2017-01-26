package regexpx

import (
	"regexp"
)

type RegexpSet []*regexp.Regexp

// Match returns true on a match, along with the index of the matching pattern.
// When no matches are found, an index of -1 is returned.
func (rs *RegexpSet) Match(s string) (bool, int) {
	for idx, r := range []*regexp.Regexp(*rs) {
		if r.FindStringIndex(s) != nil {
			return true, idx
		}
	}
	return false, -1
}

// Split slices s into all substrings separated by the first matching expression in the set.
// It returns a slice of the substrings between the matches, along with the index of the matching pattern.
// When called on an expression that contains no metacharacters, it is equivalent to strings.SplitN.
// See also regexp.Split and strings.SplitN.
func (rs *RegexpSet) Split(s string, n int) ([]string, int) {
	for idx, r := range []*regexp.Regexp(*rs) {
		if r.FindStringIndex(s) != nil {
			return r.Split(s, n), idx
		}
	}
	return nil, -1
}

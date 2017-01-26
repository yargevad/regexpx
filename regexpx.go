package regexpx

import (
	"regexp"
)

type RegexpSet []*regexp.Regexp

// Match returns true if any expression matches s, along with the corresponding index.
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
// It returns a slice of the substrings between the matches, along with the corresponding index.
// See also regexp.Split.
func (rs *RegexpSet) Split(s string, n int) ([]string, int) {
	for idx, r := range []*regexp.Regexp(*rs) {
		if r.FindStringIndex(s) != nil {
			return r.Split(s, n), idx
		}
	}
	return nil, -1
}

// Replace returns a copy of s, replacing matches with the second argument, repl.
// The index of the matching expression is also returned, or -1 when no expression matches.
// See also regexp.ReplaceAllString
func (rs *RegexpSet) Replace(s, repl string) (string, int) {
	for idx, r := range []*regexp.Regexp(*rs) {
		if r.FindStringIndex(s) != nil {
			return r.ReplaceAllString(s, repl), idx
		}
	}
	return s, -1
}

// ReplaceSubmatch returns a copy of s, replacing matches with the second argument, repl.
// The index of the matching expression is also returned, or -1 when no expression matches.
// Any submatches (aka capturing groups) are also returned
// See also the Submatch section of the regexp docs and regexp.FindStringSubmatch
func (rs *RegexpSet) ReplaceSubmatch(s, repl string) (string, int, []string) {
	for idx, r := range []*regexp.Regexp(*rs) {
		if m := r.FindStringSubmatch(s); m != nil {
			return r.ReplaceAllString(s, repl), idx, m
		}
	}
	return s, -1, nil
}

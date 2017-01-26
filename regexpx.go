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

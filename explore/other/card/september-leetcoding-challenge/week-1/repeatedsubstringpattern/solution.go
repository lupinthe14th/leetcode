package repeatedsubstringpattern

import (
	"strings"
)

// SeeAlso: https://leetcode.com/problems/repeated-substring-pattern/discuss/94334/Easy-python-solution-with-explaination
func repeatedSubstringPattern(s string) bool {
	l := len(s)
	if l == 0 {
		return false
	}
	ss := (s + s)[1 : l*2-1]
	return strings.Contains(ss, s)
}

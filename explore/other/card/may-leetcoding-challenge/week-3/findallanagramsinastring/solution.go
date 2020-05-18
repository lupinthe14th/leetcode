package findallanagramsinastring

import (
	"reflect"
)

// SeeAlso: https://leetcode.com/problems/find-all-anagrams-in-a-string/discuss/491040/go-O(N)-sliding-window
func findAnagrams(s string, p string) []int {
	if len(s) < len(p) {
		return nil
	}

	pat := make([]int, 26)
	mem := make([]int, 26)

	for i := range p {
		pat[p[i]-'a']++
		mem[s[i]-'a']++
	}
	out := make([]int, 0, len(s))
	for i := 0; i < len(s)-len(p)+1; i++ {
		if reflect.DeepEqual(pat, mem) {
			out = append(out, i)
		}

		if i+len(p) < len(s) {
			mem[s[i]-'a']--
			mem[s[i+len(p)]-'a']++
		}
	}
	return out
}

package wordpattern

import (
	"strings"
)

func wordPattern(pattern string, str string) bool {
	if len(pattern) != len(strings.Split(str, " ")) {
		return false
	}
	memo := make(map[int]string, 0)

	for i, s := range strings.Split(str, " ") {
		idx := int(pattern[i] - 'a')
		if v, ok := memo[idx]; !ok {
			memo[idx] = s
		} else {
			if v != s {
				return false
			}
		}
	}
	for i := 0; i < 26; i++ {
		base, ok := memo[i]
		if !ok {
			continue
		}
		for j := i + 1; j < 26; j++ {
			check, ok := memo[j]
			if !ok {
				continue
			}
			if base == check {
				return false
			}
		}
	}
	return true
}

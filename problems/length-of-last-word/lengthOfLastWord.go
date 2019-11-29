package lengthoflastword

import "strings"

func lengthOfLastWord(s string) int {
	if s == "" {
		return 0
	}

	lw := strings.Split(s, " ")

	if len(lw) <= 1 {
		return 0
	}

	l := strings.Split(lw[len(lw)-1], "")

	return len(l)
}

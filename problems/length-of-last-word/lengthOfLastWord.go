package lengthoflastword

import "strings"

func lengthOfLastWord(s string) int {
	if s == "" {
		return 0
	}

	ss := strings.Split(s, " ")

	var l int
	for i := len(ss) - 1; i >= 0; i-- {
		if ss[i] != "" {
			lw := strings.Split(ss[i], "")
			l = len(lw)
			break
		}
	}
	return l
}

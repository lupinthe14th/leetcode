package lengthoflastword

import (
	"strings"
)

func lengthOfLastWord(s string) int {
	ss := strings.Split(strings.TrimSpace(s), " ")
	l := len(ss)
	return len(ss[l-1])
}

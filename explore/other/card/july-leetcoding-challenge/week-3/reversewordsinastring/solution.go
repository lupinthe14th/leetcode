package reversewordsinastring

import (
	"strings"
)

func reverseWords(s string) string {
	ss := make([]string, 0)
	for _, v := range strings.Split(s, " ") {
		if len(v) != 0 {
			ss = append(ss, v)
		}
	}
	for i := len(ss)/2 - 1; i >= 0; i-- {
		opp := len(ss) - 1 - i
		ss[i], ss[opp] = ss[opp], ss[i]
	}
	return strings.Join(ss, " ")
}

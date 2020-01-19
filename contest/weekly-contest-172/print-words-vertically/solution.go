package main

import (
	"strings"
)

func printVertically(s string) []string {

	sl := strings.Split(s, " ")
	ml := maxl(sl)
	ans := make([]string, ml)

	for _, u := range sl {
		ul := len(u)
		for j := 0; j < ml; j++ {
			if j >= ul {
				ans[j] += " "
			} else {
				ans[j] += string(u[j])
			}
		}
	}

	// Trim trailing spaces
	for i := range ans {
		ans[i] = strings.TrimRight(ans[i], " ")
	}
	return ans
}

func maxl(ss []string) int {
	var ml int = -1 << 31
	for _, s := range ss {
		ml = max(ml, len(s))
	}
	return ml
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {}

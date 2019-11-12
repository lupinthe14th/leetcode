package conntandsay

import (
	"strconv"
	"strings"
)

func countAndSay(n int) string {
	if n == 1 {
		return strconv.Itoa(1)
	}
	stack := make([]string, 0, n-1)
	if n > 1 {
		s := countAndSay(n - 1)
		if len(s) == 1 {
			stack = append(stack, strconv.Itoa(len(s)))
			stack = append(stack, string(s[len(s)-1]))
		} else {
			count := 1
			for i := 0; i < len(s)-1; i++ {
				if string(s[i]) == string(s[i+1]) {
					count++
				} else {
					stack = append(stack, strconv.Itoa(count))
					stack = append(stack, string(s[i]))
					count = 1
				}
				if i == len(s)-2 {
					stack = append(stack, strconv.Itoa(count))
					stack = append(stack, string(s[i+1]))
				}

			}
		}
	}
	return strings.Join(stack, "")
}

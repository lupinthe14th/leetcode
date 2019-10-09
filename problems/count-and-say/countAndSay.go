package conntandsay

import (
	"log"
	"strconv"
	"strings"
)

func countAndSay(n int) string {
	if n == 1 {
		return strconv.Itoa(1)
	}
	count := 0
	var stack = make([]string, 0)
	if n > 1 {
		s := countAndSay(n - 1)
		i := 0
		for _, v := range s[1:] {
			log.Printf("i: %d", s[i])
			log.Printf("j: %d", v)
			if string(s[i]) == string(v) {
				count++
			}

			if count != 1 {
				stack = append(stack, string(s[i]))
				stack = append(stack, strconv.Itoa(count))
				count = 0
			} else {
				stack = append(stack, strconv.Itoa(count))
				stack = append(stack, string(s[i]))
			}

		}
	}
	log.Print(stack)
	return strings.Join(stack, "")
}

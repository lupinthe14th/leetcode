package main

import (
	"strconv"
	"strings"
)

func maximum69Number(num int) int {

	s := strconv.Itoa(num)
	sl := strings.Split(s, "")

	for i, v := range sl {
		if v != "9" {
			sl[i] = "9"
			break
		}
	}

	n, _ := strconv.Atoi(strings.Join(sl, ""))

	return n
}
func main() {}

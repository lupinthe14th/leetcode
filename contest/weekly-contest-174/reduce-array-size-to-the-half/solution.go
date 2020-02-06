package reducearraysizetothehalf

import (
	"sort"
)

func minSetSize(arr []int) int {
	m := make(map[int]int)
	for _, v := range arr {
		m[v]++
	}

	l := len(m)
	s := make([]int, 0, l)
	for _, v := range m {
		s = append(s, v)
	}

	sort.Ints(s)
	al := len(arr)
	var sum, c int
	for i := l - 1; i >= 0; i-- {
		sum += s[i]
		c++
		if sum >= al/2 {
			break
		}
	}

	return c
}

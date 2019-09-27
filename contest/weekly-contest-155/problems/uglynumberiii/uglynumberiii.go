package nthuglynumber

import (
	"sort"
)

// LCM(least common multipul)
func lcm(x, y, z int) int {
	return x * y * z
}

func nthUglyNumber(n int, a int, b int, c int) int {
	m := make(map[int]int, 0)

	l := lcm(a, b, c)
	var divA, divB, divC int
	for i := 1; ; i++ {
		divA = a * i
		divB = b * i
		divC = c * i
		if divA < l {
			m[divA]++
		}

		if divB < l {
			m[divB]++
		}
		if divC < l {
			m[divC]++
		}
		if divA > l && divB > l && divC > l {
			break
		}
	}

	numbers := make([]int, 0)

	for number := range m {
		numbers = append(numbers, number)
	}
	sort.Ints(numbers)
	return numbers[n]
}

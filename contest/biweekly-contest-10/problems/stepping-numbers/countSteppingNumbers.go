package intersectionofthreesortedarrays

import (
	"math"
	"strconv"
)

func countSteppingNumbers(low int, high int) []int {
	stack := make([]int, 0)

	for i := low; i <= high; i++ {
		if i <= 10 {
			stack = append(stack, i)
		} else {
			j := 0
			s := strconv.Itoa(i)
			for k := 1; k < len(s); k++ {
				if math.Abs(float64(s[k])-float64(s[j])) != 1 {
					break
				}
				if k == len(s)-1 {
					stack = append(stack, i)
				}
				j++
			}
		}

	}
	return stack
}

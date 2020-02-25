package closestdivisors

import "math"

// refarence: https://leetcode.com/problems/closest-divisors/discuss/517595/JavaC++Python-Easy-and-Concise
func closestDivisors(num int) []int {
	for a := int(math.Sqrt(float64(num) + 2)); a > 0; a-- {
		if (num+1)%a == 0 {
			return []int{a, (num + 1) / a}
		}
		if (num+2)%a == 0 {
			return []int{a, (num + 2) / a}
		}
	}
	return []int{}
}

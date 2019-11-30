package removeinterval

func removeInterval(intervals [][]int, toBeRemoved []int) [][]int {
	stack := make([][]int, 0)

	A := toBeRemoved[0]
	B := toBeRemoved[1]
	for _, iv := range intervals {
		if !intersect(iv, toBeRemoved) {
			stack = append(stack, iv)
		} else {

			a := iv[0]
			b := iv[1]

			if A <= a && a <= b && b <= B {
				continue
			} else if a <= A && A <= b && b <= B {
				stack = append(stack, []int{a, A})
			} else if A <= a && a <= B && B <= b {
				stack = append(stack, []int{B, b})
			} else if a <= A && A <= B && B <= b {
				stack = append(stack, []int{a, A})
				stack = append(stack, []int{B, b})
			}
		}
	}

	return stack
}

func intersect(s1, s2 []int) bool {
	x1, y1 := s1[0], s1[1]
	x2, y2 := s2[0], s2[1]

	if (x1 <= x2 && x2 <= y1) || (x1 <= y2 && y2 <= y1) || (x2 <= x1 && x1 <= y1 && y1 <= y2) {
		return true
	}
	return false
}

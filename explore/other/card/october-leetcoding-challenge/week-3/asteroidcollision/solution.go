package asteroidcollision

func asteroidCollision(asteroids []int) []int {
	stack := []int{}
	for _, asteroid := range asteroids {
		if len(stack) == 0 || asteroid > 0 {
			stack = append(stack, asteroid)
		} else {
			for {
				top := stack[len(stack)-1]
				if top < 0 {
					stack = append(stack, asteroid)
					break
				} else if top == -asteroid {
					stack = stack[:len(stack)-1]
					break
				} else if top > -asteroid {
					break
				} else {
					stack = stack[:len(stack)-1]
					if len(stack) == 0 {
						stack = append(stack, asteroid)
						break
					}
				}
			}
		}
	}
	out := make([]int, len(stack))
	for i := len(stack) - 1; i >= 0; i-- {
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		out[i] = cur
	}
	return out
}

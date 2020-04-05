package countlargestgroup

func countLargestGroup(n int) int {
	m := -1

	for i := 1; i <= n; i++ {
		n := number(i)
		m = max(m, n)
	}
	stack := make([][]int, m)
	for i := 1; i <= n; i++ {
		n := number(i)
		stack[n-1] = append(stack[n-1], i)
	}

	ml := -1
	for i := range stack {
		ml = max(ml, len(stack[i]))
	}
	c := 0
	for i := range stack {
		if len(stack[i]) == ml {
			c++
		}
	}
	return c
}

func number(n int) int {
	sum := 0
	for n > 0 {
		sum += n % 10
		n /= 10
	}
	return sum
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

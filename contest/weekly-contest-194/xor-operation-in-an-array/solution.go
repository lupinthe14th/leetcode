package xoroperationinanarray

func xorOperation(n int, start int) int {

	nums := make([]int, n)
	for i := range nums {
		nums[i] = start + 2*i
	}

	out := nums[0]
	for i := 1; i < n; i++ {
		out ^= nums[i]
	}
	return out
}

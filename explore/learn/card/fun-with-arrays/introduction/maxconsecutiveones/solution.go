package maxconsecutiveones

func findMaxConsecutiveOnes(nums []int) int {
	out, c := 0, 0
	for _, n := range nums {
		if n == 1 {
			c++
		} else {
			c = 0
		}
		out = max(out, c)
	}
	return out
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

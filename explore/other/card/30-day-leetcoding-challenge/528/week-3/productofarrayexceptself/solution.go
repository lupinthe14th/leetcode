package productofarrayexceptself

func productExceptSelf(nums []int) []int {
	c, allProd := prod(nums)
	switch {
	case c == 0:
		for i, n := range nums {
			nums[i] = allProd / n
		}
	case c == 1:
		for i, n := range nums {
			if n == 0 {
				nums[i] = allProd
			} else {
				nums[i] = 0
			}
		}
	case c == 2:
		for i := range nums {
			nums[i] = 0
		}
	}
	return nums
}

// c: 零の数
// p: 一つの零を除いた積
func prod(nums []int) (c, p int) {
	c, p = 0, 1
	for _, n := range nums {
		if n == 0 {
			c++
			continue
		}
		p *= n
	}

	if c >= 2 {
		return 2, 0
	}
	return c, p
}

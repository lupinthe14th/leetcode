package productofarrayexceptself

func productExceptSelf(nums []int) []int {
	o := make([]int, len(nums))
	for i := range nums {
		o[i] = prod(nums, i)
	}
	return o
}

func prod(nums []int, num int) int {
	ans := 1

	for i, n := range nums {
		if i != num {
			ans *= n
		}
	}
	return ans
}

package maximumproductsubarray

func maxProduct(nums []int) int {
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	out := nums[0]
	for i := range nums {
		cur := nums[i]
		out = max(out, cur)
		for j := i + 1; j < len(nums); j++ {
			cur *= nums[j]
			out = max(out, cur)
		}
	}
	return out
}

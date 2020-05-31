package maximumproductoftwoelementsinanarray

func maxProduct(nums []int) int {
	out := -1 << 31
	l := len(nums)
	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			out = max(out, (nums[i]-1)*(nums[j]-1))
		}
	}
	return out
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

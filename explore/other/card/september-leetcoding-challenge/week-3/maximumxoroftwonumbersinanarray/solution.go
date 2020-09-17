package maximumxoroftwonumbersinanarray

func findMaximumXOR(nums []int) int {
	out := 0
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	for i := range nums {
		for j := range nums {
			out = max(out, nums[i]^nums[j])
		}
	}
	return out
}

package maximumproductsubarray

// SeeAlso: https://leetcode.com/explore/challenge/card/september-leetcoding-challenge/555/week-2-september-8th-september-14th/3456/discuss/48230/Possibly-simplest-solution-with-O(n)-time-complexity
func maxProduct(nums []int) int {
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	min := func(x, y int) int {
		if x < y {
			return x
		}
		return y
	}
	// store the result that is the max we have found so far
	out := nums[0]

	// imax/imin store the max/min product of
	// subarray that ends with the current number nums[i]
	for i, imax, imin := 1, out, out; i < len(nums); i++ {
		// multiplied by a negative makes big number smaller, small number bigger
		// so we redefine the extremums by swapping them
		if nums[i] < 0 {
			imax, imin = imin, imax
		}

		// max/min product for the current number is either the current number itself
		// or the max/min by the previous number times the current one
		imax = max(nums[i], imax*nums[i])
		imin = min(nums[i], imin*nums[i])

		// the newly computed max value is a candidate for our global result
		out = max(out, imax)
	}
	return out
}

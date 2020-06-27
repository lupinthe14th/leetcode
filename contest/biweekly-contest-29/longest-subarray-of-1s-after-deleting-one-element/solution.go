package longestsubarrayof1safterdeletingoneelement

// https://leetcode.com/problems/longest-subarray-of-1s-after-deleting-one-element/discuss/708112/JavaC++Python-Sliding-Window-at-most-one-0
func longestSubarray(nums []int) int {
	out := 0
	for i, j, k := 0, 0, 1; i < len(nums); i++ {
		if nums[i] == 0 {
			k--
		}
		for k < 0 {
			if nums[j] == 0 {
				k++
			}
			j++
		}
		out = max(out, i-j)
	}
	return out
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

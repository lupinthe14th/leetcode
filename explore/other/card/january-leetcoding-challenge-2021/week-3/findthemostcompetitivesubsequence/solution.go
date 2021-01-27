package findthemostcompetitivesubsequence

func mostCompetitive(nums []int, k int) []int {
	out := make([]int, 0, k)
	for i := 0; i < len(nums); i++ {
		for len(out) > 0 && out[len(out)-1] > nums[i] && len(out)-1+len(nums)-i >= k {
			out = out[:len(out)-1]
		}
		if len(out) < k {
			out = append(out, nums[i])
		}
	}
	return out
}

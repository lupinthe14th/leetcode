package numberofgoodpairs

func numIdenticalPairs(nums []int) int {
	out := 0
	for i := range nums {
		for j := range nums {
			if i < j && nums[i] == nums[j] {
				out++

			}
		}
	}
	return out
}

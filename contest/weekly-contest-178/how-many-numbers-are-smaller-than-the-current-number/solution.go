package howmanynumbersaresmallerthanthecurrentnumber

func smallerNumbersThanCurrent(nums []int) []int {
	l := len(nums)
	var ans = make([]int, l)
	for i := 0; i < l; i++ {
		var c int
		for j := 0; j < l; j++ {
			if j != i && nums[j] < nums[i] {
				c++
			}
		}
		ans[i] = c
	}
	return ans
}

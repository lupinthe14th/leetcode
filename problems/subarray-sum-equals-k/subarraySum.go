package subarraysumequalsk

func subarraySum(nums []int, k int) int {
	var c, sum int

	rec := make(map[int]int) // prefix sum recorder
	rec[0]++                 // to take into account those subarrays that begin with index 0

	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		c += rec[sum-k]
		rec[sum]++
	}
	return c
}

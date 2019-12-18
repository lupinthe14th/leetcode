package subarraysumequalsk

func subarraySum(nums []int, k int) int {
	// 0 - len(nums)までの合計をハッシュにする
	l := len(nums)
	dp := make(map[int]int, l)
	dp[0] = nums[0]

	for i := 1; i < l; i++ {
		dp[i] = dp[i-1] + nums[i]
	}

	var c int
	const MaxUint = ^uint(0)
	const MaxInt = int(MaxUint >> 1)
	for i := 0; i < l; i++ {
		for j := 0; j < l; j++ {
			sum := MaxInt
			if i == j {
				sum = dp[i]
			} else if i < j {
				// sum(i,j)=sum(0,j)-sum(0,i), where sum(i,j) represents the sum of all the elements from index i to j-1.
				sum = dp[j] - dp[i]
			}
			if k == sum {
				c++
			}
		}
	}
	return c
}

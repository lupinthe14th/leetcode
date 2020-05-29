package countingbits

// SeeAlso: https://leetcode.com/explore/challenge/card/may-leetcoding-challenge/537/week-4-may-22nd-may-28th/3343/discuss/79539/Three-Line-Java-Solution
// An easy recurrence for this problem is f[i] = f[i/2] + i%2
func countBits(num int) []int {
	out := make([]int, num+1)
	for i := 1; i <= num; i++ {
		out[i] = out[i>>1] + i&1
	}
	return out
}

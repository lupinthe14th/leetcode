package findtwononoverlappingsubarrayseachwithtargetsum

// https://leetcode.com/problems/find-two-non-overlapping-sub-arrays-each-with-target-sum/discuss/685548/Java-Sliding-Window-with-dp-O(N)-20-lines
func minSumOfLengths(arr []int, target int) int {
	dp := make(map[int]int)
	sum := 0

	dp[0] = -1

	memo := make([]int, len(arr))
	const MAX = 1 << 31
	res := MAX

	for i := 0; i < len(arr); i++ {
		sum += arr[i]
		if i > 0 {
			memo[i] = memo[i-1]
		} else {
			memo[i] = MAX
		}

		if pre, ok := dp[sum-target]; ok {
			memo[i] = min(memo[i], i-pre)
			if pre != -1 && memo[pre] != MAX {
				res = min(res, memo[pre]+i-pre)
			}
		}
		dp[sum] = i

	}
	if res == MAX {
		return -1
	}
	return res
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

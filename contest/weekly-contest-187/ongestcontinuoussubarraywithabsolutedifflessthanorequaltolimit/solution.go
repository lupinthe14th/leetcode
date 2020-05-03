package ongestcontinuoussubarraywithabsolutedifflessthanorequaltolimit

// https://leetcode.com/problems/longest-continuous-subarray-with-absolute-diff-less-than-or-equal-to-limit/discuss/609690/Java-13-lines-with-detailed-explanation
// idea:
// For each number, maintain localMin and localMax, try to extend its rightmost
// border as far as possible. Remember to skip the continuous duplicated
// numbers.
//
// code:
// 1. initialize maxLen = 0, denote l = length
// 2. for i from 0 to l:
//    1. let localMax = nums[i], localMin = nums[i] and initialize j = i+1
//    2. while j<l:
//       1. update localMax and localMin with nums[j]
//       2. if difference of localMax and localMin is bigger than limit, it doesn't qualify anymore, just break.
//       3. otherwise j++
//    3. update maxLen with j-i
//    4. skip continuous nums[i]
// 3. return maxLen in the end
func longestSubarray(nums []int, limit int) int {
	maxLen, l := 0, len(nums)

	for i := 0; i < l; i++ {
		localMax := nums[i]
		localMin := nums[i]
		j := i + 1
		for j < l {
			localMax = max(localMax, nums[j])
			localMin = min(localMin, nums[j])
			if localMax-localMin > limit {
				break
			}
			j++
		}
		maxLen = max(maxLen, j-i)
		for i < l-1 && nums[i] == nums[i+1] {
			i++
		}
	}
	return maxLen
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

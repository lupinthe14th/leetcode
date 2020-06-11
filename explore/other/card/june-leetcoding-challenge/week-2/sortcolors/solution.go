package sortcolors

// SeeAlso
// https://leetcode.com/explore/featured/card/june-leetcoding-challenge/540/week-2-june-8th-june-14th/3357/discuss/26549/Java-solution-both-2-pass-and-1-pass
func sortColors(nums []int) {
	p1, p2, i := 0, len(nums)-1, 0
	for i <= p2 {
		switch nums[i] {
		case 0:
			nums[i] = nums[p1]
			nums[p1] = 0
			p1++
		case 2:
			nums[i] = nums[p2]
			nums[p2] = 2
			p2--
			i--
		}
		i++
	}
}

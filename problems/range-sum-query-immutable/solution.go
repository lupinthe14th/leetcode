package rangesumqueryimmutable

// NumArray is ...
type NumArray struct {
	nums []int
}

// Constructor is ...
func Constructor(nums []int) NumArray {
	return NumArray{nums: nums}
}

// SumRange is given an integer array nums, find the sum of the elements between indices i and j (i ≤ j), inclusive.
func (na *NumArray) SumRange(i int, j int) int {

	var ans int
	for i := i; i <= j; i++ {
		ans += na.nums[i]
	}
	return ans
}

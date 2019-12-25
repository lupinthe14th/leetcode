package rangesumqueryimmutable

// NumArray is ...
type NumArray struct {
	nums []int
}

// Constructor is ...
func Constructor(nums []int) NumArray {
	l := len(nums)
	sum := make([]int, l+1)
	for i := 0; i < l; i++ {
		sum[i+1] = sum[i] + nums[i]
	}
	return NumArray{nums: sum}
}

// SumRange is given an integer array nums, find the sum of the elements between indices i and j (i ≤ j), inclusive.
func (na *NumArray) SumRange(i int, j int) int {
	return na.nums[j+1] - na.nums[i]
}

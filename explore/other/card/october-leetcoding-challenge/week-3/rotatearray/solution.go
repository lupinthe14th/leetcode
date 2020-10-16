package rotatearray

func rotate(nums []int, k int) {
	t := k % len(nums)
	tmp := make([]int, len(nums))
	tmp = append(nums[len(nums)-t:], nums[:len(nums)-t]...)
	copy(nums, tmp)
}

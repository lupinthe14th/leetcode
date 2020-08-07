package findallduplicatesinanarray

// when find a number i, flip the number at position i-1 to negative.
// if the number at position i-1 is already negative, i is the number that occurs twice.
func findDuplicates(nums []int) []int {
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}
	out := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		idx := abs(nums[i]) - 1
		if nums[idx] < 0 {
			out = append(out, idx+1)
		}
		nums[idx] = -nums[idx]
	}
	return out
}

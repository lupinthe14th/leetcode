package movezeros

func moveZeroes(nums []int) {
	i := 0
	for _, x := range nums {
		if x != 0 {
			nums[i] = x
			i++
		}
	}
	for j := i; j < len(nums); j++ {
		nums[j] = 0
	}
}

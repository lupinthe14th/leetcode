package checkifall1sareatleastlengthkplacesaway

func kLengthApart(nums []int, k int) bool {
	cur := 0
	flag := false
	for i := range nums {
		if !flag && nums[i] == 1 {
			cur = i
			flag = true
			continue
		}
		if nums[i] == 1 {
			if cur != i && i-cur-1 < k {
				return false
			}
			cur = i
		}
	}
	return true
}

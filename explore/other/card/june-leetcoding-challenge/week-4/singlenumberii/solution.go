package singlenumberii

func singleNumber(nums []int) int {
	one, two := 0, 0
	for i := range nums {
		one = one ^ nums[i] & ^two
		two = two ^ nums[i] & ^one
	}
	return one
}

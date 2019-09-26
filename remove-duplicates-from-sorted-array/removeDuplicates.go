package removeduplicatesfromsortedarray

func removeDuplicates(nums []int) int {
	m := make(map[int]int, 0)
	for _, n := range nums {
		m[n]++
	}
	return len(m)
}

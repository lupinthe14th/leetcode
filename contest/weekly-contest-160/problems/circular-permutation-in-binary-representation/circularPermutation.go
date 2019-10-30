package circularpermutationinbinaryrepresentation

func circularPermutation(n int, start int) []int {
	p, l := make([]int, 0), int(1<<uint(n))

	for i := 0; i < l; i++ {
		p = append(p, start^i^i>>1)
	}
	return p
}

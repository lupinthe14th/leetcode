package circularpermutationinbinaryrepresentation

func circularPermutation(n int, start int) []int {
	l := 1 << uint(n)
	p := make([]int, 0, l)

	for i := 0; i < l; i++ {
		p = append(p, start^i^i>>1)
	}
	return p
}

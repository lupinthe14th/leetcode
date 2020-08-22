package sortarraybyparity

func sortArrayByParity(A []int) []int {

	out := make([]int, 0, len(A))

	for _, n := range A {
		if n%2 == 0 {
			out = append([]int{n}, out...)
		} else {
			out = append(out, n)
		}
	}
	return out
}

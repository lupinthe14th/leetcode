package sortarraybyparity

func sortArrayByParity(A []int) []int {

	out := make([]int, len(A))

	i := 0
	for _, n := range A {
		if n%2 == 0 {
			out[i] = n
			i++
		}
	}
	for _, n := range A {
		if n%2 != 0 {
			out[i] = n
			i++
		}
	}
	return out
}

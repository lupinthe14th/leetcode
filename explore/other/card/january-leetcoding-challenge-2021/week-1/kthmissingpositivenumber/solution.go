package kthmissingpositivenumber

func findKthPositive(arr []int, k int) int {
	i, n := 0, 1
	for k > 0 {
		if i == len(arr) {
			break
		}

		if arr[i] != n {
			k--
		} else {
			i++
		}
		n++
	}
	return n + k - 1
}

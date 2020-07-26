package duplicatezeros

func duplicateZeros(arr []int) {
	p := 0
	l := len(arr) - 1
	for i := 0; i <= l-p; i++ {
		if arr[i] == 0 {
			if i == l-p {
				arr[l] = 0
				l--
				break
			}
			p++
		}
	}

	last := l - p

	for i := last; i >= 0; i-- {
		if arr[i] == 0 {
			arr[i+p] = 0
			p--
			arr[i+p] = 0
		} else {
			arr[i+p] = arr[i]
		}
	}
}

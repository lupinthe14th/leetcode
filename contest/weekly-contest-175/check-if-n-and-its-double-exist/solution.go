package checkifnanditsdoubleexist

func checkIfExist(arr []int) bool {
	l := len(arr)
	for i := 0; i < l; i++ {
		for j := 0; j < l; j++ {
			if i == j {
				continue
			}
			if arr[i] == 2*arr[j] {
				return true
			}
		}
	}
	return false
}

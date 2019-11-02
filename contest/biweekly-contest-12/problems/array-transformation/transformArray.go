package arraytransformation

func transformArray(arr []int) []int {
	changed := true
	for changed {
		changed = false
		for i := 1; i < len(arr)-1; i++ {
			if arr[i] < arr[i-1] && arr[i] < arr[i+1] {
				changed = true
				arr[i] = arr[i] + 1
			} else if arr[i] > arr[i-1] && arr[i] > arr[i+1] {
				changed = true
				arr[i] = arr[i] - 1
			}
		}
	}
	return arr
}

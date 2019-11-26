package arraytransformation

func transformArray(arr []int) []int {
	changed := true
	for changed {
		list := make([]int, 0)
		changed = false
		for i := 0; i < len(arr); i++ {
			if i == 0 || i == len(arr)-1 {
				list = append(list, arr[i])
			} else if arr[i] < arr[i-1] && arr[i] < arr[i+1] {
				changed = true
				list = append(list, arr[i]+1)
			} else if arr[i] > arr[i-1] && arr[i] > arr[i+1] {
				changed = true
				list = append(list, arr[i]-1)
			} else {
				list = append(list, arr[i])
			}
		}
		arr = list
	}
	return arr
}

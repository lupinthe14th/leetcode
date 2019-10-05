package intersectionofthreesortedarrays

func arraysIntersection(arr1 []int, arr2 []int, arr3 []int) []int {
	arrays := [][]int{arr1, arr2, arr3}
	sum := make(map[int]int, 0)

	for _, array := range arrays {
		for _, v := range array {
			sum[v]++
		}
	}

	stack := make([]int, 0)
	for k, v := range sum {
		if v == 3 {
			stack = append(stack, k)
		}
	}
	return stack
}

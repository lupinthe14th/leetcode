package findthedistancevaluebetweentwoarrays

func findTheDistanceValue(arr1 []int, arr2 []int, d int) int {
	var ans int
	for i := 0; i < len(arr1); i++ {
		t := false
		for j := 0; j < len(arr2) && !t; j++ {
			tmp := abs(arr1[i], arr2[j])
			if tmp <= d {
				t = true
			}
		}
		if !t {
			ans++
		}
	}
	return ans
}

func abs(x, y int) int {
	ans := x - y
	if ans < 0 {
		return ans * -1
	}
	return ans
}

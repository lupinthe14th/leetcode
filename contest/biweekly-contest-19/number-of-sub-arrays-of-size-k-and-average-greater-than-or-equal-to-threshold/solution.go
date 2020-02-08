package numberofsubarraysofsizekandaveragegreaterthanorequaltothreshold

func numOfSubarrays(arr []int, k int, threshold int) int {
	var c int
	for i := 0; i <= len(arr)-k; i++ {
		if sumA(arr[i:i+k])/k >= threshold {
			c++
		}
	}
	return c
}

func sumA(arr []int) int {
	var sum int
	for _, v := range arr {
		sum += v
	}
	return sum
}

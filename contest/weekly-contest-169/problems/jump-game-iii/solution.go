package main

func canReach(arr []int, start int) bool {
	// hashmap
	l := len(arr)
	hm := make(map[int]int, l)

	for i, v := range arr {
		hm[v] = i
	}

	zidx := hm[0] // value zero index number

	for {
		if zidx == i-arr[i] || zidx == i+arr[i] {

		} else {
			return false
		}
	}
	return true
}

func main() {}

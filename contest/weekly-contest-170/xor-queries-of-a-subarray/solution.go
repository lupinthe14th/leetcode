package main

func xorQueries(arr []int, queries [][]int) []int {
	ans := make([]int, 0, len(arr))
	for _, query := range queries {
		l, r := query[0], query[1]
		var x int
		if l == r {
			x = arr[l]
		} else {
			for i := l; i < r; i++ {
				if i == l {
					x = arr[i] ^ arr[i+1]
				} else {
					x ^= arr[i+1]
				}
			}
		}
		ans = append(ans, x)
	}
	return ans
}

func main() {}

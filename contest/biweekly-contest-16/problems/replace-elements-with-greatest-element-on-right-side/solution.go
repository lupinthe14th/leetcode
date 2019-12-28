package main

func solution(arr []int) []int {
	l := len(arr)
	ans := make([]int, 0, l)
	for i := 1; i <= l; i++ {
		if i == l {
			ans = append(ans, -1)
		} else {
			max := -1 << 31
			for j := i; j < l; j++ {
				if max < arr[j] {
					max = arr[j]
				}
			}
			ans = append(ans, max)
		}
	}
	return ans
}

func main() {}

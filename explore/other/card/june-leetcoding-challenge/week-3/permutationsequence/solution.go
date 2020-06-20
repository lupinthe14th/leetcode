package permutationsequence

import (
	"sort"
)

// Refarence: https://stackoverflow.com/questions/30226438/generate-all-permutations-in-go
func getPermutation(n int, k int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = byte(i) + '1'
	}
	var helper func([]byte, int)
	res := []string{}

	helper = func(b []byte, n int) {
		if n == 1 {
			tmp := make([]byte, len(b))
			copy(tmp, b)
			res = append(res, string(tmp))
		} else {
			for i := 0; i < n; i++ {
				helper(b, n-1)
				if n%2 == 1 {
					tmp := b[i]
					b[i] = b[n-1]
					b[n-1] = tmp
				} else {
					tmp := b[0]
					b[0] = b[n-1]
					b[n-1] = tmp
				}
			}
		}
	}
	helper(b, n)

	sort.Strings(res)
	return res[k-1]
}

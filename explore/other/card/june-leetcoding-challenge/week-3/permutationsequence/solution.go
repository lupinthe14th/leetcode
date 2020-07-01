package permutationsequence

// https://leetcode.com/problems/permutation-sequence/discuss/22544/Easy-understand-Most-concise-C++-solution-minimal-memory-required
func getPermutation(n int, k int) string {
	var i, j int
	f := 1

	b := make([]byte, n)

	for i = 1; i <= n; i++ {
		f *= i
		b[i-1] += byte(i) + '0'
	}
	k--
	for i := 0; i < n; i++ {
		f /= n - i
		j = i + k/f
		tmp := b[j]
		for ; j > i; j-- {
			b[j] = b[j-1]
		}
		k %= f
		b[i] = tmp
	}
	return string(b)
}

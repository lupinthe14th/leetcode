package thekweakestrowsinamatrix

import (
	"sort"
)

type matS struct {
	idx, sum int
}

func kWeakestRows(mat [][]int, k int) []int {
	l := len(mat)
	m := make([]matS, l)
	ans := make([]int, k)
	for i := 0; i < l; i++ {
		m[i] = matS{idx: i, sum: sumS(mat[i])}
	}
	sort.Slice(m, func(i, j int) bool {
		if m[i].sum == m[j].sum {
			return m[i].idx < m[j].idx
		}
		return m[i].sum < m[j].sum
	})

	for i := 0; i < k; i++ {
		ans[i] = m[i].idx
	}
	return ans

}

func sumS(s []int) int {
	var sum int
	for _, v := range s {
		sum += v

	}
	return sum
}

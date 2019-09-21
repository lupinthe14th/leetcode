package smallestcommonelement

import (
	"sort"
)

func smallestCommonElement(mat [][]int) int {
	// mapに文字種を格納して、同じ数があった場合はカウントアップする。
	// 格納したのち、小さい文字順にソートして、その順番に行列の数と同じカウントの数字を探す。
	// 同じカウントがあった場合はその数字を返し、なかった場合は-1を返す。
	type ml struct {
		number int
		count  int
	}
	m := make(map[int]ml) // counter
	mls := make([]ml, 0)

	for i := 0; i < len(mat); i++ {
		for _, v := range mat[i] {
			n, ok := m[v]
			if ok {
				m[v] = ml{number: n.number, count: n.count + 1}
			} else {
				m[v] = ml{number: v, count: 1}
			}
		}
	}

	for _, v := range m {
		mls = append(mls, v)
	}

	sort.Slice(mls, func(i, j int) bool {
		return mls[i].number < mls[j].number
	})

	for _, c := range mls {
		if c.count == len(mat) {
			return c.number
		}
	}
	return -1

}

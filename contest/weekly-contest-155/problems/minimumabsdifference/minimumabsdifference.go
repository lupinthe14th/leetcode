package minimumabsdifference

import (
	"log"
	"sort"
)

func minimumAbsDifference(arr []int) [][]int {
	sort.Ints(arr)
	log.Print(arr)

	type Counter struct {
		abs   int
		num   [][]int
		count int
	}
	m := make(map[int]Counter)
	counters := make([]Counter, 0)

	// 隣り合わせの数との差が一番小さいのを探す
	// 配列をループしてHashTableに単語があるか確認。確認した結果を構造体に反映
	for i := 0; i < len(arr)-1; i++ {
		abs := arr[i+1] - arr[i]
		n, ok := m[abs]
		if ok {
			m[abs] = Counter{abs: n.abs, num: append(n.num, []int{arr[i], arr[i+1]}), count: n.count + 1}
		} else {
			m[abs] = Counter{abs: abs, num: [][]int{[]int{arr[i], arr[i+1]}}, count: 1}
		}
	}
	// HashTableをループして、counters構造体の配列にセット
	for _, v := range m {
		counters = append(counters, v)
	}

	log.Printf("%#v", counters)

	// counters構造体の配列をソート。
	// absが小さい順に返す。
	sort.Slice(counters, func(i, j int) bool {
		return counters[i].abs < counters[j].abs
	})

	log.Print(counters[0].num)
	return counters[0].num
}

package main

// Algorithms: BFS (breadth-first search)
// See: https://ja.wikipedia.org/w/index.php?title=幅優先探索&oldid=72818204
func canReach(arr []int, start int) bool {
	l := len(arr)
	// 訪問済管理用
	v := make(map[int]bool)

	// 空のキューの作成
	q := make([]int, 0)
	// 1. 根のノードを空のキューに加える
	q = append(q, start)

	for len(q) != 0 {
		// 2. ノードをキューから取り出す
		x := q[len(q)-1]
		q = q[:len(q)-1]

		// ノードが探索対象であれば探索をやめ真を返して終了
		if arr[x] == 0 {
			return true
		}

		// そうで無い場合、探索済であれば、もう一度2.に戻る
		if v[x] {
			continue
		}
		// 未探索であれば、探索済とする。
		v[x] = true

		// 右方向の探索の値をキューに追加する。
		// 追加する場合はカレントポジションから右方向にジャンプする値が
		// arrスライス長を超えない場合
		right := x + arr[x]
		if right < l {
			q = append(q, right)
		}

		// 同様に左方向の探索の値をキューに追加する。
		// 追加する場合はカレントポジションから左方向にジャンプする値が
		// 0以上である場合
		left := x - arr[x]
		if left >= 0 {
			q = append(q, left)
		}
	}
	// 途中で探索対象とならずにキューが空になったならば、グラフ内のすべての
	// ノードに対して処理が終わったので探索をやめ偽を返して終了
	return false
}

func main() {}

package maximumnumberofeventsthatcanbeattended

import (
	"container/heap"
	"sort"
)

// Refarence: https://leetcode.com/problems/maximum-number-of-events-that-can-be-attended/discuss/510263/JavaC++Python-Priority-Queue
// 1. Sort events increased by start time.
//    開始時間毎にイベントを並べ替えます
// 2. Priority queue pq keeps the current open events.
//    優先キュー pqは 現在のオープンイベントを保持します
//
// 3. Iterate from the day 1 to day 100000, Each day, we add new events
//    starting on day d to the queue pq. Also we remove the events that are
//    already closed.
//    1日目から100000日目まで繰り返し、毎日、d日から始まる新しいイベントを
//    キューpqに追加します。また既に閉じられているイベントを削除します。
//
// 4. Then we greedily attend the event that ends soonest.
//    貪慾法で最も早く終わるイベントに出席します。
// 5. If we can attend a meeting, we increment the result res.
//    イベントに出席できる場合はカウントを増やします。
// SeeAlso: https://golang.org/pkg/container/heap/#example__priorityQueue
// Same idea, though people worry about the time complexity of iteration all days.

// An IntHeap is a min-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func maxEvents(events [][]int) int {
	pq := make(IntHeap, 0, 1e5)
	heap.Init(&pq)
	sort.Slice(events, func(i, j int) bool {
		if events[i][0] == events[j][0] {
			return events[i][1] < events[j][1]
		}
		return events[i][0] < events[j][0]
	})
	var i, c, d int
	n := len(events)
	for pq.Len() > 0 || i < n {
		if pq.Len() == 0 {
			d = events[i][0]
		}
		for i < n && events[i][0] <= d {
			heap.Push(&pq, events[i][1])
			i++
		}

		heap.Pop(&pq)
		c++
		d++
		for pq.Len() > 0 && pq[0] < d {
			heap.Pop(&pq)
		}
	}
	return c
}

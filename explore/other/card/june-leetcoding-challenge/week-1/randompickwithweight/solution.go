package randompickwithweight

import (
	"math/rand"
)

type Solution struct {
	cum []int
	hi  int
}

func Constructor(w []int) Solution {
	cum := make([]int, len(w))
	cur := 0
	for i := range w {
		cur += w[i]
		cum[i] = cur
	}
	return Solution{cum: cum, hi: cur}
}

func (this *Solution) PickIndex() int {
	rnd := rand.Intn(this.hi)
	lo, hi := 0, len(this.cum)

	for lo < hi {
		mi := (lo + hi) >> 1
		if this.cum[mi] > rnd {
			hi = mi - 1
		} else {
			lo = mi + 1
		}
	}
	return lo
}

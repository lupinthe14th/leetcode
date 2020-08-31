package largestcomponentsizebycommonfactor

// SeeAlso:
// - https://leetcode.com/problems/largest-component-size-by-common-factor/discuss/200959/Simplest-Solution-(Union-Find-only)-No-Prime-Calculation
// - https://leetcode.com/problems/largest-component-size-by-common-factor/discuss/820164/Go-Union-Find-solution
func largestComponentSize(A []int) int {
	l := len(A)
	memo := make(map[int]int)
	uf := NewUnionFind(l)

	for i := 0; i < l; i++ {
		a := A[i]
		for j := 2; j*j <= a; j++ {
			if a%j == 0 {
				if _, ok := memo[j]; !ok {
					memo[j] = i
				} else {
					uf.union(i, memo[j])
				}
				if _, ok := memo[a/j]; !ok {
					memo[a/j] = i
				} else {
					uf.union(i, memo[a/j])
				}
			}
		}
		if _, ok := memo[a]; !ok {
			memo[a] = i
		} else {
			uf.union(i, memo[a])
		}
	}
	return uf.max
}

type UnionFind struct {
	parent, size []int
	max          int
}

func NewUnionFind(n int) *UnionFind {
	parent := make([]int, n)
	size := make([]int, n)
	max := 1
	for i := 0; i < n; i++ {
		parent[i] = i
		size[i] = 1
	}
	return &UnionFind{
		parent: parent,
		size:   size,
		max:    max,
	}
}

func (uf *UnionFind) find(x int) int {
	if x != uf.parent[x] {
		uf.parent[x] = uf.find(uf.parent[x])
	}
	return uf.parent[x]
}

func (uf *UnionFind) union(x, y int) {
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	rootX := uf.find(x)
	rootY := uf.find(y)
	if rootX != rootY {
		uf.parent[rootX] = rootY
		uf.size[rootY] += uf.size[rootX]
		uf.max = max(uf.max, uf.size[rootY])
	}
}

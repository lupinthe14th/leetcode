package leftmostcolumnwithatleastaone

type BinaryMatrix interface {
	Get(int, int) int
	Dimensions() []int
}

type binaryMatrix struct {
	mat [][]int
	c   int
}

func NewBinaryMatrix(mat [][]int) *binaryMatrix {
	return &binaryMatrix{mat: mat}
}

func (bm *binaryMatrix) Get(x, y int) int {
	if bm.c >= 1000 {
		panic("You made too many calls to BinartMatrix.get().")
	}
	bm.c++
	return bm.mat[x][y]
}

func (bm *binaryMatrix) Dimensions() []int {
	m := len(bm.mat)
	n := len(bm.mat[0])
	return []int{m, n}
}

func leftMostColumnWithOne(binaryMatrix BinaryMatrix) int {
	d := binaryMatrix.Dimensions()
	ans := make([]int, d[0])

	for i := 0; i < d[0]; i++ {
		if binaryMatrix.Get(i, d[1]-1) == 0 {
			ans[i] = -1
			continue
		}
		if binaryMatrix.Get(i, 0) == 1 {
			ans[i] = 0
			break
		}
		lo, hi := 0, d[1]-1
		for lo < hi {
			mid := (lo + hi) / 2
			if binaryMatrix.Get(i, mid) == 1 {
				hi = mid
			} else {
				lo = mid + 1
			}
		}
		ans[i] = lo
	}
	a := 1 << 31
	for i := range ans {
		if ans[i] != -1 {
			a = min(a, ans[i])
		}
	}
	if a == 1<<31 {
		return -1
	}
	return a
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func bs(nums []int) int {
	l := len(nums)
	lo, hi := 0, l
	for lo < hi {
		mid := (lo + hi) / 2
		if nums[mid] == 1 {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	return lo
}

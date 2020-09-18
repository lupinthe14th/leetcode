package maximumxoroftwonumbersinanarray

func findMaximumXOR(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// Init trie
	type trieNode struct {
		children [2]*trieNode
	}

	root := &trieNode{}

	for _, num := range nums {
		node := root
		for i := 31; i >= 0; i-- {
			curBit := (num >> uint(i)) & 1
			if node.children[curBit] == nil {
				node.children[curBit] = &trieNode{}
			}
			node = node.children[curBit]
		}
	}

	out := -1 << 31

	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	for _, num := range nums {
		node := root
		curSum := 0
		for i := 31; i >= 0; i-- {
			curBit := (num >> uint(i)) & 1
			if node.children[curBit^1] != nil {
				curSum += 1 << uint(i)
				node = node.children[curBit^1]
			} else {
				node = node.children[curBit]
			}
		}
		out = max(out, curSum)
	}
	return out
}

package main

func hammingWeight(num uint32) int {
	var c int
	for num > 0 {
		if num&1 == 1 {
			c++
		}
		num >>= 1
	}
	return c
}

func main() {}

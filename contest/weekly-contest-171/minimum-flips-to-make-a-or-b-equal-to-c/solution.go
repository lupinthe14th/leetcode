package main

// SeeAlso: https://leetcode.com/problems/minimum-flips-to-make-a-or-b-equal-to-c/discuss/479691/Java-bit-manipulation-explained-beats-100
func minFlips(a int, b int, c int) int {
	var ans int
	// while three is a least one bit in any of the numbers
	for a > 0 || b > 0 || c > 0 {
		// if c has '1' check how many in a and b has '0'
		if c&1 == 1 {
			if a&1 == 0 && b&1 == 0 {
				ans++
			}
		} else {
			// for '0' in c check how many '1' are in a and b
			if a&1 == 1 {
				ans++
			}
			if b&1 == 1 {
				ans++
			}
		}
		// move to the next bit in all three numbers
		a >>= 1
		b >>= 1
		c >>= 1
	}
	return ans
}

func main() {}

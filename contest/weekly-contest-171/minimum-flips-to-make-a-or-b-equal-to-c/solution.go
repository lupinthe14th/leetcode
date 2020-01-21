package main

// We can collect number of flips in a greedy manner - check every bit in coreesponding index of every a, b and c. Then count those that are not matched.
// 貪慾法で反転数を収集できます - すべてのa,b,cの対応するインデックスのすべてのビットをチェックします。次に一致しないものをカウントします。
//
// There are 2 cases essentially:
// 基本的に2つのケースがあります。
// 1. c bit is '1' - in this case a OR b = '1' only in case both a or b are '1'. This means that if both are '0', a or b must be flip to '1'.
//    c ビットは「1」- この場合、aとbのどちらかが「1」の場合のみ、a OR b = 「1」です。これは、両方「0」の場合にa or bどちらからを「1」に反転する必要があることを意味します。
// 2. c bit is '0' - in this case a OR b = '0' only in case both a and b are '0'. This means we flip every '1'.
//    c ビットは「0」 - この場合、aとbの両方が「0」の場合のみ、a OR b = 「0」です。これは「1」ごとに反転することを意味します。
//
// We keep iterating over numbers and shit to the right on 1 position so next iteration we're checking next set of bits.
// 数値を繰り返し処理し、1ポジションで右シフトするので、次の反復では次のビットセットをチェックします。
//
// O(max(a,b,c)) time - need to iterate max number of bits times. O(1) - no extrace space used
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

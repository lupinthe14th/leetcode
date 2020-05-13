package removekdigits

// Straightforward Using stack
// SeeAlso: https://leetcode.com/problems/remove-k-digits/discuss/88708/Straightforward-Java-Solution-Using-Stack
func removeKdigits(num string, k int) string {
	l := len(num)
	// Corner case
	if l == k {
		return "0"
	}

	stack := make([]byte, 0)

	i := 0
	for i < l {
		// whenever meet a digit which is less than the previous digit, discard the previous one
		for k > 0 && len(stack) != 0 && stack[0] > num[i] {
			stack = stack[1:]
			k--
		}
		stack = append([]byte{num[i]}, stack...)
		i++
	}

	// corner case like "1111"
	for k > 0 {
		stack = stack[1:]
		k--
	}
	// construct the number from the stack
	b := make([]byte, 0)

	for len(stack) != 0 {
		b = append(b, stack[0])
		stack = stack[1:]
	}

	for i := len(b)/2 - 1; i >= 0; i-- {
		opp := len(b) - 1 - i
		b[i], b[opp] = b[opp], b[i]
	}

	// remove all the 0 at the head
	for len(b) > 1 && b[0] == '0' {
		b = b[1:]
	}

	return string(b)
}

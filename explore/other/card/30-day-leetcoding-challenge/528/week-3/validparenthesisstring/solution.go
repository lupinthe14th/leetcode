package validparenthesisstring

func checkValidString(s string) bool {
	// Let lo, hi respectively be the smallest and largest possible number
	// of open left brackets after processing the current character in the
	// string.
	lo, hi := 0, 0
	for _, r := range s {
		// If we encounter a left bracket (c == '('), then lo++,
		// otherwise we could write a right bracket, so lo--.
		if r == '(' {
			lo++
		} else {
			lo--
		}
		// If we encounter what can be a left bracket ( c != ')'), then
		// hi++, otherwise we must write a right bracket, so hi--.
		if r != ')' {
			hi++
		} else {
			hi--
		}
		// if hi < 0, then the current prefix can't be made valid no
		// matter what our choices are.
		if hi < 0 {
			break
		}
		// Also, we can never have less than 0 open left brackets.
		lo = max(lo, 0)
	}
	// we shouldcheck that we can have exactly 0 open left brackets.
	return lo == 0
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

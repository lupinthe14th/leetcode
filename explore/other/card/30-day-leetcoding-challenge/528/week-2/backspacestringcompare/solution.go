package backspacestringcompare

// Approach: Two Pointer
//
// Intuition
//
// When writing a character, it may or may not be part of the final string
// depending on how many backspace keystrokes occur in the future.
//
// If instead we iterate though the string in reverce, then we will know how
// many backspace characters we have seen, and there fore whether the result
// includes our character.
//
// Algorithm
//
// Iterate through the string in reverse. If wee see a backspace character, the
// next non-backspace character is skipped. If a character isn't skipped, it is
// part of the final answer.
//
// See the comments in the code for more details.
func backspaceCompare(S string, T string) bool {
	i, j := len(S)-1, len(T)-1
	skipS, skipT := 0, 0
	for i >= 0 || j >= 0 { // While there may be charas in build(S) or build(T)
		for i >= 0 { // Find position of next possible char in build(S)
			if string(S[i]) == "#" {
				skipS++
				i--
			} else if skipS > 0 {
				skipS--
				i--
			} else {
				break
			}
		}
		for j >= 0 { // Find position of next possible char in build(T)
			if string(T[j]) == "#" {
				skipT++
				j--
			} else if skipT > 0 {
				skipT--
				j--
			} else {
				break
			}
		}
		// If two actual characters are different
		if i >= 0 && j >= 0 && S[i] != T[j] {
			return false
		}
		// If expecting to compare char vs nothing
		if (i >= 0) != (j >= 0) {
			return false
		}
		i--
		j--
	}
	return true
}

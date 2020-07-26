package bulbswitcheriv

func minFlips(target string) int {
	out := 0
	f := false
	for i := 0; i < len(target); i++ {
		if !f && target[i] == '1' {
			f = true
			out++
		} else if f && target[i] == '0' {
			f = false
			out++
		}
	}
	return out
}

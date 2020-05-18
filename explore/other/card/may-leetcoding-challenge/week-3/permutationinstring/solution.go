package permutationinstring

func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}

	var mem1, mem2 [26]int

	for i := range s1 {
		mem1[s1[i]-'a']++
		mem2[s2[i]-'a']++
	}

	for i := 0; i < len(s2)-len(s1)+1; i++ {
		if mem1 == mem2 {
			return true
		}

		if i+len(s1) < len(s2) {
			mem2[s2[i]-'a']--
			mem2[s2[i+len(s1)]-'a']++
		}
	}
	return false
}

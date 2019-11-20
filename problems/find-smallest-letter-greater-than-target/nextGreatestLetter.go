package findsmallestlettergreaterthantarget

func nextGreatestLetter(letters []byte, target byte) byte {
	for _, t := range letters {
		if t > target {
			return t
		}
	}
	return letters[0]
}

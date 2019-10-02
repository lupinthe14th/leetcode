package implementstrstr

func strStr(haystack string, needle string) int {
	// return 0 when needle is an empty string.
	if needle == "" {
		return 0
	}

	// return -1 when haystack is an empty string or needle string length more than haystack string length.
	// because needle is not part of haystack.
	if haystack == "" || len(haystack) < len(needle) {
		return -1
	}

	for i := 0; i <= len(haystack)-len(needle); i++ {

		if haystack[i:i+len(needle)] == needle {
			return i
		}
	}

	return -1
}

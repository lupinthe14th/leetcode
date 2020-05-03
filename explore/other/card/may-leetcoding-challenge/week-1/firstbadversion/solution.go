package firstbadversion

func firstBadVersion(n int) int {
	lo, mi, hi := 1, 0, n

	for lo < hi {
		mi = (lo + hi) / 2
		if isBadVersion(mi) {
			hi = mi
		} else {
			lo = mi + 1
		}
	}
	return lo
}

func isBadVersion(version int) bool {
	if version <= 3 {
		return false
	}
	return true
}

package decodedstringatindex

func decodeAtIndex(S string, K int) string {
	size := 0
	N := len(S)

	for i := 0; i < N; i++ {
		if '0' <= S[i] && S[i] <= '9' {
			size *= int(S[i] - '0')
		} else {
			size++
		}
	}

	out := ""
	for i := N - 1; i >= 0; i-- {
		K %= size
		if K == 0 && ('a' <= S[i] && S[i] <= 'z') {
			out = string(S[i])
			break
		}

		if '0' <= S[i] && S[i] <= '9' {
			size /= int(S[i] - '0')
		} else {
			size--
		}
	}
	return out
}

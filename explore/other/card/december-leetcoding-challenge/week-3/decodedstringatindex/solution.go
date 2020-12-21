package decodedstringatindex

func decodeAtIndex(S string, K int) string {
	buf := make([]byte, 0, K-1)
	for i := 0; i < len(S); i++ {
		if '2' <= S[i] && S[i] <= '9' {
			tmp := make([]byte, len(buf))
			copy(tmp, buf)
			for j := 0; j < int(S[i]-'0')-1; j++ {
				buf = append(buf, tmp...)
			}
		} else {
			buf = append(buf, S[i])
		}
		if len(buf) >= K {
			break
		}
	}
	return string(buf[K-1])
}

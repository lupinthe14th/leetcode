package destinationcity

func destCity(paths [][]string) string {
	in := make(map[string]bool, len(paths))
	for i := range paths {
		in[paths[i][0]] = true
	}
	var ans string
	for i := range paths {
		if _, ok := in[paths[i][1]]; !ok {
			ans = paths[i][1]
			break
		}
	}
	return ans
}

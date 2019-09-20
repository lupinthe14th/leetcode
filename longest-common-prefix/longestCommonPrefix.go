package longestcommonprefix

import (
	"sort"
	"strings"
)

func longestCommonPrefix(strs []string) string {
	if len(strs) == 1 {
		return strs[0]
	}
	if len(strs) > 1 {
		// Get the string with the fewest characters from the string slice
		// solution: use sort
		// 文字列数と単語のマップを作成する。
		// キーは文字列とし、重複排除する。
		strsM := make(map[string]int)
		for _, str := range strs {
			strsM[str] = len(str)
		}

		// sort用のスライスを作る
		type StrsS struct {
			length int
			word   string
		}

		strsS := make([]StrsS, 0, len(strs))

		for k, v := range strsM {
			strsS = append(strsS, StrsS{length: v, word: k})
		}

		sort.Slice(strsS, func(i, j int) bool { return strsS[i].length < strsS[j].length })
		// その単語の文字をその単語以外全てと比較して全てにマッチするか判定する
		// 比較対象はマップを用いる
		checkStr := strsS[0].word
		c := 0
		for i := 0; len(checkStr)-i > 0; i++ {
			for k := range strsM {
				if k != checkStr {
					// 比較文字列はループ回数分後方から削る
					if strings.HasPrefix(k, checkStr[:len(checkStr)-i]) {
						c++
						if c == len(strs) {
							// 全てにマッチしたらその文字列を返す
							return checkStr[:len(checkStr)-i]
						}
					}
				}
			}
		}
	}
	// If there is no common prefix, return an empty string ""
	return ""
}

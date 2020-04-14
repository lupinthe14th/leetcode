package htmlentityparser

import "strings"

// SeeAlso: https://leetcode.com/problems/html-entity-parser/discuss/577224/Go-strings.NewReplacer-100-solution
func entityParser(text string) string {
	r := strings.NewReplacer("&quot;", "\"", "&apos;", "'", "&amp;", "&", "&gt;", ">", "&lt;", "<", "&frasl;", "/")
	return r.Replace(text)
}

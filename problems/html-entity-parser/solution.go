package htmlentityparser

import "strings"

func entityParser(text string) string {
	// Quotation Mark: the entity is &quot; and symbol character is ".
	text = strings.ReplaceAll(text, "&quot;", "\"")
	// Single Quote Mark: the entity is &apos; and symbol character is '.
	text = strings.ReplaceAll(text, "&apos;", "'")
	// Greater Than Sign: the entity is &gt; and symbol character is >.
	text = strings.ReplaceAll(text, "&gt;", ">")
	// Less Than Sign: the entity is &lt; and symbol character is <.
	text = strings.ReplaceAll(text, "&lt;", "<")
	// Slash: the entity is &frasl; and symbol character is /.
	text = strings.ReplaceAll(text, "&frasl;", "/")
	// Ampersand: the entity is &amp; and symbol character is &.
	text = strings.ReplaceAll(text, "&amp;", "&")

	return text
}

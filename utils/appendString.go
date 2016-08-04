package utils

func AppendStringHTMLNewLine(s string, n string) string {
	if s != "" {
		s += "<br />"
	}
	return s + n
}

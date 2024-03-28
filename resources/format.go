package resources

import (
	"strings"
	"unicode"
)

func GetName(name string) string {
	sb := strings.Builder{}
	sb.WriteString(strings.ToUpper(string(name[0])))
	for _, char := range name[1:] {
		if unicode.IsUpper(char) {
			sb.WriteString(" ")
		}
		sb.WriteString(string(char))
	}
	return sb.String()
}

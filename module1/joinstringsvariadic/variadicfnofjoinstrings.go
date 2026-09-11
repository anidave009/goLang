package joinstringsvariadic

import "strings"

func MyJoin(separator string, words ...string) string {
	var builder strings.Builder

	for i, word := range words {
		//Add separator before each word except the first one
		if i > 0 {
			builder.WriteString(separator)
		}
		builder.WriteString(word)
	}
	return builder.String()
}

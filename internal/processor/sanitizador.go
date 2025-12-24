package processor

import "strings"

func Sanitizador(text string) string {
	cleanText := strings.ToLower(strings.TrimSpace(text))

	for _, word := range Ignored {
		if strings.Contains(word, cleanText) {
			return ""
		}
	}

	return cleanText
}

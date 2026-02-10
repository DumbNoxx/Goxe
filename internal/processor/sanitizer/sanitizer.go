package sanitizer

import (
	"strings"
)

// This function cleans the text by removing spaces and ignoring the words inside the filters
func Sanitizer(text string, idLog string) string {
	var infoWord string

	text = strings.TrimSpace(text)
	text = reIpLogs.ReplaceAllString(text, "")

	if len(idLog) > 0 {
		infoWord = SafeWord.ReplaceAllString(text, "")
	} else {
		infoWord = text
	}

	textSanitize := strings.ToLower(infoWord)
	infoText := reDates.ReplaceAllString(textSanitize, "")
	cleanText := strings.TrimSpace(reStatus.ReplaceAllString(infoText, ""))

	return ExtractLevelUpper(textSanitize) + cleanText
}

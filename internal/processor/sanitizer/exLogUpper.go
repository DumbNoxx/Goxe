package sanitizer

import "strings"

func ExtractLevelUpper(log string) string {
	status := reStatus.FindString(log)

	if status == "" {
		return ""
	}
	return strings.ToUpper(status)
}

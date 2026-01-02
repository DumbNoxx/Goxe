package sanitizer

import (
	"regexp"
	"strings"

	"github.com/DumbNoxx/Goxe/internal/processor/filters"
)

var (
	reStatus = regexp.MustCompile(filters.PatternsLogLevel)
	reDates  = regexp.MustCompile(strings.Join(filters.PatternsDate, "|"))
	reIpLogs = regexp.MustCompile(filters.PatternIpLogs)
)

func SafeWord(word string) *regexp.Regexp {
	return regexp.MustCompile(regexp.QuoteMeta(word) + filters.PatternsIdLogs)
}

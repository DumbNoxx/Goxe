package cluster

import (
	"strings"

	"github.com/DumbNoxx/goxe/internal/processor/sanitizer"
)

func NormalizeLog(log string) string {
	return strings.TrimSpace(sanitizer.Re.ReplaceAllString(log, "*"))
}

package cluster

import "github.com/DumbNoxx/goxe/internal/processor/sanitizer"

func Cluster(log string, idLog string) string {
	text := sanitizer.Sanitizer(log, idLog)
	normalizeLog := NormalizeLog(text)

	return normalizeLog
}

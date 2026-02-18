package exporter

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/DumbNoxx/goxe/pkg/pipelines"
)

func FileReader(logs map[string]map[string]*pipelines.LogStats, path string) {

	date := time.Now().Format("2006-01-02")

	var (
		folderCachePath string
		data            strings.Builder
	)
	pathFile := filepath.Base(path)
	fileName := strings.TrimSuffix(pathFile, filepath.Ext(pathFile))

	file := fmt.Sprintf("%s_%s_normalized.log", fileName, date)
	dir := filepath.Dir(path)

	folderCachePath = filepath.Join(dir, file)

	fmt.Fprintln(&data, "\tRESULT")
	fmt.Fprintln(&data, "----------------------------------")

	for key, stat := range logs {
		fmt.Fprintf(&data, "ORIGIN: [%s]\n", key)
		for msg, stats := range stat {
			fmt.Fprintf(&data, "- [%d] %s -- (First seen %v - Last seen %v)\n", stats.Count, msg, stats.FirstSeen.Format("15:04:05"), stats.LastSeen.Format("15:04:05"))
		}
	}

	fmt.Fprintln(&data, "----------------------------------")

	err := os.WriteFile(folderCachePath, []byte(data.String()), 0600)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("[System] Goxe: Normalization complete. Saved as %s\n", file)

}

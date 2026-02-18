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

// FileReader generates a normalized report file from the processed logs,
// saving it in the same directory as the original file.
//
// Parameters:
//
//   - logs: log map (source -> message -> statistics) to be written into the report.
//   - path: path of the original file (used to derive the directory and base filename).
//
// The function performs:
//
//   - Obtains the current date in "2006-01-02" format.
//   - Extracts the base filename (without extension) from the provided path.
//   - Constructs the output filename as "<originalName>_<date>_normalized.log".
//   - Sets the output directory to be the same as the original file's directory.
//   - Builds the report content using a strings.Builder, including a title and separators.
//   - Iterates over the logs and writes each origin and its messages, including counters and timestamps (format: "15:04:05").
//   - Writes the content to the file with 0600 permissions.
//   - If the write operation fails, the program terminates with log.Fatal.
//   - Upon success, prints a console message indicating the path where the file was saved.
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

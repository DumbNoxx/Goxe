package processor

import (
	"bufio"
	"log"
	"os"
	"sync"
	"time"
	"unsafe"

	"github.com/DumbNoxx/goxe/internal/exporter"
	"github.com/DumbNoxx/goxe/internal/processor/cluster"
	"github.com/DumbNoxx/goxe/internal/processor/sanitizer"
	"github.com/DumbNoxx/goxe/pkg/pipelines"
)

func CleanFile(file *os.File, idLog string, mu *sync.Mutex, routeFile string) {
	var (
		sanitizadedText string
		data            []byte
		logsFile        = make(map[string]map[string]*pipelines.LogStats, 100)
	)

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		data = scanner.Bytes()
		dataCluster := cluster.Cluster(data, idLog)
		sanitizadedText = unsafe.String(unsafe.SliceData(dataCluster), len(dataCluster))

		mu.Lock()
		if logsFile["file-reader"] == nil {
			logsFile["file-reader"] = make(map[string]*pipelines.LogStats)
		}
		sliceData := sanitizer.ExtractLevelUpper(data)
		word := unsafe.String(unsafe.SliceData(sliceData), len(sliceData))
		if logsFile["file-reader"][sanitizadedText] == nil {
			logsFile["file-reader"][sanitizadedText] = &pipelines.LogStats{
				Count:     0,
				FirstSeen: time.Now(),
				LastSeen:  time.Now(),
				Level:     []byte(word),
			}
		}
		logsFile["file-reader"][sanitizadedText].Count++
		logsFile["file-reader"][sanitizadedText].LastSeen = time.Now()
		mu.Unlock()
	}
	err := exporter.ShipLogs(logsFile)
	if err != nil {
		log.Fatal(err)
	}
	exporter.FileReader(logsFile, routeFile)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	clear(logsFile)
}

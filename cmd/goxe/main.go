package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"sync"
	"time"

	"github.com/DumbNoxx/Goxe/internal/ingestor"
	"github.com/DumbNoxx/Goxe/internal/options"
	"github.com/DumbNoxx/Goxe/internal/processor"
	"github.com/DumbNoxx/Goxe/pkg/pipelines"
)

var (
	versionFlag *bool
	version     = ""
)

func init() {
	versionFlag = flag.Bool("v", false, "")

}

func viewConfig() {
	dir, _ := os.UserConfigDir()
	configPath := filepath.Join(dir, "goxe", "config.json")
	initialStat, err := os.Stat(configPath)
	if err != nil {
		log.Fatal(err)
	}
	lastModified := initialStat.ModTime()

	ticker := time.NewTicker(1 * time.Second)
	for range ticker.C {
		currentStat, err := os.Stat(configPath)
		if err != nil {
			log.Fatal(err)
		}
		if currentStat.ModTime().After(lastModified) {
			fmt.Println("Config update, reload...")
			lastModified = currentStat.ModTime()
			options.Config = options.ConfigFile()
		}
	}
}

func main() {
	flag.Parse()

	if *versionFlag {
		fmt.Println(version)
		os.Exit(0)
	}

	var wg sync.WaitGroup
	pipe := make(chan *pipelines.LogEntry, 100)
	var mu sync.Mutex

	options.CacheDirGenerate()

	wg.Add(1)
	go processor.Clean(pipe, &wg, &mu)
	go ingestor.Udp(pipe, &wg)
	go viewConfig()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	close(pipe)

	wg.Wait()

}

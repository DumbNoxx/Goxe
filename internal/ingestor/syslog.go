package ingestor

import (
	"fmt"
	"net"
	"strconv"
	"sync"
	"time"

	"github.com/DumbNoxx/Goxe/internal/options"
	"github.com/DumbNoxx/Goxe/pkg/pipelines"
)

var (
	PORT string = ":" + strconv.Itoa(options.Config.Port)
)

func Udp(pipe chan<- pipelines.LogEntry, wg *sync.WaitGroup) {

	addr, err := net.ResolveUDPAddr("udp", PORT)
	if err != nil {
		fmt.Println("Error resolving address:", err)
		return
	}
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		fmt.Println("Listening error:", err)
		return
	}
	defer conn.Close()

	fmt.Printf("Server listening on port %s\n", PORT)

	buffer := make([]byte, 1024)

	for {
		n, clientAddr, err := conn.ReadFromUDP(buffer)

		if err != nil {
			fmt.Println("Error reading", err)
			continue
		}

		message := string(buffer[:n])
		host, _, _ := net.SplitHostPort(clientAddr.String())

		dates := pipelines.LogEntry{
			Source:    host,
			Content:   message,
			Timestamp: time.Now(),
			IdLog:     options.Config.IdLog,
		}

		pipe <- dates
	}

}

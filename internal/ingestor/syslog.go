package ingestor

import (
	"context"
	"fmt"
	"net"
	"strconv"
	"sync"
	"time"
	"unsafe"

	"github.com/DumbNoxx/Goxe/internal/options"
	"github.com/DumbNoxx/Goxe/pkg/pipelines"
)

var (
	PORT      string = ":" + strconv.Itoa(options.Config.Port)
	lastIp    string
	lastRawIp net.IP
)

func Udp(ctx context.Context, pipe chan<- *pipelines.LogEntry, wg *sync.WaitGroup) {
	defer wg.Done()
	addr, err := net.ResolveUDPAddr("udp", PORT)
	if err != nil {
		fmt.Println("Error resolving address:", err)
		return
	}
	conn, err := net.ListenUDP("udp", addr)
	conn.SetReadBuffer(options.Config.BufferUdpSize * 1024 * 1024)

	if err != nil {
		fmt.Println("Listening error:", err)
		return
	}
	go func() {
		<-ctx.Done()
		conn.Close()
	}()

	fmt.Printf("Goxe listening on port %s\n", PORT)

	for {
		buffer := pipelines.BufferPool.Get().([]byte)
		n, clientAddr, err := conn.ReadFromUDP(buffer)

		if err != nil {
			if ctx.Err() != nil {
				return
			}
			fmt.Println("Error reading", err)
			pipelines.BufferPool.Put(buffer)
			continue
		}

		if !clientAddr.IP.Equal(lastRawIp) {
			lastRawIp = clientAddr.IP
			lastIp = clientAddr.IP.String()
		}

		dates := pipelines.EntryPool.Get().(*pipelines.LogEntry)
		dates.Content = unsafe.String(&buffer[0], n)
		dates.Source = lastIp
		dates.Timestamp = time.Now()
		dates.IdLog = options.Config.IdLog
		dates.RawEntry = buffer

		select {
		case pipe <- dates:
		case <-ctx.Done():
			return
		}

	}

}

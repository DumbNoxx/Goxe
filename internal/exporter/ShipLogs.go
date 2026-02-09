package exporter

import (
	"encoding/json"
	"net"
	"time"

	"github.com/DumbNoxx/goxe/internal/options"
	"github.com/DumbNoxx/goxe/pkg/pipelines"
)

type DataSentTcp struct {
	Origin string       `json:"origin"`
	Data   []TcpLogSent `json:"tcpLogSent"`
}

type TcpLogSent struct {
	Count     int       `json:"count"`
	FirstSeen time.Time `json:"firstSeen"`
	LastSeen  time.Time `json:"lastSeen"`
	Message   string    `json:"message"`
}

func ShipLogs(logs map[string]map[string]*pipelines.LogStats) (err error) {
	if options.Config.ShipperConfig.Address == "" {
		return nil
	}
	conn, err := net.DialTimeout(
		options.Config.ShipperConfig.Protocol,
		options.Config.ShipperConfig.Address,
		time.Duration(options.Config.ShipperConfig.FlushInterval)*time.Second,
	)

	if err != nil {
		return err
	}
	defer conn.Close()
	for key, messages := range logs {
		var DataSentTcps DataSentTcp
		DataSentTcps.Origin = key
		DataSentTcps.Data = make([]TcpLogSent, 0, len(messages))
		for msg, stats := range messages {
			var logEntry = TcpLogSent{
				Count:     stats.Count,
				FirstSeen: stats.FirstSeen,
				LastSeen:  stats.LastSeen,
				Message:   msg,
			}

			DataSentTcps.Data = append(DataSentTcps.Data, logEntry)
		}
		data, err := json.Marshal(DataSentTcps)
		if err != nil {
			return err
		}
		_, err = conn.Write(data)
		if err != nil {
			return err
		}
	}
	return nil
}

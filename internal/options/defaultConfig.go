package options

import (
	"os"

	"github.com/DumbNoxx/goxe/pkg/options"
)

func configDefault() options.Config {
	home, _ := os.Hostname()

	return options.Config{
		Port:          1729,
		IdLog:         home,
		PatternsWords: []string{},
		GenerateLogsOptions: options.GenerateLogsOptions{
			GenerateLogsFile: false,
			Hour:             "00:00:00",
		},
		WebHookUrls: []string{},
		BurstDetectionOptions: options.BurstDetectionOptions{
			LimitBreak: 10,
		},
		ShipperConfig: options.ShipperConfig{
			Address:       "",
			FlushInterval: 30,
			Protocol:      "tcp",
		},
		ReportInterval: 60,
		BufferUdpSize:  4,
	}
}

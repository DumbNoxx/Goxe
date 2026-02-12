package reporttime

import (
	"time"

	"github.com/DumbNoxx/goxe/internal/options"
	"github.com/DumbNoxx/goxe/internal/processor"
	"github.com/DumbNoxx/goxe/internal/utils"
)

func GetReportFileTime() {
	processor.TickerReportFile.Stop()
	processor.TickerReportFile.Reset(utils.TimeReportFile)
}

func GetReportPartialTime() {
	processor.TimeReport = time.Duration(options.Config.ReportInterval * float64(time.Minute))
	processor.Ticker.Stop()
	processor.Ticker.Reset(processor.TimeReport)
}

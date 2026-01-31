package pipelines

import "time"

type LogBurst struct {
	Category      string
	WindowStart   time.Time
	Count         int
	AlertsSent    int
	LastAlertTime time.Time
}

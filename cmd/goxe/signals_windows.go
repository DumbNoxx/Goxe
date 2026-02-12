//go:build windows

package main

import (
	"os"
)

var watchSignals = []os.Signal{os.Interrupt}

func isUpdateSignal(_ os.Signal) bool {
	return false
}

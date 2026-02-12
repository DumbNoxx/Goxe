//go:build linux || darwin

package main

import (
	"os"
	"syscall"
)

// List of signals we want to listen to on Unix systems
var watchSignals = []os.Signal{os.Interrupt, syscall.SIGUSR1}

func isUpdateSignal(sig os.Signal) bool {
	return sig == syscall.SIGUSR1
}

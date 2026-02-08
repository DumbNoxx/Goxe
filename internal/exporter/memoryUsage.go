package exporter

import (
	"fmt"
	"runtime"
)

func memoryUsage() {
	var memory runtime.MemStats
	runtime.ReadMemStats(&memory)

	fmt.Printf("Memory used: %.2f Mb\n", float64(memory.Alloc)/1024/1024)
	fmt.Printf("Total program memory: %.2f Mb\n", float64(memory.Sys)/1024/1024)
	fmt.Printf("Heap in use: %.2f Mb\n", float64(memory.HeapInuse)/1024/1024)
	fmt.Println("")
}

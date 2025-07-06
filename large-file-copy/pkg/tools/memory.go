package tools

import (
	"fmt"
	"runtime"
)

// Helper function to print memory stats
func PrintMemStats(memo string) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// MemStats provides information about the Go runtime's memory allocation.
	// Check the documentation for more details: https://pkg.go.dev/runtime#MemStats
	fmt.Printf("%s: Alloc = %v MiB, TotalAlloc = %v MiB, Sys = %v MiB, NumGC = %v\n",
		memo, bToMb(m.Alloc), bToMb(m.TotalAlloc), bToMb(m.Sys), m.NumGC)
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}

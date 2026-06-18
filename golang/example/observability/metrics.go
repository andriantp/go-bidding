package main

import (
	"log"
	"sync/atomic"
)

var processedEvents uint64
var replayEvents uint64

func PrintMetrics() {
	log.Printf(
		"[metrics] processed_events=%d replay_events=%d",
		atomic.LoadUint64(&processedEvents),
		atomic.LoadUint64(&replayEvents),
	)
}

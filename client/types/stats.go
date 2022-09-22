package types

import "time"

type StatsType struct {
	Memory    MemoryStats
	CPU       CpuStats
	TimeStamp time.Time
	Error     error
}

type MemoryStats struct {
	Total  uint64
	Used   float32
	Cached float32
	Free   float32
}

type CpuStats struct {
	Usage float64
}

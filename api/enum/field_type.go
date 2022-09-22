package enum

type FieldType uint

const (
	FieldCPUUsage FieldType = iota
	FieldMemoryUsed
	FieldMemoryCached
	FieldMemoryFree
	FieldMemoryTotal
)

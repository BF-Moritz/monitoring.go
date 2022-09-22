package types

import (
	"time"

	"github.com/BF-Moritz/monitoring.go/api/enum"
)

type StatusType struct {
	Client       string
	CPUUsage     []FloatDataPoint
	MemoryTotal  []UintDataPoint
	MemoryUsed   []FloatDataPoint
	MemoryCached []FloatDataPoint
	MemoryFree   []FloatDataPoint
}

type UnparsedStatusType struct {
	Client       string
	CPUUsage     []UnparsedDataPoint
	MemoryTotal  []UnparsedDataPoint
	MemoryUsed   []UnparsedDataPoint
	MemoryCached []UnparsedDataPoint
	MemoryFree   []UnparsedDataPoint
}

func (status UnparsedDataPoint) ParseToUintDataPoint() (dataPoint UintDataPoint, err error) {
	var value uint64
	var timeStamp time.Time

	switch val := status.Value.(type) {
	case uint64:
		value = val
	case uint32:
		value = uint64(val)
	case uint16:
		value = uint64(val)
	case uint8:
		value = uint64(val)
	case uint:
		value = uint64(val)
	case int64:
		value = uint64(val)
	case int32:
		value = uint64(val)
	case int16:
		value = uint64(val)
	case int8:
		value = uint64(val)
	case int:
		value = uint64(val)
	}

	switch t := status.Time.(type) {
	case time.Time:
		timeStamp = t
	case string:
		val, err := time.Parse("2006-01-02T15:04:05.999999999Z07:00", t)
		if err != nil {
			return UintDataPoint{}, err
		}
		timeStamp = val
	}

	return UintDataPoint{
		Value: value,
		Time:  timeStamp,
	}, nil
}

func (status UnparsedDataPoint) ParseToFloatDataPoint() (dataPoint FloatDataPoint, err error) {
	var value float64
	var timeStamp time.Time

	switch val := status.Value.(type) {
	case float64:
		value = val
	case float32:
		value = float64(val)
	}

	switch t := status.Time.(type) {
	case time.Time:
		timeStamp = t
	case string:
		val, err := time.Parse("2006-01-02T15:04:05.999999999Z07:00", t)
		if err != nil {
			return FloatDataPoint{}, err
		}
		timeStamp = val
	}

	return FloatDataPoint{
		Value: value,
		Time:  timeStamp,
	}, nil
}

type UnparsedDataPoint struct {
	Value interface{}
	Time  interface{}
}

type UintDataPoint struct {
	Value uint64
	Time  time.Time
}

type FloatDataPoint struct {
	Value float64
	Time  time.Time
}

type ParseInType struct {
	Client string
	Field  enum.FieldType
	Type   enum.ParseType
	Values []UnparsedDataPoint
}

type ParseOutType struct {
	Client string
	Field  enum.FieldType
	Uints  *[]UintDataPoint
	Floats *[]FloatDataPoint
	Err    error
}

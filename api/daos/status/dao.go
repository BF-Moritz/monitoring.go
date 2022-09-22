package status

import (
	"context"
	"fmt"

	"github.com/BF-Moritz/monitoring.go/api/types"
	"github.com/BF-Moritz/monitoring.go/api/vars"
)

type DAO struct {
}

type DAOInterface interface {
	GetAll(lastMinutes uint32) (status []types.UnparsedStatusType, err error)
}

func NewDAO() DAOInterface {
	return &DAO{}
}

func (d *DAO) GetAll(lastMinutes uint32) (status []types.UnparsedStatusType, err error) {
	status = make([]types.UnparsedStatusType, 0, 5)

	result, err := vars.QueryAPI.Query(context.Background(), fmt.Sprintf(getAllSQL, lastMinutes))
	if err != nil {
		return nil, err
	}

	statusMap := make(map[string]types.UnparsedStatusType)
	var client string
	var field string

	for result.Next() {
		values := result.Record().Values()
		client = fmt.Sprintf("%v", values["client"])
		field = fmt.Sprintf("%v-%v", values["_measurement"], values["_field"])
		if _, ok := statusMap[client]; !ok {
			statusMap[client] = types.UnparsedStatusType{
				CPUUsage:     make([]types.UnparsedDataPoint, 0, 3600),
				MemoryTotal:  make([]types.UnparsedDataPoint, 0, 3600),
				MemoryUsed:   make([]types.UnparsedDataPoint, 0, 3600),
				MemoryCached: make([]types.UnparsedDataPoint, 0, 3600),
				MemoryFree:   make([]types.UnparsedDataPoint, 0, 3600),
				Client:       client,
			}
		}

		switch field {
		case "cpu-usage":
			v := statusMap[client]
			v.CPUUsage = append(v.CPUUsage, types.UnparsedDataPoint{
				Value: values["_value"],
				Time:  values["_time"],
			})
			statusMap[client] = v
		case "memory-used":
			v := statusMap[client]
			v.MemoryUsed = append(v.MemoryUsed, types.UnparsedDataPoint{
				Value: values["_value"],
				Time:  values["_time"],
			})
			statusMap[client] = v
		case "memory-cached":
			v := statusMap[client]
			v.MemoryCached = append(v.MemoryCached, types.UnparsedDataPoint{
				Value: values["_value"],
				Time:  values["_time"],
			})
			statusMap[client] = v
		case "memory-free":
			v := statusMap[client]
			v.MemoryFree = append(v.MemoryFree, types.UnparsedDataPoint{
				Value: values["_value"],
				Time:  values["_time"],
			})
			statusMap[client] = v
		case "memory-total":
			v := statusMap[client]
			v.MemoryTotal = append(v.MemoryTotal, types.UnparsedDataPoint{
				Value: values["_value"],
				Time:  values["_time"],
			})
			statusMap[client] = v
		}
	}

	for _, s := range statusMap {
		status = append(status, s)
	}

	return status, nil
}

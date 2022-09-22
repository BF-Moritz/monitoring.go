package monitoring

import (
	"time"

	api "github.com/BF-Moritz/monitoring.go/grpc_api"
	"github.com/BF-Moritz/monitoring.go/server/vars"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	influxdb2api "github.com/influxdata/influxdb-client-go/v2/api"
	"golang.org/x/net/context"
)

type Service struct {
	api.UnimplementedChatServiceServer
	InfluxWriteAPI influxdb2api.WriteAPIBlocking
}

func (s *Service) SendStatus(ctx context.Context, message *api.Ping) (*api.Pong, error) {
	timeStamp, err := time.Parse("2006-01-02T15:04:05.999999999Z07:00", message.TimeStamp)
	if err != nil {
		vars.Logger.LogError("service:monitoring.SendStatus()", "failed to parse timestamp (%s)", err)
		return nil, err
	}

	cpuPoint := influxdb2.NewPoint("cpu",
		map[string]string{"client": message.Client},
		map[string]interface{}{"usage": message.Body.CPU.Usage},
		timeStamp,
	)

	memPoint := influxdb2.NewPoint("memory",
		map[string]string{"client": message.Client},
		map[string]interface{}{
			"total":  message.Body.Memory.Total,
			"free":   message.Body.Memory.Free,
			"cached": message.Body.Memory.Cached,
			"used":   message.Body.Memory.Used,
		},
		timeStamp,
	)

	err = s.InfluxWriteAPI.WritePoint(context.Background(), cpuPoint, memPoint)
	if err != nil {
		vars.Logger.LogError("service:monitoring.SendStatus()", "failed to insert points (%s)", err)
		return nil, err
	}

	return &api.Pong{Id: message.Id}, nil
}

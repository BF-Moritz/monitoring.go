package stats

import (
	"context"
	"time"

	"github.com/BF-Moritz/monitoring.go/client/vars"
	api "github.com/BF-Moritz/monitoring.go/grpc_api"
)

func (m *StatsModule) sender() {
	for {
		time.Sleep(100 * time.Millisecond)
		m.mutex.RLock()
		if len(m.queue) > 0 {
			item := m.queue[0]

			message := api.Ping{
				Id:     uint64(time.Now().UnixMilli()),
				Client: m.clientName,
				Body: &api.StatsType{
					Memory: &api.MemoryStats{
						Total:  item.Memory.Total,
						Used:   item.Memory.Used,
						Cached: item.Memory.Cached,
						Free:   item.Memory.Free,
					},
					CPU: &api.CpuStats{
						Usage: float32(item.CPU.Usage),
					},
				},
				TimeStamp: item.TimeStamp.UTC().Format("2006-01-02T15:04:05.999999999Z07:00"),
			}

			_, err := m.client.SendStatus(context.Background(), &message)
			if err != nil {
				vars.Logger.LogError("service:stats.sender()", "failed to send stats, waiting 0.5s %s", err)
				time.Sleep(400 * time.Millisecond)
			} else {
				m.queue = m.queue[1:]
			}
		}
		m.mutex.RUnlock()
	}
}

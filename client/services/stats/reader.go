package stats

import (
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"

	"github.com/BF-Moritz/monitoring.go/client/types"
	"github.com/BF-Moritz/monitoring.go/client/vars"
)

func (m *StatsModule) reader() {
	for {
		time.Sleep(time.Duration(vars.Config.TimeBetweenReadings) * time.Second)

		cpuUsage, err := cpu.Percent(0, false)
		if err != nil {
			vars.Logger.LogError("service:stats.reader()", "failed to read CPU usage (%s)", err)
			m.channel <- types.StatsType{
				Error: err,
			}
			continue
		}

		cpuStats := types.CpuStats{
			Usage: cpuUsage[0],
		}

		v, err := mem.VirtualMemory()
		if err != nil {
			vars.Logger.LogError("service:stats.reader()", "failed to read Memory stats (%s)", err)
			m.channel <- types.StatsType{
				Error: err,
			}
			continue
		}

		memStats := types.MemoryStats{
			Total:  v.Total,
			Used:   float32(float64(v.Used) / float64(v.Total) * 100.0),
			Cached: float32(float64(v.Cached) / float64(v.Total) * 100.0),
			Free:   float32(float64(v.Free) / float64(v.Total) * 100.0),
		}

		m.channel <- types.StatsType{
			Memory:    memStats,
			CPU:       cpuStats,
			Error:     nil,
			TimeStamp: time.Now(),
		}

	}
}

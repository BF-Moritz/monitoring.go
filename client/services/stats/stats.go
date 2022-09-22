package stats

import (
	"sync"

	"github.com/BF-Moritz/monitoring.go/client/types"
	api "github.com/BF-Moritz/monitoring.go/grpc_api"
	"google.golang.org/grpc"
)

type StatsModule struct {
	queue      []types.StatsType
	mutex      sync.RWMutex
	channel    chan types.StatsType
	client     api.ChatServiceClient
	clientName string
}

func NewStatsModule(conn *grpc.ClientConn, clientName string) *StatsModule {
	return &StatsModule{
		queue:      make([]types.StatsType, 0, 100),
		mutex:      sync.RWMutex{},
		channel:    make(chan types.StatsType),
		client:     api.NewChatServiceClient(conn),
		clientName: clientName,
	}
}

func (m *StatsModule) Init() {
	go m.reader()

	go m.sender()

	for result := range m.channel {
		m.mutex.Lock()
		m.queue = append(m.queue, result)
		m.mutex.Unlock()
	}
}

func (m *StatsModule) Send(channel chan types.StatsType, mutex *sync.RWMutex) {
	go m.sender()

	for result := range channel {
		mutex.Lock()
		m.queue = append(m.queue, result)
		mutex.Unlock()
	}
}

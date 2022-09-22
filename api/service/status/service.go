package status

import (
	"github.com/BF-Moritz/monitoring.go/api/daos/status"
	"github.com/BF-Moritz/monitoring.go/api/types"
)

type Service struct {
	StatusDAO status.DAOInterface
}

type ServiceInterface interface {
	GetAll(lastMinutes uint32) (status []types.StatusType, err error)
	GetByName(name string) (status []types.StatusType, err error)
}

func NewService() ServiceInterface {

	return &Service{
		StatusDAO: status.NewDAO(),
	}
}

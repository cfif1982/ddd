package crashes

import (
	"ddd/internal/domain/clients"
	"ddd/internal/domain/contracts"
	"ddd/internal/domain/crashes"

	"github.com/google/uuid"
)

type crashRepo interface {
	GetCrash(id uuid.UUID) (*crashes.Crash, error)
	SaveCrash(c crashes.Crash) error // почему тут передаем сам объект , а не ссылку на него?
}

type contractRepo interface {
	GetContract(id uuid.UUID) (*contracts.Contract, error)
}

type clientRepo interface {
	GetClient(id uuid.UUID) (*clients.Client, error)
}

type CrashService struct {
	crashRepo    crashRepo
	contractRepo contractRepo
	clientRepo   clientRepo
}

func NewCrashService(crashRepo crashRepo, contractRepo contractRepo, clientRepo clientRepo) *CrashService {
	return &CrashService{
		crashRepo:    crashRepo,
		contractRepo: contractRepo,
		clientRepo:   clientRepo,
	}
}

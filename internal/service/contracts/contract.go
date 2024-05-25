package contracts

import (
	"ddd/internal/domain/clients"
	"ddd/internal/domain/contracts"
	"ddd/internal/domain/managers"

	"github.com/google/uuid"
)

// тут дублируются интерфейсы из Handlers - handler.go. Можно ли это в одном месте описать? Можно ли это по-другому реализовать?
type contractRepo interface {
	SaveContract(c contracts.Contract) error // почему тут передаем сам объект , а не ссылку на него?
	GetContract(id uuid.UUID) (*contracts.Contract, error)
}

type managerRepo interface {
	GetManager(id uuid.UUID) (*managers.Manager, error)
}

type clientRepo interface {
	GetClient(id uuid.UUID) (*clients.Client, error)
}

type ContractService struct {
	contractRepo contractRepo
	managerRepo  managerRepo
	clientRepo   clientRepo
}

func NewContractService(
	contractRepo contractRepo,
	managerRepo managerRepo,
	clientRepo clientRepo,
) *ContractService {
	return &ContractService{
		contractRepo: contractRepo,
		managerRepo:  managerRepo,
		clientRepo:   clientRepo,
	}
}

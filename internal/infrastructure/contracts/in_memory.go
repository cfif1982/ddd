package contracts

import (
	"ddd/internal/domain/contracts"
	"errors"

	"github.com/google/uuid"
)

// нужно ли здесь создавать свою структуру для БД и маппер в эту структуру из структуры domain/contracts?
type contractDB struct {
	managerID uuid.UUID
	clientID  uuid.UUID
	carID     uuid.UUID
	summa     int
}

type InMemoryRepo struct {
	db map[uuid.UUID]contractDB
}

func NewInMemoryRepo() *InMemoryRepo {
	return &InMemoryRepo{
		db: make(map[uuid.UUID]contractDB),
	}
}

// сюда передавать в параметрах структуру или указатель на нее?
func (r *InMemoryRepo) SaveContract(c contracts.Contract) error {

	// проверяем - есть ли уже записm в БД с таким key
	_, errFound := r.db[c.ID()]
	if errFound == true {
		return errors.New("id already exist")
	}

	r.db[c.ID()] = contractDB{
		managerID: c.ManagerID(),
		clientID:  c.ClientID(),
		summa:     c.Summa(),
	}

	return nil
}

func (r *InMemoryRepo) GetContract(id uuid.UUID) (*contracts.Contract, error) {

	cDB, errFound := r.db[id] // здесь ошибка возвращается в виде bool

	if errFound == false {
		return nil, errors.New("contract not found")
	}

	contract, err := contracts.NewContract(
		id,
		cDB.managerID,
		cDB.clientID,
		cDB.summa,
	)

	if err != nil {
		return nil, err
	}

	return contract, nil
}

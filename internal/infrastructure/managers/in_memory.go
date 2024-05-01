package managers

import (
	"ddd/internal/domain/managers"
	"errors"

	"github.com/google/uuid"
)

// нужно ли здесь создавать свою структуру для БД и маппер в эту структуру из структуры domain/contracts?
type managerDB struct {
	name  string
	email string
	bonus int
}

type InMemoryRepo struct {
	db map[uuid.UUID]managerDB
}

func NewInMemoryRepo() *InMemoryRepo {
	return &InMemoryRepo{
		db: make(map[uuid.UUID]managerDB),
	}
}

// сюда передавать в параметрах структуру или указатель на нее?
func (r *InMemoryRepo) SaveManager(m managers.Manager) error {

	// проверяем - есть ли уже записm в БД с таким key
	_, errFound := r.db[m.ID()]
	if errFound == true {
		return errors.New("id already exist")
	}

	r.db[m.ID()] = managerDB{
		name:  m.Name(),
		email: m.Email(),
		bonus: m.Bonus(),
	}

	return nil
}

func (r *InMemoryRepo) GetManager(id uuid.UUID) (*managers.Manager, error) {

	mDB, errFound := r.db[id] // здесь ошибка возвращается в виде bool

	if errFound == false {
		return nil, errors.New("manager not found")
	}

	manager, err := managers.NewManager(
		id,
		mDB.name,
		mDB.email,
		mDB.bonus,
	)

	if err != nil {
		return nil, err
	}

	return manager, nil
}

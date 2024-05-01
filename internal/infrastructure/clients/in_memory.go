package clients

import (
	"ddd/internal/domain/clients"
	"errors"

	"github.com/google/uuid"
)

// нужно ли здесь создавать свою структуру для БД и маппер в эту структуру из структуры domain/contracts?
type clientDB struct {
	name     string
	email    string
	discount int
}

type InMemoryRepo struct {
	db map[uuid.UUID]clientDB
}

func NewInMemoryRepo() *InMemoryRepo {
	return &InMemoryRepo{
		db: make(map[uuid.UUID]clientDB),
	}
}

// сюда передавать в параметрах структуру или указатель на нее?
func (r *InMemoryRepo) SaveClient(c clients.Client) error {

	// проверяем - есть ли уже записm в БД с таким key
	_, errFound := r.db[c.ID()]
	if errFound == true {
		return errors.New("id already exist")
	}

	r.db[c.ID()] = clientDB{
		name:     c.Name(),
		email:    c.Email(),
		discount: c.Discount(),
	}

	return nil
}

func (r *InMemoryRepo) GetClient(id uuid.UUID) (*clients.Client, error) {

	mDB, errFound := r.db[id] // здесь ошибка возвращается в виде bool

	if errFound == false {
		return nil, errors.New("client not found")
	}

	client, err := clients.NewClient(
		id,
		mDB.name,
		mDB.email,
		mDB.discount,
	)

	if err != nil {
		return nil, err
	}

	return client, nil
}

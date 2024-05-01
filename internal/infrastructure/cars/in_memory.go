package cars

import (
	"ddd/internal/domain/cars"
	"errors"

	"github.com/google/uuid"
)

// нужно ли здесь создавать свою структуру для БД и маппер в эту структуру из структуры domain/contracts?
type carDB struct {
	model string
	color string
	vin   string
}

type InMemoryRepo struct {
	db map[uuid.UUID]carDB
}

func NewInMemoryRepo() *InMemoryRepo {
	return &InMemoryRepo{
		db: make(map[uuid.UUID]carDB),
	}
}

// сюда передавать в параметрах структуру или указатель на нее?
func (r *InMemoryRepo) SaveCar(c cars.Car) error {

	// проверяем - есть ли уже записm в БД с таким key
	_, errFound := r.db[c.ID()]
	if errFound == true {
		return errors.New("id already exist")
	}

	r.db[c.ID()] = carDB{
		model: c.Model(),
		color: c.Color(),
		vin:   c.VIN(),
	}

	return nil
}

func (r *InMemoryRepo) GetCar(id uuid.UUID) (*cars.Car, error) {

	mDB, errFound := r.db[id] // здесь ошибка возвращается в виде bool

	if errFound == false {
		return nil, errors.New("car not found")
	}

	car, err := cars.NewCar(
		id,
		mDB.model,
		mDB.color,
		mDB.vin,
	)

	if err != nil {
		return nil, err
	}

	return car, nil
}

package contracts

import "github.com/google/uuid"

type Car struct {
	id    uuid.UUID
	model string
	vin   string
}

func NewCar(
	id uuid.UUID,
	model string,
	vin string,
) (*Car, error) {

	// проверка полей

	return &Car{
		id:    id,
		model: model,
		vin:   vin,
	}, nil
}

func (c *Car) ID() uuid.UUID {
	return c.id
}

func (c *Car) Model() string {
	return c.model
}

func (c *Car) VIN() string {
	return c.vin
}

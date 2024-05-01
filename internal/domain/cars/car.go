package cars

import "github.com/google/uuid"

type Car struct {
	id    uuid.UUID
	model string
	color string
	vin   string
}

func NewCar(id uuid.UUID, model, color, vin string) (*Car, error) {

	// проверка полей

	return &Car{
		id:    id,
		model: model,
		color: color,
		vin:   vin,
	}, nil
}

func CreateCar(model, color, vin string) (*Car, error) {

	return NewCar(uuid.New(), model, color, vin)
}

func (c *Car) ID() uuid.UUID {
	return c.id
}

func (c *Car) Model() string {
	return c.model
}

func (c *Car) Color() string {
	return c.color
}

func (c *Car) VIN() string {
	return c.vin
}

// функции проверки полей

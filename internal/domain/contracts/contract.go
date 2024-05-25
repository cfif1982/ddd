package contracts

import "github.com/google/uuid"

type Contract struct {
	id        uuid.UUID
	managerID uuid.UUID
	clientID  uuid.UUID
	summa     int
	cars      []Car
}

func NewContract(
	id uuid.UUID,
	managerID uuid.UUID,
	clientID uuid.UUID,
	summa int,
	cars []Car,
) (*Contract, error) {

	// проверка полей

	return &Contract{
		id:        id,
		managerID: managerID,
		clientID:  clientID,
		summa:     summa,
		cars:      cars,
	}, nil
}

func CreateContract(
	managerID uuid.UUID,
	clientID uuid.UUID,
	summa int,
	cars []Car,
) (*Contract, error) {

	return NewContract(
		uuid.New(),
		managerID,
		clientID,
		summa,
		cars,
	)
}

func (c *Contract) MakeTextForEmail() string {

	// Формируем текст договора
	text := "contract text"

	return text
}

func (c *Contract) ID() uuid.UUID {
	return c.id
}

func (c *Contract) ClientID() uuid.UUID {
	return c.clientID
}

func (c *Contract) ManagerID() uuid.UUID {
	return c.managerID
}

func (c *Contract) Summa() int {
	return c.summa
}

func (c *Contract) Cars() []Car {
	return c.cars
}

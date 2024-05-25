package contracts

import "github.com/google/uuid"

type Contract struct {
	id        uuid.UUID
	managerID uuid.UUID
	clientID  uuid.UUID
	summa     int
}

func NewContract(
	id uuid.UUID,
	managerID uuid.UUID,
	clientID uuid.UUID,
	summa int,
) (*Contract, error) {

	// проверка полей

	return &Contract{
		id:        id,
		managerID: managerID,
		clientID:  clientID,
		summa:     summa,
	}, nil
}

func CreateContract(
	managerID uuid.UUID,
	clientID uuid.UUID,
	summa int,
) (*Contract, error) {

	return NewContract(
		uuid.New(),
		managerID,
		clientID,
		summa,
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

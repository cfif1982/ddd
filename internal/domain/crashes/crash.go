package crashes

import (
	"github.com/google/uuid"
)

type Crash struct {
	id          uuid.UUID
	date        string
	contractID  uuid.UUID
	description string
	// нужен ли здесь clientID и carID. Мы можем их получить из contractID
}

func NewCrash(id uuid.UUID, date string, contractID uuid.UUID, description string) (*Crash, error) {

	// проверк полей

	return &Crash{
		id:          id,
		date:        date,
		contractID:  contractID,
		description: description,
	}, nil
}

func CreateCrash(contractID uuid.UUID, description string) (*Crash, error) {

	date := "current date"

	return NewCrash(uuid.New(), date, contractID, description)
}

func (c *Crash) ID() uuid.UUID {
	return c.id
}

func (c *Crash) ContractID() uuid.UUID {
	return c.contractID
}

// реализация этих функций - зависит от структуры Crash
// func (c *Crash) ClientID() uuid.UUID {
// 	return
// }

// func (c *Crash) CarID() uuid.UUID {
// 	return
// }

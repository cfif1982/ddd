package managers

import "github.com/google/uuid"

type Manager struct {
	id    uuid.UUID
	name  string
	email string
	bonus int
}

func NewManager(id uuid.UUID, name, email string, bonus int) (*Manager, error) {

	// проверки полей

	return &Manager{
		id:    id,
		name:  name,
		email: email,
		bonus: bonus,
	}, nil
}

func CreateManager(name, email string) (*Manager, error) {

	// у вновь создаваемого менеджера бонус равен 0
	bonus := 0

	return NewManager(uuid.New(), name, email, bonus)
}

func (m *Manager) SetBonus(bonus int) {
	m.bonus = bonus
}

func (m *Manager) Bonus() int {
	return m.bonus
}

func (m *Manager) ID() uuid.UUID {
	return m.id
}

func (m *Manager) Name() string {
	return m.name
}

func (m *Manager) Email() string {
	return m.email
}

// функции проверки полей

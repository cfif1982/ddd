package clients

import (
	"errors"

	"github.com/google/uuid"
)

type Client struct {
	id       uuid.UUID
	name     string
	email    string
	discount int
}

func NewClient(id uuid.UUID, name, email string, discount int) (*Client, error) {

	// проверки полей

	return &Client{
		id:       id,
		name:     name,
		email:    email,
		discount: discount,
	}, nil
}

func CretaeClient(name, email string) (*Client, error) {

	// у вновь создаваемого клиента скидка равна 0
	discount := 0

	return NewClient(uuid.New(), name, email, discount)
}

func (c *Client) SetDiscount(discount int) {
	c.discount = discount
}

func (c *Client) SendEmail(_ string) error {
	return errors.New("not implemented")
}

func (c *Client) Discount() int {
	return c.discount
}

func (c *Client) ID() uuid.UUID {
	return c.id
}

func (c *Client) Name() string {
	return c.name
}

func (c *Client) Email() string {
	return c.email
}

// функции проверки полей

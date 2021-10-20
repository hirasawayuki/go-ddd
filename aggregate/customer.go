package aggregate

import (
	"errors"

	"github.com/google/uuid"
	"github.com/hirasawayuki/go-ddd/entity"
	"github.com/hirasawayuki/go-ddd/value_object"
)

var (
	ErrInvalidPerson = errors.New("a customer has to have an valid person")
)

type Customer struct {
	person       *entity.Person
	products     []*entity.Item
	transactions []value_object.Transaction
}

func NewCustomer(name string) (Customer, error) {
	if name == "" {
		return Customer{}, ErrInvalidPerson
	}

	person := &entity.Person{
		Name: name,
		ID:   uuid.New(),
	}

	return Customer{
		person:       person,
		products:     make([]*entity.Item, 0),
		transactions: make([]value_object.Transaction, 0),
	}, nil
}

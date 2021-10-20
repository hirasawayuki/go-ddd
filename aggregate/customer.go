package aggregate

import (
	"github.com/hirasawayuki/go-ddd/entity"
	"github.com/hirasawayuki/go-ddd/value_object"
)

type Customer struct {
	person       *entity.Person
	products     []*entity.Item
	transactions []value_object.Transaction
}

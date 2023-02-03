package repo

import (
	"nats-store/internal/entity"
)

type cache struct {
	data []entity.Order
}

func (c *cache) check(ids []int) bool {
	for _, id := range ids {
		if id >= len(c.data) {
			return false
		}
	}

	return true
}

func (c *cache) create(orders ...entity.Order) {
	c.data = append(c.data, orders...)
}

func (c *cache) get(ids []int) ([]entity.Order, error) {
	if !c.check(ids) {
		return nil, ErrOrdersNotFound
	}

	res := make([]entity.Order, 0, len(ids))

	for _, id := range ids {
		res = append(res, c.data[id])
	}

	return res, nil
}

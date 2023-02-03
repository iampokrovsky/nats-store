package usecase

import "nats-store/internal/entity"

type Order interface {
	Create(order *entity.Order) error
	Get(id int) (*entity.Order, error)
}

type OrderRepo interface {
	Create(order *entity.Order) error
	Get(ids ...int) ([]entity.Order, error)
}

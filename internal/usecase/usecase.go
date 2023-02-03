package usecase

import (
	"nats-store/internal/entity"
)

type UseCase struct {
	repo OrderRepo
}

func New(repo OrderRepo) *UseCase {
	return &UseCase{repo}
}

func (uc *UseCase) Create(order *entity.Order) error {
	return uc.repo.Create(order)
}

func (uc *UseCase) Get(id int) (*entity.Order, error) {
	orders, err := uc.repo.Get(id)

	return &orders[0], err
}

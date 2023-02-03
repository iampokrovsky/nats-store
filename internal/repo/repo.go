package repo

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"log"
	"nats-store/internal/entity"
)

const (
	_defaultCapacity = 10
)

var (
	ErrOrdersNotFound = errors.New("one or more orders not found")
)

type Repo struct {
	cache
	psql
}

func New(db *sqlx.DB) *Repo {
	r := &Repo{
		cache{make([]entity.Order, 0, _defaultCapacity)},
		psql{db},
	}

	if len(r.cache.data) == 0 {
		if err := r.recoverCache(); err != nil {
			log.Fatal(err)
		}
	}

	return r
}

func (r *Repo) Create(order *entity.Order) error {
	if err := r.psql.create(*order); err != nil {
		fmt.Println("psql error")
		return err
	}

	r.cache.create(*order)

	return nil
}

func (r *Repo) Get(ids ...int) ([]entity.Order, error) {
	return r.cache.get(ids)
}

func (r *Repo) recoverCache() error {
	orders, err := r.psql.getAll()
	if err != nil {
		return err
	}

	r.cache.create(orders...)

	return nil
}

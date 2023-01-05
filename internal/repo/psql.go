package repo

import (
	"encoding/json"
	"github.com/jmoiron/sqlx"
	"github.com/pokrovsky-io/nats-store/internal/entity"
)

type psql struct {
	db *sqlx.DB
}

// TODO: Предотвратить тротлинг, накапливать значения до публикации
func (psql *psql) create(order entity.Order) error {
	var id int

	encOrder, err := json.Marshal(order)
	if err != nil {
		return err
	}

	row := psql.db.QueryRowx("INSERT INTO orders (data) VALUES ($1) RETURNING id", string(encOrder))

	return row.Scan(&id)
}

func (psql *psql) getAll() ([]entity.Order, error) {
	res := make([]entity.Order, 0, _defaultCapacity)

	rows, err := psql.db.Query("SELECT data FROM orders")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var order entity.Order
		var encOrder []byte

		if err = rows.Scan(&encOrder); err != nil {
			return nil, err
		}

		if err = json.Unmarshal(encOrder, &order); err != nil {
			return nil, err
		}

		res = append(res, order)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return res, nil
}

//// TODO Реализовать метод
//func (psql *psql) get(ids ...int) ([]entity.Order, error) {
//	res := make([]entity.Order, 0, len(ids))
//
//	if len(ids) == 0 {
//		if err := psql.db.Select(&res, "SELECT data FROM orders"); err != nil {
//			return nil, err
//		}
//
//		return nil, nil
//	}
//
//	query, args, err := sqlx.In("SELECT data FROM orders WHERE id IN ($1);", ids)
//	if err != nil {
//		return nil, err
//	}
//
//	query = psql.db.Rebind(query)
//	rows, err := psql.db.Query(query, args...)
//
//	for rows.Next() {
//		var order entity.Order
//		if err = rows.Scan(&order); err != nil {
//			return nil, err
//		}
//
//		res = append(res, order)
//	}
//
//	if err = rows.Err(); err != nil {
//		return nil, err
//	}
//
//	return res, nil
//}
